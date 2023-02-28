package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleFirebaseHostingChannel() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(ResourceFirebaseHostingChannel().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "site_id", "channel_id")

	return &schema.Resource{
		Read:   dataSourceGoogleFirebaseHostingChannelRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleFirebaseHostingChannelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, err := replaceVars(d, config, "sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceFirebaseHostingChannelRead(d, meta)
}
