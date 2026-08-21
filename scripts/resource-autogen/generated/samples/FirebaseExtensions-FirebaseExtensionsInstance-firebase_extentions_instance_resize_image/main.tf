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
resource "google_storage_bucket" "images" {
  provider                    = google-beta
  project                     = "my-project-name"
  name                        = "bucket-id"
  location                    = "US"
  uniform_bucket_level_access = true

  # Delete all objects when the bucket is deleted
  force_destroy = true
}

resource "google_firebase_extensions_instance" "resize_image" {
  provider = google-beta
  project  = "my-project-name"
  instance_id = "storage-resize-images"
  config {
    extension_ref = "firebase/storage-resize-images"
    extension_version = "0.1.37"

    # The following params apply to the firebase/storage-resize-images extension. 
    # Different extensions may have different params
    params = {
      DELETE_ORIGINAL_FILE = false
      MAKE_PUBLIC          = false
      IMAGE_TYPE           = false
      IS_ANIMATED          = true
      FUNCTION_MEMORY      = 1024
      DO_BACKFILL          = false
      IMG_SIZES            = "200x200"
      IMG_BUCKET           = google_storage_bucket.images.name
      LOCATION             = ""
    }

    system_params = {
      "firebaseextensions.v1beta.function/maxInstances"               = 3000
      "firebaseextensions.v1beta.function/memory"                     = 256
      "firebaseextensions.v1beta.function/minInstances"               = 0
      "firebaseextensions.v1beta.function/vpcConnectorEgressSettings" = "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED"
    }

    allowed_event_types = [
      "firebase.extensions.storage-resize-images.v1.complete"
    ]

    eventarc_channel = "projects/my-project-name/locations//channels/firebase"
  }
}
```
