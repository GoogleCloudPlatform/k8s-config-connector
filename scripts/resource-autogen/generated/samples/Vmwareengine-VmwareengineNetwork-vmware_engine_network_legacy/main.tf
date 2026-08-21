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
resource "google_vmwareengine_network" "vmw-engine-network" {
  provider    = google-beta
  project     = google_project_service.acceptance.project
  name        = "us-west1-default" #Legacy network IDs are in the format: {region-id}-default
  location    = "us-west1"
  type        = "LEGACY"
  description = "VMwareEngine legacy network sample"
}

resource "google_project_service" "acceptance" {
  project  = google_project.acceptance.project_id
  provider = google-beta
  service  = "vmwareengine.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

# there can be only 1 Legacy network per region for a given project,
# so creating new project for isolation in CI.
resource "google_project" "acceptance" {
  name            = "vmw-proj"
  provider        = google-beta
  project_id      = "vmw-proj"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.acceptance]

  create_duration = "60s"
}
```
