package main

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"gitlab.com/serg.dev/terraform-provider-vscale/models"
	"golang.org/x/net/context"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				DataSourcesMap: map[string]*schema.Resource{},
				ResourcesMap: map[string]*schema.Resource{
					"vscale_ssh_key": {
						CreateContext: CreateResource(&models.SSHKey{}),
						ReadContext:   ReadResource(&models.SSHKey{}),
						UpdateContext: UpdateResource(&models.SSHKey{}),
						DeleteContext: DeleteResource(&models.SSHKey{}),
						Schema:        models.SchemaSSHKey,
						Importer: &schema.ResourceImporter{
							State: schema.ImportStatePassthrough,
						},
					},
					"vscale_server": {
						CreateContext: CreateResource(&models.Server{}),
						ReadContext:   ReadResource(&models.Server{}),
						UpdateContext: UpdateResource(&models.Server{}),
						DeleteContext: DeleteResource(&models.Server{}),
						Schema:        models.SchemaServer,
						Importer: &schema.ResourceImporter{
							State: schema.ImportStatePassthrough,
						},
						Timeouts: &schema.ResourceTimeout{
							Create: schema.DefaultTimeout(5 * time.Minute),
						},
					},
				},
			}
		},
	})
}

func ReadResource(o models.Resource) schema.ReadContextFunc {
	return func(ctx context.Context, res *schema.ResourceData, m interface{}) diag.Diagnostics {
		diags := diag.Diagnostics{}
		obj := o.NewObj()
		obj.ReadTF(res)

		err := obj.Read()
		if err != nil {
			return diag.FromErr(err)
		}

		obj.WriteTF(res)
		return diags
	}

}

func UpdateResource(o models.Resource) schema.UpdateContextFunc {
	return func(ctx context.Context, res *schema.ResourceData, m interface{}) diag.Diagnostics {
		diags := diag.Diagnostics{}
		obj := o.NewObj()
		obj.ReadTF(res)

		obj.WriteTF(res)
		return diags
	}

}

func DeleteResource(o models.Resource) schema.DeleteContextFunc {
	return func(ctx context.Context, res *schema.ResourceData, m interface{}) diag.Diagnostics {
		diags := diag.Diagnostics{}
		obj := o.NewObj()
		obj.ReadTF(res)

		err := obj.Delete()
		if err != nil {
			return diag.FromErr(err)
		}
		res.SetId("")
		return diags
	}

}

func CreateResource(o models.Resource) schema.CreateContextFunc {
	return func(ctx context.Context, res *schema.ResourceData, m interface{}) diag.Diagnostics {
		diags := diag.Diagnostics{}
		obj := o.NewObj()
		err := obj.ReadTF(res)
		if err != nil {
			return diag.FromErr(err)
		}

		err = obj.Create()
		if err != nil {
			return diag.FromErr(err)
		}

		err = obj.Wait(res)
		if err != nil {
			return diag.FromErr(err)
		}

		err = obj.WriteTF(res)
		if err != nil {
			return diag.FromErr(err)
		}
		return diags
	}

}
