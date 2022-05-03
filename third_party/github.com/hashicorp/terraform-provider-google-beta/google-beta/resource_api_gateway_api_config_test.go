package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApiGatewayApiConfig_apigatewayApiConfigBasicExampleUpdated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigBasicExample(context),
			},
			{
				Config: testAccApiGatewayApiConfig_apigatewayApiConfigBasicExampleUpdated(context),
			},
		},
	})
}

func TestAccApiGatewayApiConfig_generatedPrefix(t *testing.T) {
	// Random generated id within resource
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckApiGatewayApiConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfig_generatedPrefix(context),
			},
		},
	})
}

func testAccApiGatewayApiConfig_apigatewayApiConfigBasicExampleUpdated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "tf-test-api-cfg%{random_suffix}"
  display_name = "MM Dev API Config"
  labels = {
    environment = "dev"
  }

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
}
`, context)
}

func testAccApiGatewayApiConfig_generatedPrefix(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id_prefix = "tf-test-"
  display_name = "MM Dev API Config"
  labels = {
    environment = "dev"
  }

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
}
`, context)
}
