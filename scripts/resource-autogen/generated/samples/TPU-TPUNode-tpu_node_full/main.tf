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
data "google_tpu_tensorflow_versions" "available" {
}


resource "google_tpu_node" "tpu" {
  name = "test-tpu"
  zone = "us-central1-b"

  accelerator_type = "v3-8"

  tensorflow_version = data.google_tpu_tensorflow_versions.available.versions[0]

  description = "Terraform Google Provider test TPU"
  use_service_networking = true
  network = google_service_networking_connection.private_service_connection.network

  labels = {
    foo = "bar"
  }

  scheduling_config {
    preemptible = true
  }
}

data "google_compute_network" "network" {
  name = "default"
}

resource "google_compute_global_address" "service_range" {
  name          = "my-global-address"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = data.google_compute_network.network.id
}

resource "google_service_networking_connection" "private_service_connection" {
  network                 = data.google_compute_network.network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.service_range.name]
}
```
