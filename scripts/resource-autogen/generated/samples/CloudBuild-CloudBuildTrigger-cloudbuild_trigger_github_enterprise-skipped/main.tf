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
resource "google_cloudbuild_trigger" "ghe-trigger" {
  name        = "ghe-trigger"
  location    = "us-central1"

  github {
    owner = "hashicorp"
    name  = "terraform-provider-google"
    push {
      branch = "^main$"
    }
    enterprise_config_resource_name = "projects/123456789/locations/us-central1/githubEnterpriseConfigs/configID"
  }

  filename = "cloudbuild.yaml"
}
```
