package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleFirebaseAndroidApp(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":   getTestProjectFromEnv(),
		"package_name": "android.package.app" + randString(t, 4),
		"display_name": "tf-test Display Name AndroidApp DataSource",
	}

	resourceName := "data.google_firebase_android_app.my_app"

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseAndroidApp(context),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceStateMatchesResourceStateWithIgnores(
						resourceName,
						"google_firebase_android_app.my_app",
						map[string]struct{}{
							"deletion_policy": {},
						},
					),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseAndroidApp(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_android_app" "my_app" {
  project = "%{project_id}"
  package_name = "%{package_name}"
  display_name = "%{display_name}"
  sha1_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
  sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]
}

data "google_firebase_android_app" "my_app" {
  app_id = google_firebase_android_app.my_app.app_id
}
`, context)
}
