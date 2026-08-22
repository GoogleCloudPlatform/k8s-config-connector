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

package v1alpha1

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *TLSInspectionPolicyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name == "" {
		return nil
	}
	if r.External == "" && r.Name == "" {
		return fmt.Errorf("must specify either name or external")
	}
	return nil // Mock implementation to pass compilation
}
