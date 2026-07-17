// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package identity

import (
	"context"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Identity interface {
	// The external format of the resource identity.
	String() string

	// FromExternal parses a string-formatted external reference into an identity.
	FromExternal(ref string) error
}

type IdentityV2 interface {
	Identity
	// Host returns the Host portion of the Cloud-Asset Inventory fully-qualified format (e.g. compute.googleapis.com)
	Host() string
}

// ServerGeneratedIdentity is implemented by Identity types for resources that have a server-generated ID.
type ServerGeneratedIdentity interface {
	IdentityV2
	HasIdentitySpecified() bool
}

type Resource interface {
	// GetIdentity gets the identity of a resource.
	GetIdentity(ctx context.Context, reader client.Reader) (Identity, error)
}

// StripReferencePrefixes trims optional scheme, the given host, and optional GVK version segments from a reference string.
func StripReferencePrefixes(ref string, host string) string {
	ref = strings.TrimPrefix(ref, "https://")
	ref = strings.TrimPrefix(ref, "http://")
	ref = strings.TrimPrefix(ref, "//")

	if strings.HasPrefix(ref, host+"/") {
		ref = strings.TrimPrefix(ref, host+"/")
		// Trim standard API versions
		for _, version := range []string{"v1/", "v1beta1/", "v1alpha1/", "v1beta/", "v2/"} {
			if strings.HasPrefix(ref, version) {
				ref = strings.TrimPrefix(ref, version)
				break
			}
		}
	}
	return ref
}
