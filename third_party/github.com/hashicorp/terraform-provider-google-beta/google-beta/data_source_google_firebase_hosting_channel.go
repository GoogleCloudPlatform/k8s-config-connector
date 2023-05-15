package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func DataSourceGoogleFirebaseHostingChannel() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceFirebaseHostingChannel().Schema)

	// Set 'Required' schema elements
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "site_id", "channel_id")

	return &schema.Resource{
		Read:   dataSourceGoogleFirebaseHostingChannelRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleFirebaseHostingChannelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	id, err := tpgresource.ReplaceVars(d, config, "sites/{{site_id}}/channels/{{channel_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceFirebaseHostingChannelRead(d, meta)
}
