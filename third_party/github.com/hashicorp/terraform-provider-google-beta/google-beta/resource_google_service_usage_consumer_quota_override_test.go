// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServiceUsageConsumerQuotaOverride_consumerQuotaOverrideCustomIncorrectLimitFormat(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckServiceUsageConsumerQuotaOverrideDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config:      testAccServiceUsageConsumerQuotaOverride_consumerQuotaOverrideCustomIncorrectLimitFormat(context),
				ExpectError: regexp.MustCompile("No quota limit with limitId"),
			},
		},
	})
}

func testAccServiceUsageConsumerQuotaOverride_consumerQuotaOverrideCustomIncorrectLimitFormat(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "my_project" {
  provider   = google-beta
  name       = "tf-test-project"
  project_id = "quota%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "google_service_usage_consumer_quota_override" "override" {
  provider       = google-beta
  project        = google_project.my_project.project_id
  service        = urlencode("bigquery.googleapis.com")
  metric         = urlencode("bigquery.googleapis.com/quota/query/usage")
  limit          = urlencode("1/d/{project}/{user}") # Incorrect format for the API the provider uses, correct format for the gcloud CLI
  override_value = "1"
  force          = true
}
`, context)
}
