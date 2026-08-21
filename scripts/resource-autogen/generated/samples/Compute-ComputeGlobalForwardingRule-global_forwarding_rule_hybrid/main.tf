/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
// Roughly mirrors https://cloud.google.com/load-balancing/docs/https/setting-up-ext-https-hybrid
variable "subnetwork_cidr" {
  default = "10.0.0.0/24"
}

resource "google_compute_network" "default" {
  name                    = "my-network"
}

resource "google_compute_network" "internal" {
  name                    = "my-internal-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "internal"{
  name                    = "my-subnetwork"
  network                 = google_compute_network.internal.id
  ip_cidr_range           = var.subnetwork_cidr
  region                  = "us-central1"
  private_ip_google_access= true
}

// Zonal NEG with GCE_VM_IP_PORT
resource "google_compute_network_endpoint_group" "default" {
  name                  = "default-neg"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "GCE_VM_IP_PORT"
}

// Zonal NEG with GCE_VM_IP
resource "google_compute_network_endpoint_group" "internal" {
  name                  = "internal-neg"
  network               = google_compute_network.internal.id
  subnetwork            = google_compute_subnetwork.internal.id
  zone                  = "us-central1-a"
  network_endpoint_type = "GCE_VM_IP"
}

// Hybrid connectivity NEG
resource "google_compute_network_endpoint_group" "hybrid" {
  name                  = "hybrid-neg"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "NON_GCP_PRIVATE_IP_PORT"
}

resource "google_compute_network_endpoint" "hybrid-endpoint" {
  network_endpoint_group = google_compute_network_endpoint_group.hybrid.name
  port       = google_compute_network_endpoint_group.hybrid.default_port
  ip_address = "127.0.0.1"
}

// Backend service for Zonal NEG
resource "google_compute_backend_service" "default" {
  name                  = "backend-default"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  backend {
    group = google_compute_network_endpoint_group.default.id
    balancing_mode               = "RATE"
    max_rate_per_endpoint        = 10
  }
  health_checks = [google_compute_health_check.default.id]
}

// Backgend service for Hybrid NEG
resource "google_compute_backend_service" "hybrid" {
  name                  = "backend-hybrid"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  backend {
    group                        = google_compute_network_endpoint_group.hybrid.id
    balancing_mode               = "RATE"
    max_rate_per_endpoint = 10
  }
  health_checks = [google_compute_health_check.default.id]
}

resource "google_compute_health_check" "default" {
  name               = "health-check"
  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_url_map" "default" {
  name            = "url-map-target-proxy"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }

    path_rule {
      paths   = ["/hybrid"]
      service = google_compute_backend_service.hybrid.id
    }
  }
}

resource "google_compute_target_http_proxy" "default" {
  name        = "target-proxy"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_global_forwarding_rule" "default" {
  name       = "global-rule"
  target     = google_compute_target_http_proxy.default.id
  port_range = "80"
}
```
