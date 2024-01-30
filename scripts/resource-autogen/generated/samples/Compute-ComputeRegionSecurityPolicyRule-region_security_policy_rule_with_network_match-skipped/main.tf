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
# First activate advanced network DDoS protection for the desired region
resource "google_compute_region_security_policy" "policyddosprotection" {
  provider    = google-beta

  region      = "us-west2"
  name        = "policyddosprotection"
  description = "policy for activating network DDoS protection for the desired region"
  type        = "CLOUD_ARMOR_NETWORK"
  ddos_protection_config {
    ddos_protection = "ADVANCED_PREVIEW"
  }
}

resource "google_compute_network_edge_security_service" "edge_sec_service" {
  provider        = google-beta

  region          = "us-west2"
  name            = "edgesecservice"
  description     = "linking policy to edge security service"
  security_policy = google_compute_region_security_policy.policyddosprotection.self_link
}

# Add the desired policy and custom rule.
resource "google_compute_region_security_policy" "policynetworkmatch" {
  provider    = google-beta

  region      = "us-west2"
  name        = "policyfornetworkmatch"
  description = "region security policy for network match"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "TCP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  depends_on  = [google_compute_network_edge_security_service.edge_sec_service]
}

resource "google_compute_region_security_policy_rule" "policy_rule_network_match" {
  provider        = google-beta

  region          = "us-west2"
  security_policy = google_compute_region_security_policy.policynetworkmatch.name
  description     = "custom rule for network match"
  priority        = 100
  network_match {
    src_ip_ranges = ["10.10.0.0/16"]
    user_defined_fields {
      name = "SIG1_AT_0"
      values = ["0x8F00"]
    }
  }
  action          = "allow"
  preview         = true
}
```
