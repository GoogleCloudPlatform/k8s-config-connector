/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package simpletest

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/kubebuilder-declarative-pattern/commonclient"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"

	api "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/testreconciler/simpletest/v1alpha1"
)

var _ reconcile.Reconciler = &SimpleTestReconciler{}

// SimpleTestReconciler reconciles a SimpleTest object
type SimpleTestReconciler struct {
	declarative.Reconciler
	client.Client
	Log        logr.Logger
	Scheme     *runtime.Scheme
	TestSuffix string

	watchLabels declarative.LabelMaker

	manifestController declarative.ManifestController

	applier applier.Applier
	status  declarative.Status
}

func (r *SimpleTestReconciler) setupReconciler(mgr ctrl.Manager) error {
	labels := map[string]string{
		"example-app": "simpletest",
	}

	r.watchLabels = declarative.SourceLabel(mgr.GetScheme())

	return r.Reconciler.Init(mgr, &api.SimpleTest{},
		declarative.WithObjectTransform(declarative.AddLabels(labels)),
		declarative.WithOwner(declarative.SourceAsOwner),
		declarative.WithLabels(r.watchLabels),
		declarative.WithStatus(r.status),

		// TODO: Readd prune
		//declarative.WithApplyPrune(),

		declarative.WithObjectTransform(addon.ApplyPatches),

		// Add other options for testing
		//declarative.WithApplyValidation(),

		// Don't turn on metrics, they create another watch and cause unpredictable requests
		// declarative.WithReconcileMetrics(0, nil),

		declarative.WithManifestController(r.manifestController),
		declarative.WithApplier(r.applier),

		declarative.WithHook(&sleepAfterUpdateStatusHook{}),
	)
}

func (r *SimpleTestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := r.setupReconciler(mgr); err != nil {
		return err
	}

	c, err := controller.New(fmt.Sprintf("simpletest-controller-%s", r.TestSuffix), mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to SimpleTest objects
	err = c.Watch(commonclient.SourceKind(mgr.GetCache(), &api.SimpleTest{}))
	if err != nil {
		return err
	}

	// Watch for changes to deployed objects
	err = declarative.WatchChildren(declarative.WatchChildrenOptions{Manager: mgr, Controller: c, Reconciler: r, LabelMaker: r.watchLabels})
	if err != nil {
		return err
	}

	return nil
}

// We sleep briefly after updating the status.
// This is a hack to ensure that we see the watch event, because otherwise we can start a re-reconcile based on our derived objects,
// and then reconcile again based on our own status update.  This is racy behaviour and causes non-determinism.
// We do this only in this test controller, but maybe we should have similar logic to "debounce" in all controllers.
type sleepAfterUpdateStatusHook struct {
}

var _ declarative.AfterUpdateStatus = &sleepAfterUpdateStatusHook{}

func (h *sleepAfterUpdateStatusHook) AfterUpdateStatus(ctx context.Context, op *declarative.UpdateStatusOperation) error {
	sleepFor := 100 * time.Millisecond
	klog.Infof("sleeping after apply for %v", sleepFor)
	time.Sleep(sleepFor)
	return nil
}
