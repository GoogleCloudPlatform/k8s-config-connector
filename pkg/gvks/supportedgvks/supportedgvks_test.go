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

package supportedgvks_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestAllIncludesIAMResource(t *testing.T) {
	allResources, err := supportedgvks.All(testservicemappingloader.New(t), dclmetadata.New())
	if err != nil {
		t.Fatalf("error loading all supported GVKs: %v", err)
	}
	iamResources := []schema.GroupVersionKind{
		v1beta1.IAMAuditConfigGVK,
		v1beta1.IAMPolicyGVK,
		v1beta1.IAMPolicyMemberGVK,
	}
	for _, iamResource := range iamResources {
		assertIncludesResource(t, allResources, iamResource)
	}
}

func assertIncludesResource(t *testing.T, resources []schema.GroupVersionKind, resource schema.GroupVersionKind) {
	for _, r := range resources {
		if r == resource {
			return
		}
	}
	t.Errorf("expected list of resources to contain resource '%v', but it did not", resource)
}

func TestNoResourceIsDoubleDeclared(t *testing.T) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("error creating new service mapping loader: %v", err)
	}
	serviceMetadataLoader := dclmetadata.New()
	tfBasedResources := supportedgvks.BasedOnAllServiceMappings(smLoader)
	dclBasedResources := supportedgvks.BasedOnDCL(serviceMetadataLoader)
	var u []string
	resourceMap := make(map[string]bool)
	for _, gvk := range tfBasedResources {
		resourceMap[gvk.Kind] = true
	}
	for _, gvk := range dclBasedResources {
		if _, ok := resourceMap[gvk.Kind]; ok {
			u = append(u, gvk.Kind)
		}
	}
	if len(u) != 0 {
		t.Fatalf("resources %v have been declared in both ServiceMapping(TF/KCC bridge) and ServiceMetadata(DCL/KCC bridge)", u)
	}
}
