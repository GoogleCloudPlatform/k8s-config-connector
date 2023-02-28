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
resource "google_bigtable_instance" "instance" {
  name = "bt-instance"
  cluster {
    cluster_id   = "cluster-1"
    zone         = "us-central1-a"
    num_nodes    = 3
    storage_type = "HDD"
  }
  cluster {
    cluster_id   = "cluster-2"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }
  cluster {
    cluster_id   = "cluster-3"
    zone         = "us-central1-c"
    num_nodes    = 3
    storage_type = "HDD"
  }  

  deletion_protection  = "true"
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "bt-profile"

  // Requests will be routed to any of the 3 clusters.
  multi_cluster_routing_use_any = true

  ignore_warnings               = true
}
```
