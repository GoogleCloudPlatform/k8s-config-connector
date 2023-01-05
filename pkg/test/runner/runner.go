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
	UniqueId            string
}

type SystemContext struct {
	Manager      manager.Manager
	SMLoader     *servicemappingloader.ServiceMappingLoader
	Reconciler   *testreconciler.TestReconciler
	TFProvider   *schema.Provider
	DCLConfig    *mmdcl.Config
	DCLConverter *dclconversion.Converter
}

type ShouldRunFunc func(fixture resourcefixture.ResourceFixture, mgr manager.Manager) bool
type TestCaseFunc func(t *testing.T, testContext TestContext, sysContext SystemContext)

func RunAllWithObjectCreated(t *testing.T, mgr manager.Manager, shouldRunFunc ShouldRunFunc, testCaseFunc TestCaseFunc) {
	testFunc := func(t *testing.T, testContext TestContext, sysContext SystemContext) {
		if err := sysContext.Manager.GetClient().Create(context.TODO(), testContext.CreateUnstruct); err != nil {
			t.Fatalf("error creating resource: %v", err)
		}
		resourceCleanup := sysContext.Reconciler.BuildCleanupFunc(testContext.CreateUnstruct, testreconciler.CleanupPolicyAlways)
		defer resourceCleanup()
		sysContext.Reconciler.Reconcile(testContext.CreateUnstruct, testreconciler.ExpectedSuccessfulReconcileResultFor(sysContext.Reconciler, testContext.CreateUnstruct.GroupVersionKind()), nil)
		testCaseFunc(t, testContext, sysContext)
	}
	RunAllWithDependenciesCreatedButNotObject(t, mgr, shouldRunFunc, testFunc)

}

func RunAllWithDependenciesCreatedButNotObject(t *testing.T, mgr manager.Manager, shouldRunFunc ShouldRunFunc, testCaseFunc TestCaseFunc) {
	testFunc := func(t *testing.T, testContext TestContext, sysContext SystemContext) {
		dependencyCleanup := sysContext.Reconciler.CreateAndReconcile(testContext.DependencyUnstructs, testreconciler.CleanupPolicyAlways)
		defer dependencyCleanup()
		testCaseFunc(t, testContext, sysContext)
	}
	RunAll(t, mgr, shouldRunFunc, testFunc)
}

func RunAll(t *testing.T, mgr manager.Manager, shouldRunFunc ShouldRunFunc, testCaseFunc TestCaseFunc) {
	projectId := testgcp.GetDefaultProjectID(t)
	shouldRun := func(resourceFixture resourcefixture.ResourceFixture) bool {
		return shouldRunFunc(resourceFixture, mgr)
	}
	testFunc := func(t *testing.T, fixture resourcefixture.ResourceFixture) {
		systemContext := newSystemContext(t, mgr)
		testContext := NewTestContext(t, fixture, projectId)
		setupNamespaces(t, testContext, systemContext)
		testCaseFunc(t, testContext, systemContext)
	}
	resourcefixture.RunTests(t, shouldRun, testFunc)
}

func RunSpecific(t *testing.T, fixture []resourcefixture.ResourceFixture, testCaseFunc func(t *testing.T, testContext TestContext)) {
	projectId := testgcp.GetDefaultProjectID(t)
	testFunc := func(t *testing.T, fixture resourcefixture.ResourceFixture) {
		testContext := NewTestContext(t, fixture, projectId)
		testCaseFunc(t, testContext)
	}
	resourcefixture.RunSpecificTests(t, fixture, testFunc)
}

// NewTestContext takes a resource fixture and returns a filled out TestContext
// The resources in the fixture are converted to unstructured.Unstructured and their namespaces are set equal to a
// unique generated id.
func NewTestContext(t *testing.T, fixture resourcefixture.ResourceFixture, projectId string) TestContext {
	testId := testvariable.NewUniqueId()
	initialUnstruct := bytesToUnstructured(t, fixture.Create, testId, projectId)
	name := k8s.GetNamespacedName(initialUnstruct)
	var updateUnstruct *unstructured.Unstructured
	if fixture.Update != nil {
		updateUnstruct = bytesToUnstructured(t, fixture.Update, testId, projectId)
	}
	var dependencyUnstructs []*unstructured.Unstructured
	if fixture.Dependencies != nil {
		dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
		dependencyUnstructs = make([]*unstructured.Unstructured, 0, len(dependencyYamls))
		for _, dependBytes := range dependencyYamls {
			depUnstruct := bytesToUnstructured(t, dependBytes, testId, projectId)
			dependencyUnstructs = append(dependencyUnstructs, depUnstruct)
		}
	}
	return TestContext{
		CreateUnstruct:      initialUnstruct,
		UpdateUnstruct:      updateUnstruct,
		DependencyUnstructs: dependencyUnstructs,
		ResourceFixture:     fixture,
		NamespacedName:      name,
		UniqueId:            testId,
	}
}

func bytesToUnstructured(t *testing.T, bytes []byte, testId string, projectId string) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testId, projectId)
	return test.ToUnstructWithNamespace(t, updatedBytes, testId)
}

func newSystemContext(t *testing.T, mgr manager.Manager) SystemContext {
	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	dclConfig := clientconfig.NewForIntegrationTest()
	reconciler := testreconciler.NewForDCLAndTFTestReconciler(t, mgr, tfProvider, dclConfig)
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
	}
}

func setupNamespaces(t *testing.T, testContext TestContext, systemContext SystemContext) {
	namespacesAlreadySetup := make(map[string]bool)
	testcontroller.SetupNamespaceForDefaultProject(t, systemContext.Manager.GetClient(), testContext.CreateUnstruct.GetNamespace())
	namespacesAlreadySetup[testContext.CreateUnstruct.GetNamespace()] = true
	for _, depUnstruct := range testContext.DependencyUnstructs {
		if _, ok := namespacesAlreadySetup[depUnstruct.GetNamespace()]; ok {
			continue
		}
		testcontroller.SetupNamespaceForDefaultProject(t, systemContext.Manager.GetClient(), depUnstruct.GetNamespace())
		namespacesAlreadySetup[depUnstruct.GetNamespace()] = true
	}
}
