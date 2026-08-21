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

resource "google_network_services_edge_cache_origin" "google" {
  name                 = "origin-google"
  origin_address       = "google.com"
  description          = "The default bucket for media edge test"
  max_attempts         = 2
  timeout {
    connect_timeout = "10s"
  }
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
  disable_quic         = true
  disable_http2        = true
  labels = {
    a = "b"
  }

  routing {
    host_rule {
      description = "host rule description"
      hosts = ["sslcert.tf-test.club"]
      path_matcher = "routes"
    }
    host_rule {
      description = "host rule2"
      hosts = ["sslcert.tf-test2.club"]
      path_matcher = "routes"
    }

    host_rule {
      description = "host rule3"
      hosts = ["sslcert.tf-test3.club"]
      path_matcher = "routesAdvanced"
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

    path_matcher {
      name = "routesAdvanced"
      description = "an advanced ruleset"
      route_rule {
        description = "an advanced route rule to match against"
        priority = 1
        match_rule {
          prefix_match = "/potato/"
          query_parameter_match {
            name = "debug"
            present_match = true
          }
          query_parameter_match {
            name = "state"
            exact_match = "debug"
          }
        }
        match_rule {
          full_path_match = "/apple"
        }
        header_action {
          request_header_to_add {
            header_name = "debug"
            header_value = "true"
            replace = true
          }
          request_header_to_add {
            header_name = "potato"
            header_value = "plant"
          }
          response_header_to_add {
            header_name = "potato"
            header_value = "plant"
            replace = true
          }
          request_header_to_remove {
            header_name = "prod"
          }
          response_header_to_remove {
            header_name = "prod"
          }
        }

        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
              cache_mode = "CACHE_ALL_STATIC"
              default_ttl = "3800s"
              client_ttl = "3600s"
              max_ttl = "9000s"
              cache_key_policy {
                include_protocol = true
                exclude_host = true
                included_query_parameters = ["apple", "dev", "santa", "claus"]
                included_header_names = ["banana"]
                included_cookie_names = ["orange"]
              }
              negative_caching = true
              signed_request_mode = "DISABLED"
              negative_caching_policy = {
                "500" = "3000s"
              }
          }
          url_rewrite {
            path_prefix_rewrite = "/dev"
            host_rewrite = "dev.club"
          }
          cors_policy {
            max_age = "2500s"
            allow_credentials = true
            allow_origins = ["*"]
            allow_methods = ["GET"]
            allow_headers = ["dev"]
            expose_headers = ["prod"]
          }
        }
      }
      route_rule {
        description = "a second route rule to match against"
        priority = 2
        match_rule {
          full_path_match = "/yay"
        }
        origin = google_network_services_edge_cache_origin.instance.name
        route_action {
          cdn_policy {
            cache_mode = "CACHE_ALL_STATIC"
            default_ttl = "3600s"
            cache_key_policy {
              excluded_query_parameters = ["dev"]
            }
          }
          cors_policy {
            max_age = "3000s"
            allow_headers = ["dev"]
            disabled = true
          }
        }
      }
    }
  }

  log_config {
    enable = true
    sample_rate = 0.01
  }
}
```
