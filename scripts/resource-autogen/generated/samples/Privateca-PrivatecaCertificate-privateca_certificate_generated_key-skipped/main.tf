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
  name = "default"
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

resource "tls_private_key" "cert_key" {
  algorithm = "RSA"
}

resource "google_privateca_certificate" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority = google_privateca_certificate_authority.default.certificate_authority_id
  lifetime = "86000s"
  name = "cert-1"
  config {
    subject_config  {
      subject {
        common_name = "san1.example.com"
        country_code = "us"
        organization = "google"
        organizational_unit = "enterprise"
        locality = "mountain view"
        province = "california"
        street_address = "1600 amphitheatre parkway"
      } 
      subject_alt_name {
        email_addresses = ["email@example.com"]
        ip_addresses = ["127.0.0.1"]
        uris = ["http://www.ietf.org/rfc/rfc3986.txt"]
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
      name_constraints {
        critical                  = true
        permitted_dns_names       = ["*.example.com"]
        excluded_dns_names        = ["*.deny.example.com"]
        permitted_ip_ranges       = ["10.0.0.0/8"]
        excluded_ip_ranges        = ["10.1.1.0/24"]
        permitted_email_addresses = [".example.com"]
        excluded_email_addresses  = [".deny.example.com"]
        permitted_uris            = [".example.com"]
        excluded_uris             = [".deny.example.com"]
      }
    }
    public_key {
      format = "PEM"
      key = base64encode(tls_private_key.cert_key.public_key_pem)
    }
  }
}
```
