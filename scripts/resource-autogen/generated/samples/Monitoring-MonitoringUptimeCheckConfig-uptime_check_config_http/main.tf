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
resource "google_monitoring_uptime_check_config" "http" {
  display_name = "http-uptime-check"
  timeout      = "60s"

  http_check {
    path = "some-path"
    port = "8010"
    request_method = "POST"
    content_type = "URL_ENCODED"
    body = "Zm9vJTI1M0RiYXI="
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "my-project-name"
      host       = "192.168.1.1"
    }
  }

  content_matchers {
    content = "\"example\""
    matcher = "MATCHES_JSON_PATH"
    json_path_matcher {
      json_path = "$.path"
      json_matcher = "EXACT_MATCH"
    }
  }

  checker_type = "STATIC_IP_CHECKERS"
}
```
