package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Resource interface {
	NewObj() Resource
	ReadTF(*schema.ResourceData) error
	WriteTF(*schema.ResourceData) error
	Create() error
	Read() error
	Update() error
	Delete() error
	Wait(*schema.ResourceData) error
}
