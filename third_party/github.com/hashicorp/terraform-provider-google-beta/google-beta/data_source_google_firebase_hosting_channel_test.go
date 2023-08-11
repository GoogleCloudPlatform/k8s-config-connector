// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccDataSourceGoogleFirebaseHostingChannel(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseHostingChannel(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState(
						"data.google_firebase_hosting_channel.channel",
						"google_firebase_hosting_channel.channel",
					),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseHostingChannel(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
	project  = "%{project_id}"
	site_id = "tf-test-site-with-channel%{random_suffix}"
}
	
resource "google_firebase_hosting_channel" "channel" {
	site_id = google_firebase_hosting_site.default.site_id
	channel_id = "tf-test-channel%{random_suffix}"
}

data "google_firebase_hosting_channel" "channel" {
	site_id = google_firebase_hosting_site.default.site_id
	channel_id = "tf-test-channel%{random_suffix}"

	depends_on = [google_firebase_hosting_channel.channel]
}
`, context)
}
