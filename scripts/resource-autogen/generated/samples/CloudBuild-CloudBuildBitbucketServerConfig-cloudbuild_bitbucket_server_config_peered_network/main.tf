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
data "google_project" "project" {}

resource "google_project_service" "servicenetworking" {
  service = "servicenetworking.googleapis.com"
  disable_on_destroy = false
}
 
data "google_compute_network" "vpc_network" {
    name       = "vpc-network"
    depends_on = [google_project_service.servicenetworking]
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "private-ip-alloc"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = data.google_compute_network.vpc_network.id
}

resource "google_service_networking_connection" "default" {
  network                 = data.google_compute_network.vpc_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
  depends_on              = [google_project_service.servicenetworking]
}

resource "google_cloudbuild_bitbucket_server_config" "bbs-config-with-peered-network" {
    config_id = "bbs-config"
    location = "us-central1"
    host_uri = "https://bbs.com"
    secrets {
        admin_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        read_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        webhook_secret_version_name = "projects/myProject/secrets/mybbspat/versions/1"
    }
    username = "test"
    api_key = "<api-key>"
    peered_network = replace(data.google_compute_network.vpc_network.id, data.google_project.project.name, data.google_project.project.number)
    ssl_ca = "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n"
    depends_on = [google_service_networking_connection.default]
}
```
