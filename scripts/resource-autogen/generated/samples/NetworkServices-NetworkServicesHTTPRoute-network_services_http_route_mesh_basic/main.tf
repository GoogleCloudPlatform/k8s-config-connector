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
resource "google_network_services_mesh" "default" {
  provider    = google-beta
  name        = "my-http-route"
  labels      = {
    foo = "bar"
  }
  description = "my description"
}

resource "google_network_services_http_route" "default" {
  provider               = google-beta
  name                   = "my-http-route"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  hostnames               = ["example"]
  meshes = [
    google_network_services_mesh.default.id
  ]
  rules                   {
    matches {
      query_parameters {
        query_parameter = "key"
        exact_match = "value"
      }
      full_path_match = "example"
    }
  }
}
```
