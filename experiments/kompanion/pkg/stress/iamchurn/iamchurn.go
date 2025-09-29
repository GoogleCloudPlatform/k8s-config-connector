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

package iamchurn

import (
	"context"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/kompanion/pkg/stress"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type IAMPolicyMemberChurn struct {
	RunID         string
	Config        *stress.StressConfig
	DynamicClient dynamic.Interface
}

func New(runID string, config *stress.StressConfig, dynamicClient dynamic.Interface) *IAMPolicyMemberChurn {
	return &IAMPolicyMemberChurn{
		RunID:         runID,
		Config:        config,
		DynamicClient: dynamicClient,
	}
}

func (s *IAMPolicyMemberChurn) Setup(ctx context.Context) error {
	log.Printf("Starting setup for IAMPolicyMemberChurn stress test with run ID: %s", s.RunID)

	for i := 0; i < s.Config.NumNamespaces; i++ {
		nsName := fmt.Sprintf("stress-test-%s-ns-%d", s.RunID, i)
		if err := s.createNamespace(ctx, nsName); err != nil {
			return err
		}

		for _, resourceToStress := range s.Config.ResourcesToStress {
			if err := s.createResourcesInNamespace(ctx, nsName, resourceToStress); err != nil {
				return err
			}
		}
	}

	log.Println("Setup complete.")
	return nil
}

func (s *IAMPolicyMemberChurn) Teardown(ctx context.Context) error {
	log.Printf("Starting teardown for IAMPolicyMemberChurn stress test with run ID: %s", s.RunID)
	labelSelector := fmt.Sprintf("kompanion-stress-run-id=%s", s.RunID)

	// Delete namespaces
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "namespaces"}
	if err := s.DynamicClient.Resource(gvr).DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{LabelSelector: labelSelector}); err != nil {
		return fmt.Errorf("failed to delete namespaces with label selector %s: %w", labelSelector, err)
	}

	// The deletion of namespaces should trigger the deletion of the other resources.
	// We can add explicit deletion of other resources if needed, but it's generally not necessary.

	log.Println("Teardown complete.")
	return nil
}

func (s *IAMPolicyMemberChurn) createNamespace(ctx context.Context, name string) error {
	log.Printf("Creating namespace: %s", name)
	ns := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": name,
				"labels": map[string]interface{}{
					"kompanion-stress-run-id": s.RunID,
				},
			},
		},
	}
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "namespaces"}
	_, err := s.DynamicClient.Resource(gvr).Create(ctx, ns, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create namespace %s: %w", name, err)
	}
	return nil
}

func (s *IAMPolicyMemberChurn) createResourcesInNamespace(ctx context.Context, nsName string, resourceToStress stress.ResourceToStress) error {
	log.Printf("Creating %d %s resources in namespace %s", resourceToStress.Count, resourceToStress.Kind, nsName)
	for i := 0; i < resourceToStress.Count; i++ {
		// Create IAMServiceAccount
		saName := fmt.Sprintf("stress-sa-%s-%d", s.RunID, i)
		sa := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"kind":       "IAMServiceAccount",
				"metadata": map[string]interface{}{
					"name":      saName,
					"namespace": nsName,
					"labels": map[string]interface{}{
						"kompanion-stress-run-id": s.RunID,
					},
				},
				"spec": map[string]interface{}{
					"displayName": saName,
				},
			},
		}
		gvr := schema.GroupVersionResource{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iamserviceaccounts"}
		_, err := s.DynamicClient.Resource(gvr).Namespace(nsName).Create(ctx, sa, metav1.CreateOptions{})
		if err != nil {
			return fmt.Errorf("failed to create IAMServiceAccount %s in namespace %s: %w", saName, nsName, err)
		}

		// Create IAMPolicyMember
		pmName := fmt.Sprintf("stress-pm-%s-%d", s.RunID, i)
		pm := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"kind":       "IAMPolicyMember",
				"metadata": map[string]interface{}{
					"name":      pmName,
					"namespace": nsName,
					"labels": map[string]interface{}{
						"kompanion-stress-run-id": s.RunID,
					},
				},
				"spec": map[string]interface{}{
					"memberFrom": map[string]interface{}{
						"serviceAccountRef": map[string]interface{}{
							"name": saName,
						},
					},
					"role": resourceToStress.Role,
					"resourceRef": map[string]interface{}{
						"kind":      s.Config.IAMReference.Kind,
						"name":      s.Config.IAMReference.Name,
						"namespace": s.Config.IAMReference.Namespace,
					},
				},
			},
		}
		gvr = schema.GroupVersionResource{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iampolicymembers"}
		_, err = s.DynamicClient.Resource(gvr).Namespace(nsName).Create(ctx, pm, metav1.CreateOptions{})
		if err != nil {
			return fmt.Errorf("failed to create IAMPolicyMember %s in namespace %s: %w", pmName, nsName, err)
		}
	}
	return nil
}
