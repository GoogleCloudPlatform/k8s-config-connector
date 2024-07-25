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

package resolverefs

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ComputeTargetHTTPProxy struct {
	Project                  string
	Location                 string
	ComputeTargetHTTPProxyID string
}

func (c *ComputeTargetHTTPProxy) String() string {
	if c.Location == "global" {
		return fmt.Sprintf("projects/%s/global/targetHttpProxies/%s", c.Project, c.ComputeTargetHTTPProxyID)
	}
	return fmt.Sprintf("projects/%s/location/%s/targetHttpProxies/%s", c.Project, c.Location, c.ComputeTargetHTTPProxyID)
}

func ResolveTargetHTTPProxyRef(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeTargetHTTPProxyRef) (*ComputeTargetHTTPProxy, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on ComputeNetwork reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "targetHttpProxies" {
			return &ComputeTargetHTTPProxy{
				Project:                  tokens[1],
				ComputeTargetHTTPProxyID: tokens[4]}, nil
		} else if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "location" && tokens[4] == "targetHttpProxies" {
			return &ComputeTargetHTTPProxy{
				Project:                  tokens[1],
				Location:                 tokens[3],
				ComputeTargetHTTPProxyID: tokens[5]}, nil
		}
		return nil, fmt.Errorf("format of ComputeTargetHTTPProxy external=%q was not known (use projects/<projectId>/global/targetHttpProxies/<proxyId> or projects/<projectId>/location/<location>/targetHttpProxies/<proxyId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on ComputeTargetHTTPProxy reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	computeTargetHTTPProxy := &unstructured.Unstructured{}
	computeTargetHTTPProxy.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetHTTPProxy",
	})
	if err := reader.Get(ctx, key, computeTargetHTTPProxy); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced ComputeTargetHTTPProxy %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced ComputeTargetHTTPProxy %v: %w", key, err)
	}

	computeTargetHTTPProxyID, err := GetResourceID(computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	computeTargetHTTPProxyProjectID, err := GetProjectID(ctx, reader, computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	computeTargetHTTPProxyLocation, err := GetLocation(computeTargetHTTPProxy)
	if err != nil {
		return nil, err
	}

	return &ComputeTargetHTTPProxy{
		Project:                  computeTargetHTTPProxyProjectID,
		Location:                 computeTargetHTTPProxyLocation,
		ComputeTargetHTTPProxyID: computeTargetHTTPProxyID,
	}, nil

}
