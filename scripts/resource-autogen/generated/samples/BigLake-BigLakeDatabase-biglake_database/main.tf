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
resource "google_biglake_catalog" "catalog" {
    name = "my_catalog"
    location = "US"
}

resource "google_storage_bucket" "bucket" {
  name                        = "my_bucket"
  location                    = "US"
  force_destroy               = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "metadata_folder" {
  name    = "metadata/"
  content = " "
  bucket  = google_storage_bucket.bucket.name
}

resource "google_biglake_database" "database"  {
    name = "my_database"
    catalog = google_biglake_catalog.catalog.id
    type = "HIVE"
    hive_options {
        location_uri = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.metadata_folder.name}"
        parameters = {
          "owner": "John Doe"
        }
    }
}
```
