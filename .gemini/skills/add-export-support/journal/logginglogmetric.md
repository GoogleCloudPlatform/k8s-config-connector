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

# Export Support Journal for LoggingLogMetric

## Observations
- Implementing export for `LoggingLogMetric` was already partially in place (the resource models implemented `AdapterForURL` and `Export`), but it was missing setting the `ResourceID` on the exported Spec (`lm.Spec.ResourceID = direct.LazyPtr(in.Name)`).
- Adding the `lm.Spec.ResourceID` setting ensures that the exported spec specifies the short name in `spec.resourceID` to perfectly preserve the resource's configuration across environments.
- `LoggingLogMetric` has reference-based project binding (`spec.projectRef`), so setting the project-id annotation or calling `export.SetProjectID` is correctly omitted from the export logic.
- Executing the tests locally required setting `KUBEBUILDER_ASSETS` pointing to a local envtest control plane directory when `E2E_KUBE_TARGET=envtest` is used.
- Running the tests with `GOLDEN_OBJECT_CHECKS=1` ensures the export pipeline executes and compares the output with the golden `_exported.yaml` files.
