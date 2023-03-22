package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFirebaseWebApp_firebaseWebAppFull(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
		"display_name":  "tf-test Display Name N",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: TestAccProvidersOiCS,
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
	labels     = {
		"firebase" = "enabled"
	}
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

func TestAccFirebaseWebApp_firebaseWebAppSkipDelete(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
		"display_name":  "tf-test Display Name N",
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseWebAppNotDestroyedProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseWebApp_firebaseWebAppSkipDelete(context, ""),
			},
			{
				ResourceName:            "google_firebase_web_app.skip_delete",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_policy", "project"},
			},
		},
	})
}

func testAccFirebaseWebApp_firebaseWebAppSkipDelete(context map[string]interface{}, update string) string {
	return Nprintf(`
resource "google_firebase_web_app" "skip_delete" {
	provider = google-beta
	project = "%{project_id}"
	display_name = "%{display_name} %{random_suffix}"
}
`, context)
}

func testAccCheckFirebaseWebAppNotDestroyedProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_web_app" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FirebaseBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err != nil {
				return fmt.Errorf("FirebaseWebApp doesn't exists at %s", url)
			}
		}

		return nil
	}
}
