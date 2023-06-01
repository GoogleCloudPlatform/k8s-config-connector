// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyUpdateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        acctest.GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeOrganizationSecurityPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyPreUpdateExample(context),
			},
			{
				ResourceName:      "google_compute_organization_security_policy.policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyPostUpdateExample(context),
			},
			{
				ResourceName:      "google_compute_organization_security_policy.policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyPreUpdateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
}
`, context)
}

func testAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyPostUpdateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
  description  = "Updated description."
}
`, context)
}
