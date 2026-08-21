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
resource "google_project" "my_project" {
  provider   = google-beta
  name       = "tf-test-project"
  project_id = "quota"
  org_id     = "123456789"
}

resource "google_service_usage_consumer_quota_override" "override" {
  provider       = google-beta
  project        = google_project.my_project.project_id
  service        = "servicemanagement.googleapis.com"
  metric         = urlencode("servicemanagement.googleapis.com/default_requests")
  limit          = urlencode("/min/project")
  override_value = "0"
  force          = true
}
```
