// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesTlsRoute_update(t *testing.T) {
	t.Parallel()

	tlsServiceName := fmt.Sprintf("tf-test-tls-service-%s", acctest.RandString(t, 10))
	tlsHealthCheckName := fmt.Sprintf("tf-test-tls-healthcheck-%s", acctest.RandString(t, 10))
	tlsRouteName := fmt.Sprintf("tf-test-tls-route-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesTlsRouteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesTlsRoute_basic(tlsServiceName, tlsHealthCheckName, tlsRouteName),
			},
			{
				ResourceName:      "google_network_services_tls_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkServicesTlsRoute_update(tlsServiceName, tlsHealthCheckName, tlsRouteName),
			},
			{
				ResourceName:      "google_network_services_tls_route.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesTlsRoute_basic(tlsServiceName string, tlsHealthCheckName string, tlsRouteName string) string {
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

resource "google_network_services_tls_route" "foobar" { 
  name                   = "%s"
  description             = "my description"
  rules                   {
    matches {
      sni_host = ["example.com"]
      alpn = ["http/1.1"]
    }
    action {
      destinations {
        service_name = google_compute_backend_service.foo.id
        weight = 1
      }
    }
  }
}
`, tlsServiceName, tlsHealthCheckName, tlsRouteName)
}

func testAccNetworkServicesTlsRoute_update(tlsServiceName string, tlsHealthCheckName string, tlsRouteName string) string {
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

  resource "google_network_services_tls_route" "foobar" {
    name                   = "%s"
    description             = "update description"
    rules                   {
      matches {
        sni_host = ["example.com"]
        alpn = ["http/1.1"]
      }
      action {
        destinations {
          service_name = google_compute_backend_service.foo.id
          weight = 1
        }
      }
    }
  }
`, tlsServiceName, tlsHealthCheckName, tlsRouteName)
}
