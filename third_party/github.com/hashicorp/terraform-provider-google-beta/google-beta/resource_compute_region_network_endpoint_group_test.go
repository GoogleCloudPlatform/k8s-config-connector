// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeRegionNetworkEndpointGroup_negWithServerlessDeployment(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_negWithServerlessDeployment(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.test_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_negWithServerlessDeployment(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  api_id = "tf-test-%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-config-%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-%{random_suffix}"
}

resource "google_compute_region_network_endpoint_group" "test_neg" {
  name                  = "tf-test-neg-%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  serverless_deployment {
    platform = "apigateway.googleapis.com"
    url_mask = format("<gateway>%s/hello", trimprefix(google_api_gateway_gateway.api_gw.default_hostname, "tf-test-%{random_suffix}"))
  }
}
`, context)
}
