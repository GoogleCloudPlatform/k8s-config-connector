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

package mocktests

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/clientconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type httpRoundTripperKeyType int

// httpRoundTripperKey is the key value for http.RoundTripper in a context.Context
var httpRoundTripperKey httpRoundTripperKeyType

func TestSecretManagerSecretVersion(t *testing.T) {
	h := NewHarness(t)

	dir := "testdata/secretmanager/secret_and_secretversion"

	y := h.MustReadFile(filepath.Join(dir, "input.yaml"))

	t.Logf("parsing objects")
	objects := h.ParseObjects(y)
	h.WithObjects(objects...)

	t.Logf("creating mock cloud")

	mockCloud := mockgcp.NewMockRoundTripper(t, h.Client, storage.NewInMemoryStorage())

	roundTripper := http.RoundTripper(mockCloud)

	artifacts := os.Getenv("ARTIFACTS")
	if artifacts == "" {
		t.Logf("env var ARTIFACTS is not set; will not record http log")
	} else {
		outputDir := filepath.Join(artifacts, "http-logs")

		roundTripper = test.NewHTTPRecorder(mockCloud, outputDir)
	}

	h.Ctx = context.WithValue(h.Ctx, httpRoundTripperKey, roundTripper)

	transport_tpg.DefaultHTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		t := ctx.Value(httpRoundTripperKey)
		if t != nil {
			return &http.Client{Transport: t.(http.RoundTripper)}
		}
		return inner
	}
	transport_tpg.OAuth2HTTPClientTransformer = func(ctx context.Context, inner *http.Client) *http.Client {
		t := ctx.Value(httpRoundTripperKey)
		if t != nil {
			return &http.Client{Transport: t.(http.RoundTripper)}
		}
		return inner
	}

	t.Logf("creating controller")
	mgr, err := ctrl.NewManager(h.RESTConfig(), ctrl.Options{
		MetricsBindAddress: "0",
		NewClient:          h.NewClient,
	})
	if err != nil {
		t.Fatalf("NewManager failed: %v", err)
	}

	t.Logf("creating tfprovider config")
	tfConfig := tfprovider.UnitTestConfig()

	t.Logf("creating tfprovider")
	tfProvider, err := tfprovider.New(h.Ctx, tfConfig)
	if err != nil {
		t.Fatalf("error from tfprovider.New: %v", err)
	}
	t.Logf("creating dclconfig")
	dclConfig := clientconfig.NewForIntegrationTest()
	t.Logf("creating testreconciler")
	testhelper := testreconciler.NewForDCLAndTFTestReconciler(t, mgr, tfProvider, dclConfig)

	for _, object := range objects {
		gvk := object.GetObjectKind().GroupVersionKind()
		if !strings.Contains(gvk.Group, "cnrm.cloud.google.com") {
			continue
		}
		t.Logf("creating reconciler")
		reconciler := testhelper.NewReconcilerForKind(gvk.Kind)
		t.Logf("reconciler for %v is %T", gvk, reconciler)

		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: object.GetNamespace(),
				Name:      object.GetName(),
			},
		}

		result, err := reconciler.Reconcile(h.Ctx, request)
		t.Logf("reconcile result is %#v, %#v", result, err)
		if err != nil {
			t.Errorf("reconcile failed: %v", err)
		}
	}
}
