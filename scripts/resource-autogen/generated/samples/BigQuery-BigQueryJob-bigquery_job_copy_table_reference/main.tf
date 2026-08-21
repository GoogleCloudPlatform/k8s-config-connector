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
resource "google_bigquery_table" "source" {
  deletion_protection = false
  count = length(google_bigquery_dataset.source)

  dataset_id = google_bigquery_dataset.source[count.index].dataset_id
  table_id   = "job_copy_${count.index}_table"

  schema = <<EOF
[
  {
    "name": "name",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "post_abbr",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "date",
    "type": "DATE",
    "mode": "NULLABLE"
  }
]
EOF
}

resource "google_bigquery_dataset" "source" {
  count = 2

  dataset_id                  = "job_copy_${count.index}_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "US"
}

resource "google_bigquery_table" "dest" {
  deletion_protection = false
  dataset_id = google_bigquery_dataset.dest.dataset_id
  table_id   = "job_copy_dest_table"

  schema = <<EOF
[
  {
    "name": "name",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "post_abbr",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "date",
    "type": "DATE",
    "mode": "NULLABLE"
  }
]
EOF

  encryption_configuration {
    kms_key_name = google_kms_crypto_key.crypto_key.id
  }

  depends_on = ["google_project_iam_member.encrypt_role"]
}

resource "google_bigquery_dataset" "dest" {
  dataset_id    = "job_copy_dest_dataset"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "example-key"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  name     = "example-keyring"
  location = "global"
}

data "google_project" "project" {
  project_id = "my-project-name"
}

resource "google_project_iam_member" "encrypt_role" {
  project = data.google_project.project.project_id
  role = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member = "serviceAccount:bq-${data.google_project.project.number}@bigquery-encryption.iam.gserviceaccount.com"
}

resource "google_bigquery_job" "job" {
  job_id     = "job_copy"

  copy {
    source_tables {
      table_id   = google_bigquery_table.source.0.id
    }

    source_tables {
      table_id   = google_bigquery_table.source.1.id
    }

    destination_table {
      table_id   = google_bigquery_table.dest.id
    }

    destination_encryption_configuration {
      kms_key_name = google_kms_crypto_key.crypto_key.id
    }
  }

  depends_on = ["google_project_iam_member.encrypt_role"]
}
```
