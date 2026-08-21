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

package container

import (
	"fmt"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/nasa9084/go-openapi"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

func GetContainersForGVK(gvk k8sschema.GroupVersionKind, smLoader dclmetadata.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader) ([]corekccv1alpha1.Container, error) {
	r, found := smLoader.GetResourceWithGVK(gvk)
	if !found {
		return nil, fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", gvk)
	}
	if !r.SupportsContainerAnnotations {
		return nil, nil
	}
	stv, err := dclmetadata.ToServiceTypeVersion(gvk, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting DCL ServiceTypeVersion for GroupVersionKind %v: %w", gvk, err)
	}
	dclSchema, err := schemaLoader.GetDCLSchema(stv)
	if err != nil {
		return nil, fmt.Errorf("error getting the DCL Schema for GroupVersionKind %v: %w", gvk, err)
	}
	// If the resource supports container annotations but is missing the
	// 'x-dcl-parent-container' extension, it is possible that DCL may have
	// removed the deprecated extension entirely. Since the resource _should_
	// support container annotations, construct the list of containers based
	// on the resource's hierarchical references instead.
	if _, ok := dclSchema.Extension["x-dcl-parent-container"]; !ok {
		containers, err := getContainersFromHierarchicalReferencesForGVK(gvk, smLoader, schemaLoader)
		if err != nil {
			return nil, err
		}
		if len(containers) == 0 {
			return nil, fmt.Errorf("expected resource to support containers but found none based on hierarchical references")
		}
		return containers, nil
	}
	containers, err := getContainerConfigFromDCLSchema(dclSchema)
	if err != nil {
		return nil, fmt.Errorf("error resolving the container config from DCL schema: %w", err)
	}
	return containers, nil
}

func getContainersFromHierarchicalReferencesForGVK(gvk k8sschema.GroupVersionKind, smLoader dclmetadata.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader) ([]corekccv1alpha1.Container, error) {
	hierarchicalRefs, err := dcl.GetHierarchicalReferencesForGVK(gvk, smLoader, schemaLoader)
	if err != nil {
		return nil, err
	}
	containerTypes := k8s.ContainerTypesFor(hierarchicalRefs)
	return containerConfigFromContainerTypes(containerTypes), nil
}

func getContainerConfigFromDCLSchema(schema *openapi.Schema) ([]corekccv1alpha1.Container, error) {
	container, hasContainer, err := getContainerType(schema)
	if err != nil {
		return nil, err
	}
	if !hasContainer {
		return nil, nil
	}
	containerTypes := make([]corekccv1alpha1.ContainerType, 0)
	switch container {
	case corekccv1alpha1.ContainerTypeProject,
		corekccv1alpha1.ContainerTypeFolder,
		corekccv1alpha1.ContainerTypeOrganization:
		containerTypes = append(containerTypes, corekccv1alpha1.ContainerType(container))
	default:
		return nil, fmt.Errorf("invalid container type %v", container)
	}
	return containerConfigFromContainerTypes(containerTypes), nil
}

func getContainerType(schema *openapi.Schema) (string, bool, error) {
	raw, ok := schema.Extension["x-dcl-parent-container"]
	if !ok {
		return "", false, nil
	}
	// DCL currently doesn't support resources that could have multiple container kinds.
	container, ok := raw.(string)
	if !ok {
		return "", false, fmt.Errorf("wrong type for 'x-dcl-parent-container' extension: %T, expect to have string type", raw)
	}
	return container, true, nil
}

func containerConfigFromContainerTypes(containerTypes []corekccv1alpha1.ContainerType) []corekccv1alpha1.Container {
	containers := make([]corekccv1alpha1.Container, 0)
	for _, c := range containerTypes {
		containers = append(containers, corekccv1alpha1.Container{
			// It is assumed that, for DCL resources, names of container fields
			// are always equivalent to their corresponding container types.
			TFField: string(c),
			Type:    c,
		})
	}
	return containers
}
