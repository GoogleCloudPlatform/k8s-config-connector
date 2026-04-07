package simpletest

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/kubebuilder-declarative-pattern/commonclient"
	"sigs.k8s.io/kubebuilder-declarative-pattern/ktest/httprecorder"
	"sigs.k8s.io/kubebuilder-declarative-pattern/ktest/testharness"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/loaders"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/restmapper"

	api "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/testreconciler/simpletest/v1alpha1"
)

func testSimpleReconciler(h *testharness.Harness, testdir string, applier applier.Applier, status declarative.Status) {
	ctx := context.Background()

	k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
	if err != nil {
		h.Fatalf("error building mock kube-apiserver: %v", err)
	}

	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}, "namespaces", meta.RESTScopeRoot)
	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}, "configmaps", meta.RESTScopeNamespace)
	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Event"}, "events", meta.RESTScopeNamespace)
	k8s.RegisterType(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}, "deployments", meta.RESTScopeNamespace)
	k8s.RegisterType(schema.GroupVersionKind{Group: "addons.example.org", Version: "v1alpha1", Kind: "SimpleTest"}, "simpletests", meta.RESTScopeNamespace)

	defer func() {
		if err := k8s.Stop(); err != nil {
			h.Fatalf("error closing mock kube-apiserver: %v", err)
		}
	}()

	addr, err := k8s.StartServing()
	if err != nil {
		h.Errorf("error starting mock kube-apiserver: %v", err)
	}

	klog.Infof("mock kubeapiserver will listen on %v", addr)

	var requestLog httprecorder.RequestLog
	wrapTransport := func(rt http.RoundTripper) http.RoundTripper {
		return httprecorder.NewRecorder(rt, &requestLog)
	}
	restConfig := &rest.Config{
		Host:          addr.String(),
		WrapTransport: wrapTransport,
	}

	scheme := runtime.NewScheme()
	if err := api.AddToScheme(scheme); err != nil {
		h.Fatalf("error from AddToScheme: %v", err)
	}

	logger := klogr.New()
	mgrOpt := ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,

		// MapperProvider provides the rest mapper used to map go types to Kubernetes APIs
		MapperProvider: restmapper.NewControllerRESTMapper,

		Logger: logger,
	}
	err = commonclient.SetMetricsBindAddress(&mgrOpt, "0")
	if err != nil {
		h.Errorf("error configuring manager: %v", err)
	}
	mgr, err := ctrl.NewManager(restConfig, mgrOpt)
	if err != nil {
		h.Fatalf("error starting manager: %v", err)
	}

	reconciler := &SimpleTestReconciler{
		TestSuffix: testdir,
		Client:     mgr.GetClient(),
		Scheme:     mgr.GetScheme(),
		applier:    applier,
		status:     status,
	}

	mc, err := loaders.NewManifestLoader("testdata/channels")
	if err != nil {
		h.Fatalf("error from loaders.NewManifestLoader: %v", err)
	}
	reconciler.manifestController = mc

	if err = reconciler.SetupWithManager(mgr); err != nil {
		h.Fatalf("error creating reconciler: %v", err)
	}

	if h.FileExists(filepath.Join(testdir, "before.yaml")) {
		before := string(h.MustReadFile(filepath.Join(testdir, "before.yaml")))
		if err := k8s.AddObjectsFromManifest(before); err != nil {
			h.Fatalf("error precreating objects: %v", err)
		}
	}

	mgrContext, mgrStop := context.WithCancel(ctx)
	go func() {
		time.Sleep(1 * time.Second)
		// time.Sleep(200 * time.Second)
		mgrStop()
	}()
	if err := mgr.Start(mgrContext); err != nil {
		h.Fatalf("error starting manager: %v", err)
	}

	h.Logf("replacing old url prefix %q", "http://"+restConfig.Host)
	requestLog.ReplaceURLPrefix("http://"+restConfig.Host, "http://kube-apiserver")
	requestLog.RemoveUserAgent()
	requestLog.SortGETs()
	// Workaround for non-determinism in https://github.com/kubernetes/kubernetes/blob/79a62d62350fb600f97d1f6309c3274515b3587a/staging/src/k8s.io/client-go/tools/cache/reflector.go#L301
	requestLog.RegexReplaceURL(h.T, "&timeoutSeconds=.*&", "&timeoutSeconds=<replaced>&")
	h.Logf("replacing real timestamp in request and response to a fake value")
	requestLog.ReplaceTimestamp()
	requestLog.ReplaceHeader("Date", "(removed)")

	requests := requestLog.FormatHTTP(false)

	h.CompareGoldenFile(filepath.Join(testdir, "expected-http.yaml"), requests)
}
