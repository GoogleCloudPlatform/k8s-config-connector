// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package securityscanner_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecurityScannerScanConfig_scanConfigUpdate(t *testing.T) {
	t.Parallel()

	firstAddressSuffix := acctest.RandString(t, 10)
	secondAddressSuffix := acctest.RandString(t, 10)
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

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecurityScannerScanConfigDestroyProducer(t),
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
	return acctest.Nprintf(`
resource "google_compute_address" "scanner_static_ip" {
  name     = "tf-test-scan-static-ip-%{random_suffix}"
}

resource "google_compute_address" "scanner_static_ip_update" {
  name     = "tf-test-scan-static-ip-%{random_suffix2}"
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
