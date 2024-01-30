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
locals {
  project = "my-project-name" # Google Cloud Platform Project ID
}

data "google_project" "project" {
  provider = google-beta
}

resource "google_storage_bucket" "bucket" {
  provider = google-beta

  name     = "${local.project}-gcf-source"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  provider = google-beta

  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "function-source.zip"  # Add path to the zipped function source code
}

resource "google_project_service_identity" "ea_sa" {
  provider = google-beta

  project = data.google_project.project.project_id
  service = "eventarc.googleapis.com"
}

resource "google_artifact_registry_repository" "unencoded-ar-repo" {
  provider = google-beta

  repository_id = "ar-repo"
  location = "us-central1"
  format = "DOCKER"
}

resource "google_artifact_registry_repository_iam_binding" "binding" {
  provider = google-beta

  location = google_artifact_registry_repository.encoded-ar-repo.location
  repository = google_artifact_registry_repository.encoded-ar-repo.name
  role = "roles/artifactregistry.admin"
  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcf-admin-robot.iam.gserviceaccount.com",
  ]
}

resource "google_kms_crypto_key_iam_binding" "gcf_cmek_keyuser" {
  provider = google-beta

  crypto_key_id = "cmek-key"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcf-admin-robot.iam.gserviceaccount.com",
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-artifactregistry.iam.gserviceaccount.com",
    "serviceAccount:service-${data.google_project.project.number}@gs-project-accounts.iam.gserviceaccount.com",
    "serviceAccount:service-${data.google_project.project.number}@serverless-robot-prod.iam.gserviceaccount.com",
    "serviceAccount:${google_project_service_identity.ea_sa.email}",
  ]

  depends_on = [
    google_project_service_identity.ea_sa
  ]
}

resource "google_artifact_registry_repository" "encoded-ar-repo" {
  provider = google-beta

  location = "us-central1"
  repository_id = "cmek-repo"
  format = "DOCKER"
  kms_key_name = "cmek-key"
  depends_on = [
    google_kms_crypto_key_iam_binding.gcf_cmek_keyuser
  ]
}

resource "google_cloudfunctions2_function" "function" {
  provider = google-beta

  name = "function-cmek"
  location = "us-central1"
  description = "CMEK function"
  kms_key_name = "cmek-key"

  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point
    docker_repository = google_artifact_registry_repository.encoded-ar-repo.id

    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }

  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60
  }

  depends_on = [
    google_kms_crypto_key_iam_binding.gcf_cmek_keyuser
  ]

}
```
