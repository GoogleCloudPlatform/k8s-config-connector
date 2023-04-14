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
# [START functions_v2_basic_auditlogs]
# This example follows the examples shown in this Google Cloud Community blog post
# https://medium.com/google-cloud/applying-a-path-pattern-when-filtering-in-eventarc-f06b937b4c34
# and the docs:
# https://cloud.google.com/eventarc/docs/path-patterns

resource "google_storage_bucket" "source-bucket" {
  name     = "gcf-source-bucket"
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.source-bucket.name
  source = "function-source.zip"  # Add path to the zipped function source code
}

resource "google_service_account" "account" {
  account_id   = "gcf-sa"
  display_name = "Test Service Account - used for both the cloud function and eventarc trigger in the test"
}

# Note: The right way of listening for Cloud Storage events is to use a Cloud Storage trigger.
# Here we use Audit Logs to monitor the bucket so path patterns can be used in the example of
# google_cloudfunctions2_function below (Audit Log events have path pattern support)
resource "google_storage_bucket" "audit-log-bucket" {
  name     = "gcf-auditlog-bucket"
  location = "us-central1"  # The trigger must be in the same location as the bucket
  uniform_bucket_level_access = true
}

# Permissions on the service account used by the function and Eventarc trigger
resource "google_project_iam_member" "invoking" {
  project = "my-project-name"
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.account.email}"
}

resource "google_project_iam_member" "event-receiving" {
  project = "my-project-name"
  role    = "roles/eventarc.eventReceiver"
  member  = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_project_iam_member.invoking]
}

resource "google_project_iam_member" "artifactregistry-reader" {
  project = "my-project-name"
  role     = "roles/artifactregistry.reader"
  member   = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_project_iam_member.event-receiving]
}

resource "google_cloudfunctions2_function" "function" {
  depends_on = [
    google_project_iam_member.event-receiving,
    google_project_iam_member.artifactregistry-reader,
  ]
  name = "gcf-function"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime     = "nodejs12"
    entry_point = "entryPoint" # Set the entry point in the code
    environment_variables = {
      BUILD_CONFIG_TEST = "build_test"
    }
    source {
      storage_source {
        bucket = google_storage_bucket.source-bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 3
    min_instance_count = 1
    available_memory    = "256M"
    timeout_seconds     = 60
    environment_variables = {
        SERVICE_CONFIG_TEST = "config_test"
    }
    ingress_settings = "ALLOW_INTERNAL_ONLY"
    all_traffic_on_latest_revision = true
    service_account_email = google_service_account.account.email
  }

  event_trigger {
    trigger_region = "us-central1" # The trigger must be in the same location as the bucket
    event_type = "google.cloud.audit.log.v1.written"
    retry_policy = "RETRY_POLICY_RETRY"
    service_account_email = google_service_account.account.email
    event_filters {
      attribute = "serviceName"
      value = "storage.googleapis.com"
    }
    event_filters {
      attribute = "methodName"
      value = "storage.objects.create"
    }
    event_filters {
      attribute = "resourceName"
      value = "/projects/_/buckets/${google_storage_bucket.audit-log-bucket.name}/objects/*.txt" # Path pattern selects all .txt files in the bucket
      operator = "match-path-pattern" # This allows path patterns to be used in the value field
    }
  }
}
# [END functions_v2_basic_auditlogs]
```
