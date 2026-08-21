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
resource "google_project" "owner_project" {
  project_id      = "tf-test%{random_suffix}"
  name            = "tf-test%{random_suffix}"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}


resource "google_project_service" "compute" {
  project = google_project.owner_project.project_id
  service = "compute.googleapis.com"
  disable_on_destroy = false
}

resource "google_project" "guest_project" {
  project_id      = "tf-test-2%{random_suffix}"
  name            = "tf-test-2%{random_suffix}"
  org_id          = "123456789"
}

resource "google_organization_policy" "shared_reservation_org_policy" {
  org_id     = "123456789"
  constraint = "constraints/compute.sharedReservationsOwnerProjects"
  list_policy {
    allow {
      values = ["projects/${google_project.owner_project.number}"]
    }
  }
}

resource "google_compute_reservation" "gce_reservation" {
  project = google_project.owner_project.project_id
  name = "gce-shared-reservation"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      min_cpu_platform = "Intel Cascade Lake"
      machine_type     = "n2-standard-2"
    }
  }
  share_settings {
    share_type = "SPECIFIC_PROJECTS"
    project_map {
      id = google_project.guest_project.project_id
      project_id = google_project.guest_project.project_id
    }
  }
  depends_on = [google_organization_policy.shared_reservation_org_policy,google_project_service.compute]
}
```
