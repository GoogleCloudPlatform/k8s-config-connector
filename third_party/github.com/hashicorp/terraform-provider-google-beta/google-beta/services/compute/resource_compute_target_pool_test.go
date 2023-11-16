// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeTargetPool_basic(t *testing.T) {
	t.Parallel()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeTargetPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeTargetPool_basic(acctest.RandString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeTargetPoolExists(
						t, "google_compute_target_pool.foo"),
					testAccCheckComputeTargetPoolHealthCheck("google_compute_target_pool.foo", "google_compute_http_health_check.foobar"),
					testAccCheckComputeTargetPoolExists(
						t, "google_compute_target_pool.bar"),
					testAccCheckComputeTargetPoolHealthCheck("google_compute_target_pool.bar", "google_compute_http_health_check.foobar"),
				),
			},
			{
				ResourceName:      "google_compute_target_pool.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeTargetPool_update(t *testing.T) {
	t.Parallel()

	tpname := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	name1 := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	name2 := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeTargetPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// Create target pool with no instances attached
				Config: testAccComputeTargetPool_update(tpname, "", name1, name2),
			},
			{
				ResourceName:      "google_compute_target_pool.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Add the two instances to the pool
				Config: testAccComputeTargetPool_update(tpname,
					`google_compute_instance.foo.self_link, google_compute_instance.bar.self_link`,
					name1, name2),
			},
			{
				ResourceName:      "google_compute_target_pool.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Reversing the order of instances or changing import format shouldn't matter
				Config: testAccComputeTargetPool_update(tpname,
					fmt.Sprintf(`google_compute_instance.bar.self_link, "us-central1-a/%s"`, name1),
					name1, name2),
				PlanOnly: true,
			},
		},
	})
}

func TestAccComputeTargetPool_withSecurityPolicy(t *testing.T) {
	tpname := fmt.Sprintf("tf-tp-test-%s", acctest.RandString(t, 10))
	ddosPolicy := fmt.Sprintf("tf-tp-ddos-pol-test-%s", acctest.RandString(t, 10))
	edgeSecService := fmt.Sprintf("tf-tp-edge-sec-test-%s", acctest.RandString(t, 10))
	pol1 := fmt.Sprintf("tf-tp-pol1-test-%s", acctest.RandString(t, 10))
	pol2 := fmt.Sprintf("tf-tp-pol2-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeTargetPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				// Create target pool with no security policy attached
				Config: testAccComputeTargetPool_withSecurityPolicy(ddosPolicy, edgeSecService, pol1, pol2, tpname, "\"\""),
			},
			{
				ResourceName:      "google_compute_target_pool.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Add the first security policy to the pool
				Config: testAccComputeTargetPool_withSecurityPolicy(ddosPolicy, edgeSecService, pol1, pol2, tpname,
					`google_compute_region_security_policy.policytargetpool1.self_link`),
			},
			{
				ResourceName:      "google_compute_target_pool.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Change to the second security policy in the pool
				Config: testAccComputeTargetPool_withSecurityPolicy(ddosPolicy, edgeSecService, pol1, pol2, tpname,
					`google_compute_region_security_policy.policytargetpool2.self_link`),
			},
			{
				ResourceName:      "google_compute_target_pool.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Clean the security policy from the pool
				Config: testAccComputeTargetPool_withSecurityPolicy(ddosPolicy, edgeSecService, pol1, pol2, tpname, "\"\""),
			},
			{
				ResourceName:      "google_compute_target_pool.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckComputeTargetPoolDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		config := acctest.GoogleProviderConfig(t)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_target_pool" {
				continue
			}

			_, err := config.NewComputeClient(config.UserAgent).TargetPools.Get(
				config.Project, config.Region, rs.Primary.Attributes["name"]).Do()
			if err == nil {
				return fmt.Errorf("TargetPool still exists")
			}
		}

		return nil
	}
}

func testAccCheckComputeTargetPoolExists(t *testing.T, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := acctest.GoogleProviderConfig(t)

		found, err := config.NewComputeClient(config.UserAgent).TargetPools.Get(
			config.Project, config.Region, rs.Primary.Attributes["name"]).Do()
		if err != nil {
			return err
		}

		if found.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("TargetPool not found")
		}

		return nil
	}
}

func testAccCheckComputeTargetPoolHealthCheck(targetPool, healthCheck string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		targetPoolRes, ok := s.RootModule().Resources[targetPool]
		if !ok {
			return fmt.Errorf("Not found: %s", targetPool)
		}

		healthCheckRes, ok := s.RootModule().Resources[healthCheck]
		if !ok {
			return fmt.Errorf("Not found: %s", healthCheck)
		}

		hcLink := healthCheckRes.Primary.Attributes["self_link"]
		if tpgresource.ConvertSelfLinkToV1(targetPoolRes.Primary.Attributes["health_checks.0"]) != hcLink {
			return fmt.Errorf("Health check not set up. Expected %q to equal %q", tpgresource.ConvertSelfLinkToV1(targetPoolRes.Primary.Attributes["health_checks.0"]), hcLink)
		}

		return nil
	}
}

func testAccComputeTargetPool_basic(suffix string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_http_health_check" "foobar" {
  name = "healthcheck-test-%s"
  host = "example.com"
}

resource "google_compute_instance" "foobar" {
  name         = "tf-test-%s"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_target_pool" "foo" {
  description      = "Resource created for Terraform acceptance testing"
  instances        = [google_compute_instance.foobar.self_link, "us-central1-b/bar"]
  name             = "tpool-test-%s"
  session_affinity = "CLIENT_IP_PROTO"
  health_checks = [
    google_compute_http_health_check.foobar.name,
  ]
}

resource "google_compute_target_pool" "bar" {
  description = "Resource created for Terraform acceptance testing"
  name        = "tpool-test-2-%s"
  health_checks = [
    google_compute_http_health_check.foobar.self_link,
  ]
}
`, suffix, suffix, suffix, suffix)
}

func testAccComputeTargetPool_update(tpname, instances, name1, name2 string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foo" {
  description = "Resource created for Terraform acceptance testing"
  name        = "%s"
  instances   = [%s]
}

resource "google_compute_instance" "foo" {
  name         = "%s"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance" "bar" {
  name         = "%s"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}
`, tpname, instances, name1, name2)
}

func testAccComputeTargetPool_withSecurityPolicy(ddosPolicy, edgeSecService, pol1, pol2, tpname, polToSet string) string {
	return fmt.Sprintf(`
resource "google_compute_region_security_policy" "policyddosprotection" {
  region      = "us-south1"
  name        = "%s"
  description = "region security policy for load balancers target pool"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}

resource "google_compute_network_edge_security_service" "edge_sec_service" {
  name            = "%s"
  region          = "us-south1"
  description     = "edge security service with security policy"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}

resource "google_compute_region_security_policy" "policytargetpool1" {
  region      = "us-south1"
  name        = "%s"
  description = "region security policy one"
  type        = "CLOUD_ARMOR_NETWORK"
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}

resource "google_compute_region_security_policy" "policytargetpool2" {
  region      = "us-south1"
  name        = "%s"
  description = "region security policy two"
  type        = "CLOUD_ARMOR_NETWORK"
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}

resource "google_compute_target_pool" "foo" {
  region          = "us-south1"
  description     = "Setting SecurityPolicy to targetPool"
  name            = "%s"
  security_policy = %s
}
`, ddosPolicy, edgeSecService, pol1, pol2, tpname, polToSet)
}
