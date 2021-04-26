package network

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Network struct {
	ApiUrl       string
	Url          string
	Request      *http.Request
	Response     *http.Response
	Client       *http.Client
	Method       string
	RequestBody  []byte
	ResponseBody []byte
	Token        string
	TimeoutSec   int
	Debug        bool
}

func (o *Network) setEnv() {
	token, ok := os.LookupEnv("VSCALE_TOKEN")
	if ok {
		o.Token = token
	}
	url, ok := os.LookupEnv("VSCALE_URL")
	if ok {
		o.Url = fmt.Sprintf("%s%s", url, o.Url)
	}
	logMode, ok := os.LookupEnv("TF_LOG")
	if ok && logMode == "DEBUG" {
		o.Debug = true
	}

}

func (o *Network) Do() (err error) {
	if o.TimeoutSec == 0 {
		o.TimeoutSec = 5
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(o.TimeoutSec)*time.Second)
	defer cancel()

	o.setEnv()

	o.Request, err = http.NewRequestWithContext(ctx, o.Method, o.Url, bytes.NewBuffer(o.RequestBody))
	if err != nil {
		return err
	}
	o.Request.Header.Set("Content-Type", "application/json;charset=UTF-8'")
	o.Request.Header.Set("X-Token", o.Token)
	o.Client = &http.Client{}
	if o.Debug {
		bb := o.jsonLog(o.RequestBody)
		log.Printf("API Request: %v", bb.String())
	}
	o.Response, err = o.Client.Do(o.Request)
	if err != nil {
		return err
	}
	o.ResponseBody, err = ioutil.ReadAll(o.Response.Body)
	if err != nil {
		return err
	}
	err = o.Response.Body.Close()
	if err != nil {
		return err
	}
	if o.Debug {
		bb := o.jsonLog(o.ResponseBody)
		log.Printf("API Response [%d]: %v", o.Response.StatusCode, bb.String())
	}
	return nil
}

func (o *Network) jsonLog(data []byte) *bytes.Buffer {
	bb := &bytes.Buffer{}
	err := json.Indent(bb, data, "", "\t")
	if err != nil {
		log.Println(err)
		return nil
	}
	return bb
}
