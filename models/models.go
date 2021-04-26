package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Resource interface {
	NewObj() Resource
	ReadTF(*schema.ResourceData)
	WriteTF(*schema.ResourceData)
	Create() error
	Read() error
	Update() error
	Delete() error
}
