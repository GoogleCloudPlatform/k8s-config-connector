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
resource "google_compute_network" "network" {
  name                    = "my-router-net"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnetwork" {
  name          = "my-router-sub"
  network       = google_compute_network.network.self_link
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
}

resource "google_compute_address" "addr_intf" {
  name         = "my-router-addr-intf"
  region       = google_compute_subnetwork.subnetwork.region
  subnetwork   = google_compute_subnetwork.subnetwork.id
  address_type = "INTERNAL"
}

resource "google_compute_address" "addr_intf_redundant" {
  name         = "my-router-addr-intf-red"
  region       = google_compute_subnetwork.subnetwork.region
  subnetwork   = google_compute_subnetwork.subnetwork.id
  address_type = "INTERNAL"
}

resource "google_compute_address" "addr_peer" {
  name         = "my-router-addr-peer"
  region       = google_compute_subnetwork.subnetwork.region
  subnetwork   = google_compute_subnetwork.subnetwork.id
  address_type = "INTERNAL"
}

resource "google_compute_instance" "instance" {
  name           = "router-appliance"
  zone           = "us-central1-a"
  machine_type   = "e2-medium"
  can_ip_forward = true

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network_ip = google_compute_address.addr_peer.address
    subnetwork = google_compute_subnetwork.subnetwork.self_link
  }
}

resource "google_network_connectivity_hub" "hub" {
  name = "my-router-hub"
}

resource "google_network_connectivity_spoke" "spoke" {
  name     = "my-router-spoke"
  location = google_compute_subnetwork.subnetwork.region
  hub      = google_network_connectivity_hub.hub.id

  linked_router_appliance_instances {
    instances {
      virtual_machine = google_compute_instance.instance.self_link
      ip_address      = google_compute_address.addr_peer.address
    }
    site_to_site_data_transfer = false
  }
}

resource "google_compute_router" "router" {
  name    = "my-router-router"
  region  = google_compute_subnetwork.subnetwork.region
  network = google_compute_network.network.self_link
  bgp {
    asn = 64514
  }
}

resource "google_compute_router_interface" "interface_redundant" {
  name               = "my-router-intf-red"
  region             = google_compute_router.router.region
  router             = google_compute_router.router.name
  subnetwork         = google_compute_subnetwork.subnetwork.self_link
  private_ip_address = google_compute_address.addr_intf_redundant.address
}

resource "google_compute_router_interface" "interface" {
  name                = "my-router-intf"
  region              = google_compute_router.router.region
  router              = google_compute_router.router.name
  subnetwork          = google_compute_subnetwork.subnetwork.self_link
  private_ip_address  = google_compute_address.addr_intf.address
  redundant_interface = google_compute_router_interface.interface_redundant.name
}

resource "google_compute_router_peer" "peer" {
  name                      = "my-router-peer"
  router                    = google_compute_router.router.name
  region                    = google_compute_router.router.region
  interface                 = google_compute_router_interface.interface.name
  router_appliance_instance = google_compute_instance.instance.self_link
  peer_asn                  = 65513
  peer_ip_address           = google_compute_address.addr_peer.address
}
```
