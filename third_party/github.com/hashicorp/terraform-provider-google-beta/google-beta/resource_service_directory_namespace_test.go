package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServiceDirectoryNamespace_serviceDirectoryNamespaceUpdateExample(t *testing.T) {
	t.Parallel()

	project := getTestProjectFromEnv()
	location := "us-central1"
	testId := fmt.Sprintf("tf-test-example-namespace%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDirectoryNamespaceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryNamespace_basic(location, testId),
			},
			{
				ResourceName:      "google_service_directory_namespace.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_namespace.example",
				// {{project}}/{{location}}/{{namespace_id}}
				ImportStateId:     fmt.Sprintf("%s/%s/%s", project, location, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName: "google_service_directory_namespace.example",
				// {{location}}/{{namespace_id}}
				ImportStateId:     fmt.Sprintf("%s/%s", location, testId),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccServiceDirectoryNamespace_update(location, testId),
			},
			{
				ResourceName:      "google_service_directory_namespace.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccServiceDirectoryNamespace_basic(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"
}
`, testId, location)
}

func testAccServiceDirectoryNamespace_update(location, testId string) string {
	return fmt.Sprintf(`
resource "google_service_directory_namespace" "example" {
  namespace_id = "%s"
  location     = "%s"

  labels = {
    key = "value"
    foo = "bar"
  }
}
`, testId, location)
}
