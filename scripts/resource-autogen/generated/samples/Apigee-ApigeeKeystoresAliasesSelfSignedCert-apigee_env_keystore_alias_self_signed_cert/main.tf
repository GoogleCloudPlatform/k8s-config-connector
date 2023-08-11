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
  project_id      = "my-project"
  name            = "my-project"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
  depends_on = [google_project_service.apigee]
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
  depends_on = [google_project_service.servicenetworking]
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee,
  ]
}

resource "google_apigee_environment" "apigee_environment_keystore_ss_alias" {
  org_id       = google_apigee_organization.apigee_org.id
  name         = "env-name"
  description  = "Apigee Environment"
  display_name = "environment-1"
}

resource "google_apigee_env_keystore" "apigee_environment_keystore_alias" {
  name       = "env-keystore"
  env_id     = google_apigee_environment.apigee_environment_keystore_ss_alias.id
}

resource "google_apigee_keystores_aliases_self_signed_cert" "apigee_environment_keystore_ss_alias" {
  environment			      = google_apigee_environment.apigee_environment_keystore_ss_alias.name
  org_id				        = google_apigee_organization.apigee_org.name
  keystore				      = google_apigee_env_keystore.apigee_environment_keystore_alias.name
  alias                 = "alias"
  key_size              = 1024
  sig_alg               = "SHA512withRSA"
  cert_validity_in_days = 4
  subject {
      common_name = "selfsigned_example"
      country_code = "US"
      locality = "TX"
      org = "CCE"
      org_unit = "PSO"
  }   
}
```
