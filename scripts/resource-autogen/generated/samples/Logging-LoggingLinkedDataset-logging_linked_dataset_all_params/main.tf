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
resource "google_logging_project_bucket_config" "logging_linked_dataset" {
  location         = "global"
  project          = "my-project-name"
  enable_analytics = true
  bucket_id        = "tftest%{random_suffix}"
}

resource "google_logging_linked_dataset" "logging_linked_dataset" {
  link_id     = "tftest%{random_suffix}"
  bucket      = "tftest%{random_suffix}"
  parent      = "projects/my-project-name"
  location    = "global"
  description = "Linked dataset test"

  # Using forced dependency in order to test use of the bucket name by itself without
  # referencing the entire resource link, which is what is outputted by the 
  # google_logging_project_bucket_config resource. Use of the outputted ID is tested in
  # the basic example.
  depends_on = ["google_logging_project_bucket_config.logging_linked_dataset"]
}
```
