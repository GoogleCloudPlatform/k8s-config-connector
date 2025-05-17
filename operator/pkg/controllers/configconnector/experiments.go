package configconnector

import (
	"context"
	"fmt"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
)

func (r *Reconciler) transformForExperiments() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		cc, ok := o.(*corev1beta1.ConfigConnector)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnector, but it was of type %T", o)
		}

		if err := r.applyExperiments(ctx, cc, m); err != nil {
			return fmt.Errorf("error applying experiment transforms: %w", err)
		}
		return nil
	}
}

func (r *Reconciler) applyExperiments(ctx context.Context, cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	log := log.FromContext(ctx)

	for _, experiment := range cc.Spec.Experiments {
		key := strings.ToLower(experiment)
		switch key {
		case "no-monitoring-services":
			if err := r.applyExperimentNoMonitoringServices(ctx, cc, m); err != nil {
				return err
			}

		default:
			log.Info("ignoring unknown experiment", "key", key)
			// TODO: add to status?
		}
	}

	return nil
}

func (r *Reconciler) applyExperimentNoMonitoringServices(ctx context.Context, cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	// log := log.FromContext(ctx)

	transformed := make([]*manifest.Object, 0, len(m.Items))
	for _, obj := range m.Items {
		keep := true
		if obj.Kind == "Service" && obj.Group == "" {
			switch obj.GetName() {
			case "cnrm-resource-stats-recorder-service":
				keep = false
			case "cnrm-deletiondefender":
				keep = true // deletion defender is a webhook; when a CRD is deleted it will go mark all the resources as abandoned.
			}
		}

		if keep {
			transformed = append(transformed, obj)
		}
	}
	m.Items = transformed
	return nil
}
