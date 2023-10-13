#!/bin/bash
# Copyright 2023 Google LLC
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


set -e

cd third_party/googleapis

# We need to rewrite various proto import paths, to avoid conflicts with the version that is in use in KCC itself.
# We also need to be a little careful not to rewrite the google.api.http annotations, which are used by grpc-go.

rm -rf mockgcp/ && mkdir -p mockgcp/api/

mv google/cloud/ mockgcp/
mv google/iam/ mockgcp/
mv google/api/serviceusage/ mockgcp/api/

cd mockgcp

find . -type f -print0 | xargs -0 sed -i -e "s@google/iam/@mockgcp/iam/@g"
find . -type f -print0 | xargs -0 sed -i -e "s@google\.iam@mockgcp.iam@g"

find . -type f -print0 | xargs -0 sed -i -e "s@google/api/serviceusage/@mockgcp/api/serviceusage/@g"
find . -type f -print0 | xargs -0 sed -i -e "s@google\.api\.serviceusage@mockgcp.api.serviceusage@g"

find . -type f -print0 | xargs -0 sed -i -e "s@google/cloud/@mockgcp/cloud/@g"
find . -type f -print0 | xargs -0 sed -i -e "s@google\.cloud@mockgcp.cloud@g"
