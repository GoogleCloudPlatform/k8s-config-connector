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
resource "google_binary_authorization_attestor" "attestor" {
  name = "attestor"
  attestation_authority_note {
    note_reference = google_container_analysis_note.note.name
    public_keys {
      id = data.google_kms_crypto_key_version.version.id
      pkix_public_key {
        public_key_pem      = data.google_kms_crypto_key_version.version.public_key[0].pem
        signature_algorithm = data.google_kms_crypto_key_version.version.public_key[0].algorithm
      }
    }
  }
}

resource "google_container_analysis_note" "note" {
  name = "attestation-note"
  attestation_authority {
    hint {
      human_readable_name = "Attestor Note"
    }
  }
}

data "google_kms_key_ring" "keyring" {
  name = "my-key-ring"
  location = "global"
}

data "google_kms_crypto_key" "crypto-key" {
  name     = "my-key"
  key_ring = data.google_kms_key_ring.keyring.id
}

data "google_kms_crypto_key_version" "version" {
  crypto_key = data.google_kms_crypto_key.crypto-key.id
}

resource "google_container_analysis_occurrence" "occurrence" {
  resource_uri = "gcr.io/my-project/my-image"
  note_name = google_container_analysis_note.note.id

  // See "Creating Attestations" Guide for expected
  // payload and signature formats.
  attestation {
    serialized_payload = filebase64("path/to/my/payload.json")
    signatures {
      public_key_id = data.google_kms_crypto_key_version.version.id
      serialized_payload = filebase64("path/to/my/payload.json.sig")
    }
  }
}
```
