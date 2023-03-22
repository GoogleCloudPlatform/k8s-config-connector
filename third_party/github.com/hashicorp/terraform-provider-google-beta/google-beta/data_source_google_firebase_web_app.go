package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleFirebaseWebApp() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(ResourceFirebaseWebApp().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "app_id")

	// Allow specifying a project
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleFirebaseWebAppRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleFirebaseWebAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	appId := d.Get("app_id")
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	name := fmt.Sprintf("projects/%s/webApps/%s", project, appId.(string))
	d.SetId(name)
	if err := d.Set("name", name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	return resourceFirebaseWebAppRead(d, meta)
}
