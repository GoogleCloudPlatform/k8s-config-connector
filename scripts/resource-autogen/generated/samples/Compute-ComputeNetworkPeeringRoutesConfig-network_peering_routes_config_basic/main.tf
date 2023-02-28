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
resource "google_compute_network_peering_routes_config" "peering_primary_routes" {
  peering = google_compute_network_peering.peering_primary.name
  network = google_compute_network.network_primary.name

  import_custom_routes = true
  export_custom_routes = true
}

resource "google_compute_network_peering" "peering_primary" {
  name         = "primary-peering"
  network      = google_compute_network.network_primary.id
  peer_network = google_compute_network.network_secondary.id

  import_custom_routes = true
  export_custom_routes = true
}

resource "google_compute_network_peering" "peering_secondary" {
  name         = "secondary-peering"
  network      = google_compute_network.network_secondary.id
  peer_network = google_compute_network.network_primary.id
}

resource "google_compute_network" "network_primary" {
  name                    = "primary-network"
  auto_create_subnetworks = "false"
}

resource "google_compute_network" "network_secondary" {
  name                    = "secondary-network"
  auto_create_subnetworks = "false"
}
```
