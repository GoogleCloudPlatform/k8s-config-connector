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
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	dclconversion "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/testcontext"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	cloudresourcemanagerv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"sigs.k8s.io/controller-runtime/pkg/log"

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
		dependencyCleanup := sysContext.Reconciler.CreateAndReconcile(ctx, testContext.DependencyUnstructs, testreconciler.CleanupPolicyAlways)
		defer dependencyCleanup()
		testCaseFunc(ctx, t, testContext, sysContext)
	}
	RunAll(ctx, t, mgr, shouldRunFunc, testFunc)
}

func RunAll(ctx context.Context, t *testing.T, mgr manager.Manager, shouldRunFunc ShouldRunFunc, testCaseFunc TestCaseFunc) {
	var kccConfig kccmanager.Config

	var project testgcp.GCPProject
	// TODO: Try to centralize and deduplicate this logic somewhere?
	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" {
		t.Logf("creating mock gcp")

		mockCloud := mockgcp.NewMockRoundTripper(t, mgr.GetClient(), storage.NewInMemoryStorage())

		httpRoundTripper := http.RoundTripper(mockCloud)

		ctx = testcontext.WithHTTPRoundTripper(ctx, httpRoundTripper)

		kccConfig.HTTPClient = &http.Client{Transport: httpRoundTripper}
		kccConfig.GCPAccessToken = "dummytoken"

		// Pre-create project
		crm, err := cloudresourcemanagerv1.NewService(ctx, option.WithHTTPClient(kccConfig.HTTPClient))
		if err != nil {
			t.Fatalf("error building cloudresourcemanagerv1 client: %v", err)
		}
		req := &cloudresourcemanager.Project{
			ProjectId: "mock-project",
		}
		op, err := crm.Projects.Create(req).Context(ctx).Do()
		if err != nil {
			t.Fatalf("error creating project: %v", err)
		}
		if !op.Done {
			t.Fatalf("expected mock create project operation to be done immediately")
		}
		found, err := crm.Projects.Get(req.ProjectId).Context(ctx).Do()
		if err != nil {
			t.Fatalf("error reading created project: %v", err)
		}
		project = testgcp.GCPProject{
			ProjectID:     found.ProjectId,
			ProjectNumber: found.ProjectNumber,
		}
	} else if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "real" || targetGCP == "" {
		t.Logf("targeting real GCP")
		project = testgcp.GetDefaultProject(t)
	} else {
		t.Fatalf("E2E_GCP_TARGET=%q not supported", targetGCP)
	}

	shouldRun := func(resourceFixture resourcefixture.ResourceFixture) bool {
		return shouldRunFunc(resourceFixture, mgr)
	}
	testFunc := func(ctx context.Context, t *testing.T, fixture resourcefixture.ResourceFixture) {
		systemContext := newSystemContext(ctx, t, mgr, kccConfig)
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
	testId := testvariable.NewUniqueId()
	initialUnstruct := bytesToUnstructured(t, fixture.Create, testId, project)
	name := k8s.GetNamespacedName(initialUnstruct)
	var updateUnstruct *unstructured.Unstructured
	if fixture.Update != nil {
		updateUnstruct = bytesToUnstructured(t, fixture.Update, testId, project)
	}
	var dependencyUnstructs []*unstructured.Unstructured
	if fixture.Dependencies != nil {
		dependencyYamls := testyaml.SplitYAML(t, fixture.Dependencies)
		dependencyUnstructs = make([]*unstructured.Unstructured, 0, len(dependencyYamls))
		for _, dependBytes := range dependencyYamls {
			depUnstruct := bytesToUnstructured(t, dependBytes, testId, project)
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

func bytesToUnstructured(t *testing.T, bytes []byte, testId string, project testgcp.GCPProject) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testId, project)
	return test.ToUnstructWithNamespace(t, updatedBytes, testId)
}

func newSystemContext(ctx context.Context, t *testing.T, mgr manager.Manager, kccConfig kccmanager.Config) SystemContext {
	smLoader := testservicemappingloader.New(t)

	// Allow for capture of http requests during a test.
	transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := testcontext.HTTPRoundTripperFromContext(ctx); t != nil {
			ret = &http.Client{Transport: t}
		}
		if artifacts := os.Getenv("ARTIFACTS"); artifacts == "" {
			log := log.FromContext(ctx)
			log.Info("env var ARTIFACTS is not set; will not record http log")
		} else {
			outputDir := filepath.Join(artifacts, "http-logs")
			t := test.NewHTTPRecorder(ret.Transport, outputDir)
			ret = &http.Client{Transport: t}
		}
		return ret
	}

	transport_tpg.OAuth2HTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := testcontext.HTTPRoundTripperFromContext(ctx); t != nil {
			ret = &http.Client{Transport: t}
		}
		if artifacts := os.Getenv("ARTIFACTS"); artifacts == "" {
			log := log.FromContext(ctx)
			log.Info("env var ARTIFACTS is not set; will not record http log")
		} else {
			outputDir := filepath.Join(artifacts, "http-logs")
			t := test.NewHTTPRecorder(ret.Transport, outputDir)
			ret = &http.Client{Transport: t}
		}
		return ret
	}

	// Bootstrap the Google Terraform provider
	tfCfg := tfprovider.NewConfig()
	tfCfg.UserProjectOverride = kccConfig.UserProjectOverride
	tfCfg.BillingProject = kccConfig.BillingProject
	tfCfg.GCPAccessToken = kccConfig.GCPAccessToken

	tfProvider, err := tfprovider.New(ctx, tfCfg)
	if err != nil {
		t.Fatalf("error creating TF provider: %v", err)
	}

	var dclConfig *mmdcl.Config
	{
		dclOptions := clientconfig.Options{
			UserProjectOverride: kccConfig.UserProjectOverride,
			BillingProject:      kccConfig.BillingProject,
			HTTPClient:          kccConfig.HTTPClient,
			UserAgent:           "kcc/dev",
		}

		// Log DCL requests
		if artifacts := os.Getenv("ARTIFACTS"); artifacts != "" {
			outputDir := filepath.Join(artifacts, "http-logs")
			if dclOptions.HTTPClient == nil {
				httpClient, err := google.DefaultClient(ctx, gcp.ClientScopes...)
				if err != nil {
					t.Fatalf("error creating the http client to be used by DCL: %v", err)
				}
				dclOptions.HTTPClient = httpClient
			}
			t := test.NewHTTPRecorder(dclOptions.HTTPClient.Transport, outputDir)
			dclOptions.HTTPClient = &http.Client{Transport: t}
		}

		dclClient, err := clientconfig.New(ctx, dclOptions)
		if err != nil {
			t.Fatalf("error from NewForIntegrationTest: %v", err)
		}
		dclConfig = dclClient
	}

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
