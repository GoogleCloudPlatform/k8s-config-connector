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
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testrunner "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/runner"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type ResourceLevelTestFunc func(ctx context.Context, t *testing.T, testID string, mgr manager.Manager, rc IAMResourceContext, refResource *unstructured.Unstructured, resourceRef v1beta1.ResourceReference)

// Runs a resource level test against all resources.
// testFunc will be executed once for each resource that supports IAMPolicy.
// shouldRunFunc is optional, it can be supplied to skip tests for resources that don't support a given operation
// (e.g. deleting the IAMPolicy on a storage bucket)
func RunResourceLevelTest(ctx context.Context, t *testing.T, mgr manager.Manager, iamTestFunc ResourceLevelTestFunc, shouldRunFunc resourcefixture.ShouldRunFunc) {
	t.Parallel()
	kindToIamPolicyResourceContext := getKindToIamPolicyResourceContextMap()
	testFunc := buildTestFunc(kindToIamPolicyResourceContext, iamTestFunc)
	shouldRun := buildShouldRunFunc(kindToIamPolicyResourceContext, shouldRunFunc)
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testFunc)
}

// Runs a resource level test against all resources, but creates an external
// resource reference to point to the referenced resource instead of a regular
// resource reference.
func RunResourceLevelTestWithExternalRef(ctx context.Context, t *testing.T, mgr manager.Manager, iamTestFunc ResourceLevelTestFunc, shouldRunFunc resourcefixture.ShouldRunFunc) {
	t.Parallel()
	kindToIamPolicyResourceContext := getKindToIamPolicyResourceContextMap()
	testFunc := buildTestFuncWithExternalRef(kindToIamPolicyResourceContext, iamTestFunc)
	shouldRun := buildShouldRunFunc(kindToIamPolicyResourceContext, shouldRunFunc)
	testrunner.RunAllWithObjectCreated(ctx, t, mgr, shouldRun, testFunc)
}

func buildTestFunc(kindToIamPolicyResourceContext map[string]IAMResourceContext, testFunc ResourceLevelTestFunc) testrunner.TestCaseFunc {
	return func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		rc := kindToIamPolicyResourceContext[testContext.ResourceFixture.GVK.Kind]
		refResource := testContext.CreateUnstruct
		resourceRef := NewResourceRef(refResource)
		testFunc(ctx, t, testContext.UniqueID, sysContext.Manager, rc, refResource, resourceRef)
	}
}

func buildTestFuncWithExternalRef(kindToIamPolicyResourceContext map[string]IAMResourceContext, testFunc ResourceLevelTestFunc) testrunner.TestCaseFunc {
	return func(ctx context.Context, t *testing.T, testContext testrunner.TestContext, sysContext testrunner.SystemContext) {
		rc := kindToIamPolicyResourceContext[testContext.ResourceFixture.GVK.Kind]
		refResource := testContext.CreateUnstruct
		resourceRef, err := NewExternalRef(refResource, sysContext.TFProvider, sysContext.SMLoader)
		if err != nil {
			t.Fatal(err)
		}
		testFunc(ctx, t, testContext.UniqueID, sysContext.Manager, rc, refResource, resourceRef)
	}
}

func buildShouldRunFunc(kindToIamPolicyResourceContext map[string]IAMResourceContext, additionalShouldRunFunc resourcefixture.ShouldRunFunc) testrunner.ShouldRunFunc {
	return func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool {
		if additionalShouldRunFunc != nil {
			if !additionalShouldRunFunc(fixture) {
				return false
			}
		}

		if fixture.Type != resourcefixture.Basic {
			return false
		}

		rc, ok := kindToIamPolicyResourceContext[fixture.GVK.Kind]
		if !ok {
			return false
		}
		// If there is no specific test case name defined, run all the test cases for the kind.
		if rc.Name == "" {
			return true
		}
		return rc.Name == fixture.Name
	}
}

func getKindToIamPolicyResourceContextMap() map[string]IAMResourceContext {
	results := make(map[string]IAMResourceContext)
	for _, r := range ResourceContexts {
		results[r.Kind] = r
	}
	return results
}
