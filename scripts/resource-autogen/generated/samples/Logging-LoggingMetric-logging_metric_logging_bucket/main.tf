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
resource "google_logging_project_bucket_config" "logging_metric" {
    location  = "global"
    project   = "my-project-name"
    bucket_id = "_Default"
}

resource "google_logging_metric" "logging_metric" {
  name        = "my-(custom)/metric"
  filter      = "resource.type=gae_app AND severity>=ERROR"
  bucket_name = google_logging_project_bucket_config.logging_metric.id
}
```
