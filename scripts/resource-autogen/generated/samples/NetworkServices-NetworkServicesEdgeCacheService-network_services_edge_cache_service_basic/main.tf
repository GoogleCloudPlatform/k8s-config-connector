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
resource "google_storage_bucket" "dest" {
  name          = "my-bucket"
  location      = "US"
  force_destroy = true
}

resource "google_network_services_edge_cache_origin" "instance" {
  name                 = "my-origin"
  origin_address       = google_storage_bucket.dest.url
  description          = "The default bucket for media edge test"
  max_attempts         = 2
  timeout {
    connect_timeout = "10s"
  }
}

resource "google_network_services_edge_cache_service" "instance" {
  name                 = "my-service"
  description          = "some description"
  routing {
    host_rule {
      description = "host rule description"
      hosts = ["sslcert.tf-test.club"]
      path_matcher = "routes"
    }
    path_matcher {
      name = "routes"
      route_rule {
        description = "a route rule to match against"
        priority = 1
        match_rule {
          prefix_match = "/"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
              cache_mode = "CACHE_ALL_STATIC"
              default_ttl = "3600s"
          }
        }
        header_action {
          response_header_to_add {
            header_name = "x-cache-status"
            header_value = "{cdn_cache_status}"
          }
        }
      }
    }
  }
}
```
