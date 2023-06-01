// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccProjectServiceIdentity_basic(t *testing.T) {
	t.Parallel()

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testGoogleProjectServiceIdentity_basic(),
				Check: resource.ComposeTestCheckFunc(
					// Email field for logging service identity will be empty for as long as
					// `gcloud beta services identity create --service=logging.googleapis.com` doesn't return an email address
					resource.TestCheckNoResourceAttr("google_project_service_identity.log_sa", "email"),
				),
			},
		},
	})
}

func testGoogleProjectServiceIdentity_basic() string {
	return `
data "google_project" "project" {}

# Service which has an email returned from service identity API
resource "google_project_service_identity" "hc_sa" {
  project = data.google_project.project.project_id
  service = "healthcare.googleapis.com"
}

resource "google_project_iam_member" "hc_sa_bq_jobuser" {
  project = data.google_project.project.project_id
  role    = "roles/bigquery.jobUser"
  member  = "serviceAccount:${google_project_service_identity.hc_sa.email}"
}

# Service which does NOT have email returned from service identity API
# Email attribute will be null - correct as of 2022-12-13
resource "google_project_service_identity" "log_sa" {
  project = data.google_project.project.project_id
  service = "logging.googleapis.com"
}
`
}
