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
resource "google_container_cluster" "primary" {
  name               = "transform-rule-cluster"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "my-project-name.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "transform-rule"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_restore_plan" "transform_rule" {
  name = "transform-rule-rp"
  description = "copy nginx env variables"
  labels = {
    "app" = "nginx"
  }
  location = "us-central1"
  backup_plan = google_gke_backup_backup_plan.basic.id
  cluster = google_container_cluster.primary.id
  restore_config {
    excluded_namespaces {
      namespaces = ["my-ns"]
    }
    namespaced_resource_restore_mode = "DELETE_AND_RESTORE"
    volume_data_restore_policy = "RESTORE_VOLUME_DATA_FROM_BACKUP"
    cluster_resource_restore_scope {
      excluded_group_kinds {
        resource_group = "apiextension.k8s.io"
        resource_kind = "CustomResourceDefinition"
      }
    }
    cluster_resource_conflict_policy = "USE_EXISTING_VERSION"
    transformation_rules {
      description = "Copy environment variables from the nginx container to the install init container."
      resource_filter {
        group_kinds {
          resource_kind = "Pod"
          resource_group = ""
        }
        json_path = ".metadata[?(@.name == 'nginx')]"
      }
      field_actions {
        op = "COPY"
        path = "/spec/initContainers/0/env"
        from_path = "/spec/containers/0/env"
      }
    }
  }
}
```
