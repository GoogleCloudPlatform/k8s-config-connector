// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataSourceGoogleFirebaseWebApp(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":   envvar.GetTestProjectFromEnv(),
		"display_name": "tf_test Display Name WebApp DataSource",
	}

	resourceName := "data.google_firebase_web_app.my_app"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseWebApp(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores(
						resourceName,
						"google_firebase_web_app.my_app",
						map[string]struct{}{
							"deletion_policy": {},
						},
					),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseWebApp(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_web_app" "my_app" {
  project = "%{project_id}"
  display_name = "%{display_name}"
  deletion_policy = "DELETE"
}

data "google_firebase_web_app" "my_app" {
  app_id = google_firebase_web_app.my_app.app_id
}

data "google_firebase_web_app" "my_app_project" {
  project = "%{project_id}"
  app_id = google_firebase_web_app.my_app.app_id
}
`, context)
}
