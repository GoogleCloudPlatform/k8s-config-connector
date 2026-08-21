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
resource "google_monitoring_custom_service" "customsrv" {
  service_id = "custom-srv-request-slos"
  display_name = "My Custom Service"
}

resource "google_monitoring_slo" "request_based_slo" {
  service = google_monitoring_custom_service.customsrv.service_id
  slo_id = "consumed-api-slo"
  display_name = "Terraform Test SLO with request based SLI (good total ratio)"

  goal = 0.9
  rolling_period_days = 30

  request_based_sli {
    distribution_cut {
          distribution_filter = "metric.type=\"serviceruntime.googleapis.com/api/request_latencies\" resource.type=\"api\"  "
          range {
            max = 0.5
          }
        }
  }
}
```
