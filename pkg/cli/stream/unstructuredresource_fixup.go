// Copyright 2022 Google LLC
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

package stream

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const namePath = "metadata.name"

// UnstructuredResourceFixupStream will make any necessary edits to a resource to make sure it will work with K8s and KCC
// For example, ensuring the metadata.name is a K8s legal value (some GCP resource's allow values in 'name' with
// characters that are not allowed by K8s in metadata.name)
type UnstructuredResourceFixupStream struct {
	unstructStream UnstructuredStream
}

func NewUnstructuredResourceFixupStream(unstructuredStream UnstructuredStream) *UnstructuredResourceFixupStream {
	return &UnstructuredResourceFixupStream{
		unstructStream: unstructuredStream,
	}
}

func (s *UnstructuredResourceFixupStream) Next(ctx context.Context) (*unstructured.Unstructured, error) {
	u, err := s.unstructStream.Next(ctx)
	if err != nil {
		return nil, err
	}
	if err := defaultNameIfNotPresent(u); err != nil {
		return nil, err
	}
	if err := handleSpecialCases(u); err != nil {
		return nil, err
	}
	if err := ensureNameIsK8sLegal(u); err != nil {
		return nil, err
	}
	return u, nil
}

// if the resource does not have a name, fill it with the value in the resourceID field
// Not all resources support resourceID yet, temporarily fill in such resources with a generated name until the feature
// is completely rolled out
func defaultNameIfNotPresent(u *unstructured.Unstructured) error {
	_, ok, err := unstructured.NestedString(u.Object, strings.Split(namePath, ".")...)
	if err != nil {
		return fmt.Errorf("error retrieving %v from unstruct: %w", namePath, err)
	}
	if ok {
		return nil
	}
	val, ok, err := unstructured.NestedString(u.Object, strings.Split(k8s.ResourceIDFieldPath, ".")...)
	if err != nil {
		return fmt.Errorf("error retrieving %v from unstruct: %w", k8s.ResourceIDFieldPath, err)
	}
	if !ok {
		val = randomid.New().String()
	}
	if err := unstructured.SetNestedField(u.Object, val, strings.Split(namePath, ".")...); err != nil {
		return fmt.Errorf("error setting %v to '%v': %w", namePath, val, err)
	}
	return nil
}

// GCP resources names are not always valid K8s resource names, this function modifies the metadata.name to ensure it
// is always K8s legal.
func ensureNameIsK8sLegal(u *unstructured.Unstructured) error {
	name, ok, err := unstructured.NestedString(u.Object, strings.Split(namePath, ".")...)
	if err != nil {
		return fmt.Errorf("error retrieving %v from unstruct: %w", namePath, err)
	}
	if !ok {
		return fmt.Errorf("unexpected empty value for %v in resource", namePath)
	}
	// From https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions,
	// a custom resource's name must be a valid DNS subdomain name
	k8sName := k8s.ValueToDNSSubdomainName(name)
	if k8sName == name {
		return nil
	}
	if err := unstructured.SetNestedField(u.Object, k8sName, strings.Split(namePath, ".")...); err != nil {
		return fmt.Errorf("error setting %v to '%v': %w", namePath, k8sName, err)
	}
	return nil
}

// Some resources need additional handling
func handleSpecialCases(u *unstructured.Unstructured) error {
	resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
	switch u.GroupVersionKind().GroupKind() {
	// Table name is only unique in a dataset, avoid collisions
	case schema.GroupKind{Group: "bigquery.cnrm.cloud.google.com", Kind: "BigQueryTable"}:
		datasetID, _, _ := unstructured.NestedString(u.Object, "spec", "datasetRef", "external")
		tableID := resourceID

		if datasetID == "" {
			return fmt.Errorf("unexpected empty value for spec.datasetRef.external in BigQueryTable resource")
		}
		if tableID == "" {
			return fmt.Errorf("unexpected empty value for spec.resourceID in BigQueryTable resource")
		}
		u.SetName(datasetID + "-" + tableID)
	}
	return nil
}
