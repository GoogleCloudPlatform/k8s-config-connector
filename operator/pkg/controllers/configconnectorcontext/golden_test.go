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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/golden"

	customizev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1alpha1"
	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
)

func TestGoldenConfigConnectorContext(t *testing.T) {
	env := &envtest.Environment{
		CRDInstallOptions: envtest.CRDInstallOptions{
			Paths:              paths.GetOperatorCRDsPaths(),
			ErrorIfPathMissing: true,
		},
	}

	rewriteObjects := func(u *unstructured.Unstructured) {
		// Nothing to rewrite currently
	}

	goldenOptions := golden.ValidatorOptions{
		RewriteObjects:     rewriteObjects,
		EnvtestEnvironment: env,
		ManagerOptions: manager.Options{
			Metrics: metricsserver.Options{BindAddress: "0"},
		},
	}
	goldenOptions.WithSchema(
		v1beta1.AddToScheme,
		customizev1alpha1.AddToScheme,
		customizev1beta1.AddToScheme,
		corev1.AddToScheme,
		appsv1.AddToScheme,
		apiextensionsv1.AddToScheme,
	)

	v := golden.NewValidator(t, goldenOptions)

	repoPath := "../../../channels"

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
