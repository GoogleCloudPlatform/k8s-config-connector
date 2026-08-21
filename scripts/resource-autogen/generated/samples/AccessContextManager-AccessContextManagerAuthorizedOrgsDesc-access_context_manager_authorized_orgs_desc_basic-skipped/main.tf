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
resource "google_access_context_manager_authorized_orgs_desc" "authorized-orgs-desc" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}/authorizedOrgsDescs/fakeDescName"
  authorization_type = "AUTHORIZATION_TYPE_TRUST"
  asset_type = "ASSET_TYPE_CREDENTIAL_STRENGTH"
  authorization_direction = "AUTHORIZATION_DIRECTION_TO"
  orgs = ["organizations/12345", "organizations/98765"]
}

resource "google_access_context_manager_access_policy" "test-access" {
  parent = "organizations/"
  title  = "my policy"
}
```
