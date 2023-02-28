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
resource "google_compute_network" "default" {
  name = "psc-network"
}

resource "google_compute_subnetwork" "default" {
  name          = "psc-subnetwork"
  ip_cidr_range = "10.0.0.0/16"
  region        = "europe-west4"
  network       = google_compute_network.default.id
}

resource "google_compute_subnetwork" "psc_subnetwork" {
  name          = "psc-subnetwork-nat"
  ip_cidr_range = "10.1.0.0/16"
  region        = "europe-west4"
  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.default.id
}

resource "google_compute_health_check" "default" {
  name = "psc-healthcheck"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
resource "google_compute_region_backend_service" "default" {
  name   = "psc-backend"
  region = "europe-west4"

  health_checks = [google_compute_health_check.default.id]
}

resource "google_compute_forwarding_rule" "default" {
  name   = "psc-forwarding-rule"
  region = "europe-west4"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.default.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}

resource "google_compute_service_attachment" "default" {
  name        = "psc-service-attachment"
  region      = "europe-west4"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol = false
  connection_preference = "ACCEPT_AUTOMATIC"
  nat_subnets           = [google_compute_subnetwork.psc_subnetwork.self_link]
  target_service        = google_compute_forwarding_rule.default.self_link
}

resource "google_compute_region_network_endpoint_group" "psc_neg_service_attachment" {
  name                  = "psc-neg"
  region                = "europe-west4"

  network_endpoint_type = "PRIVATE_SERVICE_CONNECT"
  psc_target_service    = google_compute_service_attachment.default.self_link

  network               = google_compute_network.default.self_link
  subnetwork            = google_compute_subnetwork.default.self_link
}
```
