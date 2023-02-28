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

  default_route_action {
    retry_policy {
      retry_conditions = [
        "5xx",
        "gateway-error",
      ]
      num_retries = 3
      per_try_timeout {
        seconds = 0
        nanos = 500
      }
    }
    request_mirror_policy {
      backend_service = google_compute_region_backend_service.home.id
    }
    weighted_backend_services {
      backend_service = google_compute_region_backend_service.login.id
      weight = 200
      header_action {
        request_headers_to_add {
          header_name = "foo-request-1"
          header_value = "bar"
          replace = true
        }
        request_headers_to_remove = ["fizz"]
        response_headers_to_add {
          header_name = "foo-response-1"
          header_value = "bar"
          replace = true
        }
        response_headers_to_remove = ["buzz"]
      }
    }
    weighted_backend_services {
      backend_service = google_compute_region_backend_service.home.id
      weight = 100
      header_action {
        request_headers_to_add {
          header_name = "foo-request-1"
          header_value = "bar"
          replace = true
        }
        request_headers_to_add {
          header_name = "foo-request-2"
          header_value = "bar"
          replace = true
        }
        request_headers_to_remove = ["fizz"]
        response_headers_to_add {
          header_name = "foo-response-2"
          header_value = "bar"
          replace = true
        }
        response_headers_to_add {
          header_name = "foo-response-1"
          header_value = "bar"
          replace = true
        }
        response_headers_to_remove = ["buzz"]
      }
    }
    url_rewrite {
      host_rewrite = "dev.example.com"
      path_prefix_rewrite = "/v1/api/"
    }
  
    cors_policy {
      disabled = false
      allow_credentials = true
      allow_headers = [
        "foobar"
      ]
      allow_methods = [
        "GET",
        "POST",
      ]
      allow_origins = [
        "example.com"
      ]
      expose_headers = [
        "foobar"
      ]
      max_age = 60
    }
    fault_injection_policy {
      delay {
        fixed_delay {
          seconds = 0
          nanos = 500
        }
        percentage = 0.5
      }
      abort {
        http_status = 500
        percentage = 0.5
      }
    }
    timeout {
      seconds = 0
      nanos = 500
    }
  }

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
