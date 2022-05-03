package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApiGatewayApi_apigatewayApiBasicExampleUpdated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayApiDestroyProducer(t),
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
	return Nprintf(`
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
