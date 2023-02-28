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
resource "google_compute_global_network_endpoint_group" "external_proxy" {
  provider = google-beta
  name                  = "network-endpoint"
  network_endpoint_type = "INTERNET_FQDN_PORT"
  default_port          = "443"
}

resource "google_compute_global_network_endpoint" "proxy" {
  provider = google-beta
  global_network_endpoint_group = google_compute_global_network_endpoint_group.external_proxy.id
  fqdn                          = "test.example.com"
  port                          = google_compute_global_network_endpoint_group.external_proxy.default_port
}

resource "google_compute_backend_service" "default" {
  provider = google-beta
  name                            = "backend-service"
  enable_cdn                      = true
  timeout_sec                     = 10
  connection_draining_timeout_sec = 10

  custom_request_headers          = ["host: ${google_compute_global_network_endpoint.proxy.fqdn}"]
  custom_response_headers         = ["X-Cache-Hit: {cdn_cache_status}"]

  backend {
    group = google_compute_global_network_endpoint_group.external_proxy.id
  }
}
```
