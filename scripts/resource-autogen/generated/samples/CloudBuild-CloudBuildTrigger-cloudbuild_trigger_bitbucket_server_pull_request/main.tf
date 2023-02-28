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
resource "google_cloudbuild_trigger" "bbs-pull-request-trigger" {
  name        = "terraform-bbs-pull-request-trigger"
  location    = "us-central1"

  bitbucket_server_trigger_config {
    repo_slug = "terraform-provider-google"
    project_key = "STAG"
    bitbucket_server_config_resource = "projects/123456789/locations/us-central1/bitbucketServerConfigs/myBitbucketConfig"
    pull_request {
        branch = "^master$"
        invert_regex = false
        comment_control = "COMMENTS_ENABLED"
    }
  }

  filename = "cloudbuild.yaml"
}
```
