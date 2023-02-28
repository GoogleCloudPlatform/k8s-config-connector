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
resource "google_data_loss_prevention_stored_info_type" "large" {
	parent = "projects/my-project-name"
	description = "Description"
	display_name = "Displayname"

	large_custom_dictionary {
		cloud_storage_file_set {
			url = "gs://${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
		}
		output_path {
			path = "gs://${google_storage_bucket.bucket.name}/output/dictionary.txt"
		}
	}
}

resource "google_storage_bucket" "bucket" {
  name          = "tf-test-bucket"
  location      = "US"
  force_destroy = true
}

resource "google_storage_bucket_object" "object" {
  name   = "tf-test-object"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/dlp/words.txt"
}
```
