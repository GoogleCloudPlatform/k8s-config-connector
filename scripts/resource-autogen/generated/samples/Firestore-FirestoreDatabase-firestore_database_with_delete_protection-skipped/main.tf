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
resource "google_firestore_database" "database" {
  project                           = google_project.project.project_id
  name                              = "my-database"
  location_id                       = "nam5"
  type                              = "FIRESTORE_NATIVE"

  # Prevents accidental deletion of the database.
  # To delete the database, first set this field to `DELETE_PROTECTION_DISABLED`, apply the changes.
  # Then delete the database resource and apply the changes again.
  delete_protection_state           = "DELETE_PROTECTION_ENABLED"

  depends_on = [google_project_service.firestore]
}
```
