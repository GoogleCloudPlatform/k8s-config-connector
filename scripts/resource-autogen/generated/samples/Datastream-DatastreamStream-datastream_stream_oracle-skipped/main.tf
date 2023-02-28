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
resource "google_datastream_connection_profile" "source" {
    display_name          = "Oracle Source"
    location              = "us-central1"
    connection_profile_id = "source-profile"

    oracle_profile {
        hostname = "hostname"
        port     = 1521
        username = "user"
        password = "pass"
        database_service = "ORCL"
    }
}

resource "google_datastream_connection_profile" "destination" {
    display_name          = "BigQuery Destination"
    location              = "us-central1"
    connection_profile_id = "destination-profile"

    bigquery_profile {}
}

resource "google_datastream_stream" "stream5" {
    display_name = "Oracle to BigQuery"
    location     = "us-central1"
    stream_id    = "my-stream"
    desired_state = "RUNNING"

    source_config {
        source_connection_profile = google_datastream_connection_profile.source.id
        oracle_source_config {
            max_concurrent_cdc_tasks = 8
            max_concurrent_backfill_tasks = 12
            include_objects {
                oracle_schemas {
                    schema = "schema"
                    oracle_tables {
                        table = "table"
                        oracle_columns {
                            column = "column"
                        }
                    }
                }
            }
            exclude_objects {
                oracle_schemas {
                    schema = "schema"
                    oracle_tables {
                        table = "table"
                        oracle_columns {
                            column = "column"
                        }
                    }
                }
            }
            drop_large_objects {}
		    }
    }

    destination_config {
        destination_connection_profile = google_datastream_connection_profile.destination.id
        bigquery_destination_config {
            data_freshness = "900s"
            source_hierarchy_datasets {
                dataset_template {
                    location = "us-central1"
                }
            }
        }
    }

    backfill_all {
        oracle_excluded_objects {
            oracle_schemas {
                schema = "schema"
                oracle_tables {
                    table = "table"
                    oracle_columns {
                        column = "column"
                    }
                }
            }
        }
    }
}
```
