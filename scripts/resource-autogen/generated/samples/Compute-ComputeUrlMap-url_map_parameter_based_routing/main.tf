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
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap"
  description = "parameter-based routing example"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts = ["*"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_service.default.id

    route_rules {
      priority = 1
      service = google_compute_backend_service.service-a.id
      match_rules {
        prefix_match = "/"
        ignore_case = true
        query_parameter_matches {
          name = "abtest"
          exact_match = "a"
        }
      }
    }
    route_rules {
      priority = 2
      service = google_compute_backend_service.service-b.id
      match_rules {
        ignore_case = true
        prefix_match = "/"
        query_parameter_matches {
          name = "abtest"
          exact_match = "b"
        }
      }
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "default"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "service-a" {
  name        = "service-a"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "service-b" {
  name        = "service-b"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
```
