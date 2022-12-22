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
  # Mark for testing to avoid service networking connection usage that is not cleaned up
  options = {
  	prober_test_run = "true"
  }
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
  # Mark for testing to avoid service networking connection usage that is not cleaned up
  options = {
  	prober_test_run = "true"
  }
}
`, instanceName)
}

func TestAccDataFusionInstanceEnterprise_update(t *testing.T) {
	t.Parallel()

	instanceName := fmt.Sprintf("tf-test-%s", randString(t, 10))

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataFusionInstanceEnterprise_basic(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataFusionInstanceEnterprise_updated(instanceName),
			},
			{
				ResourceName:      "google_data_fusion_instance.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataFusionInstanceEnterprise_basic(instanceName string) string {
	return fmt.Sprintf(`
resource "google_data_fusion_instance" "foobar" {
  name   = "%s"
  region = "us-central1"
  type   = "ENTERPRISE"
  # Mark for testing to avoid service networking connection usage that is not cleaned up
  options = {
  	prober_test_run = "true"
  }
}
`, instanceName)
}

func testAccDataFusionInstanceEnterprise_updated(instanceName string) string {
	return fmt.Sprintf(`
resource "google_data_fusion_instance" "foobar" {
  name                          = "%s"
  region                        = "us-central1"
  type                          = "ENTERPRISE"
  enable_stackdriver_monitoring = true
  enable_stackdriver_logging    = true
  enable_rbac                   = true

  labels = {
    label1 = "value1"
    label2 = "value2"
  }
  # Mark for testing to avoid service networking connection usage that is not cleaned up
  options = {
  	prober_test_run = "true"
  }
}
`, instanceName)
}
