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
resource "google_edgecontainer_cluster" "default" {
  name = "local-control-plane-cluster"
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

  external_load_balancer_ipv4_address_pools = ["10.100.0.0-10.100.0.10"]

  control_plane {
    local {
      node_location = "us-central1-edge-example-edgesite"
      node_count = 1
      machine_filter = "machine-name"
      shared_deployment_policy = "ALLOWED"
    }
  }
}

data "google_project" "project" {}
```
