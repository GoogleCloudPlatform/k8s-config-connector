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
resource "google_dns_managed_zone" "sd-zone" {
  provider = google-beta

  name        = "peering-zone"
  dns_name    = "services.example.com."
  description = "Example private DNS Service Directory zone"

  visibility = "private"

  service_directory_config {
    namespace {
      namespace_url = google_service_directory_namespace.example.id
    }
  }
}

resource "google_service_directory_namespace" "example" {
  provider = google-beta

  namespace_id = "example"
  location     = "us-central1"
}

resource "google_compute_network" "network" {
  provider = google-beta

  name                    = "network"
  auto_create_subnetworks = false
}
```
