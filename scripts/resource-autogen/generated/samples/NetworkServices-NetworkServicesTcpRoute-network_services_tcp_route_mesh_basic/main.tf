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
resource "google_compute_backend_service" "default" {
  provider               = google-beta
  name          = "my-backend-service"
  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  provider               = google-beta
  name               = "backend-service-health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_network_services_mesh" "default" {
  provider    = google-beta
  name        = "my-tcp-route"
  labels      = {
    foo = "bar"
  }
  description = "my description"
}


resource "google_network_services_tcp_route" "default" {
  provider               = google-beta
  name                   = "my-tcp-route"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  meshes = [
    google_network_services_mesh.default.id
  ]
  rules                   {
    matches {
      address = "10.0.0.1/32"
      port = "8081"
    }
    action {
      destinations {
        service_name = google_compute_backend_service.default.id
        weight = 1
      }
      original_destination = false
    }
  }
}
```
