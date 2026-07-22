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

# Journal for ComputeNodeTemplate Direct Migration Validation

## 2026-07-21 - ComputeNodeTemplate Direct Takeover Normalization
- **cpuOvercommitType (Unspecified Optional Field):**
  - GCP API defaults `cpuOvercommitType` to `"NONE"` when not explicitly provided.
  - The KRM spec (`create.yaml`) for `ComputeNodeTemplate` does not define `cpuOvercommitType`.
  - To prevent the direct controller from seeing this as a diff (and attempting to clear/update the field), we normalized `maskedActual.CpuOvercommitType = nil` when `clonedDesired.CpuOvercommitType == nil` inside `compareComputeNodeTemplate`.

- **region (Fully-Qualified URL/URI Normalization):**
  - GCP API returns `region` as a fully-qualified URI: `https://www.googleapis.com/compute/v1/projects/${projectId}/regions/us-central1`.
  - In KRM, the user specifies the region name directly, e.g., `us-central1`.
  - We normalized the `Region` field on both `maskedActual` and `clonedDesired` to their last path component using the package-level helper `lastComponent` inside `compareComputeNodeTemplate`.
