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

// Package identity lets Config Connector pick GCP credentials per KRM object.
//
// With a Resolver, each reconciler calls ForObject after reading its object,
// and its GCP calls use the returned credentials. Nothing falls back to the
// default credentials. Without a Resolver, nothing changes.
package identity

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Target is the KRM object that GCP calls are made for.
type Target struct {
	// GroupKind is the object's kind; its TypeMeta can be empty.
	GroupKind schema.GroupKind

	// Object is the reconciled object. Do not modify it.
	Object client.Object
}

func (t Target) String() string {
	if t.Object == nil {
		return t.GroupKind.String()
	}
	return fmt.Sprintf("%s %s/%s", t.GroupKind.String(), t.Object.GetNamespace(), t.Object.GetName())
}

// Credentials are the GCP credentials a Resolver returns.
type Credentials struct {
	// ID names the identity, such as a service account email. Some clients
	// are cached per ID, so it must change when the identity does.
	ID string

	// TokenSource mints the identity's tokens. It should cache them, as
	// oauth2.ReuseTokenSource does.
	TokenSource oauth2.TokenSource
}

// Resolver picks the GCP credentials for a KRM object.
type Resolver interface {
	// Resolve returns target's credentials. An error fails the reconcile.
	Resolve(ctx context.Context, target Target) (*Credentials, error)
}

// ResolverFunc adapts a function to the Resolver interface.
type ResolverFunc func(ctx context.Context, target Target) (*Credentials, error)

// Resolve calls f.
func (f ResolverFunc) Resolve(ctx context.Context, target Target) (*Credentials, error) {
	return f(ctx, target)
}

type resolverKey struct{}

type credentialsKey struct{}

// WithResolver returns ctx carrying resolver, or ctx if resolver is nil.
func WithResolver(ctx context.Context, resolver Resolver) context.Context {
	if resolver == nil {
		return ctx
	}
	return context.WithValue(ctx, resolverKey{}, resolver)
}

// ResolverFromContext returns the resolver recorded by WithResolver, or nil.
func ResolverFromContext(ctx context.Context) Resolver {
	resolver, _ := ctx.Value(resolverKey{}).(Resolver)
	return resolver
}

// ForObject returns ctx carrying target's credentials from the resolver in
// ctx, or ctx if there is no resolver. It fails if the resolver errors or
// returns no credentials.
func ForObject(ctx context.Context, target Target) (context.Context, error) {
	resolver := ResolverFromContext(ctx)
	if resolver == nil {
		return ctx, nil
	}
	creds, err := resolver.Resolve(ctx, target)
	if err != nil {
		return ctx, fmt.Errorf("resolving GCP credentials for %v: %w", target, err)
	}
	if creds == nil || creds.ID == "" || creds.TokenSource == nil {
		return ctx, fmt.Errorf("resolving GCP credentials for %v: resolver returned no credentials", target)
	}
	return context.WithValue(ctx, credentialsKey{}, creds), nil
}

// CredentialsFromContext returns the credentials recorded by ForObject.
func CredentialsFromContext(ctx context.Context) (*Credentials, bool) {
	creds, ok := ctx.Value(credentialsKey{}).(*Credentials)
	return creds, ok
}
