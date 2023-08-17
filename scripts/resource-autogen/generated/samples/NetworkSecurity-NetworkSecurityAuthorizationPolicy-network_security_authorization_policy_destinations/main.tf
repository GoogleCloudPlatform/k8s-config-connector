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
resource "google_network_security_authorization_policy" "default" {
  provider               = google-beta
  name                   = "my-authorization-policy"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  action                 = "ALLOW"
  rules {
    sources {
      principals = ["namespace/*"]
      ip_blocks = ["1.2.3.0/24"]
    }
    destinations {
      hosts = ["mydomain.*"]
      ports = [8080]
      methods = ["GET"]
      http_header_match {
        header_name = ":method"
        regex_match = "GET"
      }
    }
  }
}
```
