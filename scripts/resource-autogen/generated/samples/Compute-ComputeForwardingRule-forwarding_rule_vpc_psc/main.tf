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
// Forwarding rule for VPC private service connect
resource "google_compute_forwarding_rule" "default" {
  provider              = google-beta
  name                  = "psc-endpoint"
  region                = "us-central1"
  load_balancing_scheme = ""
  target                = google_compute_service_attachment.producer_service_attachment.id
  network               = google_compute_network.consumer_net.name
  ip_address            = google_compute_address.consumer_address.id
}

// Consumer service endpoint

resource "google_compute_network" "consumer_net" {
  provider                = google-beta
  name                    = "consumer-net"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "consumer_subnet" {
  provider      = google-beta
  name          = "consumer-net"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_address" "consumer_address" {
  name         = "website-ip-1"
  provider     = google-beta
  region       = "us-central1"
  subnetwork   = google_compute_subnetwork.consumer_subnet.id
  address_type = "INTERNAL"
}


// Producer service attachment

resource "google_compute_network" "producer_net" {
  provider                = google-beta
  name                    = "producer-net"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "producer_subnet" {
  provider      = google-beta
  name          = "producer-net"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_subnetwork" "psc_producer_subnet" {
  provider      = google-beta
  name          = "producer-psc-net"
  ip_cidr_range = "10.1.0.0/16"
  region        = "us-central1"

  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_service_attachment" "producer_service_attachment" {
  provider    = google-beta
  name        = "producer-service"
  region      = "us-central1"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol = true
  connection_preference = "ACCEPT_AUTOMATIC"
  nat_subnets           = [google_compute_subnetwork.psc_producer_subnet.name]
  target_service        = google_compute_forwarding_rule.producer_target_service.id


}

resource "google_compute_forwarding_rule" "producer_target_service" {
  provider = google-beta
  name     = "producer-forwarding-rule"
  region   = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.producer_service_backend.id
  all_ports             = true
  network               = google_compute_network.producer_net.name
  subnetwork            = google_compute_subnetwork.producer_subnet.name


}

resource "google_compute_region_backend_service" "producer_service_backend" {
  provider = google-beta
  name     = "producer-service-backend"
  region   = "us-central1"

  health_checks = [google_compute_health_check.producer_service_health_check.id]
}

resource "google_compute_health_check" "producer_service_health_check" {
  provider = google-beta
  name     = "producer-service-health-check"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
```
