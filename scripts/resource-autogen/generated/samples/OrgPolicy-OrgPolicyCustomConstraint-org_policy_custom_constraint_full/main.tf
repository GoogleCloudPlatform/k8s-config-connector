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
resource "google_org_policy_custom_constraint" "constraint" {
  provider = google-beta

  name         = "custom.disableGkeAutoUpgrade"
  parent       = "organizations/123456789"
  display_name = "Disable GKE auto upgrade"
  description  = "Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced."

  action_type    = "ALLOW"
  condition      = "resource.management.autoUpgrade == false"
  method_types   = ["CREATE", "UPDATE"]
  resource_types = ["container.googleapis.com/NodePool"]
}

resource "google_org_policy_policy" "bool" {
  provider = google-beta

  name   = "organizations/123456789/policies/${google_org_policy_custom_constraint.constraint.name}"
  parent = "organizations/123456789"

  spec {
    rules {
      enforce = "TRUE"
    }
  }
}
```
