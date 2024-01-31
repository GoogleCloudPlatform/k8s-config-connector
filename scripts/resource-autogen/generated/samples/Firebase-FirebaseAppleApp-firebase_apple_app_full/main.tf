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
resource "google_firebase_apple_app" "full" {
  provider = google-beta
  project = "my-project-name"
  display_name = "Display Name Full"
  bundle_id = "apple.app.12345"
  app_store_id = "12345"
  team_id = "9987654321"
  api_key_id = google_apikeys_key.apple.uid
}

resource "google_apikeys_key" "apple" {
  provider = google-beta

  name         = "api-key"
  display_name = "Display Name Full"
  project = "my-project-name"
  
  restrictions {
    ios_key_restrictions {
      allowed_bundle_ids = ["apple.app.12345"]
    }
  }
}
```
