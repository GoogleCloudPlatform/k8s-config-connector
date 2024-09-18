// Copyright 2024 Google LLC
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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Resolver = &BigQueryConnectionConnectionRef{}

type BigQueryConnectionConnectionRef struct {
	// A reference to an externally managed BigQueryConnectionConnection resource.
	// Should be in the format `projects/<projectID>/locations/<location>/connections/<connectionID>`.
	External string `json:"external,omitempty"`

	// The `name` of a `BigQueryConnectionConnection` resource.
	Name string `json:"name,omitempty"`
	// The `namespace` of a `BigQueryConnectionConnection` resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *BigQueryConnectionConnectionRef) ValidateExternal() error {
	if r.External == "" {
		return fmt.Errorf("external not specified")
	}
	r.External = strings.TrimPrefix(r.External, "/")
	tokens := strings.Split(r.External, "/")
	if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "connections" {
		return fmt.Errorf("format of BigQueryConnectionConnection external=%q was not known (use projects/<projectId>/locations/<location>/connections/<connectionID>)", r.External)
	}
	return nil
}

func (r *BigQueryConnectionConnectionRef) Resolve(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (string, error) {
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", fmt.Errorf("BigQueryConnectionConnection is not ready yet.")
	}
	return actualExternalRef, nil
}
