// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (obj *ArtifactRegistryRepository) GetIdentity(ctx context.Context, reader client.Reader) (string, error) {
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", err
	}
	projectID, err := v1beta1.ResolveProjectID(ctx, reader, &unstructured.Unstructured{Object: u})
	if err != nil {
		return "", err
	}
	location := obj.Spec.Location
	if location == "" {
		return "", fmt.Errorf("spec.location is required")
	}
	resourceID := ""
	if obj.Spec.ResourceID != nil {
		resourceID = *obj.Spec.ResourceID
	}
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", projectID, location, resourceID), nil
}

func (obj *ArtifactRegistryRepository) ParseIdentity(id string) (string, string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 6 || parts[0] != "projects" || parts[2] != "locations" || parts[4] != "repositories" {
		return "", "", "", fmt.Errorf("format of ArtifactRegistryRepository identity is expected to be projects/{project}/locations/{location}/repositories/{repository}, but got %s", id)
	}
	return parts[1], parts[3], parts[5], nil
}
