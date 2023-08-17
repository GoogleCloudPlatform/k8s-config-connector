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
data "google_project" "project" {
}

resource "google_compute_network" "default" {
  name = "vpc-network"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          =  "private-ip-alloc"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}


resource "google_database_migration_service_connection_profile" "alloydbprofile" {
  location = "us-central1"
  connection_profile_id = "my-profileid"
  display_name = "my-profileid_display"
  labels = { 
    foo = "bar" 
  }
  alloydb {
    cluster_id = "tf-test-dbmsalloycluster%{random_suffix}"
    settings {
      initial_user {
        user = "alloyuser%{random_suffix}"
        password = "alloypass%{random_suffix}"
      }
      vpc_network = google_compute_network.default.id
      labels  = { 
        alloyfoo = "alloybar" 
      }
      primary_instance_settings {
        id = "priminstid"
        machine_config {
          cpu_count = 2
        }
        database_flags = { 
        }
        labels = { 
          alloysinstfoo = "allowinstbar" 
        }
      }
    }
  }

  depends_on = [google_service_networking_connection.vpc_connection]
}
```
