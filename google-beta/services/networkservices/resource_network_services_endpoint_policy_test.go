// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package networkservices_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesEndpointPolicy_update(t *testing.T) {
	t.Parallel()

	endpointPolicyName := fmt.Sprintf("tf-test-endpoint-policy-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEndpointPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEndpointPolicy_basic(endpointPolicyName),
			},
			{
				ResourceName:      "google_network_services_endpoint_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesEndpointPolicy_update(endpointPolicyName),
			},
			{
				ResourceName:      "google_network_services_endpoint_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesEndpointPolicy_basic(endpointPolicyName string) string {
	return fmt.Sprintf(`
resource "google_network_services_endpoint_policy" "foobar" {
  name                   = "%s"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  type                   = "SIDECAR_PROXY"
  traffic_port_selector {
    ports = ["8081"]
  }
  endpoint_matcher {
    metadata_label_matcher {
      metadata_label_match_criteria = "MATCH_ANY"
      metadata_labels {
        label_name = "foo"
        label_value = "bar"
      }
    }
  }
}
`, endpointPolicyName)
}

func testAccNetworkServicesEndpointPolicy_update(endpointPolicyName string) string {
	return fmt.Sprintf(`
resource "google_network_services_endpoint_policy" "foobar" {
  name        = "%s"
  labels                 = {
    foo = "barbar"
    baz = "qux"
  }
  description            = "update description"
  type                   = "GRPC_SERVER"
  endpoint_matcher {
    metadata_label_matcher {
      metadata_label_match_criteria = "MATCH_ALL"
      metadata_labels {
        label_name = "baz"
        label_value = "bux"
      }
    }
  } 
}
`, endpointPolicyName)
}
