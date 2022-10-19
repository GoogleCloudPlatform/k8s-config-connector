package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccBigqueryDatapolicyDataPolicy_bigqueryDatapolicyDataPolicyUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckBigqueryDatapolicyDataPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryDatapolicyDataPolicy_bigqueryDatapolicyDataPolicyBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_datapolicy_data_policy.data_policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
			{
				Config: testAccBigqueryDatapolicyDataPolicy_bigqueryDatapolicyDataPolicyUpdate(context),
			},
		},
	})
}

func testAccBigqueryDatapolicyDataPolicy_bigqueryDatapolicyDataPolicyUpdate(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_datapolicy_data_policy" "data_policy" {
    provider = google-beta
    location         = "us-central1"
    data_policy_id   = "tf_test_data_policy%{random_suffix}"
    policy_tag       = google_data_catalog_policy_tag.policy_tag_updated.name
    data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
  }

  resource "google_data_catalog_policy_tag" "policy_tag" {
    provider = google-beta
    taxonomy     = google_data_catalog_taxonomy.taxonomy.id
    display_name = "Low security"
    description  = "A policy tag normally associated with low security items"
  }

  resource "google_data_catalog_policy_tag" "policy_tag_updated" {
    provider = google-beta
    taxonomy     = google_data_catalog_taxonomy.taxonomy.id
    display_name = "Low security updated"
    description  = "A policy tag normally associated with low security items"
  }  
  
  resource "google_data_catalog_taxonomy" "taxonomy" {
    provider = google-beta
    region                 = "us-central1"
    display_name           = "taxonomy%{random_suffix}"
    description            = "A collection of policy tags"
    activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
  }
`, context)
}
