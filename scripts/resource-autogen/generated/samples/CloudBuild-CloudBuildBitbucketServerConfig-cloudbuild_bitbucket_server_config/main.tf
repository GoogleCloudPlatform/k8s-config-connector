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
resource "google_cloudbuild_bitbucket_server_config" "bbs-config" {
    config_id = "mybbsconfig"
    location = "us-central1"
    host_uri = "https://bbs.com"
    secrets {
        admin_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        read_access_token_version_name = "projects/myProject/secrets/mybbspat/versions/1"
        webhook_secret_version_name = "projects/myProject/secrets/mybbspat/versions/1"
    }
    username = "test"
    api_key = "<api-key>"
}
```
