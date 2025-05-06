package configconnectorcontext

import (
	"context"
	"fmt"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cluster"
)

func (r *Reconciler) transformForExperiments(kube client.Client) declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnector, but it was of type %T", o)
		}

		cc := &corev1beta1.ConfigConnector{}
		ccName := client.ObjectKey{
			Namespace: "", Name: corev1beta1.ConfigConnectorAllowedName,
		}
		if err := kube.Get(ctx, ccName, cc); err != nil {
			return fmt.Errorf("getting cluster ConfigConnector object %v: %w", ccName, err)
		}

		if err := r.applyExperiments(ctx, cc, ccc, m); err != nil {
			return fmt.Errorf("error applying experiment transforms: %w", err)
		}
		return nil
	}
}

func (r *Reconciler) applyExperiments(ctx context.Context, cc *corev1beta1.ConfigConnector, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	log := log.FromContext(ctx)

	for _, experiment := range cc.Spec.Experiments {
		key := strings.ToLower(experiment)
		switch key {
		case "no-monitoring-services":
			if err := r.applyExperimentNoMonitoringServices(ctx, cc, ccc, m); err != nil {
				return err
			}

		default:
			log.Info("ignoring unknown experiment", "key", key)
			// TODO: add to status?
		}
	}

	return nil
}

func (r *Reconciler) applyExperimentNoMonitoringServices(ctx context.Context, cc *corev1beta1.ConfigConnector, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	log := log.FromContext(ctx)

	nsID, err := cluster.GetNamespaceID(ctx, k8s.OperatorNamespaceIDConfigMapNN, r.client, ccc.Namespace)
	if err != nil {
		return fmt.Errorf("error getting namespace id for namespace %v: %w", ccc.Namespace, err)
	}
	
	for _, obj := range m.Items {
		if obj.Kind == "Service" && obj.Group == "" {
			useHeadless := false
			switch obj.GetName() {
			case "cnrm-manager-" + nsID:
				useHeadless = true
			default:
				log.Info("ignoring unknown service %q in use-headless-service-for-monitoring", obj.GetName())
				useHeadless = false
			}

			if useHeadless {
				if err := obj.SetNestedField("None", "spec", "clusterIP"); err != nil {
					return fmt.Errorf("setting spec.clusterIP=None: %w", err)
				}
			}
		}
	}

	return nil
}
