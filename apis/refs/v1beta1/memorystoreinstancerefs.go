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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MemorystoreInstanceServiceAttachmentRef struct {
	// TODO
	ServiceAttachmentExternal string `json:"serviceAttachmentExternal,omitempty"`
	// TODO
	MemorystoreInstanceName string `json:"memorystoreInstanceName,omitempty"`
	// TODO
	MemorystoreInstanceNamespace string `json:"memorystoreInstanceNamespace,omitempty"`
	// TODO
	MemorystoreInstanceServiceAttachmentIndex int `json:"memorystoreInstanceServiceAttachmentIndex,omitempty"`
}

func ResolveMemorystoreInstanceServiceAttachment(ctx context.Context, reader client.Reader, defaultNamespace string, ref *MemorystoreInstanceServiceAttachmentRef) error {
	if ref == nil {
		return nil
	}

	if ref.ServiceAttachmentExternal != "" {
		return nil
	}

	if ref.MemorystoreInstanceName == "" {
		return fmt.Errorf("must specify either MemorystoreInstanceName or ServiceAttachmentExternal on reference")
	}

	if ref.MemorystoreInstanceServiceAttachmentIndex < 0 {
		return fmt.Errorf("MemorystoreInstanceServiceAttachmentIndex must be non-negative integer")
	}

	key := types.NamespacedName{
		Namespace: ref.MemorystoreInstanceNamespace,
		Name:      ref.MemorystoreInstanceName,
	}
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}

	memorystoreInstance := &unstructured.Unstructured{}
	memorystoreInstance.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "memorystore.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "MemorystoreInstance",
	})
	if err := reader.Get(ctx, key, memorystoreInstance); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(memorystoreInstance.GroupVersionKind(), key)
		}
		return fmt.Errorf("error reading referenced MemorystoreInstance %v: %w", key, err)
	}

	// Read status.observedState.pscAttachmentDetails[MemorystoreInstanceServiceAttachmentIndex]
	// to retrieve the service attachment external.
	pscAttachmentDetails, found, err := unstructured.NestedSlice(memorystoreInstance.Object, "status", "observedState", "pscAttachmentDetails")
	if err != nil {
		return fmt.Errorf("getting status.observedState.pscAttachmentDetails[]: %w", err)
	}
	if !found {
		return k8s.NewReferenceNotFoundError(memorystoreInstance.GroupVersionKind(), key)
	}
	if len(pscAttachmentDetails) <= ref.MemorystoreInstanceServiceAttachmentIndex {
		return fmt.Errorf("MemorystoreInstanceServiceAttachmentIndex is out of range")
	}

	pscAttachmentDetail, ok := pscAttachmentDetails[ref.MemorystoreInstanceServiceAttachmentIndex].(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed getting status.observedState.pscAttachmentDetails[%d]", ref.MemorystoreInstanceServiceAttachmentIndex)
	}
	serviceAttachmentExternal, ok := pscAttachmentDetail["serviceAttachment"].(string)
	if !ok {
		return fmt.Errorf("failed getting status.observedState.pscAttachmentDetails[%d].serviceAttachment", ref.MemorystoreInstanceServiceAttachmentIndex)
	}

	ref.ServiceAttachmentExternal = serviceAttachmentExternal
	return nil
}
