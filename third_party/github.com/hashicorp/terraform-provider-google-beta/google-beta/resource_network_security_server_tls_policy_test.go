// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNetworkSecurityServerTlsPolicy_update(t *testing.T) {
	t.Parallel()

	serverTlsPolicyName := fmt.Sprintf("tf-test-server-tls-policy-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityServerTlsPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityServerTlsPolicy_basic(serverTlsPolicyName),
			},
			{
				ResourceName:      "google_network_security_server_tls_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityServerTlsPolicy_update(serverTlsPolicyName),
			},
			{
				ResourceName:      "google_network_security_server_tls_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityServerTlsPolicy_basic(serverTlsPolicyName string) string {
	return fmt.Sprintf(`
  resource "google_network_security_server_tls_policy" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description            = "my description"
		allow_open             = "false"
    server_certificate {
      certificate_provider_instance {
        plugin_instance = "google_cloud_private_spiffe"
      }
    }
  }
`, serverTlsPolicyName)
}

func testAccNetworkSecurityServerTlsPolicy_update(serverTlsPolicyName string) string {
	return fmt.Sprintf(`
  resource "google_network_security_server_tls_policy" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description            = "updated description"
    allow_open             = "false"
    server_certificate {
      certificate_provider_instance {
        plugin_instance = "google_cloud_private_spiffe"
      }
    }
  }
`, serverTlsPolicyName)
}
