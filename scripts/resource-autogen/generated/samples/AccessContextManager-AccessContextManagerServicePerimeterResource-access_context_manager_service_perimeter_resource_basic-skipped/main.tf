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
resource "google_access_context_manager_service_perimeter_resource" "service-perimeter-resource" {
  perimeter_name = google_access_context_manager_service_perimeter.service-perimeter-resource.name
  resource = "projects/987654321"
}

resource "google_access_context_manager_service_perimeter" "service-perimeter-resource" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/restrict_all"
  title  = "restrict_all"
  status {
    restricted_services = ["storage.googleapis.com"]
  }

  lifecycle {
    ignore_changes = [status[0].resources]
  }
}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "my policy"
}
```
