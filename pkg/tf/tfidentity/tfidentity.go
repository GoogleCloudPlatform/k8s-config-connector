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

// Package tfidentity applies the identity.Resolver to Terraform calls.
//
// Many Terraform resources build their GCP clients from the provider meta and
// send requests without a context, so the credentials cannot be picked per
// request. Instead, each identity gets its own copy of the provider, with its
// own meta.
package tfidentity

import (
	"context"
	"fmt"
	"sync"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/identity"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

type copyKey struct {
	provider *tfschema.Provider
	id       string
}

var (
	copiesMu sync.Mutex
	copies   = map[copyKey]*tfschema.Provider{}
	// bases maps each copy back to the provider it was made from, so passing a
	// copy to Provider again does not copy it twice.
	bases = map[*tfschema.Provider]*tfschema.Provider{}
)

// Provider returns the Terraform provider to use for calls made under ctx.
//
// It returns provider unchanged when ctx has no identity.Resolver. Otherwise
// it returns a copy of provider whose meta authenticates with the credentials
// identity.ForObject recorded in ctx, and fails if there are none. The copy
// shares the schema and resource maps with provider. Copies are cached by
// credentials ID for the life of the process.
func Provider(ctx context.Context, provider *tfschema.Provider) (*tfschema.Provider, error) {
	if identity.ResolverFromContext(ctx) == nil {
		return provider, nil
	}
	creds, ok := identity.CredentialsFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no GCP credentials for this Terraform call: it is not made under a context from identity.ForObject")
	}
	return copyFor(provider, creds)
}

func copyFor(provider *tfschema.Provider, creds *identity.Credentials) (*tfschema.Provider, error) {
	copiesMu.Lock()
	defer copiesMu.Unlock()
	if base, ok := bases[provider]; ok {
		provider = base
	}
	key := copyKey{provider: provider, id: creds.ID}
	if cp, ok := copies[key]; ok {
		return cp, nil
	}
	base, ok := provider.Meta().(*transport_tpg.Config)
	if !ok {
		return nil, fmt.Errorf("terraform provider meta is %T, want *transport.Config", provider.Meta())
	}
	meta, err := base.CopyWithTokenSource(creds.TokenSource)
	if err != nil {
		return nil, fmt.Errorf("building terraform config for %q: %w", creds.ID, err)
	}
	cp := *provider
	cp.SetMeta(meta)
	copies[key] = &cp
	bases[&cp] = provider
	return &cp, nil
}
