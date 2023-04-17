package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleFirebaseHostingChannel(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleFirebaseHostingChannel(context),
				Check: resource.ComposeTestCheckFunc(
					CheckDataSourceStateMatchesResourceState(
						"data.google_firebase_hosting_channel.channel",
						"google_firebase_hosting_channel.channel",
					),
				),
			},
		},
	})
}

func testAccDataSourceGoogleFirebaseHostingChannel(context map[string]interface{}) string {
	return Nprintf(`
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
