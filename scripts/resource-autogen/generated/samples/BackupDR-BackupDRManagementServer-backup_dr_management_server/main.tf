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
  provider = google-beta
  name = "vpc-network"
}

resource "google_compute_global_address" "private_ip_address" {
  provider = google-beta
  name          = "vpc-network"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 20
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "default" {
  provider = google-beta
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name]
}

resource "google_backup_dr_management_server" "ms-console" {
  provider = google-beta
  location = "us-central1"
  name     = "ms-console"
  type     = "BACKUP_RESTORE" 
  networks {
    network      = google_compute_network.default.id
    peering_mode = "PRIVATE_SERVICE_ACCESS"
  }
  depends_on = [ google_service_networking_connection.default ]
}
```
