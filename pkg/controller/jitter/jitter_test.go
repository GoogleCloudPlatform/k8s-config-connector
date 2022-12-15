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

package jitter_test

import (
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"

	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func TestGenerateTFJitteredReenqueuePeriod(t *testing.T) {
	t.Parallel()
	gvk := schema.GroupVersionKind{
		Group:   "test1.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "Test1Foo",
	}

	smLoader := servicemappingloader.NewFromServiceMappings(test.FakeServiceMappings())
	duration := jitter.GenerateJitteredReenqueuePeriod(gvk, smLoader, nil)

	if duration > time.Duration(100)/2*3*time.Second || duration < time.Duration(100)/2*time.Second {
		t.Fatalf("got unexpected time duration %v for gvk %v", duration, gvk)
	}
}

func TestGenerateDCLJitteredReenqueuePeriod(t *testing.T) {
	t.Parallel()
	gvk := schema.GroupVersionKind{
		Group:   "test2.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Test2Bar",
	}
	var desiredIntervalInSeconds uint32 = 100

	serviceList := []dclmetadata.ServiceMetadata{
		{
			Name:                 "Test2",
			APIVersion:           "v1beta1",
			ServiceNameUsedByDCL: "test2",
			DCLVersion:           "ga",
			Resources: []dclmetadata.Resource{
				{
					Kind:                            "Test2Bar",
					ReconciliationIntervalInSeconds: &desiredIntervalInSeconds,
				},
			},
		},
	}

	serviceMetadataLoader := dclmetadata.NewFromServiceList(serviceList)
	duration := jitter.GenerateJitteredReenqueuePeriod(gvk, nil, serviceMetadataLoader)

	if duration > time.Duration(desiredIntervalInSeconds)/2*3*time.Second || duration < time.Duration(desiredIntervalInSeconds)/2*time.Second {
		t.Fatalf("got unexpected time duration %v for gvk %v", duration, gvk)
	}
}

func TestGenerateIAMJitteredReenqueuePeriod(t *testing.T) {
	t.Parallel()
	gvk := iamv1beta1.IAMPolicyGVK

	duration := jitter.GenerateJitteredReenqueuePeriod(gvk, nil, nil)

	if duration > k8s.MeanReconcileReenqueuePeriod/2*3 || duration < k8s.MeanReconcileReenqueuePeriod/2 {
		t.Fatalf("got unexpected time duration %v for gvk %v", duration, gvk)
	}
}

func TestGenerateDefaultJitteredReenqueuePeriod(t *testing.T) {
	t.Parallel()
	// gvk1 from TF servicemapping
	gvk1 := schema.GroupVersionKind{
		Group:   "test1.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "Test1Bar",
	}

	// gvk2 from DCL service metadata
	gvk2 := schema.GroupVersionKind{
		Group:   "dcltest1.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "DclTest2Bar",
	}

	// gvk3 none TF/DCL
	gvk3 := schema.GroupVersionKind{
		Group:   "foobar.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "FooBarBaz",
	}

	smLoader := servicemappingloader.NewFromServiceMappings(test.FakeServiceMappings())
	serviceList := []dclmetadata.ServiceMetadata{
		{
			Name:                 "DclTest1",
			APIVersion:           "v1beta1",
			ServiceNameUsedByDCL: "DclTest1",
			DCLVersion:           "ga",
			Resources: []dclmetadata.Resource{
				{
					Kind: "DclTest1Bar",
				},
			},
		},
	}

	serviceMetadataLoader := dclmetadata.NewFromServiceList(serviceList)

	// Test default value for TF resources
	duration := jitter.GenerateJitteredReenqueuePeriod(gvk1, smLoader, serviceMetadataLoader)

	if duration > k8s.MeanReconcileReenqueuePeriod/2*3 || duration < k8s.MeanReconcileReenqueuePeriod/2 {
		t.Fatalf("got unexpected time duration %v for gvk %v", duration, gvk1)
	}

	// Test default value for DCL resources
	duration = jitter.GenerateJitteredReenqueuePeriod(gvk2, smLoader, serviceMetadataLoader)

	if duration > k8s.MeanReconcileReenqueuePeriod/2*3 || duration < k8s.MeanReconcileReenqueuePeriod/2 {
		t.Fatalf("got unexpected time duration %v for gvk %v", duration, gvk2)
	}

	// Test default value for GVK not found in servicemapping/service metadata
	duration = jitter.GenerateJitteredReenqueuePeriod(gvk3, smLoader, serviceMetadataLoader)

	if duration > k8s.MeanReconcileReenqueuePeriod/2*3 || duration < k8s.MeanReconcileReenqueuePeriod/2 {
		t.Fatalf("got unexpected time duration %v for gvk %v", duration, gvk2)
	}
}
