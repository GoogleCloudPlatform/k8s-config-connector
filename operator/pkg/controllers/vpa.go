// Copyright 2025 Google LLC
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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	vpaGVK = schema.GroupVersionKind{
		Group:   "autoscaling.k8s.io",
		Version: "v1",
		Kind:    "VerticalPodAutoscaler",
	}
)

func EnsureVPAForStatefulSet(ctx context.Context, c client.Client, sts *appsv1.StatefulSet) error {
	return ensureVPA(ctx, c, sts.Namespace, sts.Name, "StatefulSet")
}

func EnsureVPAForDeployment(ctx context.Context, c client.Client, deployment *appsv1.Deployment) error {
	return ensureVPA(ctx, c, deployment.Namespace, deployment.Name, "Deployment")
}

func ensureVPA(ctx context.Context, c client.Client, namespace, name, kind string) error {
	vpa := &unstructured.Unstructured{}
	vpa.SetGroupVersionKind(vpaGVK)
	vpa.SetNamespace(namespace)
	vpa.SetName(name)

	err := c.Get(ctx, client.ObjectKeyFromObject(vpa), vpa)
	if err == nil {
		// VPA already exists
		return nil
	}
	if !apierrors.IsNotFound(err) {
		return fmt.Errorf("error getting VPA %s/%s: %w", namespace, name, err)
	}

	// VPA does not exist, create it
	vpa.Object["spec"] = map[string]interface{}{
		"targetRef": map[string]interface{}{
			"apiVersion": appsv1.SchemeGroupVersion.String(),
			"kind":       kind,
			"name":       name,
		},
		"updatePolicy": map[string]interface{}{
			"updateMode":  "Auto",
			"minReplicas": int64(1),
		},
	}

	if err := c.Create(ctx, vpa); err != nil {
		return fmt.Errorf("error creating VPA %s/%s: %w", namespace, name, err)
	}

	return nil
}

func GetVPARecommendations(ctx context.Context, c client.Client, namespace, name string) (map[string]corev1.ResourceRequirements, error) {
	vpa := &unstructured.Unstructured{}
	vpa.SetGroupVersionKind(vpaGVK)
	vpa.SetNamespace(namespace)
	vpa.SetName(name)

	if err := c.Get(ctx, client.ObjectKeyFromObject(vpa), vpa); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting VPA %s/%s: %w", namespace, name, err)
	}

	recommendation, found, err := unstructured.NestedMap(vpa.Object, "status", "recommendation")
	if err != nil {
		return nil, fmt.Errorf("error getting status.recommendation from VPA %s/%s: %w", namespace, name, err)
	}
	if !found {
		return nil, nil
	}

	containerRecommendations, found, err := unstructured.NestedSlice(recommendation, "containerRecommendations")
	if err != nil {
		return nil, fmt.Errorf("error getting containerRecommendations from VPA %s/%s: %w", namespace, name, err)
	}
	if !found {
		return nil, nil
	}

	result := make(map[string]corev1.ResourceRequirements)
	for _, r := range containerRecommendations {
		rec, ok := r.(map[string]interface{})
		if !ok {
			continue
		}
		containerName, found, err := unstructured.NestedString(rec, "containerName")
		if err != nil || !found {
			continue
		}

		target, found, err := unstructured.NestedMap(rec, "target")
		if err != nil || !found {
			continue
		}

		resources := corev1.ResourceRequirements{
			Limits:   make(corev1.ResourceList),
			Requests: make(corev1.ResourceList),
		}

		for k, v := range target {
			quantity, err := resource.ParseQuantity(fmt.Sprintf("%v", v))
			if err != nil {
				continue
			}
			// VPA target recommendations are usually treated as both requests and limits,
			// or just requests depending on configuration.
			// For "Auto" mode, it usually recommends requests.
			// VPA 'target' is the recommended request.
			// VPA 'uncappedTarget' is the recommended request without min/max policy constraints.
			// VPA 'lowerBound' and 'upperBound' are confidence intervals.
			// Usually, we set Requests = Target.
			// If we want to set Limits, we might need to handle it carefully (e.g. maintain ratio or just set same as requests if we want Guaranteed QoS).
			// The user said "honors the CPU and Memory limits specified in the VerticalPodAutoscaler".
			// I will assume they mean the 'target' recommendation should be applied.
			// And typically VPA recommends Requests.
			resources.Requests[corev1.ResourceName(k)] = quantity
			resources.Limits[corev1.ResourceName(k)] = quantity
		}
		result[containerName] = resources
	}

	return result, nil
}
