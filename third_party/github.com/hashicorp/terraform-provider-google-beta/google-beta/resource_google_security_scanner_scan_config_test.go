package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecurityScannerScanConfig_scanConfigUpdate(t *testing.T) {
	t.Parallel()

	firstAddressSuffix := randString(t, 10)
	secondAddressSuffix := randString(t, 10)
	context := map[string]interface{}{
		"random_suffix":       firstAddressSuffix,
		"random_suffix2":      secondAddressSuffix,
		"static_address_name": "scanner_static_ip",
		"user_agent":          "CHROME_LINUX",
		"export":              "ENABLED",
		"max_qps":             10,
	}
	updateContext := map[string]interface{}{
		"random_suffix":       firstAddressSuffix,
		"random_suffix2":      secondAddressSuffix,
		"static_address_name": "scanner_static_ip_update",
		"user_agent":          "CHROME_ANDROID",
		"export":              "DISABLED",
		"max_qps":             20,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityScannerScanConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityScannerScanConfig(context),
			},
			{
				ResourceName:      "google_security_scanner_scan_config.terraform-scan-config",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecurityScannerScanConfig(updateContext),
			},
		},
	})
}

func testAccSecurityScannerScanConfig(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_address" "scanner_static_ip" {
  name     = "scan-static-ip-%{random_suffix}"
}

resource "google_compute_address" "scanner_static_ip_update" {
  name     = "scan-static-ip-%{random_suffix2}"
}

resource "google_security_scanner_scan_config" "terraform-scan-config" {
  display_name                      = "terraform-scan-config-%{random_suffix}"
  max_qps                           = %{max_qps}
  starting_urls                     = ["http://${google_compute_address.%{static_address_name}.address}"]
  target_platforms                  = ["COMPUTE"]
  user_agent                        = "%{user_agent}"
  export_to_security_command_center = "%{export}"
}
`, context)
}
