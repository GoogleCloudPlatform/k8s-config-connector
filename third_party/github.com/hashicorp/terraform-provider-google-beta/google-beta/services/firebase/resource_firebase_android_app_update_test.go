// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseAndroidApp_update(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"package_name":  "android.package.app" + acctest.RandString(t, 4),
		"random_suffix": acctest.RandString(t, 10),
		"display_name":  "tf-test Display Name N",
	}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAndroidApp(context, "", "key1"),
			},
			{
				Config: testAccFirebaseAndroidApp(context, "2", "key2"),
			},
		},
	})
}

func testAccFirebaseAndroidApp(context map[string]interface{}, update string, apiKeyLabel string) string {
	context["display_name"] = context["display_name"].(string) + update
	context["api_key_label"] = apiKeyLabel
	return acctest.Nprintf(`
resource "google_firebase_android_app" "update" {
	provider = google-beta
	project  = "%{project_id}"

	package_name  = "%{package_name}"
	display_name  = "%{display_name} %{random_suffix}"
	sha1_hashes   = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
	sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]
	api_key_id    = google_apikeys_key.%{api_key_label}.uid
}

resource "google_apikeys_key" "key1" {
	provider = google-beta
	project  = "%{project_id}"

	name         = "tf-test-api-key1%{random_suffix}"
	display_name = "Test api key 1"
  
	restrictions {
		android_key_restrictions {
			allowed_applications {
				package_name     = "%{package_name}"
				sha1_fingerprint = "2145bdf698b8715039bd0e83f2069bed435ac21c"
			}
		}
    }
}

resource "google_apikeys_key" "key2" {
	provider = google-beta
	project = "%{project_id}"

	name         = "tf-test-api-key2%{random_suffix}"
	display_name = "Test api key 2"
  
	restrictions {
		android_key_restrictions {
			allowed_applications {
				package_name     = "%{package_name}"
				sha1_fingerprint = "2145bdf698b8715039bd0e83f2069bed435ac21c"
			}
		}
    }
}
`, context)
}
