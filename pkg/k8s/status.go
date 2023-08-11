// Copyright 2022 Google LLC
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

package k8s

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var renameStatusFieldsWithReservedNamesExcludeList = map[string]bool{
	// This resource was already released with no 'resource' prefix prepended to `.status.conditions` field in CRD.
	// Added to the list to avoid breaking change.
	"google_storage_default_object_access_control": true,
}

func RenameStatusFieldsWithReservedNames(status *apiextensions.JSONSchemaProps) (*apiextensions.JSONSchemaProps, error) {
	statusCopy := status.DeepCopy()
	for field := range ReservedStatusFieldNames() {
		renamedField := RenameStatusFieldWithReservedName(field)

		// Error out if status has fields that collide with the renames
		// themselves (e.g. "resourceConditions"). We error out here even if
		// there are no fields that need to be renamed in order to defend
		// against the case where the resource does introduce a name collision
		// later (e.g. suppose status has "resourceConditions", but not
		// "conditions", but then adds "conditions" later).
		// TODO(b/213888818): Handle fields that collide with the renames.
		if _, ok := statusCopy.Properties[renamedField]; ok {
			return nil, fmt.Errorf("status schema already has a field named "+
				"'%v' which is the rename meant for fields that collide "+
				"with the reserved name '%v'", renamedField, field)
		}

		// Rename fields that collide with reserved names
		if schema, ok := statusCopy.Properties[field]; ok {
			statusCopy.Properties[renamedField] = schema
			delete(statusCopy.Properties, field)
		}
	}

	return statusCopy, nil
}

func RenameStatusFieldsWithReservedNamesIfResourceNotExcluded(tfResourceName string, status *apiextensions.JSONSchemaProps) (*apiextensions.JSONSchemaProps, error) {
	if shouldSkip(tfResourceName) {
		return status, nil
	}
	return RenameStatusFieldsWithReservedNames(status)
}

func ReservedStatusFieldNames() map[string]bool {
	reservedFieldNames := make(map[string]bool)

	// Status field names that are in use
	crdSchema := crdboilerplate.GetOpenAPIV3SchemaSkeleton()
	for fieldName := range crdSchema.Properties["status"].Properties {
		reservedFieldNames[fieldName] = true
	}

	// Status field names that might be used in the future
	for _, fieldName := range ReservedStatusFieldNamesForFutureUse {
		reservedFieldNames[fieldName] = true
	}
	return reservedFieldNames
}

func RenameStatusFieldWithReservedName(field string) string {
	return "resource" + text.UppercaseInitial(field)
}

func RenameStatusFieldWithReservedNameIfResourceNotExcluded(tfResourceName, field string) string {
	if shouldSkip(tfResourceName) {
		return field
	}
	return RenameStatusFieldWithReservedName(field)
}

// shouldSkip returns true if the tfResourceName is included in renameStatusFieldsWithReservedNamesExcludeList.
func shouldSkip(tfResourceName string) bool {
	if _, found := renameStatusFieldsWithReservedNamesExcludeList[tfResourceName]; found {
		return true
	}
	return false
}
