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

package testservicemetadataloader

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
)

func FakeServiceMetadata() []metadata.ServiceMetadata {
	return []metadata.ServiceMetadata{
		{
			Name:       "Test1",
			APIVersion: "v1alpha1",
			DCLVersion: "beta",
			Resources: []metadata.Resource{
				{Kind: "Test1Foo", Releasable: true},
				{Kind: "Test1Bar", Releasable: true},
				{Kind: "Test1IDPConfig", DCLType: "IdpConfig", Releasable: true},
				{Kind: "Test1NotYetSupportedKind", Releasable: false},
				{Kind: "Test1NoLabelsExtension", Releasable: true},
			},
		},
		{
			Name:       "Test2",
			APIVersion: "v1alpha1",
			DCLVersion: "beta",
			Resources: []metadata.Resource{
				{Kind: "Test2Baz", Releasable: true},
				{Kind: "Test2NoName", Releasable: true},
			},
		},
		{
			Name:                 "Test3",
			APIVersion:           "v1alpha1",
			DCLVersion:           "beta",
			ServiceNameUsedByDCL: "DCLTest3",
			Resources: []metadata.Resource{
				{Kind: "Test3Qux", Releasable: true},
			},
		},
		{
			Name:       "Test4",
			APIVersion: "v1alpha1",
			DCLVersion: "beta",
			Resources: []metadata.Resource{
				{Kind: "Test4ProjectContainer", SupportsContainerAnnotations: true, Releasable: true},
				{Kind: "Test4OrganizationContainer", SupportsContainerAnnotations: true, Releasable: true},
				{Kind: "Test4NoContainer", Releasable: true},
				// TODO(b/186159460): Remove the SupportsHierarchicalReferences
				// flag below once all resources support hierarchical
				// references
				{Kind: "Test4NoContainerExtensionButSupportsContainers", SupportsContainerAnnotations: true, SupportsHierarchicalReferences: true, Releasable: true},
				{Kind: "Test4NoContainerExtensionAndSupportsContainersButHasNoHierarchicalRefs", SupportsContainerAnnotations: true, Releasable: true},
			},
		},
		{
			Name:       "Test5",
			APIVersion: "v1alpha1",
			DCLVersion: "beta",
			Resources: []metadata.Resource{
				// TODO(b/186159460): Remove the SupportsHierarchicalReferences
				// flags below once all resources support hierarchical
				// references
				{Kind: "Test5NoHierarchicalRef", Releasable: true},
				{Kind: "Test5ProjectRef", Releasable: true, SupportsHierarchicalReferences: true},
				{Kind: "Test5FolderRef", Releasable: true, SupportsHierarchicalReferences: true},
				{Kind: "Test5OrganizationRef", Releasable: true, SupportsHierarchicalReferences: true},
				{Kind: "Test5MultipleRefs", Releasable: true, SupportsHierarchicalReferences: true},
				{Kind: "Test5TwoRefs", Releasable: true, SupportsHierarchicalReferences: true},
			},
		},
		{
			Name:       "Test6",
			APIVersion: "v1alpha1",
			DCLVersion: "beta",
			Resources: []metadata.Resource{
				// TODO(b/186159460): Remove the SupportsHierarchicalReferences
				// flags below once all resources support hierarchical
				// references
				{Kind: "Test6NoContainerOrHierarchicalRef", Releasable: true},
				{Kind: "Test6OnlyContainer", Releasable: true, SupportsContainerAnnotations: true},
				{Kind: "Test6OnlyHierarchicalRef", Releasable: true, SupportsHierarchicalReferences: true},
				{Kind: "Test6BothContainerAndHierarchicalRef", Releasable: true, SupportsContainerAnnotations: true, SupportsHierarchicalReferences: true},
			},
		},
		{
			Name:       "Test7",
			APIVersion: "v1alpha1",
			DCLVersion: "beta",
			Resources: []metadata.Resource{
				{Kind: "Test7AlphaResource", Releasable: true, DCLVersion: "alpha"},
				{Kind: "Test7BetaResource", Releasable: true},
				{Kind: "Test7GaResource", Releasable: true, DCLVersion: "ga"},
			},
		},
	}
}

// FakeServiceMetadataWithHierarchicalResources returns a ServiceMetadata list
// which includes hierarchical resources to allow for the testing of resources
// that reference hierarchical resources (e.g. "Cloudresourcemanager/Project")
func FakeServiceMetadataWithHierarchicalResources() []metadata.ServiceMetadata {
	return append(FakeServiceMetadata(),
		metadata.ServiceMetadata{
			Name:                 "ResourceManager",
			APIVersion:           "v1beta1",
			DCLVersion:           "ga",
			ServiceNameUsedByDCL: "cloudresourcemanager",
			Resources: []metadata.Resource{
				{Kind: "Project", Releasable: false},
				{Kind: "Folder", Releasable: false},
				{Kind: "Organization", Releasable: false},
				{Kind: "BillingAccount", Releasable: false},
			},
		})
}

func NewForUnitTest() metadata.ServiceMetadataLoader {
	return metadata.NewFromServiceList(FakeServiceMetadata())
}
