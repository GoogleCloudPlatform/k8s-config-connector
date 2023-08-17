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
resource "google_network_security_server_tls_policy" "default" {
  provider               = google-beta
  name                   = "my-server-tls-policy"
  labels                 = {
    foo = "bar"
  }
  description            = "my description"
  allow_open             = "false"
  server_certificate {
    certificate_provider_instance {
        plugin_instance = "google_cloud_private_spiffe"
      }
  }
  mtls_policy {
    client_validation_ca {
      grpc_endpoint {
        target_uri = "unix:mypath"
      }
    }
    client_validation_ca {
      grpc_endpoint {
        target_uri = "unix:abc/mypath"
      }
    }
    client_validation_ca {
      certificate_provider_instance {
        plugin_instance = "google_cloud_private_spiffe"
      }
    }
  }
}
```
