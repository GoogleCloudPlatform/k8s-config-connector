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

package controllers

import (
	"context"
	"fmt"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

var (
	ValidConfigConnectorNamespacedName = types.NamespacedName{Name: k8s.ConfigConnectorAllowedName}
)

func GetConfigConnector(ctx context.Context, client client.Client, nn types.NamespacedName) (*corev1beta1.ConfigConnector, error) {
	cc := &corev1beta1.ConfigConnector{}
	if err := client.Get(ctx, nn, cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func RemoveOperatorFinalizer(o metav1.Object) (found bool) {
	var finalizers []string
	for _, f := range o.GetFinalizers() {
		if f != k8s.OperatorFinalizer {
			finalizers = append(finalizers, f)
		} else {
			found = true
		}
	}
	if found {
		o.SetFinalizers(finalizers)
	}
	return found
}

func EnsureOperatorFinalizer(o metav1.Object) (found bool) {
	for _, f := range o.GetFinalizers() {
		if f == k8s.OperatorFinalizer {
			return true
		}
	}
	o.SetFinalizers(append(o.GetFinalizers(), k8s.OperatorFinalizer))
	return false
}

func AnnotateServiceAccountObject(object *manifest.Object, gsa string) (*manifest.Object, error) {
	u := object.UnstructuredObject()
	annotations := u.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[k8s.WorkloadIdentityAnnotation] = gsa
	u.SetAnnotations(annotations)
	return manifest.NewObject(u)
}

func DeleteObject(ctx context.Context, c client.Client, obj client.Object) error {
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	name := obj.(metav1.Object).GetName()
	if err := c.Delete(ctx, obj, &client.DeleteOptions{}); err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("error deleting %v %v: %v", kind, name, err)
	}
	return removeOperatorFinalizerIfPresent(ctx, c, obj)
}

/*
 * some of the critical resources, such as role bindings, are in the customer namespace, we protect them with an
 * operator finalizer to ensure that the operator can control the time when they are removed.
 */
func removeOperatorFinalizerIfPresent(ctx context.Context, c client.Client, obj client.Object) error {
	found := RemoveOperatorFinalizer(obj)
	if !found {
		return nil
	}
	if err := c.Update(ctx, obj); err != nil {
		return fmt.Errorf("error removing operator finalizer from %v %v: %w",
			obj.GetObjectKind().GroupVersionKind().Kind, obj.GetName(), err)
	}
	return nil
}

func IsControllerManagerStatefulSet(obj *manifest.Object) bool {
	if obj.Kind != "StatefulSet" {
		return false
	}
	labels := obj.UnstructuredObject().GetLabels()
	return labels[k8s.KCCSystemComponentLabel] == k8s.KCCControllerManagerComponent
}

func IsControllerManagerService(obj *manifest.Object) bool {
	if obj.Kind != "Service" {
		return false
	}
	return obj.GetName() == k8s.NamespacedManagerServiceTmpl
}
