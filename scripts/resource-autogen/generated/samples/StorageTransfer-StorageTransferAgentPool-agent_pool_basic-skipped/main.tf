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
data "google_storage_transfer_project_service_account" "default" {
  project = "my-project-name"
}

resource "google_project_iam_member" "pubsub_editor_role" {
  project = "my-project-name"
  role    = "roles/pubsub.editor"
  member  = "serviceAccount:${data.google_storage_transfer_project_service_account.default.email}"
}

resource "google_storage_transfer_agent_pool" "example" {
  name         = "agent-pool-example"
  display_name = "Source A to destination Z"
  bandwidth_limit {
    limit_mbps = "120"
  }

  depends_on = [google_project_iam_member.pubsub_editor_role]
}
```
