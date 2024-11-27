// Copyright 2024 Google LLC
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
package parent

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Parent interface {
	// The external format of the Parent.
	String() string
	// Verify that the desired parent (from .spec) matches the actual parent in .status.externalRef.
	// This ensures the parent remains unchanged.
	// We currently don't enforce parent immutability using a webhook or CRD CEL due to legacy reasons.
	MatchActual(Parent) error
}

// ParentBuilder builds a Parent object from a ParentRef.
// - ParentRef is the Config Connector API reference for identifying a resource's logical parent.
// - The Parent object provides helper functions for parent-related logic in direct reconciliation.
type ParentBuilder interface {
	// Parent API reference builds its corresponding Parent object.
	Build(ctx context.Context, reader client.Reader, othernamespace string, parent Parent) error
}
