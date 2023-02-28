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
data "google_project" "project" {
}

data "google_container_attached_versions" "versions" {
	location       = "us-west1"
	project        = data.google_project.project.project_id
}

resource "google_container_attached_cluster" "primary" {
  name     = "basic"
  location = "us-west1"
  project = data.google_project.project.project_id
  description = "Test cluster"
  distribution = "aks"
  oidc_config {
      issuer_url = "https://oidc.issuer.url"
  }
  platform_version = data.google_container_attached_versions.versions.valid_versions[0]
  fleet {
    project = "projects/${data.google_project.project.number}"
  }
}
```
