package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeOrganizationSecurityPolicy_organizationSecurityPolicyUpdateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeOrganizationSecurityPolicyDestroyProducer(t),
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
