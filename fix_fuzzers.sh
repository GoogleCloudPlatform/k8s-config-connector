# Copyright 2026 Google LLC
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


# bigqueryreservation assignment
sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".principal")\n        return f/' pkg/controller/direct/bigqueryreservation/assignment_fuzzer.go

# redis cluster
sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".rotate_server_certificate")\n        f.Unimplemented_NotYetTriaged(".server_ca_mode")\n        f.Unimplemented_NotYetTriaged(".server_ca_pool")\n        return f/' pkg/controller/direct/redis/cluster/cluster_fuzzer.go

# sql instance
sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".nodes")\n        f.Unimplemented_NotYetTriaged(".switch_transaction_logs_to_cloud_storage_enabled")\n        f.Unimplemented_NotYetTriaged(".satisfies_pzi")\n        f.Unimplemented_NotYetTriaged(".include_replicas_for_major_version_upgrade")\n        return f/' pkg/controller/direct/sql/sqlinstance_fuzzer.go

# compute future reservation
sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".params")\n        return f/' pkg/controller/direct/compute/futurereservation_fuzzer.go

# dataflow flex template job
sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".additional_pipeline_options")\n        return f/' pkg/controller/direct/dataflow/flextemplatejob_fuzzer.go

# alloydb cluster
sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".dataplex_config")\n        f.Unimplemented_NotYetTriaged(".dataplex_config.enabled")\n        return f/' pkg/controller/direct/alloydb/cluster_fuzzer.go


