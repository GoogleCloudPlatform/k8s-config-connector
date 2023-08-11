// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkServicesServiceBinding_update(t *testing.T) {
	t.Parallel()

	serviceNamespace := fmt.Sprintf("tf-test-service-namespace-%s", acctest.RandString(t, 10))
	serviceName := fmt.Sprintf("tf-test-service-%s", acctest.RandString(t, 10))
	serviceBindingName := fmt.Sprintf("tf-test-service-binding-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesServiceBindingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesServiceBinding_create(serviceNamespace, serviceName, serviceBindingName),
			},
			{
				ResourceName:      "google_network_services_service_binding.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkServicesServiceBinding_create(serviceNamespace string, serviceName string, serviceBindingName string) string {
	return fmt.Sprintf(`

	resource "google_service_directory_namespace" "foo" {
    namespace_id = "%s"
    location     = "us-central1"
  }
  resource "google_service_directory_service" "bar" {
    service_id = "%s"
    namespace  = google_service_directory_namespace.foo.id

    metadata = {
      stage  = "prod"
      region = "us-central1"
    }
  }
  resource "google_network_services_service_binding" "foobar" {
    name        = "%s"
    description = "my description"
	  service = google_service_directory_service.bar.id
  }
`, serviceNamespace, serviceName, serviceBindingName)
}
