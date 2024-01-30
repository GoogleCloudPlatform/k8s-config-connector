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
resource "google_compute_region_security_policy" "default" {
  provider    = google-beta
  
  region      = "us-west2"
  name        = "policywithmultiplerules"
  description = "basic region security policy"
  type        = "CLOUD_ARMOR"
}

resource "google_compute_region_security_policy_rule" "policy_rule_one" {
  provider = google-beta
  
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.default.name
  description     = "new rule one"
  priority        = 100
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["10.10.0.0/16"]
    }
  }
  action          = "allow"
  preview         = true
}

resource "google_compute_region_security_policy_rule" "policy_rule_two" {
  provider = google-beta
  
  region          = "us-west2"
  security_policy = google_compute_region_security_policy.default.name
  description     = "new rule two"
  priority        = 101
  match {
    versioned_expr = "SRC_IPS_V1"
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
    }
  }
  action          = "allow"
  preview         = true
}
```
