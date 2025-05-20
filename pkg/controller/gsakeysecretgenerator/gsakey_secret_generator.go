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

package gsakeysecretgenerator

import (
	"context"
	"encoding/base64"
	"fmt"

	kontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

/*
This controller watches service account key CRD and reconciles the corresponding secrets against private keys.
The default behaviour is to generate a secret from the private key; the secret has the same name as service account key object
and lives in the same namespace. Users can annotate the service account key object with cnrm.cloud.google.com/disable-secret-creation: true
to disable the secret creation if needed.
*/
const controllerName = "gsakeysecretgenerator"
const createGsaKeySecretAnnotation = "cnrm.cloud.google.com/create-gsa-key-secret"
const gsaKeySecretDataKeyAnnotation = "cnrm.cloud.google.com/gsa-key-secret-data-key"
const eventMessageTemplate = "secret %v in namespace %v %v"

var logger = klog.Log.WithName(controllerName)

func Add(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition, deps *kontroller.Deps) error {
	if deps.JitterGen == nil {
		deps.JitterGen = &jitter.SimpleJitterGenerator{}
	}

	r := newReconciler(mgr, crd, deps.JitterGen)
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       crd.Spec.Names.Kind,
			"apiVersion": k8s.GetAPIVersionFromCRD(crd),
		},
	}
	_, err := builder.
		ControllerManagedBy(mgr).
		Named(controllerName).
		WithOptions(controller.Options{MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles, RateLimiter: ratelimiter.NewRateLimiter()}).
		For(obj, builder.OnlyMetadata).
		Build(r)
	if err != nil {
		return fmt.Errorf("error creating new controller: %w", err)
	}
	logger.Info("added a controller for service-account-key-to-secret")
	return nil
}

// newReconciler returns a new reconcile.Reconciler.
func newReconciler(mgr manager.Manager, crd *apiextensions.CustomResourceDefinition, jg jitter.Generator) reconcile.Reconciler {
	return &ReconcileSecret{
		Client:     mgr.GetClient(),
		kind:       crd.Spec.Names.Kind,
		apiVersion: k8s.GetAPIVersionFromCRD(crd),
		recorder:   mgr.GetEventRecorderFor(controllerName),
		jitterGen:  jg,
	}
}

type ReconcileSecret struct {
	client.Client
	kind       string
	apiVersion string
	recorder   record.EventRecorder
	jitterGen  jitter.Generator
}

func (r *ReconcileSecret) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	logger.Info("starting reconcile", "resource", request.NamespacedName)
	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       r.kind,
			"apiVersion": r.apiVersion,
		},
	}
	err := r.Get(ctx, request.NamespacedName, u)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, fmt.Errorf("error getting KCC object from API server: %w", err)
	}

	// exit early if annotation says not creating the secret
	if val, ok := k8s.GetAnnotation(createGsaKeySecretAnnotation, u); ok && val == "false" {
		return reconcile.Result{}, nil
	}
	// if service account key object is marked as deleted, no action needed
	if !u.GetDeletionTimestamp().IsZero() {
		return reconcile.Result{}, nil
	}

	// if private_key status field is not filled, skip it
	key, ok, err := getPrivateKey(u)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error finding private key from service account key resource %v: %w", request.NamespacedName, err)
	}
	if !ok {
		logger.Info("no private key is found from service account key. No secret will be created.", "resource", request.NamespacedName)
		return reconcile.Result{}, nil
	}

	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error decoding the private key: %w", err)
	}
	secretDataKey := "key.json"
	if val, ok := k8s.GetAnnotation(gsaKeySecretDataKeyAnnotation, u); ok {
		secretDataKey = val
	}
	secret := &corev1.Secret{
		Type: corev1.SecretTypeOpaque,
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name,
			Namespace: request.Namespace,
			Labels: map[string]string{
				label.CnrmManagedKey: "true",
			},
			OwnerReferences: []metav1.OwnerReference{{
				APIVersion: r.apiVersion,
				Kind:       r.kind,
				Name:       request.Name,
				UID:        u.GetUID(),
			}},
		},
		Data: map[string][]byte{
			secretDataKey: b,
		},
	}

	originalSecret := &corev1.Secret{}
	if err = r.Get(ctx, request.NamespacedName, originalSecret); err == nil {
		logger.Info("updating the secret", "resource", request.NamespacedName)
		if err = r.Update(ctx, secret); err != nil {
			r.recorder.Eventf(u, corev1.EventTypeWarning, k8s.UpdateFailed, eventMessageTemplate, u.GetName(), u.GetNamespace(), fmt.Errorf("update call failed: %w", err))
			return reconcile.Result{}, err
		}
		jitteredPeriod := r.jitterGen.WatchJitteredTimeout()
		logger.Info("successfully finished reconcile", "time to next reconciliation", jitteredPeriod)
		return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
	}

	if !errors.IsNotFound(err) {
		return reconcile.Result{}, err
	}
	logger.Info("creating the secret", "resource", request.NamespacedName)
	if err = r.Create(ctx, secret); err != nil {
		r.recorder.Eventf(u, corev1.EventTypeWarning, k8s.CreateFailed, eventMessageTemplate, u.GetName(), u.GetNamespace(), fmt.Sprintf(k8s.CreateFailedMessageTmpl, err))
		return reconcile.Result{}, err
	}
	r.recorder.Eventf(u, corev1.EventTypeNormal, k8s.Created, eventMessageTemplate, u.GetName(), u.GetNamespace(), k8s.CreatedMessage)
	jitteredPeriod := r.jitterGen.WatchJitteredTimeout()
	logger.Info("successfully finished reconcile", "time to next reconciliation", jitteredPeriod)
	return reconcile.Result{RequeueAfter: jitteredPeriod}, nil
}

func getPrivateKey(obj *unstructured.Unstructured) (string, bool, error) {
	if val, found, err := unstructured.NestedString(obj.Object, "status", "privateKey"); found || err != nil {
		return val, found, err
	}
	return unstructured.NestedString(obj.Object, "status", "privateKeyEncrypted")
}
