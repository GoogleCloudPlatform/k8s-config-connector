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
resource "google_compute_ha_vpn_gateway" "vpn-gateway" {
  name           = "test-ha-vpngw"
  network        = google_compute_network.network.id
  vpn_interfaces {
      id                      = 0
      interconnect_attachment = google_compute_interconnect_attachment.attachment1.self_link
  }
  vpn_interfaces {
      id                      = 1
      interconnect_attachment = google_compute_interconnect_attachment.attachment2.self_link
  }
}

resource "google_compute_interconnect_attachment" "attachment1" {
  name                     = "test-interconnect-attachment1"
  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
  type                     = "PARTNER"
  router                   = google_compute_router.router.id
  encryption               = "IPSEC"
  ipsec_internal_addresses = [
    google_compute_address.address1.self_link,
  ]
}

resource "google_compute_interconnect_attachment" "attachment2" {
  name                     = "test-interconnect-attachment2"
  edge_availability_domain = "AVAILABILITY_DOMAIN_2"
  type                     = "PARTNER"
  router                   = google_compute_router.router.id
  encryption               = "IPSEC"
  ipsec_internal_addresses = [
    google_compute_address.address2.self_link,
  ]
}

resource "google_compute_address" "address1" {
  name          = "test-address1"
  address_type  = "INTERNAL"
  purpose       = "IPSEC_INTERCONNECT"
  address       = "192.168.1.0"
  prefix_length = 29
  network       = google_compute_network.network.self_link
}

resource "google_compute_address" "address2" {
  name          = "test-address2"
  address_type  = "INTERNAL"
  purpose       = "IPSEC_INTERCONNECT"
  address       = "192.168.2.0"
  prefix_length = 29
  network       = google_compute_network.network.self_link
}

resource "google_compute_router" "router" {
  name                          = "test-router"
  network                       = google_compute_network.network.name
  encrypted_interconnect_router = true
  bgp {
    asn = 16550
  }
}

resource "google_compute_network" "network" {
  name                    = "test-network"
  auto_create_subnetworks = false
}
```
