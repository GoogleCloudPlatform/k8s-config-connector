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
resource "google_compute_region_url_map" "regionurlmap" {
  region = "us-central1"

  name        = "regionurlmap"
  description = "a description"

  default_service = google_compute_region_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      service = google_compute_region_backend_service.home.id
    }

    path_rule {
      paths   = ["/login"]
      service = google_compute_region_backend_service.login.id
    }
  }

  test {
    service = google_compute_region_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_region_backend_service" "login" {
  region = "us-central1"

  name        = "login"
  protocol    = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_backend_service" "home" {
  region = "us-central1"

  name        = "home"
  protocol    = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_health_check" "default" {
  region = "us-central1"

  name               = "health-check"
  check_interval_sec = 1
  timeout_sec        = 1
  http_health_check {
    port         = 80
    request_path = "/"
  }
}
```
