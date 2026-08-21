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
resource "google_alloydb_cluster" "full" {
  cluster_id   = "alloydb-cluster-full"
  location     = "us-central1"
  network      = "projects/${data.google_project.project.number}/global/networks/${google_compute_network.default.name}"

  initial_user {
    user     = "alloydb-cluster-full"
    password = "alloydb-cluster-full"
  }

  automated_backup_policy {
    location      = "us-central1"
    backup_window = "1800s"
    enabled       = true

    weekly_schedule {
      days_of_week = ["MONDAY"]

      start_times {
        hours   = 23
        minutes = 0
        seconds = 0
        nanos   = 0
      }
    }

    quantity_based_retention {
      count = 1
    }

    labels = {
      test = "alloydb-cluster-full"
    }
  }

  labels = {
    test = "alloydb-cluster-full"
  }
}

data "google_project" "project" {}

resource "google_compute_network" "default" {
  name = "alloydb-cluster-full"
}
```
