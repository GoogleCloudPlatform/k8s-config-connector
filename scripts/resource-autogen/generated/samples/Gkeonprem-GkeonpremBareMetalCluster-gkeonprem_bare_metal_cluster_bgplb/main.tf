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
resource "google_gkeonprem_bare_metal_cluster" "cluster-bgplb" {
  provider = google-beta
  name = "cluster-bgplb"
  location = "us-west1"
  admin_cluster_membership = "projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"
  bare_metal_version = "1.12.3"
  network_config {
    island_mode_cidr {
      service_address_cidr_blocks = ["172.26.0.0/16"]
      pod_address_cidr_blocks = ["10.240.0.0/13"]
    }
    advanced_networking = true
    multiple_network_interfaces_config {
      enabled = true
    }
    sr_iov_config {
      enabled = true
    }
  }
  control_plane {
    control_plane_node_pool_config {
      node_pool_config {
        labels = {}
        operating_system = "LINUX"
        node_configs {
          labels = {}
          node_ip = "10.200.0.9"
        }
        taints {
          key = "test-key"
          value = "test-value"
          effect = "NO_EXECUTE"
        }
      }
    }
    api_server_args {
      argument = "test-argument"
      value = "test-value"
    }
  }
  load_balancer {
    port_config {
      control_plane_load_balancer_port = 443
    }
    vip_config {
      control_plane_vip = "10.200.0.13"
      ingress_vip = "10.200.0.14"
    }
    bgp_lb_config {
      asn = 123456
      bgp_peer_configs {
        asn = 123457
        ip_address = "10.0.0.1"
        control_plane_nodes = ["test-node"]
      }
      address_pools {
        pool = "pool1"
        addresses = [
          "10.200.0.14/32",
          "10.200.0.15/32",
          "10.200.0.16/32",
          "10.200.0.17/32",
          "10.200.0.18/32",
          "fd00:1::f/128",
          "fd00:1::10/128",
          "fd00:1::11/128",
          "fd00:1::12/128"
        ]
      }
      load_balancer_node_pool_config {
        node_pool_config {
          labels = {}
          operating_system = "LINUX"
          node_configs {
            labels = {}
            node_ip = "10.200.0.9"
          }
          taints {
            key = "test-key"
            value = "test-value"
            effect = "NO_EXECUTE"
          }
          kubelet_config {
            registry_pull_qps = 10
            registry_burst = 12
            serialize_image_pulls_disabled = true
          }
        }
      }
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
  security_config {
    authorization {
      admin_users {
        username = "admin@hashicorptest.com"
      }
    }
  }
  proxy {
    uri = "http://test-domain/test"
    no_proxy = ["127.0.0.1"]
  }
  cluster_operations {
    enable_application_logs = true
  }
  maintenance_config {
    maintenance_address_cidr_blocks = ["192.168.0.1/20"] 
  }
  node_config {
    max_pods_per_node = 10
    container_runtime = "CONTAINERD"
  }
  node_access_config {
    login_user = "test@example.com"
  }
  os_environment_config {
    package_repo_excluded = true
  }
}
```
