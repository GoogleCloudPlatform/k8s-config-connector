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
	"fmt"

	iamapi "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// All returns GroupVersionKinds corresponding to all the GCP resources
// supported by KCC.
func All(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader) ([]schema.GroupVersionKind, error) {
	return resourcesWithDirect(smLoader, serviceMetaLoader, true)
}

func AllWithoutDirect(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader) []schema.GroupVersionKind {
	return resources(smLoader, serviceMetaLoader, true)
}

// ManualResources returns GroupVersionKinds for all the manually configured KCC
// resources.
func ManualResources(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader) ([]schema.GroupVersionKind, error) {
	return resourcesWithDirect(smLoader, serviceMetaLoader, false)
}

func resourcesWithDirect(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader, includesAutoGen bool) ([]schema.GroupVersionKind, error) {
	gvks := resources(smLoader, serviceMetaLoader, includesAutoGen)

	directGVKs, err := DirectResources()
	if err != nil {
		return nil, fmt.Errorf("error getting direct resource GVKs: %w", err)
	}
	for _, gvk := range gvks {
		if _, ok := directGVKs[gvk]; ok {
			delete(directGVKs, gvk)
		}
	}
	for gvk, _ := range directGVKs {
		gvks = append(gvks, gvk)
	}
	return gvks, nil
}

func resources(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader, includesAutoGen bool) []schema.GroupVersionKind {
	gvks := dynamicTypes(smLoader, serviceMetaLoader, includesAutoGen)
	gvks = append(gvks, BasedOnHandwrittenIAMTypes()...)
	return gvks
}

func DirectResources() (map[schema.GroupVersionKind]bool, error) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		return nil, fmt.Errorf("error loading crds: %w", err)
	}
	handWrittenIAMTypes := make(map[schema.GroupVersionKind]bool)
	directResources := make(map[schema.GroupVersionKind]bool)
	for _, gvk := range BasedOnHandwrittenIAMTypes() {
		handWrittenIAMTypes[gvk] = true
	}
	for _, crd := range crds {
		if crd.ObjectMeta.Labels["cnrm.cloud.google.com/tf2crd"] == "true" {
			continue
		}
		if crd.ObjectMeta.Labels["cnrm.cloud.google.com/dcl2crd"] == "true" {
			continue
		}
		versions := crd.Spec.Versions
		highestVersion := k8s.KCCAPIVersionV1Alpha1
		for _, version := range versions {
			if version.Name == k8s.KCCAPIVersionV1Beta1 {
				highestVersion = k8s.KCCAPIVersionV1Beta1
			}
		}
		gvk := schema.GroupVersionKind{
			Group:   crd.Spec.Group,
			Kind:    crd.Spec.Names.Kind,
			Version: highestVersion,
		}
		if _, ok := handWrittenIAMTypes[gvk]; ok {
			continue
		}
		directResources[gvk] = true
	}
	return directResources, nil
}

// AllDynamicTypes returns GroupVersionKinds generated from:
// 1) Terraform schemas (with ServiceMappings metadata layer)
// 2) DCL OpenAPI schemas
func AllDynamicTypes(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader) []schema.GroupVersionKind {
	return dynamicTypes(smLoader, serviceMetaLoader, true)
}

func dynamicTypes(smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader metadata.ServiceMetadataLoader, includesAutoGen bool) []schema.GroupVersionKind {
	gvks := basedOnServiceMappings(smLoader, includesAutoGen)
	gvks = append(gvks, BasedOnDCL(serviceMetaLoader)...)
	return gvks
}

func BasedOnAllServiceMappings(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	return basedOnServiceMappings(smLoader, true)
}

func BasedOnManualServiceMappings(smLoader *servicemappingloader.ServiceMappingLoader) []schema.GroupVersionKind {
	return basedOnServiceMappings(smLoader, false)
}

func basedOnServiceMappings(smLoader *servicemappingloader.ServiceMappingLoader, includesAutoGen bool) []schema.GroupVersionKind {
	gvkSet := make(map[schema.GroupVersionKind]bool)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			if !includesAutoGen && rc.AutoGenerated {
				continue
			}
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
