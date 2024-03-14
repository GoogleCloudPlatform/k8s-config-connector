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

package create

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-logr/logr"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	cloudresourcemanagerv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	exportparameters "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testenvironment "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/environment"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/webhook"
	cnrmwebhook "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"
)

type Harness struct {
	*testing.T
	Ctx context.Context

	Events  *test.MemoryEventSink
	Project testgcp.GCPProject

	client     client.Client
	restConfig *rest.Config

	// gcpAccessToken is set to the oauth2 token to use for GCP, primarily when GCP is mocked.
	gcpAccessToken string
	kccConfig      kccmanager.Config
}

type httpRoundTripperKeyType int

// httpRoundTripperKey is the key value for http.RoundTripper in a context.Context
var httpRoundTripperKey httpRoundTripperKeyType

// NewHarnessWithManager builds a Harness for an existing manager.
// deprecated: Prefer NewHarness, which can construct a manager and mock gcp etc.
func NewHarnessWithManager(ctx context.Context, t *testing.T, mgr manager.Manager) *Harness {
	h := &Harness{
		T:      t,
		Ctx:    ctx,
		client: mgr.GetClient(),
	}
	return h
}

func NewHarness(ctx context.Context, t *testing.T) *Harness {
	ctx, ctxCancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		ctxCancel()
	})

	log := log.FromContext(ctx)

	h := &Harness{
		T:   t,
		Ctx: ctx,
	}

	kccConfig := kccmanager.Config{}
	// Prevent manager from binding to a port to serve prometheus metrics
	// since creating multiple managers for tests will fail if more than
	// one manager tries to bind to the same port.
	kccConfig.ManagerOptions.MetricsBindAddress = "0"
	// Prevent manager from binding to a port to serve health probes since
	// creating multiple managers for tests will fail if more than one
	// manager tries to bind to the same port.
	kccConfig.ManagerOptions.HealthProbeBindAddress = "0"
	// supply a concrete client to disable the default behavior of caching
	kccConfig.ManagerOptions.NewClient = nocache.NoCacheClientFunc
	kccConfig.StateIntoSpecDefaultValue = k8s.StateIntoSpecDefaultValueV1Beta1

	var webhooks []cnrmwebhook.Config

	loadCRDs := true
	if targetKube := os.Getenv("E2E_KUBE_TARGET"); targetKube == "envtest" {
		whCfgs, err := testwebhook.GetTestCommonWebhookConfigs()
		if err != nil {
			h.Fatalf("error getting common wehbook configs: %v", err)
		}
		webhooks = append(webhooks, whCfgs...)

		env := &envtest.Environment{
			ControlPlaneStartTimeout: time.Minute,
			ControlPlaneStopTimeout:  time.Minute,
		}

		testenvironment.ConfigureWebhookInstallOptions(env, whCfgs)

		h.Logf("starting envtest apiserver")
		restConfig, err := env.Start()
		if err != nil {
			h.Fatalf("error starting test environment: %v", err)
		}

		t.Cleanup(func() {
			if err := env.Stop(); err != nil {
				h.Errorf("error stopping envtest environment: %v", err)
			}
		})

		h.restConfig = restConfig

		kccConfig.ManagerOptions.Port = env.WebhookInstallOptions.LocalServingPort
		kccConfig.ManagerOptions.Host = env.WebhookInstallOptions.LocalServingHost
		kccConfig.ManagerOptions.CertDir = env.WebhookInstallOptions.LocalServingCertDir
	} else if targetKube := os.Getenv("E2E_KUBE_TARGET"); targetKube == "mock" {
		k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
		if err != nil {
			h.Fatalf("error building mock kube-apiserver: %v", err)
		}

		addr, err := k8s.StartServing()
		if err != nil {
			h.Fatalf("error starting mock kube-apiserver: %v", err)
		}

		t.Cleanup(func() {
			if err := k8s.Stop(); err != nil {
				h.Errorf("error stopping mock kube-apiserver: %v", err)
			}
		})

		h.restConfig = &rest.Config{
			Host: addr.String(),
			ContentConfig: rest.ContentConfig{
				ContentType: "application/json",
			},
			// gotta go fast during tests -- we don't really care about overwhelming our test API server
			QPS:   1000.0,
			Burst: 2000.0,
		}

	} else {
		t.Fatalf("E2E_KUBE_TARGET=%q not supported", targetKube)
	}

	if h.client == nil {
		client, err := client.New(h.restConfig, client.Options{})
		if err != nil {
			h.Fatalf("error building client: %v", err)
		}
		h.client = client
	}

	logging.SetupLogger()

	if loadCRDs {
		crds, err := crdloader.LoadAllCRDs()
		if err != nil {
			h.Fatalf("error loading crds: %v", err)
		}
		{
			var wg sync.WaitGroup
			for i := range crds {
				crd := &crds[i]
				wg.Add(1)
				log.V(2).Info("loading crd", "name", crd.GetName())
				go func() {
					defer wg.Done()
					if err := h.client.Create(ctx, crd.DeepCopy()); err != nil {
						h.Fatalf("error creating crd %v: %v", crd.GroupVersionKind(), err)
					}
					h.waitForCRDReady(crd)
				}()
			}
			wg.Wait()
		}
	}

	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" {
		t.Logf("creating mock gcp")

		mockCloud := mockgcp.NewMockRoundTripper(t, h.client, storage.NewInMemoryStorage())

		roundTripper := http.RoundTripper(mockCloud)

		ctx = context.WithValue(ctx, httpRoundTripperKey, roundTripper)
		h.Ctx = ctx

		kccConfig.HTTPClient = &http.Client{Transport: roundTripper}

		// Also hook the oauth2 library
		ctx = context.WithValue(ctx, oauth2.HTTPClient, kccConfig.HTTPClient)
		h.Ctx = ctx

		h.gcpAccessToken = "dummytoken"
		kccConfig.GCPAccessToken = h.gcpAccessToken
	} else if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "real" {
		t.Logf("targeting real GCP")
	} else {
		t.Fatalf("E2E_GCP_TARGET=%q not supported", targetGCP)
	}

	if os.Getenv("E2E_GCP_TARGET") == "mock" {
		// Some fixed-value fake org-ids for testing.
		// We used fixed values so that the output is predictable (for golden testing)
		testgcp.TestFolderID.Set("123451001")
		testgcp.TestFolder2ID.Set("123451002")
		testgcp.TestOrgID.Set("123450001")
		testgcp.TestBillingAccountID.Set("123456-777777-000001")
		testgcp.IAMIntegrationTestsOrganizationID.Set("123450002")
		testgcp.IAMIntegrationTestsBillingAccountID.Set("123456-777777-000002")
		testgcp.TestAttachedClusterName.Set("xks-cluster")

		crm := h.getCloudResourceManagerClient(kccConfig.HTTPClient)
		req := &cloudresourcemanagerv1.Project{
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
		project := testgcp.GCPProject{
			ProjectID:     found.ProjectId,
			ProjectNumber: found.ProjectNumber,
		}
		testgcp.TestKCCAttachedClusterProject.Set("mock-project")
		h.Project = project
	} else {
		h.Project = testgcp.GetDefaultProject(t)
	}

	eventSink := test.NewMemoryEventSink()
	ctx = test.AddSinkToContext(ctx, eventSink)
	h.Ctx = ctx

	h.Events = eventSink

	eventSinks := test.EventSinksFromContext(ctx)

	// Set up event sink for logging to a file, if ARTIFACTS env var is set
	if artifacts := os.Getenv("ARTIFACTS"); artifacts != "" {
		outputDir := filepath.Join(artifacts, "http-logs")
		eventSinks = append(eventSinks, test.NewDirectoryEventSink(outputDir))
	} else {
		log.Info("env var ARTIFACTS is not set; will not record http log")
	}

	// Intercept (and log) DCL requests
	if len(eventSinks) != 0 {
		if kccConfig.HTTPClient == nil {
			httpClient, err := google.DefaultClient(ctx, gcp.ClientScopes...)
			if err != nil {
				t.Fatalf("error creating the http client to be used by DCL: %v", err)
			}
			kccConfig.HTTPClient = httpClient
		}
		t := test.NewHTTPRecorder(kccConfig.HTTPClient.Transport, eventSinks...)
		kccConfig.HTTPClient = &http.Client{Transport: t}
	}

	// Intercept (and log) TF requests
	transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := ctx.Value(httpRoundTripperKey); t != nil {
			ret = &http.Client{Transport: t.(http.RoundTripper)}
		}
		if len(eventSinks) != 0 {
			t := test.NewHTTPRecorder(ret.Transport, eventSinks...)
			ret = &http.Client{Transport: t}
		}
		return ret
	}

	// Intercept (and log) TF oauth requests
	transport_tpg.OAuth2HTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		ret := inner
		if t := ctx.Value(httpRoundTripperKey); t != nil {
			ret = &http.Client{Transport: t.(http.RoundTripper)}
		}
		if len(eventSinks) != 0 {
			t := test.NewHTTPRecorder(ret.Transport, eventSinks...)
			ret = &http.Client{Transport: t}
		}
		return ret
	}

	h.kccConfig = kccConfig
	// We must cancel the manager Context before cancelling the envtest Context
	// Create a context specifically for this, and register the test cleanup function
	// after the envtest cleanup function (these run last-in, first-out).
	// See https://github.com/kubernetes-sigs/controller-runtime/issues/1571#issuecomment-945535598
	var ctrlManagerShutdown sync.WaitGroup
	mgrContext, mgrContextCancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		mgrContextCancel()
		// Wait for the manager to exit, cancel doesn't wait for the exist.
		// Otherwise the manager may still connect to kube-apiserver
		// during its shutdown, blocking the shutdown of kube-apiserver.
		t.Log("waiting for controller-runtime manager shutdown")
		ctrlManagerShutdown.Wait()
		t.Log("controller-runtime manager is shutdown")
	})
	kccConfig.ManagerOptions.Logger = filterLogs(log)

	mgr, err := kccmanager.New(mgrContext, h.restConfig, kccConfig)
	if err != nil {
		t.Fatalf("error creating new manager: %v", err)
	}
	if len(webhooks) > 0 {
		server := mgr.GetWebhookServer()
		for _, cfg := range webhooks {
			handler := cfg.HandlerFunc(mgr)
			server.Register(cfg.Path, &webhook.Admission{Handler: handler})
		}
	}

	// Register the deletion defender controller.
	if err := registration.Add(mgr, nil, nil, nil, nil, registration.RegisterDeletionDefenderController, nil); err != nil {
		t.Fatalf("error adding registration controller for deletion defender controllers: %v", err)
	}
	// Start the manager, Start(...) is a blocking operation so it needs to be done asynchronously.
	ctrlManagerShutdown.Add(1)
	go func() {
		defer ctrlManagerShutdown.Done()

		err := mgr.Start(mgrContext)
		if err != nil {
			t.Errorf("error from mgr.Start: %v", err)
		}
	}()

	return h
}

// ExportParams returns the default parameters.Parameters to use for an export
func (h *Harness) ExportParams() exportparameters.Parameters {
	var exportParams exportparameters.Parameters
	exportParams.GCPAccessToken = h.gcpAccessToken
	return exportParams
}

func (h *Harness) getCloudResourceManagerClient(httpClient *http.Client) *cloudresourcemanagerv1.Service {
	s, err := cloudresourcemanagerv1.NewService(h.Ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		h.Fatalf("error building cloudresourcemanagerv1 client: %v", err)
	}
	return s
}

func (h *Harness) GetClient() client.Client {
	return h.client
}

func MaybeSkip(t *testing.T, name string, resources []*unstructured.Unstructured) {
	if os.Getenv("E2E_GCP_TARGET") == "mock" {
		for _, resource := range resources {
			gvk := resource.GroupVersionKind()

			switch gvk.Group {
			case "certificatemanager.cnrm.cloud.google.com":
				continue
			}

			switch gvk.GroupKind() {
			case schema.GroupKind{Group: "apikeys.cnrm.cloud.google.com", Kind: "APIKeysKey"}:

			case schema.GroupKind{Group: "cloudfunctions.cnrm.cloud.google.com", Kind: "CloudFunctionsFunction"}:

			case schema.GroupKind{Group: "containerattached.cnrm.cloud.google.com", Kind: "ContainerAttachedCluster"}:

			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeAddress"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeDisk"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetwork"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNodeGroup"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNodeTemplate"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSubnetwork"}:
				// ok

			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPartialPolicy"}:
			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicy"}:
			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicyMember"}:
			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMServiceAccount"}:

			case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerCluster"}:
			case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerNodePool"}:

			case schema.GroupKind{Group: "edgenetwork.cnrm.cloud.google.com", Kind: "EdgeNetworkNetwork"}:
			case schema.GroupKind{Group: "edgenetwork.cnrm.cloud.google.com", Kind: "EdgeNetworkSubnet"}:

			case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesMesh"}:

			case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACAPool"}:

			case schema.GroupKind{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "Project"}:
				// ok

			case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecret"}:
			case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecretVersion"}:

			case schema.GroupKind{Group: "", Kind: "Secret"}:

			case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:
				if os.Getenv("GOLDEN_REQUEST_CHECKS") != "" {
					// no golden log for this resource  yet
					t.Skipf("gk %v/%v does not support golden request check; skipping", gvk.GroupKind(), name)
				}
			case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "ServiceIdentity"}:

			case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageBucket"}:

			case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagKey"}:

			case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAITensorboard"}:

			default:
				t.Skipf("gk %v not suppported by mock gcp %v; skipping", gvk.GroupKind(), name)
			}
		}
	}
}

func (h *Harness) waitForCRDReady(obj client.Object) {
	logger := log.FromContext(h.Ctx)

	apiVersion, kind := obj.GetObjectKind().GroupVersionKind().ToAPIVersionAndKind()
	name := obj.GetName()
	namespace := obj.GetNamespace()

	id := types.NamespacedName{Name: name, Namespace: namespace}
	if err := wait.PollImmediate(2*time.Second, 1*time.Minute, func() (bool, error) {
		u := &unstructured.Unstructured{}
		u.SetAPIVersion(apiVersion)
		u.SetKind(kind)
		logger.V(2).Info("Testing to see if resource is ready", "kind", kind, "id", id)
		if err := h.GetClient().Get(h.Ctx, id, u); err != nil {
			logger.Info("Error getting resource", "kind", kind, "id", id, "error", err)
			return false, err
		}
		objectStatus := dynamic.GetObjectStatus(h.T, u)
		// CRDs do not have observedGeneration
		for _, condition := range objectStatus.Conditions {
			if condition.Type == "Established" && condition.Status == "True" {
				logger.V(2).Info("crd is ready", "kind", kind, "id", id)
				return true, nil
			}
		}
		// This resource is not completely ready. Let's keep polling.
		logger.V(2).Info("CRD is not ready", "kind", kind, "id", id, "conditions", objectStatus.Conditions)
		return false, nil
	}); err != nil {
		h.Errorf("error while polling for ready on %v %v: %v", kind, id, err)
		return
	}
}

func (h *Harness) CompareGoldenFile(p string, got string, normalizers ...func(s string) string) {
	test.CompareGoldenFile(h.T, p, got, normalizers...)
}

func (h *Harness) MustReadFile(p string) []byte {
	return test.MustReadFile(h.T, p)
}

// IgnoreComments is a normalization function that strips comments.
func (h *Harness) IgnoreComments(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			lines[i] = ""
		}
	}
	s = strings.Join(lines, "\n")
	return strings.TrimSpace(s)
}

// ReplaceString is a normalization function that replaces a string, useful for e.g. project IDs.
func (h *Harness) ReplaceString(from, to string) func(string) string {
	return func(s string) string {
		return strings.ReplaceAll(s, from, to)
	}
}

func filterLogs(log logr.Logger) logr.Logger {
	f := &filterSink{sink: log.GetSink()}
	f.IgnoreMessages = sets.New[string]()
	f.IgnoreMessages.Insert("Registered controller")
	f.IgnoreMessages.Insert("Registered deletion-defender controller")
	f.IgnoreMessages.Insert("Starting Controller")
	f.IgnoreMessages.Insert("Starting EventSource")
	f.IgnoreMessages.Insert("Starting workers")
	f.IgnoreMessages.Insert("Shutdown signal received, waiting for all workers to finish")
	f.IgnoreMessages.Insert("All workers finished")
	return log.WithSink(f)
}

type filterSink struct {
	IgnoreMessages sets.Set[string]
	sink           logr.LogSink
}

// Init implements logr.LogSink
func (s *filterSink) Init(info logr.RuntimeInfo) {
	s.sink.Init(info)
}

// Enabled implements logr.LogSink
func (s *filterSink) Enabled(level int) bool {
	return s.sink.Enabled(level)
}

// Info implements logr.LogSink
func (s *filterSink) Info(level int, msg string, args ...any) {
	if s.IgnoreMessages.Has(msg) {
		return
	}
	s.sink.Info(level, msg, args...)
}

// WithValues implements logr.LogSink
func (s *filterSink) WithValues(keysAndValues ...any) logr.LogSink {
	return &filterSink{IgnoreMessages: s.IgnoreMessages, sink: s.sink.WithValues(keysAndValues...)}
}

// WithName implements logr.LogSink
func (s *filterSink) WithName(name string) logr.LogSink {
	return &filterSink{IgnoreMessages: s.IgnoreMessages, sink: s.sink.WithName(name)}
}

// Error implements logr.LogSink
func (s *filterSink) Error(err error, msg string, args ...any) {
	s.sink.Error(err, msg, args...)
}
