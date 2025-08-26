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

package common

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// This function should be called if the typed object has `spec.labels` field.
func ComputeLabels_ToProto(mapCtx *direct.MapContext, u *unstructured.Unstructured) map[string]string {
	var newLabels map[string]string
	specLabels, found, err := unstructured.NestedStringMap(u.Object, "spec", "labels")
	if err != nil {
		mapCtx.Errorf("retrieve %s: %s `spec.labels` field: %w", u.GroupVersionKind().Kind, u.GetName(), err)
		return nil
	}
	if specLabels != nil {
		newLabels = specLabels
	} else if found {
		newLabels = map[string]string{}
	} else {
		newLabels = u.GetLabels()
	}
	// No matter where the labels come from, sanitize them based on GCP label validation.
	newLabels = label.SanitizeGCPLabels(newLabels)
	return newLabels
}
