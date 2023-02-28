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
resource "google_compute_network" "net" {
  name                    = "my-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  name                     = "my-subnetwork"
  region                   = "us-central1"
  network                  = google_compute_network.net.id
  ip_cidr_range            = "10.0.0.0/22"
  private_ip_google_access = true
}

resource "google_dataproc_metastore_service" "default" {
  service_id = "metastore-srv"
  location   = "us-central1"

  hive_metastore_config {
    version = "3.1.2"
  }

  network_config {
    consumers {
      subnetwork = google_compute_subnetwork.subnet.id
    }
  }
}
```
