// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package refs

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// TrimComputeURIPrefix trims known GCP Compute Engine URL and URI prefixes
// to normalize the resource path to projects/{{project}}/... format.
// This is robust and ensures unknown values/prefixes are not silently ignored.
func TrimComputeURIPrefix(ref string) string {
	return refsv1beta1.TrimComputeURIPrefix(ref)
}
