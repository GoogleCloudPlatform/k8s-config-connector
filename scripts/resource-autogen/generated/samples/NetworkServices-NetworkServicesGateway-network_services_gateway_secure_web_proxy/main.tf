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
resource "google_certificate_manager_certificate" "default" {
  name        = "my-certificate"
  location    = "us-central1"
  self_managed {
    pem_certificate = file("test-fixtures/certificatemanager/cert.pem")
    pem_private_key = file("test-fixtures/certificatemanager/private-key.pem")
  }
}

resource "google_compute_network" "default" {
  name                    = "my-network"
  routing_mode            = "REGIONAL"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "my-subnetwork-name"
  purpose       = "PRIVATE"
  ip_cidr_range = "10.128.0.0/20"
  region        = "us-central1"
  network       = google_compute_network.default.id
  role          = "ACTIVE"
}

resource "google_compute_subnetwork" "proxyonlysubnet" {
  name          = "my-proxy-only-subnetwork"
  purpose       = "REGIONAL_MANAGED_PROXY"
  ip_cidr_range = "192.168.0.0/23"
  region        = "us-central1"
  network       = google_compute_network.default.id
  role          = "ACTIVE"
}

resource "google_network_security_gateway_security_policy" "default" {
  name        = "my-policy-name"
  location    = "us-central1"
}

resource "google_network_security_gateway_security_policy_rule" "default" {
  name                    = "my-policyrule-name"
  location                = "us-central1"
  gateway_security_policy = google_network_security_gateway_security_policy.default.name
  enabled                 = true  
  priority                = 1
  session_matcher         = "host() == 'example.com'"
  basic_profile           = "ALLOW"
}

resource "google_network_services_gateway" "default" {
  name                                 = "my-gateway1"
  location                             = "us-central1"
  addresses                            = ["10.128.0.99"]
  type                                 = "SECURE_WEB_GATEWAY"
  ports                                = [443]
  scope                                = "my-default-scope1"
  certificate_urls                     = [google_certificate_manager_certificate.default.id]
  gateway_security_policy              = google_network_security_gateway_security_policy.default.id
  network                              = google_compute_network.default.id
  subnetwork                           = google_compute_subnetwork.default.id
  delete_swg_autogen_router_on_destroy = true
  depends_on                           = [google_compute_subnetwork.proxyonlysubnet]
}
```
