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

package main

import (
	"context"
	"testing"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	spannerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/spanner/v1beta1"
	computeclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/typed/compute/v1beta1"
	spannerclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/typed/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"

	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var mgr manager.Manager

func TestSpannerInstanceGoClient(t *testing.T) {
	t.Parallel()
	client := spannerclient.NewForConfigOrDie(mgr.GetConfig())
	testId := testvariable.NewUniqueId()
	spannerInstance := spannerv1beta1.SpannerInstance{
		// TypeMeta (Kind/APIVersion) is automatically filled out
		ObjectMeta: v1.ObjectMeta{
			Name:      "spannerinstance",
			Namespace: testId,
			Labels:    map[string]string{"Key": "value"},
		},
		Spec: spannerv1beta1.SpannerInstanceSpec{
			Config:      "regional-us-west1",
			DisplayName: "Unique Display Name",
		},
		Status: spannerv1beta1.SpannerInstanceStatus{},
	}
	testcontroller.SetupNamespaceForDefaultProject(t, mgr.GetClient(), testId)
	if _, err := client.SpannerInstances(testId).Create(context.TODO(), &spannerInstance, v1.CreateOptions{}); err != nil {
		t.Fatalf("Error creating SpannerInstance: %v", err)
	}

	name := k8s.GetNamespacedName(&spannerInstance)
	u := unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "SpannerInstance",
			"apiVersion": "spanner.cnrm.cloud.google.com/v1beta1",
		},
	}
	if err := mgr.GetClient().Get(context.TODO(), name, &u); err != nil {
		t.Fatalf("Error getting spanner instance: %v", err)
	}
	s, found, err := unstructured.NestedString(u.Object, "spec", "displayName")
	if err != nil {
		t.Errorf("Error getting nested field: %v ", err)
	}
	if !found {
		t.Errorf("Nested field not found")
	}

	// Validate displayName fields match
	if s != spannerInstance.Spec.DisplayName {
		t.Errorf("Results mismatch: got %v, expected %v", s, spannerInstance.Spec.DisplayName)
	}

	// Cleanup
	if err = client.SpannerInstances(testId).Delete(context.TODO(), "spannerinstance", v1.DeleteOptions{}); err != nil {
		t.Errorf("Error deleting Spanner Instance: %v", err)
	}
	err = mgr.GetClient().Get(context.TODO(), name, &u)
	if err == nil || !errors.IsNotFound(err) {
		t.Errorf("Unexpected error value: '%v'", err)
	}
}

func TestComputeInstanceGoClient(t *testing.T) {
	client := computeclient.NewForConfigOrDie(mgr.GetConfig())
	testId := testvariable.NewUniqueId()
	autoDelete := true
	bootDiskSize := 20
	bootDiskType := "pd-ssd"
	machineType := "n1-standard-1"
	zone := "europe-west1-b"
	computeInstance := computev1beta1.ComputeInstance{
		// TypeMeta (Kind/APIVersion) is automatically filled out
		ObjectMeta: v1.ObjectMeta{
			Name:      "computeinstance",
			Namespace: testId,
			Labels:    map[string]string{"key": "value"},
		},
		Spec: computev1beta1.ComputeInstanceSpec{
			BootDisk: &computev1beta1.InstanceBootDisk{
				AutoDelete: &autoDelete,
				InitializeParams: &computev1beta1.InstanceInitializeParams{
					Size: &bootDiskSize,
					SourceImageRef: &v1alpha1.ResourceRef{
						External: "debian-cloud/debian-9",
					},
					Type: &bootDiskType,
				},
			},
			MachineType: &machineType,
			Zone:        &zone,
			NetworkInterface: []computev1beta1.InstanceNetworkInterface{
				{
					NetworkRef: &v1alpha1.ResourceRef{
						External: "default",
					},
					SubnetworkRef: &v1alpha1.ResourceRef{
						External: "default",
					},
				},
			},
		},
	}
	testcontroller.SetupNamespaceForDefaultProject(t, mgr.GetClient(), testId)
	if _, err := client.ComputeInstances(testId).Create(context.TODO(), &computeInstance, v1.CreateOptions{}); err != nil {
		t.Fatalf("Error creating ComputeInstance: %v", err)
	}

	name := k8s.GetNamespacedName(&computeInstance)
	u := unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "ComputeInstance",
			"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
		},
	}
	if err := mgr.GetClient().Get(context.TODO(), name, &u); err != nil {
		t.Fatalf("Error getting compute instance: %v", err)
	}

	s, found, err := unstructured.NestedString(u.Object, "spec", "bootDisk", "initializeParams", "type")
	if err != nil {
		t.Errorf("Error getting nested field: %v ", err)
	}
	if !found {
		t.Errorf("Nested field not found")
	}

	// Validate fields match
	if s != bootDiskType {
		t.Errorf("Results mismatch: got %v, expected %v", s, bootDiskType)
	}

	// Cleanup
	if err = client.ComputeInstances(testId).Delete(context.TODO(), "computeinstance", v1.DeleteOptions{}); err != nil {
		t.Errorf("Error deleting Compute Instance: %v", err)
	}
	err = mgr.GetClient().Get(context.TODO(), name, &u)
	if err == nil || !errors.IsNotFound(err) {
		t.Errorf("Unexpected error value: '%v'", err)
	}
}

func TestMain(m *testing.M) {
	testmain.TestMainForIntegrationTests(m, &mgr)
}
