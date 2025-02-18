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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type IAMServiceAccountRef struct {
	/* The `email` field of an `IAMServiceAccount` resource. */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
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

	sa := &unstructured.Unstructured{}
	sa.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "iam.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "IAMServiceAccount",
	})
	if err := reader.Get(ctx, key, sa); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced IAMServiceAccount %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced IAMServiceAccount %v: %w", key, err)
	}
	resource, err := k8s.NewResource(sa)
	if err != nil {
		return nil, fmt.Errorf("error converting unstructured to resource: %w", err)
	}
	if !k8s.IsResourceReady(resource) {
		return nil, k8s.NewReferenceNotReadyError(sa.GroupVersionKind(), key)
	}

	email, _, err := unstructured.NestedString(sa.Object, "status", "email")
	if err != nil {
		return nil, fmt.Errorf("reading status.email from IAMServiceAccount %v: %w", key, err)
	}
	// if the status.email not populated, should we construct the email from spec.resourceID or metadata.name.
	if email == "" {
		return nil, fmt.Errorf("status.email is empty from IAMServiceAccount %v, expected not-empty", key)
	}

	return &serviceAccountInfo{External: email}, nil
}
