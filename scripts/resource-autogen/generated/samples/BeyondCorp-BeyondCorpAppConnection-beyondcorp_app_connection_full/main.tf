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
resource "google_service_account" "service_account" {
  account_id   = "my-account"
  display_name = "Test Service Account"
}

resource "google_beyondcorp_app_gateway" "app_gateway" {
  name = "my-app-gateway"
  type = "TCP_PROXY"
  host_type = "GCP_REGIONAL_MIG"
}

resource "google_beyondcorp_app_connector" "app_connector" {
  name = "my-app-connector"
  principal_info {
    service_account {
     email = google_service_account.service_account.email
    }
  }
}

resource "google_beyondcorp_app_connection" "app_connection" {
  name = "my-app-connection"
  type = "TCP_PROXY"
  display_name = "some display name"
  application_endpoint {
    host = "foo-host"
    port = 8080
  }
  connectors = [google_beyondcorp_app_connector.app_connector.id]
  gateway {
    app_gateway = google_beyondcorp_app_gateway.app_gateway.id
  }
  labels = {
    foo = "bar"
    bar = "baz"
  }
}
```
