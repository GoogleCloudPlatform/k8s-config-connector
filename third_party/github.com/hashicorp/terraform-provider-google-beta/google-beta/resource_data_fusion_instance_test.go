package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataFusionInstance_update(t *testing.T) {
	t.Parallel()

	instanceName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstance_basic(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataFusionInstance_updated(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataFusionInstance_basic(instanceName string) string {
	return fmt.Sprintf(`
resource "google_data_fusion_instance" "foobar" {
  name   = "%s"
  region = "us-central1"
  type   = "BASIC"
  version = "6.1.1"
}
`, instanceName)
}

func testAccDataFusionInstance_updated(instanceName string) string {
	return fmt.Sprintf(`
resource "google_data_fusion_instance" "foobar" {
  name                          = "%s"
  region                        = "us-central1"
  type                          = "DEVELOPER"
  enable_stackdriver_monitoring = true
  enable_stackdriver_logging    = true

  labels = {
    label1 = "value1"
    label2 = "value2"
  }
  version = "6.2.0"
}
`, instanceName)
}
