package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGoogleRuntimeconfigVariable() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(resourceRuntimeconfigVariable().Schema)
	addRequiredFieldsToSchema(dsSchema, "name")
	addRequiredFieldsToSchema(dsSchema, "parent")
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleRuntimeconfigVariableRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleRuntimeconfigVariableRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, err := replaceVars(d, config, "projects/{{project}}/configs/{{parent}}/variables/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceRuntimeconfigVariableRead(d, meta)
}
