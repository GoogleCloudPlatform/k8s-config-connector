package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIAMBetaWorkloadIdentityPoolProvider() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceIAMBetaWorkloadIdentityPoolProvider().Schema)
	addRequiredFieldsToSchema(dsSchema, "workload_identity_pool_id")
	addRequiredFieldsToSchema(dsSchema, "workload_identity_pool_provider_id")
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceIAMBetaWorkloadIdentityPoolProviderRead,
		Schema: dsSchema,
	}
}

func dataSourceIAMBetaWorkloadIdentityPoolProviderRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceIAMBetaWorkloadIdentityPoolProviderRead(d, meta)

}
