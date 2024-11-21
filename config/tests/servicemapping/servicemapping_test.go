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

package servicemapping_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"
	autogenloader "github.com/GoogleCloudPlatform/k8s-config-connector/scripts/resource-autogen/servicemapping/servicemappingloader"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	emptyTypeConfig = v1alpha1.TypeConfig{}
	emptyIAMConfig  = v1alpha1.IAMConfig{}
)

func TestIDTemplateCanBeUsedToMatchResourceNameShouldHaveValue(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		for _, rc := range sm.Spec.Resources {
			if rc.Direct {
				continue
			}
			if rc.IDTemplateCanBeUsedToMatchResourceName == nil {
				t.Fatalf("resource config '%v' is missing required field 'IDTemplateCanBeUsedToMatchResourceName'",
					rc.Name)
			}
		}
	}
}

func TestNamingConventions(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		switch sm.Spec.Name {
		case "CloudBuild", "CloudIdentity", "CloudIOT", "CloudTasks", "CloudAsset", "CloudIDS", "CloudFunctions2":
			// CloudBuildTrigger was added before this test was put in so for historical
			// reasons we ignore CloudBuild service mappings (until we make a fix).
			// CloudIdentity is a resource that we decided should keep the "Cloud"
			// prefix to prevent confusion from a service just being named "Identity".
			// CloudIOT, CloudTasks, CloudAsset, CloudIDS and CloudFunctions2
			// are service names coming from TF types via auto-generation. KCC
			// doesn't manually remove 'Cloud' from service names during
			// auto-generation.
			continue
		}
		if strings.HasPrefix(sm.Spec.Name, "Cloud") {
			t.Fatalf("invalid service mapping name '%v': 'Cloud' should be dropped from any service name of which it is not an integral part", sm.Spec.Name)
		}
		for _, rc := range sm.Spec.Resources {
			if strings.HasPrefix(rc.Kind, "Cloud") {
				t.Fatalf("invalid resource kind '%v': 'Cloud' should be dropped from the service portion of any resource name", rc.Kind)
			}
		}
	}
}

func TestServiceHostName(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		hostName := sm.Spec.ServiceHostName
		if hostName == "" {
			t.Fatalf("unexpected empty value for ServiceHostName for service mapping '%v'", sm.Name)
		}
		if !strings.HasSuffix(hostName, "googleapis.com") {
			t.Fatalf("unexpected empty value for ServiceHostName for service mapping '%v': expected suffix of 'googleapis.com'", sm.Name)
		}
	}
}

func TestIAMPolicyMappings(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		for _, rc := range sm.Spec.Resources {
			rc := rc
			// TODO (b/221463073): disable ComputeBackendService until
			// ComputeRegionBackendService gets IAM support.
			if rc.Kind == "ComputeBackendService" {
				continue
			}
			if rc.Direct { // Do not check for direct resource
				continue
			}
			t.Run(rc.Kind, func(t *testing.T) {
				t.Parallel()
				// IAMConfig is not supported for the auto-generated v1alpha1 CRDs.
				if isAutogenAlphaResource(&sm, &rc) {
					return
				}
				testIamPolicyMappings(t, rc)
			})
		}
	}
}

func TestIAMPolicyMappingsForKindsWithMultipleResourceConfigs(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			kindToRCs := make(map[string][]v1alpha1.ResourceConfig)
			for _, rc := range sm.Spec.Resources {
				kindToRCs[rc.Kind] = append(kindToRCs[rc.Kind], rc)
			}
			for kind, rcs := range kindToRCs {
				if len(rcs) < 2 {
					continue
				}
				kind := kind
				rcs := rcs
				t.Run(kind, func(t *testing.T) {
					t.Parallel()
					assertAllHaveEmptyOrNonEmptyIAMConfigButNotBoth(t, kind, rcs)
					assertAllHaveSameValueForSupportsConditions(t, kind, rcs)
					assertAllOrNoneSupportAuditConfigs(t, kind, rcs)
				})
			}
		})
	}
}

func TestKindsWithMultipleResourceConfigsHaveSameDescriptionsForSameReferences(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			kindToRCs := make(map[string][]v1alpha1.ResourceConfig)
			for _, rc := range sm.Spec.Resources {
				kindToRCs[rc.Kind] = append(kindToRCs[rc.Kind], rc)
			}
			for kind, rcs := range kindToRCs {
				if len(rcs) < 2 {
					continue
				}
				kind := kind
				rcs := rcs
				t.Run(kind, func(t *testing.T) {
					t.Parallel()
					assertAllHaveSameDescriptionsForSameReferences(t, kind, rcs)
				})
			}
		})
	}
}

func assertAllHaveSameDescriptionsForSameReferences(t *testing.T, kind string, rcs []v1alpha1.ResourceConfig) {
	t.Helper()
	tfFieldToDescription := make(map[string]string)
	for _, rc := range rcs {
		for _, ref := range rc.ResourceReferences {
			if _, ok := tfFieldToDescription[ref.TFField]; !ok {
				tfFieldToDescription[ref.TFField] = ref.Description
				continue
			}
			description := tfFieldToDescription[ref.TFField]
			if ref.Description != description {
				t.Errorf("all ResourceConfigs of kind %v must have the same descriptions "+
					"for all resource references with the same tfField, but not "+
					"all resource references with tfField %v have the same descriptions", kind, ref.TFField)
			}
		}
	}
}

func TestResourcesListedAlphabetically(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			var prev string
			for _, curr := range sm.Spec.Resources {
				if prev > curr.Name {
					t.Errorf("resources not listed alphabetically: %v listed before %v", prev, curr.Name)
				}
				prev = curr.Name
			}
		})
	}
}

func TestTerraformFieldsAreInResourceSchema(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				tfResource := provider.ResourcesMap[rc.Name]
				// Add all the fields that are considered reference fields
				// to this array
				fields := []string{
					rc.MetadataMapping.Name,
					rc.MetadataMapping.Labels,
					rc.ServerGeneratedIDField,
				}
				for _, refConfig := range rc.ResourceReferences {
					fields = append(fields, refConfig.TFField)
				}
				for _, d := range rc.Directives {
					fields = append(fields, d)
				}
				for _, f := range rc.IgnoredFields {
					fields = append(fields, f)
				}
				if rc.IgnoredOutputOnlySpecFields != nil {
					for _, o := range *rc.IgnoredOutputOnlySpecFields {
						fields = append(fields, o)
					}
				}
				for _, c := range rc.Containers {
					fields = append(fields, c.TFField)
				}
				if rc.ObservedFields != nil {
					for _, o := range *rc.ObservedFields {
						fields = append(fields, o)
					}
				}
				// Check the fields to ensure they're in the schema
				for _, f := range fields {
					if f == "" {
						continue
					}
					if !tfresource.TFResourceHasField(tfResource, f) {
						// TODO(b/278948939): Remove once the unknown fields are cleaned up in google_apigee_addons_config.
						if rc.Name == "google_apigee_addons_config" {
							t.Logf("field '%v' mentioned in ServiceMapping for the auto-generated v1alpha1 resource '%v' but is not found in resource schema", f, rc.Name)
						} else {
							t.Errorf("field '%v' mentioned in ServiceMapping for '%v' but is not found in resource schema", f, rc.Name)
						}
					}
				}
			}
		})
	}
}

func TestReferencedTargetFieldsAreInReferencedResourceSchema(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	kindToTFResources := createKindToTFResourcesMap(serviceMappings)
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					testReferencedTargetFieldsAreInReferencedResourceSchema(t, rc, provider, kindToTFResources)
				})
			}
		})
	}
	fmt.Println(kindToTFResources)
}

func testReferencedTargetFieldsAreInReferencedResourceSchema(t *testing.T, rc v1alpha1.ResourceConfig, provider *schema.Provider, kindToTFResources map[string][]string) {
	t.Helper()
	for _, ref := range rc.ResourceReferences {
		for _, tc := range typeConfigsOf(ref) {
			if tc.TargetField == "" {
				// If no TargetField is specified, then either this is a
				// complex reference or the TargetField is the referenced
				// resource's metadata.name (in which case there is no need to
				// check the referenced resource's Terraform schema)
				continue
			}
			if tc.GVK.Kind == "" {
				t.Errorf("kind %v has a resource reference with a targetField specified as %v but has no kind specified", rc.Kind, tc.TargetField)
				continue
			}
			if rc.Name == "google_dns_record_set" && tc.GVK.Kind == "ComputeForwardingRule" && tc.TargetField == "location" {
				// https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/1399#issuecomment-2014218170
				continue
			}
			for _, referencedTFResourceName := range kindToTFResources[tc.GVK.Kind] {
				referencedTFResource := provider.ResourcesMap[referencedTFResourceName]
				if !tfresource.TFResourceHasField(referencedTFResource, tc.TargetField) {
					t.Errorf("kind %v has a resource reference with kind %v and targetField %v, "+
						"but this field does not exist in the Terraform resource %v",
						rc.Kind, tc.GVK.Kind, tc.TargetField, referencedTFResourceName)
				}
			}
		}
	}
}

func TestResourceReferencesAreValid(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		for _, rc := range sm.Spec.Resources {
			rc := rc
			t.Run(rc.Kind, func(t *testing.T) {
				t.Parallel()
				if isAutogenAlphaResource(&sm, &rc) {
					return
				}
				validateResourceReferences(t, rc)
			})
		}
	}
}

func validateResourceReferences(t *testing.T, rc v1alpha1.ResourceConfig) {
	if len(rc.ResourceReferences) == 0 {
		return
	}
	assertHasAtMostOneReferenceConfigPerField(t, rc)
	for _, refConfig := range rc.ResourceReferences {
		if len(refConfig.Types) == 0 {
			assertTypeConfig(t, rc, refConfig, refConfig.TypeConfig)
		} else {
			if !reflect.DeepEqual(refConfig.TypeConfig, emptyTypeConfig) {
				t.Errorf("should not fill the inline TypeConfig if Types is specified")
			}
			for _, typeConfig := range refConfig.Types {
				assertTypeConfig(t, rc, refConfig, typeConfig)
			}
			for _, typeConfig := range refConfig.Types {
				if typeConfig.Key == "" {
					t.Errorf("the ReferenceConfig for tfField %v has multiple types, but not all types have a key specified, like: %+v", refConfig.TFField, typeConfig)
				}
			}
		}
	}
}

func assertHasAtMostOneReferenceConfigPerField(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	tfFields := make(map[string]bool)
	for _, refConfig := range rc.ResourceReferences {
		tfField := refConfig.TFField
		if tfField == "" {
			t.Errorf("tfField value doesn't exist for the reference config")
		}
		if _, ok := tfFields[tfField]; ok {
			t.Errorf("tfField %v has more than one reference config", tfField)
		}
		tfFields[tfField] = true
	}
}

func assertTypeConfig(t *testing.T, rc v1alpha1.ResourceConfig, ref v1alpha1.ReferenceConfig, tc v1alpha1.TypeConfig) {
	gvkUnspecified := tc.GVK.Group == "" && tc.GVK.Version == "" && tc.GVK.Kind == ""
	if gvkUnspecified && tc.JSONSchemaType == "" {
		t.Errorf("the TypeConfig for tfField %v doesn't have either a GVK or a JSONSchemaType", ref.TFField)
	}
	if !gvkUnspecified && tc.JSONSchemaType != "" {
		t.Errorf("the TypeConfig for tfField %v has both GVK and JSONSchemaType defined; they should be mutually exclusive", ref.TFField)
	}
	if !gvkUnspecified {
		validateTypeConfigGVK(t, rc, ref, tc)
	}
}

func validateTypeConfigGVK(t *testing.T, rc v1alpha1.ResourceConfig, ref v1alpha1.ReferenceConfig, tc v1alpha1.TypeConfig) {
	gvk := tc.GVK
	if gvk.Kind == "" {
		t.Fatalf("invalid resource reference '%v' on resource '%v' with key '%v': the field 'kind' must have a value", ref.TFField, rc.Kind, tc.Key)
	}
	if gvk.Group == "" {
		t.Fatalf("invalid resource reference '%v' on resource '%v' with key '%v': the field 'group' must have a value", ref.TFField, rc.Kind, tc.Key)
	}
	if gvk.Version == "" {
		t.Fatalf("invalid resource reference '%v' on resource '%v' with key '%v': the field 'version' must have a value", ref.TFField, rc.Kind, tc.Key)
	}
	// this is needed because there is a resource reference to a Kind that doesn't exist yet (BillingAccount)
	// when a billing service mapping is added, delete to the "end code block" comment
	billingGroup := "billing.cnrm.cloud.google.com"
	if gvk.Group == billingGroup {
		_, err := testservicemappingloader.New(t).GetServiceMapping(billingGroup)
		if err == nil {
			t.Fatalf("a service mapping for billing has been added -- delete this code block (see comment above)")
		}
		return
	}
	// end code block

	// This is needed because there is a resource reference to a Kind that doesn't exist yet (Organization)
	// when an organization resource is added, delete to the "end code block" comment
	resourceManagerGroup := "resourcemanager.cnrm.cloud.google.com"
	if gvk.Group == resourceManagerGroup {
		sm, err := testservicemappingloader.New(t).GetServiceMapping(resourceManagerGroup)
		if err != nil {
			t.Fatalf("expected resource manager service mapping but there was none")
		}
		for _, r := range sm.Spec.Resources {
			if r.Kind == "Organization" {
				t.Fatalf("a resource for organizations has been added -- delete this code block (see comment above)")
			}
		}
		return
	}
	// end code block

	// This list of ignored GVK is to allow certain resources to have
	// external-only resource references (DCL-based resources or unsupported
	// resources).
	ignoredGVKList := []k8sschema.GroupVersionKind{
		{
			Group:   "networksecurity.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "NetworkSecurityClientTLSPolicy",
		},
		{
			Group:   "cloudbuild.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "CloudBuildGithubEnterpriseConfig",
		},
		{
			Group:   "cloudbuild.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "CloudBuildBitbucketServerConfig",
		},
		{
			Group:   "cloudbuild.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "CloudBuildV2Repository",
		},
		{
			Group:   "certificatemanager.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "CertificateManagerCertificateIssuanceConfig",
		},
	}
	for _, g := range ignoredGVKList {
		if gvk == g {
			return
		}
	}

	crd, err := crdloader.GetCRD(gvk.Group, gvk.Version, gvk.Kind)
	if err != nil {
		t.Fatalf("bad resource reference '%v' on resource '%v': error getting crd: %v", ref.TFField, rc.Kind, err)
	}
	crdGvk := k8sschema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: k8s.GetVersionFromCRD(crd),
		Kind:    crd.Spec.Names.Kind,
	}
	if gvk != crdGvk {
		t.Fatalf("crd and service mappings reference mismatch for reference '%v' on resource '%v' with key '%v': service mappings '%v', crd '%v'",
			ref.TFField, rc.Kind, tc.Key, gvk, crdGvk)
	}
}

func TestHierarchicalReferences(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					testHierarchicalReferences(t, rc)
				})
			}
		})
	}
}

func testHierarchicalReferences(t *testing.T, rc v1alpha1.ResourceConfig) {
	// TODO(b/193177782): Delete this if-block once all resources support
	// hierarchical references.
	if !krmtotf.SupportsHierarchicalReferences(&rc) {
		return
	}
	assertHasAtMostOneOfEachHierarchicalReferenceType(t, rc)
	for _, hierarchicalRef := range rc.HierarchicalReferences {
		assertHasRootLevelResourceReference(t, rc, hierarchicalRef.Key)
	}
	for _, container := range rc.Containers {
		assertHasHierarchicalReferenceForContainerType(t, rc, container.Type)
	}
}

func assertHasAtMostOneOfEachHierarchicalReferenceType(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	supportedTypes := make(map[v1alpha1.HierarchicalReferenceType]bool)
	for _, hierarchicalRef := range rc.HierarchicalReferences {
		if _, ok := supportedTypes[hierarchicalRef.Type]; ok {
			t.Fatalf("kind %v has more than one hierarchical reference with type %v", rc.Kind, hierarchicalRef.Type)
		}
		supportedTypes[hierarchicalRef.Type] = true
	}
}

func assertHasRootLevelResourceReference(t *testing.T, rc v1alpha1.ResourceConfig, key string) {
	t.Helper()
	if strings.Contains(key, ".") {
		t.Fatalf("key %v is a path, not a field", key)
	}
	for _, ref := range rc.ResourceReferences {
		if strings.Contains(ref.TFField, ".") {
			// Reference is not at the root-level of the spec.
			continue
		}
		if krmtotf.GetKeyForReferenceField(&ref) == key {
			return
		}
	}
	t.Fatalf("kind %v does not have a root-level resource reference with key %v", rc.Kind, key)
}

func assertHasHierarchicalReferenceForContainerType(t *testing.T, rc v1alpha1.ResourceConfig, containerType v1alpha1.ContainerType) {
	t.Helper()
	hierarchicalType := k8s.HierarchicalReferenceTypeFor(containerType)
	for _, hierarchicalRef := range rc.HierarchicalReferences {
		if hierarchicalRef.Type == hierarchicalType {
			return
		}
	}
	t.Fatalf("kind %v has a container of type %v, but no hierarchical reference of type %v", rc.Kind, containerType, hierarchicalType)
}

func TestHierarchicalReferencesForKindsWithMultipleResourceConfigs(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			kindToRCs := make(map[string][]v1alpha1.ResourceConfig)
			for _, rc := range sm.Spec.Resources {
				kindToRCs[rc.Kind] = append(kindToRCs[rc.Kind], rc)
			}
			for kind, rcs := range kindToRCs {
				if len(rcs) < 2 {
					continue
				}
				kind := kind
				rcs := rcs
				t.Run(kind, func(t *testing.T) {
					t.Parallel()
					assertAllHaveSameHierarchicalReferences(t, kind, rcs)
				})
			}
		})
	}
}

func assertAllHaveSameHierarchicalReferences(t *testing.T, kind string, rcs []v1alpha1.ResourceConfig) {
	t.Helper()
	if len(rcs) == 0 {
		return
	}
	hierarchicalRefs := rcs[0].HierarchicalReferences
	for _, rc := range rcs {
		if !reflect.DeepEqual(rc.HierarchicalReferences, hierarchicalRefs) {
			t.Errorf("not all ResourceConfigs of kind %v have the same HierarchicalReferences configuration", kind)
		}
	}
}

func TestMustHaveIDTemplateOrServerGeneratedId(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		for _, rc := range sm.Spec.Resources {
			rc := rc
			t.Run(rc.Kind, func(t *testing.T) {
				t.Parallel()
				assertIDTemplateOrServerGeneratedID(t, rc)
			})
		}
	}
}

func assertIDTemplateOrServerGeneratedID(t *testing.T, rc v1alpha1.ResourceConfig) {
	if !rc.Direct && rc.IDTemplate == "" && rc.ServerGeneratedIDField == "" {
		t.Fatalf("resource kind '%v' with name '%v' has neither id template or server generated ID defined: at least one must be present", rc.Kind, rc.Name)
	}
}

func TestIDTemplate(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					if rc.IDTemplate == "" {
						return
					}

					// The resource contains neither a user-specified ID nor a
					// server-generated ID.
					if rc.MetadataMapping.Name == "" && rc.ServerGeneratedIDField == "" {
						return
					}

					// The idTemplate should contain either the user-specified
					// ID or the server-generated ID.
					if (IDTemplateContainsMetadataName(rc) &&
						!IDTemplateContainsServerGeneratedIDField(rc)) ||
						(!IDTemplateContainsMetadataName(rc) &&
							IDTemplateContainsServerGeneratedIDField(rc)) {
						return
					}

					t.Fatalf("idTemplate of resource kind '%v' with name "+
						"'%v' contains 0 or 2 field names defined in "+
						"'metadata.name' and 'serverGeneratedIDField': "+
						"exactly one should be contained", rc.Kind, rc.Name)
				})
			}
		})
	}
}

func IDTemplateContainsMetadataName(rc v1alpha1.ResourceConfig) bool {
	return strings.Contains(rc.IDTemplate,
		fmt.Sprintf("{{%v}}", rc.MetadataMapping.Name))
}

func IDTemplateContainsServerGeneratedIDField(rc v1alpha1.ResourceConfig) bool {
	return strings.Contains(rc.IDTemplate,
		fmt.Sprintf("{{%v}}", rc.ServerGeneratedIDField))
}

func TestMutableButUnreadableFields(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					testMutableButUnreadableFields(t, rc, provider)
				})
			}
		})
	}
}

func testIamPolicyMappings(t *testing.T, rc v1alpha1.ResourceConfig) {
	if rc.IAMConfig.PolicyName == "" {
		assertIAMConfigIsEmpty(t, rc)
		assertIAMConfigShouldBeEmpty(t, rc)
	} else {
		assertIAMConfigValueIsValid(t, rc)
	}
}

func assertIAMConfigShouldBeEmpty(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	// TODO: Implement IAMPolicy support for:
	//  - BigQueryDataset (b/167223329)
	//  - ComputeDisk (b/168609794)
	switch rc.Name {
	case "google_bigquery_dataset", "google_compute_region_disk", "google_compute_disk":
		return
	}
	tfIamPolicyResourceName, tfIamPolicyResource := getAssociatedTerraformIAMPolicyResource(rc)
	if tfIamPolicyResource != nil {
		t.Errorf("kind '%v' is missing a valid IAMConfig, but a valid terraform IAM Policy resource '%v' exists",
			tfIamPolicyResourceName, tfIamPolicyResourceName)
	}
}

func assertAllHaveEmptyOrNonEmptyIAMConfigButNotBoth(t *testing.T, kind string, rcs []v1alpha1.ResourceConfig) {
	t.Helper()
	if len(rcs) == 0 {
		return
	}
	hasEmptyIAM := hasEmptyIAMConfig(rcs[0])
	for _, rc := range rcs {
		if hasEmptyIAMConfig(rc) != hasEmptyIAM {
			t.Errorf("all ResourceConfigs of kind %v must all have an empty or non-empty iamConfig, but not a mixture of both", kind)
		}
	}
}

func assertIAMConfigIsEmpty(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	if !hasEmptyIAMConfig(rc) {
		t.Errorf("invalid argument, iamConfig for resource '%v' is non-empty", rc.Kind)
	}
}

func assertIAMConfigValueIsValid(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	if rc.IAMConfig.ReferenceField.Name == "" {
		t.Errorf("invalid value for Name: value must be present")
	}

	tfIamPolicyResourceName, tfIamPolicyResource := getAssociatedTerraformIAMPolicyResource(rc)
	if rc.IAMConfig.PolicyName != tfIamPolicyResourceName {
		// if this exceptional case is valid then manually add an entry in formatAssociatedTerraformIAMPolicyResourceName(...) to return the correct value
		t.Fatalf("tf iampolicy name mismatch for kind '%v': value of '%v' does not match expected value of '%v'",
			rc.Kind, rc.IAMConfig.PolicyName, tfIamPolicyResourceName)
	}
	_, ok := tfIamPolicyResource.Schema[rc.IAMConfig.ReferenceField.Name]
	if !ok {
		t.Errorf("kind '%v' has an invalid value for ReferenceFieldName '%v': the terraform resource '%v' does not contain any field with that"+
			" name", rc.Kind, rc.IAMConfig.ReferenceField.Name, tfIamPolicyResourceName)
	}

	tfIamPolicyMemberResourceName, tfIamPolicyMemberResource := getAssociatedTerraformIAMPolicyMemberResource(rc)
	if rc.IAMConfig.PolicyMemberName != tfIamPolicyMemberResourceName {
		// if this exceptional case is valid then manually add an entry in formatAssociatedTerraformIAMPolicyMemberResourceName(...) to return the correct value
		t.Fatalf("tf iampolicy member name mismatch for kind '%v': value of '%v' does not match expected value of '%v'",
			rc.Kind, rc.IAMConfig.PolicyMemberName, tfIamPolicyMemberResourceName)
	}
	_, ok = tfIamPolicyMemberResource.Schema[rc.IAMConfig.ReferenceField.Name]
	if !ok {
		t.Errorf("kind '%v' has an invalid value for ReferenceFieldName '%v': the terraform resource '%v' does not contain any field with that"+
			" name", rc.Kind, rc.IAMConfig.ReferenceField.Name, tfIamPolicyMemberResourceName)
	}

	if rc.IAMConfig.AuditConfigName != "" {
		tfIamAuditConfigResourceName, tfIamAuditConfigResource := getAssociatedTerraformIAMAuditConfigResource(rc)
		if rc.IAMConfig.AuditConfigName != tfIamAuditConfigResourceName {
			// if this exceptional case is valid then manually add an entry in formatAssociatedTerraformIAMAuditConfigResourceName(...) to return the correct value
			t.Fatalf("tf auditconfig name mismatch for kind '%v': value of '%v' does not match expected value of '%v'",
				rc.Kind, rc.IAMConfig.AuditConfigName, tfIamAuditConfigResourceName)
		}
		_, ok = tfIamAuditConfigResource.Schema[rc.IAMConfig.ReferenceField.Name]
		if !ok {
			t.Errorf("kind '%v' has an invalid value for ReferenceFieldName '%v': the terraform resource '%v' does not contain any field with that"+
				" name", rc.Kind, rc.IAMConfig.ReferenceField.Name, tfIamAuditConfigResourceName)
		}
	}

	assertValidAndUsableIAMReferenceValueType(t, rc)
}

func assertValidAndUsableIAMReferenceValueType(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	value := rc.IAMConfig.ReferenceField.Type
	switch value {
	case v1alpha1.IAMReferenceTypeName:
	case v1alpha1.IAMReferenceTypeId:
		if rc.IDTemplate == "" && rc.ServerGeneratedIDField == "" {
			msg := "to use this value type, either the IDTemplate or ServerGeneratedIDField fields must contain a value"
			t.Errorf("invalid usage of reference value type '%v': %v", value, msg)
		}
	default:
		t.Errorf("unknown value type value: %v", value)
	}
}

func assertAllHaveSameValueForSupportsConditions(t *testing.T, kind string, rcs []v1alpha1.ResourceConfig) {
	t.Helper()
	if len(rcs) == 0 {
		return
	}
	supportsConditions := rcs[0].IAMConfig.SupportsConditions
	for _, rc := range rcs {
		if rc.IAMConfig.SupportsConditions != supportsConditions {
			t.Errorf("not all ResourceConfigs of kind %v have the same value for iamConfig.supportsConditions", kind)
		}
	}
}

func assertAllOrNoneSupportAuditConfigs(t *testing.T, kind string, rcs []v1alpha1.ResourceConfig) {
	t.Helper()
	if len(rcs) == 0 {
		return
	}
	supportsAuditConfigs := rcs[0].IAMConfig.AuditConfigName != ""
	for _, rc := range rcs {
		rcSupportsAuditConfigs := rc.IAMConfig.AuditConfigName != ""
		if rcSupportsAuditConfigs != supportsAuditConfigs {
			t.Errorf("all ResourceConfigs of kind %v must support or not support IAM audit configs, but not a mixture of both", kind)
		}
	}
}

func getAssociatedTerraformIAMPolicyResource(rc v1alpha1.ResourceConfig) (string, *schema.Resource) {
	schemaProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	tfIamPolicyResourceName := formatAssociatedTerraformIAMPolicyResourceName(rc)
	return tfIamPolicyResourceName, schemaProvider.ResourcesMap[tfIamPolicyResourceName]
}

func formatAssociatedTerraformIAMPolicyResourceName(rc v1alpha1.ResourceConfig) string {
	switch rc.Name {
	case "google_compute_instance_from_template":
		return "google_compute_instance_iam_policy"
	default:
		return fmt.Sprintf("%v_iam_policy", rc.Name)

	}
}

func getAssociatedTerraformIAMPolicyMemberResource(rc v1alpha1.ResourceConfig) (string, *schema.Resource) {
	schemaProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	tfIamPolicyResourceName := formatAssociatedTerraformIAMPolicyMemberResourceName(rc)
	return tfIamPolicyResourceName, schemaProvider.ResourcesMap[tfIamPolicyResourceName]
}

func formatAssociatedTerraformIAMPolicyMemberResourceName(rc v1alpha1.ResourceConfig) string {
	switch rc.Name {
	case "google_compute_instance_from_template":
		return "google_compute_instance_iam_member"
	default:
		return fmt.Sprintf("%v_iam_member", rc.Name)
	}
}

func getAssociatedTerraformIAMAuditConfigResource(rc v1alpha1.ResourceConfig) (string, *schema.Resource) {
	schemaProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	tfIamAuditConfigResourceName := formatAssociatedTerraformIAMAuditConfigResourceName(rc)
	return tfIamAuditConfigResourceName, schemaProvider.ResourcesMap[tfIamAuditConfigResourceName]
}

func formatAssociatedTerraformIAMAuditConfigResourceName(rc v1alpha1.ResourceConfig) string {
	return fmt.Sprintf("%v_iam_audit_config", rc.Name)
}

func hasEmptyIAMConfig(rc v1alpha1.ResourceConfig) bool {
	return reflect.DeepEqual(rc.IAMConfig, emptyIAMConfig)
}

func createKindToTFResourcesMap(sms []v1alpha1.ServiceMapping) map[string][]string {
	kindToTFResources := make(map[string][]string)
	for _, sm := range sms {
		for _, rc := range sm.Spec.Resources {
			if _, ok := kindToTFResources[rc.Kind]; !ok {
				kindToTFResources[rc.Kind] = make([]string, 0)
			}
			kindToTFResources[rc.Kind] = slice.IncludeString(kindToTFResources[rc.Kind], rc.Name)
		}
	}
	return kindToTFResources
}

func TestIAMMemberReferenceConfig(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					iamMemberRefConfig := rc.IAMMemberReferenceConfig
					if iamMemberRefConfig.TargetField != "" {
						testIAMMemberReferenceConfig(t, rc, provider)
					}
				})
			}
		})
	}
}

func testIAMMemberReferenceConfig(t *testing.T, rc v1alpha1.ResourceConfig, provider *schema.Provider) {
	tfResource := provider.ResourcesMap[rc.Name]
	targetField := rc.IAMMemberReferenceConfig.TargetField
	if !tfresource.TFResourceHasField(tfResource, targetField) {
		t.Errorf("kind %v has its iamMemberReference.targetField set to %v, "+
			"but no such field exists in the Terraform resource %v",
			rc.Kind, targetField, rc.Name)
	}
}

func TestResourceIDForKindsWithMultipleResourceConfigs(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			kindToRCs := make(map[string][]v1alpha1.ResourceConfig)
			for _, rc := range sm.Spec.Resources {
				kindToRCs[rc.Kind] = append(kindToRCs[rc.Kind], rc)
			}
			for kind, rcs := range kindToRCs {
				if len(rcs) < 2 {
					continue
				}
				kind := kind
				rcs := rcs
				t.Run(kind, func(t *testing.T) {
					t.Parallel()
					assertAllHaveSameResourceIDConfigs(t, kind, rcs)
				})
			}
		})
	}
}

func assertAllHaveSameResourceIDConfigs(t *testing.T, kind string, rcs []v1alpha1.ResourceConfig) {
	t.Helper()
	if len(rcs) == 0 {
		return
	}

	targetField := rcs[0].ResourceID.TargetField
	valueTemplate := rcs[0].ResourceID.ValueTemplate
	for _, rc := range rcs {
		if rc.ResourceID.TargetField != targetField || rc.ResourceID.ValueTemplate != valueTemplate {
			t.Fatalf("not all ResourceConfigs of kind %v have the same value for resourceID.targetField or resourceID.valueTemplate", kind)
		}
	}
}

func TestVersionForKindsWithMultipleResourceConfigs(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			kindToRCs := make(map[string][]v1alpha1.ResourceConfig)
			for _, rc := range sm.Spec.Resources {
				kindToRCs[rc.Kind] = append(kindToRCs[rc.Kind], rc)
			}
			for kind, rcs := range kindToRCs {
				if len(rcs) < 2 {
					continue
				}
				kind := kind
				rcs := rcs
				t.Run(kind, func(t *testing.T) {
					t.Parallel()
					assertAllHaveSameVersion(t, kind, rcs, &sm)
				})
			}
		})
	}
}

func assertAllHaveSameVersion(t *testing.T, kind string, rcs []v1alpha1.ResourceConfig, sm *v1alpha1.ServiceMapping) {
	t.Helper()
	if len(rcs) == 0 {
		return
	}

	version := sm.GetVersionFor(&rcs[0])

	for _, rc := range rcs {
		if newVersion := sm.GetVersionFor(&rc); newVersion != version {
			t.Fatalf("ResourceConfigs of kind %v have more than one version: %v, %v", kind, version, newVersion)
		}
	}
}

func testMutableButUnreadableFields(t *testing.T, rc v1alpha1.ResourceConfig, provider *schema.Provider) {
	tfResource := provider.ResourcesMap[rc.Name]
	for _, field := range rc.MutableButUnreadableFields {
		tfSchema, err := tfresource.GetTFSchemaForField(tfResource, field)
		if err != nil {
			t.Fatalf("error getting Terraform schema for field '%v': %v", field, err)
		}
		if !tfresource.IsConfigurableField(tfSchema) {
			t.Fatalf("'%v' is marked mutable-but-unreadable, but cannot have mutable-but-unreadable fields that are not configurable", field)
		}
		if tfSchema.ForceNew {
			t.Fatalf("'%v' is marked mutable-but-unreadable, but cannot have mutable-but-unreadable fields that are marked immutable", field)
		}
		if tfresource.IsFieldNestedInList(tfResource, field) {
			t.Fatalf("'%v' is marked mutable-but-unreadable, but cannot have mutable-but-unreadable fields that are nested in lists", field)
		}
		if slice.StringSliceContains(rc.IgnoredFields, field) {
			t.Fatalf("'%v' is marked mutable-but-unreadable, but cannot have mutable-but-unreadable fields that are marked as ignored fields", field)
		}
		if slice.StringSliceContains(rc.Directives, field) {
			t.Fatalf("'%v' is marked mutable-but-unreadable, but cannot have mutable-but-unreadable fields that are marked as directives", field)
		}
		if field == rc.MetadataMapping.Name || field == rc.MetadataMapping.Labels {
			t.Fatalf("'%v' is marked mutable-but-unreadable, but cannot have mutable-but-unreadable fields that are marked as metadata fields", field)
		}
		for _, resourceRef := range rc.ResourceReferences {
			if field == resourceRef.TFField {
				t.Fatalf("'%v' is marked mutable-but-unreadable, but cannot have mutable-but-unreadable fields that are marked as resource references", field)
			}
		}
	}
}

func typeConfigsOf(resourceRef v1alpha1.ReferenceConfig) []v1alpha1.TypeConfig {
	if len(resourceRef.Types) == 0 {
		return []v1alpha1.TypeConfig{resourceRef.TypeConfig}
	}
	return resourceRef.Types
}

func TestResourceID(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					if rc.ResourceID.TargetField == "" {
						// Resource ID field is not supported so no test is
						// needed.
						return
					}

					// Empty idTemplate field means that the resource uses
					// the server-generated ID as the ID template.
					if rc.IDTemplate == "" {
						testServerGeneratedResourceID(t, rc)
						return
					}

					// If the idTemplate contains the TF field that
					// metadata.name maps to, then the resource has a
					// user-specified resource ID.
					// If the idTemplate contains the TF field that
					// status.[serverGeneratedIDField] maps to, then the
					// resource has a server-generated resource ID.
					// Otherwise, resourceID should not be supported.
					if strings.Contains(rc.IDTemplate,
						fmt.Sprintf("{{%v}}", rc.MetadataMapping.Name)) {
						testUserSpecifiedResourceID(t, rc)
					} else if strings.Contains(rc.IDTemplate,
						fmt.Sprintf("{{%v}}", rc.ServerGeneratedIDField)) {
						testServerGeneratedResourceID(t, rc)
					} else {
						t.Fatalf("resourceID in ResourceConfig %s shouldn't "+
							"be supported if the resource has neither a "+
							"user-specified ID nor a server-generated ID",
							rc.Name)
					}
				})
			}
		})
	}
}

func testUserSpecifiedResourceID(t *testing.T, rc v1alpha1.ResourceConfig) {
	if rc.ResourceID.TargetField != rc.MetadataMapping.Name {
		t.Fatalf("targetField of user-specified resourceID in "+
			"ResourceConfig %s is different from value of "+
			"metadataMapping.name", rc.Name)
	}
	if rc.ResourceID.ValueTemplate != rc.MetadataMapping.NameValueTemplate {
		t.Fatalf("valueTemplate of user-specified resourceID in "+
			"ResourceConfig %s is different from value of "+
			"metadataMapping.nameValueTemplate", rc.Name)
	}
}

func testServerGeneratedResourceID(t *testing.T, rc v1alpha1.ResourceConfig) {
	if rc.ResourceID.TargetField != rc.ServerGeneratedIDField {
		t.Fatalf("targetField of server-generated resourceID in "+
			"ResourceConfig %s is different from value of "+
			"serverGeneratedIDField", rc.Name)
	}
}

// TestUnreadableResourcesShouldHaveZeroReconciliationInterval ensures that all resources that are
// unreadable have set ReconciliationIntervalInSeconds to 0.
func TestUnreadableResourcesShouldHaveZeroReconciliationInterval(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					if rc.Unreadable == nil || *rc.Unreadable == false {
						return
					}
					if rc.ReconciliationIntervalInSeconds == nil || *rc.ReconciliationIntervalInSeconds != 0 {
						t.Fatalf("resource config '%v' is marked 'Unreadable', but field 'ReconciliationIntervalInSeconds' is not set to 0", rc.Name)
					}
				})
			}
		})
	}
}

// TestReconciliationIntervalConsistency makes sure the configured reconciliation intervals have
// the same value for all resource configs mapped to the same GVK.
func TestReconciliationIntervalConsistency(t *testing.T) {
	smLoader := testservicemappingloader.New(t)
	for _, gvk := range supportedgvks.BasedOnAllServiceMappings(smLoader) {
		rcs, err := smLoader.GetResourceConfigs(gvk)
		if err != nil || len(rcs) < 2 {
			// only check for GVKs mapped to multiple resource configs
			continue
		}
		var ri *uint32
		for _, rc := range rcs {
			if rc.ReconciliationIntervalInSeconds == nil {
				// ReconciliationIntervalInSeconds not configured
				continue
			}
			if ri == nil {
				// first time seeing ReconciliationIntervalInSeconds for this GVK
				ri = new(uint32)
				*ri = *rc.ReconciliationIntervalInSeconds
				continue
			}
			if *ri != *rc.ReconciliationIntervalInSeconds {
				t.Errorf("the configured reconciliation intervals "+
					"should have the same value for all resource configs "+
					"mapped to GVK %v", gvk)
			}
		}
	}
}

func isAutogenAlphaResource(sm *v1alpha1.ServiceMapping, rc *v1alpha1.ResourceConfig) bool {
	if sm.GetVersionFor(rc) == k8s.KCCAPIVersionV1Alpha1 && rc.AutoGenerated {
		return true
	}
	return false
}

func TestDCLBasedResourceIsTrueIFFIsDCLBasedResource(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	referencedDCLResources := make([]k8sschema.GroupVersionKind, 0)
	referencedTFResources := make([]k8sschema.GroupVersionKind, 0)
	for _, sm := range serviceMappings {
		for _, r := range sm.Spec.Resources {
			if r.AutoGenerated {
				continue
			}
			for _, rr := range r.ResourceReferences {
				if rr.DCLBasedResource {
					referencedDCLResources = append(referencedDCLResources, rr.GVK)
				} else {
					referencedTFResources = append(referencedTFResources, rr.GVK)
				}
			}
		}
	}
	smLoader := dclmetadata.New()
	for _, gvk := range referencedDCLResources {
		r, found := smLoader.GetResourceWithGVK(gvk)
		if !found || !r.Releasable {
			t.Errorf("%v is listed in servicemappings as a resource reference with "+
				"`DCLBasedResource: true`, but it is not a DCL-based resource", gvk)
		}
	}
	for _, gvk := range referencedTFResources {
		r, found := smLoader.GetResourceWithGVK(gvk)
		if found && r.Releasable {
			t.Errorf("%v is listed in servicemappings as a resource reference with "+
				"`DCLBasedResource: false`, but it is a DCL-based resource", gvk)
		}
	}
}

// TestV1alpha1ToV1beta1IsSetForManuallyConfiguredAndAllowlistedResources
// verifies that IFF when a resource is allowlisted and auto-generated, but is
// also manually configured under config/servicemappings, it's a resource under
// v1alpha1 -> v1beta1 conversion, and `v1alpha1ToV1beta1: true` must be set.
func TestV1alpha1ToV1beta1IsSetForManuallyConfiguredAndAllowlistedResources(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	autoGenSMMap, err := autogenloader.GetServiceMappingMap()
	if err != nil {
		t.Fatalf("error getting auto-generated service mappings: %v", err)
	}

	for _, sm := range serviceMappings {
		autoGenSM, ok := autoGenSMMap[sm.Name]
		if !ok {
			continue
		}
		autoGenRCMap := make(map[string]bool)
		for _, rc := range autoGenSM.Spec.Resources {
			autoGenRCMap[rc.Name] = true
		}
		for _, r := range sm.Spec.Resources {
			isV1alpha1ToV1beta1 := r.V1alpha1ToV1beta1 != nil && *r.V1alpha1ToV1beta1 == true
			if r.AutoGenerated && isV1alpha1ToV1beta1 {
				t.Errorf("resource config %v is auto-generated "+
					"and allowlisted, but has `v1alpha1ToV1beta1: true`: "+
					"`v1alpha1ToV1beta1` should be unset", r.Name)
				continue
			}
			if !r.AutoGenerated {
				if _, ok := autoGenRCMap[r.Name]; !ok {
					continue
				}
				if !isV1alpha1ToV1beta1 {
					t.Errorf("resource config %v is manually configured "+
						"and allowlisted, but doesn't have `v1alpha1ToV1beta1: "+
						"true`", r.Name)
				}
			}
		}
	}
}

// TestStorageVersionIsSetAndValidIFFV1alpha1ToV1beta1IsSet verifies that a
// valid `storageVersion` is set and only set when `v1alpha1ToV1beta1: true`.
// Currently, storage version is required and only required when the resource is
// under v1alpha1 -> v1beta1 conversion.
func TestStorageVersionIsSetAndValidIFFV1alpha1ToV1beta1IsSet(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		for _, r := range sm.Spec.Resources {
			isV1alpha1ToV1beta1 := false
			if r.V1alpha1ToV1beta1 != nil && *r.V1alpha1ToV1beta1 {
				isV1alpha1ToV1beta1 = true
			}
			hasStorageVersion := false
			hasValidStorageVersion := false
			if r.StorageVersion != nil {
				hasStorageVersion = true
				if crdgeneration.IsValidStorageVersion(*r.StorageVersion) {
					hasValidStorageVersion = true
				}
			}
			if isV1alpha1ToV1beta1 && hasValidStorageVersion ||
				!isV1alpha1ToV1beta1 && !hasStorageVersion {
				continue
			}
			if isV1alpha1ToV1beta1 {
				// if this is a direct resource, the storage version is defiend
				// in the kubebuilder tooling
				if r.Direct {
					continue
				}
				if hasStorageVersion {
					t.Errorf("Resource config %v has `v1alpha1ToV1beta1: "+
						"true` but doesn't have a valid `storageVersion`: "+
						"must be %v or %v", r.Name, k8s.KCCAPIVersionV1Alpha1,
						k8s.KCCAPIVersionV1Beta1)
					continue
				}

				t.Errorf("Resource config %v has `v1alpha1ToV1beta1: "+
					"true` but doesn't have a `storageVersion`", r.Name)
				continue
			}
			if hasStorageVersion {
				t.Errorf("Resource config %v has `storageVersion` set "+
					"but doesn't have a `v1alpha1ToV1beta1: true`", r.Name)
			}
		}
	}
}

func TestObservedFields(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				tfResource := provider.ResourcesMap[rc.Name]
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					if rc.ObservedFields == nil {
						return
					}
					if len(*rc.ObservedFields) == 0 {
						t.Errorf("kind %v has an empty observed fields slice", rc.Kind)
						return
					}
					for _, f := range *rc.ObservedFields {
						if f == "" {
							t.Errorf("kind %v has an empty observed field", rc.Kind)
							return
						}
						// Nested fields (configurable or output-only) will be
						// under 'spec' if their top-level parent fields are
						// configurable.
						topLevelField := strings.Split(f, ".")[0]
						fieldSchema, err := tfresource.GetTFSchemaForField(tfResource, topLevelField)
						if err != nil {
							t.Errorf("error getting TF schema for observed field %v in kind %v", f, rc.Kind)
						}
						if !tfresource.IsConfigurableField(fieldSchema) {
							t.Errorf("observed field %v in kind %v is not configurable", f, rc.Kind)
						}
					}
					assertNoDuplicatesInObservedFieldsSlice(t, rc)
					// TODO(b/314840974): Remove after reference fields are supported as observed fields.
					assertObservedFieldsNotReferences(t, rc)
					// TODO(b/314841141): Remove after label fields are supported as observed fields.
					assertObservedFieldsNotLabels(t, rc)
					// TODO(b/314842047): Remove after name fields are supported as observed fields.
					assertObservedFieldsNotName(t, rc)
					// TODO(b/314841744): Remove after sensitive fields are supported as observed fields.
					assertObservedFieldsNotSensitive(t, rc, tfResource)
				})
			}
		})
	}
}

func assertNoDuplicatesInObservedFieldsSlice(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	if len(*rc.ObservedFields) == 0 {
		t.Errorf("kind %v has no observed field", rc.Kind)
		return
	}
	observedFieldMap := make(map[string]bool)
	for _, field := range *rc.ObservedFields {
		if _, ok := observedFieldMap[field]; ok {
			t.Errorf("kind %v contains duplicated observed field %v", rc.Kind, field)
			continue
		}
		observedFieldMap[field] = true
	}
	return
}

func assertObservedFieldsNotReferences(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	referenceFields := make(map[string]bool)
	for _, refConfig := range rc.ResourceReferences {
		referenceFields[refConfig.TFField] = true
	}
	for _, field := range *rc.ObservedFields {
		if _, ok := referenceFields[field]; ok {
			t.Errorf("kind %v contains observed field %v: reference "+
				"fields are not supported as observed fields", rc.Kind, field)
		}
	}
}

func assertObservedFieldsNotLabels(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	for _, field := range *rc.ObservedFields {
		if field != "" && field == rc.MetadataMapping.Labels {
			t.Errorf("kind %v contains observed field %v: labels "+
				"field is not supported as ab observed field", rc.Kind, field)
		}
	}
}

func assertObservedFieldsNotName(t *testing.T, rc v1alpha1.ResourceConfig) {
	t.Helper()
	for _, field := range *rc.ObservedFields {
		if field != "" && (field == rc.MetadataMapping.Name ||
			field == rc.ServerGeneratedIDField) {
			t.Errorf("kind %v contains observed field %v: name "+
				"field is not supported as an observed field", rc.Kind, field)
		}
	}
}

func assertObservedFieldsNotSensitive(t *testing.T, rc v1alpha1.ResourceConfig, tfResource *schema.Resource) {
	t.Helper()
	for _, field := range *rc.ObservedFields {
		tfSchema, err := tfresource.GetTFSchemaForField(tfResource, field)
		if err != nil {
			t.Errorf("error getting Terraform schema for observed "+
				"field %v in kind %v: %v", field, rc.Kind, err)
		}
		if tfresource.IsSensitiveField(tfSchema) {
			t.Errorf("kind %v contains observed field %v: sensitive "+
				"fields are not supported as observed fields", rc.Kind, field)
		}
	}

}

func TestAlphaResourceAreNotReferencedByStableResource(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	for _, sm := range serviceMappings {
		for _, rc := range sm.Spec.Resources {
			rc := rc
			t.Run(rc.Kind, func(t *testing.T) {
				t.Parallel()
				if sm.GetVersionFor(&rc) == k8s.KCCAPIVersionV1Alpha1 {
					return
				}
				assertReferencedResourcesNotAlpha(t, &rc)
			})
		}
	}
}

func assertReferencedResourcesNotAlpha(t *testing.T, rc *v1alpha1.ResourceConfig) {
	for _, refConfig := range rc.ResourceReferences {
		if len(refConfig.Types) == 0 {
			if refConfig.TypeConfig.GVK.Version == k8s.KCCAPIVersionV1Alpha1 {
				t.Errorf("cannot reference %s resource %s in stable resource %s", k8s.KCCAPIVersionV1Alpha1, refConfig.TypeConfig.GVK.Kind, rc.Kind)
			}
		} else {
			for _, typeConfig := range refConfig.Types {
				if typeConfig.GVK.Version == k8s.KCCAPIVersionV1Alpha1 {
					t.Errorf("cannot reference %s resource %s in stable resource %s", k8s.KCCAPIVersionV1Alpha1, typeConfig.GVK.Kind, rc.Kind)
				}
			}
		}
	}
}

func TestIgnoredOutputOnlySpecFields(t *testing.T) {
	t.Parallel()
	serviceMappings := testservicemappingloader.New(t).GetServiceMappings()
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	for _, sm := range serviceMappings {
		sm := sm
		t.Run(sm.Name, func(t *testing.T) {
			t.Parallel()
			for _, rc := range sm.Spec.Resources {
				tfResource := provider.ResourcesMap[rc.Name]
				rc := rc
				t.Run(rc.Kind, func(t *testing.T) {
					t.Parallel()
					if rc.IgnoredOutputOnlySpecFields == nil {
						return
					}
					if len(*rc.IgnoredOutputOnlySpecFields) == 0 {
						t.Errorf("kind %v has an empty IgnoredOutputOnlySpecFields slice", rc.Kind)
						return
					}
					for _, f := range *rc.IgnoredOutputOnlySpecFields {
						if f == "" {
							t.Errorf("kind %v has an empty value in IgnoredOutputOnlySpecFields slice", rc.Kind)
							return
						}
						fieldSchema, err := tfresource.GetTFSchemaForField(tfResource, f)
						if err != nil {
							t.Errorf("error getting TF schema for output-only spec field %v in kind %v", f, rc.Kind)
						}
						if tfresource.IsConfigurableField(fieldSchema) {
							t.Errorf("output-only spec field %v in kind %v is configurable", f, rc.Kind)
						}
					}
				})
			}
		})
	}
}
