// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package networkservices_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNetworkServicesGrpcRoute_update(t *testing.T) {
	t.Parallel()

	grpcRouteName := fmt.Sprintf("tf-test-grpc-route-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesGrpcRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesGrpcRoute_basic(grpcRouteName),
			},
			{
				ResourceName:      "google_network_services_grpc_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesGrpcRoute_update(grpcRouteName),
			},
			{
				ResourceName:      "google_network_services_grpc_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesGrpcRoute_basic(grpcRouteName string) string {
	return fmt.Sprintf(`
  resource "google_network_services_grpc_route" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description             = "my description"
    hostnames               = ["example"]
    rules                   {
      matches {
        headers {
          key = "key"
          value = "value"
        }
      }
      action {
        retry_policy {
            retry_conditions = ["cancelled"]
            num_retries = 1
        }
      }
    }
    rules                   {
      matches {
        headers {
          key = "key"
          value = "value"
        }
      }
      action {
        fault_injection_policy {
          delay {
            fixed_delay = "1s"
            percentage = 1
          }
          abort {
            http_status = 500
            percentage = 1
          }
        }
      }
    }
  }
`, grpcRouteName)
}

func testAccNetworkServicesGrpcRoute_update(grpcRouteName string) string {
	return fmt.Sprintf(`
  resource "google_network_services_grpc_route" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description             = "updated description"
    hostnames               = ["example"]
    rules                   {
      matches {
        headers {
          key = "key"
          value = "value"
        }
      }
      action {
        retry_policy {
            retry_conditions = ["cancelled"]
            num_retries = 2
        }
      }
    }
    rules                   {
      matches {
        headers {
          key = "key1"
          value = "value1"
        }
      }
      action {
        retry_policy {
            retry_conditions = ["connect-failure"]
            num_retries = 1
        }
      }
    }
  }
`, grpcRouteName)
}
