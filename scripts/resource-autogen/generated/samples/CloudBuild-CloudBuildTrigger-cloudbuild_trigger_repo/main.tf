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
resource "google_cloudbuildv2_connection" "my-connection" {
  provider = google-beta
  location = "us-central1"
  name = "my-connection"

  github_config {
    app_installation_id = 123123
    authorizer_credential {
      oauth_token_secret_version = "projects/my-project/secrets/github-pat-secret/versions/latest"
    }
  }
}

resource "google_cloudbuildv2_repository" "my-repository" {
  provider = google-beta
  name = "my-repo"
  parent_connection = google_cloudbuildv2_connection.my-connection.id
  remote_uri = "https://github.com/myuser/my-repo.git"
}

resource "google_cloudbuild_trigger" "repo-trigger" {
  provider = google-beta
  location = "us-central1"

  repository_event_config {
    repository = google_cloudbuildv2_repository.my-repository.id
    push {
      branch = "feature-.*"
    }
  }

  filename = "cloudbuild.yaml"
}
```
