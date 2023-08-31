// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseAppleApp_update(t *testing.T) {
	t.Parallel()
	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"bundle_id":     "apple.app.12345",
		"random_suffix": acctest.RandString(t, 10),
		"display_name":  "tf-test Display Name N",
	}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAppleApp(context, 12345, "1", "key1"),
			},
			{
				Config: testAccFirebaseAppleApp(context, 67890, "2", "key2"),
			},
		},
	})
}

func testAccFirebaseAppleApp(context map[string]interface{}, appStoreId int, delta string, apiKeyLabel string) string {
	context["display_name"] = context["display_name"].(string) + delta
	context["app_store_id"] = appStoreId
	context["team_id"] = "123456789" + delta
	context["api_key_label"] = apiKeyLabel
	return acctest.Nprintf(`
resource "google_firebase_apple_app" "update" {
  provider = google-beta
  project  = "%{project_id}"

  bundle_id    = "%{bundle_id}"
  display_name = "%{display_name} %{random_suffix}"
  app_store_id = "%{app_store_id}"
  team_id      = "%{team_id}"
  api_key_id   = google_apikeys_key.%{api_key_label}.uid
}

resource "google_apikeys_key" "key1" {
  provider = google-beta
  project  = "%{project_id}"

  name         = "tf-test-api-key1%{random_suffix}"
  display_name = "Test api key 1"
  
  restrictions {
    ios_key_restrictions {
      allowed_bundle_ids = ["%{bundle_id}"]
    }
  }
}

resource "google_apikeys_key" "key2" {
  provider = google-beta
  project  = "%{project_id}"

  name         = "tf-test-api-key2%{random_suffix}"
  display_name = "Test api key 2"
  
  restrictions {
    ios_key_restrictions {
      allowed_bundle_ids = ["%{bundle_id}"]
    }
  }
}
`, context)
}
