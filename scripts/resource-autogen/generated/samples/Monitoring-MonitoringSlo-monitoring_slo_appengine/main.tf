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
data "google_monitoring_app_engine_service" "default" {
  module_id = "default"
}

resource "google_monitoring_slo" "appeng_slo" {
  service = data.google_monitoring_app_engine_service.default.service_id

  slo_id = "ae-slo"
  display_name = "Terraform Test SLO for App Engine"

  goal = 0.9
  calendar_period = "DAY"

  basic_sli {
    latency {
      threshold = "1s"
    }
  }

  user_labels = {
    my_key       = "my_value"
    my_other_key = "my_other_value"
  }
}
```
