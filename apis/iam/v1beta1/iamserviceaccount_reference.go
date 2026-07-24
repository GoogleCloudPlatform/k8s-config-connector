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

// IAMServiceAccountRef is a reference to a GCP IAMServiceAccount in standard relative resource name format.
type IAMServiceAccountRef struct {
	// A reference to an externally managed IAMServiceAccount resource.
	// Must be in the standard GCP relative resource name format: "projects/{{projectID}}/serviceAccounts/{{serviceAccountID}}"
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
	_, _, err := parseRelativeResourceName(ref)
	return err
}

func (r *IAMServiceAccountRef) ParseExternalToIdentity() (identity.Identity, error) {
	project, accountID, err := parseRelativeResourceName(r.External)
	if err != nil {
		return nil, err
	}
	return &IAMServiceAccountIdentity{
		Project: project,
		Account: accountID,
	}, nil
}

// AsEmail returns the service account in email format: {{serviceAccountID}}@{{projectID}}.iam.gserviceaccount.com
func (r *IAMServiceAccountRef) AsEmail() (string, error) {
	project, accountID, err := parseRelativeResourceName(r.External)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", accountID, project), nil
}

// AsRelativeResourceName returns the service account in relative resource name format: projects/{{projectID}}/serviceAccounts/{{serviceAccountID}}
func (r *IAMServiceAccountRef) AsRelativeResourceName() (string, error) {
	project, accountID, err := parseRelativeResourceName(r.External)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("projects/%s/serviceAccounts/%s", project, accountID), nil
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

		// Normalize to standard relative resource name
		project, accountID, err := parseEmailAddress(email)
		if err != nil {
			return fmt.Errorf("parsing service account email %q: %w", email, err)
		}
		r.SetExternal(fmt.Sprintf("projects/%s/serviceAccounts/%s", project, accountID))
	}

	return r.ValidateExternal(r.GetExternal())
}

var _ refs.Ref = &LegacyIAMServiceAccountRef{}

// LegacyIAMServiceAccountRef is a reference to a GCP IAMServiceAccount in legacy email format.
type LegacyIAMServiceAccountRef struct {
	// A reference to an externally managed IAMServiceAccount resource.
	// Must be in the legacy email format: "{{serviceAccountID}}@{{projectID}}.iam.gserviceaccount.com"
	External string `json:"external,omitempty"`

	// The name of an IAMServiceAccount resource.
	Name string `json:"name,omitempty"`

	// The namespace of an IAMServiceAccount resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&LegacyIAMServiceAccountRef{}, &IAMServiceAccount{})
}

func (r *LegacyIAMServiceAccountRef) GetGVK() schema.GroupVersionKind {
	return IAMServiceAccountGVK
}

func (r *LegacyIAMServiceAccountRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *LegacyIAMServiceAccountRef) GetExternal() string {
	return r.External
}

func (r *LegacyIAMServiceAccountRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *LegacyIAMServiceAccountRef) ValidateExternal(ref string) error {
	_, _, err := parseEmailAddress(ref)
	return err
}

func (r *LegacyIAMServiceAccountRef) ParseExternalToIdentity() (identity.Identity, error) {
	project, accountID, err := parseEmailAddress(r.External)
	if err != nil {
		return nil, err
	}
	return &IAMServiceAccountIdentity{
		Project: project,
		Account: accountID,
	}, nil
}

// AsEmail returns the service account in email format: {{serviceAccountID}}@{{projectID}}.iam.gserviceaccount.com
func (r *LegacyIAMServiceAccountRef) AsEmail() (string, error) {
	project, accountID, err := parseEmailAddress(r.External)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", accountID, project), nil
}

// AsRelativeResourceName returns the service account in relative resource name format: projects/{{projectID}}/serviceAccounts/{{serviceAccountID}}
func (r *LegacyIAMServiceAccountRef) AsRelativeResourceName() (string, error) {
	project, accountID, err := parseEmailAddress(r.External)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("projects/%s/serviceAccounts/%s", project, accountID), nil
}

func (r *LegacyIAMServiceAccountRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
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

		// Normalize to legacy email format
		project, accountID, err := parseEmailAddress(email)
		if err != nil {
			return fmt.Errorf("parsing service account email %q: %w", email, err)
		}
		r.SetExternal(fmt.Sprintf("%s@%s.iam.gserviceaccount.com", accountID, project))
	}

	return r.ValidateExternal(r.GetExternal())
}

func parseRelativeResourceName(ref string) (project, accountID string, err error) {
	ref = strings.TrimPrefix(ref, "https://iam.googleapis.com/")
	ref = strings.TrimPrefix(ref, "/")

	if !strings.HasPrefix(ref, "projects/") {
		return "", "", fmt.Errorf("invalid IAMServiceAccount relative resource name external=%q (must start with projects/)", ref)
	}

	parts := strings.Split(ref, "/")
	if len(parts) != 4 || parts[2] != "serviceAccounts" {
		return "", "", fmt.Errorf("format of IAMServiceAccount reference external=%q was not known (expected projects/{project}/serviceAccounts/{account})", ref)
	}

	project = parts[1]
	accountID = parts[3]
	if project == "" || accountID == "" {
		return "", "", fmt.Errorf("invalid IAMServiceAccount reference %q: project or account is empty", ref)
	}

	if strings.Contains(accountID, "@") {
		return "", "", fmt.Errorf("invalid IAMServiceAccount reference %q: account ID should not contain '@'", ref)
	}

	return project, accountID, nil
}

func parseEmailAddress(ref string) (project, accountID string, err error) {
	if strings.Contains(ref, "/") {
		return "", "", fmt.Errorf("invalid IAMServiceAccount email %q: email address cannot contain '/'", ref)
	}
	parts := strings.Split(ref, "@")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("format of IAMServiceAccount reference external=%q was not known (use email address, i.e. {{serviceAccountID}}@{{projectID}}.iam.gserviceaccount.com)", ref)
	}
	accountID = parts[0]
	rest := parts[1]
	suffix := ".iam.gserviceaccount.com"
	if !strings.HasSuffix(rest, suffix) {
		return "", "", fmt.Errorf("format of IAMServiceAccount reference external=%q was not known (expected suffix %s)", ref, suffix)
	}
	project = strings.TrimSuffix(rest, suffix)
	if accountID == "" || project == "" {
		return "", "", fmt.Errorf("invalid IAMServiceAccount email: account or project is empty in %q", ref)
	}
	return project, accountID, nil
}
