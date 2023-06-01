// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIAMBetaWorkloadIdentityPool_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIAMBetaWorkloadIdentityPoolBasic(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_iam_workload_identity_pool.foo", "google_iam_workload_identity_pool.bar"),
				),
			},
		},
	})
}

func testAccDataSourceIAMBetaWorkloadIdentityPoolBasic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_iam_workload_identity_pool" "bar" {
	workload_identity_pool_id = "bar-pool-%{random_suffix}"
	display_name              = "Name of pool"
	description               = "Identity pool for automated test"
	disabled                  = true
}

data "google_iam_workload_identity_pool" "foo" {
	workload_identity_pool_id = google_iam_workload_identity_pool.bar.workload_identity_pool_id
}
`, context)
}
