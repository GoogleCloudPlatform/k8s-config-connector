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
  description = "a description"

  default_service = google_compute_backend_bucket.static.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "mysite"
  }

  path_matcher {
    name            = "mysite"
    default_service = google_compute_backend_bucket.static.id

    route_rules {
      match_rules {
        path_template_match = "/xyzwebservices/v2/xyz/users/{username=*}/carts/{cartid=**}"
      }
      service = google_compute_backend_service.cart-backend.id
      priority = 1
      route_action {
        url_rewrite {
          path_template_rewrite = "/{username}-{cartid}/"
        }
      }
    }

    route_rules {
      match_rules {
        path_template_match = "/xyzwebservices/v2/xyz/users/*/accountinfo/*"
      }
      service = google_compute_backend_service.user-backend.id
      priority = 2
    }
  }
}

resource "google_compute_backend_service" "cart-backend" {
  name        = "cart-service"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10
  load_balancing_scheme = "EXTERNAL_MANAGED"

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "user-backend" {
  name        = "user-service"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10
  load_balancing_scheme = "EXTERNAL_MANAGED"

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "health-check"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_backend_bucket" "static" {
  name        = "static-asset-backend-bucket"
  bucket_name = google_storage_bucket.static.name
  enable_cdn  = true
}

resource "google_storage_bucket" "static" {
  name     = "static-asset-bucket"
  location = "US"
}
```
