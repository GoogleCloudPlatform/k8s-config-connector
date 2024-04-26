// Copyright 2024 Google LLC
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

package utils

import (
	"fmt"
	"time"

	compositionv1alpha1 "google.com/composition/api/v1alpha1"
	"google.com/composition/internal/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(compositionv1alpha1.AddToScheme(scheme))
}

func StartLocalController(config *rest.Config, imageRegistry string) error {
	mgr, err := ctrl.NewManager(config, ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
	})

	if err != nil {
		return fmt.Errorf("Unable to start manager: %w", err)
	}

	if err = (&controller.CompositionReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		ImageRegistry: imageRegistry,
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create Composition controller: %w", err)
	}
	if err = (&controller.ContextReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create Context controller: %w", err)
	}
	if err = (&controller.PlanReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create Plan controller: %w", err)
	}

	//if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
	//	return fmt.Errorf("Problem running manager: %w", err)
	//}
	go mgr.Start(ctrl.SetupSignalHandler())
	time.Sleep(time.Second * 5)
	return nil
}

func GetUnstructuredObj(g, v, k, ns, n string) *unstructured.Unstructured {
	cr := unstructured.Unstructured{}
	cr.SetGroupVersionKind(schema.GroupVersionKind{Group: g, Version: v, Kind: k})
	cr.SetName(n)
	cr.SetNamespace(ns)
	return &cr
}

func GetCompositionObj(ns, n string) *unstructured.Unstructured {
	return GetUnstructuredObj("composition.google.com", "v1alpha1", "Composition", ns, n)
}

func GetPlanObj(ns, n string) *unstructured.Unstructured {
	return GetUnstructuredObj("composition.google.com", "v1alpha1", "Plan", ns, n)
}

func GetConfigMapObj(ns, n string) *unstructured.Unstructured {
	return GetUnstructuredObj("", "v1", "ConfigMap", ns, n)
}

func GetValidationFailedCondition(reason, message string) *metav1.Condition {
	return &metav1.Condition{
		Message: message,
		Reason:  reason,
		Type:    string(compositionv1alpha1.ValidationFailed),
	}
}

func GetReadyCondition(reason, message string) *metav1.Condition {
	return &metav1.Condition{
		Message: message,
		Reason:  reason,
		Type:    string(compositionv1alpha1.Ready),
	}
}

func GetErrorCondition(reason, message string) *metav1.Condition {
	return &metav1.Condition{
		Message: message,
		Reason:  reason,
		Type:    string(compositionv1alpha1.Error),
	}
}

func GetWaitingCondition(reason, message string) *metav1.Condition {
	return &metav1.Condition{
		Message: message,
		Reason:  reason,
		Type:    string(compositionv1alpha1.Waiting),
	}
}

func GetContextObj(context map[string]string) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "composition.google.com/v1alpha1",
			"kind":       "Context",
			"metadata": map[string]interface{}{
				"name":      "context",
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"project": context["gcpProject"],
			},
		},
	}
}
