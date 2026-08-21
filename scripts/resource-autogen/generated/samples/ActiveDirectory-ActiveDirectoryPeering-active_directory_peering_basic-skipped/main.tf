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
resource "google_active_directory_peering" "ad-domain-peering" {
    provider           = google-beta
    domain_resource    = google_active_directory_domain.ad-domain.name
    peering_id         = "ad-domain-peering"
    authorized_network = google_compute_network.peered-network.id
    labels             = {
        foo = "bar"
    }
}

resource "google_active_directory_domain" "ad-domain" {
    provider            = google-beta
    domain_name         = "ad.test.hashicorptest.com"
    locations           = ["us-central1"]
    reserved_ip_range   = "192.168.255.0/24"
    authorized_networks = [google_compute_network.source-network.id]
}

resource "google_compute_network" "peered-network" {
    provider = google-beta
    project  = google_project_service.compute.project
    name     = "ad-peered-network"
}

resource "google_compute_network" "source-network" {
    provider = google-beta
    name     = "ad-network"
}

resource "google_project_service" "compute" {
    provider = google-beta
    project  = google_project.peered-project.project_id
    service  = "compute.googleapis.com"
}

resource "google_project" "peered-project" {
    provider        = google-beta
    name            = "my-peered-project"
    project_id      = "my-peered-project"
    org_id          = "123456789"
    billing_account = "000000-0000000-0000000-000000"
}
```
