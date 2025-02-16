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

#!/bin/bash
set -x
set -o nounset

echo "KIND: $KIND, the config connector kind" # e.g. CloudbuildWorkerPool
echo "PROTO: $PROTO, the gcp proto name" # e.g. WorkerPool
echo "PACKAGE: $PACKAGE, the gcp proto package or service name" # e.g. google.devtools.cloudbuild.v1
echo "SERVICE: $SERVICE, the config connector api group name" # e.g. cloudbuild
 

INS_PATH="dev/tools/controllerbuilder/cmd/codebot/examples/reconciler-granular/2-mapper-case"

while true; do
    "${INS_PATH}"/verify.sh
    if [[ $? -eq 0 ]]; then
        echo passed
        exit 0
    fi 

    # Step 1
    "${INS_PATH}"/1-add-mapper.sh

    "${INS_PATH}"/3-llm-fuzz.sh

done