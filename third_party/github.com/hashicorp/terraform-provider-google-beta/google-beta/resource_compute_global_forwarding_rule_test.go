// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccComputeGlobalForwardingRule_updateTarget(t *testing.T) {
	t.Parallel()

	fr := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	proxy := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	proxyUpdated := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	backend := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	hc := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_httpProxy(fr, "proxy", proxy, proxyUpdated, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"google_compute_global_forwarding_rule.forwarding_rule", "target", regexp.MustCompile(proxy+"$")),
				),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"port_range", "target"},
			},
			{
				Config: testAccComputeGlobalForwardingRule_httpProxy(fr, "proxy2", proxy, proxyUpdated, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"google_compute_global_forwarding_rule.forwarding_rule", "target", regexp.MustCompile(proxyUpdated+"$")),
				),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"port_range", "target"},
			},
		},
	})
}

func TestAccComputeGlobalForwardingRule_ipv6(t *testing.T) {
	t.Parallel()

	fr := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	proxy := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	backend := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	hc := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_ipv6(fr, proxy, backend, hc, urlmap),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"google_compute_global_forwarding_rule.forwarding_rule", "ip_version", "IPV6"),
				),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"port_range", "target"},
			},
		},
	})
}

func TestAccComputeGlobalForwardingRule_labels(t *testing.T) {
	t.Parallel()

	fr := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	proxy := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	backend := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	hc := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))
	urlmap := fmt.Sprintf("forwardrule-test-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_labels(fr, proxy, backend, hc, urlmap),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"port_range", "target"},
			},
			{
				Config: testAccComputeGlobalForwardingRule_labelsUpdated(fr, proxy, backend, hc, urlmap),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"port_range", "target"},
			},
		},
	})
}

func TestAccComputeGlobalForwardingRule_internalLoadBalancing(t *testing.T) {
	t.Parallel()

	fr := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	proxy := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	backend := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	hc := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	urlmap := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	igm := fmt.Sprintf("tf-test-%s", RandString(t, 10))
	it := fmt.Sprintf("tf-test-%s", RandString(t, 10))

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_internalLoadBalancing(fr, proxy, backend, hc, urlmap, igm, it),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"port_range", "target"},
			},
			{
				Config: testAccComputeGlobalForwardingRule_internalLoadBalancingUpdate(fr, proxy, backend, hc, urlmap, igm, it),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.forwarding_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_httpProxy(fr, targetProxy, proxy, proxy2, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "TCP"
  name        = "%s"
  port_range  = "80"
  target      = google_compute_target_http_proxy.%s.self_link
}

resource "google_compute_target_http_proxy" "proxy" {
  description = "Resource created for Terraform acceptance testing"
  name        = "%s"
  url_map     = google_compute_url_map.url_map.self_link
}

resource "google_compute_target_http_proxy" "proxy2" {
  description = "Resource created for Terraform acceptance testing"
  name        = "%s"
  url_map     = google_compute_url_map.url_map.self_link
}

resource "google_compute_backend_service" "backend" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_url_map" "url_map" {
  name            = "%s"
  default_service = google_compute_backend_service.backend.self_link
  host_rule {
    hosts        = ["mysite.com", "myothersite.com"]
    path_matcher = "boop"
  }
  path_matcher {
    default_service = google_compute_backend_service.backend.self_link
    name            = "boop"
    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.backend.self_link
    }
  }
  test {
    host    = "mysite.com"
    path    = "/*"
    service = google_compute_backend_service.backend.self_link
  }
}
`, fr, targetProxy, proxy, proxy2, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_labels(fr, proxy, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  name       = "%s"
  target     = google_compute_target_http_proxy.proxy.self_link
  port_range = "80"

  labels = {
    my-label          = "a-value"
    a-different-label = "my-second-label-value"
  }
}

resource "google_compute_target_http_proxy" "proxy" {
  description = "Resource created for Terraform acceptance testing"
  name        = "%s"
  url_map     = google_compute_url_map.urlmap.self_link
}

resource "google_compute_backend_service" "backend" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_url_map" "urlmap" {
  name            = "%s"
  default_service = google_compute_backend_service.backend.self_link
  host_rule {
    hosts        = ["mysite.com", "myothersite.com"]
    path_matcher = "boop"
  }
  path_matcher {
    default_service = google_compute_backend_service.backend.self_link
    name            = "boop"
    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.backend.self_link
    }
  }
  test {
    host    = "mysite.com"
    path    = "/*"
    service = google_compute_backend_service.backend.self_link
  }
}
`, fr, proxy, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_labelsUpdated(fr, proxy, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  name       = "%s"
  target     = google_compute_target_http_proxy.proxy.self_link
  port_range = "80"

  labels = {
    my-label          = "a-new-value"
    a-different-label = "my-third-label-value"
  }
}

resource "google_compute_target_http_proxy" "proxy" {
  description = "Resource created for Terraform acceptance testing"
  name        = "%s"
  url_map     = google_compute_url_map.urlmap.self_link
}

resource "google_compute_backend_service" "backend" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_url_map" "urlmap" {
  name            = "%s"
  default_service = google_compute_backend_service.backend.self_link
  host_rule {
    hosts        = ["mysite.com", "myothersite.com"]
    path_matcher = "boop"
  }
  path_matcher {
    default_service = google_compute_backend_service.backend.self_link
    name            = "boop"
    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.backend.self_link
    }
  }
  test {
    host    = "mysite.com"
    path    = "/*"
    service = google_compute_backend_service.backend.self_link
  }
}
`, fr, proxy, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_ipv6(fr, proxy, backend, hc, urlmap string) string {
	return fmt.Sprintf(`
resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  description = "Resource created for Terraform acceptance testing"
  ip_protocol = "TCP"
  name        = "%s"
  port_range  = "80"
  target      = google_compute_target_http_proxy.proxy.self_link
  ip_version  = "IPV6"
}

resource "google_compute_target_http_proxy" "proxy" {
  description = "Resource created for Terraform acceptance testing"
  name        = "%s"
  url_map     = google_compute_url_map.urlmap.self_link
}

resource "google_compute_backend_service" "backend" {
  name          = "%s"
  health_checks = [google_compute_http_health_check.zero.self_link]
}

resource "google_compute_http_health_check" "zero" {
  name               = "%s"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_url_map" "urlmap" {
  name            = "%s"
  default_service = google_compute_backend_service.backend.self_link
  host_rule {
    hosts        = ["mysite.com", "myothersite.com"]
    path_matcher = "boop"
  }
  path_matcher {
    default_service = google_compute_backend_service.backend.self_link
    name            = "boop"
    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.backend.self_link
    }
  }
  test {
    host    = "mysite.com"
    path    = "/*"
    service = google_compute_backend_service.backend.self_link
  }
}
`, fr, proxy, backend, hc, urlmap)
}

func testAccComputeGlobalForwardingRule_internalLoadBalancing(fr, proxy, backend, hc, urlmap, igm, it string) string {
	return fmt.Sprintf(`
resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  name                  = "%s"
  target                = google_compute_target_http_proxy.default.self_link
  port_range            = "8080"
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  ip_address            = "0.0.0.0"
  metadata_filters {
    filter_match_criteria = "MATCH_ANY"
    filter_labels {
      name  = "PLANET"
      value = "NEPTUNE"
    }
  }
}

resource "google_compute_target_http_proxy" "default" {
  name        = "%s"
  description = "a description"
  url_map     = google_compute_url_map.default.self_link
}

resource "google_compute_backend_service" "default" {
  name                  = "%s"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"

  backend {
    group                 = google_compute_instance_group_manager.igm.instance_group
    balancing_mode        = "RATE"
    capacity_scaler       = 0.4
    max_rate_per_instance = 50
  }

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_health_check" "default" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_url_map" "default" {
  name            = "%s"
  description     = "a description"
  default_service = google_compute_backend_service.default.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.self_link

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.self_link
    }
  }
}

data "google_compute_image" "debian_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_instance_group_manager" "igm" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.instance_template.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-internal-igm"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "instance_template" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }
}
`, fr, proxy, backend, hc, urlmap, igm, it)
}

func testAccComputeGlobalForwardingRule_internalLoadBalancingUpdate(fr, proxy, backend, hc, urlmap, igm, it string) string {
	return fmt.Sprintf(`
resource "google_compute_global_forwarding_rule" "forwarding_rule" {
  name                  = "%s"
  target                = google_compute_target_http_proxy.default.self_link
  port_range            = "8080"
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
  ip_address            = "0.0.0.0"
  metadata_filters {
    filter_match_criteria = "MATCH_ANY"
    filter_labels {
      name  = "PLANET"
      value = "NEPTUNE"
    }
    filter_labels {
      name  = "PLANET"
      value = "JUPITER"
    }
  }
  metadata_filters {
    filter_match_criteria = "MATCH_ALL"
    filter_labels {
      name  = "STAR"
      value = "PROXIMA CENTAURI"
    }
    filter_labels {
      name  = "SPECIES"
      value = "ALIEN"
    }
  }
}

resource "google_compute_target_http_proxy" "default" {
  name        = "%s"
  description = "a description"
  url_map     = google_compute_url_map.default.self_link
}

resource "google_compute_backend_service" "default" {
  name                  = "%s"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"

  backend {
    group                 = google_compute_instance_group_manager.igm.instance_group
    balancing_mode        = "RATE"
    capacity_scaler       = 0.4
    max_rate_per_instance = 50
  }

  health_checks = [google_compute_health_check.default.self_link]
}

resource "google_compute_health_check" "default" {
  name               = "%s"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_url_map" "default" {
  name            = "%s"
  description     = "a description"
  default_service = google_compute_backend_service.default.self_link

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.self_link

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.self_link
    }
  }
}

data "google_compute_image" "debian_image" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_instance_group_manager" "igm" {
  name = "%s"
  version {
    instance_template = google_compute_instance_template.instance_template.self_link
    name              = "primary"
  }
  base_instance_name = "tf-test-internal-igm"
  zone               = "us-central1-f"
  target_size        = 1
}

resource "google_compute_instance_template" "instance_template" {
  name         = "%s"
  machine_type = "e2-medium"

  network_interface {
    network = "default"
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }
}
`, fr, proxy, backend, hc, urlmap, igm, it)
}
