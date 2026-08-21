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
resource "google_dns_managed_zone" "peering-zone" {
  name        = "peering-zone"
  dns_name    = "peering.example.com."
  description = "Example private DNS peering zone"

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network-source.id
    }
  }

  peering_config {
    target_network {
      network_url = google_compute_network.network-target.id
    }
  }
}

resource "google_compute_network" "network-source" {
  name                    = "network-source"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-target" {
  name                    = "network-target"
  auto_create_subnetworks = false
}
```
