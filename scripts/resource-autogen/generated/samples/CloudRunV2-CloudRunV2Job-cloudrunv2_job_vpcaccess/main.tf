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
resource "google_cloud_run_v2_job" "default" {
  name     = "cloudrun-job"
  location = "us-central1"

  template {
    template{
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
      vpc_access{
        connector = google_vpc_access_connector.connector.id
        egress = "ALL_TRAFFIC"
      }
    }
  }

  lifecycle {
    ignore_changes = [
      launch_stage,
    ]
  }
}

resource "google_vpc_access_connector" "connector" {
  name          = "run-vpc"
  subnet {
    name = google_compute_subnetwork.custom_test.name
  }
  machine_type = "e2-standard-4"
  min_instances = 2
  max_instances = 3
  region        = "us-central1"
}
resource "google_compute_subnetwork" "custom_test" {
  name          = "run-subnetwork"
  ip_cidr_range = "10.2.0.0/28"
  region        = "us-central1"
  network       = google_compute_network.custom_test.id
}
resource "google_compute_network" "custom_test" {
  name                    = "run-network"
  auto_create_subnetworks = false
}
```
