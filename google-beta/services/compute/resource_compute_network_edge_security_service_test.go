// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeNetworkEdgeSecurityService_update(t *testing.T) {
	t.Parallel()

	pName := fmt.Sprintf("tf-test-security-policy-%s", acctest.RandString(t, 10))
	nesName := fmt.Sprintf("tf-test-edge-security-services-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNetworkEdgeSecurityServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkEdgeSecurityService_basic(pName, nesName),
			},
			{
				ResourceName:      "google_compute_network_edge_security_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkEdgeSecurityService_update(pName, nesName),
			},
			{
				ResourceName:      "google_compute_network_edge_security_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkEdgeSecurityService_basic(pName, nesName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "foobar" {
  name        = "%s"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR_NETWORK"
  # can only exist one of this resource per region
  region      = "us-central1"
}

resource "google_compute_network_edge_security_service" "foobar" {
  name            = "%s"
  region          = "us-central1"
  description     = "My basic resource using security policy"
  security_policy = google_compute_region_security_policy.foobar.self_link
}
`, pName, nesName)
}

func testAccNetworkEdgeSecurityService_update(pName, nesName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "foobar" {
  name        = "%s"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR_NETWORK"
  region      = "us-central1"
}

resource "google_compute_network_edge_security_service" "foobar" {
  name            = "%s"
  region          = "us-central1"
  description     = "My basic updated resource using security policy"
  security_policy = google_compute_region_security_policy.foobar.self_link
}
`, pName, nesName)
}
