package restmapper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"
)

func TestRESTMapping(t *testing.T) {
	k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":0")
	if err != nil {
		t.Fatalf("error building mock kube-apiserver: %v", err)
	}

	k8s.RegisterType(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}, "namespaces", meta.RESTScopeRoot)

	defer func() {
		if err := k8s.Stop(); err != nil {
			t.Fatalf("error closing mock kube-apiserver: %v", err)
		}
	}()

	addr, err := k8s.StartServing()
	if err != nil {
		t.Errorf("error starting mock kube-apiserver: %v", err)
	}

	klog.Infof("mock kubeapiserver will listen on %v", addr)

	restConfig := &rest.Config{
		Host: addr.String(),
	}

	restMapper, err := NewForTest(restConfig)
	if err != nil {
		t.Fatalf("error from NewForTest: %v", err)
	}
	gvk := schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}
	restMapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		t.Fatalf("error from RESTMapping(%v): %v", gvk, err)
	}

	got := fmt.Sprintf("resource:%v\ngvk:%v\nscope:%v", restMapping.Resource, restMapping.GroupVersionKind, restMapping.Scope.Name())
	want := `
resource:/v1, Resource=namespaces
gvk:/v1, Kind=Namespace
scope:root
`
	got = strings.TrimSpace(got)
	want = strings.TrimSpace(want)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("RESTMapping(%v) diff (-want +got):\n%s", gvk, diff)
	}
}
