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
  provider                = google-beta
  name                    = "workstation-cluster"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "workstation-cluster"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}

resource "google_compute_disk" "my_source_disk" {
  provider = google-beta
  name     = "workstation-config"
  size     = 10
  type     = "pd-ssd"
  zone     = "us-central1-a"
}

resource "google_compute_snapshot" "my_source_snapshot" {
  provider    = google-beta
  name        = "workstation-config"
  source_disk = google_compute_disk.my_source_disk.name
  zone        = "us-central1-a"
}

resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "workstation-cluster"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"
}

resource "google_workstations_workstation_config" "default" {
  provider               = google-beta
  workstation_config_id  = "workstation-config"
  workstation_cluster_id = google_workstations_workstation_cluster.default.workstation_cluster_id
  location               = google_workstations_workstation_cluster.default.location

  persistent_directories {
    mount_path = "/home"

    gce_pd {
      source_snapshot = google_compute_snapshot.my_source_snapshot.id
      reclaim_policy  = "DELETE"
    }
  }
}
```
