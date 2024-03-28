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
  provider = google-beta
  name               = "basiccluster"
  location           = "us-central1-a"
  initial_node_count = 1
}

resource "google_gke_hub_membership" "membershiprbacrolebinding" {
  provider = google-beta
  membership_id = "tf-test-membership%{random_suffix}"
  endpoint {
    gke_cluster {
      resource_link = "//container.googleapis.com/${google_container_cluster.primary.id}"
    }
  }

  depends_on = [google_container_cluster.primary]
}

resource "google_gke_hub_membership_rbac_role_binding" "membershiprbacrolebinding" {
  provider = google-beta
  membership_rbac_role_binding_id = "tf-test-membership-rbac-role-binding%{random_suffix}"
  membership_id = "tf-test-membership%{random_suffix}"
  user = "service-${data.google_project.project.number}@gcp-sa-anthossupport.iam.gserviceaccount.com"
  role {
    predefined_role = "ANTHOS_SUPPORT"
  }
  location = "global"
  depends_on = [google_gke_hub_membership.membershiprbacrolebinding]
}

data "google_project" "project" {
  provider = google-beta
}
```
