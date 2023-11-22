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
  name = "default"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.0.0.0/16"]
    services_ipv4_cidr_blocks = ["10.1.0.0/16"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }
}

resource "google_edgecontainer_node_pool" "node_pool" {
  name = "nodepool-1"
  cluster = google_edgecontainer_cluster.cluster.name
  location = "us-central1"
  node_location = "us-central1-edge-example-edgesite"
  node_count = 3
}

resource "google_edgecontainer_vpn_connection" "default" {
  depends_on = [google_edgecontainer_node_pool.node_pool]
  name = "vpn-connection-1"
  location = "us-central1"
  cluster = "projects/${data.google_project.project.number}/locations/us-east1/clusters/${google_edgecontainer_cluster.cluster.name}"
  vpc = google_compute_network.vpc.name
  enable_high_availability = true
}

resource "google_compute_network" "vpc" {
  name = "example-vpc"
}

data "google_project" "project" {}
```
