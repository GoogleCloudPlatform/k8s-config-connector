package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseProject_destroyAndReapply(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        GetTestOrgFromEnv(t),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseProject_firebaseProjectBasicExample(context),
			},
			{
				Config: testAccFirebaseProject_firebaseProjectBasicExampleDestroyed(context),
			},
			{
				Config: testAccFirebaseProject_firebaseProjectBasicExample(context),
			},
			{
				ResourceName:      "google_firebase_project.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFirebaseProject_firebaseProjectBasicExampleDestroyed(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "default" {
  provider = google-beta

  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"

  labels = {
    "firebase" = "enabled"
  }
}
`, context)
}
