// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirebaseHostingChannel_firebasehostingChannelUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckFirebaseHostingChannelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelBasic(context),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelTtl(context, "8600s"),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelTtl(context, "86400s"),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelRetainedReleaseCount(context, 30),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelRetainedReleaseCount(context, 20),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelLabels(context),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelMultipleFields(context),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
			{
				Config: testAccFirebaseHostingChannel_firebasehostingChannelBasic(context),
			},
			{
				ResourceName:            "google_firebase_hosting_channel.update",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl", "site_id", "channel_id"},
			},
		},
	})
}

func testAccFirebaseHostingChannel_firebasehostingChannelBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_channel" "update" {
  provider = google-beta
  site_id = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-update%{random_suffix}"
}
`, context)
}

func testAccFirebaseHostingChannel_firebasehostingChannelTtl(context map[string]interface{}, ttl string) string {
	context["ttl"] = ttl
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_channel" "update" {
  provider = google-beta
  site_id = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-update%{random_suffix}"
  ttl = "%{ttl}"
}
`, context)
}

func testAccFirebaseHostingChannel_firebasehostingChannelRetainedReleaseCount(context map[string]interface{}, retainedReleaseCount int) string {
	context["retained_release_count"] = retainedReleaseCount
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_channel" "update" {
  provider = google-beta
  site_id = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-update%{random_suffix}"
  ttl = "86400s"
  retained_release_count = %{retained_release_count}
}
`, context)
}

func testAccFirebaseHostingChannel_firebasehostingChannelLabels(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_channel" "update" {
  provider = google-beta
  site_id = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-update%{random_suffix}"
  ttl = "86400s"
  retained_release_count = 10
  labels = {
    "some-key": "some-value"
  }
}
`, context)
}

func testAccFirebaseHostingChannel_firebasehostingChannelMultipleFields(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_firebase_hosting_site" "default" {
  provider = google-beta
  project  = "%{project_id}"
  site_id = "tf-test-site-with-channel%{random_suffix}"
}

resource "google_firebase_hosting_channel" "update" {
  provider = google-beta
  site_id = google_firebase_hosting_site.default.site_id
  channel_id = "tf-test-channel-update%{random_suffix}"
  ttl = "86400s"
  retained_release_count = 40
  labels = {
    "some-key-2": "some-value-2"
  }
}
`, context)
}
