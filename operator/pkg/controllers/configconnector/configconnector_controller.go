// Copyright 2022 Google LLC
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
	"context"
	"fmt"
	"time"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	cnrmmanifest "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/preflight"
	corekcck8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	admissionregistration "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const controllerName = "configconnector-controller"

// ConfigConnectorReconciler reconciles a ConfigConnector object.

// ConfigConnectorReconciler watches 'ConfigConnector' kind and is responsible for managing the lifecycle of KCC resource CRDs and other shared components like webhook, deletion defender, recorder.
// If it’s configured to run KCC in cluster mode, ConfigConnectorReconciler also deploys the global controller manager workload;
// If it's configured to run KCC in namespaced mode, ConfigConnectorReconciler ensures the global controller manager workload not existing.
type ConfigConnectorReconciler struct {
	reconciler *declarative.Reconciler
	client     client.Client
	recorder   record.EventRecorder
	labelMaker declarative.LabelMaker
	log        logr.Logger
}

func Add(mgr ctrl.Manager, repoPath string) error {
	r, err := newReconciler(mgr, repoPath)
	if err != nil {
		return err
	}

	// Create a new ConfigConnector controller.
	obj := &corev1beta1.ConfigConnector{}
	_, err = builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: 1}).
		For(obj, builder.OnlyMetadata).
		Build(r)
	if err != nil {
		return err
	}

	return nil
}

func newReconciler(mgr ctrl.Manager, repoPath string) (*ConfigConnectorReconciler, error) {
	repo := cnrmmanifest.NewLocalRepository(repoPath)
	manifestLoader := cnrmmanifest.NewManifestLoader(repo)
	preflight := preflight.NewCompositePreflight([]declarative.Preflight{
		preflight.NewNameChecker(mgr.GetClient(), k8s.ConfigConnectorAllowedName),
		preflight.NewUpgradeChecker(mgr.GetClient(), repo),
	})

	r := &ConfigConnectorReconciler{
		reconciler: &declarative.Reconciler{},
		client:     mgr.GetClient(),
		recorder:   mgr.GetEventRecorderFor(controllerName),
		labelMaker: declarative.SourceLabel(mgr.GetScheme()),
		log:        ctrl.Log.WithName(controllerName),
	}

	err := r.reconciler.Init(mgr, &corev1beta1.ConfigConnector{},
		declarative.WithLabels(r.labelMaker),
		declarative.WithPreserveNamespace(),
		declarative.WithManifestController(manifestLoader),
		declarative.WithOwner(declarative.SourceAsOwner),
		declarative.WithObjectTransform(r.transformForClusterMode()),
		declarative.WithObjectTransform(r.handleConfigConnectorLifecycle()),
		declarative.WithStatus(&declarative.StatusBuilder{
			PreflightImpl: preflight,
		}),
	)
	return r, err
}

func (r *ConfigConnectorReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	r.log.Info("reconciling the ConfigConnector object", "name", req.Name)
	_, err := controllers.GetConfigConnector(ctx, r.client, req.NamespacedName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("ConfigConnector not found in API server; skipping the reconciliation", "name", req.Name)
			return reconcile.Result{}, nil
		}
	}
	_, reconciliationErr := r.reconciler.Reconcile(ctx, req)
	if reconciliationErr != nil {
		if err := r.handleReconcileFailed(ctx, req.NamespacedName, reconciliationErr); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, reconciliationErr
	}
	r.log.Info("successfully finished reconcile", "ConfigConnector", req.Name)
	return reconcile.Result{RequeueAfter: corekcck8s.MeanReconcileReenqueuePeriod}, r.handleReconcileSucceeded(ctx, req.NamespacedName)
}

func (r *ConfigConnectorReconciler) handleReconcileFailed(ctx context.Context, nn types.NamespacedName, reconcileErr error) error {
	cc, err := controllers.GetConfigConnector(ctx, r.client, nn)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("ConfigConnector not found in API server; skipping the handling of failed reconciliation", "name", nn.Name, "reconcile error", reconcileErr)
			return nil
		}
		r.log.Info("error getting ConfigConnector object", "name", nn.Name, "reconcile error", reconcileErr)
		return fmt.Errorf("error getting ConfigConnector object %v: %w", nn.Name, err)
	}
	msg := fmt.Sprintf(k8s.ReconcileErrMsgTmpl, reconcileErr)
	r.recordEvent(cc, corev1.EventTypeWarning, k8s.UpdateFailed, msg)
	cc.SetCommonStatus(v1alpha1.CommonStatus{
		Healthy: false,
		Errors:  []string{msg},
	})
	return r.updateConfigConnectorStatus(ctx, cc)
}

func (r *ConfigConnectorReconciler) handleReconcileSucceeded(ctx context.Context, nn types.NamespacedName) error {
	cc, err := controllers.GetConfigConnector(ctx, r.client, nn)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("ConfigConnector not found in API server; skipping the handling of successful reconciliation", "name", nn.Name)
			return nil
		}
		return fmt.Errorf("error getting ConfigConnector object %v: %v", nn.Name, err)
	}
	r.recordEvent(cc, corev1.EventTypeNormal, k8s.UpToDate, k8s.UpToDateMessage)
	cc.SetCommonStatus(v1alpha1.CommonStatus{
		Healthy: true,
		Errors:  []string{},
	})
	return r.updateConfigConnectorStatus(ctx, cc)
}

// Handle the lifecycle of the given components under different conditions:
// 1) If the ConfigConnector object is pending deletion, ensure all deployed k8s components by CC controller don't exist or are deleted
// 2) If the ConfigConnector object is active, and if it’s cluster mode, verify that all per-namespace controller manager workloads are deleted,
//    then ensure KCC system components (shared components like CRDs and webhooks as well as cluster-mode-only components) are created.
// 3) If the ConfigConnector object is active, and if it’s namespaced mode, first remove cluster mode only components if any,
//    then ensure shared KCC system components (excluding cluster-mode-only components) are created.
func (r *ConfigConnectorReconciler) handleConfigConnectorLifecycle() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		r.log.Info("handling the lifecycle of the ConfigConnector object", "name", o.GetName())
		cc, ok := o.(*corev1beta1.ConfigConnector)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnector, but it was not. Object: %v", o)
		}
		// On delete
		if !cc.GetDeletionTimestamp().IsZero() {
			r.log.Info("finalizing the deletion of Config Connector system components deployed by ConfigConnector controller", "name", cc.Name)
			if err := r.finalizeSystemComponentsDeletion(ctx, r.client); err != nil {
				return errors.Wrap(err, "error finalizing the deletion of Config Connector system components deployed by ConfigConnector controller")
			}
			if controllers.RemoveOperatorFinalizer(cc) {
				if err := r.client.Update(ctx, cc); err != nil {
					return fmt.Errorf("error removing %v finalizer from ConfigConnector object %v: %v", k8s.OperatorFinalizer, cc.GetName(), err)
				}
			}
			// Nothing needs to apply when it's a delete ops.
			m.Items = nil
			return nil
		}
		// On apply
		if !controllers.EnsureOperatorFinalizer(cc) {
			if err := r.client.Update(ctx, cc); err != nil {
				return fmt.Errorf("error adding %v finalizer in ConfigConnector object %v: %v", k8s.OperatorFinalizer, cc.GetName(), err)
			}
			// Create the cnrm-systm namespace first; this is done to prevent the creation of components from failing due to the cnrm-system namespace not existing yet.
			if err := createCNRMSystemNamespace(ctx, r.client, m); err != nil {
				return fmt.Errorf("error creating %v namespace: %w", k8s.CNRMSystemNamespace, err)
			}
		}

		if cc.GetMode() == k8s.ClusterMode {
			// Verify that all per-namespace controller manager pods are removed, then continue the reconciliation.
			// This is done to avoid having more than one controller reconciling the same object.
			if err := r.verifyPerNamespaceControllerManagerPodsAreDeleted(ctx, r.client); err != nil {
				return fmt.Errorf("error waiting for all per-namespace controller manager pods to be removed: %w", err)
			}
		} else {
			if err := r.removeClusterModeOnlyComponents(ctx, r.client); err != nil {
				return err
			}
		}
		return nil
	}
}

func (r *ConfigConnectorReconciler) transformForClusterMode() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		cc, ok := o.(*corev1beta1.ConfigConnector)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnector, but it was not. Object: %v", o)
		}
		if cc.GetMode() == k8s.ClusterMode {
			if cc.Spec.GoogleServiceAccount != "" {
				// workload identity version
				if err := r.objectTransformForWorkloadIdentity(cc, m); err != nil {
					return errors.Wrap(err, "error transforming loadedManifest for workload identity version")
				}
			} else {
				// gcp identity version
				if err := r.objectTransformForGCPIdentity(cc, m); err != nil {
					return errors.Wrap(err, "error transforming loadedManifest for gcp identity version")
				}
			}
		}
		return nil
	}
}

func (r *ConfigConnectorReconciler) objectTransformForWorkloadIdentity(cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	transformed := make([]*manifest.Object, 0, len(m.Items))
	for _, obj := range m.Items {
		if obj.Kind == rbacv1.ServiceAccountKind && obj.GetName() == k8s.KCCControllerManagerComponent {
			r.log.Info("annotating controller manager service account with workload identity annotation")
			processed, err := controllers.AnnotateServiceAccountObject(obj, cc.Spec.GoogleServiceAccount)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error annotating ServiceAccount %v/%v", obj.UnstructuredObject().GetNamespace(), obj.UnstructuredObject().GetName()))
			}
			transformed = append(transformed, processed)
		} else {
			transformed = append(transformed, obj)
		}
	}
	m.Items = transformed
	return nil
}

func (r *ConfigConnectorReconciler) objectTransformForGCPIdentity(cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	transformed := make([]*manifest.Object, 0, len(m.Items))
	for _, obj := range m.Items {
		if controllers.IsControllerManagerStatefulSet(obj) {
			processed, err := setSecretVolume(obj, cc.Spec.CredentialSecretName)
			if err != nil {
				return err
			}
			transformed = append(transformed, processed)
		} else {
			transformed = append(transformed, obj)
		}
	}
	m.Items = transformed
	return nil
}

func createCNRMSystemNamespace(ctx context.Context, c client.Client, m *manifest.Objects) error {
	for _, obj := range m.Items {
		if obj.Kind == "Namespace" && obj.GetName() == k8s.CNRMSystemNamespace {
			if err := c.Create(ctx, obj.UnstructuredObject()); err != nil && !apierrors.IsAlreadyExists(err) {
				return err
			}
			break
		}
	}
	return nil
}

func (r *ConfigConnectorReconciler) removeClusterModeOnlyComponents(ctx context.Context, c client.Client) error {
	r.log.Info("removing controller manager components for cluster mode")
	svc := &corev1.Service{}
	svc.Namespace = k8s.CNRMSystemNamespace
	svc.Name = k8s.ControllerManagerService
	if err := controllers.DeleteObject(ctx, c, svc); err != nil {
		return err
	}

	sts := &appsv1.StatefulSet{}
	sts.Namespace = k8s.CNRMSystemNamespace
	sts.Name = k8s.KCCControllerManagerComponent
	if err := controllers.DeleteObject(ctx, c, sts); err != nil {
		return err
	}

	sc := &corev1.ServiceAccount{}
	sc.Namespace = k8s.CNRMSystemNamespace
	sc.Name = k8s.KCCControllerManagerComponent
	if err := controllers.DeleteObject(ctx, c, sc); err != nil {
		return err
	}
	return nil
}

func setSecretVolume(object *manifest.Object, secretName string) (*manifest.Object, error) {
	u := object.UnstructuredObject()
	volumes, ok, err := unstructured.NestedSlice(u.Object, "spec", "template", "spec", "volumes")
	if err != nil || !ok || len(volumes) == 0 {
		return nil, fmt.Errorf("couldn't find volumes from StatefulSet %v: %v", u.GetName(), err)
	}
	for _, volume := range volumes {
		volume := volume.(map[string]interface{})
		if volume["name"] == "gcp-service-account" {
			if err := unstructured.SetNestedField(volume, secretName, "secret", "secretName"); err != nil {
				return nil, fmt.Errorf("error setting the secret volume for StatefulSet %v: %v", u.GetName(), err)
			}
			if err := unstructured.SetNestedSlice(u.Object, volumes, "spec", "template", "spec", "volumes"); err != nil {
				return nil, fmt.Errorf("error setting the secret volume for StatefulSet %v: %v", u.GetName(), err)
			}
			return manifest.NewObject(u)
		}
	}
	return nil, fmt.Errorf("couldn't find the gcp-service-account volume to set for StatefulSet %v", u.GetName())
}

func (r *ConfigConnectorReconciler) verifyPerNamespaceControllerManagerPodsAreDeleted(ctx context.Context, c client.Client) error {
	podLabelSelector, err := labels.Parse(k8s.KCCControllerPodLabelSelectorRaw)
	if err != nil {
		return fmt.Errorf("error parsing '%v' as a label selector: %v", k8s.KCCControllerPodLabelSelectorRaw, err)
	}
	podList := &corev1.PodList{}
	podOpts := &client.ListOptions{
		Namespace:     k8s.CNRMSystemNamespace,
		LabelSelector: podLabelSelector,
		Limit:         100,
	}
	if err := c.List(ctx, podList, podOpts); err != nil {
		return fmt.Errorf("error listing controller manager pods: %w", err)
	}
	podNames := make([]string, 0, len(podList.Items))
	for _, p := range podList.Items {
		podNames = append(podNames, p.Name)
	}
	r.log.Info("verifying that per-namespace controller manager pods are deleted", "namespace", k8s.CNRMSystemNamespace, "pods", podNames)
	if len(podList.Items) == 0 {
		return nil
	}
	// Return nil if the only pod in the list is the controller manager pod for cluster mode.
	if len(podList.Items) == 1 && podList.Items[0].Name == k8s.ControllerManagerPodForClusterMode {
		return nil
	}
	return fmt.Errorf("per-namespace controller manager pods are not yet deleted by configconnectorcontext controller, reenquee the reconcilation for another attempt later; "+
		"remaining pods include, but may not be limited to %v", podNames)
}

func (r *ConfigConnectorReconciler) finalizeSystemComponentsDeletion(ctx context.Context, c client.Client) error {
	// Delete the global controller manager workload (deployed by the ConfigConnector controller when in cluster mode) if any.
	sts := &appsv1.StatefulSet{}
	sts.Namespace = k8s.CNRMSystemNamespace
	sts.Name = k8s.KCCControllerManagerComponent
	if err := controllers.DeleteObject(ctx, c, sts); err != nil {
		return err
	}

	// Wait until all the controller manager pods are removed first; this is an additional safeguard to ensure that we will NOT delete any resource on uninstallation.
	b := wait.Backoff{
		Duration: 1 * time.Second,
		Factor:   1.2,
		Steps:    12,
	}

	podLabelSelector, err := labels.Parse(k8s.KCCControllerPodLabelSelectorRaw)
	if err != nil {
		return fmt.Errorf("error parsing '%v' as a label selector: %v", k8s.KCCControllerPodLabelSelectorRaw, err)
	}
	podList := &corev1.PodList{}
	podOpts := &client.ListOptions{
		Namespace:     k8s.CNRMSystemNamespace,
		LabelSelector: podLabelSelector,
	}
	if err := wait.ExponentialBackoff(b, func() (done bool, err error) {
		if err := c.List(ctx, podList, podOpts); err != nil {
			return false, errors.Wrap(err, "error listing controller manager pods")
		}
		podNames := make([]string, 0, len(podList.Items))
		for _, p := range podList.Items {
			podNames = append(podNames, p.Name)
		}
		r.log.Info("waiting for controller manager pods to be deleted", "pods", podNames)
		return len(podList.Items) == 0, nil
	}); err != nil {
		return errors.Wrap(err, "error waiting for controller manager pods to be deleted")
	}
	r.log.Info("Successfully deleted all controller manager StatefulSets")

	// KCC adds finalizers to all its CRs. We need to delete each CRD that KCC
	// owns, and then wait for its controllers to finalize the deletion. Once the
	// CRDs are all cleaned up, we can delete the rest of KCC.
	b = wait.Backoff{
		Duration: 1 * time.Second,
		Factor:   1.2,
		Steps:    12,
	}
	if err := wait.ExponentialBackoff(b, func() (done bool, err error) {
		allDeleted := true
		crdList, _, err := k8s.ListCRDs(ctx, c, "")
		if err != nil {
			return false, err
		}
		for _, crd := range crdList {
			allDeleted = false
			r.log.Info("deleting CRD", "name", crd.GetName())
			if err := c.Delete(ctx, &crd); err != nil && !apierrors.IsNotFound(err) {
				return false, fmt.Errorf("error deleting CRD %v: %v", crd.GetName(), err)
			}
		}
		return allDeleted, nil
	}); err != nil {
		return errors.Wrap(err, "error waiting for CRDs to be deleted")
	}
	r.log.Info("Successfully finalized resource CRDs deletion.")

	// Specifically delete other resources set up by KCC on the fly.
	resources := []client.Object{
		&admissionregistration.ValidatingWebhookConfiguration{
			ObjectMeta: metav1.ObjectMeta{
				Name: "validating-webhook.cnrm.cloud.google.com",
			},
		},
		&admissionregistration.ValidatingWebhookConfiguration{
			ObjectMeta: metav1.ObjectMeta{
				Name: "abandon-on-uninstall.cnrm.cloud.google.com",
			},
		},
		&admissionregistration.MutatingWebhookConfiguration{
			ObjectMeta: metav1.ObjectMeta{
				Name: "mutating-webhook.cnrm.cloud.google.com",
			},
		},
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "cnrm-validating-webhook",
				Namespace: k8s.CNRMSystemNamespace,
			},
		},
		&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "abandon-on-uninstall",
				Namespace: k8s.CNRMSystemNamespace,
			},
		},
	}

	for _, obj := range resources {
		if err := controllers.DeleteObject(ctx, r.client, obj); err != nil {
			return err
		}
	}
	return nil
}

func (r *ConfigConnectorReconciler) updateConfigConnectorStatus(ctx context.Context, cc *corev1beta1.ConfigConnector) error {
	if err := r.client.Status().Update(ctx, cc); err != nil {
		if apierrors.IsConflict(err); err != nil {
			return fmt.Errorf("couldn't update ConfigConnector on API server due to conflict")
		}
		return fmt.Errorf("failed to update ConfigConnector on API server: %v", err)
	}
	return nil
}

func (r *ConfigConnectorReconciler) recordEvent(cc *corev1beta1.ConfigConnector, eventtype, reason, message string) {
	r.recorder.Event(cc, eventtype, reason, message)
}
