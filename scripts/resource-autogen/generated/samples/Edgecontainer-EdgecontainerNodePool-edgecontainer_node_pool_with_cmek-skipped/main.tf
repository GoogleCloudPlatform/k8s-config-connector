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
resource "google_edgecontainer_cluster" "cluster" {
  name = "default"
  location = "us-central1"

  authorization {
    admin_users {
      username = "admin@hashicorptest.com"
    }
  }

  networking {
    cluster_ipv4_cidr_blocks = ["10.0.0.0/16"]
    services_ipv4_cidr_blocks = ["10.1.0.0/16"]
  }

  fleet {
    project = "projects/${data.google_project.project.number}"
  }
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-edgecontainer.iam.gserviceaccount.com"
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "key"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  name     = "keyring"
  location = "us-central1"
}

resource "google_edgecontainer_node_pool" "default" {
  depends_on = [google_kms_crypto_key_iam_member.crypto_key]

  name = "nodepool-1"
  cluster = google_edgecontainer_cluster.cluster.name
  location = "us-central1"
  node_location = "us-central1-edge-example-edgesite"
  node_count = 3

  local_disk_encryption {
    kms_key = google_kms_crypto_key.crypto_key.id
  }
}

data "google_project" "project" {}
```
