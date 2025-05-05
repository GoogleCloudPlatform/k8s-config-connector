package configconnector

import (
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/golden"

	customizev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1alpha1"
	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
)

func TestGoldenConfigConnector(t *testing.T) {
	env := &envtest.Environment{
		CRDInstallOptions: envtest.CRDInstallOptions{
			Paths: []string{
				"../../../config/crd/bases",
			},
			ErrorIfPathMissing: true,
		},
	}

	v := golden.NewValidator(t, env, v1beta1.AddToScheme, customizev1alpha1.AddToScheme, customizev1beta1.AddToScheme, corev1.AddToScheme, appsv1.AddToScheme)

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
	reconciler, err := Add(v.Manager(), ccOptions)
	if err != nil {
		t.Fatalf("adding controller: %v", err)
	}

	trimCRDs := func(u *unstructured.Unstructured) {
		if u.GetKind() == "CustomResourceDefinition" {
			unstructured.RemoveNestedField(u.Object, "spec", "versions")
		}
	}
	v.Validate(reconciler.reconciler, golden.ValidateOptions{
		RewriteObjects: trimCRDs,
	})
}
