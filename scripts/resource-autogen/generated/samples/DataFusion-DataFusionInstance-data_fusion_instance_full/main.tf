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
resource "google_data_fusion_instance" "extended_instance" {
  name                          = "my-instance"
  description                   = "My Data Fusion instance"
  display_name                  = "My Data Fusion instance"
  region                        = "us-central1"
  type                          = "BASIC"
  enable_stackdriver_logging    = true
  enable_stackdriver_monitoring = true
  private_instance              = true
  dataproc_service_account      = data.google_app_engine_default_service_account.default.email

  labels = {
    example_key = "example_value"
  }

  network_config {
    network       = "default"
    ip_allocation = "${google_compute_global_address.private_ip_alloc.address}/${google_compute_global_address.private_ip_alloc.prefix_length}"
  }

  accelerators {
    accelerator_type = "CDC"
    state = "ENABLED"
  }
  
}

data "google_app_engine_default_service_account" "default" {
}

resource "google_compute_network" "network" {
  name = "datafusion-full-network"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "datafusion-ip-alloc"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 22
  network       = google_compute_network.network.id
}
```
