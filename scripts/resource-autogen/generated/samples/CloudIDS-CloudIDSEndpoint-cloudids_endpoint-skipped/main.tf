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
	name = "tf-test-my-network"
}
resource "google_compute_global_address" "service_range" {
	name          = "address"
	purpose       = "VPC_PEERING"
	address_type  = "INTERNAL"
	prefix_length = 16
	network       = google_compute_network.default.id
}
resource "google_service_networking_connection" "private_service_connection" {
	network                 = google_compute_network.default.id
	service                 = "servicenetworking.googleapis.com"
	reserved_peering_ranges = [google_compute_global_address.service_range.name]
}

resource "google_cloud_ids_endpoint" "example-endpoint" {
    name     = "test"
    location = "us-central1-f"
    network  = google_compute_network.default.id
    severity = "INFORMATIONAL"
    depends_on = [google_service_networking_connection.private_service_connection]
}
```
