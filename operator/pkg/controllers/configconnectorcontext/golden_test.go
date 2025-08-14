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

package configconnectorcontext

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test"
)

func TestGoldenConfigConnectorContext(t *testing.T) {
	h := test.NewHarness(t)

	rewriteObjects := func(u *unstructured.Unstructured) {
		// Nothing to rewrite currently
	}

	imagePrefix := "foobar.local"

	var imageTransform *controllers.ImageTransform
	if imagePrefix != "" {
		imageTransform = controllers.NewImageTransform(imagePrefix)
	}
	ccOptions := &ReconcilerOptions{
		RepoPath:       h.RepoPath(),
		ImageTransform: imageTransform,
	}

	v := h.KDPValidator(rewriteObjects)

	v.Validate(func(mgr manager.Manager) (*declarative.Reconciler, error) {
		r, err := newReconciler(mgr, ccOptions)
		return r.reconciler, err
	})
}
