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
resource "google_certificate_manager_certificate_map" "certificate_map" {
  name        = "cert-map-entry"
  description = "My acceptance test certificate map"
   labels      = {
    "terraform" : true,
    "acc-test"  : true,
  }
}

resource "google_certificate_manager_certificate_map_entry" "default" {
  name        = "cert-map-entry"
  description = "My acceptance test certificate map entry"
  map = google_certificate_manager_certificate_map.certificate_map.name 
  labels      = {
    "terraform" : true,
    "acc-test"  : true,
  }
  certificates = [google_certificate_manager_certificate.certificate.id]
  matcher = "PRIMARY"
}

resource "google_certificate_manager_certificate" "certificate" {
  name        = "cert-map-entry"
  description = "The default cert"
  scope       = "DEFAULT"
  managed {
    domains = [
      google_certificate_manager_dns_authorization.instance.domain,
      google_certificate_manager_dns_authorization.instance2.domain,
      ]
    dns_authorizations = [
      google_certificate_manager_dns_authorization.instance.id,
      google_certificate_manager_dns_authorization.instance2.id,
      ]
  }
}


resource "google_certificate_manager_dns_authorization" "instance" {
  name        = "dns-auth"
  description = "The default dnss"
  domain      = "subdomain.hashicorptest.com"
}

resource "google_certificate_manager_dns_authorization" "instance2" {
  name        = "dns-auth2"
  description = "The default dnss"
  domain      = "subdomain2.hashicorptest.com"
}
```
