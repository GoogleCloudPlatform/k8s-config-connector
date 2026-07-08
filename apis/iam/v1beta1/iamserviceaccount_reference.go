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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &IAMServiceAccountRef{}

// IAMServiceAccountRef is a reference to a GCP IAMServiceAccount.
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
	refs.Register(&IAMServiceAccountRef{}, &IAMServiceAccount{})
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
	_, _, err := parseIAMServiceAccountEmail(ref)
	return err
}

func (r *IAMServiceAccountRef) ParseExternalToIdentity() (identity.Identity, error) {
	account, project, err := parseIAMServiceAccountEmail(r.External)
	if err != nil {
		return nil, err
	}
	return &IAMServiceAccountIdentity{
		Project: project,
		Account: account,
	}, nil
}

func (r *IAMServiceAccountRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.GetExternal() == "" {
		key := r.GetNamespacedName()
		if key.Namespace == "" {
			key.Namespace = defaultNamespace
		}
		u := &IAMServiceAccount{}
		if err := reader.Get(ctx, key, u); err != nil {
			if apierrors.IsNotFound(err) {
				return k8s.NewReferenceNotFoundError(schema.GroupVersionKind{
					Group:   "iam.cnrm.cloud.google.com",
					Version: "v1beta1",
					Kind:    "IAMServiceAccount",
				}, key)
			}
			return fmt.Errorf("reading referenced IAMServiceAccount %s: %w", key, err)
		}

		// Read status.email
		email := common.ValueOf(u.Status.Email)
		if email == "" {
			return k8s.NewReferenceNotReadyError(schema.GroupVersionKind{
				Group:   "iam.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "IAMServiceAccount",
			}, key)
		}
		r.SetExternal(email)
	}

	return r.ValidateExternal(r.GetExternal())
}

func parseIAMServiceAccountEmail(email string) (account, project string, err error) {
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
