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

package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	ResourceSettingsHashAnnotation = "cnrm.cloud.google.com/resource-settings-hash"
)

// ComputeResourceSettingsHash calculates a deterministic SHA-256 hash for the given resourceSettings.
// Returns an empty string if both ccSettings and cccSettings are nil or empty.
func ComputeResourceSettingsHash(ccSettings *corev1beta1.ResourceSettings, cccSettings *corev1beta1.ResourceSettings) (string, error) {
	if (ccSettings == nil || (len(ccSettings.Resources) == 0 && ccSettings.Mode == "")) &&
		(cccSettings == nil || (len(cccSettings.Resources) == 0 && cccSettings.Mode == "")) {
		return "", nil
	}

	payload := struct {
		CC  *corev1beta1.ResourceSettings `json:"cc,omitempty"`
		CCC *corev1beta1.ResourceSettings `json:"ccc,omitempty"`
	}{
		CC:  ccSettings,
		CCC: cccSettings,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling resourceSettings for hash computation: %w", err)
	}
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:]), nil
}

// ApplyResourceSettingsHashToPodTemplate sets or deletes the resourceSettings hash annotation
// on the Pod template of a StatefulSet unstructured object.
func ApplyResourceSettingsHashToPodTemplate(u *unstructured.Unstructured, ccSettings *corev1beta1.ResourceSettings, cccSettings *corev1beta1.ResourceSettings) error {
	hash, err := ComputeResourceSettingsHash(ccSettings, cccSettings)
	if err != nil {
		return err
	}

	annotations, _, err := unstructured.NestedStringMap(u.Object, "spec", "template", "metadata", "annotations")
	if err != nil {
		return fmt.Errorf("failed to get pod template annotations: %w", err)
	}

	if hash != "" {
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations[ResourceSettingsHashAnnotation] = hash
	} else if annotations != nil {
		delete(annotations, ResourceSettingsHashAnnotation)
	}

	if annotations != nil {
		if err := unstructured.SetNestedStringMap(u.Object, annotations, "spec", "template", "metadata", "annotations"); err != nil {
			return fmt.Errorf("failed to set pod template annotations: %w", err)
		}
	}
	return nil
}
