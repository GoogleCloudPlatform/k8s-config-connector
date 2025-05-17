// Copyright 2025 Google LLC
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

package mockgcptests

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcpregistry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/log"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	cloudresourcemanagerv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GCPTargetMode string

const (
	GCPTargetModeReal GCPTargetMode = "real"
	GCPTargetModeMock GCPTargetMode = "mock"
	GCPTargetModeVCR  GCPTargetMode = "vcr"
)

type Harness struct {
	*testing.T

	MockGCP            mockgcp.Interface
	registeredServices mockgcpregistry.Normalizer

	HTTPClient *http.Client

	Ctx context.Context

	Events *test.MemoryEventSink

	Project GCPProject

	// GRPCUnaryClientInterceptor is the GRPC interceptor for use in tests.
	GRPCUnaryClientInterceptor grpc.UnaryClientInterceptor

	gcpAccessToken string

	Endpoint string

	GcloudConfiguration string

	proxy         *Proxy
	ProxyEndpoint *net.TCPAddr

	goldenFiles []string

	// GCPTarget is the GCP mode to use (real, mock, vcr)
	// If not set, will use the E2E_GCP_TARGET env var
	GCPTarget GCPTargetMode
}

type HarnessOption func(*Harness)

func WithGCPTarget(gcpTarget GCPTargetMode) HarnessOption {
	return func(h *Harness) {
		h.GCPTarget = gcpTarget
	}
}
func NewHarness(ctx context.Context, t *testing.T, opts ...HarnessOption) *Harness {
	ctx, close := context.WithCancel(ctx)

	h := &Harness{
		T:   t,
		Ctx: ctx,
	}

	for _, opt := range opts {
		opt(h)
	}

	if h.GCPTarget == "" {
		h.GCPTarget = GCPTargetMode(os.Getenv("E2E_GCP_TARGET"))
	}

	t.Cleanup(h.cleanup)
	t.Cleanup(close)
	return h
}

// cleanup is called when the test cleans up
func (t *Harness) cleanup() {

}

func (h *Harness) RegisteredServices() mockgcpregistry.Normalizer {
	return h.registeredServices
}

func (h *Harness) CompareGoldenFile(p string, got string, normalizers ...func(s string) string) {
	abs, err := filepath.Abs(p)
	if err != nil {
		h.Fatalf("error converting path %q to absolute path: %v", p, err)
	}
	h.goldenFiles = append(h.goldenFiles, abs)

	test.CompareGoldenFile(h.T, p, got, normalizers...)
}

func (t *Harness) getCloudResourceManagerClient(httpClient *http.Client) *cloudresourcemanagerv1.Service {
	s, err := cloudresourcemanagerv1.NewService(t.Ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("error building cloudresourcemanagerv1 client: %v", err)
	}
	return s
}

func (t *Harness) Init() {
	// t := h.t
	// log := klog.FromContext(ctx)

	ctx := t.Ctx

	var mockCloudGRPCClientConnection *grpc.ClientConn
	if t.GCPTarget == GCPTargetModeMock {
		t.Logf("creating mock gcp")

		var kubeClient client.Client // TODO: We should replace this, it didn't work
		mockCloud := mockgcp.NewMockRoundTripperForTest(t.T, kubeClient, storage.NewInMemoryStorage())

		mockCloudGRPCClientConnection = mockCloud.NewGRPCConnection(ctx)
		t.MockGCP = mockCloud
		t.registeredServices = mockCloud.(mockgcpregistry.Normalizer)

		roundTripper := http.RoundTripper(mockCloud)

		// ctx = context.WithValue(ctx, httpRoundTripperKey, roundTripper)
		// t.Ctx = ctx

		httpClient := &http.Client{Transport: roundTripper}
		// kccConfig.HTTPClient = httpClient
		t.HTTPClient = httpClient

		// Also hook the oauth2 library
		ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
		t.Ctx = ctx

		t.gcpAccessToken = "dummytoken"
		// kccConfig.GCPAccessToken = h.gcpAccessToken
	} else if t.GCPTarget == GCPTargetModeReal {
		t.Logf("targeting real GCP")
		// } else if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "vcr" {
		// 	t.Logf("creating vcr test")

		// We create registered services, even though we only use it for replacements
		var kubeClient client.Client // TODO: We should replace this, it didn't work
		mockCloud := mockgcp.NewMockRoundTripperForTest(t.T, kubeClient, storage.NewInMemoryStorage())
		t.registeredServices = mockCloud.(mockgcpregistry.Normalizer)
	} else {
		t.Fatalf("E2E_GCP_TARGET=%q not supported", t.GCPTarget)
	}

	if t.GCPTarget == GCPTargetModeMock {
		// // Some fixed-value fake org-ids for testing.
		// // We used fixed values so that the output is predictable (for golden testing)
		// testgcp.TestFolderID.Set("123451001")
		// testgcp.TestFolder2ID.Set("123451002")
		// testgcp.TestOrgID.Set("123450001")
		// testgcp.IsolatedTestOrgName.Set("isolated-test-org.example.com")
		// testgcp.TestBillingAccountID.Set("123456-777777-000001")
		// testgcp.TestBillingAccountIDForBillingResources.Set("123456-777777-000003")
		// testgcp.IAMIntegrationTestsOrganizationID.Set("123450002")
		// testgcp.IAMIntegrationTestsBillingAccountID.Set("123456-777777-000002")
		// testgcp.TestAttachedClusterName.Set("xks-cluster")
		// testgcp.TestDependentNoNetworkProjectID.Set("mock-project")
		// testgcp.TestDependentOrgProjectID.Set("example-project-01")
		// testgcp.TestDependentFolderProjectID.Set("example-project-02")
		// testgcp.FirestoreTestProject.Set("cnrm-test-firestore")
		// testgcp.IdentityPlatformTestProject.Set("kcc-identity-platform")
		// testgcp.RecaptchaEnterpriseTestProject.Set("kcc-recaptcha-enterprise")

		crm := t.getCloudResourceManagerClient(t.HTTPClient)
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
		project := GCPProject{
			ProjectID:     found.ProjectId,
			ProjectNumber: found.ProjectNumber,
		}
		testgcp.TestKCCAttachedClusterProject.Set("mock-project")
		testgcp.TestKCCAttachedClusterPlatformVersion.Set("1.30.0-gke.1")
		t.Project = project
		// } else if os.Getenv("E2E_GCP_TARGET") == "vcr" && os.Getenv("VCR_MODE") == "replay" {
		// 	t.gcpAccessToken = "dummytoken"
		// 	// kccConfig.GCPAccessToken = h.gcpAccessToken

		// 	t.Project = testgcp.GCPProject{
		// 		ProjectID:     "example-project",
		// 		ProjectNumber: 123456789,
		// 	}
		// 	testgcp.TestDependentOrgProjectID.Set("example-project-01")
		// 	testgcp.TestDependentFolderProjectID.Set("example-project-02")
		// 	testgcp.FirestoreTestProject.Set("cnrm-test-firestore")
		// 	testgcp.IdentityPlatformTestProject.Set("kcc-identity-platform")
		// 	testgcp.RecaptchaEnterpriseTestProject.Set("kcc-recaptcha-enterprise")
		// 	testgcp.TestOrgID.Set("123450001")
		// 	testgcp.TestBillingAccountID.Set("123456-777777-000001")
		// 	testgcp.TestBillingAccountIDForBillingResources.Set("123456-777777-000003")
	} else {
		t.Project = GetDefaultProject(t.T)
	}

	eventSink := test.NewMemoryEventSink()
	ctx = test.AddSinkToContext(ctx, eventSink)
	t.Ctx = ctx

	t.Events = eventSink

	eventSinks := test.EventSinksFromContext(ctx)

	// Set up event sink for logging to a file, if ARTIFACTS env var is set
	if artifacts := os.Getenv("ARTIFACTS"); artifacts != "" {
		outputDir := filepath.Join(artifacts, "http-logs")
		eventSinks = append(eventSinks, test.NewDirectoryEventSink(outputDir))
	} else {
		log.Info("env var ARTIFACTS is not set; will not record http log")
	}

	if t.GCPTarget == GCPTargetModeVCR {
		t.Fatalf("vcr target not supported here")
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

		t.GRPCUnaryClientInterceptor = grpcUnaryInterceptor

		// Intercept (and log) DCL and direct(non TF) requests
		if len(eventSinks) != 0 {
			if t.HTTPClient == nil {
				httpClient, err := google.DefaultClient(ctx, gcp.ClientScopes...)
				if err != nil {
					t.Fatalf("error creating the http client to be used by DCL: %v", err)
				}
				t.HTTPClient = httpClient
			}
			transport := test.NewHTTPRecorder(t.HTTPClient.Transport, eventSinks...)
			t.HTTPClient = &http.Client{Transport: transport}
		}

	}
}

func (t *Harness) StartProxy(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	httpClient := t.HTTPClient
	t.proxy = NewProxy(httpClient)

	httpListener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("net.Listen failed: %v", err)
	}

	httpServer := &http.Server{}
	httpServer.Handler = t.proxy
	go func() {
		if err := httpServer.Serve(httpListener); err != nil {
			if err != http.ErrServerClosed {
				t.Errorf("error from http proxy server: %v", err)
			}
		}
	}()

	go func() {
		if err := t.proxy.ListenAndServeHTTPS(ctx); err != nil {
			if err != http.ErrServerClosed {
				t.Errorf("error from http proxy server: %v", err)
			}
		}
	}()

	t.Cleanup(func() {
		httpServer.Close()
		cancel()
	})

	t.ProxyEndpoint = httpListener.Addr().(*net.TCPAddr)
}

func (t *Harness) RunCommand(cmdline string) {
	args := strings.Fields(cmdline)
	cmd := exec.CommandContext(t.Ctx, args[0], args[1:]...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	t.Logf("running command %q", cmdline)
	if err := cmd.Run(); err != nil {
		t.Logf("stdout: %v", stdout.String())
		t.Logf("stderr: %v", stderr.String())
		t.Fatalf("running command %q: %v", strings.Join(args, " "), err)
	}
}

func (h *Harness) GCPAuthorization() oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: h.gcpAccessToken})
}

// GCPHTTPClient is the http.Client to use when talking to GCP
// It is wired up to our mocks for tests.
func (h *Harness) GCPHTTPClient() *http.Client {
	return h.HTTPClient
}
