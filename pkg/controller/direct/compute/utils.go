// Copyright 2025 Google LLC
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

package compute

import (
	"reflect"
	"strings"
)

/*
IsSelfLinkEqual Terraform and mock uses the /beta/ endpoints, while direct controller uses /v1/.
Compute resources(i.e. ComputeServiceAttachment) might be managed by legacy controller and still use beta endpoint.

(Might be a bug/intended behavior in Compute service)
When v1 resource references to a beta resource, after creation the version in selfLink of the referenced resource changed from beta to v1.

Compare resource selfLink by eliminating the version to avoid version mismatching.
todo(yuhou): Should direct controller use the same version that TF uses to avoid this mixed version issue in Compute?
*/

func IsSelfLinkEqual(a, b *string) bool {
	if reflect.DeepEqual(a, b) {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	aVal := *a
	bVal := *b

	for _, basePath := range []string{"https://compute.googleapis.com/compute", "https://www.googleapis.com/compute"} {
		for _, version := range []string{"/beta/", "/v1/"} {
			prefix := basePath + version
			if strings.HasPrefix(aVal, prefix) {
				aVal = strings.TrimPrefix(aVal, prefix)
			}
			if strings.HasPrefix(bVal, prefix) {
				bVal = strings.TrimPrefix(bVal, prefix)
			}
		}
	}
	return aVal == bVal
}
