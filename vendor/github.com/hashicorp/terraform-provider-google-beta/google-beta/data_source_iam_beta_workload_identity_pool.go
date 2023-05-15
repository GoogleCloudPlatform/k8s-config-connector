package google

import (
	"fmt"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIAMBetaWorkloadIdentityPool() *schema.Resource {

	dsSchema := tpgresource.DatasourceSchemaFromResourceSchema(ResourceIAMBetaWorkloadIdentityPool().Schema)
	tpgresource.AddRequiredFieldsToSchema(dsSchema, "workload_identity_pool_id")
	tpgresource.AddOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceIAMBetaWorkloadIdentityPoolRead,
		Schema: dsSchema,
	}
}

func dataSourceIAMBetaWorkloadIdentityPoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceIAMBetaWorkloadIdentityPoolRead(d, meta)

}
