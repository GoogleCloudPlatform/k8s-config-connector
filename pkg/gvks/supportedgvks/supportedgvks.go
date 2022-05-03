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

package supportedgvks

import (
	iamapi "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// All returns GroupVersionKinds corresponding to all the GCP resources
// supported by KCC.
func All(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader) []schema.GroupVersionKind {
	gvks := DynamicTypes(smLoader, serviceMetaLoader)
	gvks = append(gvks, BasedOnHandwrittenIAMTypes()...)
	return gvks
}

// DynamicTypes returns GroupVersionKinds generated from:
// 1) Terraform schemas (with ServiceMappings metadata layer)
// 2) DCL OpenAPI schemas
func DynamicTypes(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader) []schema.GroupVersionKind {
	gvks := BasedOnServiceMappings(smLoader)
	gvks = append(gvks, BasedOnDCL(serviceMetaLoader)...)
	return gvks
}

func BasedOnServiceMappings(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	gvkSet := make(map[schema.GroupVersionKind]bool)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			gvk := krmtotf.GVKForResource(&sm, &rc)
			gvkSet[gvk] = true
		}
	}
	return k8s.GVKSetToList(gvkSet)
}

func BasedOnHandwrittenIAMTypes() []schema.GroupVersionKind {
	return []schema.GroupVersionKind{
		iamapi.IAMAuditConfigGVK,
		iamapi.IAMPolicyGVK,
		iamapi.IAMPolicyMemberGVK,
		iamapi.IAMPartialPolicyGVK,
	}
}

func BasedOnDCL(serviceMetaLoader metadata.ServiceMetadataLoader) []schema.GroupVersionKind {
	serviceList := serviceMetaLoader.GetAllServiceMetadata()
	gvkList := make([]schema.GroupVersionKind, 0)
	for _, service := range serviceList {
		for _, resource := range service.Resources {
			if resource.Releasable {
				gvk := metadata.GVKForResource(service, resource)
				gvkList = append(gvkList, gvk)
			}
		}
	}
	return gvkList
}
