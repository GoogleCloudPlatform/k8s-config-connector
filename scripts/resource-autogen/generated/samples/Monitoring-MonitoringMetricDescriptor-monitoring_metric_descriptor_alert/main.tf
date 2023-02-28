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
resource "google_monitoring_metric_descriptor" "with_alert" {
  description = "Daily sales records from all branch stores."
  display_name = "metric-descriptor"
  type = "custom.googleapis.com/stores/daily_sales"
  metric_kind = "GAUGE"
  value_type = "DOUBLE"
  unit = "{USD}"
}

resource "google_monitoring_alert_policy" "alert_policy" {
  display_name = "metric-descriptor"
  combiner     = "OR"
  conditions {
    display_name = "test condition"
    condition_threshold {
      filter     = "metric.type=\"${google_monitoring_metric_descriptor.with_alert.type}\" AND resource.type=\"gce_instance\""
      duration   = "60s"
      comparison = "COMPARISON_GT"
    }
  }
}
```
