// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeForwardingRule_update(t *testing.T) {
	t.Parallel()

	poolName := fmt.Sprintf("tf-%s", acctest.RandString(t, 10))
	ruleName := fmt.Sprintf("tf-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_basic(poolName, ruleName),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeForwardingRule_update(poolName, ruleName),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeForwardingRule_ip(t *testing.T) {
	t.Parallel()

	addrName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	poolName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	ruleName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	addressRefFieldRaw := "address"
	addressRefFieldID := "id"

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_ip(addrName, poolName, ruleName, addressRefFieldID),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ip_address"}, // ignore ip_address because we've specified it by ID
			},
			{
				Config: testAccComputeForwardingRule_ip(addrName, poolName, ruleName, addressRefFieldRaw),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeForwardingRule_internalTcpUdpLbWithMigBackendExampleUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_internalTcpUdpLbWithMigBackendExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.google_compute_forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
			{
				Config: testAccComputeForwardingRule_internalTcpUdpLbWithMigBackendExampleUpdate(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.google_compute_forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
		},
	})
}

func TestAccComputeForwardingRule_networkTier(t *testing.T) {
	t.Parallel()

	poolName := fmt.Sprintf("tf-%s", acctest.RandString(t, 10))
	ruleName := fmt.Sprintf("tf-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_networkTier(poolName, ruleName),
			},

			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeForwardingRule_serviceDirectoryRegistrations(t *testing.T) {
	t.Parallel()

	poolName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	ruleName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	svcDirNamespace := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_serviceDirectoryRegistrations(poolName, ruleName, svcDirNamespace, serviceName),
			},

			{
				ResourceName:      "google_compute_forwarding_rule.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeForwardingRule_forwardingRuleVpcPscExampleUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleVpcPscExample(context),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeForwardingRule_forwardingRuleVpcPscExampleUpdate(context, true),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeForwardingRule_forwardingRuleVpcPscExampleUpdate(context, false),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeForwardingRule_forwardingRuleRegionalSteeringExampleUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleRegionalSteeringExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.steering",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
			{
				Config: testAccComputeForwardingRule_forwardingRuleRegionalSteeringExampleUpdate(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.steering",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_basic(poolName, ruleName string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foo-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "foo-%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "UDP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.foo-tp.self_link
  labels = {
    "foo" = "bar"
  }
}
`, poolName, ruleName)
}

func testAccComputeForwardingRule_update(poolName, ruleName string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foo-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "foo-%s"
}

resource "google_compute_target_pool" "bar-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "bar-%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "UDP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.bar-tp.self_link
  labels = {
    "baz" = "qux"
  }
}
`, poolName, poolName, ruleName)
}

func testAccComputeForwardingRule_ip(addrName, poolName, ruleName, addressRefFieldValue string) string {
	return fmt.Sprintf(`
resource "google_compute_address" "foo" {
  name = "%s"
}

resource "google_compute_target_pool" "foobar-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_address  = google_compute_address.foo.%s
  ip_protocol = "TCP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.foobar-tp.self_link
}
`, addrName, poolName, addressRefFieldValue, ruleName)
}

func testAccComputeForwardingRule_internalTcpUdpLbWithMigBackendExampleUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
# Internal TCP/UDP load balancer with a managed instance group backend

# VPC
resource "google_compute_network" "ilb_network" {
  name                    = "tf-test-l4-ilb-network%{random_suffix}"
  provider                = google-beta
  auto_create_subnetworks = false
}

# backed subnet
resource "google_compute_subnetwork" "ilb_subnet" {
  name          = "tf-test-l4-ilb-subnet%{random_suffix}"
  provider      = google-beta
  ip_cidr_range = "10.0.1.0/24"
  region        = "europe-west1"
  network       = google_compute_network.ilb_network.id
}

# forwarding rule
resource "google_compute_forwarding_rule" "google_compute_forwarding_rule" {
  name                  = "tf-test-l4-ilb-forwarding-rule%{random_suffix}"
  backend_service       = google_compute_region_backend_service.default.id
  provider              = google-beta
  region                = "europe-west1"
  ip_protocol           = "TCP"
  load_balancing_scheme = "INTERNAL"
  all_ports             = true
  allow_global_access   = false
  network               = google_compute_network.ilb_network.id
  subnetwork            = google_compute_subnetwork.ilb_subnet.id
}

# backend service
resource "google_compute_region_backend_service" "default" {
  name                  = "tf-test-l4-ilb-backend-subnet%{random_suffix}"
  provider              = google-beta
  region                = "europe-west1"
  protocol              = "TCP"
  load_balancing_scheme = "INTERNAL"
  health_checks         = [google_compute_region_health_check.default.id]
  backend {
    group           = google_compute_region_instance_group_manager.mig.instance_group
    balancing_mode  = "CONNECTION"
  }
}

# instance template
resource "google_compute_instance_template" "instance_template" {
  name         = "tf-test-l4-ilb-mig-template%{random_suffix}"
  provider     = google-beta
  machine_type = "e2-small"
  tags         = ["allow-ssh","allow-health-check"]

  network_interface {
    network    = google_compute_network.ilb_network.id
    subnetwork = google_compute_subnetwork.ilb_subnet.id
    access_config {
      # add external ip to fetch packages
    }
  }
  disk {
    source_image = "debian-cloud/debian-10"
    auto_delete  = true
    boot         = true
  }

  # install nginx and serve a simple web page
  metadata = {
    startup-script = <<-EOF1
      #! /bin/bash
      set -euo pipefail

      export DEBIAN_FRONTEND=noninteractive
      apt-get update
      apt-get install -y nginx-light jq

      NAME=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/hostname")
      IP=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip")
      METADATA=$(curl -f -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/attributes/?recursive=True" | jq 'del(.["startup-script"])')

      cat <<EOF > /var/www/html/index.html
      <pre>
      Name: $NAME
      IP: $IP
      Metadata: $METADATA
      </pre>
      EOF
    EOF1
  }
  lifecycle {
    create_before_destroy = true
  }
}

# health check
resource "google_compute_region_health_check" "default" {
  name     = "tf-test-l4-ilb-hc%{random_suffix}"
  provider = google-beta
  region   = "europe-west1"
  http_health_check {
    port = "80"
  }
}

# MIG
resource "google_compute_region_instance_group_manager" "mig" {
  name     = "tf-test-l4-ilb-mig1%{random_suffix}"
  provider = google-beta
  region   = "europe-west1"
  version {
    instance_template = google_compute_instance_template.instance_template.id
    name              = "primary"
  }
  base_instance_name = "vm"
  target_size        = 2
}

# allow all access from health check ranges
resource "google_compute_firewall" "fw_hc" {
  name          = "tf-test-l4-ilb-fw-allow-hc%{random_suffix}"
  provider      = google-beta
  direction     = "INGRESS"
  network       = google_compute_network.ilb_network.id
  source_ranges = ["130.211.0.0/22", "35.191.0.0/16", "35.235.240.0/20"]
  allow {
    protocol = "tcp"
  }
  target_tags = ["allow-health-check"]
}

# allow communication within the subnet 
resource "google_compute_firewall" "fw_ilb_to_backends" {
  name          = "tf-test-l4-ilb-fw-allow-ilb-to-backends%{random_suffix}"
  provider      = google-beta
  direction     = "INGRESS"
  network       = google_compute_network.ilb_network.id
  source_ranges = ["10.0.1.0/24"]
  allow {
    protocol = "tcp"
  }
  allow {
    protocol = "udp"
  }
  allow {
    protocol = "icmp"
  }
}

# allow SSH
resource "google_compute_firewall" "fw_ilb_ssh" {
  name          = "tf-test-l4-ilb-fw-ssh%{random_suffix}"
  provider      = google-beta
  direction     = "INGRESS"
  network       = google_compute_network.ilb_network.id
  allow {
    protocol = "tcp"
    ports = ["22"]
  }
  target_tags   = ["allow-ssh"]
  source_ranges = ["0.0.0.0/0"]
}

# test instance
resource "google_compute_instance" "vm_test" {
  name         = "tf-test-l4-ilb-test-vm%{random_suffix}"
  provider     = google-beta
  zone         = "europe-west1-b"
  machine_type = "e2-small"
  network_interface {
    network    = google_compute_network.ilb_network.id
    subnetwork = google_compute_subnetwork.ilb_subnet.id
  }
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }
}
`, context)
}

func testAccComputeForwardingRule_networkTier(poolName, ruleName string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foobar-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "UDP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.foobar-tp.self_link

  network_tier = "STANDARD"
}
`, poolName, ruleName)
}

func testAccComputeForwardingRule_serviceDirectoryRegistrations(poolName, ruleName, svcDirNamespace, serviceName string) string {
	return fmt.Sprintf(`
resource "google_compute_target_pool" "foo-tp" {
  description = "Resource created for Terraform acceptance testing"
  instances   = ["us-central1-a/foo", "us-central1-b/bar"]
  name        = "foo-%s"
}

resource "google_compute_forwarding_rule" "foobar" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "UDP"
  name        = "%s"
  port_range  = "80-81"
  target      = google_compute_target_pool.foo-tp.self_link

  service_directory_registrations {
    namespace = google_service_directory_namespace.examplens.namespace_id
    service = google_service_directory_service.examplesvc.service_id
  }
}

resource "google_service_directory_namespace" "examplens" {
  namespace_id = "%s"
  location     = "us-central1"
}

resource "google_service_directory_service" "examplesvc" {
  service_id = "%s"
  namespace  = google_service_directory_namespace.examplens.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}
`, poolName, ruleName, svcDirNamespace, serviceName)
}

func testAccComputeForwardingRule_forwardingRuleVpcPscExampleUpdate(context map[string]interface{}, preventDestroy bool) string {
	context["lifecycle_block"] = ""
	if preventDestroy {
		context["lifecycle_block"] = `
		lifecycle {
			prevent_destroy = true
		}`
	}

	return acctest.Nprintf(`
// Forwarding rule for VPC private service connect
resource "google_compute_forwarding_rule" "default" {
  name                    = "tf-test-psc-endpoint%{random_suffix}"
  region                  = "us-central1"
  load_balancing_scheme   = ""
  target                  = google_compute_service_attachment.producer_service_attachment.id
  network                 = google_compute_network.consumer_net.name
  ip_address              = google_compute_address.consumer_address.id
  allow_psc_global_access = false
  %{lifecycle_block}
}

// Consumer service endpoint

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-consumer-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-consumer-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_address" "consumer_address" {
  name         = "tf-test-website-ip%{random_suffix}-1"
  region       = "us-central1"
  subnetwork   = google_compute_subnetwork.consumer_subnet.id
  address_type = "INTERNAL"
}


// Producer service attachment

resource "google_compute_network" "producer_net" {
  name                    = "tf-test-producer-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "producer_subnet" {
  name          = "tf-test-producer-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_subnetwork" "psc_producer_subnet" {
  name          = "tf-test-producer-psc-net%{random_suffix}"
  ip_cidr_range = "10.1.0.0/16"
  region        = "us-central1"

  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_service_attachment" "producer_service_attachment" {
  name        = "tf-test-producer-service%{random_suffix}"
  region      = "us-central1"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol = true
  connection_preference = "ACCEPT_AUTOMATIC"
  nat_subnets           = [google_compute_subnetwork.psc_producer_subnet.name]
  target_service        = google_compute_forwarding_rule.producer_target_service.id
}

resource "google_compute_forwarding_rule" "producer_target_service" {
  name     = "tf-test-producer-forwarding-rule%{random_suffix}"
  region   = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.producer_service_backend.id
  all_ports             = true
  network               = google_compute_network.producer_net.name
  subnetwork            = google_compute_subnetwork.producer_subnet.name
}

resource "google_compute_region_backend_service" "producer_service_backend" {
  name     = "tf-test-producer-service-backend%{random_suffix}"
  region   = "us-central1"

  health_checks = [google_compute_health_check.producer_service_health_check.id]
}

resource "google_compute_health_check" "producer_service_health_check" {
  name     = "tf-test-producer-service-health-check%{random_suffix}"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func testAccComputeForwardingRule_forwardingRuleRegionalSteeringExampleUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_forwarding_rule" "steering" {
  name = "tf-test-steering-rule%{random_suffix}"
  region = "us-central1"
  ip_address = google_compute_address.basic.self_link
  backend_service = google_compute_region_backend_service.external.self_link
  load_balancing_scheme = "EXTERNAL"
  source_ip_ranges = ["35.121.88.0/24", "36.187.239.137"]
  depends_on = [google_compute_forwarding_rule.external]
}

resource "google_compute_address" "basic" {
  name = "tf-test-website-ip%{random_suffix}"
  region = "us-central1"
}

resource "google_compute_region_backend_service" "external" {
  name = "tf-test-service-backend%{random_suffix}"
  region = "us-central1"
  load_balancing_scheme = "EXTERNAL"
}

resource "google_compute_forwarding_rule" "external" {
  name = "tf-test-external-forwarding-rule%{random_suffix}"
  region = "us-central1"
  ip_address = google_compute_address.basic.self_link
  backend_service = google_compute_region_backend_service.external.self_link
  load_balancing_scheme = "EXTERNAL"
}
`, context)
}
