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
resource "google_compute_region_security_policy" "region-sec-policy-user-defined-fields" {
  provider    = google-beta  

  name        = "my-sec-policy-user-defined-fields"
  description = "with user defined fields"
  type        = "CLOUD_ARMOR_NETWORK"
  user_defined_fields {
    name = "SIG1_AT_0"
    base = "UDP"
    offset = 8
    size = 2
    mask = "0x8F00"
  }
  user_defined_fields {
    name = "SIG2_AT_8"
    base = "UDP"
    offset = 16
    size = 4
    mask = "0xFFFFFFFF"
  }
}
```
