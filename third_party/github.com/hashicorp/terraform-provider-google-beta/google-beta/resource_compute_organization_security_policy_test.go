// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyUpdateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
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
	return acctest.Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
}
`, context)
}

func testAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyPostUpdateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_organization_security_policy" "policy" {
  display_name = "tf-test%{random_suffix}"
  parent       = "organizations/%{org_id}"
  description  = "Updated description."
}
`, context)
}
