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

package leasable

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nasa9084/go-openapi"
)

func ResourceConfigSupportsLeasing(rc *v1alpha1.ResourceConfig, tfResourceMap map[string]*tfschema.Resource) (ok bool, err error) {
	// Disable leasing for DataflowJob since the leasing mechanism needs to
	// modify the resource's labels regularly, which, for DataflowJob, would
	// trigger a job update on each modification.
	if rc.Kind == "DataflowJob" {
		return false, nil
	}
	labelsField := rc.MetadataMapping.Labels
	if labelsField == "" {
		return false, nil
	}
	tfResource, ok := tfResourceMap[rc.Name]
	if !ok {
		return false, fmt.Errorf("unknown resource %v", rc.Name)
	}
	labelsFieldSchema, err := tfresource.GetTFSchemaForField(tfResource, labelsField)
	if err != nil {
		return false, fmt.Errorf("error getting schema for field '%v' of resource '%v': %w", labelsField, rc.Name, err)
	}
	labelsFieldIsMutable := !labelsFieldSchema.ForceNew
	return labelsFieldIsMutable, nil
}

func DCLSchemaSupportsLeasing(schema *openapi.Schema) (bool, error) {
	labelsField, s, found, err := extension.GetLabelsFieldSchema(schema)
	if err != nil {
		return false, fmt.Errorf("error getting schema for field %v of resource '%v': %w", labelsField, schema.Title, err)
	}
	if !found {
		return false, nil
	}
	isImmutable, err := extension.IsImmutableField(s)
	if err != nil {
		return false, err
	}
	return !isImmutable, nil
}
