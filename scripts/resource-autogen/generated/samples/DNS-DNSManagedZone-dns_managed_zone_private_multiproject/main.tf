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
resource "google_dns_managed_zone" "private-zone" {
  name        = "private-zone"
  dns_name    = "multiproject.private.example.com."
  description = "Example private DNS zone"
  labels = {
    foo = "bar"
  }

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network_1_project_1.id
    }
    networks {
      network_url = google_compute_network.network_2_project_1.id
    }
    networks {
      network_url = google_compute_network.network_1_project_2.id
    }
    networks {
      network_url = google_compute_network.network_2_project_2.id
    }
  }

  depends_on = [
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_project" "project_1" {
  name            = "project-1"
  project_id      = "project-1"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}

resource "google_project" "project_2" {
  name            = "project-2"
  project_id      = "project-2"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}

resource "google_compute_network" "network_1_project_1" {
  name                    = "network-1"
  project                 = google_project.project_1.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
  ]
}

resource "google_compute_network" "network_2_project_1" {
  name                    = "network-2"
  project                 = google_project.project_1.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
  ]
}

resource "google_compute_network" "network_1_project_2" {
  name                    = "network-1"
  project                 = google_project.project_2.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_compute_network" "network_2_project_2" {
  name                    = "network-2"
  project                 = google_project.project_2.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_project_service" "compute_project_1" {
  project    = google_project.project_1.project_id
  service    = "compute.googleapis.com"
  depends_on = [
    google_project.project_1,
  ]
}

resource "google_project_service" "compute_project_2" {
  project    = google_project.project_2.project_id
  service    = "compute.googleapis.com"
  depends_on = [
    google_project_service.dns_project_1
  ]
}

resource "google_project_service" "dns_project_1" {
  project    = google_project.project_1.project_id
  service    = "dns.googleapis.com"
  depends_on = [
    google_project_service.compute_project_1
  ]
}

resource "google_project_service" "dns_project_2" {
  project    = google_project.project_2.project_id
  service    = "dns.googleapis.com"
  depends_on = [
    google_project_service.compute_project_2,
  ]
}
```
