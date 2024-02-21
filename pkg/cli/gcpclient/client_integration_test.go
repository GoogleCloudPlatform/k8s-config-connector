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

//go:build integration
// +build integration

package gcpclient_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/gcpclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testconstants "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/constants"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestGetNotFound(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	serviceUsageSM, err := smLoader.GetServiceMapping("serviceusage.cnrm.cloud.google.com")
	if err != nil {
		t.Fatalf("error getting service mapping: %v", err)
	}
	serviceResource := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": fmt.Sprintf("serviceusage.cnrm.cloud.google.com/%v", serviceUsageSM.Spec.Version),
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name": "invalid.service.name.googleapis.com",
			},
		},
	}
	projectID := testgcp.GetDefaultProjectID(t)
	applyProjectRefOrAnnotation(t, smLoader, serviceResource, projectID)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.Config{})
	client := gcpclient.New(tfProvider, smLoader)
	result, err := client.Get(ctx, serviceResource)
	if err != gcpclient.ErrNotFound {
		t.Fatalf("unexpected error value: got '%v', want '%v", err, gcpclient.ErrNotFound)
	}
	if result != nil {
		t.Fatalf("unexpected result value: got '%v', want '%v'", result, nil)
	}
}

func init() {
	// TODO: Can we initialize this globally once somewhere?
	// (It shouldn't _really_ matter if we have multiple instances, it's just leaky and confusing)
	transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		artifacts := os.Getenv("ARTIFACTS")
		if artifacts == "" {
			log.Printf("env var ARTIFACTS is not set; will not record http log")
			return inner
		}
		outputDir := filepath.Join(artifacts, "http-logs")
		t := test.NewHTTPRecorder(inner.Transport, test.NewDirectoryEventSink(outputDir))
		return &http.Client{Transport: t}
	}
	transport_tpg.OAuth2HTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		artifacts := os.Getenv("ARTIFACTS")
		if artifacts == "" {
			return inner
		}
		outputDir := filepath.Join(artifacts, "http-logs")
		t := test.NewHTTPRecorder(inner.Transport, test.NewDirectoryEventSink(outputDir))
		return &http.Client{Transport: t}
	}
}

func TestCRUD(t *testing.T) {
	ctx := context.TODO()

	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.Config{})
	client := gcpclient.New(tfProvider, smLoader)
	projectID := testgcp.GetDefaultProjectID(t)
	testFunc := func(ctx context.Context, t *testing.T, testContext testrunner.TestContext) {
		var resources []*unstructured.Unstructured
		for _, d := range testContext.DependencyUnstructs {
			resolveAPIServerDependenciesIfKCCManaged(t, smLoader, tfProvider, resources, d)
			if k8s.IsManagedByKCC(d.GroupVersionKind()) {
				applyProjectRefOrAnnotation(t, smLoader, d, projectID)
				defer buildDeleteFunc(t, client, d)()
				d = clientApply(t, client, d)
			}
			resources = append(resources, d)
		}
		createUnstruct := testContext.CreateUnstruct
		applyProjectRefOrAnnotation(t, smLoader, createUnstruct, projectID)
		resolveAPIServerDependenciesIfKCCManaged(t, smLoader, tfProvider, resources, createUnstruct)
		defer buildDeleteFunc(t, client, createUnstruct)()
		clientApply(t, client, createUnstruct)
		clientGet(t, ctx, client, createUnstruct)
		if testContext.UpdateUnstruct != nil {
			resolveAPIServerDependenciesIfKCCManaged(t, smLoader, tfProvider, resources, testContext.UpdateUnstruct)
			applyProjectRefOrAnnotation(t, smLoader, testContext.UpdateUnstruct, projectID)
			clientApply(t, client, testContext.UpdateUnstruct)
		}
		clientDelete(t, client, createUnstruct)
		for i := len(testContext.DependencyUnstructs) - 1; i >= 0; i-- {
			d := testContext.DependencyUnstructs[i]
			if !k8s.IsManagedByKCC(d.GroupVersionKind()) {
				continue
			}
			clientDelete(t, client, d)
		}
	}
	supportedFixtures := getSupportedResourcesSetCover(t, client)
	testrunner.RunSpecific(ctx, t, supportedFixtures, testFunc)
}

func clientGet(t *testing.T, ctx context.Context, client gcpclient.Client, u *unstructured.Unstructured) *unstructured.Unstructured {
	t.Helper()
	u, err := client.Get(ctx, u)
	if err != nil {
		t.Fatalf("error getting unstructured %s object %s: %v", u.GroupVersionKind().Kind, k8s.GetNamespacedName(u), err)
	}
	return u
}

func clientApply(t *testing.T, client gcpclient.Client, u *unstructured.Unstructured) *unstructured.Unstructured {
	t.Helper()
	newUnstruct, err := client.Apply(u)
	if err != nil {
		t.Fatalf("error applying %s object %s: %v", u.GroupVersionKind().Kind, k8s.GetNamespacedName(u), err)
	}
	t.Logf("applied %s object %s", u.GroupVersionKind().Kind, k8s.GetNamespacedName(u))
	return newUnstruct
}

func clientDelete(t *testing.T, client gcpclient.Client, u *unstructured.Unstructured) {
	t.Helper()
	if !shouldDelete(u) {
		return
	}
	if err := client.Delete(u); err != nil {
		t.Fatalf("error deleting %s object %s: %v", u.GroupVersionKind().Kind, k8s.GetNamespacedName(u), err)
	}
	t.Logf("deleted %s object %s", u.GroupVersionKind().Kind, k8s.GetNamespacedName(u))
}

// the intention is for the returned delete function to be deferred as a catch all for resource leaks when a test fails
func buildDeleteFunc(t *testing.T, client gcpclient.Client, u *unstructured.Unstructured) func() {
	if !shouldDelete(u) {
		return func() {}
	}
	return func() {
		if err := client.Delete(u); err != nil {
			// do not Fatal so that other delete functions will still run
			t.Errorf("error deleting %s object %s: %v", u.GroupVersionKind().Kind, k8s.GetNamespacedName(u), err)
		}
	}
}

func shouldDelete(u *unstructured.Unstructured) bool {
	return testgcp.ResourceSupportsDeletion(u.GetKind()) && !k8s.HasAbandonAnnotation(u)
}

func getSupportedResourcesSetCover(t *testing.T, gcpClient gcpclient.Client) []resourcefixture.ResourceFixture {
	lightFilter := func(name string, testType resourcefixture.TestType) bool {
		if test.StringMatchesRegexList(t, testconstants.TestNameRegexToSkipForTestCRUD, name) {
			return false
		}
		return testType == resourcefixture.Basic
	}
	heavyFilter := func(fixture resourcefixture.ResourceFixture) bool {
		return allKindsAreSupported(t, gcpClient, fixture)
	}
	return resourcefixture.GetFilteredSetCover(t, lightFilter, heavyFilter)
}

func allKindsAreSupported(t *testing.T, client gcpclient.Client, fixture resourcefixture.ResourceFixture) bool {
	u := test.ToUnstruct(t, fixture.Create)
	if !client.IsSupported(u.GroupVersionKind().Kind) {
		return false
	}
	for _, d := range testyaml.SplitYAML(t, fixture.Dependencies) {
		u := test.ToUnstruct(t, d)
		if !client.IsSupported(u.GroupVersionKind().Kind) {
			return false
		}
	}
	return true
}

func applyProjectRefOrAnnotation(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, u *unstructured.Unstructured, projectID string) {
	rc, err := smLoader.GetResourceConfig(u)
	if err != nil {
		t.Fatalf("error getting ResourceConfig: %v", err)
	}
	switch {
	case krmtotf.SupportsHierarchicalReferences(rc):
		applyProjectRef(t, rc, u, projectID)
	case len(rc.Containers) > 0:
		applyProjectAnnotation(t, rc, u, projectID)
	}
}

func applyProjectRef(t *testing.T, rc *corekccv1alpha1.ResourceConfig, u *unstructured.Unstructured, projectID string) {
	r, err := k8s.NewResource(u)
	if err != nil {
		t.Fatalf("error creating resource from unstructured: %v", err)
	}

	// If the resource already has a hierarchical reference, no modification is required
	ref, _, err := k8s.GetHierarchicalReference(r, rc.HierarchicalReferences)
	if err != nil {
		t.Fatalf("error getting hierarchical reference from resource: %v", err)
	}
	if ref != nil {
		return
	}

	h := k8s.HierarchicalReferenceWithType(rc.HierarchicalReferences, corekccv1alpha1.HierarchicalReferenceTypeProject)
	if h == nil {
		return
	}
	if err := k8s.SetHierarchicalReference(r, h, projectID); err != nil {
		t.Fatalf("error setting hierarchical reference on resource: %v", err)
	}

	unstruct, err := r.MarshalAsUnstructured()
	if err != nil {
		t.Fatalf("error marshalling resource to unstructured: %v", err)
	}
	*u = *unstruct
}

func applyProjectAnnotation(t *testing.T, rc *corekccv1alpha1.ResourceConfig, u *unstructured.Unstructured, projectID string) {
	// If the resource already has a container annotation, no modification is required
	val, _, err := k8s.GetContainerAnnotation(u.GetAnnotations(), k8s.ContainerTypes(rc.Containers))
	if err != nil {
		t.Fatalf("error getting container annotation from object: %v", err)
	}
	if val != "" {
		return
	}
	if k8s.IsProjectScoped(rc.Containers) {
		k8s.SetAnnotation(k8s.ProjectIDAnnotation, projectID, u)
	}
}
