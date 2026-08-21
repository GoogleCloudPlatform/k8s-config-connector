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

package testrunner

import (
	"context"
	"net/http"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	dclconversion "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type TestContext struct {
	CreateUnstruct      *unstructured.Unstructured
	UpdateUnstruct      *unstructured.Unstructured
	DependencyUnstructs []*unstructured.Unstructured
	ResourceFixture     resourcefixture.ResourceFixture
	NamespacedName      types.NamespacedName
	UniqueID            string
}

type SystemContext struct {
	Manager      manager.Manager
	SMLoader     *servicemappingloader.ServiceMappingLoader
	Reconciler   *testreconciler.TestReconciler
	TFProvider   *schema.Provider
	DCLConfig    *mmdcl.Config
	DCLConverter *dclconversion.Converter
	HttpClient   *http.Client
}

type ShouldRunFunc func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool
type TestCaseFunc func(ctx context.Context, t *testing.T, testContext TestContext, sysContext SystemContext)

func RunAllWithObjectCreated(ctx context.Context, t *testing.T, mgr manager.Manager, shouldRunFunc ShouldRunFunc, testCaseFunc TestCaseFunc) {
	testFunc := func(ctx context.Context, t *testing.T, testContext TestContext, sysContext SystemContext) {
		if err := sysContext.Manager.GetClient().Create(ctx, testContext.CreateUnstruct); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
		resourceCleanup := sysContext.Reconciler.BuildCleanupFunc(ctx, testContext.CreateUnstruct, testreconciler.CleanupPolicyAlways)
		defer resourceCleanup()
		sysContext.Reconciler.Reconcile(ctx, testContext.CreateUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(sysContext.Reconciler, testContext.CreateUnstruct), nil)
		testCaseFunc(ctx, t, testContext, sysContext)
	}
	RunAllWithDependenciesCreatedButNotObject(ctx, t, mgr, shouldRunFunc, testFunc)
}

func RunAllWithDependenciesCreatedButNotObject(ctx context.Context, t *testing.T, mgr manager.Manager, shouldRunFunc ShouldRunFunc, testCaseFunc TestCaseFunc) {
	testFunc := func(ctx context.Context, t *testing.T, testContext TestContext, sysContext SystemContext) {
		cleanupFuncs := make([]func(), 0, len(testContext.DependencyUnstructs))
		for _, u := range testContext.DependencyUnstructs {
			if err := sysContext.Manager.GetClient().Create(ctx, u); err != nil {
				t.Fatalf("error creating dependecy '%v' for resource '%v/%v': %v", u.GetKind(), testContext.CreateUnstruct.GetName(), testContext.CreateUnstruct.GetKind(), err)
			}
			cleanupFuncs = append(cleanupFuncs, sysContext.Reconciler.BuildCleanupFunc(ctx, u, testreconciler.CleanupPolicyAlways))
			sysContext.Reconciler.Reconcile(ctx, u, testreconciler.ExpectedSuccessfulReconcileResultFor(sysContext.Reconciler, u), nil)
		}
		dependencyCleanup := func() {
			for i := len(cleanupFuncs) - 1; i >= 0; i-- {
				cleanupFuncs[i]()
			}
		}
		defer dependencyCleanup()
		testCaseFunc(ctx, t, testContext, sysContext)
	}
	RunAll(ctx, t, mgr, shouldRunFunc, testFunc)
}

func RunAll(ctx context.Context, t *testing.T, mgr manager.Manager, shouldRunFunc ShouldRunFunc, testCaseFunc TestCaseFunc) {
	shouldRun := func(resourceFixture resourcefixture.ResourceFixture) bool {
		return shouldRunFunc(resourceFixture, mgr)
	}
	testFunc := func(ctx context.Context, t *testing.T, fixture resourcefixture.ResourceFixture) {
		project := testgcp.GetDefaultProject(t)
		systemContext := newSystemContext(ctx, t, mgr)
		testContext := NewTestContext(t, fixture, project)
		setupNamespaces(t, testContext, systemContext, project)
		testCaseFunc(ctx, t, testContext, systemContext)
	}
	resourcefixture.RunTests(ctx, t, shouldRun, testFunc)
}

func RunSpecific(ctx context.Context, t *testing.T, fixture []resourcefixture.ResourceFixture, testCaseFunc func(ctx context.Context, t *testing.T, testContext TestContext)) {
	testFunc := func(ctx context.Context, t *testing.T, fixture resourcefixture.ResourceFixture) {
		project := testgcp.GetDefaultProject(t)
		testContext := NewTestContext(t, fixture, project)
		testCaseFunc(ctx, t, testContext)
	}
	resourcefixture.RunSpecificTests(ctx, t, fixture, testFunc)
}

// NewTestContext takes a resource fixture and returns a filled out TestContext
// The resources in the fixture are converted to unstructured.Unstructured and their namespaces are set equal to a
// unique generated id.
func NewTestContext(t *testing.T, fixture resourcefixture.ResourceFixture, project testgcp.GCPProject) TestContext {
	testID := testvariable.NewUniqueID()
	initialUnstruct := bytesToUnstructured(t, fixture.Create, testID, project)
	name := k8s.GetNamespacedName(initialUnstruct)
	var updateUnstruct *unstructured.Unstructured
	if fixture.Update != nil {
		updateUnstruct = bytesToUnstructured(t, fixture.Update, testID, project)
	}
	var dependencyUnstructs []*unstructured.Unstructured
	if fixture.Dependencies != nil {
		dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
		dependencyUnstructs = make([]*unstructured.Unstructured, 0, len(dependencyYamls))
		for _, dependBytes := range dependencyYamls {
			depUnstruct := bytesToUnstructured(t, dependBytes, testID, project)
			dependencyUnstructs = append(dependencyUnstructs, depUnstruct)
		}
	}
	return TestContext{
		CreateUnstruct:      initialUnstruct,
		UpdateUnstruct:      updateUnstruct,
		DependencyUnstructs: dependencyUnstructs,
		ResourceFixture:     fixture,
		NamespacedName:      name,
		UniqueID:            testID,
	}
}

func bytesToUnstructured(t *testing.T, bytes []byte, testID string, project testgcp.GCPProject) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testID, project)
	return test.ToUnstructWithNamespace(t, updatedBytes, testID)
}

func newSystemContext(ctx context.Context, t *testing.T, mgr manager.Manager) SystemContext {
	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatalWithContext(ctx, tfprovider.DefaultConfig)
	dclConfig, httpClient := clientconfig.NewConfigAndClientForIntegrationTest()
	reconciler := testreconciler.NewTestReconciler(t, mgr, tfProvider, dclConfig, httpClient)
	serviceMetadataLoader := dclmetadata.New()
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		t.Fatalf("error creating a new DCL schema laoder: %v", dclSchemaLoader)
	}
	dclConverter := dclconversion.New(dclSchemaLoader, serviceMetadataLoader)
	return SystemContext{
		Manager:      mgr,
		Reconciler:   reconciler,
		SMLoader:     smLoader,
		TFProvider:   tfProvider,
		DCLConfig:    dclConfig,
		DCLConverter: dclConverter,
		HttpClient:   httpClient,
	}
}

func setupNamespaces(t *testing.T, testContext TestContext, systemContext SystemContext, project testgcp.GCPProject) {
	namespacesAlreadySetup := make(map[string]bool)
	testcontroller.SetupNamespaceForProject(t, systemContext.Manager.GetClient(), testContext.CreateUnstruct.GetNamespace(), project.ProjectID)
	namespacesAlreadySetup[testContext.CreateUnstruct.GetNamespace()] = true
	for _, depUnstruct := range testContext.DependencyUnstructs {
		if _, ok := namespacesAlreadySetup[depUnstruct.GetNamespace()]; ok {
			continue
		}
		testcontroller.SetupNamespaceForProject(t, systemContext.Manager.GetClient(), depUnstruct.GetNamespace(), project.ProjectID)
		namespacesAlreadySetup[depUnstruct.GetNamespace()] = true
	}
}
