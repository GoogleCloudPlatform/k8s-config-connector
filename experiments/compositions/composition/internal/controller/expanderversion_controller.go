/*
Copyright 2024.

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

package controller

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"

	semver "github.com/Masterminds/semver/v3"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	compositionv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/api/v1alpha1"
)

// ExpanderVersionReconciler reconciles a ExpanderVersion object
type ExpanderVersionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=composition.google.com,resources=expanderversions,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=composition.google.com,resources=expanderversions/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=composition.google.com,resources=expanderversions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ExpanderVersion object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *ExpanderVersionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Got a new request!", "request", req)

	var ev compositionv1alpha1.ExpanderVersion
	if err := r.Client.Get(ctx, req.NamespacedName, &ev); err != nil {
		logger.Error(err, "unable to fetch ExpanderVersion")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Grab status for comparison later
	oldStatus := ev.Status.DeepCopy()

	// Try updating status before returning
	defer func() {
		if !reflect.DeepEqual(oldStatus, ev.Status) {
			newStatus := ev.Status.DeepCopy()
			err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				nn := types.NamespacedName{Namespace: ev.Namespace, Name: ev.Name}
				err := r.Client.Get(ctx, nn, &ev)
				if err != nil {
					return err
				}
				ev.Status = *newStatus.DeepCopy()
				return r.Client.Status().Update(ctx, &ev)
			})
			if err != nil {
				logger.Error(err, "unable to update Composition status")
			}
		}
	}()

	logger = logger.WithName(ev.Name).WithName(fmt.Sprintf("%d", ev.Generation))

	logger.Info("Validating ExpanderVersion object")
	if !ev.Validate() {
		logger.Info("Validation Failed")
		return ctrl.Result{}, fmt.Errorf("Validation failed")
	}

	ev.Status.ClearCondition(compositionv1alpha1.Error)
	logger.Info("Processing ExpanderVersion object")
	r.processExpanderVersion(&ev, logger)
	return ctrl.Result{}, nil
}

func (r *ExpanderVersionReconciler) processExpanderVersion(
	ev *compositionv1alpha1.ExpanderVersion, logger logr.Logger,
) {
	if ev.Status.VersionMap == nil {
		ev.Status.VersionMap = make(map[string]string)
	}

	expander := strings.TrimPrefix(ev.Name, "composition-")
	image := ev.Spec.Image
	if image == "" {
		image = fmt.Sprintf("expander-%s", expander)
	}
	semVerVersions := []*semver.Version{}
	for _, r := range ev.Spec.ValidVersions {
		v, err := semver.NewVersion(r)
		if err != nil {
			logger.Info("Error parsing version: %s", err)
			continue
		}

		semVerVersions = append(semVerVersions, v)
		key := "v" + strings.TrimPrefix(r, "v")
		value := ""

		if ev.Spec.Type == compositionv1alpha1.ExpanderTypeJob {
			value = fmt.Sprintf("%s/%s:%s", ev.Spec.ImageRegistry, image, key)
		} else {
			svcVersion := strings.Replace(key, ".", "-", -1)
			value = fmt.Sprintf("composition-%s-%s:8443", expander, svcVersion)
		}
		ev.Status.VersionMap[key] = value
	}

	sort.Sort(semver.Collection(semVerVersions))
	latest := "v" + semVerVersions[len(semVerVersions)-1].String()
	ev.Status.VersionMap["latest"] = ev.Status.VersionMap[latest]
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExpanderVersionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&compositionv1alpha1.ExpanderVersion{}).
		Complete(r)
}
