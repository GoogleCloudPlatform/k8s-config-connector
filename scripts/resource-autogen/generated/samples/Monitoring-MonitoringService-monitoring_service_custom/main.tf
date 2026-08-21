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
resource "google_monitoring_custom_service" "custom" {
  service_id = "custom-srv"
  display_name = "My Custom Service custom-srv"

  telemetry {
  	resource_name = "//product.googleapis.com/foo/foo/services/test"
  }

  user_labels = {
    my_key       = "my_value"
    my_other_key = "my_other_value"
  }
}
```
