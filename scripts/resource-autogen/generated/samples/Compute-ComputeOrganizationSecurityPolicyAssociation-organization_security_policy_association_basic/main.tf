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
resource "google_folder" "security_policy_target" {
  provider     = google-beta
  display_name = "tf-test-secpol-%{random_suffix}"
  parent       = "organizations/123456789"
}

resource "google_compute_organization_security_policy" "policy" {
  provider = google-beta
  display_name = "tf-test%{random_suffix}"
  parent       = google_folder.security_policy_target.name
}

resource "google_compute_organization_security_policy_rule" "policy" {
  provider = google-beta
  policy_id = google_compute_organization_security_policy.policy.id
  action = "allow"

  direction = "INGRESS"
  enable_logging = true
  match {
    config {
      src_ip_ranges = ["192.168.0.0/16", "10.0.0.0/8"]
      layer4_config {
        ip_protocol = "tcp"
        ports = ["22"]
      }
      layer4_config {
        ip_protocol = "icmp"
      }
    }
  }
  priority = 100
}

resource "google_compute_organization_security_policy_association" "policy" {
  provider = google-beta
  name          = "tf-test%{random_suffix}"
  attachment_id = google_compute_organization_security_policy.policy.parent
  policy_id     = google_compute_organization_security_policy.policy.id
}
```
