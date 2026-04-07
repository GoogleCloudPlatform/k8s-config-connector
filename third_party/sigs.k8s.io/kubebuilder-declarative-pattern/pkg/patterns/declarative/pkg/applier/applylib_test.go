package applier

import (
	"context"
	"net/http"
	"path/filepath"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/kubebuilder-declarative-pattern/applylib/applyset"
	"sigs.k8s.io/kubebuilder-declarative-pattern/ktest/httprecorder"
	"sigs.k8s.io/kubebuilder-declarative-pattern/ktest/testharness"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/restmapper"
)

func fakeParent() *unstructured.Unstructured {
	parent := &unstructured.Unstructured{}
	parent.SetKind("ConfigMap")
	parent.SetAPIVersion("v1")
	parent.SetName("test")
	parent.SetNamespace("default")
	return parent
}

func TestApplySetApplier(t *testing.T) {
	patchOptions := metav1.PatchOptions{FieldManager: "kdp-test"}
	applier := NewApplySetApplier(patchOptions, metav1.DeleteOptions{}, ApplysetOptions{})
	runApplierGoldenTests(t, "testdata/applylib", false, applier)
}

func runApplierGoldenTests(t *testing.T, testDir string, interceptHTTPServer bool, applier Applier) {
	testharness.RunGoldenTests(t, testDir, func(h *testharness.Harness, testdir string) {
		ctx := context.Background()
		t := h.T

		env := &envtest.Environment{}
		k8s := testharness.NewTestKubeAPIServer(t, ctx, env)

		var apiserverRequestLog httprecorder.RequestLog
		if interceptHTTPServer {
			k8s.AddProxyAndRecordToLog(&apiserverRequestLog)
		}

		restConfig := k8s.RESTConfig()

		var requestLog httprecorder.RequestLog
		wrapTransport := func(rt http.RoundTripper) http.RoundTripper {
			return httprecorder.NewRecorder(rt, &requestLog)
		}
		restConfig.WrapTransport = wrapTransport

		if h.FileExists(filepath.Join(testdir, "before.yaml")) {
			before := string(h.MustReadFile(filepath.Join(testdir, "before.yaml")))
			if err := k8s.AddObjectsFromManifest(before); err != nil {
				t.Fatalf("error precreating objects: %v", err)
			}
		}
		p := filepath.Join(testdir, "manifest.yaml")
		manifestYAML := string(h.MustReadFile(p))
		objects, err := manifest.ParseObjects(ctx, manifestYAML)
		if err != nil {
			t.Errorf("error parsing manifest %q: %v", p, err)
		}

		restMapper, err := restmapper.NewForTest(restConfig)
		if err != nil {
			t.Fatalf("error from controllerrestmapper.NewForTest: %v", err)
		}

		parent := fakeParent()
		parentGVK := parent.GroupVersionKind()
		parentRESTMapping, err := restMapper.RESTMapping(parentGVK.GroupKind(), parentGVK.Version)
		if err != nil {
			t.Fatalf("error getting restmapping for parent %v", err)
		}
		k8s.AddObject(parent)

		client := k8s.Client()
		options := ApplierOptions{
			Objects:    objects.GetItems(),
			RESTConfig: restConfig,
			RESTMapper: restMapper,
			ParentRef:  applyset.NewParentRef(parent, parent.GetName(), parent.GetNamespace(), parentRESTMapping),
			Client:     client,
		}
		if err := applier.Apply(ctx, options); err != nil {
			t.Fatalf("error from applier.Apply: %v", err)
		}

		httprecorder.NormalizeKubeRequestLog(t, &requestLog, restConfig)
		requests := requestLog.FormatHTTP(false)
		h.CompareGoldenFile(filepath.Join(testdir, "expected.yaml"), requests)

		if interceptHTTPServer {
			httprecorder.NormalizeKubeRequestLog(t, &apiserverRequestLog, restConfig)
			apiserverRequests := apiserverRequestLog.FormatHTTP(false)
			h.CompareGoldenFile(filepath.Join(testdir, "expected-apiserver.yaml"), apiserverRequests)
		}
	})
}
