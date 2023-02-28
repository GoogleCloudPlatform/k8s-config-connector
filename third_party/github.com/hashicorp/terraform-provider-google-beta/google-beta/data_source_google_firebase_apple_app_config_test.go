package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceGoogleFirebaseAppleAppConfig(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":   getTestProjectFromEnv(),
		"bundle_id":    "apple.app." + randString(t, 5),
		"display_name": "tf-test Display Name AppleAppConfig DataSource",
		"app_store_id": 12345,
		"team_id":      1234567890,
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseAppleAppConfig(context),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceFirebaseAppleAppConfigCheck("data.google_firebase_apple_app_config.my_app_config"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseAppleAppConfig(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_apple_app" "my_app_config" {
  project = "%{project_id}"
  bundle_id = "%{bundle_id}"
  display_name = "%{display_name}"
  app_store_id = "%{app_store_id}"
  team_id = "%{team_id}"
}

data "google_firebase_apple_app_config" "my_app_config" {
  app_id = google_firebase_apple_app.my_app_config.app_id
}
`, context)
}

func testAccDataSourceFirebaseAppleAppConfigCheck(datasourceName string) resource.TestCheckFunc {
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
