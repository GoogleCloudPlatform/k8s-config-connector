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
resource "google_network_services_edge_cache_origin" "fallback" {
  name                 = "my-fallback"
  origin_address       = "fallback.example.com"
  description          = "The default bucket for media edge test"
  max_attempts         = 3
  protocol = "HTTP"
  port = 80

  retry_conditions = [
    "CONNECT_FAILURE",
    "NOT_FOUND",
    "HTTP_5XX",
    "FORBIDDEN",
  ]
  timeout {
    connect_timeout = "10s"
    max_attempts_timeout = "20s"
    response_timeout = "60s"
    read_timeout = "5s"
  }
  origin_override_action {
    url_rewrite {
      host_rewrite = "example.com"
    }
    header_action {
      request_headers_to_add {
        header_name = "x-header"
	header_value = "value"
	replace = true
      }
    }
  }
  origin_redirect {
    redirect_conditions = [
      "MOVED_PERMANENTLY",
      "FOUND",
      "SEE_OTHER",
      "TEMPORARY_REDIRECT",
      "PERMANENT_REDIRECT",
    ]
  }
}

resource "google_network_services_edge_cache_origin" "default" {
  name                 = "my-origin"
  origin_address       = "gs://media-edge-default"
  failover_origin      = google_network_services_edge_cache_origin.fallback.id
  description          = "The default bucket for media edge test"
  max_attempts         = 2
  labels = {
    a = "b"
  }

  timeout {
    connect_timeout = "10s"
  }
}
```
