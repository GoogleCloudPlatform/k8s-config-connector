// Copyright 2025 Google LLC
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

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Identity interface {
	// The external format of the resource identity.
	String() string

	// FromExternal parses a string-formatted external reference into an identity.
	FromExternal(ref string) error
}

type Resource interface {
	// GetIdentity gets the identity of a resource.
	GetIdentity(ctx context.Context, reader client.Reader) (Identity, error)

	// GetParentIdentity gets the parent identity of a resource.
	GetParentIdentity(ctx context.Context, reader client.Reader) (Identity, error)
}
