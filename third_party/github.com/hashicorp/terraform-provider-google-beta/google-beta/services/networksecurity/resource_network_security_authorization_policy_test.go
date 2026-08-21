// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package networksecurity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNetworkSecurityAuthorizationPolicy_update(t *testing.T) {
	t.Parallel()

	authorizationPolicyName := fmt.Sprintf("tf-test-authorization-policy-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
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
