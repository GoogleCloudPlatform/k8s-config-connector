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


sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".node_count")\n        f.Unimplemented_NotYetTriaged(".tags")\n        return f/' pkg/controller/direct/sql/sqlinstance_fuzzer.go

sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".private_config.custom_host_config")\n        f.Unimplemented_NotYetTriaged(".private_config.custom_host_config.html")\n        return f/' pkg/controller/direct/securesourcemanager/instance_fuzzer.go

sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".natural_language_query_understanding_config")\n        return f/' pkg/controller/direct/discoveryengine/fuzzers.go

sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".connection_endpoints[].service_class_id")\n        return f/' pkg/controller/direct/compute/networkattachment_fuzzer.go

sed -i 's/return f/        f.Unimplemented_NotYetTriaged(".environment_config.execution_config.resource_manager_tags")\n        return f/' pkg/controller/direct/dataproc/batch_fuzzer.go


