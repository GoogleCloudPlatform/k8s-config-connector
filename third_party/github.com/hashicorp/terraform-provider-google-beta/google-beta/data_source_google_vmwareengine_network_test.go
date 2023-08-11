// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVmwareEngineNetwork_basic(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"region":          envvar.GetTestRegionFromEnv(),
		"random_suffix":   acctest.RandString(t, 10),
		"organization":    envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckVmwareengineNetworkDestroyProducer(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVmwareEngineNetworkConfig(context),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceStateWithIgnores("data.google_vmwareengine_network.ds", "google_vmwareengine_network.nw", map[string]struct{}{}),
				),
			},
		},
	})
}

func testAccDataSourceVmwareEngineNetworkConfig(context map[string]interface{}) string {
	return acctest.Nprintf(`
 # there can be only 1 Legacy network per region for a given project, so creating new project to isolate tests.
resource "google_project" "acceptance" {
  name            = "tf-test-%{random_suffix}"
  provider        = google-beta
  project_id      = "tf-test-%{random_suffix}"
  org_id          = "%{organization}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "acceptance" {
  project  = google_project.acceptance.project_id
  provider = google-beta
  service  = "vmwareengine.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.acceptance]

  create_duration = "60s"
}

resource "google_vmwareengine_network" "nw" {
  project     = google_project_service.acceptance.project
  name        = "%{region}-default" #Legacy network IDs are in the format: {region-id}-default
  provider    = google-beta
  location    = "%{region}"
  type        = "LEGACY"
  description = "VMwareEngine legacy network sample"
}

data "google_vmwareengine_network" "ds" {
  name     = google_vmwareengine_network.nw.name
  project  = google_project_service.acceptance.project
  provider = google-beta
  location = "%{region}"
  depends_on = [
    google_vmwareengine_network.nw,
  ]
}
`, context)
}
