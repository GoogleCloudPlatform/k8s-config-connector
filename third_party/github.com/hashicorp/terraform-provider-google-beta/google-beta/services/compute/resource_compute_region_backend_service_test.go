// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package compute_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeRegionBackendService_basic(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	extraCheckName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_basic(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_basicModified(
					serviceName, checkName, extraCheckName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_ilbBasic_withUnspecifiedProtocol(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_ilbBasic_withUnspecifiedProtocol(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_withBackendInternal(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	igName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	itName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_withInvalidInternalBackend(
					serviceName, igName, itName, checkName),
				ExpectError: regexp.MustCompile(`capacity_scaler" cannot be set for non-managed backend service`),
			},
			{
				Config: testAccComputeRegionBackendService_withBackend(
					serviceName, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_region_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_withBackend(
					serviceName, igName, itName, checkName, 20),
			},
			{
				ResourceName:      "google_compute_region_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_withBackendInternalManaged(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	igmName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	hcName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_internalManagedMultipleBackends(serviceName, igmName, hcName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_withBackendMultiNic(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	net1Name := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	net2Name := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	igName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	itName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_withBackendMultiNic(
					serviceName, net1Name, net2Name, igName, itName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_region_backend_service.lipsum",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_withConnectionDrainingAndUpdate(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_withConnectionDraining(serviceName, checkName, 10),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_basic(serviceName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_ilbUpdateBasic(t *testing.T) {
	t.Parallel()

	backendName := fmt.Sprintf("foo-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("bar-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_ilbBasic(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_ilbUpdateBasic(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_ilbUpdateFull(t *testing.T) {
	t.Parallel()

	randString := acctest.RandString(t, 10)

	backendName := fmt.Sprintf("foo-%s", randString)
	checkName := fmt.Sprintf("bar-%s", randString)
	igName := fmt.Sprintf("tf-test-%s", randString)
	instanceName := fmt.Sprintf("tf-test-%s", randString)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_ilbFull(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_ilbUpdateFull(backendName, igName, instanceName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_withBackendAndIAP(t *testing.T) {
	backendName := fmt.Sprintf("foo-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("bar-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_ilbBasicwithIAP(backendName, checkName),
			},
			{
				ResourceName:            "google_compute_region_backend_service.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret"},
			},
			{
				Config: testAccComputeRegionBackendService_ilbBasic(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_UDPFailOverPolicyUpdate(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	checkName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_UDPFailOverPolicyHasDrain(serviceName, "TCP", "true", checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_UDPFailOverPolicyHasDrain(serviceName, "TCP", "false", checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_UDPFailOverPolicy(serviceName, "UDP", "false", checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionBackendService_subsettingUpdate(t *testing.T) {
	t.Parallel()

	randString := acctest.RandString(t, 10)

	backendName := fmt.Sprintf("foo-%s", randString)
	checkName := fmt.Sprintf("bar-%s", randString)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_ilbWithSubsetting(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_ilbNoSubsetting(backendName, checkName),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionBackendService_ilbBasic_withUnspecifiedProtocol(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  protocol              = "UNSPECIFIED"
  load_balancing_scheme = "INTERNAL"
  region        = "us-central1"
}

resource "google_compute_health_check" "health_check" {
  name     = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeRegionBackendService_ilbBasic(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  port_name             = "http"
  protocol              = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy    = "RING_HASH"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos   = 1234
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  name     = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeRegionBackendService_ilbUpdateBasic(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  port_name             = "https"
  protocol              = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy    = "RANDOM"
  circuit_breakers {
    max_connections = 10
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  name     = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeRegionBackendService_ilbFull(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  port_name             = "http"
  protocol              = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy    = "MAGLEV"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos   = 1234
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  name     = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeRegionBackendService_ilbUpdateFull(serviceName, igName, instanceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  port_name             = "https"
  protocol              = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy    = "MAGLEV"
  backend {
    balancing_mode  = "UTILIZATION"
    capacity_scaler = 0.5
    description     = "The backend"
    group           = google_compute_instance_group.group.self_link
    max_rate        = 6
    max_utilization = 0.5
  }
  circuit_breakers {
    connect_timeout {
      seconds = 3
      nanos   = 4
    }
    max_connections             = 11
    max_requests_per_connection = 12
    max_pending_requests        = 13
    max_requests                = 14
    max_retries                 = 15
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 12
      }
      name = "mycookie2"
      path = "mycookie2/path"
    }
    minimum_ring_size = 16
  }
  log_config {
    enable      = true
    sample_rate = 0.5
  }
  outlier_detection {
    base_ejection_time {
      seconds = 0
      nanos   = 5
    }
    consecutive_errors                    = 1
    consecutive_gateway_failure           = 3
    enforcing_consecutive_errors          = 4
    enforcing_consecutive_gateway_failure = 5
    enforcing_success_rate                = 6
    interval {
      seconds = 7
    }
    max_ejection_percent        = 99
    success_rate_minimum_hosts  = 98
    success_rate_request_volume = 97
    success_rate_stdev_factor   = 1800
  }
}

resource "google_compute_instance_group" "group" {
  name      = "%s"
  instances = [google_compute_instance.ig_instance.self_link]

  named_port {
    name = "http"
    port = "8080"
  }

  named_port {
    name = "https"
    port = "8443"
  }
}

data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_instance" "ig_instance" {
  name         = "%s"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_health_check" "health_check" {
  name = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, igName, instanceName, checkName)
}

func testAccComputeRegionBackendService_UDPFailOverPolicy(serviceName, protocol, failover, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_health_check.zero.self_link]
  region        = "us-central1"

  protocol = "%s"
  failover_policy {
      # Disable connection drain on failover cannot be set when the protocol is UDP
      drop_traffic_if_unhealthy = "%s"
  }
}

resource "google_compute_health_check" "zero" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, serviceName, protocol, failover, checkName)
}

func testAccComputeRegionBackendService_UDPFailOverPolicyHasDrain(serviceName, protocol, failover, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_health_check.zero.self_link]
  region        = "us-central1"

  protocol = "%s"
  failover_policy {
      # Disable connection drain on failover cannot be set when the protocol is UDP
      drop_traffic_if_unhealthy = "%s"
      disable_connection_drain_on_failover = "%s"
  }
}

resource "google_compute_health_check" "zero" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, serviceName, protocol, failover, failover, checkName)
}

func testAccComputeRegionBackendService_basic(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_health_check.zero.self_link]
  region        = "us-central1"
}

resource "google_compute_health_check" "zero" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, serviceName, checkName)
}

func testAccComputeRegionBackendService_basicModified(serviceName, checkOne, checkTwo string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name          = "%s"
  health_checks = [google_compute_health_check.one.self_link]
  region        = "us-central1"
}

resource "google_compute_health_check" "zero" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = 443
  }
}

resource "google_compute_health_check" "one" {
  name               = "%s"
  check_interval_sec = 30
  timeout_sec        = 30

  tcp_health_check {
    port = 443
  }
}
`, serviceName, checkOne, checkTwo)
}

func testAccComputeRegionBackendService_withBackend(
	serviceName, igName, itName, checkName string, timeout int64) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_region_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  protocol    = "TCP"
  region      = "us-central1"
  timeout_sec = %v

  backend {
    group    = google_compute_instance_group_manager.foobar.instance_group
    failover = true
  }

  failover_policy {
    disable_connection_drain_on_failover = true
    drop_traffic_if_unhealthy            = true
    failover_ratio                       = 0.4
  }

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_health_check" "default" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = 443
  }
}
`, serviceName, timeout, igName, itName, checkName)
}

func testAccComputeRegionBackendService_withBackendMultiNic(
	serviceName, net1Name, net2Name, igName, itName, checkName string, timeout int64) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_region_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  protocol    = "TCP"
  region      = "us-central1"
  timeout_sec = %v

  backend {
    group = google_compute_instance_group_manager.foobar.instance_group
    failover = true
  }

  failover_policy {
    disable_connection_drain_on_failover = true
    drop_traffic_if_unhealthy            = true
    failover_ratio                       = 0.4
  }

  network = google_compute_network.network2.self_link

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_network" "network1" {
  name                            = "%s"
  auto_create_subnetworks         = false
}

resource "google_compute_subnetwork" "subnet1" {
  name                     = "%s"
  ip_cidr_range            = "10.0.1.0/24"
  region                   = "us-central1"
  private_ip_google_access = true
  network                  = google_compute_network.network1.self_link
}

resource "google_compute_network" "network2" {
  name                            = "%s"
  auto_create_subnetworks         = false
}

resource "google_compute_subnetwork" "subnet2" {
  name                     = "%s"
  ip_cidr_range            = "10.0.2.0/24"
  region                   = "us-central1"
  private_ip_google_access = true
  network                  = google_compute_network.network2.self_link
}

resource "google_compute_instance_group_manager" "foobar" {
  name               = "%s"
  version {
    instance_template  = google_compute_instance_template.foobar.self_link
    name               = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    subnetwork = google_compute_subnetwork.subnet1.self_link
  }

  network_interface {
    subnetwork = google_compute_subnetwork.subnet2.self_link
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_health_check" "default" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = 443
  }
}
`, serviceName, timeout, net1Name, net1Name, net2Name, net2Name, igName, itName, checkName)
}

func testAccComputeRegionBackendService_withInvalidInternalBackend(
	serviceName, igName, itName, checkName string) string {
	return fmt.Sprintf(`
data "google_compute_image" "my_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_region_backend_service" "lipsum" {
  name        = "%s"
  description = "Hello World 1234"
  port_name   = "http"
  protocol    = "TCP"
  region      = "us-central1"

  backend {
    group    = google_compute_instance_group_manager.foobar.instance_group
    capacity_scaler = 1.0
  }

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_instance_group_manager" "foobar" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.foobar.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-foobar"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "foobar" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.my_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_health_check" "default" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = 443
  }
}
`, serviceName, igName, itName, checkName)
}

func testAccComputeRegionBackendService_internalManagedMultipleBackends(serviceName, igmName, hcName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "default" {
  name        = "%s"
  load_balancing_scheme = "INTERNAL_MANAGED"

  backend {
    group          = google_compute_region_instance_group_manager.rigm1.instance_group
    balancing_mode = "UTILIZATION"
  }

  backend {
    group          = google_compute_region_instance_group_manager.rigm2.instance_group
    balancing_mode = "UTILIZATION"
    capacity_scaler = 1.0
  }

  region      = "us-central1"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.self_link]
}

data "google_compute_image" "debian_image" {
  family   = "debian-11"
  project  = "debian-cloud"
}

resource "google_compute_region_instance_group_manager" "rigm1" {
  name     = "%s-1"
  region   = "us-central1"
  version {
    instance_template = google_compute_instance_template.instance_template.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-internal-glb"
  target_size        = 1
}

resource "google_compute_region_instance_group_manager" "rigm2" {
  name     = "%s-2"
  region   = "us-central1"
  version {
    instance_template = google_compute_instance_template.instance_template.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-internal-glb"
  target_size        = 1
}

resource "google_compute_instance_template" "instance_template" {
  name         = "%s-template"
  machine_type = "e2-medium"

  network_interface {
    network    = "default"
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }
}

resource "google_compute_region_health_check" "default" {
  name   = "%s"
  region = "us-central1"

  http_health_check {
    port_specification = "USE_SERVING_PORT"
  }
}
`, serviceName, igmName, igmName, igmName, hcName)
}

func testAccComputeRegionBackendService_withConnectionDraining(serviceName, checkName string, drainingTimeout int64) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                            = "%s"
  health_checks                   = [google_compute_health_check.zero.self_link]
  region                          = "us-central1"
  connection_draining_timeout_sec = %v
}

resource "google_compute_health_check" "zero" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, serviceName, drainingTimeout, checkName)
}

func testAccComputeRegionBackendService_ilbBasicwithIAP(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  port_name             = "http"
  protocol              = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy    = "RING_HASH"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos   = 1234
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }

  iap {
    oauth2_client_id     = "test"
    oauth2_client_secret = "test"
  }
}

resource "google_compute_health_check" "health_check" {
  name     = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeRegionBackendService_ilbWithSubsetting(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  protocol              = "TCP"
  load_balancing_scheme = "INTERNAL"  
  subsetting {
	  policy = "CONSISTENT_HASH_SUBSETTING"
  }
}

resource "google_compute_health_check" "health_check" {
  name     = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func testAccComputeRegionBackendService_ilbNoSubsetting(serviceName, checkName string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  health_checks         = [google_compute_health_check.health_check.self_link]
  protocol              = "TCP"
  load_balancing_scheme = "INTERNAL"  
}

resource "google_compute_health_check" "health_check" {
  name     = "%s"
  http_health_check {
    port = 80
  }
}
`, serviceName, checkName)
}

func TestAccComputeRegionBackendService_withSecurityPolicy(t *testing.T) {
	t.Parallel()

	serviceName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))
	polName := fmt.Sprintf("tf-test-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_withSecurityPolicy(serviceName, polName, "google_compute_region_security_policy.policy.self_link"),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionBackendService_withSecurityPolicy(serviceName, polName, "\"\""),
			},
			{
				ResourceName:      "google_compute_region_backend_service.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionBackendService_withSecurityPolicy(serviceName, polName, polLink string) string {
	return fmt.Sprintf(`
resource "google_compute_region_backend_service" "foobar" {
  name                  = "%s"
  region                = "us-central1"
  security_policy       = %s
  load_balancing_scheme = "EXTERNAL_MANAGED"
}

resource "google_compute_region_security_policy" "policy" {
  name        = "%s"
  region      = "us-central1"
  description = "basic security policy"
  type        = "CLOUD_ARMOR"
}
`, serviceName, polLink, polName)
}
