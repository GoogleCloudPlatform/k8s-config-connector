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

package identity

import (
	"context"
	"errors"
	"strings"
	"testing"

	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var policyMemberGK = schema.GroupKind{Group: "iam.cnrm.cloud.google.com", Kind: "IAMPolicyMember"}

func staticToken(token string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token, TokenType: "Bearer"})
}

func newObject(namespace, name string, annotations map[string]string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetNamespace(namespace)
	u.SetName(name)
	u.SetAnnotations(annotations)
	return u
}

// namespaceResolver hands out a token named after the object's namespace.
var namespaceResolver = ResolverFunc(func(_ context.Context, target Target) (*Credentials, error) {
	namespace := target.Object.GetNamespace()
	return &Credentials{ID: "sa@" + namespace, TokenSource: staticToken("token-" + namespace)}, nil
})

// ctxFor returns a context that carries namespaceResolver's credentials for an
// object in namespace.
func ctxFor(t *testing.T, namespace string) context.Context {
	t.Helper()
	ctx, err := ForObject(WithResolver(context.Background(), namespaceResolver), Target{GroupKind: policyMemberGK, Object: newObject(namespace, "member", nil)})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	return ctx
}

func TestForObjectPassesTheObject(t *testing.T) {
	var got Target
	resolver := ResolverFunc(func(_ context.Context, target Target) (*Credentials, error) {
		got = target
		project := target.Object.GetAnnotations()["cnrm.cloud.google.com/project-id"]
		return &Credentials{ID: "sa@" + project, TokenSource: staticToken(project)}, nil
	})
	obj := newObject("team-a", "member", map[string]string{"cnrm.cloud.google.com/project-id": "project-a"})

	ctx, err := ForObject(WithResolver(context.Background(), resolver), Target{GroupKind: policyMemberGK, Object: obj})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	if got.Object != obj || got.GroupKind != policyMemberGK {
		t.Errorf("resolver got %v, want the reconciled object", got)
	}
	creds, ok := CredentialsFromContext(ctx)
	if !ok || creds.ID != "sa@project-a" {
		t.Errorf("CredentialsFromContext = %v, %v; want sa@project-a", creds, ok)
	}
}

func TestForObjectWithoutResolver(t *testing.T) {
	ctx := context.Background()
	got, err := ForObject(WithResolver(ctx, nil), Target{GroupKind: policyMemberGK, Object: newObject("team-a", "member", nil)})
	if err != nil {
		t.Fatalf("ForObject: %v", err)
	}
	if got != ctx {
		t.Errorf("ForObject without a resolver changed the context")
	}
	if _, ok := CredentialsFromContext(got); ok {
		t.Errorf("ForObject without a resolver recorded credentials")
	}
}

func TestForObjectFailsEarly(t *testing.T) {
	tests := []struct {
		name    string
		creds   *Credentials
		err     error
		wantErr string
	}{
		{name: "resolver error", err: errors.New("hook unreachable"), wantErr: "hook unreachable"},
		{name: "no credentials", wantErr: "no credentials"},
		{name: "no ID", creds: &Credentials{TokenSource: staticToken("x")}, wantErr: "no credentials"},
		{name: "no token source", creds: &Credentials{ID: "sa@team-a"}, wantErr: "no credentials"},
	}
	for _, tc := range tests {
		resolver := ResolverFunc(func(context.Context, Target) (*Credentials, error) {
			return tc.creds, tc.err
		})
		ctx, err := ForObject(WithResolver(context.Background(), resolver), Target{GroupKind: policyMemberGK, Object: newObject("team-a", "member", nil)})
		if err == nil || !strings.Contains(err.Error(), tc.wantErr) || !strings.Contains(err.Error(), "IAMPolicyMember.iam.cnrm.cloud.google.com team-a/member") {
			t.Errorf("%s: ForObject error = %v, want one naming the object and containing %q", tc.name, err, tc.wantErr)
		}
		if _, ok := CredentialsFromContext(ctx); ok {
			t.Errorf("%s: ForObject recorded credentials despite the error", tc.name)
		}
	}
}
