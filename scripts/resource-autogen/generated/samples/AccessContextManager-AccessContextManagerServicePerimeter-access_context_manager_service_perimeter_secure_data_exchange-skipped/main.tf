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
resource "google_access_context_manager_service_perimeters" "secure-data-exchange" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"

  service_perimeters {
    name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/"
    title  = ""
    status {
      restricted_services = ["storage.googleapis.com"]
    }
  }

  service_perimeters {
    name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/servicePerimeters/"
    title  = ""
    status {
      restricted_services = ["bigtable.googleapis.com"]
      		vpcAccessibleServices = {
			enableRestriction = true
			allowedServices = ["bigquery.googleapis.com"]
		}
    }
  }
}

resource "google_access_context_manager_access_level" "access-level" {
  parent = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}"
  name   = "accessPolicies/${google_access_context_manager_access_policy.access-policy.name}/accessLevels/secure_data_exchange"
  title  = "secure_data_exchange"
  basic {
    conditions {
      device_policy {
        require_screen_lock = false
        os_constraints {
          os_type = "DESKTOP_CHROME_OS"
        }
      }
      regions = [
        "CH",
        "IT",
        "US",
      ]
    }
  }
}

resource "google_access_context_manager_access_policy" "access-policy" {
  parent = "organizations/123456789"
  title  = "my policy"
}

resource "google_access_context_manager_service_perimeter" "test-access" {
  parent         = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}"
  name           = "accessPolicies/${google_access_context_manager_access_policy.test-access.name}/servicePerimeters/%s"
  title          = "%s"
  perimeter_type = "PERIMETER_TYPE_REGULAR"
  status {
    restricted_services = ["bigquery.googleapis.com", "storage.googleapis.com"]
		access_levels       = [google_access_context_manager_access_level.access-level.name]

		vpc_accessible_services {
			enable_restriction = true
			allowed_services   = ["bigquery.googleapis.com", "storage.googleapis.com"]
		}

		ingress_policies {
			ingress_from {
				sources {
					access_level = google_access_context_manager_access_level.test-access.name
				}
				identity_type = "ANY_IDENTITY"
			}

			ingress_to {
				resources = [ "*" ]
				operations {
					service_name = "bigquery.googleapis.com"

					method_selectors {
						method = "BigQueryStorage.ReadRows"
					}

					method_selectors {
						method = "TableService.ListTables"
					}

					method_selectors {
						permission = "bigquery.jobs.get"
					}
				}

				operations {
					service_name = "storage.googleapis.com"

					method_selectors {
						method = "google.storage.objects.create"
					}
				}
			}
		}

		egress_policies {
			egress_from {
				identity_type = "ANY_USER_ACCOUNT"
			}
		}
  }
}
```
