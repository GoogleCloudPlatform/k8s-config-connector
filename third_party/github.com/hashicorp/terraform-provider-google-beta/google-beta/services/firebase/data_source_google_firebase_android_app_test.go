// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataSourceGoogleFirebaseAndroidApp(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":   envvar.GetTestProjectFromEnv(),
		"package_name": "android.package.app" + acctest.RandString(t, 4),
		"display_name": "tf-test Display Name AndroidApp DataSource",
	}

	resourceName := "data.google_firebase_android_app.my_app"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseAndroidApp(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores(
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
	return acctest.Nprintf(`
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

data "google_firebase_android_app" "my_app_project" {
  project = "%{project_id}"
  app_id = google_firebase_android_app.my_app.app_id
}
`, context)
}
