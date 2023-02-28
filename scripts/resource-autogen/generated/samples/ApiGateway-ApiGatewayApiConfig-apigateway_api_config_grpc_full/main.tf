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
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "api-cfg"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "cfg"

  grpc_services {
    file_descriptor_set {
      path = "api_descriptor.pb"
      contents = filebase64("test-fixtures/apigateway/api_descriptor.pb")
    }
    source {
      path = "bookstore.proto"
      contents = filebase64("test-fixtures/apigateway/bookstore.proto")
    }
  }
  managed_service_configs {
    path = "api_config.yaml"
    contents = base64encode(<<-EOF
      type: google.api.Service
      config_version: 3
      name: ${google_api_gateway_api.api_cfg.managed_service}
      title: gRPC API example
      apis:
        - name: endpoints.examples.bookstore.Bookstore
      usage:
        rules:
        - selector: endpoints.examples.bookstore.Bookstore.ListShelves
          allow_unregistered_calls: true
      backend:
        rules:
          - selector: "*"
            address: grpcs://example.com
            disable_auth: true

    EOF
    )
  }
  lifecycle {
    create_before_destroy = true
  }
}
```
