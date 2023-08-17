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
  provider = google-beta
  name      = "my-basic-ca-pool"
  location  = "us-central1"
  tier     = "DEVOPS"
  publishing_options {
    publish_ca_cert = false
    publish_crl = false
  }
  issuance_policy {
    maximum_lifetime = "1209600s"
    baseline_values {
      ca_options {
        is_ca = false
      }
      key_usage {
        base_key_usage {}
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
}


resource "google_privateca_certificate_authority" "default" {
  provider = google-beta
  pool = google_privateca_ca_pool.default.name
  certificate_authority_id = "my-basic-certificate-authority"
  location = "us-central1"
  lifetime = "86400s"
  type = "SELF_SIGNED"
  deletion_protection = false
  skip_grace_period = true
  ignore_active_certificates_on_deletion = true
  config {
    subject_config {
      subject {
        organization = "Test LLC"
        common_name = "my-ca"
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
          server_auth = false
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
}

resource "google_project_service_identity" "ns_sa" {
  provider = google-beta

  service = "networksecurity.googleapis.com"
}

resource "google_privateca_ca_pool_iam_member" "tls_inspection_permission" {
  provider = google-beta

  ca_pool = google_privateca_ca_pool.default.id
  role = "roles/privateca.certificateManager"
  member = "serviceAccount:${google_project_service_identity.ns_sa.email}"
}

resource "google_network_security_tls_inspection_policy" "default" {
  provider = google-beta
  name     = "my-tls-inspection-policy"
  location = "us-central1"
  ca_pool  = google_privateca_ca_pool.default.id
  depends_on = [google_privateca_ca_pool.default, google_privateca_certificate_authority.default, google_privateca_ca_pool_iam_member.tls_inspection_permission]
}

resource "google_network_security_gateway_security_policy" "default" {
  provider    = google-beta
  name        = "my-gateway-security-policy"
  location    = "us-central1"
  description = "my description"
  tls_inspection_policy = google_network_security_tls_inspection_policy.default.id
  depends_on = [google_network_security_tls_inspection_policy.default]
}
```
