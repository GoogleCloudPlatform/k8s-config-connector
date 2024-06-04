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

package references

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ComputeNetwork struct {
	Project          string
	ComputeNetworkID string
}

func (c *ComputeNetwork) String() string {
	return fmt.Sprintf("projects/%s/global/networks/%s", c.Project, c.ComputeNetworkID)
}

func ResolveComputeNetwork(ctx context.Context, reader client.Reader, src client.Object, ref *v1beta1.ComputeNetworkRef) (*ComputeNetwork, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on computenetwork reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "networks" {
			return &ComputeNetwork{
				Project:          tokens[1],
				ComputeNetworkID: tokens[4]}, nil
		}
		return nil, fmt.Errorf("format of computenetwork external=%q was not known (use projects/<projectId>/global/networks/<networkid>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on computenetwork reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computenetwork := &unstructured.Unstructured{}
	computenetwork.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	})
	if err := reader.Get(ctx, key, computenetwork); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeNetwork %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeNetwork %v: %w", key, err)
	}

	computenetworkID, _, err := unstructured.NestedString(computenetwork.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from ComputeNetwork %v: %w", key, err)
	}
	if computenetworkID == "" {
		computenetworkID = computenetwork.GetName()
	}

	projectID := ""
	// TODO: where is the const var.
	annotations := computenetwork.GetAnnotations()
	fromAnnotation, ok := annotations["cnrm.cloud.google.com/project-id"]
	if ok {
		projectID = fromAnnotation
	} else if computenetwork.GetNamespace() != "" {
		projectID = computenetwork.GetNamespace()
	}
	return &ComputeNetwork{
		Project:          projectID,
		ComputeNetworkID: computenetworkID,
	}, nil
}
