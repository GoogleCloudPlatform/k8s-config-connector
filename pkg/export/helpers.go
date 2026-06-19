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

package export

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

// SetProjectID sets the cnrm.cloud.google.com/project-id annotation on an unstructured object.
func SetProjectID(u *unstructured.Unstructured, projectID string) {
	if projectID == "" {
		return
	}
	annotations := u.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[k8s.ProjectIDAnnotation] = projectID
	u.SetAnnotations(annotations)
}

// SetLabels sets the metadata labels on an unstructured object from GCP labels.
func SetLabels(u *unstructured.Unstructured, labels map[string]string) {
	if len(labels) == 0 {
		return
	}
	u.SetLabels(labels)
}
