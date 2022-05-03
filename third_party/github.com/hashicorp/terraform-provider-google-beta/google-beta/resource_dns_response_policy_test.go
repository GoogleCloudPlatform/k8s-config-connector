package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDNSResponsePolicy_update(t *testing.T) {
	t.Parallel()

	responsePolicySuffix := randString(t, 10)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDNSResponsePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDnsResponsePolicy_privateUpdate(responsePolicySuffix, "network-1"),
			},
			{
				ResourceName:      "google_dns_response_policy.example-response-policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsResponsePolicy_privateUpdate(responsePolicySuffix, "network-2"),
			},
			{
				ResourceName:      "google_dns_response_policy.example-response-policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDnsResponsePolicy_privateUpdate(suffix, network string) string {
	return fmt.Sprintf(`
resource "google_dns_response_policy" "example-response-policy" {
  provider = google-beta

  response_policy_name = "tf-test-response-policy-%s"

  networks {
    network_url = google_compute_network.%s.self_link
  }
}

resource "google_compute_network" "network-1" {
  provider = google-beta

  name                    = "tf-test-network-1-%s"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  provider = google-beta
	
  name                    = "tf-test-network-2-%s"
  auto_create_subnetworks = false
}
`, suffix, network, suffix, suffix)
}
