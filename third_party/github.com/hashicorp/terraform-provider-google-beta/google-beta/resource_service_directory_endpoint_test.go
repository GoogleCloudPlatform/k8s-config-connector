// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServiceDirectoryEndpoint_serviceDirectoryEndpointUpdateExample(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	location := "us-central1"
	testId := fmt.Sprintf("tf-test-example-endpoint%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckServiceDirectoryEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryEndpoint_basic(location, testId),
			},
			{
				ResourceName:      "google_service_directory_endpoint.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_endpoint.example",
				// {{project}}/{{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}
				ImportStateId:     fmt.Sprintf("%s/%s/%s/%s/%s", project, location, testId, testId, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_endpoint.example",
				// {{location}}/{{namespace_id}}/{{service_id}}/{{endpoint_id}}
				ImportStateId:     fmt.Sprintf("%s/%s/%s/%s", location, testId, testId, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccServiceDirectoryEndpoint_update(location, testId),
			},
			{
				ResourceName:      "google_service_directory_endpoint.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccServiceDirectoryEndpoint_basic(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"
}

resource "google_service_directory_service" "example" {
  service_id = "%s"
  namespace  = google_service_directory_namespace.example.id
}

resource "google_service_directory_endpoint" "example" {
  endpoint_id = "%s"
  service     = google_service_directory_service.example.id
}
`, testId, location, testId, testId)
}

func testAccServiceDirectoryEndpoint_update(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"
}

resource "google_service_directory_service" "example" {
  service_id = "%s"
  namespace  = google_service_directory_namespace.example.id
}

resource "google_service_directory_endpoint" "example" {
  endpoint_id = "%s"
  service     = google_service_directory_service.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }

  address = "1.2.3.4"
  port    = 5353
}
`, testId, location, testId, testId)
}
