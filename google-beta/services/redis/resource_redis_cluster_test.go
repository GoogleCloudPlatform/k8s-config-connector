// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package redis_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

// Validate that replica count is updated for the cluster
func TestAccRedisCluster_updateReplicaCount(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with replica count 1
				Config: createOrUpdateRedisCluster(name /* replicaCount = */, 1 /* shardCount = */, 3, true),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update replica count to 2
				Config: createOrUpdateRedisCluster(name /* replicaCount = */, 2 /* shardCount = */, 3, true),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(name /* replicaCount = */, 2 /* shardCount = */, 3, false),
			},
		},
	})
}

// Validate that shard count is updated for the cluster
func TestAccRedisCluster_updateShardCount(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-test-%d", acctest.RandInt(t))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckRedisClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// create cluster with shard count 3
				Config: createOrUpdateRedisCluster(name /* replicaCount = */, 1 /* shardCount = */, 3, true),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// update shard count to 5
				Config: createOrUpdateRedisCluster(name /* replicaCount = */, 1 /* shardCount = */, 5, true),
			},
			{
				ResourceName:            "google_redis_cluster.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"psc_configs"},
			},
			{
				// clean up the resource
				Config: createOrUpdateRedisCluster(name /* replicaCount = */, 1 /* shardCount = */, 5, false),
			},
		},
	})
}

func createOrUpdateRedisCluster(name string, replicaCount int, shardCount int, preventDestroy bool) string {
	lifecycleBlock := ""
	if preventDestroy {
		lifecycleBlock = `
		lifecycle {
			prevent_destroy = true
		}`
	}
	return fmt.Sprintf(`
resource "google_redis_cluster" "test" {
        provider = google-beta
        name           = "%s"
	replica_count = %d
	shard_count = %d
  region         = "us-central1"
	psc_configs {
			network = google_compute_network.producer_net.id
	}
	depends_on = [
          google_network_connectivity_service_connection_policy.default
        ]
	%s
}

resource "google_network_connectivity_service_connection_policy" "default" {
  provider = google-beta
  name = "%s"
  location = "us-central1"
  service_class = "gcp-memorystore-redis"
  description   = "my basic service connection policy"
  network = google_compute_network.producer_net.id
  psc_config {
    subnetworks = [google_compute_subnetwork.producer_subnet.id]
  }
}

resource "google_compute_subnetwork" "producer_subnet" {
  provider      = google-beta
  name          = "%s"
  ip_cidr_range = "10.0.0.248/29"
  region        = "us-central1"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_network" "producer_net" {
  provider                = google-beta
  name                    = "%s"
  auto_create_subnetworks = false
}
`, name, replicaCount, shardCount, lifecycleBlock, name, name, name)
}
