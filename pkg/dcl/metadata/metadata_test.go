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

package metadata_test

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

func TestNoDuplicateService(t *testing.T) {
	t.Parallel()
	smLoader := metadata.New()
	allServices := smLoader.GetAllServiceMetadata()
	m := make(map[string]bool)
	for _, s := range allServices {
		serviceName := metadata.CanonicalizeServiceName(s.Name)
		if _, ok := m[metadata.CanonicalizeServiceName(serviceName)]; ok {
			t.Fatalf("found duplicate service: %s", s.Name)
		}
		m[serviceName] = true
	}
}

func TestServiceNamedDifferentlyInDCLAndKCC(t *testing.T) {
	t.Parallel()
	serviceList := []metadata.ServiceMetadata{
		{
			Name:                 "KMS",
			APIVersion:           "v1beta1",
			ServiceNameUsedByDCL: "cloudkms",
		},
	}
	smLoader := metadata.NewFromServiceList(serviceList)
	serviceNameInDCL := "cloudkms"
	_, found := smLoader.GetServiceMetadata(serviceNameInDCL)
	if !found {
		t.Fatalf("service metadata for service %v is not found", serviceNameInDCL)
	}
	serviceNameInKCC := "KMS"
	_, found = smLoader.GetServiceMetadata(serviceNameInKCC)
	if !found {
		t.Fatalf("service metadata for service %v is not found", serviceNameInKCC)
	}
}

func TestResourceNamesAreCaseInsensitiveEqual(t *testing.T) {
	t.Parallel()
	smLoader := metadata.New()
	allServices := smLoader.GetAllServiceMetadata()
	for _, s := range allServices {
		for _, r := range s.Resources {
			gvk := metadata.GVKForResource(s, r)
			kindWithoutService := k8s.KindWithoutServicePrefix(gvk)
			if strings.ToLower(kindWithoutService) != strings.ToLower(r.DCLType) {
				t.Fatalf("Kind %v (with service prefix ignored) is not the same (case-insensitive) with the DCL Type %v, there might be a typo.", r.Kind, r.DCLType)
			}
		}
	}
}

func TestServicesAndResourcesListedAlphabetically(t *testing.T) {
	t.Parallel()
	smLoader := metadata.New()
	allServices := smLoader.GetAllServiceMetadata()
	var prevServiceName string
	for _, s := range allServices {
		if prevServiceName > s.Name {
			t.Errorf("services are not listed alphabetically: %v listed before %v", prevServiceName, s.Name)
		}
		prevServiceName = s.Name
		var preResourceKind string
		for _, r := range s.Resources {
			if preResourceKind > r.Kind {
				t.Errorf("resource are not listed alphabetically: %v listed before %v", preResourceKind, r.Kind)
			}
			preResourceKind = r.Kind
		}
	}
}

// TODO(b/186159460): Delete this test once all DCL-based resources support
// hierarchial references and the SupportsHierarchicalReferences flag is
// deleted since that will allow for all resources that should support
// hierarchical references to automatically support hierarchical references.
func TestResourceSupportsHierarchicalReferencesIfHasParentReferenceField(t *testing.T) {
	t.Parallel()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a DCL schema loader: %v", err)
	}
	smLoader := metadata.New()
	allServices := smLoader.GetAllServiceMetadata()
	for _, s := range allServices {
		for _, r := range s.Resources {
			s := s
			r := r
			t.Run(r.Kind, func(t *testing.T) {
				t.Parallel()
				if !r.Releasable {
					if r.SupportsHierarchicalReferences {
						t.Fatalf("kind %v should not have SupportsHierarchicalReferences=true since it is not Releasable", r.Kind)
					}
					return
				}

				gvk := metadata.GVKForResource(s, r)
				schema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, smLoader, dclSchemaLoader)
				if err != nil {
					t.Fatalf("error getting the DCL schema for GroupVersionKind %v: %v", gvk, err)
				}

				hasParentReferenceField := false
				for _, f := range dcl.ParentReferenceFields() {
					s, ok := schema.Properties[f]
					if ok && !s.ReadOnly {
						hasParentReferenceField = true
						break
					}
				}

				if hasParentReferenceField && !r.SupportsHierarchicalReferences {
					// Don't fail the test for resources that existed before
					// hierarchical references support was added for DCL-based
					// resources. The intention is to support hierarchical
					// references for these resources on a rolling basis.
					if _, ok := kindsThatPrecededHierarchicalRefs[r.Kind]; ok {
						return
					}
					t.Fatalf("kind %v has a parent reference field, but has SupportsHierarchicalReferences=false", r.Kind)
				} else if !hasParentReferenceField && r.SupportsHierarchicalReferences {
					t.Fatalf("kind %v doesn't have a parent reference field, but has SupportsHierarchicalReferences=true", r.Kind)
				}
			})
		}
	}
}

func TestResourceSupportsContainerAnnotationsOnlyIfPrecededHierarchicalReferences(t *testing.T) {
	t.Parallel()
	smLoader := metadata.New()
	allServices := smLoader.GetAllServiceMetadata()
	for _, s := range allServices {
		for _, r := range s.Resources {
			r := r
			t.Run(r.Kind, func(t *testing.T) {
				t.Parallel()
				if !r.SupportsContainerAnnotations {
					return
				}
				if !r.Releasable {
					t.Fatalf("kind %v should not have SupportsContainerAnnotations=true since it is not Releasable", r.Kind)
				}
				if _, ok := kindsThatPrecededHierarchicalRefs[r.Kind]; !ok {
					t.Fatalf("kind %v should not have SupportsContainerAnnotations=true since it is not one of the resources that preceded hierarchical references support", r.Kind)
				}
			})
		}
	}
}

// kindsThatPrecededHierarchicalRefs lists the kinds that had been supported by
// the KCC-DCL bridge before hierarchical references support was added for
// DCL-based resources.
var kindsThatPrecededHierarchicalRefs = map[string]bool{
	"CloudSchedulerJob":                    true,
	"ContainerAnalysisNote":                true,
	"DataFusionInstance":                   true,
	"DataprocAutoscalingPolicy":            true,
	"DataprocCluster":                      true,
	"DataprocWorkflowTemplate":             true,
	"GKEHubMembership":                     true,
	"IAPBrand":                             true,
	"IAPIdentityAwareProxyClient":          true,
	"IdentityPlatformOAuthIDPConfig":       true,
	"IdentityPlatformTenant":               true,
	"IdentityPlatformTenantOAuthIDPConfig": true,
	"MonitoringGroup":                      true,
	"NetworkSecurityClientTLSPolicy":       true,
	"NetworkSecurityServerTLSPolicy":       true,
	"OSConfigGuestPolicy":                  true,
}

func TestServiceListApiVersionsAreAllV1Beta1(t *testing.T) {
	t.Parallel()
	smLoader := metadata.New()
	allServices := smLoader.GetAllServiceMetadata()
	for _, s := range allServices {
		if s.APIVersion != k8s.KCCAPIVersionV1Beta1 {
			t.Errorf("service %v should have APIVersion set to %v using the constant k8s.KCCAPIVersionV1Beta1", s.Name, k8s.KCCAPIVersionV1Beta1)
		}
	}
}
