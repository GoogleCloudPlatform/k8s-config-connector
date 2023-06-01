// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesTcpRoute_update(t *testing.T) {
	t.Parallel()

	tcpServiceName := fmt.Sprintf("tf-test-tcp-service-%s", RandString(t, 10))
	tcpHealthCheckName := fmt.Sprintf("tf-test-tcp-healthcheck-%s", RandString(t, 10))
	tcpRouteName := fmt.Sprintf("tf-test-tcp-route-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTcpRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTcpRoute_basic(tcpServiceName, tcpHealthCheckName, tcpRouteName),
			},
			{
				ResourceName:      "google_network_services_tcp_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesTcpRoute_update(tcpServiceName, tcpHealthCheckName, tcpRouteName),
			},
			{
				ResourceName:      "google_network_services_tcp_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesTcpRoute_basic(tcpServiceName string, tcpHealthCheckName string, tcpRouteName string) string {
	return fmt.Sprintf(`
resource "google_compute_backend_service" "foo" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.bar.id]
}

resource "google_compute_http_health_check" "bar" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_tcp_route" "foobar" { 
  name                   = "%s"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  rules                   {
    matches {
      address = "10.0.0.1/32"
      port = "8081"
    }
    action {
      destinations {
        service_name = google_compute_backend_service.foo.id
        weight = 1
      }
      original_destination = false
    }
  }
}
`, tcpServiceName, tcpHealthCheckName, tcpRouteName)
}

func testAccNetworkServicesTcpRoute_update(tcpServiceName string, tcpHealthCheckName string, tcpRouteName string) string {
	return fmt.Sprintf(`
  resource "google_compute_backend_service" "foo" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.bar.id]
}

resource "google_compute_http_health_check" "bar" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_tcp_route" "foobar" {
  name                   = "%s"
  labels                 = {
    foo = "bar"
  }
  description             = "update description"
  rules                   {
    matches {
      address = "10.0.0.1/32"
      port = "8081"
    }
    action {
      destinations {
        service_name = google_compute_backend_service.foo.id
        weight = 1
      }
      original_destination = false
    }
  }
}
`, tcpServiceName, tcpHealthCheckName, tcpRouteName)
}
