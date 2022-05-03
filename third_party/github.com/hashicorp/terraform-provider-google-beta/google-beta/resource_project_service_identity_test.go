package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccProjectServiceIdentity_basic(t *testing.T) {
	t.Parallel()

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testGoogleProjectServiceIdentity_basic(),
			},
		},
	})
}

func testGoogleProjectServiceIdentity_basic() string {
	return `
data "google_project" "project" {}

resource "google_project_service_identity" "hc_sa" {
  project = data.google_project.project.project_id
  service = "healthcare.googleapis.com"
}

resource "google_project_iam_member" "hc_sa_bq_jobuser" {
  project = data.google_project.project.project_id
  role    = "roles/bigquery.jobUser"
  member  = "serviceAccount:${google_project_service_identity.hc_sa.email}"
}`
}
