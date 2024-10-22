#!/bin/bash

# Copyright 2024 Google LLC
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


set -o errexit
set -o nounset
set -o pipefail

set -x

REPO_ROOT=$(git rev-parse --show-toplevel)
cd ${REPO_ROOT}/mockgcp

cd tools/patch-proto

# This is an example proto patch; the maintenance_update_policy field has now been added upstream,
# so we don't need it, but we are keeping it as a comment here for reference.
#
# # Use our proto patch tool to add the missing maintenanceUpdatePolicy field for alloydb
# go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/alloydb/v1beta/resources.proto --message Cluster <<EOF
#   // MaintenanceUpdatePolicy defines the policy for system updates.
#   message MaintenanceUpdatePolicy {
#     // MaintenanceWindow specifies a preferred day and time for maintenance.
#     message MaintenanceWindow {
#       // Preferred day of the week for maintenance, e.g. MONDAY, TUESDAY, etc.
#       google.type.DayOfWeek day = 1;

#       // Preferred time to start the maintenance operation on the specified day.
#       // Maintenance will start within 1 hour of this time.
#       google.type.TimeOfDay start_time = 2;
#     }

#     // Preferred windows to perform maintenance. Currently limited to 1.
#     repeated MaintenanceWindow maintenance_windows = 1;
#   }

#   // The maintenance update policy determines when to allow or deny updates.
#   MaintenanceUpdatePolicy maintenance_update_policy = 32;
# EOF

# Use our proto patch tool to add the missing pscConfig field for AlloyDB cluster.
go run . --file ${REPO_ROOT}/mockgcp/third_party/googleapis/google/cloud/alloydb/v1beta/resources.proto --message Cluster <<EOF

  // PscConfig contains PSC related configuration at a cluster level.
  message PscConfig {
    // Optional. Create an instance that allows connections from Private Service
    // Connect endpoints to the instance.
    bool psc_enabled = 1 [(google.api.field_behavior) = OPTIONAL];
  }

  // Optional. The configuration for Private Service Connect (PSC) for the cluster.
  PscConfig psc_config = 31 [(google.api.field_behavior) = OPTIONAL];
EOF
