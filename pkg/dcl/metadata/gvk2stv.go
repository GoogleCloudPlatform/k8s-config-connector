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

package metadata

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ToServiceTypeVersion returns the DCL ServiceTypeVersion given the KRM GroupVersionKind.
func ToServiceTypeVersion(gvk k8sschema.GroupVersionKind, loader ServiceMetadataLoader) (dclunstruct.ServiceTypeVersion, error) {
	service := groupToService(gvk.Group)
	sm, found := loader.GetServiceMetadata(service)
	if !found {
		return dclunstruct.ServiceTypeVersion{}, fmt.Errorf("ServiceMetadata for service %v is not found", service)
	}
	r, found := sm.GetResourceWithKind(gvk.Kind)
	if !found {
		return dclunstruct.ServiceTypeVersion{}, fmt.Errorf("resource with kind %v not supported in service %v", gvk.Kind, service)
	}
	return dclunstruct.ServiceTypeVersion{
		Service: sm.ServiceNameUsedByDCL,
		Type:    r.DCLType,
		Version: r.DCLVersion,
	}, nil
}

// ToGroupVersionKind returns the KRM GroupVersionKind given the DCL ServiceTypeVersion.
func ToGroupVersionKind(stv dclunstruct.ServiceTypeVersion, loader ServiceMetadataLoader) (k8sschema.GroupVersionKind, error) {
	sm, found := loader.GetServiceMetadata(stv.Service)
	if !found {
		return k8sschema.GroupVersionKind{}, fmt.Errorf("ServiceMetadata for service %v is not found", stv.Service)
	}
	r, found := sm.GetResourceWithType(stv.Type)
	if !found {
		return k8sschema.GroupVersionKind{}, fmt.Errorf("resource with DCL type %v not supported in service %v", stv.Type, stv.Service)
	}
	return GVKForResource(sm, r), nil
}

func IsDCLBasedResourceKind(gvk k8sschema.GroupVersionKind, loader ServiceMetadataLoader) bool {
	r, found := loader.GetResourceWithGVK(gvk)
	if !found {
		return false
	}
	return r.Releasable
}

func groupToService(group string) string {
	return strings.TrimSuffix(group, k8s.APIDomainSuffix)
}
