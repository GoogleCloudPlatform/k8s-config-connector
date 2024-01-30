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
resource "google_compute_network_attachment" "default" {
    provider = google-beta
    name = "basic-network-attachment"
    region = "us-central1"
    description = "basic network attachment description"
    connection_preference = "ACCEPT_MANUAL"

    subnetworks = [
        google_compute_subnetwork.default.self_link
    ]

    producer_accept_lists = [
        google_project.accepted_producer_project.project_id
    ]

    producer_reject_lists = [
        google_project.rejected_producer_project.project_id
    ]
}

resource "google_compute_network" "default" {
    provider = google-beta
    name = "basic-network"
    auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
    provider = google-beta
    name = "basic-subnetwork"
    region = "us-central1"

    network = google_compute_network.default.id
    ip_cidr_range = "10.0.0.0/16"
}

resource "google_project" "rejected_producer_project" {
    provider = google-beta
    project_id      = "prj-rejected%{random_suffix}"
    name            = "prj-rejected%{random_suffix}"
    org_id          = "123456789"
    billing_account = "000000-0000000-0000000-000000"
}

resource "google_project" "accepted_producer_project" {
    provider = google-beta
    project_id      = "prj-accepted%{random_suffix}"
    name            = "prj-accepted%{random_suffix}"
    org_id          = "123456789"
    billing_account = "000000-0000000-0000000-000000"
}
```
