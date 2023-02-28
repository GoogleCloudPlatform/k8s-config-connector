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
  name = "my-pool"
  location = "us-central1"
  tier = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = false
    publish_crl = true
  }
  labels = {
    foo = "bar"
  }
  issuance_policy {
    allowed_key_types {
      elliptic_curve {
        signature_algorithm = "ECDSA_P256"
      }
    }
    allowed_key_types {
      rsa {
        min_modulus_size = 5
        max_modulus_size = 10
      }
    }
    maximum_lifetime = "50000s"
    allowed_issuance_modes {
      allow_csr_based_issuance = true
      allow_config_based_issuance = true
    }
    identity_constraints {
      allow_subject_passthrough = true
      allow_subject_alt_names_passthrough = true
      cel_expression {
        expression = "subject_alt_names.all(san, san.type == DNS || san.type == EMAIL )"
        title = "My title"
      }
    }
    baseline_values {
      aia_ocsp_servers = ["example.com"]
      additional_extensions {
        critical = true
        value = "asdf"
        object_id {
          object_id_path = [1, 7]
        }
      }
      policy_ids {
        object_id_path = [1, 5]
      }
      policy_ids {
        object_id_path = [1, 5, 7]
      }
      ca_options {
        is_ca = true
        max_issuer_path_length = 10
      }
      key_usage {
        base_key_usage {
          digital_signature = true
          content_commitment = true
          key_encipherment = false
          data_encipherment = true
          key_agreement = true
          cert_sign = false
          crl_sign = true
          decipher_only = true
        }
        extended_key_usage {
          server_auth = true
          client_auth = false
          email_protection = true
          code_signing = true
          time_stamping = true
        }
      }
    }
  }
}
```
