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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BigQueryConnectionConnectionRef{}

// BigQueryConnectionConnectionRef is a reference to a BigQueryConnectionConnection.
type BigQueryConnectionConnectionRef struct {
	// A reference to an externally managed BigQueryConnectionConnection resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/connections/{{connectionID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryConnectionConnection resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryConnectionConnection resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BigQueryConnectionConnectionRef{})
}

func (r *BigQueryConnectionConnectionRef) GetGVK() schema.GroupVersionKind {
	return BigQueryConnectionConnectionGVK
}

func (r *BigQueryConnectionConnectionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigQueryConnectionConnectionRef) GetExternal() string {
	return r.External
}

func (r *BigQueryConnectionConnectionRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *BigQueryConnectionConnectionRef) ValidateExternal(ref string) error {
	id := &BigQueryConnectionConnectionIdentity{}
	external := strings.TrimPrefix(ref, "/")
	if err := id.FromExternal(external); err != nil {
		return err
	}
	return nil
}

func (r *BigQueryConnectionConnectionRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BigQueryConnectionConnectionIdentity{}
	external := strings.TrimPrefix(r.External, "/")
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BigQueryConnectionConnectionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*BigQueryConnectionConnection](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromBigQueryConnectionConnectionSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NewBigQueryConnectionConnectionRef builds a BigQueryConnectionConnectionRef from the ConfigConnector BigQueryConnectionConnection object.
func NewBigQueryConnectionConnectionRef(ctx context.Context, reader client.Reader, obj *BigQueryConnectionConnection) (*BigQueryConnectionConnectionRef, error) {
	id := &BigQueryConnectionConnectionRef{}

	specIdentity, err := getIdentityFromBigQueryConnectionConnectionSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Validate status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &BigQueryConnectionConnectionIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		// Validate spec parent and resourceID field if the resource is already reconciled with a GCP Connection resource.
		if specIdentity.Project != statusIdentity.Project {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", specIdentity.Project, statusIdentity.Project)
		}
		if specIdentity.Location != statusIdentity.Location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", specIdentity.Location, statusIdentity.Location)
		}
		if specIdentity.Connection != "" && specIdentity.Connection != statusIdentity.Connection {
			return nil, fmt.Errorf("cannot reset `spec.resourceID` to %s, since it has already acquired the Connection %s",
				specIdentity.Connection, statusIdentity.Connection)
		}
		id.External = externalRef
		return id, nil
	}
	id.External = specIdentity.String()
	return id, nil
}

// Parent returns the parent path in external form ("projects/{project}/locations/{location}").
func (r *BigQueryConnectionConnectionRef) Parent() (string, error) {
	if r.External != "" {
		id := &BigQueryConnectionConnectionIdentity{}
		external := strings.TrimPrefix(r.External, "/")
		if err := id.FromExternal(external); err != nil {
			return "", err
		}
		return "projects/" + id.Project + "/locations/" + id.Location, nil
	}
	return "", fmt.Errorf("BigQueryConnectionConnectionRef not normalized to External form or not created from `New()`")
}

// ConnectionID returns the connection ID, a boolean indicating whether the connection ID is specified by user (or generated by service), and an error.
func (r *BigQueryConnectionConnectionRef) ConnectionID() (string, bool, error) {
	if r.External != "" {
		id := &BigQueryConnectionConnectionIdentity{}
		external := strings.TrimPrefix(r.External, "/")
		if err := id.FromExternal(external); err != nil {
			return "", false, err
		}
		return id.Connection, id.Connection != "", nil
	}
	return "", false, fmt.Errorf("BigQueryConnectionConnectionRef not normalized to External form or not created from `New()`")
}

// NormalizedExternal provision the "External" value.
// Kept for backward compatibility with older callers.
func (r *BigQueryConnectionConnectionRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

// BigQueryConnectionServiceAccountRef specifies a reference to a service account generated by the BigQueryConnectionConnection resource.
// Supported connection types are: cloudSQL, spark, and cloudResource.
type BigQueryConnectionServiceAccountRef struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
	// Type field specifies the connection type of the BigQueryConnectionConnection resource, whose service account is to be bound to the role.
	// +kubebuilder:validation:Enum=spark;cloudSQL;cloudResource
	Type string `json:"type"`
}

// ResolveServiceAccountID resolves the service account ID from a BigQueryConnection resource reference.
// Supported connection types are: cloudSQL, spark, and cloudResource.
// This function extracts the service account ID from the referenced BigQueryConnection resource based on the connection type.
//
// Input parameters:
//   - namespace: Default namespace to use if not specified in the reference. This is usually the namespace of the resource referencing the BigQueryConnection
//   - ref: Reference to the BigQueryConnection (Name, Namespace, Type)
func ResolveServiceAccountID(ctx context.Context, reader client.Reader, namespace string, ref *BigQueryConnectionServiceAccountRef) (string, error) {
	key := types.NamespacedName{
		Name:      ref.Name,
		Namespace: ref.Namespace,
	}
	if key.Namespace == "" { // use the namespace of the IAM resource if a namespace is not provided in the reference
		key.Namespace = namespace
	}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(BigQueryConnectionConnectionGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", BigQueryConnectionConnectionGVK, key, err)

	}

	// a map of supported connection types to the corresponding status paths to service acocunt ID
	serviceAccountPaths := map[string][]string{
		"cloudSQL":      {"status", "observedState", "cloudSQL", "serviceAccountID"},
		"spark":         {"status", "observedState", "spark", "serviceAccountID"},
		"cloudResource": {"status", "observedState", "cloudResource", "serviceAccountID"},
	}
	path, supported := serviceAccountPaths[ref.Type]
	if !supported {
		return "", fmt.Errorf("invalid bigqueryconnectionconnectionRef.type '%s'. Supported types are: cloudSQL, spark, cloudResource", ref.Type)
	}

	sa, found, err := unstructured.NestedString(u.Object, path...)
	if err != nil {
		return "", fmt.Errorf("failed to access serviceAccountID field in BigQueryConnection %s/%s: %w", key.Namespace, key.Name, err)
	}
	if !found {
		pathStr := strings.Join(path, ".")
		return "", fmt.Errorf("BigQueryConnection %s is not ready - field '%s' is missing", key.String(), pathStr)
	}

	return sa, nil
}
