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
resource "google_alloydb_cluster" "source" {
  cluster_id = "alloydb-source-cluster"
  location   = "us-central1"
  network    = data.google_compute_network.default.id

  initial_user {
    password = "alloydb-source-cluster"
  }
}

resource "google_alloydb_instance" "source" {
  cluster       = google_alloydb_cluster.source.name
  instance_id   = "alloydb-instance"
  instance_type = "PRIMARY"

  machine_config {
    cpu_count = 2
  }

  depends_on = [google_service_networking_connection.vpc_connection]
}

resource "google_alloydb_backup" "source" {
  backup_id    = "alloydb-backup"
  location     = "us-central1"
  cluster_name = google_alloydb_cluster.source.name

  depends_on = [google_alloydb_instance.source]
}

resource "google_alloydb_cluster" "restored_from_backup" {
  cluster_id            = "alloydb-backup-restored"
  location              = "us-central1"
  network               = data.google_compute_network.default.id
  restore_backup_source {
    backup_name = google_alloydb_backup.source.name
  }
}

resource "google_alloydb_cluster" "restored_via_pitr" {
  cluster_id             = "alloydb-pitr-restored"
  location               = "us-central1"
  network                = data.google_compute_network.default.id

  restore_continuous_backup_source {
    cluster = google_alloydb_cluster.source.name
    point_in_time = "2023-08-03T19:19:00.094Z"
  }
}

data "google_project" "project" {}

data "google_compute_network" "default" {
  name = "alloydb-network"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          =  "alloydb-source-cluster"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = data.google_compute_network.default.id
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = data.google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}
```
