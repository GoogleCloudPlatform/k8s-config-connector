package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkSecurityGatewaySecurityPolicy_update(t *testing.T) {
	t.Parallel()

	gatewaySecurityPolicyName := fmt.Sprintf("tf-test-gateway-sp-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityGatewaySecurityPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityGatewaySecurityPolicy_basic(gatewaySecurityPolicyName),
			},
			{
				ResourceName:      "google_network_security_gateway_security_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityGatewaySecurityPolicy_update(gatewaySecurityPolicyName),
			},
			{
				ResourceName:      "google_network_security_gateway_security_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityGatewaySecurityPolicy_basic(gatewaySecurityPolicyName string) string {
	return fmt.Sprintf(`
resource "google_network_security_gateway_security_policy" "foobar" {
  name        = "%s"
  location    = "us-central1"
  description = "my description"
}
`, gatewaySecurityPolicyName)
}

func testAccNetworkSecurityGatewaySecurityPolicy_update(gatewaySecurityPolicyName string) string {
	return fmt.Sprintf(`
resource "google_network_security_gateway_security_policy" "foobar" {
  name        = "%s"
  location    = "us-central1"
  description = "update description"
}
`, gatewaySecurityPolicyName)
}
