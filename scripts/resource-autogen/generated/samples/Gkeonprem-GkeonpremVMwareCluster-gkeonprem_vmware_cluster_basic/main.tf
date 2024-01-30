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
resource "google_gkeonprem_vmware_cluster" "cluster-basic" {
  provider = google-beta
  name = "cluster-basic"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  description = "test cluster"
  on_prem_version = "1.13.1-gke.35"
  annotations = {}
  network_config {
    service_address_cidr_blocks = ["10.96.0.0/12"]
    pod_address_cidr_blocks = ["192.168.0.0/16"]
    dhcp_ip_config {
      enabled = true
    }
  }
  control_plane_node {
     cpus = 4
     memory = 8192
     replicas = 1
  }
  load_balancer {
    vip_config {
      control_plane_vip = "10.251.133.5"
      ingress_vip = "10.251.135.19"
    }
    metal_lb_config {
      address_pools {
        pool = "ingress-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
        avoid_buggy_ips = true
      }
      address_pools {
        pool = "lb-test-ip"
        manual_assign = "true"
        addresses = ["10.251.135.19"]
        avoid_buggy_ips = true
      }
    }
  }
}
```
