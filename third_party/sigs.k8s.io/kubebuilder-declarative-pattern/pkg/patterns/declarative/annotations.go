package declarative

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

// AddAnnotations returns an ObjectTransform that adds annotations to all the objects
func AddAnnotations(annotations map[string]string) ObjectTransform {
	return func(ctx context.Context, o DeclarativeObject, manifest *manifest.Objects) error {
		log := log.Log
		for _, o := range manifest.Items {
			log.WithValues("object", o).WithValues("annotations", annotations).V(1).Info("add annotations to object")
			o.AddAnnotations(annotations)
		}

		return nil
	}
}
