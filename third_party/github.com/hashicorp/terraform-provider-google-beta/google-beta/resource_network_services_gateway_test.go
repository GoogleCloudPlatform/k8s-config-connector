package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesGateway_update(t *testing.T) {
	t.Parallel()

	gatewayName := fmt.Sprintf("tf-test-gateway-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesGateway_basic(gatewayName),
			},
			{
				ResourceName:      "google_network_services_gateway.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesGateway_update(gatewayName),
			},
			{
				ResourceName:      "google_network_services_gateway.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesGateway_basic(gatewayName string) string {
	return fmt.Sprintf(`
resource "google_network_services_gateway" "foobar" {
  name        = "%s"
  scope       = "default-scope-update"
  type        = "OPEN_MESH"
  ports       = [443]
  description = "my description"
}
`, gatewayName)
}

func testAccNetworkServicesGateway_update(gatewayName string) string {
	return fmt.Sprintf(`
resource "google_network_services_gateway" "foobar" {
  name        = "%s"
  scope       = "default-scope-update"
  type        = "OPEN_MESH"
  ports       = [443]
  description = "update description"
  labels      = {
    foo = "bar"
  } 
}
`, gatewayName)
}
