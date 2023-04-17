package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseHostingSite_firebasehostingSiteUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
		"site_id":       "tf-test-site-update-app",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseHostingSiteDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingSite_firebasehostingSiteBeforeUpdate(context),
			},
			{
				ResourceName:            "google_firebase_hosting_site.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"site_id"},
			},
			{
				Config: testAccFirebaseHostingSite_firebasehostingSiteAfterUpdate(context),
			},
			{
				ResourceName:            "google_firebase_hosting_site.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"site_id"},
			},
		},
	})
}

func testAccFirebaseHostingSite_firebasehostingSiteBeforeUpdate(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_web_app" "before" {
  provider = google-beta
  project  = "%{project_id}"
  display_name = "tf-test Test web app before for Firebase Hosting"
  deletion_policy = "DELETE"
}

resource "google_firebase_hosting_site" "update" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "%{site_id}%{random_suffix}"
  app_id = google_firebase_web_app.before.app_id
}
`, context)
}

func testAccFirebaseHostingSite_firebasehostingSiteAfterUpdate(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_web_app" "after" {
  provider = google-beta
  project  = "%{project_id}"
  display_name = "tf-test Test web app after for Firebase Hosting"
  deletion_policy = "DELETE"
}

resource "google_firebase_hosting_site" "update" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "%{site_id}%{random_suffix}"
  app_id = google_firebase_web_app.after.app_id
}
`, context)
}
