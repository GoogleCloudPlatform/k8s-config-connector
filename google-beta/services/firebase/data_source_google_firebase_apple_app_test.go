// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataSourceGoogleFirebaseAppleApp(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":   envvar.GetTestProjectFromEnv(),
		"bundle_id":    "apple.app." + acctest.RandString(t, 5),
		"display_name": "tf-test Display Name AppleApp DataSource",
		"app_store_id": 12345,
		"team_id":      1234567890,
	}

	resourceName := "data.google_firebase_apple_app.my_app"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseAppleApp(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores(
						resourceName,
						"google_firebase_apple_app.my_app",
						map[string]struct{}{
							"deletion_policy": {},
						},
					),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseAppleApp(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_apple_app" "my_app" {
  project = "%{project_id}"
  bundle_id = "%{bundle_id}"
  display_name = "%{display_name}"
  app_store_id = "%{app_store_id}"
  team_id = "%{team_id}"
}

data "google_firebase_apple_app" "my_app" {
  app_id = google_firebase_apple_app.my_app.app_id
}

data "google_firebase_apple_app" "my_app_project" {
  project = "%{project_id}"
  app_id = google_firebase_apple_app.my_app.app_id
}
`, context)
}
