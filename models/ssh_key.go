package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/serg.dev/terraform-provider-vscale/network"
)

type SSHKey struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

func (o *SSHKey) NewObj() Resource {
	return &SSHKey{}
}

func (o *SSHKey) ReadTF(res *schema.ResourceData) error {
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
	o.Key = res.Get("key").(string)
	return nil
}

func (o *SSHKey) WriteTF(res *schema.ResourceData) error {
	res.SetId(strconv.Itoa(o.Id))
	res.Set("name", o.Name)
	res.Set("key", o.Key)
	return nil
}

func (o *SSHKey) Create() error {
	requestMap := map[string]string{
		"name": o.Name,
		"key":  o.Key,
	}

	requestBody, err := json.Marshal(requestMap)
	if err != nil {
		return err
	}

	nw := network.Network{
		Url:         "/sshkeys",
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

func (o *SSHKey) Read() error {
	nw := network.Network{Method: "GET", Url: "/sshkeys"}
	err := nw.Do()
	if err != nil {
		return err
	}

	if nw.Response.StatusCode != 200 {
		return fmt.Errorf("%s", nw.Response.Header["Vscale-Error-Message"])
	}

	keys := make([]*SSHKey, 0)
	err = json.Unmarshal(nw.ResponseBody, &keys)
	if err != nil {
		return err
	}

	for _, v := range keys {
		if o.Id == v.Id {
			o.Name = v.Name
			o.Key = v.Key
		}
	}

	return nil
}

func (o *SSHKey) Update() error {
	return fmt.Errorf("not implemented in api")
}

func (o *SSHKey) Delete() error {
	nw := network.Network{
		Method: "DELETE",
		Url:    fmt.Sprintf("/sshkeys/%d", o.Id),
	}
	err := nw.Do()
	if err != nil {
		return err
	}
	if nw.Response.StatusCode != 204 {
		return fmt.Errorf("%s", nw.Response.Header["Vscale-Error-Message"])
	}
	return nil
}
