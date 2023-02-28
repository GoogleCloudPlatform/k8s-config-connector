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
resource "google_compute_instance_group_named_port" "my_port" {
  group = google_container_cluster.my_cluster.node_pool[0].instance_group_urls[0]
  zone = "us-central1-a"

  name = "http"
  port = 8080
}

resource "google_compute_instance_group_named_port" "my_ports" {
  group = google_container_cluster.my_cluster.node_pool[0].instance_group_urls[0]
  zone = "us-central1-a"

  name = "https"
  port = 4443
}

resource "google_compute_network" "container_network" {
  name                    = "container-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "container_subnetwork" {
  name                     = "container-subnetwork"
  region                   = "us-central1"
  network                  = google_compute_network.container_network.name
  ip_cidr_range            = "10.0.36.0/24"
}

resource "google_container_cluster" "my_cluster" {
  name               = "my-cluster"
  location           = "us-central1-a"
  initial_node_count = 1

  network    = google_compute_network.container_network.name
  subnetwork = google_compute_subnetwork.container_subnetwork.name

  ip_allocation_policy {
    cluster_ipv4_cidr_block  = "/19"
    services_ipv4_cidr_block = "/22"
  }
}
```
