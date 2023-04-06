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
resource "google_bigquery_capacity_commitment" "commitment" {
	capacity_commitment_id = "capacity-tf-test%{random_suffix}"

	location   = "us-west2"
	slot_count = 100
	plan       = "FLEX_FLAT_RATE"
	edition    = "ENTERPRISE"
}

resource "time_sleep" "wait_61_seconds" {
	depends_on = [google_bigquery_capacity_commitment.commitment]
    
	# Only needed for CI tests to be able to tear down the commitment once it's expired
    create_duration = "61s"
}
```
