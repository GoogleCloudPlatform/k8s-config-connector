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
resource "google_certificate_manager_dns_authorization" "default" {
  name        = "dns-auth"
  description = "The default dnss"
  domain      = "subdomain.hashicorptest.com"
}

output "record_name_to_insert" {
 value = google_certificate_manager_dns_authorization.default.dns_resource_record.0.name
}

output "record_type_to_insert" {
 value = google_certificate_manager_dns_authorization.default.dns_resource_record.0.type
}

output "record_data_to_insert" {
 value = google_certificate_manager_dns_authorization.default.dns_resource_record.0.data
}
```
