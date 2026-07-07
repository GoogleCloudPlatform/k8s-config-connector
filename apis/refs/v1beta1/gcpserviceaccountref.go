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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
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

// IAMServiceAccountRef is the reference type for brownfield/legacy resources.
// It uses/normalizes to the Service Account email address in the "External" field for backwards compatibility.
// For newly created greenfield resources, please use ServiceAccountRef instead.
type IAMServiceAccountRef struct {
	/* The `email` field of an `IAMServiceAccount` resource. */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

func (r *IAMServiceAccountRef) Resolve(ctx context.Context, reader client.Reader, src client.Object) error {
	if r == nil {
		return nil
	}

	serviceAccountInfo, err := resolveServiceAccount(ctx, reader, src, r.Name, r.Namespace, r.External)
	if err != nil {
		return err
	}
	*r = IAMServiceAccountRef{External: serviceAccountInfo.External}
	return nil
}

// ServiceAccountRef is the reference type for greenfield resources.
// It advocates and normalizes to the relative resource name (i.e. "projects/{{project}}/serviceAccounts/{{email}}") in the "External" field.
// It also supports the legacy Service Account email address format (e.g., "email@project.iam.gserviceaccount.com").
type ServiceAccountRef struct {
	// A reference to an externally managed IAMServiceAccount resource.
	// Recommended format is the relative resource name: "projects/{{projectID}}/serviceAccounts/{{serviceAccountID}}".
	// The legacy format (email) is also supported.
	External string `json:"external,omitempty"`

	// The name of an IAMServiceAccount resource.
	Name string `json:"name,omitempty"`

	// The namespace of an IAMServiceAccount resource.
	Namespace string `json:"namespace,omitempty"`
}

var _ Ref = &ServiceAccountRef{}

func (r *ServiceAccountRef) GetGVK() schema.GroupVersionKind {
	return IAMServiceAccountGVK
}

func (r *ServiceAccountRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ServiceAccountRef) GetExternal() string {
	return r.External
}

func (r *ServiceAccountRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ServiceAccountRef) ValidateExternal(ref string) error {
	// Support both legacy email format and standard relative resource name
	if strings.Contains(ref, "@") && !strings.Contains(ref, "/") {
		return nil
	}
	id := &serviceAccountIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ServiceAccountRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		statusName, _, err := unstructured.NestedString(u.Object, "status", "name")
		if err == nil && statusName != "" {
			return statusName
		}
		email, _, err := unstructured.NestedString(u.Object, "status", "email")
		if err == nil && email != "" {
			projectID, err := ResolveProjectID(ctx, reader, u)
			if err == nil && projectID != "" {
				return fmt.Sprintf("projects/%s/serviceAccounts/%s", projectID, email)
			}
		}
		return ""
	}
	return NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

type MetricsGcpServiceAccountRef struct {
	/* The Email of the Google Cloud Service Account (GSA) used for exporting Config Sync metrics to Cloud Monitoring. The GSA should have the Monitoring Metric Writer(roles/monitoring.metricWriter) IAM role. The Kubernetes ServiceAccount `default` in the namespace `config-management-monitoring` should be bound to the GSA. Allowed value: The `email` field of an `IAMServiceAccount` resource. */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

func (r *MetricsGcpServiceAccountRef) Resolve(ctx context.Context, reader client.Reader, src client.Object) error {
	if r == nil {
		return nil
	}

	serviceAccountInfo, err := resolveServiceAccount(ctx, reader, src, r.Name, r.Namespace, r.External)
	if err != nil {
		return err
	}
	*r = MetricsGcpServiceAccountRef{External: serviceAccountInfo.External}
	return nil
}

type serviceAccountInfo struct {
	External string
}

func resolveServiceAccount(ctx context.Context, reader client.Reader, src client.Object, name, namespace, external string) (*serviceAccountInfo, error) {
	if external != "" {
		if name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on an IAMServiceAccount reference")
		}

		if strings.Contains(external, "@") {
			return &serviceAccountInfo{External: external}, nil
		}
		return nil, fmt.Errorf("format of IAMServiceAccount reference external=%q was not known (use email address)", external)
	}

	if name == "" {
		return nil, fmt.Errorf("must specify either name or external on an IAMServiceAccount reference")
	}

	key := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computenetwork := &unstructured.Unstructured{}
	computenetwork.SetGroupVersionKind(IAMServiceAccountGVK)
	if err := reader.Get(ctx, key, computenetwork); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced IAMServiceAccount %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced IAMServiceAccount %v: %w", key, err)
	}

	email, _, err := unstructured.NestedString(computenetwork.Object, "status", "email")
	if err != nil {
		return nil, fmt.Errorf("reading status.email from IAMServiceAccount %v: %w", key, err)
	}
	// if the status.email not populated, should we construct the email from spec.resourceID or metadata.name.
	if email == "" {
		return nil, fmt.Errorf("status.email is empty from IAMServiceAccount %v, expected not-empty", key)
	}

	return &serviceAccountInfo{External: email}, nil
}

type serviceAccountIdentity struct {
	Project string
	Account string
}

func (p *serviceAccountIdentity) String() string {
	return serviceAccountFormat.ToString(*p)
}

func (p *serviceAccountIdentity) FromExternal(ref string) error {
	parsed, match, err := serviceAccountFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of serviceAccountIdentity external=%q was not known (use %s): %w", ref, serviceAccountFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of serviceAccountIdentity external=%q was not known (use %s)", ref, serviceAccountFormat.CanonicalForm())
	}
	*p = *parsed
	return nil
}

func (p *serviceAccountIdentity) Host() string {
	return serviceAccountFormat.Host()
}

var (
	serviceAccountFormatTemplate = "projects/{project}/serviceAccounts/{account}"
	serviceAccountFormat         = gcpurls.Template[serviceAccountIdentity](
		"iam.googleapis.com",
		serviceAccountFormatTemplate,
	)
)
