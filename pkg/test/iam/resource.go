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

package testiam

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewResourceRef(refResource *unstructured.Unstructured) v1beta1.ResourceReference {
	return v1beta1.ResourceReference{
		Kind:       refResource.GetKind(),
		APIVersion: refResource.GetAPIVersion(),
		Name:       refResource.GetName(),
	}
}

func NewExternalRef(refResource *unstructured.Unstructured, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader) (v1beta1.ResourceReference, error) {
	gvk := refResource.GroupVersionKind()
	sm, err := smLoader.GetServiceMapping(gvk.Group)
	if err != nil {
		return v1beta1.ResourceReference{}, err
	}
	r, err := krmtotf.NewResource(refResource, sm, provider)
	if err != nil {
		return v1beta1.ResourceReference{}, err
	}
	var id string
	// TODO(kcc-eng): As import ID resolution is dependent on the API server,
	//  for resources that require resolution from references, we cannot
	//  call r.GetImportID before those resources exist in etcd. For now,
	//  for SpannerDatabase, manually set its external reference. We should
	//  restructure this such that either the references already exist
	//  before this point, or external IDs are manually supplied.
	if refResource.GetKind() == "SpannerDatabase" {
		project, ok := k8s.GetAnnotation(k8s.ProjectIDAnnotation, refResource)
		if !ok {
			return v1beta1.ResourceReference{}, fmt.Errorf("referenced resource does not have the annotation %v", k8s.ProjectIDAnnotation)
		}
		instance, ok, err := unstructured.NestedString(refResource.Object, "spec", "instanceRef", "name")
		if err != nil || !ok {
			return v1beta1.ResourceReference{}, fmt.Errorf("error getting instance reference for SpannerDatabase")
		}
		id = fmt.Sprintf("projects/%v/instances/%v/databases/%v", project, instance, r.GetName())
	} else {
		id, err = r.GetImportID(nil, smLoader)
		if err != nil {
			return v1beta1.ResourceReference{}, err
		}
	}
	return v1beta1.ResourceReference{
		Kind:       refResource.GetKind(),
		APIVersion: refResource.GetAPIVersion(),
		External:   id,
	}, nil
}

// isDynamicGVK returns true if the GVK is either TF-based or DCL-based.
func isDynamicGVK(gvk schema.GroupVersionKind, smLoader *servicemappingloader.ServiceMappingLoader, serviceMetaLoader dclmetadata.ServiceMetadataLoader) bool {
	dynamicGVKs := supportedgvks.AllDynamicTypes(smLoader, serviceMetaLoader)
	dynamicGVKMap := make(map[schema.GroupVersionKind]bool)
	for _, gvk := range dynamicGVKs {
		dynamicGVKMap[gvk] = true
	}
	_, ok := dynamicGVKMap[gvk]
	return ok
}

func FixtureSupportsIAMAuditConfigs(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, fixture resourcefixture.ResourceFixture) bool {
	t.Helper()
	if !isDynamicGVK(fixture.GVK, smLoader, serviceMetadataLoader) {
		// Only dynamic GVKs support IAM.
		// Direct GVKs don't support IAM yet.
		// Handwritten GVKs are the IAM resources themselves.
		return false
	}
	// IAM support is not implemented for DCL based resources
	if dclmetadata.IsDCLBasedResourceKind(fixture.GVK, serviceMetadataLoader) {
		return false
	}
	project := testgcp.GCPProject{
		ProjectID:     "project-name",
		ProjectNumber: 1234,
	}
	ns := project.ProjectID
	unstruct := test.ToUnstructWithNamespace(t,
		testcontroller.ReplaceTestVars(t, fixture.Create, "testid", project),
		ns)
	rc, err := smLoader.GetResourceConfig(unstruct)
	if err != nil {
		t.Fatalf("error getting service mapping: %v", err)
	}
	if err != nil {
		t.Fatalf("error getting resource config: %v", err)
	}
	return rc.IAMConfig.AuditConfigName != ""
}

func FixtureSupportsIAMPolicy(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader, fixture resourcefixture.ResourceFixture) bool {
	t.Helper()
	if !isDynamicGVK(fixture.GVK, smLoader, serviceMetadataLoader) {
		// Only dynamic GVKs support IAM.
		// Direct GVKs don't support IAM yet.
		// Handwritten GVKs are the IAM resources themselves.
		return false
	}
	if dclmetadata.IsDCLBasedResourceKind(fixture.GVK, serviceMetadataLoader) {
		dclSchema, err := dclschemaloader.GetDCLSchemaForGVK(fixture.GVK, serviceMetadataLoader, dclSchemaLoader)
		if err != nil {
			t.Fatalf("error getting DCLSchema: %v", err)
		}
		supportsIAM, err := extension.HasIam(dclSchema)
		if err != nil {
			t.Fatalf("error checking if DCLSchema supports IAM: %v", err)
		}
		return supportsIAM
	}
	project := testgcp.GCPProject{
		ProjectID:     "project-name",
		ProjectNumber: 1234,
	}
	ns := project.ProjectID
	unstruct := test.ToUnstructWithNamespace(t,
		testcontroller.ReplaceTestVars(t, fixture.Create, "testid", project),
		ns)
	rc, err := smLoader.GetResourceConfig(unstruct)
	if err != nil {
		t.Fatalf("error getting service mapping: %v", err)
	}
	if err != nil {
		t.Fatalf("error getting resource config: %v", err)
	}
	return rc.IAMConfig.PolicyName != ""
}
