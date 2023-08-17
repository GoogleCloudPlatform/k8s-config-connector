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
resource "google_gkeonprem_bare_metal_admin_cluster" "admin-cluster-basic" {
  provider = google-beta
  name = "my-cluster"
  location = "us-west1"
  bare_metal_version = "1.13.4"
  network_config {
    island_mode_cidr {
      service_address_cidr_blocks = ["172.26.0.0/16"]
      pod_address_cidr_blocks = ["10.240.0.0/13"]
    }
  }
  node_config {
      max_pods_per_node = 250
  }
  control_plane {
    control_plane_node_pool_config {
      node_pool_config {
        labels = {}
        operating_system = "LINUX"
        node_configs {
            labels  = {}
            node_ip = "10.200.0.2"
        }
        node_configs {
            labels  = {}
            node_ip = "10.200.0.3"
        }
        node_configs {
            labels  = {}
            node_ip = "10.200.0.4"
        }
      }
    }
  }
  load_balancer {
    port_config {
      control_plane_load_balancer_port = 443
    }
    vip_config {
      control_plane_vip = "10.200.0.5"
    }
  }
  storage {
    lvp_share_config {
      lvp_config {
        path = "/mnt/localpv-share"
        storage_class = "local-shared"
      }
      shared_path_pv_count = 5
    }
    lvp_node_mounts_config {
      path = "/mnt/localpv-disk"
      storage_class = "local-disks"
    }
  }
  node_access_config {
    login_user = "root"
  }
}
```
