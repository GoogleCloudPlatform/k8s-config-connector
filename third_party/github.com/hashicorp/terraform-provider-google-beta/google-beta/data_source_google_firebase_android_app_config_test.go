// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"
)

func TestAccDataSourceGoogleFirebaseAndroidAppConfig(t *testing.T) {
	t.Parallel()
	// Framework-based resources and datasources don't work with VCR yet
	acctest.SkipIfVcr(t)

	context := map[string]interface{}{
		"project_id":   envvar.GetTestProjectFromEnv(),
		"package_name": "android.app." + acctest.RandString(t, 5),
		"display_name": "tf-test Display Name AndroidAppConfig DataSource",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseAndroidAppConfig(context),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceFirebaseAndroidAppConfigCheck("data.google_firebase_android_app_config.my_app_config"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseAndroidAppConfig(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_android_app" "my_app_config" {
  project = "%{project_id}"
  package_name = "%{package_name}"
  display_name = "%{display_name}"
}

data "google_firebase_android_app_config" "my_app_config" {
  project = "%{project_id}"
  app_id = google_firebase_android_app.my_app_config.app_id
}
`, context)
}

func testAccDataSourceFirebaseAndroidAppConfigCheck(datasourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[datasourceName]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", datasourceName)
		}

		if ds.Primary.Attributes["config_filename"] == "" {
			return fmt.Errorf("config filename not found in data source")
		}

		if ds.Primary.Attributes["config_file_contents"] == "" {
			return fmt.Errorf("config file contents not found in data source")
		}

		return nil
	}
}
