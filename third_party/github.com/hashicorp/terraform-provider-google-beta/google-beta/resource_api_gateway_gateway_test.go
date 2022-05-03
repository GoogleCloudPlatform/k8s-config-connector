package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApiGatewayGateway_apigatewayGatewayBasicExampleUpdated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayGateway_apigatewayGatewayBasicExample(context),
			},
			{
				Config: testAccApiGatewayGateway_apigatewayGatewayBasicExampleUpdated(context),
			},
		},
	})
}

func testAccApiGatewayGateway_apigatewayGatewayBasicExampleUpdated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_gw" {
  provider = google-beta
  api_id = "tf-test-api-gw%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_gw" {
  provider = google-beta
  api = google_api_gateway_api.api_gw.api_id
  api_config_id = "tf-test-api-gw%{random_suffix}"
	lifecycle {
    create_before_destroy = true
  }

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_gw.id
  gateway_id = "tf-test-api-gw%{random_suffix}"
  display_name = "MM Dev API Gateway"
  labels = {
    environment = "dev"
  }
}
`, context)
}
