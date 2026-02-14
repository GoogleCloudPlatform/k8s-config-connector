// Copyright 2025 Google LLC
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

// MemorystoreInstanceRef defines the resource reference to MemorystoreInstance, which "External" field
// holds the GCP identifier for the KRM object.
type MemorystoreInstanceRef struct {
	// A reference to an externally managed MemorystoreInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	// +kubebuilder:validation:Pattern=^projects\/[^/]+\/locations\/[^/]+\/instances\/[^/]+$
	External string `json:"external,omitempty"`

	// The name of a MemorystoreInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MemorystoreInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

func ResolveMemorystoreInstance(ctx context.Context, reader client.Reader, src client.Object, ref *MemorystoreInstanceRef) (*MemorystoreInstanceRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on memorystoreinstance reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
			return &MemorystoreInstanceRef{
				External: fmt.Sprintf("projects/%s/locations/%s/instances/%s", tokens[1], tokens[3], tokens[5]),
			}, nil
		}
		return nil, fmt.Errorf("format of memorystoreinstance external=%q was not known (use projects/<projectId>/locations/<location>/instances/<instanceId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on memorystoreinstance reference")
	}
	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	instanceObj := &unstructured.Unstructured{}
	instanceObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "memorystore.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "MemorystoreInstance",
	})
	if err := reader.Get(ctx, key, instanceObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced MemorystoreInstance %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced MemorystoreInstance %v: %w", key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(instanceObj.Object, "status", "externalRef")
	if err != nil {
		return nil, fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return nil, k8s.NewReferenceNotReadyError(instanceObj.GroupVersionKind(), key)
	}
	return &MemorystoreInstanceRef{
		External: actualExternalRef,
	}, nil
}

// MemorystoreInstanceBackupRef defines the resource reference to MemorystoreInstanceBackup, which "External" field
// holds the GCP identifier for the KRM object.
type MemorystoreInstanceBackupRef struct {
	// A reference to an externally managed MemorystoreInstanceBackup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/backupCollections/{{backupCollectionName}}/backups/{{backupID}}".
	// +kubebuilder:validation:Pattern=^projects\/[^/]+\/locations\/[^/]+\/backupCollections\/[^/]+\/backups\/[^/]+$
	External string `json:"external,omitempty"`

	// The name of a MemorystoreInstanceBackup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MemorystoreInstanceBackup resource.
	Namespace string `json:"namespace,omitempty"`
}

func ResolveMemorystoreInstanceBackup(ctx context.Context, reader client.Reader, src client.Object, ref *MemorystoreInstanceBackupRef) (*MemorystoreInstanceBackupRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on memorystoreinstance reference")
		}

		tokens := strings.Split(ref.External, "/")
		if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupCollections" && tokens[6] == "backups" {
			return &MemorystoreInstanceBackupRef{
				External: fmt.Sprintf("projects/%s/locations/%s/backupCollections/%s/backups/%s", tokens[1], tokens[3], tokens[5], tokens[7]),
			}, nil
		}
		return nil, fmt.Errorf("format of memorystoreinstance external=%q was not known (use projects/<projectId>/locations/<location>/backupCollections/<backupCollectionName>/backups/<backupId>)", ref.External)
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on memorystoreinstancebackup reference")
	}
	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	backupObj := &unstructured.Unstructured{}
	backupObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "memorystore.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "MemorystoreInstanceBackup",
	})
	if err := reader.Get(ctx, key, backupObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced MemorystoreInstanceBackup %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced MemorystoreInstanceBackup %v: %w", key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(backupObj.Object, "status", "externalRef")
	if err != nil {
		return nil, fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return nil, k8s.NewReferenceNotReadyError(backupObj.GroupVersionKind(), key)
	}
	return &MemorystoreInstanceBackupRef{
		External: actualExternalRef,
	}, nil
}
