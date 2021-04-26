package models

import "encoding/json"

type Address struct {
	Address string `json:"address"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}

func (o *Address) IsEmpty() bool {
	if o.Address == "" && o.Gateway == "" && o.Netmask == "" {
		return true
	}
	return false
}

func (o *Address) ToMap() (map[string]string, error) {
	objBytes, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	err = json.Unmarshal(objBytes, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
