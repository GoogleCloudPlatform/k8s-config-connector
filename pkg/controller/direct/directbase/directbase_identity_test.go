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

package directbase

import (
	"context"
	"errors"
	"strings"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	"golang.org/x/oauth2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// credentialsModel records the credentials in the context it is called with.
type credentialsModel struct {
	called bool
	creds  *identity.Credentials
}

func (m *credentialsModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	m.called = true
	m.creds, _ = identity.CredentialsFromContext(ctx)
	return nil, errors.New("stop after AdapterForObject")
}

func (m *credentialsModel) AdapterForURL(ctx context.Context, url string) (Adapter, error) {
	return nil, errors.New("not implemented")
}

func reconcileWithResolver(t *testing.T, resolver identity.Resolver) (*credentialsModel, error) {
	t.Helper()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetAnnotations(map[string]string{"cnrm.cloud.google.com/project-id": "project-a"})

	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(u).
		WithRuntimeObjects(u).
		Build()
	model := &credentialsModel{}
	r := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(k8sClient, record.NewFakeRecorder(100)),
		Client:           k8sClient,
		scheme:           scheme,
		gvk:              testGVK,
		model:            model,
		jitterGenerator:  &mockJitterGenerator{},
	}
	ctx := identity.WithResolver(context.Background(), resolver)
	_, err := r.Reconcile(ctx, reconcile.Request{NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "test-resource"}})
	return model, err
}

func TestReconcile_IdentityResolver_UsesObjectCredentials(t *testing.T) {
	var got identity.Target
	resolver := identity.ResolverFunc(func(_ context.Context, target identity.Target) (*identity.Credentials, error) {
		got = target
		project := target.Object.GetAnnotations()["cnrm.cloud.google.com/project-id"]
		return &identity.Credentials{
			ID:          "sa@" + project,
			TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: project}),
		}, nil
	})

	model, _ := reconcileWithResolver(t, resolver)

	if got.GroupKind != testGVK.GroupKind() || got.Object == nil || got.Object.GetNamespace() != "test-ns" || got.Object.GetName() != "test-resource" {
		t.Errorf("resolver got %v, want the reconciled object", got)
	}
	if !model.called {
		t.Fatalf("AdapterForObject was not called")
	}
	if model.creds == nil || model.creds.ID != "sa@project-a" {
		t.Errorf("AdapterForObject got credentials %v, want sa@project-a", model.creds)
	}
}

func TestReconcile_IdentityResolver_FailsEarly(t *testing.T) {
	resolver := identity.ResolverFunc(func(context.Context, identity.Target) (*identity.Credentials, error) {
		return nil, errors.New("hook unreachable")
	})

	model, err := reconcileWithResolver(t, resolver)

	if err == nil || !strings.Contains(err.Error(), "hook unreachable") {
		t.Errorf("Reconcile error = %v, want the resolver error", err)
	}
	if model.called {
		t.Errorf("AdapterForObject was called despite the resolver error")
	}
}
