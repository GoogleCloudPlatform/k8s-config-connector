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
resource "google_project" "default" {
	provider = google-beta

	project_id = "my-project"
	name       = "my-project"
	org_id     = "123456789"

	labels = {
		"firebase" = "enabled"
	}
}

resource "google_firebase_project" "default" {
	provider = google-beta
	project  = google_project.default.project_id
}

resource "google_firebase_web_app" "basic" {
	provider = google-beta
	project = google_project.default.project_id
	display_name = "Display Name Basic"
	deletion_policy = "DELETE"

	depends_on = [google_firebase_project.default]
}

data "google_firebase_web_app_config" "basic" {
  provider   = google-beta
  web_app_id = google_firebase_web_app.basic.app_id
}

resource "google_storage_bucket" "default" {
    provider = google-beta
    name     = "fb-webapp-"
    location = "US"
}

resource "google_storage_bucket_object" "default" {
    provider = google-beta
    bucket = google_storage_bucket.default.name
    name = "firebase-config.json"

    content = jsonencode({
        appId              = google_firebase_web_app.basic.app_id
        apiKey             = data.google_firebase_web_app_config.basic.api_key
        authDomain         = data.google_firebase_web_app_config.basic.auth_domain
        databaseURL        = lookup(data.google_firebase_web_app_config.basic, "database_url", "")
        storageBucket      = lookup(data.google_firebase_web_app_config.basic, "storage_bucket", "")
        messagingSenderId  = lookup(data.google_firebase_web_app_config.basic, "messaging_sender_id", "")
        measurementId      = lookup(data.google_firebase_web_app_config.basic, "measurement_id", "")
    })
}
```
