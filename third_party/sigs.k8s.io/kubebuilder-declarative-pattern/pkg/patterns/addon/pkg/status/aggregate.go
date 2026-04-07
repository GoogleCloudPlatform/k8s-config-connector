/*
Copyright 2019 The Kubernetes Authors.

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

package status

import (
	"context"
	"fmt"
	"reflect"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/utils"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"
)

const successfulDeployment = appsv1.DeploymentAvailable

// NewAggregator provides an implementation of declarative.Reconciled that
// aggregates the status of deployed objects to configure the 'Healthy'
// field on an addon that derives from CommonStatus
//
// TODO: Create a version that doesn't require the unused client arg
func NewAggregator(_ client.Client) *aggregator {
	return &aggregator{}
}

type aggregator struct {
}

func (a *aggregator) BuildStatus(ctx context.Context, info *declarative.StatusInfo) error {
	log := log.FromContext(ctx)

	statusHealthy := true
	statusErrors := []string{}

	shouldComputeHealthFromObjects := info.Manifest != nil && info.LiveObjects != nil
	if info.Err != nil {
		statusHealthy = false
		shouldComputeHealthFromObjects = false
	}

	if shouldComputeHealthFromObjects {
		for _, o := range info.Manifest.GetItems() {
			gvk := o.GroupVersionKind()
			nn := o.NamespacedName()

			log := log.WithValues("kind", gvk.Kind).WithValues("name", nn.Name).WithValues("namespace", nn.Namespace)

			healthy := true

			var err error
			switch gvk.Group + "/" + gvk.Kind {
			case "/Service":
				healthy, err = a.serviceIsHealthy(ctx, info.LiveObjects, gvk, nn)
			case "extensions/Deployment", "apps/Deployment":
				healthy, err = a.deploymentIsHealthy(ctx, info.LiveObjects, gvk, nn)
			default:
				log.V(4).Info("type not implemented for status aggregation, skipping")
			}

			statusHealthy = statusHealthy && healthy
			if err != nil {
				statusErrors = append(statusErrors, fmt.Sprintf("%v", err))
			}
		}
	}

	log.WithValues("status", statusHealthy).V(2).Info("built status")

	currentStatus, err := utils.GetCommonStatus(info.Subject)
	if err != nil {
		return err
	}

	status := currentStatus
	status.Healthy = statusHealthy
	status.Errors = statusErrors
	status.ObservedGeneration = info.Subject.GetGeneration()

	if !reflect.DeepEqual(status, currentStatus) {
		err := utils.SetCommonStatus(info.Subject, status)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *aggregator) deploymentIsHealthy(ctx context.Context, liveObjects declarative.LiveObjectReader, gvk schema.GroupVersionKind, nn types.NamespacedName) (bool, error) {
	u, err := liveObjects(ctx, gvk, nn)
	if err != nil {
		return false, fmt.Errorf("error reading deployment: %w", err)
	}

	dep := &appsv1.Deployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, dep); err != nil {
		return false, fmt.Errorf("error converting deployment from unstructured: %w", err)
	}

	for _, cond := range dep.Status.Conditions {
		if cond.Type == successfulDeployment && cond.Status == corev1.ConditionTrue {
			return true, nil
		}
	}

	return false, fmt.Errorf("deployment does not meet condition: %s", successfulDeployment)
}

func (a *aggregator) serviceIsHealthy(ctx context.Context, liveObjects declarative.LiveObjectReader, gvk schema.GroupVersionKind, nn types.NamespacedName) (bool, error) {
	_, err := liveObjects(ctx, gvk, nn)
	if err != nil {
		return false, fmt.Errorf("error reading service: %w", err)
	}

	return true, nil
}
