package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var (
	SchemaSSHKey map[string]*schema.Schema
)

func init() {
	SchemaSSHKey = map[string]*schema.Schema{
		"id":   {Type: schema.TypeString, Computed: true},
		"name": {Type: schema.TypeString, Optional: true},
		"key":  {Type: schema.TypeString, Optional: true},
	}
}
