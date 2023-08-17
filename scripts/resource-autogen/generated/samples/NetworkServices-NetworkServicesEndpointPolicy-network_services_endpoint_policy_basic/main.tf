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
resource "google_network_services_endpoint_policy" "default" {
  provider               = google-beta
  name                   = "my-endpoint-policy"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  type                   = "SIDECAR_PROXY"
  traffic_port_selector {
    ports = ["8081"]
  }
  endpoint_matcher {
    metadata_label_matcher {
      metadata_label_match_criteria = "MATCH_ANY"
      metadata_labels {
        label_name = "foo"
        label_value = "bar"
      }
    }
  }
}
  
```
