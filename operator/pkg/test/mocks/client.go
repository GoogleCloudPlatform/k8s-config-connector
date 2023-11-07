package mocks

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/kubebuilder-declarative-pattern/mockkubeapiserver"
)

func GetMockClient(t *testing.T) client.Client {
	k8s, err := mockkubeapiserver.NewMockKubeAPIServer(":http")
	if err != nil {
		t.Fatalf("error building mock kube-apiserver: %v", err)
	}

	addr, err := k8s.StartServing()
	if err != nil {
		t.Errorf("error starting mock kube-apiserver: %v", err)
	}

	if addr == nil {
		t.Fatalf("address of the mock kube-apiserver is nil")
	}

	t.Cleanup(func() {
		if err := k8s.Stop(); err != nil {
			t.Errorf("error stopping mock kube-apiserver: %v", err)
		}
	})

	restConfig := &rest.Config{
		Host: addr.String(),
	}

	mgr, err := manager.New(restConfig, manager.Options{
		Scheme: runtime.NewScheme(),
	})
	if err != nil {
		t.Fatalf("error building manager: %v", err)
	}
	return mgr.GetClient()
}
