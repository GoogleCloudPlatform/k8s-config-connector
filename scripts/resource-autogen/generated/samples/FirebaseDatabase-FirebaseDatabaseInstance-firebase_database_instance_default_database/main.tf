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
resource "google_project" "default" {
  provider = google-beta
  project_id = "rtdb-project"
  name       = "rtdb-project"
  org_id     = "123456789"
  labels     = {
    "firebase" = "enabled"
  }
}

resource "google_firebase_project" "default" {
  provider = google-beta
  project  = google_project.default.project_id
}

resource "google_project_service" "firebase_database" {
  provider = google-beta
  project  = google_firebase_project.default.project
  service  = "firebasedatabase.googleapis.com"
}

resource "google_firebase_database_instance" "default" {
  provider = google-beta
  project  = google_firebase_project.default.project
  region   = "us-central1"
  instance_id = "rtdb-project-default-rtdb"
  type     = "DEFAULT_DATABASE"
  depends_on = [google_project_service.firebase_database]
}
```
