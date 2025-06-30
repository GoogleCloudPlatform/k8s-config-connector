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

package configconnector

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test"
)

func TestGoldenConfigConnector(t *testing.T) {
	h := test.NewHarness(t)

	rewriteObjects := func(u *unstructured.Unstructured) {
		// Don't output the bulk of the CRD data, just to keep the output small
		if u.GetKind() == "CustomResourceDefinition" {
			unstructured.RemoveNestedField(u.Object, "metadata", "creationTimestamp")
			unstructured.SetNestedField(u.Object, "(removed)", "spec")                    //nolint:errcheck
			unstructured.SetNestedField(u.Object, "(removed)", "metadata", "annotations") //nolint:errcheck
			unstructured.SetNestedField(u.Object, "(removed)", "metadata", "labels")      //nolint:errcheck
			// Note: the operator is setting the status on the CRDs, but it probably should not be doing so
		}

		// Note: the operator is doing some namespace manipulation
		if u.GetKind() == "Namespace" {
			unstructured.RemoveNestedField(u.Object, "metadata", "creationTimestamp")
			unstructured.RemoveNestedField(u.Object, "metadata", "resourceVersion")
			unstructured.RemoveNestedField(u.Object, "metadata", "uid")
			u.SetManagedFields(nil)
		}
	}

	v := h.KDPValidator(rewriteObjects)

	repoPath := h.RepoPath()

	imagePrefix := "foobar.local"

	var imageTransform *controllers.ImageTransform
	if imagePrefix != "" {
		imageTransform = controllers.NewImageTransform(imagePrefix)
	}
	ccOptions := &ReconcilerOptions{
		RepoPath:       repoPath,
		ImageTransform: imageTransform,
	}

	v.Validate(func(mgr manager.Manager) (*declarative.Reconciler, error) {
		r, err := newReconciler(mgr, ccOptions)
		return r.reconciler, err
	})
}
