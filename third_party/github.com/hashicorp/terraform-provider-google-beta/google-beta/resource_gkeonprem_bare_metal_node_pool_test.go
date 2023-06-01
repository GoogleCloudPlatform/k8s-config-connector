// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdate(t *testing.T) {
	// TODO: https://github.com/hashicorp/terraform-provider-google/issues/14417
	t.Skip()

	t.Parallel()

	context := map[string]interface{}{}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckGkeonpremBareMetalNodePoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdateStart(context),
			},
			{
				ResourceName:      "google_gkeonprem_bare_metal_node_pool.nodepool",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdate(context),
			},
			{
				ResourceName:      "google_gkeonprem_bare_metal_node_pool.nodepool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdateStart(context map[string]interface{}) string {
	return Nprintf(`

  resource "google_gkeonprem_bare_metal_cluster" "cluster" {
    provider = google-beta
    name = "cluster"
    location = "us-west1"
    admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
    bare_metal_version = "1.12.3"
    network_config {
      island_mode_cidr {
        service_address_cidr_blocks = ["172.26.0.0/16"]
        pod_address_cidr_blocks = ["10.240.0.0/13"]
      }
    }
    control_plane {
      control_plane_node_pool_config {
        node_pool_config {
          labels = {}
          operating_system = "LINUX"
          node_configs {
            labels = {}
            node_ip = "10.200.0.9"
          }
        }
      }
    }
    load_balancer {
      port_config {
        control_plane_load_balancer_port = 443
      }
      vip_config {
        control_plane_vip = "10.200.0.13"
        ingress_vip = "10.200.0.14"
      }
      metal_lb_config {
        address_pools {
          pool = "pool1"
          addresses = [
            "10.200.0.14/32",
            "10.200.0.15/32",
            "10.200.0.16/32",
            "10.200.0.17/32",
            "10.200.0.18/32",
            "fd00:1::f/128",
            "fd00:1::10/128",
            "fd00:1::11/128",
            "fd00:1::12/128"
          ]
        }
      }
    }
    storage {
      lvp_share_config {
        lvp_config {
          path = "/mnt/localpv-share"
          storage_class = "local-shared"
        }
        shared_path_pv_count = 5
      }
      lvp_node_mounts_config {
        path = "/mnt/localpv-disk"
        storage_class = "local-disks"
      }
    }
    security_config {
      authorization {
        admin_users {
          username = "admin@hashicorptest.com"
        }
      }
    }
  }

  resource "google_gkeonprem_bare_metal_node_pool" "nodepool" {
    provider = google-beta
    location = "us-west1"
    name = "nodepool"
    bare_metal_cluster = google_gkeonprem_bare_metal_cluster.cluster.name
    annotations = {}
    node_pool_config {
      operating_system = "LINUX"
      labels = {}
      node_configs {
        node_ip = "10.200.0.11"
        labels = {}
      }
    }
  }
`, context)
}

func testAccGkeonpremBareMetalNodePool_bareMetalNodePoolUpdate(context map[string]interface{}) string {
	return Nprintf(`

  resource "google_gkeonprem_bare_metal_cluster" "cluster" {
    provider = google-beta
    name = "cluster"
    location = "us-west1"
    admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
    bare_metal_version = "1.12.3"
    network_config {
      island_mode_cidr {
        service_address_cidr_blocks = ["172.26.0.0/16"]
        pod_address_cidr_blocks = ["10.240.0.0/13"]
      }
    }
    control_plane {
      control_plane_node_pool_config {
        node_pool_config {
          labels = {}
          operating_system = "LINUX"
          node_configs {
            labels = {}
            node_ip = "10.200.0.9"
          }
        }
      }
    }
    load_balancer {
      port_config {
        control_plane_load_balancer_port = 443
      }
      vip_config {
        control_plane_vip = "10.200.0.13"
        ingress_vip = "10.200.0.14"
      }
      metal_lb_config {
        address_pools {
          pool = "pool1"
          addresses = [
            "10.200.0.14/32",
            "10.200.0.15/32",
            "10.200.0.16/32",
            "10.200.0.17/32",
            "10.200.0.18/32",
            "fd00:1::f/128",
            "fd00:1::10/128",
            "fd00:1::11/128",
            "fd00:1::12/128"
          ]
        }
      }
    }
    storage {
      lvp_share_config {
        lvp_config {
          path = "/mnt/localpv-share"
          storage_class = "local-shared"
        }
        shared_path_pv_count = 5
      }
      lvp_node_mounts_config {
        path = "/mnt/localpv-disk"
        storage_class = "local-disks"
      }
    }
    security_config {
      authorization {
        admin_users {
          username = "admin@hashicorptest.com"
        }
      }
    }
  }

  resource "google_gkeonprem_bare_metal_node_pool" "nodepool" {
    provider = google-beta
    location = "us-west1"
    name = "nodepool"
    bare_metal_cluster = google_gkeonprem_bare_metal_cluster.cluster.name
    annotations = {}
    node_pool_config {
      operating_system = "LINUX"
      labels = {}
      node_configs {
        node_ip = "10.200.0.12"
        labels = {}
      }
    }
  }
`, context)
}

func testAccCheckGkeonpremBareMetalNodePoolDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gkeonprem_bare_metal_node_pool" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GkeonpremBasePath}}projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("GkeonpremBareMetalNodePool still exists at %s", url)
			}
		}

		return nil
	}
}
