// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNetworkSecurityClientTlsPolicy_update(t *testing.T) {
	t.Parallel()

	clientTlsPolicyName := fmt.Sprintf("tf-test-client-tls-policy-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityClientTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityClientTlsPolicy_basic(clientTlsPolicyName),
			},
			{
				ResourceName:      "google_network_security_client_tls_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityClientTlsPolicy_update(clientTlsPolicyName),
			},
			{
				ResourceName:      "google_network_security_client_tls_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityClientTlsPolicy_basic(clientTlsPolicyName string) string {
	return fmt.Sprintf(`
  resource "google_network_security_client_tls_policy" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description            = "my description"
    sni                    = "secure.example.com"
    client_certificate {
      certificate_provider_instance {
        plugin_instance = "google_cloud_private_spiffe"
      }
    }
    server_validation_ca {
    	grpc_endpoint {
      	target_uri = "unix:mypath"
      }
    }
  }
`, clientTlsPolicyName)
}

func testAccNetworkSecurityClientTlsPolicy_update(clientTlsPolicyName string) string {
	return fmt.Sprintf(`
  resource "google_network_security_client_tls_policy" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description            = "updated description"
    sni                    = "secure1.example.com"
    client_certificate {
      certificate_provider_instance {
        plugin_instance = "google_cloud"
      }
    }
    server_validation_ca {
    	grpc_endpoint {
      	target_uri = "unix:mypath1"
      }
    }
    server_validation_ca {
    	grpc_endpoint {
      	target_uri = "unix:mypath2"
      }
    }
  }
`, clientTlsPolicyName)
}
