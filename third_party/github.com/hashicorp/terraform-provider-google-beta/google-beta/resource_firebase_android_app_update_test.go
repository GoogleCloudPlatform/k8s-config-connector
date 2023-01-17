package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseAndroidApp_update(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"package_name":  "android.package.app" + randString(t, 4),
		"random_suffix": randString(t, 10),
		"display_name":  "Display Name N",
	}
	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAndroidApp(context, ""),
			},
			{
				Config: testAccFirebaseAndroidApp(context, "2"),
			},
		},
	})
}

func testAccFirebaseAndroidApp(context map[string]interface{}, update string) string {
	context["display_name"] = context["display_name"].(string) + update
	return Nprintf(`
resource "google_firebase_android_app" "update" {
        provider = google-beta
        project = "%{project_id}"
        package_name = "%{package_name}"
        display_name = "%{display_name} %{random_suffix}"
        sha1_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
        sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]
}
`, context)
}
