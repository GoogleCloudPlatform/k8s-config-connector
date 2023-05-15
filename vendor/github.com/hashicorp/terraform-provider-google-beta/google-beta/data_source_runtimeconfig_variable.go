package google

import (
	"fmt"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleRuntimeconfigVariable() *schema.Resource {

	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceRuntimeconfigVariable().Schema)
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "name")
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "parent")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleRuntimeconfigVariableRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleRuntimeconfigVariableRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/configs/{{parent}}/variables/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceRuntimeconfigVariableRead(d, meta)
}
