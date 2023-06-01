// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccDNSResponsePolicy_update(t *testing.T) {
	t.Parallel()

	responsePolicySuffix := RandString(t, 10)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDNSResponsePolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDnsResponsePolicy_privateUpdate(responsePolicySuffix, "network-1"),
			},
			{
				ResourceName:      "google_dns_response_policy.example-response-policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsResponsePolicy_privateUpdate(responsePolicySuffix, "network-2"),
			},
			{
				ResourceName:      "google_dns_response_policy.example-response-policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDnsResponsePolicy_privateUpdate(suffix, network string) string {
	return fmt.Sprintf(`
resource "google_dns_response_policy" "example-response-policy" {
  provider = google-beta

  response_policy_name = "tf-test-response-policy-%s"

  networks {
    network_url = google_compute_network.%s.self_link
  }
  gke_clusters {
	gke_cluster_name = google_container_cluster.cluster-1.id
  }
}

resource "google_compute_network" "network-1" {
  provider = google-beta

  name                    = "tf-test-network-1-%s"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  provider = google-beta
	
  name                    = "tf-test-network-2-%s"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnetwork-1" {
  provider = google-beta

  name                     = google_compute_network.network-1.name
  network                  = google_compute_network.network-1.name
  ip_cidr_range            = "10.0.36.0/24"
  region                   = "us-central1"
  private_ip_google_access = true

  secondary_ip_range {
    range_name    = "pod"
    ip_cidr_range = "10.0.0.0/19"
  }

  secondary_ip_range {
    range_name    = "svc"
    ip_cidr_range = "10.0.32.0/22"
  }
}

resource "google_container_cluster" "cluster-1" {
  provider = google-beta

  name               = "tf-test-cluster-1-%s"
  location           = "us-central1-c"
  initial_node_count = 1

  networking_mode = "VPC_NATIVE"
  default_snat_status {
    disabled = true
  }
  network    = google_compute_network.network-1.name
  subnetwork = google_compute_subnetwork.subnetwork-1.name

  private_cluster_config {
    enable_private_endpoint = true
    enable_private_nodes    = true
    master_ipv4_cidr_block  = "10.42.0.0/28"
    master_global_access_config {
      enabled = true
	}
  }
  master_authorized_networks_config {
  }
  ip_allocation_policy {
    cluster_secondary_range_name  = google_compute_subnetwork.subnetwork-1.secondary_ip_range[0].range_name
    services_secondary_range_name = google_compute_subnetwork.subnetwork-1.secondary_ip_range[1].range_name
  }
}
`, suffix, network, suffix, suffix, suffix)
}
