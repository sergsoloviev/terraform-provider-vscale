package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/serg.dev/terraform-provider-vscale/network"
)

type Server struct {
	Id             int      `json:"ctid"`
	Name           string   `json:"name"`
	Hostname       string   `json:"hostname"`
	Location       string   `json:"location"`
	Rplan          string   `json:"rplan"`
	Image          string   `json:"made_from"`
	PublicAddress  *Address `json:"public_address,omitempty"`
	PrivateAddress *Address `json:"private_address,omitempty"`
	Keys           []*SSHKey
	Status         string `json:"status"`
	Active         bool   `json:"active"`
	Locked         bool   `json:"locked"`
}

func (o *Server) NewObj() Resource {
	return &Server{}
}

func (o *Server) ReadTF(res *schema.ResourceData) (err error) {
	idString := res.Get("id").(string)
	if idString != "" {
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Println(err)
		} else {
			o.Id = id
		}
	}
	o.Name = res.Get("name").(string)
	o.Image = res.Get("image").(string)
	o.Location = res.Get("location").(string)
	o.Rplan = res.Get("rplan").(string)
	o.Hostname = res.Get("hostname").(string)
	o.Status = res.Get("status").(string)
	o.Active, err = strconv.ParseBool(res.Get("active").(string))
	if err != nil {
		log.Println(err)
	}
	o.Locked, err = strconv.ParseBool(res.Get("locked").(string))
	if err != nil {
		log.Println(err)
	}

	err = o.getKeys(res.Get("keys"))
	if err != nil {
		return err
	}
	return nil
}

func (o *Server) getKeys(keynames interface{}) error {
	nw := network.Network{Method: "GET", Url: "/sshkeys"}
	err := nw.Do()
	if err != nil {
		return err
	}
	if nw.Response.StatusCode != 200 {
		return fmt.Errorf("%s", nw.Response.Header["Vscale-Error-Message"])
	}
	if o.Keys == nil {
		o.Keys = make([]*SSHKey, 0)
	}
	allkeys := make([]*SSHKey, 0)
	err = json.Unmarshal(nw.ResponseBody, &allkeys)
	if err != nil {
		return err
	}
	for _, v := range allkeys {
		for _, name := range keynames.([]interface{}) {
			if v.Name == name {
				o.Keys = append(o.Keys, v)
			}
		}
	}
	return nil
}

func (o *Server) WriteTF(res *schema.ResourceData) error {
	res.SetId(strconv.Itoa(o.Id))
	res.Set("name", o.Name)
	res.Set("image", o.Image)
	res.Set("location", o.Location)
	res.Set("rplan", o.Rplan)
	res.Set("hostname", o.Hostname)
	res.Set("status", o.Status)
	res.Set("active", strconv.FormatBool(o.Active))
	res.Set("locked", strconv.FormatBool(o.Locked))

	pubAddr, err := o.PublicAddress.ToMap()
	if err != nil {
		return err
	}
	err = res.Set("public_address", pubAddr)
	if err != nil {
		return err
	}
	privAddr, err := o.PrivateAddress.ToMap()
	if err != nil {
		return err
	}
	err = res.Set("private_address", privAddr)
	if err != nil {
		return err
	}
	//res.Set("", o)
	//res.Set("", o)

	return nil
}

func (o *Server) Create() error {
	requestMap := map[string]interface{}{
		"do_start":  true,
		"make_from": o.Image,
		"rplan":     o.Rplan,
		"name":      o.Name,
		"location":  o.Location,
	}

	if len(o.Keys) > 0 {
		keyList := make([]int, len(o.Keys))
		for k, v := range o.Keys {
			keyList[k] = v.Id
		}
		requestMap["keys"] = keyList
	}

	requestBody, err := json.Marshal(requestMap)
	if err != nil {
		return err
	}

	nw := network.Network{
		Url:         "/scalets",
		Method:      "POST",
		RequestBody: requestBody,
		Debug:       true,
	}

	err = nw.Do()
	if err != nil {
		return err
	}

	if nw.Response.StatusCode != 201 {
		return fmt.Errorf("%s", nw.Response.Header["Vscale-Error-Message"])
	}

	err = json.Unmarshal(nw.ResponseBody, &o)
	if err != nil {
		return err
	}

	return nil
}

func (o *Server) Read() error {
	nw := network.Network{Method: "GET", Url: fmt.Sprintf("/scalets/%d", o.Id)}
	err := nw.Do()
	if err != nil {
		return err
	}

	if nw.Response.StatusCode != 200 {
		return fmt.Errorf("%s", nw.Response.Header["Vscale-Error-Message"])
	}

	err = json.Unmarshal(nw.ResponseBody, &o)
	if err != nil {
		return err
	}

	if o.PrivateAddress.IsEmpty() {
		o.PrivateAddress = nil
	}
	if o.PublicAddress.IsEmpty() {
		o.PublicAddress = nil
	}
	//for _, v := range keys {
	//}
	return nil
}

func (o *Server) Update() error {
	return nil
}

func (o *Server) Delete() error {
	nw := network.Network{
		Method: "DELETE",
		Url:    fmt.Sprintf("/scalets/%d", o.Id),
	}
	err := nw.Do()
	if err != nil {
		return err
	}
	// not 202||204?
	if nw.Response.StatusCode != 200 {
		return fmt.Errorf("%s", nw.Response.Header["Vscale-Error-Message"])
	}
	return nil
}

func (o *Server) Wait(res *schema.ResourceData) error {
	stateChangeConf := &resource.StateChangeConf{
		Pending:                   []string{"queued"},
		Target:                    []string{"started"},
		Timeout:                   res.Timeout(schema.TimeoutCreate),
		Delay:                     10 * time.Second,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 3,
		Refresh: func() (interface{}, string, error) {
			err := o.Read()
			if err != nil {
				return 0, o.Status, err
			}
			return 1, o.Status, nil
		},
	}
	_, err := stateChangeConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for example instance (%s) to be created: %s", res.Id(), err)
	}
	return nil
}
