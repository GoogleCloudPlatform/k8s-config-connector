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
resource "google_network_services_http_route" "default" {
  provider               = google-beta
  name                   = "my-http-route"
  labels                 = {
    foo = "bar"
  }
  description             = "my description"
  hostnames               = ["example"]
  rules                   {
    matches {
      headers {
        header = "header"
        invert_match = false
        regex_match = "header-value"
      }
      query_parameters {
        query_parameter = "key"
        exact_match = "value"
      }
      prefix_match = "example"
      ignore_case = false
    }
    matches {
      headers {
        header = "header"
        invert_match = false
        present_match = true
      }
      query_parameters {
        query_parameter = "key"
        regex_match = "value"
      }
      regex_match = "example"
      ignore_case = false
    }
    matches {
      headers {
        header = "header"
        invert_match = false
        present_match = true
      }
      query_parameters {
        query_parameter = "key"
        present_match = true
      }
      full_path_match = "example"
      ignore_case = false
    }
    action {
      redirect {
        host_redirect = "new-host"
        path_redirect =  "new-path"
        prefix_rewrite =  "new-prefix"
        https_redirect = true
        strip_query = true
        port_redirect = 8081
      }
      url_rewrite {
        path_prefix_rewrite = "new-prefix"
        host_rewrite = "new-host"
      }
      retry_policy {
          retry_conditions = ["server_error"]
          num_retries = 1
          per_try_timeout =  "1s"
      }
      request_mirror_policy {
        destination {
          service_name = "new"
          weight = 1
        }
      }
      cors_policy {
        allow_origins = ["example"]
        allow_methods = ["GET", "PUT"]
        allow_headers = ["version", "type"]
        expose_headers = ["version", "type"]
        max_age = "1s"
        allow_credentials = true
        disabled = false
      }
    }
  }
}
```
