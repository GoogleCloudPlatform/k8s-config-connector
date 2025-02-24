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
// Package datafusion contains methods and types for handling datafusion GCP resources.
package beta

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// This custom update mask function removes fields which are not allowed in the update mask but are nevertheless mutable.
func (op *updateInstanceUpdateInstanceOperation) UpdateMask() string {
	allowedValues := []string{"EnableOptions", "EnableRbac", "EnableStackdriverLogging", "EnableStackdriverMonitoring"}
	maskDiffs := make([]*dcl.FieldDiff, 0, len(allowedValues))
	for _, fieldDiff := range op.FieldDiffs {
		if dcl.StringSliceContains(fieldDiff.FieldName, allowedValues) {
			maskDiffs = append(maskDiffs, fieldDiff)
		}
	}
	return dcl.UpdateMask(maskDiffs)
}
