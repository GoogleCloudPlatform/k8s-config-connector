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
data "google_project" "project" {
}

resource "google_sql_database_instance" "instance" {
    name             = "my-instance"
    database_version = "MYSQL_8_0"
    region           = "us-central1"
    settings {
        tier = "db-f1-micro"
        backup_configuration {
            enabled            = true
            binary_log_enabled = true
        }

        ip_configuration {

            // Datastream IPs will vary by region.
            authorized_networks {
                value = "34.71.242.81"
            }

            authorized_networks {
                value = "34.72.28.29"
            }

            authorized_networks {
                value = "34.67.6.157"
            }

            authorized_networks {
                value = "34.67.234.134"
            }

            authorized_networks {
                value = "34.72.239.218"
            }
        }
    }

    deletion_protection  = true
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name     = "user"
    instance = google_sql_database_instance.instance.name
    host     = "%"
    password = random_password.pwd.result
}

resource "google_datastream_connection_profile" "source_connection_profile" {
    display_name          = "Source connection profile"
    location              = "us-central1"
    connection_profile_id = "source-profile"

    mysql_profile {
        hostname = google_sql_database_instance.instance.public_ip_address
        username = google_sql_user.user.name
        password = google_sql_user.user.password
    }
}

resource "google_storage_bucket" "bucket" {
  name                        = "my-bucket"
  location                    = "US"
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_iam_member" "viewer" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.objectViewer"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_storage_bucket_iam_member" "creator" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.objectCreator"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_storage_bucket_iam_member" "reader" {
    bucket = google_storage_bucket.bucket.name
    role   = "roles/storage.legacyBucketReader"
    member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-datastream.iam.gserviceaccount.com"
}

resource "google_datastream_connection_profile" "destination_connection_profile" {
    display_name          = "Connection profile"
    location              = "us-central1"
    connection_profile_id = "destination-profile"

    gcs_profile {
        bucket    = google_storage_bucket.bucket.name
        root_path = "/path"
    }
}

resource "google_datastream_stream" "default" {
    stream_id = "my-stream"
    location = "us-central1"
    display_name = "my stream"
    source_config {
        source_connection_profile = google_datastream_connection_profile.source_connection_profile.id
        mysql_source_config {
          max_concurrent_backfill_tasks = 15
        }
    }
    destination_config {
        destination_connection_profile = google_datastream_connection_profile.destination_connection_profile.id
        gcs_destination_config {
            avro_file_format {}
        }
    }

    backfill_none {
    }
}
```
