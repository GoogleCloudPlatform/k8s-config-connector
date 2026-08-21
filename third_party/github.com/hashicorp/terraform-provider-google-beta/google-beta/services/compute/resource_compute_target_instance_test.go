// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeTargetInstance_withSecurityPolicy(t *testing.T) {
	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeTargetInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetInstance_withSecurityPolicy(context, "google_compute_region_security_policy.regionsecuritypolicy1.self_link", true),
			},
			{
				ResourceName:            "google_compute_target_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "zone"},
			},
			{
				Config: testAccComputeTargetInstance_withSecurityPolicy(context, "google_compute_region_security_policy.regionsecuritypolicy2.self_link", true),
			},
			{
				ResourceName:            "google_compute_target_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "zone"},
			},
			{
				Config: testAccComputeTargetInstance_withSecurityPolicy(context, "\"\"", true),
			},
			{
				ResourceName:            "google_compute_target_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance", "zone"},
			},
			{
				Config: testAccComputeTargetInstance_withSecurityPolicy(context, "\"\"", false),
			},
		},
	})
}

func testAccComputeTargetInstance_withSecurityPolicy(context map[string]interface{}, policySet string, preventDestroy bool) string {
	context["policy_set"] = policySet
	context["lifecycle_block"] = ""
	if preventDestroy {
		context["lifecycle_block"] = `
		lifecycle {
			prevent_destroy = true
		}`
	}

	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-up-pol-net%{random_suffix}"
  auto_create_subnetworks = false
  routing_mode            = "REGIONAL"
}
      
resource "google_compute_subnetwork" "default" {
  provider                   = google-beta
  name                       = "tf-test-up-pol-subnet%{random_suffix}"
  ip_cidr_range              = "10.1.2.0/24"
  network                    = google_compute_network.default.id
  private_ipv6_google_access = "DISABLE_GOOGLE_ACCESS"
  purpose                    = "PRIVATE"
  region                     = "southamerica-east1"
  stack_type                 = "IPV4_ONLY"
}

data "google_compute_image" "vmimage" {
  provider = google-beta
  family   = "debian-11"
  project  = "debian-cloud"
}

resource "google_compute_instance" "target-vm" {
  provider     = google-beta
  name         = "tf-test-up-pol-target-vm%{random_suffix}"
  machine_type = "e2-medium"
  zone         = "southamerica-east1-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.vmimage.self_link
    }
  }

  network_interface {       
    network = google_compute_network.default.self_link
    subnetwork = google_compute_subnetwork.default.self_link
    access_config {
    }
  }
}

resource "google_compute_region_security_policy" "policyddosprotection" {
  provider    = google-beta
  region      = "southamerica-east1"
  name        = "tf-test-up-pol-policyddos%{random_suffix}"
  description = "ddos protection security policy to set target instance"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}

resource "google_compute_network_edge_security_service" "edge_sec_service" {
  provider        = google-beta
  region          = "southamerica-east1"
  name            = "tf-test-up-pol-edgesec%{random_suffix}"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}

resource "google_compute_region_security_policy" "regionsecuritypolicy1" {
  provider    = google-beta
  name        = "tf-test-up-pol-region-secpolicy1%{random_suffix}"
  region      = "southamerica-east1"
  description = "basic security policy one for target instance"
  type        = "CLOUD_ARMOR_NETWORK"
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}

resource "google_compute_region_security_policy" "regionsecuritypolicy2" {
  provider    = google-beta
  name        = "tf-test-up-pol-region-secpolicy2%{random_suffix}"
  region      = "southamerica-east1"
  description = "basic security policy two for target instance"
  type        = "CLOUD_ARMOR_NETWORK"
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}

resource "google_compute_target_instance" "default" {
  provider        = google-beta
  name            = "tf-test-up-pol-target-instance%{random_suffix}"
  zone            = "southamerica-east1-a"
  instance        = google_compute_instance.target-vm.id
  security_policy = %{policy_set}
  %{lifecycle_block}
}
`, context)
}
