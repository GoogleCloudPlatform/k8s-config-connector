// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package firebase_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccFirebaseWebApp_firebaseWebAppFull(t *testing.T) {
	// TODO: https://github.com/hashicorp/terraform-provider-google/issues/14158
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
		"display_name":  "tf-test Display Name N",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck: func() { acctest.AccTestPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"google": {
						VersionConstraint: "4.58.0",
						Source:            "hashicorp/google-beta",
					},
				},
				Config: testAccFirebaseWebApp_firebaseWebAppFull(context, "", "key1"),
			},
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"google": {
						VersionConstraint: "4.58.0",
						Source:            "hashicorp/google-beta",
					},
				},
				Config: testAccFirebaseWebApp_firebaseWebAppFull(context, "2", "key2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "api_key"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "auth_domain"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "storage_bucket"),
				),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
				Config:                   testAccFirebaseWebApp_firebaseWebAppFull(context, "", "key1"),
			},
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
				Config:                   testAccFirebaseWebApp_firebaseWebAppFull(context, "2", "key2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "api_key"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "auth_domain"),
					resource.TestCheckResourceAttrSet("data.google_firebase_web_app_config.default", "storage_bucket"),
				),
			},
		},
	})
}

func testAccFirebaseWebApp_firebaseWebAppFull(context map[string]interface{}, update string, apiKeyLabel string) string {
	context["display_name"] = context["display_name"].(string) + update
	context["api_key_label"] = apiKeyLabel
	return acctest.Nprintf(`
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

resource "google_apikeys_key" "key1" {
	provider     = google-beta
	name         = "tf-test-api-key1%{random_suffix}"
	display_name = "Test api key 1"
	project      = google_project.default.project_id

	restrictions {
		browser_key_restrictions {
			allowed_referrers = ["*"]
		}
	}
}

resource "google_apikeys_key" "key2" {
	provider     = google-beta
	name         = "tf-test-api-key2%{random_suffix}"
	display_name = "Test api key 2"
	project      = google_project.default.project_id

	restrictions {
		browser_key_restrictions {
			allowed_referrers = ["*"]
		}
	}
}

resource "google_firebase_web_app" "default" {
	provider = google-beta
	project = google_project.default.project_id
	display_name = "%{display_name} %{random_suffix}"
	api_key_id = google_apikeys_key.%{api_key_label}.uid

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
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
		"display_name":  "tf-test Display Name N",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseWebAppNotDestroyedProducer(t),
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
	return acctest.Nprintf(`
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

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{FirebaseBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				return fmt.Errorf("FirebaseWebApp doesn't exists at %s", url)
			}
		}

		return nil
	}
}
