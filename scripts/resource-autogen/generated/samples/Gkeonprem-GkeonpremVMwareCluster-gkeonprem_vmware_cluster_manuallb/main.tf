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
resource "google_gkeonprem_vmware_cluster" "cluster-manuallb" {
  provider = google-beta
  name = "cluster-manuallb"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  description = "test cluster"
  on_prem_version = "1.13.1-gke.35"
  annotations = {}
  network_config {
    service_address_cidr_blocks = ["10.96.0.0/12"]
    pod_address_cidr_blocks = ["192.168.0.0/16"]
    host_config {
      dns_servers = ["10.254.41.1"]
      ntp_servers = ["216.239.35.8"]
      dns_search_domains = ["test-domain"]
    }
  
    static_ip_config {
      ip_blocks {
        netmask = "255.255.252.0"
        gateway = "10.251.31.254"
        ips {
          ip = "10.251.30.153"
          hostname = "test-hostname1"
        }
        ips {
          ip = "10.251.31.206"
          hostname = "test-hostname2"
        }
        ips {
          ip = "10.251.31.193"
          hostname = "test-hostname3"
        }
        ips { 
          ip = "10.251.30.230"
          hostname = "test-hostname4"
        }
      }
    }
  }
  control_plane_node {
     cpus = 4
     memory = 8192
     replicas = 1
     auto_resize_config {
      enabled = true
     }
  }
  load_balancer {
    vip_config {
      control_plane_vip = "10.251.133.5"
      ingress_vip = "10.251.135.19"
    }
    manual_lb_config {
      ingress_http_node_port = 30005
      ingress_https_node_port = 30006
      control_plane_node_port = 30007
      konnectivity_server_node_port = 30008
    }
  }
  dataplane_v2 {
    dataplane_v2_enabled = true
    windows_dataplane_v2_enabled = true
    advanced_networking = true
  }
  vm_tracking_enabled = true
  enable_control_plane_v2 = true
  upgrade_policy {
    control_plane_only = true
  }
  authorization {
    admin_users {
      username = "testuser@gmail.com"
    }
  }
  anti_affinity_groups {
    aag_config_disabled = true
  }
  auto_repair_config {
    enabled = true
  }
}
```
