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
resource "google_certificate_manager_certificate" "default" {
  name        = "issuance-config-cert"
  description = "The default cert"
  scope       = "EDGE_CACHE"
  managed {
    domains = [
        "terraform.subdomain1.com"
      ]
    issuance_config = google_certificate_manager_certificate_issuance_config.issuanceconfig.id
  }
}



# creating certificate_issuance_config to use it in the managed certificate
resource "google_certificate_manager_certificate_issuance_config" "issuanceconfig" {
  name    = "issuanceconfigtestterraform"
  description = "sample description for the certificate issuanceConfigs"
  certificate_authority_config {
    certificate_authority_service_config {
        ca_pool = google_privateca_ca_pool.pool.id
    }
  }
  lifetime = "1814400s"
  rotation_window_percentage = 34
  key_algorithm = "ECDSA_P256"
  depends_on=[google_privateca_certificate_authority.ca_authority]
}
  
resource "google_privateca_ca_pool" "pool" {
  name     = "my-ca-pool"
  location = "us-central1"
  tier     = "ENTERPRISE"
}

resource "google_privateca_certificate_authority" "ca_authority" {
  location = "us-central1"
  pool = google_privateca_ca_pool.pool.name
  certificate_authority_id = "my-ca"
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
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = true
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
```
