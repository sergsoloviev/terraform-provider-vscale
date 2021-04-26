package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var (
	SchemaSSHKey map[string]*schema.Schema
	SchemaServer map[string]*schema.Schema
)

func init() {
	SchemaSSHKey = map[string]*schema.Schema{
		"id":   {Type: schema.TypeString, Computed: true},
		"name": {Type: schema.TypeString, Optional: true},
		"key":  {Type: schema.TypeString, Optional: true},
	}

	SchemaServer = map[string]*schema.Schema{
		"id":       {Type: schema.TypeString, Computed: true},
		"name":     {Type: schema.TypeString, Required: true},
		"location": {Type: schema.TypeString, Required: true},
		"image":    {Type: schema.TypeString, Required: true},
		"rplan":    {Type: schema.TypeString, Required: true},
		"hostname": {Type: schema.TypeString, Computed: true},
		"status":   {Type: schema.TypeString, Computed: true},
		"active":   {Type: schema.TypeString, Computed: true},
		"locked":   {Type: schema.TypeString, Computed: true},
		"keys": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"private_address": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"public_address": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
