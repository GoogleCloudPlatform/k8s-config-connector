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
resource "google_access_context_manager_service_perimeter" "service-perimeter" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/restrict_bigquery_dryrun_storage"
  title  = "restrict_bigquery_dryrun_storage"

  # Service 'bigquery.googleapis.com' will be restricted.
  status {
    restricted_services = ["bigquery.googleapis.com"]
  }

  # Service 'storage.googleapis.com' will be in dry-run mode.
  spec {
    restricted_services = ["storage.googleapis.com"]
  }

  use_explicit_dry_run_spec = true

}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "my policy"
}
```
