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
resource "google_service_directory_namespace" "default" {
  provider     = google-beta
  namespace_id = "my-namespace"
  location     = "us-central1"
}

resource "google_service_directory_service" "default" {
  provider   = google-beta
  service_id = "my-service"
  namespace  = google_service_directory_namespace.default.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }
}

resource "google_network_services_service_binding" "default" {
  provider    = google-beta
  name        = "my-service-binding"
  labels      = {
    foo = "bar"
  }
  description = "my description"
  service = google_service_directory_service.default.id
}
```
