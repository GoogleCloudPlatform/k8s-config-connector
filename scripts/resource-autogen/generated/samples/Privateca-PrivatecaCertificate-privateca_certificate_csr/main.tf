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
resource "google_privateca_ca_pool" "default" {
  location = "us-central1"
  name = "my-pool"
  tier = "ENTERPRISE"
}

resource "google_privateca_certificate_authority" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority_id = "my-authority"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-certificate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    x509_config {
      ca_options {
        # is_ca *MUST* be true for certificate authorities
        is_ca = true
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
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable CA deletion related safe checks for easier cleanup.
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true
}


resource "google_privateca_certificate" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority = google_privateca_certificate_authority.default.certificate_authority_id
  name = "my-certificate"
  lifetime = "860s"
  pem_csr = file("test-fixtures/rsa_csr.pem")
}
```
