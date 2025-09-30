#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# This script runs a query in 2-day intervals for the last 30 days.

# Loop 15 times (15 * 2 days = 30 days)
for i in {0..14}
do
  # Calculate the start and end days ago for each interval
  DAYS_AGO_END=$((i * 2))
  DAYS_AGO_START=$(((i * 2) + 2))

  # Get timestamps in ISO 8601 format (UTC), at the start of the day
  START_TIME=$(date -u -d "${DAYS_AGO_START} days ago" +"%Y-%m-%dT00:00:00Z")
  END_TIME=$(date -u -d "${DAYS_AGO_END} days ago" +"%Y-%m-%dT00:00:00Z")

  echo "--- Querying from ${START_TIME} to ${END_TIME} ---"

  # Define the query.
  # FIX: Corrected the FROM clause to use explicit project.dataset.table syntax.
  read -r -d '' QUERY <<EOF
SELECT
  JSON_VALUE(labels, '$.instance_id') AS instance_id,
  JSON_VALUE(labels, '$.location') AS region,
  JSON_VALUE(labels, '$.consumer_project_number') AS consumer_project_number
FROM
  \`apricot-api-prod\`.\`global._Default\`._Default
WHERE
  log_name = "projects/apricot-api-prod/logs/memorystore.googleapis.com%2Fagents"
  AND CONTAINS_SUBSTR(json_payload, "[/] Parsing Redis ELF/Loader info")
  AND timestamp >= "${START_TIME}"
  AND timestamp < "${END_TIME}"
GROUP BY
  1, 2, 3
EOF

  # Run the query using the bq command-line tool
  bq query --use_legacy_sql=false "${QUERY}"

done
