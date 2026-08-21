// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package apigateway_test

import (
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApiGatewayApi_apigatewayApiBasicExampleUpdated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckApiGatewayApiDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApi_apigatewayApiBasicExample(context),
			},
			{
				Config: testAccApiGatewayApi_apigatewayApiBasicExampleUpdated(context),
			},
		},
	})
}

func testAccApiGatewayApi_apigatewayApiBasicExampleUpdated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "api%{random_suffix}"
  display_name = "Magical API"
  labels = {
	environment = "dev"
  }
}
`, context)
}
