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
data "google_client_config" "current" {}

resource "google_project_service" "apigee" {
  project = data.google_client_config.current.project
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = data.google_client_config.current.project
  service = "compute.googleapis.com"
}

resource "google_project_service" "servicenetworking" {
  project = data.google_client_config.current.project
  service = "servicenetworking.googleapis.com"
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = data.google_client_config.current.project
  depends_on = [ google_project_service.compute ]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = data.google_client_config.current.project
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
}

resource "google_apigee_organization" "org" {
  analytics_region   = "us-central1"
  project_id         = data.google_client_config.current.project
  authorized_network = google_compute_network.apigee_network.id
  billing_type       = "EVALUATION"
  depends_on         = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee
  ]
}

resource "google_apigee_addons_config" "test_organization" {
  org = google_apigee_organization.org.name

  addons_config {
    integration_config {
      enabled = true
    }
    api_security_config {
      enabled = true
    }
    connectors_platform_config {
      enabled = true
    }
    monetization_config {
      enabled = true
    }
    advanced_api_ops_config {
      enabled = true
    }
  }
}
```
