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
resource "google_folder" "my_folder" {
  display_name = "my-folder"
  parent       = "organizations/123456789"
}

resource "google_folder_access_approval_settings" "folder_access_approval" {
  folder_id           = google_folder.my_folder.folder_id
  notification_emails = ["testuser@example.com", "example.user@example.com"]

  enrolled_services {
  	cloud_product = "all"
  }
}
```
