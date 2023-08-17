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
resource "google_sql_database_instance" "postgresqldb" {
  name             = "my-database"
  database_version = "POSTGRES_12"
  settings {
    tier = "db-custom-2-13312"
  }
  deletion_protection = false
}

resource "google_sql_ssl_cert" "sql_client_cert" {
  common_name = "my-cert"
  instance    = google_sql_database_instance.postgresqldb.name

  depends_on = [google_sql_database_instance.postgresqldb]
}

resource "google_sql_user" "sqldb_user" {
  name     = "my-username"
  instance = google_sql_database_instance.postgresqldb.name
  password = "my-password"


  depends_on = [google_sql_ssl_cert.sql_client_cert]
}

resource "google_database_migration_service_connection_profile" "postgresprofile" {
  location = "us-central1"
  connection_profile_id = "my-profileid"
  display_name = "my-profileid_display"
  labels = { 
    foo = "bar" 
  }
  postgresql {
    host = google_sql_database_instance.postgresqldb.ip_address.0.ip_address
    port = 5432
    username = google_sql_user.sqldb_user.name
    password = google_sql_user.sqldb_user.password
    ssl {
      client_key = google_sql_ssl_cert.sql_client_cert.private_key
      client_certificate = google_sql_ssl_cert.sql_client_cert.cert
      ca_certificate = google_sql_ssl_cert.sql_client_cert.server_ca_cert
    }
    cloud_sql_id = "my-database"
  }
  depends_on = [google_sql_user.sqldb_user]
}
```
