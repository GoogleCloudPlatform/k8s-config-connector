// Copyright 2026 Google LLC
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

package deletiondefender

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/go-logr/logr"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestReconcilePreservesDeletionSafety(t *testing.T) {
	for _, tc := range []struct {
		name         string
		uninstalling bool
		crdMissing   bool
	}{
		{name: "normal deletion keeps the resource controller finalizer"},
		{name: "uninstalling CRD abandons the underlying resource", uninstalling: true},
		{name: "cannot remove protection without checking the CRD", crdMissing: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			deletedAt := metav1.NewTime(time.Unix(1, 0))
			obj := &unstructured.Unstructured{}
			obj.SetAPIVersion("iam.cnrm.cloud.google.com/v1beta1")
			obj.SetKind("IAMPolicyMember")
			obj.SetName("test-policy")
			obj.SetNamespace("test-namespace")
			obj.SetDeletionTimestamp(&deletedAt)
			// Keep the fake API object available after the KCC finalizers are removed.
			const holdFinalizer = "test.example.com/hold"
			obj.SetFinalizers([]string{k8s.DeletionDefenderFinalizerName, k8s.ControllerFinalizerName, holdFinalizer})
			c := fake.NewClientBuilder().WithObjects(obj).Build()
			crd := &apiextensions.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{Name: "iampolicymembers.iam.cnrm.cloud.google.com"},
			}
			if tc.uninstalling {
				crd.SetDeletionTimestamp(&deletedAt)
			}
			server := httptest.NewServer(&crdHandler{crd: crd, missing: tc.crdMissing})
			defer server.Close()
			clientSet, err := clientset.NewForConfig(&rest.Config{Host: server.URL})
			if err != nil {
				t.Fatal(err)
			}
			r := &Reconciler{Client: c, clientSet: clientSet, crd: crd, gvk: obj.GroupVersionKind(), logger: logr.Discard()}
			key := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
			if _, err := r.Reconcile(ctx, reconcile.Request{NamespacedName: key}); (err != nil) != tc.crdMissing {
				t.Fatalf("Reconcile() error = %v, want error %v", err, tc.crdMissing)
			}
			got := &unstructured.Unstructured{}
			got.SetGroupVersionKind(obj.GroupVersionKind())
			if err := c.Get(ctx, key, got); err != nil {
				t.Fatal(err)
			}
			if has := k8s.HasFinalizer(got, k8s.DeletionDefenderFinalizerName); has != tc.crdMissing {
				t.Errorf("deletion defender finalizer present = %v, want %v", has, tc.crdMissing)
			}
			if has := k8s.HasFinalizer(got, k8s.ControllerFinalizerName); has != !tc.uninstalling {
				t.Errorf("resource controller finalizer present = %v, want %v", has, !tc.uninstalling)
			}
			if !k8s.HasFinalizer(got, holdFinalizer) {
				t.Error("unrelated finalizer was removed")
			}
			wantPolicy := ""
			if tc.uninstalling {
				wantPolicy = k8s.DeletionPolicyAbandon
			}
			if gotPolicy := got.GetAnnotations()[k8s.DeletionPolicyAnnotation]; gotPolicy != wantPolicy {
				t.Errorf("deletion policy = %q, want %q", gotPolicy, wantPolicy)
			}
		})
	}
}

type crdHandler struct {
	crd     *apiextensions.CustomResourceDefinition
	missing bool
}

func (h *crdHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if h.missing {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(metav1.Status{Status: metav1.StatusFailure, Reason: metav1.StatusReasonNotFound, Code: http.StatusNotFound})
		return
	}
	_ = json.NewEncoder(w).Encode(h.crd)
}
