// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package networkservices_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesHttpRoute_update(t *testing.T) {
	t.Parallel()

	httpRouteName := fmt.Sprintf("tf-test-http-route-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesHttpRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesHttpRoute_basic(httpRouteName),
			},
			{
				ResourceName:      "google_network_services_http_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesHttpRoute_update(httpRouteName),
			},
			{
				ResourceName:      "google_network_services_http_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesHttpRoute_basic(httpRouteName string) string {
	return fmt.Sprintf(`
resource "google_network_services_http_route" "foobar" {
  name                   = "%s"
  description             = "my description"
  hostnames               = ["example"]
  rules                   {
    matches {
      query_parameters {
        query_parameter = "key"
        exact_match = "value"
      }
      full_path_match = "example"
    }
  }
}
`, httpRouteName)
}

func testAccNetworkServicesHttpRoute_update(httpRouteName string) string {
	return fmt.Sprintf(`
resource "google_network_services_http_route" "foobar" {
  name        = "%s"
  description = "update description"
  labels      = {
    foo = "bar"
  }
  hostnames               = ["example"]
  rules                   {
    matches {
      query_parameters {
        query_parameter = "key"
        exact_match = "value"
      }
      full_path_match = "example"
    }
  } 
}
`, httpRouteName)
}
