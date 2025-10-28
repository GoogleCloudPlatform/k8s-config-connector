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

package configconnectorcontext

import (
	"context"
	"fmt"
	"strings"

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	cnrmmanifest "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/manifest"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/preflight"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cluster"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

const controllerName = "configconnectorcontext-controller"

// ReconcilerOptions holds configuration options for the reconciler
type ReconcilerOptions struct {
	RepoPath                  string
	ImageTransform            *controllers.ImageTransform
	ManagerNamespaceIsolation string
}

// Reconciler reconciles a ConfigConnectorContext object.
//
// From the high level, the Reconciler watches `ConfigConnectorContext` kind
// and is responsible for managing the lifecycle of per-namespace KCC components (e.g. Service, StatefulSet, ServiceAccount and RoleBindings)
// independently with multiple workers.
// Reconciler also watches "NamespacedControllerResource" kind and apply
// customizations specified in "NamespacedControllerResource" CRs to per-namespace KCC components.
type Reconciler struct {
	reconciler                *declarative.Reconciler
	client                    client.Client
	recorder                  record.EventRecorder
	labelMaker                declarative.LabelMaker
	log                       logr.Logger
	customizationWatcher      *controllers.CustomizationWatcher
	jitterGen                 jitter.Generator
	managerNamespaceIsolation string
}

func Add(mgr ctrl.Manager, opt *ReconcilerOptions) error {
	r, err := newReconciler(mgr, opt)
	if err != nil {
		return err
	}

	// Create a new ConfigConnectorContext controller.
	obj := &corev1beta1.ConfigConnectorContext{}
	_, err = builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: 20}).
		WatchesRawSource(source.TypedChannel(r.customizationWatcher.Events(), &handler.EnqueueRequestForObject{})).
		For(obj, builder.OnlyMetadata).
		Build(r)
	if err != nil {
		return err
	}

	return nil
}

func newReconciler(mgr ctrl.Manager, opt *ReconcilerOptions) (*Reconciler, error) {
	repo := cnrmmanifest.NewLocalRepository(opt.RepoPath)
	manifestLoader := cnrmmanifest.NewPerNamespaceManifestLoader(repo)
	preflight := preflight.NewCompositePreflight([]declarative.Preflight{
		preflight.NewNameChecker(mgr.GetClient(), corev1beta1.ConfigConnectorContextAllowedName),
		preflight.NewUpgradeChecker(mgr.GetClient(), repo),
		preflight.NewConfigConnectorContextChecker(),
	})

	r := &Reconciler{
		reconciler:                &declarative.Reconciler{},
		client:                    mgr.GetClient(),
		recorder:                  mgr.GetEventRecorderFor(controllerName),
		labelMaker:                SourceLabel(),
		log:                       ctrl.Log.WithName(controllerName),
		jitterGen:                 &jitter.SimpleJitterGenerator{},
		managerNamespaceIsolation: opt.ManagerNamespaceIsolation,
	}

	r.customizationWatcher = controllers.NewWithDynamicClient(
		dynamic.NewForConfigOrDie(mgr.GetConfig()),
		controllers.CustomizationWatcherOptions{
			TriggerGVRs: controllers.NamespacedCustomizationCRsToWatch,
			Log:         r.log,
		})

	options := []declarative.ReconcilerOption{
		declarative.WithPreserveNamespace(),
		declarative.WithManifestController(manifestLoader),
		declarative.WithObjectTransform(r.transformNamespacedComponents()),
	}

	if opt.ManagerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
		options = append(options, declarative.WithObjectTransform(r.transformPerNamespaceComponents()))
	}

	options = append(options,
		declarative.WithObjectTransform(r.addLabels()),
		declarative.WithObjectTransform(r.handleCCContextLifecycle()),
		declarative.WithObjectTransform(r.applyNamespacedCustomizations()),
		declarative.WithStatus(&declarative.StatusBuilder{
			PreflightImpl: preflight,
		}))

	if opt.ImageTransform != nil {
		options = append(options, declarative.WithObjectTransform(opt.ImageTransform.RemapImages))
	}

	err := r.reconciler.Init(mgr, &corev1beta1.ConfigConnectorContext{}, options...)
	return r, err
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	r.log.Info("reconciling ConfigConnectorContext", "name", req.Name, "namespace", req.Namespace)
	_, err := r.getConfigConnectorContext(ctx, req.NamespacedName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("ConfigConnectorContext not found in API server; skipping the reconciliation", "name", req.NamespacedName)
			return reconcile.Result{}, nil
		}
	}
	_, reconciliationErr := r.reconciler.Reconcile(ctx, req)
	if reconciliationErr != nil {
		if err := r.handleReconcileFailed(ctx, req.NamespacedName, reconciliationErr); err != nil {
			return reconcile.Result{}, fmt.Errorf("error handling reconciled failed: %w, original reconciliation error: %w", err, reconciliationErr)
		}
		return reconcile.Result{}, reconciliationErr
	}
	// Setup watch for customization CRDs if not already done so in the previous reconciliations.
	// When there is a change detected on a customization CR, raises an event on ConfigConnectorContext CR.
	if err := r.customizationWatcher.EnsureWatchStarted(ctx, req.NamespacedName); err != nil {
		r.log.Error(err, "ensure watch start for customization CRDs failed")
		// Don't fail entire reconciliation if we cannot start watch for customization CRDs.
		// return reconcile.Result{}, err
	}
	jitteredPeriod := r.jitterGen.WatchJitteredTimeout()
	r.log.Info("successfully finished reconcile", "ConfigConnectorContext", req.NamespacedName, "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, r.handleReconcileSucceeded(ctx, req.NamespacedName)
}

func (r *Reconciler) getConfigConnectorContext(ctx context.Context, nn types.NamespacedName) (*corev1beta1.ConfigConnectorContext, error) {
	ccc := &corev1beta1.ConfigConnectorContext{}
	if err := r.client.Get(ctx, nn, ccc); err != nil {
		return nil, err
	}
	return ccc, nil
}

func (r *Reconciler) handleReconcileFailed(ctx context.Context, nn types.NamespacedName, reconcileErr error) error {
	ccc, err := r.getConfigConnectorContext(ctx, nn)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("ConfigConnectorContext not found in API server; skipping the handling of failed reconciliation", "namespace", nn.Namespace, "name", nn.Name)
			return nil
		}
		r.log.Info("error getting ConfigConnectorContext object", "namespace", nn.Namespace, "name", nn.Name, "reconcile error", reconcileErr)
		return fmt.Errorf("error getting ConfigConnectorContext object %v/%v: %w", nn.Namespace, nn.Name, err)
	}

	msg := fmt.Errorf("error during reconciliation: %w", reconcileErr).Error()
	r.recorder.Event(ccc, corev1.EventTypeWarning, k8s.UpdateFailed, msg)
	r.log.Info("surfacing error messages in status...", "namespace", nn.Namespace, "name", nn.Name, "error", msg)
	status := ccc.GetCommonStatus()
	status.Healthy = false
	status.Errors = []string{msg}
	status.ObservedGeneration = ccc.Generation
	ccc.SetCommonStatus(status)
	return r.updateConfigConnectorContextStatus(ctx, ccc)
}

func (r *Reconciler) handleReconcileSucceeded(ctx context.Context, nn types.NamespacedName) error {
	ccc, err := r.getConfigConnectorContext(ctx, nn)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("ConfigConnectorContext not found in API server; skipping the handling of successful reconciliation", "namespace", nn.Namespace, "name", nn.Name)
			return nil
		}
		return fmt.Errorf("error getting ConfigConnectorContext object %v/%v: %w", nn.Namespace, nn.Name, err)
	}

	r.recorder.Event(ccc, corev1.EventTypeNormal, k8s.UpToDate, k8s.UpToDateMessage)
	status := ccc.GetCommonStatus()
	status.Healthy = true
	status.Errors = []string{}
	status.ObservedGeneration = ccc.Generation
	ccc.SetCommonStatus(status)
	return r.updateConfigConnectorContextStatus(ctx, ccc)
}

func (r *Reconciler) updateConfigConnectorContextStatus(ctx context.Context, ccc *corev1beta1.ConfigConnectorContext) error {
	if err := r.client.Status().Update(ctx, ccc); err != nil {
		return fmt.Errorf("failed to update ConfigConnectorContext %v/%v on API server: %w", ccc.Namespace, ccc.Name, err)
	}
	return nil
}

func (r *Reconciler) transformNamespacedComponents() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was not. Object: %v", o)
		}
		transformedObjects, err := transformNamespacedComponentTemplates(ctx, r.client, ccc, m.Items)
		if err != nil {
			return fmt.Errorf("error transforming namespaced components: %w", err)
		}
		m.Items = transformedObjects
		return nil
	}
}

func (r *Reconciler) transformPerNamespaceComponents() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was not. Object: %v", o)
		}
		transformedObjects, err := transformPerNamespaceComponentTemplates(ctx, r.client, ccc, m.Items)
		if err != nil {
			return fmt.Errorf("error transforming per namespace components: %w", err)
		}
		m.Items = transformedObjects
		return nil
	}
}

// Add labels that will be used for the controller to dynamically watch on deployed KCC components.
func (r *Reconciler) addLabels() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, manifest *manifest.Objects) error {
		labels := r.labelMaker(ctx, o)
		for _, o := range manifest.Items {
			o.AddLabels(labels)
		}
		return nil
	}
}

// Handle the lifecycle of the given per-namespace components under different conditions:
// 1) If the ConfigConnector object is not found or pending deletion, this is the uninstallation case, finalize the deletion of per-namespace components.
// 2) If the ConfigConnector object says the cluster mode, finalize the deletion of per-namespace components,
//
//	returns some error like “ConfigConnector runs in cluster mode, this CCC is ignored and should be deleted”.
//
// 3) If the ConfigConnector object says the namespaced mode, and if this ConfigConnectorContext object is active, verify that the controller manager workload for the cluster mode is deleted and the ‘cnrm-system’ namespace is created,
//
//	then ensure per-namespace components are created.
//
// 4) If the ConfigConnector object says the namespaced mode, and if this ConfigConnectorContext object is pending deletion, verify that all KCC resource CRs are deleted, then finalize the deletion of per-namespace components.
func (r *Reconciler) handleCCContextLifecycle() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
		if !ok {
			return fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was not. Object: %v", o)
		}
		var isCCObjectNotFound bool
		cc, err := controllers.GetConfigConnector(ctx, r.client, controllers.ValidConfigConnectorNamespacedName)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				return fmt.Errorf("error getting the ConfigConnector object %v: %w", controllers.ValidConfigConnectorNamespacedName, err)
			}
			isCCObjectNotFound = true
		}
		if r.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated && ccc.Spec.ManagerNamespace == "" {
			return fmt.Errorf("error in ConfigConnectorContext object %v: dedicated manager namespace mode enabled, but spec.ManagerNamespace is not configured", controllers.ValidConfigConnectorNamespacedName)
		}
		if r.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationShared && ccc.Spec.ManagerNamespace != "" {
			return fmt.Errorf("error in ConfigConnectorContext object %v: dedicated manager namespace mode disabled, but spec.ManagerNamespace=%s is configured", controllers.ValidConfigConnectorNamespacedName, ccc.Spec.ManagerNamespace)
		}
		if isCCObjectNotFound || !cc.GetDeletionTimestamp().IsZero() {
			return r.finalizeSystemComponentsDeletion(ctx, ccc, m)
		}
		if cc.GetMode() == k8s.ClusterMode {
			return r.handleCCContextLifecycleForClusterMode(ctx, ccc, m)
		}
		return r.handleCCContextLifecycleForNamespacedMode(ctx, ccc, m)
	}
}

func (r *Reconciler) finalizeCCContextDeletion(ctx context.Context, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	r.log.Info("ConfigConnectorContext for namespace is marked for deletion; verifying all CNRM resources have been deleted...", "namespace", ccc.Namespace)
	kindToCount, err := getCNRMResourceCounts(ctx, r.client, ccc.Namespace)
	if err != nil {
		return fmt.Errorf("error verifying the Config Connector resource counts in namespace '%v': %w", ccc.Namespace, err)
	}
	if len(kindToCount) > 0 {
		r.log.Info("Cannot finalize deletion of ConfigConnectorContext: there are still Config Connector resource(s) in the namespace.",
			"namespace", ccc.Namespace, "numKindsWithResources", len(kindToCount))
		return formatCNRMResourcesPresentError(kindToCount)
	}
	if err := r.finalizeNamespacedComponentsDeletion(ctx, ccc, m); err != nil {
		return err
	}
	if err := cluster.DeleteNamespaceID(ctx, k8s.OperatorNamespaceIDConfigMapNN, r.client, ccc.Namespace); err != nil {
		return err
	}
	if r.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
		if err := r.deleteManagerNamepace(ctx, ccc.Spec.ManagerNamespace); err != nil {
			return err
		}
	}
	r.log.Info("Successfully finalized ConfigConnectorContext deletion...", "name", ccc.Name, "namespace", ccc.Namespace)
	// Nothing needs to apply when it's a delete ops.
	m.Items = nil
	return nil
}

func (r *Reconciler) finalizeNamespacedComponentsDeletion(ctx context.Context, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	r.log.Info("finalizing namespaced components deletion...", "namespace", ccc.Namespace)
	if err := removeNamespacedComponents(ctx, r.client, m.Items); err != nil {
		return fmt.Errorf("error finalizing ConfigConnectorContext %v/%v deletion: %w", ccc.Namespace, ccc.Name, err)
	}
	if controllers.RemoveOperatorFinalizer(ccc) {
		if err := r.client.Update(ctx, ccc); err != nil {
			return fmt.Errorf("error removing %v finalizer in ConfigConnectorContext object %v/%v: %w", k8s.OperatorFinalizer, ccc.Namespace, ccc.GetName(), err)
		}
	}
	return nil
}

func (r *Reconciler) handleCCContextLifecycleForClusterMode(ctx context.Context, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	// On the cluster mode, clean up namespaced components associated with the ConfigConnectorContext.
	if err := r.finalizeNamespacedComponentsDeletion(ctx, ccc, m); err != nil {
		return err
	}
	return fmt.Errorf("ConfigConnector is in cluster-mode, this ConfigConnectorContext object does not serve any purpose and should be removed")
}

func (r *Reconciler) handleCCContextLifecycleForNamespacedMode(ctx context.Context, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	if !ccc.GetDeletionTimestamp().IsZero() {
		return r.finalizeCCContextDeletion(ctx, ccc, m)
	}
	// Verify that the controller manager pod for cluster mode is removed, then continue the reconciliation.
	// This is done to avoid having more than one controller reconciling the same object.
	if err := r.verifyControllerManagerPodForClusterModeIsDeleted(ctx); err != nil {
		return err
	}
	if err := r.verifyCNRMSystemNamespaceIsActive(ctx); err != nil {
		return err
	}
	if r.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
		if err := r.verifyManagerNamespaceIsActive(ctx, ccc.Spec.ManagerNamespace); err != nil {
			return err
		}
	}

	if !controllers.EnsureOperatorFinalizer(ccc) {
		if err := r.client.Update(ctx, ccc); err != nil {
			return fmt.Errorf("error adding %v finalizer in ConfigConnectorContext object %v: %w", k8s.OperatorFinalizer, client.ObjectKeyFromObject(ccc), err)
		}
	}
	return nil
}

func (r *Reconciler) finalizeSystemComponentsDeletion(ctx context.Context, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	r.log.Info("deleting namespaced components on uninstallation", "namespace", ccc.Namespace)
	if err := r.finalizeNamespacedComponentsDeletion(ctx, ccc, m); err != nil {
		return err
	}
	m.Items = nil
	if ccc.GetDeletionTimestamp().IsZero() {
		return fmt.Errorf("the ConfigConnector object %v is not found or pending deletion; this ConfigConnectorContext object does not serve any purpose and should be removed", controllers.ValidConfigConnectorNamespacedName)
	}
	return nil
}

func (r *Reconciler) verifyControllerManagerPodForClusterModeIsDeleted(ctx context.Context) error {
	sts := &appsv1.StatefulSet{}
	sts.Namespace = k8s.CNRMSystemNamespace
	sts.Name = k8s.KCCControllerManagerComponent
	stsKey := client.ObjectKeyFromObject(sts)

	pod := &corev1.Pod{}
	pod.Namespace = k8s.CNRMSystemNamespace
	pod.Name = k8s.ControllerManagerPodForClusterMode
	podKey := client.ObjectKeyFromObject(pod)

	r.log.Info("verifying that cluster mode workload is deleted...", "StatefulSet", stsKey, "Pod", podKey)
	err := r.client.Get(ctx, stsKey, sts)
	if err == nil {
		return fmt.Errorf("statefulset %v is not yet deleted, reenquee the reconciliation for another attempt later", stsKey)
	}
	if !apierrors.IsNotFound(err) {
		return fmt.Errorf("error getting the StatefulSet %v: %w", stsKey, err)
	}

	err = r.client.Get(ctx, podKey, pod)
	if err == nil {
		return fmt.Errorf("pod %v is not yet deleted, reenquee the reconciliation for another attempt later", stsKey)
	}
	if !apierrors.IsNotFound(err) {
		return fmt.Errorf("error getting the pod %v: %w", podKey, err)
	}

	return nil
}

func (r *Reconciler) verifyCNRMSystemNamespaceIsActive(ctx context.Context) error {
	r.log.Info("verifying that ConfigConnector system namespace is active...", "system namespace", k8s.CNRMSystemNamespace)
	n := &corev1.Namespace{}
	key := types.NamespacedName{
		Name: k8s.CNRMSystemNamespace,
	}
	if err := r.client.Get(ctx, key, n); err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("ConfigConnector system namespace %v is not created by configconnector controller yet, reenquee the reconciliation for another attempt later", k8s.CNRMSystemNamespace)
		}

		return fmt.Errorf("error getting the ConfigConnector system namespace %v: %w", key, err)
	}
	if !n.GetDeletionTimestamp().IsZero() {
		return fmt.Errorf("ConfigConnector system namespace %v is pending deletion, stop the reconciliation", k8s.CNRMSystemNamespace)
	}
	return nil
}

func (r *Reconciler) verifyManagerNamespaceIsActive(ctx context.Context, managerNamespace string) error {
	r.log.Info("verifying that ConfigConnector manager namespace is active...", "manager namespace", managerNamespace)
	n := &corev1.Namespace{}
	key := types.NamespacedName{
		Name: managerNamespace,
	}
	if err := r.client.Get(ctx, key, n); err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("verifying that ConfigConnector manager namespace not found, creating...", "manager namespace", managerNamespace)
			ns := &corev1.Namespace{}
			ns.SetName(managerNamespace)
			ns.SetLabels(map[string]string{k8s.ManagedByKCCLabel: "true"})
			if err := r.client.Create(ctx, ns); err != nil {
				return fmt.Errorf("error creating the ConfigConnector manager namespace %v: %w", ns, err)
			}
			return nil
		}
		return fmt.Errorf("error getting the ConfigConnector manager namespace %v: %w", key, err)
	}
	if !n.GetDeletionTimestamp().IsZero() {
		return fmt.Errorf("ConfigConnector manager namespace %v is pending deletion, stop the reconciliation", managerNamespace)
	}
	return nil
}

func (r *Reconciler) deleteManagerNamepace(ctx context.Context, managerNamespace string) error {
	r.log.Info("deleting ConfigConnector manager namespace...", "manager namespace", managerNamespace)
	n := &corev1.Namespace{}
	key := types.NamespacedName{
		Name: managerNamespace,
	}
	if err := r.client.Get(ctx, key, n); err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("not found ConfigConnector manager namespace...", "manager namespace", managerNamespace)
			return nil
		}
		return err
	}
	if !n.GetDeletionTimestamp().IsZero() {
		r.log.Info("ConfigConnector manager namespace already deleted...", "manager namespace", managerNamespace)
		return nil
	}
	managed, found := n.GetLabels()[k8s.ManagedByKCCLabel]
	if found && managed == "true" {
		r.log.Info("managed ConfigConnector manager namespace...", "manager namespace", managerNamespace)
		if err := r.client.Delete(ctx, n); err != nil {
			return fmt.Errorf("error deleting the ConfigConnector manager namespace %v: %w", n, err)
		}
	} else {
		r.log.Info("manager namespace is not managed by ConfigConnector, skip deletion...", "manager namespace", managerNamespace)
	}
	return nil
}

// applyNamespacedCustomizations fetches and applies all namespace-scoped customization CRDs.
func (r *Reconciler) applyNamespacedCustomizations() declarative.ObjectTransform {
	return func(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
		if err := r.fetchAndApplyAllNamespacedControllerResources(ctx, o, m); err != nil {
			r.log.Error(err, "error applying all namespaced controller resource customization CRs")
			// Don't fail entire reconciliation if we cannot apply customization CRs.
			// return err
		}
		if err := r.fetchAndApplyAllNamespacedControllerReconcilers(ctx, o, m); err != nil {
			r.log.Error(err, "error applying all namespaced controller reconciler customization CRs")
			// Don't fail entire reconciliation if we cannot apply customization CRs.
			// return err
		}
		return nil
	}
}

func (r *Reconciler) fetchAndApplyAllNamespacedControllerResources(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
	ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
	if !ok {
		return fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was not. Object: %v", o)
	}
	// List all the NamespacedControllerResource CRs in the same namespace as ConfigConnectorContext object.
	crs, err := controllers.ListNamespacedControllerResources(ctx, r.client, ccc.Namespace)
	if err != nil {
		return err
	}
	// Apply all the NamespacedControllerResource CRs in the same namespace as ConfigConnectorContext object.
	for _, cr := range crs {
		if cr.Namespace != ccc.Namespace {
			// this shouldn't happen!
			r.log.Error(fmt.Errorf("unexpected namespace for NamespacedControllerResource object"), "expected namespace", ccc.Namespace, "got namespace", cr.Namespace)
		}
		r.log.Info("applying namespace-scoped controller resource customization", "Namespace", cr.Namespace, "Name", cr.Name)
		if err := r.applyNamespacedControllerResource(ctx, &cr, m); err != nil {
			return err
		}
	}
	return nil
}

// applyNamespacedControllerResource applies customizations specified in NamespacedControllerResource CR.
func (r *Reconciler) applyNamespacedControllerResource(ctx context.Context, cr *customizev1beta1.NamespacedControllerResource, m *manifest.Objects) error {
	if cr.Name != "cnrm-controller-manager" {
		msg := fmt.Sprintf("resource customization for controller %s is not supported", cr.Name)
		r.log.Info(msg)
		return r.handleApplyNamespacedControllerResourceFailed(ctx, cr.Namespace, cr.Name, msg)
	}
	controllerGVK := schema.GroupVersionKind{
		Group:   appsv1.SchemeGroupVersion.Group,
		Version: appsv1.SchemeGroupVersion.Version,
		Kind:    "StatefulSet",
	}
	if err := controllers.ApplyContainerResourceCustomization(true, m, cr.Name, controllerGVK, cr.Spec.Containers, nil); err != nil {
		r.log.Error(err, "failed to apply customization", "Namespace", cr.Namespace, "Name", cr.Name)
		return r.handleApplyNamespacedControllerResourceFailed(ctx, cr.Namespace, cr.Name, fmt.Sprintf("failed to apply customization %s: %v", cr.Name, err))
	}
	return r.handleApplyNamespacedControllerResourceSucceeded(ctx, cr.Namespace, cr.Name)
}

func (r *Reconciler) handleApplyNamespacedControllerResourceFailed(ctx context.Context, namespace, name string, msg string) error {
	cr, err := controllers.GetNamespacedControllerResource(ctx, r.client, namespace, name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("NamespacedControllerResource object not found; skipping the handling of failed customization apply", "namespace", namespace, "name", name)
			return nil
		}
		// Don't fail entire reconciliation if we cannot get NamespacedControllerResource object.
		// return fmt.Errorf("error getting NamespacedControllerResource object %v/%v: %v", namespace, name, err)
		r.log.Error(err, "error getting NamespacedControllerResource object %v", "Namespace", namespace, "Name", name)
		return nil
	}
	status := cr.GetCommonStatus()
	status.Healthy = false
	status.Errors = []string{msg}
	status.ObservedGeneration = cr.Generation
	cr.SetCommonStatus(status)
	return r.updateNamespacedControllerResourceStatus(ctx, cr)
}

func (r *Reconciler) handleApplyNamespacedControllerResourceSucceeded(ctx context.Context, namespace, name string) error {
	cr, err := controllers.GetNamespacedControllerResource(ctx, r.client, namespace, name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("NamespacedControllerResource object not found; skipping the handling of succeeded customization apply", "namespace", namespace, "name", name)
			return nil
		}
		// Don't fail entire reconciliation if we cannot get NamespacedControllerResource object.
		// return fmt.Errorf("error getting NamespacedControllerResource object %v/%v: %v", namespace, name, err)
		r.log.Error(err, "error getting NamespacedControllerResource object %v", "Namespace", namespace, "Name", name)
		return nil
	}
	status := cr.GetCommonStatus()
	status.Healthy = true
	status.Errors = []string{}
	status.ObservedGeneration = cr.Generation
	cr.SetCommonStatus(status)
	return r.updateNamespacedControllerResourceStatus(ctx, cr)
}

func (r *Reconciler) updateNamespacedControllerResourceStatus(ctx context.Context, cr *customizev1beta1.NamespacedControllerResource) error {
	if err := r.client.Status().Update(ctx, cr); err != nil {
		r.log.Error(err, "failed to update NamespacedControllerResource", "Namespace", cr.Namespace, "Name", cr.Name, "Object", cr)
		// Don't fail entire reconciliation if we cannot update NamespacedControllerResource status.
		// return fmt.Errorf("failed to update NamespacedControllerResource %v/%v: %v", cr.Namespace, cr.Name, err)
	}
	return nil
}

func (r *Reconciler) fetchAndApplyAllNamespacedControllerReconcilers(ctx context.Context, o declarative.DeclarativeObject, m *manifest.Objects) error {
	ccc, ok := o.(*corev1beta1.ConfigConnectorContext)
	if !ok {
		return fmt.Errorf("expected the resource to be a ConfigConnectorContext, but it was not. Object: %v", o)
	}
	// List all the NamespacedControllerReconciler CRs in the same namespace as ConfigConnectorContext object.
	crs, err := controllers.ListNamespacedControllerReconcilers(ctx, r.client, ccc.Namespace)
	if err != nil {
		return err
	}
	// Apply all the NamespacedControllerReconciler CRs in the same namespace as ConfigConnectorContext object.
	for _, cr := range crs {
		if cr.Namespace != ccc.Namespace {
			// this shouldn't happen!
			r.log.Error(fmt.Errorf("unexpected namespace for NamespacedControllerReconciler object"), "expected namespace", ccc.Namespace, "got namespace", cr.Namespace)
		}
		r.log.Info("applying namespace-scoped controller reconciler customization", "Namespace", cr.Namespace, "Name", cr.Name)
		if err := r.applyNamespacedControllerReconciler(ctx, &cr, m); err != nil {
			return err
		}
	}
	return nil
}

// applyNamespacedControllerReconciler applies customizations specified in NamespacedControllerReconciler CR.
func (r *Reconciler) applyNamespacedControllerReconciler(ctx context.Context, cr *customizev1beta1.NamespacedControllerReconciler, m *manifest.Objects) error {
	if err := controllers.ApplyContainerRateLimit(m, cr.Name, cr.Spec.RateLimit); err != nil {
		msg := fmt.Sprintf("failed to apply rate limit customization %s: %v", cr.Name, err)
		return r.handleApplyNamespacedControllerReconcilerFailed(ctx, cr.Namespace, cr.Name, msg)
	}
	if err := controllers.ApplyContainerPprof(m, cr.Name, cr.Spec.Pprof); err != nil {
		msg := fmt.Sprintf("failed to apply pprof customization %s: %v", cr.Name, err)
		return r.handleApplyNamespacedControllerReconcilerFailed(ctx, cr.Namespace, cr.Name, msg)
	}
	return r.handleApplyNamespacedControllerReconcilerSucceeded(ctx, cr.Namespace, cr.Name)
}

func (r *Reconciler) handleApplyNamespacedControllerReconcilerFailed(ctx context.Context, namespace, name string, msg string) error {
	cr, err := controllers.GetNamespacedControllerReconciler(ctx, r.client, namespace, name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("NamespacedControllerReconciler object not found; skipping the handling of failed customization apply", "namespace", namespace, "name", name)
			return nil
		}
		r.log.Info("error getting NamespacedControllerReconciler object", "namespace", namespace, "name", name, "error", err, "reconciler error", msg)
		return err
	}
	cr.Status.CommonStatus = v1alpha1.CommonStatus{
		Healthy: false,
		Errors:  []string{msg},
	}
	return r.updateNamespacedControllerReconcilerStatus(ctx, cr)
}

func (r *Reconciler) handleApplyNamespacedControllerReconcilerSucceeded(ctx context.Context, namespace, name string) error {
	cr, err := controllers.GetNamespacedControllerReconciler(ctx, r.client, namespace, name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.log.Info("NamespacedControllerReconciler object not found; skipping the handling of failed customization apply", "namespace", namespace, "name", name)
			return nil
		}
		r.log.Info("error getting NamespacedControllerReconciler object", "namespace", namespace, "name", name)
		return err
	}
	cr.SetCommonStatus(v1alpha1.CommonStatus{
		Healthy: true,
		Errors:  []string{},
	})
	return r.updateNamespacedControllerReconcilerStatus(ctx, cr)
}

func (r *Reconciler) updateNamespacedControllerReconcilerStatus(ctx context.Context, cr *customizev1beta1.NamespacedControllerReconciler) error {
	if err := r.client.Status().Update(ctx, cr); err != nil {
		r.log.Error(err, "failed to update NamespacedControllerReconciler", "namespace", cr.Namespace, "name", cr.Name)
		return fmt.Errorf("failed to update NamespacedControllerReconciler %v/%v: %w", cr.Namespace, cr.Name, err)
	}
	return nil
}

func getCNRMResourceCounts(ctx context.Context, kubeClient client.Client, namespace string) (map[string]int64, error) {
	kindToCount := make(map[string]int64)
	pageToken := ""
	var list []apiextensions.CustomResourceDefinition
	var err error
	for ok := true; ok; ok = pageToken != "" {
		list, pageToken, err = k8s.ListCRDs(ctx, kubeClient, pageToken)
		if err != nil {
			return nil, err
		}
		for _, crd := range list {
			if len(crd.Spec.Versions) == 0 {
				continue
			}
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: crd.Spec.Versions[0].Name,
				Kind:    crd.Spec.Names.Kind,
			}
			count, err := countResourcesForGVK(ctx, kubeClient, namespace, gvk)
			if err != nil {
				return nil, err
			}
			if count == 0 {
				continue
			}
			kindToCount[gvk.Kind] = count
		}
	}
	return kindToCount, nil
}

func countResourcesForGVK(ctx context.Context, kubeClient client.Client, namespace string, gvk schema.GroupVersionKind) (int64, error) {
	listOpts := &client.ListOptions{
		Limit:     1000,
		Namespace: namespace,
		Raw:       &metav1.ListOptions{},
	}
	resourceCount := int64(0)
	for ok := true; ok; ok = listOpts.Continue != "" {
		list := unstructured.UnstructuredList{}
		list.SetGroupVersionKind(gvk)
		if err := kubeClient.List(ctx, &list, listOpts); err != nil {
			return 0, fmt.Errorf("error listing loadedManifest for gvk '%v': %w", gvk, err)
		}
		resourceCount += int64(len(list.Items))
		listOpts.Continue = list.GetContinue()
	}
	return resourceCount, nil
}

func formatCNRMResourcesPresentError(kindToCount map[string]int64) error {
	totalCount := int64(0)
	kindCountStrings := make([]string, 0, len(kindToCount))
	for kind, count := range kindToCount {
		totalCount += count
		kindCountStrings = append(kindCountStrings, fmt.Sprintf("%v %v(s)", count, kind))
	}
	msg := "cannot finalize deletion until all Config Connector resources in namespace have been removed: there are %v Config Connector resource(s) in namespace (%v)"
	return fmt.Errorf(msg, totalCount, strings.Join(kindCountStrings, ", "))
}
