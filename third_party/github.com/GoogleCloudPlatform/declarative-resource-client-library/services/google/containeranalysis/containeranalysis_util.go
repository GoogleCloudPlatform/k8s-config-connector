// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package containeranalysis contains utilities for working with GCP's container analysis resources.
package containeranalysis

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// analyzeNoteDiff checks whether more than one of the kind oneof in Note have been changed.
// If they have, it adds a "requires recreate" diff to the list of diffs, so that the resource
// will be destroyed and recreated.  This is not implemented directly in resource_internal because
// this is pretty unusual.
func analyzeNoteDiff(desired, actual *Note, diffs []*dcl.FieldDiff) []*dcl.FieldDiff {
	// the following kinds of field cannot be modified in the same resource.
	// If we find more than one being modified, we will need to add a "forceRecreate"
	// to the list.
	maxOneOf := []string{"Vulnerability", "Build", "BaseImage", "Package", "Deployable", "Discovery", "AttestationAuthority"}
	count := 0
	for _, diff := range diffs {
		if dcl.StringSliceContains(diff.FieldName, maxOneOf) {
			count++
		}
	}
	if count > 1 {
		diffs = append(diffs, &dcl.FieldDiff{FieldName: "Kind", ResultingOperation: []string{"Recreate"}})
	}
	return diffs
}
