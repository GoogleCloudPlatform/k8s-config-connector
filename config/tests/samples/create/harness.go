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
	"errors"
	"fmt"
	"io"
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
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	exportparameters "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/export/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dynamic"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/registration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
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

	registeredServices mockgcpregistry.Normalizer

	Events     *test.MemoryEventSink
	KubeEvents *test.MemoryEventSink

	Project testgcp.GCPProject

	VCRRecorderNonTF *recorder.Recorder
	VCRRecorderTF    *recorder.Recorder
	VCRRecorderOauth *recorder.Recorder

	client     client.Client
	restConfig *rest.Config

	// gcpAccessToken is set to the oauth2 token to use for GCP, primarily when GCP is mocked.
	gcpAccessToken string
	kccConfig      kccmanager.Config

	// goldenFiles tracks the golden files we checked, so we can look for "extra" golden files.
	goldenFiles []string

	// MockGCP holds our mockgcp instance, if we are running against mockgcp
	MockGCP mockgcp.Interface

	// some fields that can be set by options
	vcrPath    string
	filterCRDs func(gk schema.GroupKind) bool
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

type HarnessOption func(*Harness)

func FilterCRDs(filterCRDs func(gk schema.GroupKind) bool) HarnessOption {
	return func(h *Harness) {
		h.filterCRDs = filterCRDs
	}
}

func WithVCRPath(vcrPath string) HarnessOption {
	return func(h *Harness) {
		h.vcrPath = vcrPath
	}
}
func NewHarness(ctx context.Context, t *testing.T, opts ...HarnessOption) *Harness {
	ctx, ctxCancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		ctxCancel()
	})
	log := log.FromContext(ctx)

	h := &Harness{
		T:   t,
		Ctx: ctx,
	}

	for _, opt := range opts {
		opt(h)
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
	// configure caching
	nocache.OnlyCacheCCAndCCC(&kccConfig.ManagerOptions)

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

		if pprofPath := os.Getenv("KUBEAPISERVER_CAPTURE_PPROF"); pprofPath != "" {
			pprofDone := make(chan error)
			t.Cleanup(func() {
				err := <-pprofDone
				if err != nil {
					t.Errorf("pprof failed: %v", err)
				}
			})
			doPprof := func() error {
				url := env.ControlPlane.GetAPIServer().SecureServing.URL("https", "debug/pprof/profile")
				url.RawQuery = "seconds=30"
				t.Logf("profiling with url %v", url)
				httpClient, err := rest.HTTPClientFor(restConfig)
				if err != nil {
					return fmt.Errorf("building http client: %w", err)
				}
				req, err := http.NewRequest("GET", url.String(), nil)
				if err != nil {
					return fmt.Errorf("error building request: %w", err)
				}
				response, err := httpClient.Do(req)
				if err != nil {
					return fmt.Errorf("doing pprof request: %w", err)
				}
				defer response.Body.Close()

				if response.StatusCode != 200 {
					return fmt.Errorf("unexpected response from pprof: %v", response.Status)
				}
				b, err := io.ReadAll(response.Body)
				if err != nil {
					return fmt.Errorf("reading pprof response: %w", err)
				}
				pprofName := strings.ReplaceAll(t.Name(), "/", "-") + ".pprof"
				pprofPath := filepath.Join(pprofPath, pprofName)
				if err := os.WriteFile(pprofPath, b, 0644); err != nil {
					return fmt.Errorf("writing pprof file %q: %w", pprofPath, err)
				}
				return nil
			}
			go func() {
				err := doPprof()
				if err != nil {
					t.Logf("error from pprof: %v", err)
				}
				pprofDone <- err
			}()
		}
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

	// Set up eventSinks for logging GCP and kube requests
	eventSinks := test.EventSinksFromContext(ctx)

	// Set up event sink for logging to a file, if ARTIFACTS env var is set
	if artifacts := os.Getenv("ARTIFACTS"); artifacts != "" {
		outputDir := filepath.Join(artifacts, "http-logs")
		eventSinks = append(eventSinks, test.NewDirectoryEventSink(outputDir))
	} else {
		log.Info("env var ARTIFACTS is not set; will not record http log")
	}

	// Set up logging of k8s requests
	logKubeRequests := true
	if logKubeRequests {
		kubeEvents := test.NewMemoryEventSink()
		h.KubeEvents = kubeEvents

		// Don't log these to general events (for now)
		kubeEventSinks := append(eventSinks, kubeEvents)

		wrapTransport := func(rt http.RoundTripper) http.RoundTripper {
			t := test.NewHTTPRecorder(rt, kubeEventSinks...)
			return t
		}
		h.restConfig.Wrap(wrapTransport)
	}

	// Set up capture of GCP requests
	{
		eventSink := test.NewMemoryEventSink()
		ctx = test.AddSinkToContext(ctx, eventSink)
		eventSinks = append(eventSinks, eventSink)
		h.Ctx = ctx

		h.Events = eventSink
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
			var errsMutex sync.Mutex
			var errs []error

			for i := range crds {
				crd := &crds[i]
				if h.filterCRDs != nil {
					gk := k8s.GetGroupKindFromCRD(crd)
					if !h.filterCRDs(gk) {
						continue
					}
				}
				wg.Add(1)
				log.V(2).Info("loading crd", "name", crd.GetName())

				go func() {
					defer wg.Done()
					if err := h.client.Create(ctx, crd.DeepCopy()); err != nil {
						errsMutex.Lock()
						defer errsMutex.Unlock()
						errs = append(errs, fmt.Errorf("error creating crd %v: %w", crd.GroupVersionKind(), err))
					}
					h.waitForCRDReady(crd)
				}()
			}
			wg.Wait()
			if len(errs) != 0 {
				h.Fatalf("error creating crds: %v", errors.Join(errs...))
			}
		}
	}

	var mockCloudGRPCClientConnection *grpc.ClientConn
	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" {
		t.Logf("creating mock gcp")

		mockCloud := mockgcp.NewMockRoundTripperForTest(t, h.client, storage.NewInMemoryStorage())

		mockCloudGRPCClientConnection = mockCloud.NewGRPCConnection(ctx)
		h.MockGCP = mockCloud

		roundTripper := http.RoundTripper(mockCloud)

		ctx = context.WithValue(ctx, httpRoundTripperKey, roundTripper)
		h.Ctx = ctx

		kccConfig.HTTPClient = &http.Client{Transport: roundTripper}

		// Also hook the oauth2 library
		ctx = context.WithValue(ctx, oauth2.HTTPClient, kccConfig.HTTPClient)
		h.Ctx = ctx

		h.gcpAccessToken = "dummytoken"
		kccConfig.GCPAccessToken = h.gcpAccessToken

		h.registeredServices = mockCloud.(mockgcpregistry.Normalizer)
	} else if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "real" {
		t.Logf("targeting real GCP")

		// We create registered services, even though we only use it for replacements
		var kubeClient client.Client // TODO: We should replace this, it didn't work
		mockCloud := mockgcp.NewMockRoundTripperForTest(t, kubeClient, storage.NewInMemoryStorage())
		h.registeredServices = mockCloud.(mockgcpregistry.Normalizer)
	} else if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "vcr" {
		t.Logf("creating vcr test")
	} else {
		t.Fatalf("E2E_GCP_TARGET=%q not supported", targetGCP)
	}

	if os.Getenv("E2E_GCP_TARGET") == "mock" {
		// Some fixed-value fake org-ids for testing.
		// We used fixed values so that the output is predictable (for golden testing)
		testgcp.TestFolderID.Set("123451001")
		testgcp.TestFolder2ID.Set("123451002")
		testgcp.TestOrgID.Set("123450001")
		testgcp.IsolatedTestOrgName.Set("isolated-test-org.example.com")
		testgcp.TestBillingAccountID.Set("123456-777777-000001")
		testgcp.TestBillingAccountIDForBillingResources.Set("123456-777777-000003")
		testgcp.IAMIntegrationTestsOrganizationID.Set("123450002")
		testgcp.IAMIntegrationTestsBillingAccountID.Set("123456-777777-000002")
		testgcp.TestAttachedClusterName.Set("xks-cluster")
		testgcp.TestDependentNoNetworkProjectID.Set("mock-project")
		testgcp.TestDependentOrgProjectID.Set("example-project-01")
		testgcp.TestDependentFolderProjectID.Set("example-project-02")
		testgcp.FirestoreTestProject.Set("cnrm-test-firestore")
		testgcp.IdentityPlatformTestProject.Set("kcc-identity-platform")
		testgcp.RecaptchaEnterpriseTestProject.Set("kcc-recaptcha-enterprise")

		crm := h.getCloudResourceManagerClient(kccConfig.HTTPClient)
		req := &cloudresourcemanagerv1.Project{
			ProjectId: "mock-project",
		}
		op, err := crm.Projects.Create(req).Context(ctx).Do()
		if err != nil {
			t.Fatalf("error creating project: %v", err)
		}

		for i := 0; i < 10; i++ {
			if op.Done {
				break
			}
			time.Sleep(100 * time.Millisecond)
			latest, err := crm.Operations.Get(op.Name).Context(ctx).Do()
			if err != nil {
				t.Fatalf("error getting operation %q: %v", op.Name, err)
			}
			op = latest
		}
		if !op.Done {
			t.Fatalf("expected mock create project operation to be done")
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
		testgcp.TestKCCAttachedClusterPlatformVersion.Set("1.30.0-gke.1")
		h.Project = project
	} else if os.Getenv("E2E_GCP_TARGET") == "vcr" && os.Getenv("VCR_MODE") == "replay" {
		h.gcpAccessToken = "dummytoken"
		kccConfig.GCPAccessToken = h.gcpAccessToken

		h.Project = testgcp.GCPProject{
			ProjectID:     "example-project",
			ProjectNumber: 123456789,
		}
		testgcp.TestDependentOrgProjectID.Set("example-project-01")
		testgcp.TestDependentFolderProjectID.Set("example-project-02")
		testgcp.FirestoreTestProject.Set("cnrm-test-firestore")
		testgcp.IdentityPlatformTestProject.Set("kcc-identity-platform")
		testgcp.RecaptchaEnterpriseTestProject.Set("kcc-recaptcha-enterprise")
		testgcp.TestOrgID.Set("123450001")
		testgcp.TestBillingAccountID.Set("123456-777777-000001")
		testgcp.TestBillingAccountIDForBillingResources.Set("123456-777777-000003")
	} else {
		h.Project = testgcp.GetDefaultProject(t)
	}

	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "vcr" {
		// Initialize VCR recorder
		inputMode := os.Getenv("VCR_MODE")
		var vcrMode recorder.Mode
		if inputMode == "record" {
			vcrMode = recorder.ModeRecordOnly
		} else if inputMode == "replay" {
			vcrMode = recorder.ModeReplayOnly
		} else {
			t.Fatalf("[VCR] VCR_MODE should be set to record or replay; value %q is not known", inputMode)
		}
		path := filepath.Join(h.vcrPath, "_vcr_cassettes")
		// In replay mode, RealTransport is unnecessary because we simply replay existing cassettes.
		opts := &recorder.Options{
			CassetteName: filepath.Join(path, "nontf"),
			Mode:         vcrMode,
		}
		// In record mode, use the real GCP HTTP client's transport as the recorder's transport.
		// This way, the recorder is able to capture the real request/response pairs.
		if inputMode == "record" {
			// Intercept (and log) DCL and direct(non TF) requests
			if kccConfig.HTTPClient == nil {
				httpClient, err := google.DefaultClient(ctx, gcp.ClientScopes...)
				if err != nil {
					t.Fatalf("error creating the http client to be not used by TF: %v", err)
				}
				kccConfig.HTTPClient = httpClient
			}
			opts.RealTransport = kccConfig.HTTPClient.Transport
		}
		r, err := recorder.NewWithOptions(opts)
		if err != nil {
			t.Fatalf("[VCR] Failed create non TF vcr recorder: %v", err)
		}
		h.VCRRecorderNonTF = r
		kccConfig.HTTPClient = &http.Client{Transport: h.VCRRecorderNonTF}

		// Intercept (and log) TF requests
		transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
			ret := inner
			if t := ctx.Value(httpRoundTripperKey); t != nil {
				ret = &http.Client{Transport: t.(http.RoundTripper)}
			}
			opts := &recorder.Options{
				CassetteName:  filepath.Join(path, "tf"),
				Mode:          vcrMode,
				RealTransport: ret.Transport,
			}
			r, err := recorder.NewWithOptions(opts)
			if err != nil {
				t.Fatalf("[VCR] Failed create TF vcr recorder: %v", err)
			}
			h.VCRRecorderTF = r
			ret = &http.Client{Transport: h.VCRRecorderTF}
			return ret
		}
		// Intercept (and log) OAuth requests
		transport_tpg.OAuth2HTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
			ret := inner
			if t := ctx.Value(httpRoundTripperKey); t != nil {
				ret = &http.Client{Transport: t.(http.RoundTripper)}
			}
			opts := &recorder.Options{
				CassetteName:  filepath.Join(path, "oauth"),
				Mode:          vcrMode,
				RealTransport: ret.Transport,
			}
			r, err := recorder.NewWithOptions(opts)
			if err != nil {
				t.Fatalf("[VCR] Failed create Oauth vcr recorder: %v", err)
			}
			h.VCRRecorderOauth = r
			ret = &http.Client{Transport: h.VCRRecorderOauth}
			return ret
		}
	} else {
		// Intercept (and log) GRPC requests
		grpcUnaryInterceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			entry := &test.LogEntry{}

			entry.Request.URL = method
			entry.Request.Method = "GRPC"

			if req != nil {
				requestBytes, _ := protojson.Marshal(req.(proto.Message))
				entry.Request.Body = string(requestBytes)
			}

			if mockCloudGRPCClientConnection != nil {
				cc = mockCloudGRPCClientConnection
			}
			err := invoker(ctx, method, req, reply, cc, opts...)

			if reply != nil {
				replyBytes, _ := protojson.Marshal(reply.(proto.Message))
				entry.Response.Body = string(replyBytes)
			}

			if err != nil {
				entry.Response.Status = fmt.Sprintf("error: %v", err)
			} else {
				entry.Response.Status = "OK"
			}

			for _, eventSink := range eventSinks {
				eventSink.AddHTTPEvent(ctx, entry)
			}
			return err
		}

		transport_tpg.GRPCUnaryClientInterceptor = grpcUnaryInterceptor

		kccConfig.GRPCUnaryClientInterceptor = grpcUnaryInterceptor

		// Intercept (and log) DCL and direct(non TF) requests
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

	krmtotf.SetUserAgentForTerraformProvider()

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
	if err := registration.Add(mgr, &controller.Deps{}, registration.RegisterDeletionDefenderController); err != nil {
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

	// Wait for the webhook server to start (mgr.Start runs asynchronously)
	webhookWaitStart := time.Now()
	webhookTimeout := 10 * time.Second
	for {
		webhookStarted := mgr.GetWebhookServer().StartedChecker()
		req := &http.Request{}
		err := webhookStarted(req)
		if err == nil {
			break
		}
		if time.Since(webhookWaitStart) > webhookTimeout {
			t.Fatalf("webhook did not start within %v timeout", webhookTimeout)
		}
		t.Logf("waiting for webhook to start (%v)", err)
		time.Sleep(100 * time.Millisecond)
	}

	return h
}

func (h *Harness) RegisteredServices() mockgcpregistry.Normalizer {
	return h.registeredServices
}

// ExportParams returns the default parameters.Parameters to use for an export
func (h *Harness) ExportParams() exportparameters.Parameters {
	var exportParams exportparameters.Parameters
	exportParams.GCPAccessToken = h.gcpAccessToken
	exportParams.HTTPClient = h.kccConfig.HTTPClient
	return exportParams
}

func (h *Harness) getCloudResourceManagerClient(httpClient *http.Client) *cloudresourcemanagerv1.Service {
	s, err := cloudresourcemanagerv1.NewService(h.Ctx, option.WithHTTPClient(httpClient), option.WithUserAgent(gcp.KCCUserAgent()))
	if err != nil {
		h.Fatalf("error building cloudresourcemanagerv1 client: %v", err)
	}
	return s
}

func (h *Harness) GetClient() client.Client {
	return h.client
}

func (h *Harness) GetRESTConfig() *rest.Config {
	return h.restConfig
}

func MaybeSkip(t *testing.T, name string, resources []*unstructured.Unstructured) {
	if os.Getenv("E2E_GCP_TARGET") == "mock" {
		for _, resource := range resources {
			gvk := resource.GroupVersionKind()

			// Special fake types for testing
			if gvk.Group == "" && gvk.Kind == "RunCLI" {
				continue
			}
			if gvk.Group == "" && gvk.Kind == "MockGCPBackdoor" {
				continue
			}

			switch gvk.Group {
			case "core.cnrm.cloud.google.com":
				continue
			case "certificatemanager.cnrm.cloud.google.com":
				continue
			}

			switch gvk.GroupKind() {
			case schema.GroupKind{Group: "", Kind: "Secret"}:

			case schema.GroupKind{Group: "alloydb.cnrm.cloud.google.com", Kind: "AlloyDBCluster"}:
			case schema.GroupKind{Group: "alloydb.cnrm.cloud.google.com", Kind: "AlloyDBInstance"}:

			case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEndpointAttachment"}:
			case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEnvgroup"}:
			case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEnvgroupAttachment"}:
			case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeEnvironment"}:
			case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeInstance"}:
			case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeInstanceAttachment"}:
			case schema.GroupKind{Group: "apigee.cnrm.cloud.google.com", Kind: "ApigeeOrganization"}:

			case schema.GroupKind{Group: "apikeys.cnrm.cloud.google.com", Kind: "APIKeysKey"}:

			case schema.GroupKind{Group: "artifactregistry.cnrm.cloud.google.com", Kind: "ArtifactRegistryRepository"}:

			case schema.GroupKind{Group: "backupdr.cnrm.cloud.google.com", Kind: "BackupDRBackupPlan"}:

			case schema.GroupKind{Group: "backupdr.cnrm.cloud.google.com", Kind: "BackupDRBackupVault"}:

			case schema.GroupKind{Group: "backupdr.cnrm.cloud.google.com", Kind: "BackupDRManagementServer"}:

			case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryDataset"}:
			case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryTable"}:
			case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryRoutine"}:

			case schema.GroupKind{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Kind: "BigQueryAnalyticsHubDataExchange"}:
			case schema.GroupKind{Group: "bigqueryanalyticshub.cnrm.cloud.google.com", Kind: "BigQueryAnalyticsHubListing"}:

			case schema.GroupKind{Group: "bigqueryconnection.cnrm.cloud.google.com", Kind: "BigQueryConnectionConnection"}:

			case schema.GroupKind{Group: "bigquerydatatransfer.cnrm.cloud.google.com", Kind: "BigQueryDataTransferConfig"}:
			case schema.GroupKind{Group: "bigqueryreservation.cnrm.cloud.google.com", Kind: "BigQueryReservationReservation"}:

			case schema.GroupKind{Group: "gkehub.cnrm.cloud.google.com", Kind: "GKEHubFeature"}:
			case schema.GroupKind{Group: "gkehub.cnrm.cloud.google.com", Kind: "GKEHubMembership"}:
			case schema.GroupKind{Group: "gkehub.cnrm.cloud.google.com", Kind: "GKEHubFeatureMembership"}:

			case schema.GroupKind{Group: "cloudbuild.cnrm.cloud.google.com", Kind: "CloudBuildWorkerPool"}:

			case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableInstance"}:
			case schema.GroupKind{Group: "bigtable.cnrm.cloud.google.com", Kind: "BigtableTable"}:

			case schema.GroupKind{Group: "cloudfunctions.cnrm.cloud.google.com", Kind: "CloudFunctionsFunction"}:
			case schema.GroupKind{Group: "cloudids.cnrm.cloud.google.com", Kind: "CloudIDSEndpoint"}:

			case schema.GroupKind{Group: "cloudidentity.cnrm.cloud.google.com", Kind: "CloudIdentityGroup"}:
			case schema.GroupKind{Group: "cloudidentity.cnrm.cloud.google.com", Kind: "CloudIdentityMembership"}:

			case schema.GroupKind{Group: "containerattached.cnrm.cloud.google.com", Kind: "ContainerAttachedCluster"}:

			case schema.GroupKind{Group: "composer.cnrm.cloud.google.com", Kind: "ComposerEnvironment"}:

			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeAddress"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeBackendService"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeDisk"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeFirewallPolicy"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeFirewallPolicyRule"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeForwardingRule"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeHealthCheck"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeInstance"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeImage"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNetwork"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNodeGroup"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeNodeTemplate"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeManagedSSLCertificate"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeServiceAttachment"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSSLCertificate"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeSubnetwork"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetGRPCProxy"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetHTTPProxy"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetHTTPSProxy"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetVPNGateway"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeVPNGateway"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetHTTPProxy"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetSSLProxy"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeTargetTCPProxy"}:
			case schema.GroupKind{Group: "compute.cnrm.cloud.google.com", Kind: "ComputeURLMap"}:
				// ok

			case schema.GroupKind{Group: "container.cnrm.cloud.google.com", Kind: "ContainerCluster"}:
			case schema.GroupKind{Group: "container.cnrm.cloud.google.com", Kind: "ContainerNodePool"}:

			case schema.GroupKind{Group: "containeranalysis.cnrm.cloud.google.com", Kind: "ContainerAnalysisNote"}:

			case schema.GroupKind{Group: "dataflow.cnrm.cloud.google.com", Kind: "DataflowFlexTemplateJob"}:

			case schema.GroupKind{Group: "dataform.cnrm.cloud.google.com", Kind: "DataformRepository"}:

			case schema.GroupKind{Group: "discoveryengine.cnrm.cloud.google.com", Kind: "DiscoveryEngineDataStore"}:

			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPartialPolicy"}:
			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicy"}:
			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicyMember"}:
			case schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMServiceAccount"}:

			case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerCluster"}:
			case schema.GroupKind{Group: "edgecontainer.cnrm.cloud.google.com", Kind: "EdgeContainerNodePool"}:

			case schema.GroupKind{Group: "edgenetwork.cnrm.cloud.google.com", Kind: "EdgeNetworkNetwork"}:
			case schema.GroupKind{Group: "edgenetwork.cnrm.cloud.google.com", Kind: "EdgeNetworkSubnet"}:

			case schema.GroupKind{Group: "firestore.cnrm.cloud.google.com", Kind: "FirestoreDatabase"}:

			case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSKeyRing"}:
			case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSCryptoKey"}:
			case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSAutokeyConfig"}:
			case schema.GroupKind{Group: "kms.cnrm.cloud.google.com", Kind: "KMSKeyHandle"}:

			case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
			case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogBucket"}:
			case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogSink"}:
			case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogView"}:
			//case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLink"}:

			case schema.GroupKind{Group: "managedkafka.cnrm.cloud.google.com", Kind: "ManagedKafkaCluster"}:
			case schema.GroupKind{Group: "managedkafka.cnrm.cloud.google.com", Kind: "ManagedKafkaTopic"}:

			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringAlertPolicy"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringDashboard"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringGroup"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringMetricDescriptor"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringMonitoredProject"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringNotificationChannel"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringService"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringServiceLevelObjective"}:
			case schema.GroupKind{Group: "monitoring.cnrm.cloud.google.com", Kind: "MonitoringUptimeCheckConfig"}:

			case schema.GroupKind{Group: "networkconnectivity.cnrm.cloud.google.com", Kind: "NetworkConnectivityServiceConnectionPolicy"}:

			case schema.GroupKind{Group: "networkservices.cnrm.cloud.google.com", Kind: "NetworkServicesMesh"}:

			case schema.GroupKind{Group: "notebooks.cnrm.cloud.google.com", Kind: "NotebookEnvironment"}:
			case schema.GroupKind{Group: "notebooks.cnrm.cloud.google.com", Kind: "NotebookInstance"}:

			case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACAPool"}:
			case schema.GroupKind{Group: "privateca.cnrm.cloud.google.com", Kind: "PrivateCACertificateAuthority"}:

			case schema.GroupKind{Group: "privilegedaccessmanager.cnrm.cloud.google.com", Kind: "PrivilegedAccessManagerEntitlement"}:

			case schema.GroupKind{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubSchema"}:
			case schema.GroupKind{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubSubscription"}:
			case schema.GroupKind{Group: "pubsub.cnrm.cloud.google.com", Kind: "PubSubTopic"}:

			case schema.GroupKind{Group: "redis.cnrm.cloud.google.com", Kind: "RedisInstance"}:
			case schema.GroupKind{Group: "redis.cnrm.cloud.google.com", Kind: "RedisCluster"}:

			case schema.GroupKind{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "Folder"}:
			case schema.GroupKind{Group: "resourcemanager.cnrm.cloud.google.com", Kind: "Project"}:

			case schema.GroupKind{Group: "pubsublite.cnrm.cloud.google.com", Kind: "PubSubLiteReservation"}:
			case schema.GroupKind{Group: "pubsublite.cnrm.cloud.google.com", Kind: "PubSubLiteSubscription"}:
			case schema.GroupKind{Group: "pubsublite.cnrm.cloud.google.com", Kind: "PubSubLiteTopic"}:
				// ok

			case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecret"}:
			case schema.GroupKind{Group: "secretmanager.cnrm.cloud.google.com", Kind: "SecretManagerSecretVersion"}:

			case schema.GroupKind{Group: "securesourcemanager.cnrm.cloud.google.com", Kind: "SecureSourceManagerInstance"}:
			case schema.GroupKind{Group: "securesourcemanager.cnrm.cloud.google.com", Kind: "SecureSourceManagerRepository"}:

			case schema.GroupKind{Group: "servicedirectory.cnrm.cloud.google.com", Kind: "ServiceDirectoryNamespace"}:
			case schema.GroupKind{Group: "servicedirectory.cnrm.cloud.google.com", Kind: "ServiceDirectoryService"}:

			case schema.GroupKind{Group: "servicenetworking.cnrm.cloud.google.com", Kind: "ServiceNetworkingConnection"}:

			case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "Service"}:

			case schema.GroupKind{Group: "serviceusage.cnrm.cloud.google.com", Kind: "ServiceIdentity"}:

			case schema.GroupKind{Group: "sql.cnrm.cloud.google.com", Kind: "SQLDatabase"}:
			case schema.GroupKind{Group: "sql.cnrm.cloud.google.com", Kind: "SQLInstance"}:
			case schema.GroupKind{Group: "sql.cnrm.cloud.google.com", Kind: "SQLUser"}:

			case schema.GroupKind{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerDatabase"}:
			case schema.GroupKind{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerInstance"}:
			case schema.GroupKind{Group: "spanner.cnrm.cloud.google.com", Kind: "SpannerBackupSchedule"}:

			case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageBucket"}:
			case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageNotification"}:

			case schema.GroupKind{Group: "storage.cnrm.cloud.google.com", Kind: "StorageManagedFolder"}:

			case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagKey"}:
			case schema.GroupKind{Group: "tags.cnrm.cloud.google.com", Kind: "TagsTagValue"}:

			case schema.GroupKind{Group: "cloudtasks.cnrm.cloud.google.com", Kind: "TasksQueue"}:

			case schema.GroupKind{Group: "workflows.cnrm.cloud.google.com", Kind: "WorkflowsWorkflow"}:

			case schema.GroupKind{Group: "workstations.cnrm.cloud.google.com", Kind: "WorkstationCluster"}:
			case schema.GroupKind{Group: "workstations.cnrm.cloud.google.com", Kind: "WorkstationConfig"}:
			case schema.GroupKind{Group: "workstations.cnrm.cloud.google.com", Kind: "Workstation"}:

			case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIDataset"}:
			case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAITensorboard"}:
			case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIEndpoint"}:
			case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIMetadataStore"}:
			case schema.GroupKind{Group: "vertexai.cnrm.cloud.google.com", Kind: "VertexAIFeaturestore"}:

			case schema.GroupKind{Group: "vmwareengine.cnrm.cloud.google.com", Kind: "VMwareEngineNetwork"}:

			case schema.GroupKind{Group: "vpcaccess.cnrm.cloud.google.com", Kind: "VPCAccessConnector"}:

			default:
				t.Skipf("gk %v not suppported by mock gcp %v; skipping", gvk.GroupKind(), name)
			}
		}
	}
	if os.Getenv("E2E_GCP_TARGET") == "vcr" {
		// TODO(yuhou): use a cleaner way(resource kind) to manage the allow list for vcr
		switch name {
		// update test data requires regeneration of the vcr log, skip the test for now.
		// case "fullalloydbcluster":
		case "apikeyskeybasic":
		case "artifactregistryrepository":
		case "bigqueryconnectionconnection":
		case "bigqueryjob":
		case "bigquerytable":
		case "custombudget":
		case "certificatemanagercertificatemapentry":
		case "httpsfunction":
		case "cloudschedulerjob":
		case "globalcomputeforwardingrule":
		case "containernodepool":
		case "containeranalysisnote":
		case "dataproccluster":
		case "cloudstoragepathstoredinfotype":
		case "dnsrecordset":
		case "eventarctrigger":
		case "firestoreindex":
		case "identityplatformoauthidpconfig":
		case "kmscryptokey":
		case "logginglogview":
		case "memcacheinstance":
		case "monitoringalertpolicy":
		case "networkconnectivityhub":
		case "networkservicesgrpcroute":
		case "osconfigguestpolicy":
		case "basicpubsubsubscription":
		case "pubsublitereservation":
		case "androidrecaptchaenterprisekey":
		case "redisinstance":
		case "runservicebasic":
		case "secretmanagersecretversion":
		case "servicedirectorynamespace":
		case "servicenetworkingconnection":
		case "sourcereporepository":
		case "spannerdatabase":
		case "computenodegroup":
		case "computenodetemplate":
		case "privatecacapool":

		case "projectinorg":
		default:
			t.Skipf("test %v not suppported by vcr; skipping", name)
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

func (h *Harness) NoExtraGoldenFiles(glob string) {
	gotFiles, err := filepath.Glob(glob)
	if err != nil {
		h.Fatalf("error matching glob %q: %v", glob, err)
	}

	goldenFilesSet := sets.New(h.goldenFiles...)

	for _, gotFile := range gotFiles {
		abs, err := filepath.Abs(gotFile)
		if err != nil {
			h.Fatalf("error getting absolute path for %q: %v", gotFile, err)
		}
		if goldenFilesSet.Has(abs) {
			continue
		}

		h.Errorf("found extra file %q", gotFile)

		if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
			if err := os.Remove(abs); err != nil {
				h.Errorf("error removing extra file %q", abs)
			}
		}
	}
}

func (h *Harness) CompareGoldenFile(p string, got string, normalizers ...func(s string) string) {
	abs, err := filepath.Abs(p)
	if err != nil {
		h.Fatalf("error converting path %q to absolute path: %v", p, err)
	}
	h.goldenFiles = append(h.goldenFiles, abs)

	test.CompareGoldenFile(h.T, p, got, normalizers...)
}

func (h *Harness) MustReadFile(p string) []byte {
	return test.MustReadFile(h.T, p)
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
