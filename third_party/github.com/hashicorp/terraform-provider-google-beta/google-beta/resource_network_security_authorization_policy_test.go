// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkSecurityAuthorizationPolicy_update(t *testing.T) {
	t.Parallel()

	authorizationPolicyName := fmt.Sprintf("tf-test-authorization-policy-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAuthorizationPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAuthorizationPolicy_basic(authorizationPolicyName),
			},
			{
				ResourceName:      "google_network_security_authorization_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityAuthorizationPolicy_update(authorizationPolicyName),
			},
			{
				ResourceName:      "google_network_security_authorization_policy.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityAuthorizationPolicy_basic(authorizationPolicyName string) string {
	return fmt.Sprintf(`
  resource "google_network_security_authorization_policy" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description            = "my description"
    action                 = "ALLOW"
    rules {
      sources {
        principals = ["namespace/*"]
        ip_blocks = ["1.2.3.0/24"]
      }
    }
  }
`, authorizationPolicyName)
}

func testAccNetworkSecurityAuthorizationPolicy_update(authorizationPolicyName string) string {
	return fmt.Sprintf(`
  resource "google_network_security_authorization_policy" "foobar" {
    name                   = "%s"
    labels                 = {
      foo = "bar"
    }
    description            = "updated description"
    action                 = "DENY"
    rules {
      sources {
        principals = ["namespace1/*"]
        ip_blocks = ["1.2.3.0/24"]
      }
    }
  }
`, authorizationPolicyName)
}
