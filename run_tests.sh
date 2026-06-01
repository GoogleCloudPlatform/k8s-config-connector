#!/bin/bash
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

UNIT_TEST_PACKAGES=$(go list ./pkg/... ./cmd/... ./config/tests/...  ./scripts/resource-autogen/... ./scripts/generate-google3-docs/... ./tests/... | grep -v tests/e2e | grep -v pkg/cli/preview)
go test -count=1 $UNIT_TEST_PACKAGES | grep -i "FAIL"
