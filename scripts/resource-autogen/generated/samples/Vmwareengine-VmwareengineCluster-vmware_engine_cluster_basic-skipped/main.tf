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
resource "google_vmwareengine_cluster" "vmw-engine-ext-cluster" {
  provider = google-beta
  name     = "ext-cluster"
  parent   = google_vmwareengine_private_cloud.cluster-pc.id
  node_type_configs {
    node_type_id = "standard-72"
    node_count   = 3
  }
}

resource "google_vmwareengine_private_cloud" "cluster-pc" {
  provider    = google-beta
  location    = "us-west1-a"
  name        = "sample-pc"
  description = "Sample test PC."
  network_config {
    management_cidr       = "192.168.30.0/24"
    vmware_engine_network = google_vmwareengine_network.cluster-nw.id
  }

  management_cluster {
    cluster_id = "sample-mgmt-cluster"
    node_type_configs {
      node_type_id = "standard-72"
      node_count   = 3
    }
  }
}

resource "google_vmwareengine_network" "cluster-nw" {
  provider    = google-beta
  name        = "us-west1-default"
  location    = "us-west1"
  type        = "LEGACY"
  description = "PC network description."
}
```
