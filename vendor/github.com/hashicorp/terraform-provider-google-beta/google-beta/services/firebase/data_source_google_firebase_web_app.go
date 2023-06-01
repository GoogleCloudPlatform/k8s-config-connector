// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleFirebaseWebApp() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceFirebaseWebApp().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "app_id")

	// Allow specifying a project
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleFirebaseWebAppRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleFirebaseWebAppRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	appId := d.Get("app_id")
	project, err := tpgresource.GetProject(d, config)
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
