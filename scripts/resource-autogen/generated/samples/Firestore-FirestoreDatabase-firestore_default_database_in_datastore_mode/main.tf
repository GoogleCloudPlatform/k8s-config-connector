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
resource "google_project" "project" {
	project_id = "tf-test%{random_suffix}"
	name       = "tf-test%{random_suffix}"
	org_id     = "123456789"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]
  create_duration = "60s"
}

resource "google_project_service" "firestore" {
  project = google_project.project.project_id
  service = "firestore.googleapis.com"
  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_firestore_database" "datastore_mode_database" {
    project = google_project.project.project_id

    name = "(default)"

    location_id = "nam5"
    type        = "DATASTORE_MODE"

    depends_on = [google_project_service.firestore]
}
```
