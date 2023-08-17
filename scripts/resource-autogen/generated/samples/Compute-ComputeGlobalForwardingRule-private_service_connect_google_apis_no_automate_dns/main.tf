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
  provider      = google-beta
  project       = "my-project-name"
  name          = "my-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "vpc_subnetwork" {
  provider                 = google-beta
  project                  = google_compute_network.network.project
  name                     = "my-subnetwork"
  ip_cidr_range            = "10.2.0.0/16"
  region                   = "us-central1"
  network                  = google_compute_network.network.id
  private_ip_google_access = true
}

resource "google_compute_global_address" "default" {
  provider      = google-beta
  project       = google_compute_network.network.project
  name          = "global-psconnect-ip"
  address_type  = "INTERNAL"
  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.network.id
  address       = "100.100.100.106"
}

resource "google_compute_global_forwarding_rule" "default" {
  provider      = google-beta
  project       = google_compute_network.network.project
  name          = "globalrule"
  target        = "all-apis"
  network       = google_compute_network.network.id
  ip_address    = google_compute_global_address.default.id
  load_balancing_scheme = ""
  no_automate_dns_zone  = false
}
```
