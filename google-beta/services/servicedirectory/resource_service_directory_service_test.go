// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package servicedirectory_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServiceDirectoryService_serviceDirectoryServiceUpdateExample(t *testing.T) {
	t.Parallel()

	project := envvar.GetTestProjectFromEnv()
	location := "us-central1"
	testId := fmt.Sprintf("tf-test-example-service%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckServiceDirectoryServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryService_basic(location, testId),
			},
			{
				ResourceName:      "google_service_directory_service.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_service.example",
				// {{project}}/{{location}}/{{namespace_id}}/{{service_id}}
				ImportStateId:     fmt.Sprintf("%s/%s/%s/%s", project, location, testId, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_service.example",
				// {{location}}/{{namespace_id}}/{{service_id}}
				ImportStateId:     fmt.Sprintf("%s/%s/%s", location, testId, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccServiceDirectoryService_update(location, testId),
			},
			{
				ResourceName:      "google_service_directory_service.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccServiceDirectoryService_basic(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"
}

resource "google_service_directory_service" "example" {
  service_id = "%s"
  namespace  = google_service_directory_namespace.example.id
}
`, testId, location, testId)
}

func testAccServiceDirectoryService_update(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"
}

resource "google_service_directory_service" "example" {
  service_id = "%s"
  namespace  = google_service_directory_namespace.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}
`, testId, location, testId)
}
