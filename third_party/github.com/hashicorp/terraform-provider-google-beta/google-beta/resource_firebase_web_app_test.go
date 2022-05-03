package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseWebApp_firebaseWebAppFull(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
		"display_name":  "Display Name N",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseWebApp_firebaseWebAppFull(context, ""),
			},
			{
				Config: testAccFirebaseWebApp_firebaseWebAppFull(context, "2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "api_key"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "auth_domain"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "storage_bucket"),
				),
			},
		},
	})
}

func testAccFirebaseWebApp_firebaseWebAppFull(context map[string]interface{}, update string) string {
	context["display_name"] = context["display_name"].(string) + update
	return Nprintf(`
resource "google_project" "default" {
	provider = google-beta

	project_id = "tf-test%{random_suffix}"
	name       = "tf-test%{random_suffix}"
	org_id     = "%{org_id}"
}

resource "google_firebase_project" "default" {
	provider = google-beta
	project  = google_project.default.project_id
}

resource "google_firebase_web_app" "default" {
	provider = google-beta
	project = google_project.default.project_id
	display_name = "%{display_name} %{random_suffix}"

	depends_on = [google_firebase_project.default]
}

data "google_firebase_web_app_config" "default" {
	provider   = google-beta
	web_app_id = google_firebase_web_app.default.app_id
}
`, context)
}
