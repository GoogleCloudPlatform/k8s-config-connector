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
resource "google_edgecontainer_cluster" "cluster" {
  name = "tf-lcp-cluster"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.96.0.0/17"]
    services_ipv4_cidr_blocks = ["10.0.0.0/21"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }

  external_load_balancer_ipv4_address_pools = ["10.100.68.100-10.100.68.102"]

  control_plane {
    local {
      node_location = "us-central1-edge-den29"
      node_count = 1
      machine_filter = "den29-06"
      shared_deployment_policy = "ALLOWED"
    }
  }
}

resource "google_edgecontainer_node_pool" "default" {
  name = "nodepool-1"
  cluster = google_edgecontainer_cluster.cluster.name
  location = "us-central1"
  node_location = "us-central1-edge-den29"
  machine_filter = "NOT name:den29-01"
  node_count = 1
}

data "google_project" "project" {}
```
