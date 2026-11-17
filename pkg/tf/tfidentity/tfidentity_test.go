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

package tfidentity

import (
	"context"
	"net/http"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func newTestProvider(ctx context.Context) (*tfschema.Provider, *transport_tpg.Config) {
	base := &transport_tpg.Config{
		Context:        ctx,
		Client:         &http.Client{},
		BatchingConfig: &transport_tpg.BatchingConfig{},
	}
	provider := &tfschema.Provider{ResourcesMap: map[string]*tfschema.Resource{"google_folder_iam_member": {}}}
	provider.SetMeta(base)
	return provider, base
}

// byNamespace hands out an identity named after the object's namespace.
var byNamespace = identity.ResolverFunc(func(_ context.Context, target identity.Target) (*identity.Credentials, error) {
	namespace := target.Object.GetNamespace()
	return &identity.Credentials{
		ID:          "sa@" + namespace,
		TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: namespace}),
	}, nil
})

// forObject returns ctx with byNamespace's credentials for an object, as the
// TF controller's context would carry them.
func forObject(t *testing.T, ctx context.Context, namespace, name string) context.Context {
	t.Helper()
	obj := &unstructured.Unstructured{}
	obj.SetNamespace(namespace)
	obj.SetName(name)
	ctx, err := identity.ForObject(identity.WithResolver(ctx, byNamespace), identity.Target{Object: obj})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	return ctx
}

func TestProvider(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	provider, base := newTestProvider(ctx)

	metaFor := func(ctx context.Context) *transport_tpg.Config {
		t.Helper()
		p, err := Provider(ctx, provider)
		if err != nil {
			t.Fatalf("Provider: %v", err)
		}
		if len(p.ResourcesMap) != len(provider.ResourcesMap) {
			t.Fatalf("Provider copy does not share the resource map")
		}
		return p.Meta().(*transport_tpg.Config)
	}

	if got := metaFor(ctx); got != base {
		t.Errorf("without a resolver, Provider returned a copy, want the provider meta")
	}

	teamA := metaFor(forObject(t, ctx, "team-a", "one"))
	if teamA == base {
		t.Fatalf("with credentials, Provider returned the provider meta, want a copy")
	}
	if again := metaFor(forObject(t, ctx, "team-a", "two")); again != teamA {
		t.Errorf("same credentials ID returned a different copy, want the cached one")
	}
	if teamB := metaFor(forObject(t, ctx, "team-b", "one")); teamB == teamA {
		t.Errorf("different credentials IDs share a copy")
	}

	teamACtx := forObject(t, ctx, "team-a", "one")
	copyA, err := Provider(teamACtx, provider)
	if err != nil {
		t.Fatalf("Provider: %v", err)
	}
	if again, err := Provider(teamACtx, copyA); err != nil || again != copyA {
		t.Errorf("Provider on its own copy = %p, %v; want the same copy %p", again, err, copyA)
	}
	teamBCtx := forObject(t, ctx, "team-b", "one")
	copyB, err := Provider(teamBCtx, provider)
	if err != nil {
		t.Fatalf("Provider: %v", err)
	}
	if fromA, err := Provider(teamBCtx, copyA); err != nil || fromA != copyB {
		t.Errorf("Provider for team-b from team-a's copy = %p, %v; want team-b's copy %p", fromA, err, copyB)
	}

	if p, err := Provider(identity.WithResolver(ctx, byNamespace), provider); err == nil {
		t.Errorf("Provider with a resolver but no credentials = %p, want an error", p)
	}
}
