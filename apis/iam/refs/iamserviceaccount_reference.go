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

package iamrefs

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var IAMServiceAccountGVK = schema.GroupVersionKind{
	Group:   "iam.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "IAMServiceAccount",
}

var _ refsv1beta1.Ref = &IAMServiceAccountRef{}

// IAMServiceAccountRef is a reference to a GCP IAMServiceAccount.
// +k8s:deepcopy-gen=true
type IAMServiceAccountRef struct {
	// A reference to an externally managed IAMServiceAccount resource.
	// Should be in the format "{{serviceAccountID}}@{{projectID}}.iam.gserviceaccount.com".
	External string `json:"external,omitempty"`

	// The name of an IAMServiceAccount resource.
	Name string `json:"name,omitempty"`

	// The namespace of an IAMServiceAccount resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&IAMServiceAccountRef{}, nil)
}

func (r *IAMServiceAccountRef) GetGVK() schema.GroupVersionKind {
	return IAMServiceAccountGVK
}

func (r *IAMServiceAccountRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *IAMServiceAccountRef) GetExternal() string {
	return r.External
}

func (r *IAMServiceAccountRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *IAMServiceAccountRef) ValidateExternal(ref string) error {
	_, _, err := ParseIAMServiceAccountEmail(ref)
	return err
}

func (r *IAMServiceAccountRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.GetExternal() == "" {
		key := r.GetNamespacedName()
		if key.Namespace == "" {
			key.Namespace = defaultNamespace
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(IAMServiceAccountGVK)
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(IAMServiceAccountGVK, key)
			}
			return fmt.Errorf("reading referenced IAMServiceAccount %s: %w", key, err)
		}

		// Read status.email
		email, _, err := unstructured.NestedString(u.Object, "status", "email")
		if err != nil {
			return fmt.Errorf("reading status.email: %w", err)
		}
		if email == "" {
			return k8s.NewReferenceNotReadyError(IAMServiceAccountGVK, key)
		}
		r.SetExternal(email)
	}

	return r.ValidateExternal(r.GetExternal())
}

func ParseIAMServiceAccountEmail(email string) (account, project string, err error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("format of IAMServiceAccount reference external=%q was not known (use email address, i.e. {{serviceAccountID}}@{{projectID}}.iam.gserviceaccount.com)", email)
	}
	account = parts[0]
	rest := parts[1]
	suffix := ".iam.gserviceaccount.com"
	if !strings.HasSuffix(rest, suffix) {
		return "", "", fmt.Errorf("format of IAMServiceAccount reference external=%q was not known (expected suffix %s)", email, suffix)
	}
	project = strings.TrimSuffix(rest, suffix)
	if account == "" || project == "" {
		return "", "", fmt.Errorf("invalid IAMServiceAccount email: account or project is empty in %q", email)
	}
	return account, project, nil
}
