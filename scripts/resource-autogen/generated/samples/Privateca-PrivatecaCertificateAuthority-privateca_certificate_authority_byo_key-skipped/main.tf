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
resource "google_project_service_identity" "privateca_sa" {
  service = "privateca.googleapis.com"
}

resource "google_kms_crypto_key_iam_binding" "privateca_sa_keyuser_signerverifier" {
  crypto_key_id = "projects/keys-project/locations/us-central1/keyRings/key-ring/cryptoKeys/crypto-key"
  role          = "roles/cloudkms.signerVerifier"

  members = [
    "serviceAccount:${google_project_service_identity.privateca_sa.email}",
  ]
}

resource "google_kms_crypto_key_iam_binding" "privateca_sa_keyuser_viewer" {
  crypto_key_id = "projects/keys-project/locations/us-central1/keyRings/key-ring/cryptoKeys/crypto-key"
  role          = "roles/viewer"
  members = [
    "serviceAccount:${google_project_service_identity.privateca_sa.email}",
  ]
}

resource "google_privateca_certificate_authority" "default" {
  // This example assumes this pool already exists.
  // Pools cannot be deleted in normal test circumstances, so we depend on static pools
  pool = "ca-pool"
  certificate_authority_id = "my-certificate-authority"
  location = "us-central1"
  deletion_protection = "true"
  key_spec {
    cloud_kms_key_version = "projects/keys-project/locations/us-central1/keyRings/key-ring/cryptoKeys/crypto-key/cryptoKeyVersions/1"
  }

  config  {
    subject_config  {
      subject {
        organization = "Example, Org."
        common_name  = "Example Authority"
      }
    }
    x509_config {
      ca_options {
        # is_ca *MUST* be true for certificate authorities
        is_ca = true
        max_issuer_path_length = 10
      }
      key_usage {
        base_key_usage {
          # cert_sign and crl_sign *MUST* be true for certificate authorities
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }

  depends_on = [
    google_kms_crypto_key_iam_binding.privateca_sa_keyuser_signerverifier,
    google_kms_crypto_key_iam_binding.privateca_sa_keyuser_viewer,
  ]
}
```
