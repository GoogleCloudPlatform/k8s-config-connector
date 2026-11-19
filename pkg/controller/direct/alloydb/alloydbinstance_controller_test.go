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

package alloydb

import (
	"context"
	"reflect"
	"testing"
	"time"

	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	computev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func newTestProjectMapper() *projects.ProjectMapper {
	cache := projects.NewProjectCache(nil, time.Hour)
	cache.InsertForTest("my-project", 123456789)
	cache.InsertForTest("other-project", 987654321)
	return projects.NewProjectMapper(cache)
}

func TestAlloyDBInstancePSCMappings(t *testing.T) {
	mapCtx := &direct.MapContext{}
	spec := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{
				{External: "123456789"},
			},
			PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "my-project"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/my-project/global/networks/default"},
				},
			},
			PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
				{
					NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{
						External: "projects/my-project/regions/us-central1/networkAttachments/my-na",
					},
				},
			},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_ENABLED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("ENABLED"),
		},
	}

	protoInst := AlloyDBInstanceSpec_ToProto(mapCtx, spec)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in AlloyDBInstanceSpec_ToProto: %v", err)
	}
	if protoInst.GetPscInstanceConfig().GetPscAutoDnsState() != pb.Instance_PSC_AUTO_DNS_STATE_ENABLED {
		t.Errorf("got PscAutoDnsState = %v, want %v", protoInst.GetPscInstanceConfig().GetPscAutoDnsState(), pb.Instance_PSC_AUTO_DNS_STATE_ENABLED)
	}
	if protoInst.GetPscInstanceConfig().GetPscAutoConnectionPolicyState() != pb.Instance_PscInstanceConfig_ENABLED {
		t.Errorf("got PscAutoConnectionPolicyState = %v, want %v", protoInst.GetPscInstanceConfig().GetPscAutoConnectionPolicyState(), pb.Instance_PscInstanceConfig_ENABLED)
	}

	roundTripSpec := AlloyDBInstanceSpec_FromProto(mapCtx, protoInst)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in AlloyDBInstanceSpec_FromProto: %v", err)
	}
	if !reflect.DeepEqual(roundTripSpec.PSCInstanceConfig, spec.PSCInstanceConfig) {
		t.Errorf("roundTripSpec.PSCInstanceConfig = %+v, want %+v", roundTripSpec.PSCInstanceConfig, spec.PSCInstanceConfig)
	}

	// Verify ObservedState mapping for PSCInstanceConfig and PSCInstanceInfo.
	protoInst.PscInstanceConfig.ServiceAttachmentLink = "projects/123/regions/us-central1/serviceAttachments/sa-1"
	protoInst.PscInstanceConfig.PscDnsName = "uid.uid.us-central1.alloydb-psc.goog."
	protoInst.PscInstanceConfig.PscAutoConnections[0].IpAddress = "10.0.0.10"
	protoInst.PscInstanceConfig.PscAutoConnections[0].Status = "ACTIVE"
	protoInst.PscInstanceConfig.PscAutoConnections[0].ConsumerNetworkStatus = "VALID"
	protoInst.PscInstanceInfo = &pb.Instance_PscInstanceInfo{
		EffectivePscAutoDnsEnabled:       true,
		PscAutoDnsNames:                  []string{"inst.cluster.us-central1.my-project.alloydb.goog."},
		EffectivePscAutoConnectionPolicy: true,
		ServiceConnectionPolicy:          "projects/my-project/locations/us-central1/serviceConnectionPolicies/default-scp",
	}

	status := AlloyDBInstanceStatus_FromProto(mapCtx, protoInst)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in AlloyDBInstanceStatus_FromProto: %v", err)
	}
	if status.ObservedState == nil || status.ObservedState.PSCInstanceConfig == nil || status.ObservedState.PSCInstanceInfo == nil {
		t.Fatalf("expected ObservedState.PSCInstanceConfig and ObservedState.PSCInstanceInfo to be populated, got %+v", status.ObservedState)
	}
	if !direct.ValueOf(status.ObservedState.PSCInstanceInfo.EffectivePSCAutoDNSEnabled) {
		t.Errorf("got EffectivePSCAutoDNSEnabled = false, want true")
	}
	if !direct.ValueOf(status.ObservedState.PSCInstanceInfo.EffectivePSCAutoConnectionPolicy) {
		t.Errorf("got EffectivePSCAutoConnectionPolicy = false, want true")
	}
	if got := direct.ValueOf(status.ObservedState.PSCInstanceInfo.ServiceConnectionPolicy); got != "projects/my-project/locations/us-central1/serviceConnectionPolicies/default-scp" {
		t.Errorf("got ServiceConnectionPolicy = %q, want projects/my-project/locations/us-central1/serviceConnectionPolicies/default-scp", got)
	}
}

func TestResolvePSCReferences(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()

	projObj := &unstructured.Unstructured{}
	projObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Project",
	})
	projObj.SetNamespace("test-ns")
	projObj.SetName("k8s-project")
	_ = unstructured.SetNestedField(projObj.Object, "resolved-proj-id", "spec", "resourceID")

	netObj := &unstructured.Unstructured{}
	netObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	})
	netObj.SetNamespace("test-ns")
	netObj.SetName("k8s-net")
	_ = unstructured.SetNestedField(netObj.Object, "https://www.googleapis.com/compute/v1/projects/resolved-proj-id/global/networks/k8s-net", "status", "selfLink")

	naObj := &unstructured.Unstructured{}
	naObj.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "ComputeNetworkAttachment",
	})
	naObj.SetNamespace("test-ns")
	naObj.SetName("k8s-na")
	_ = unstructured.SetNestedField(naObj.Object, "projects/resolved-proj-id/regions/europe-west1/networkAttachments/k8s-na", "status", "externalRef")

	reader := fake.NewClientBuilder().WithScheme(scheme).WithObjects(projObj, netObj, naObj).Build()

	inst := &krm.AlloyDBInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-inst",
			Namespace: "test-ns",
		},
		Spec: krm.AlloyDBInstanceSpec{
			PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
				AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{
					{Name: "k8s-project"},
					{External: "projects/other-project"},
				},
				PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
					{
						ConsumerProjectRef: &refsv1beta1.ProjectRef{Name: "k8s-project"},
						ConsumerNetworkRef: &computerefs.ComputeNetworkRef{Name: "k8s-net"},
					},
					{
						ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "projects/123456789"},
						ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "https://www.googleapis.com/compute/v1/projects/123456789/global/networks/ext-net"},
					},
				},
				PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
					{
						NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{Name: "k8s-na"},
					},
					{
						NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{
							External: "https://www.googleapis.com/compute/v1/projects/123456789/regions/europe-west1/networkAttachments/ext-na",
						},
					},
				},
			},
		},
	}

	if err := resolvePSCReferences(ctx, reader, inst); err != nil {
		t.Fatalf("resolvePSCReferences failed: %v", err)
	}

	psc := inst.Spec.PSCInstanceConfig
	if got := psc.AllowedConsumerProjectRefs[0].External; got != "resolved-proj-id" {
		t.Errorf("AllowedConsumerProjectRefs[0].External = %q, want %q", got, "resolved-proj-id")
	}
	if got := psc.AllowedConsumerProjectRefs[1].External; got != "other-project" {
		t.Errorf("AllowedConsumerProjectRefs[1].External = %q, want %q", got, "other-project")
	}
	if got := psc.PSCAutoConnections[0].ConsumerProjectRef.External; got != "resolved-proj-id" {
		t.Errorf("PSCAutoConnections[0].ConsumerProjectRef.External = %q, want %q", got, "resolved-proj-id")
	}
	if got := psc.PSCAutoConnections[0].ConsumerNetworkRef.External; got != "projects/resolved-proj-id/global/networks/k8s-net" {
		t.Errorf("PSCAutoConnections[0].ConsumerNetworkRef.External = %q, want %q", got, "projects/resolved-proj-id/global/networks/k8s-net")
	}
	if got := psc.PSCAutoConnections[1].ConsumerProjectRef.External; got != "123456789" {
		t.Errorf("PSCAutoConnections[1].ConsumerProjectRef.External = %q, want %q", got, "123456789")
	}
	if got := psc.PSCAutoConnections[1].ConsumerNetworkRef.External; got != "projects/123456789/global/networks/ext-net" {
		t.Errorf("PSCAutoConnections[1].ConsumerNetworkRef.External = %q, want %q", got, "projects/123456789/global/networks/ext-net")
	}
	if got := psc.PSCInterfaceConfigs[0].NetworkAttachmentResourceRef.External; got != "projects/resolved-proj-id/regions/europe-west1/networkAttachments/k8s-na" {
		t.Errorf("PSCInterfaceConfigs[0].NetworkAttachmentResourceRef.External = %q, want %q", got, "projects/resolved-proj-id/regions/europe-west1/networkAttachments/k8s-na")
	}
	if got := psc.PSCInterfaceConfigs[1].NetworkAttachmentResourceRef.External; got != "projects/123456789/regions/europe-west1/networkAttachments/ext-na" {
		t.Errorf("PSCInterfaceConfigs[1].NetworkAttachmentResourceRef.External = %q, want %q", got, "projects/123456789/regions/europe-west1/networkAttachments/ext-na")
	}
}

func TestCompareInstancePSC(t *testing.T) {
	ctx := context.Background()
	mapper := newTestProjectMapper()

	actual := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{
				{External: "123456789"},
			},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_ENABLED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("ENABLED"),
		},
	}
	desired := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{
				{External: "123456789"},
				{External: "987654321"},
			},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_DISABLED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("DISABLED"),
		},
	}

	paths, err := compareInstance(ctx, actual, desired, mapper)
	if err != nil {
		t.Fatalf("compareInstance returned unexpected error: %v", err)
	}
	wantPaths := []string{
		"psc_instance_config.allowed_consumer_projects",
		"psc_instance_config.psc_auto_dns_state",
		"psc_instance_config.psc_auto_connection_policy_state",
	}
	if !reflect.DeepEqual(paths, wantPaths) {
		t.Errorf("compareInstance paths = %v, want %v", paths, wantPaths)
	}
}

func TestCompareInstancePSCNormalizationAndUnspecified(t *testing.T) {
	ctx := context.Background()
	mapper := newTestProjectMapper()

	// GCP returns project numbers in allowedConsumerProjects and project IDs in pscAutoConnections.
	actual := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{
				{External: "123456789"},
				{External: "987654321"},
			},
			PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "my-project"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/my-project/global/networks/net-a"},
				},
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "other-project"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/other-project/global/networks/net-b"},
				},
			},
			PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
				{
					NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{
						External: "projects/123456789/regions/us-central1/networkAttachments/my-na",
					},
				},
			},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_ENABLED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("ENABLED"),
		},
	}

	// Desired uses project IDs + duplicates + reversed order in allowedConsumerProjectRefs,
	// project numbers + full Compute URI + duplicates + reversed order in pscAutoConnections,
	// project ID + full Compute URI in pscInterfaceConfigs,
	// and *_UNSPECIFIED on enums when actual has non-nil values.
	desired := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{
				{External: "other-project"},
				{External: "my-project"},
				{External: "123456789"},
			},
			PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "987654321"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "https://www.googleapis.com/compute/v1/projects/987654321/global/networks/net-b"},
				},
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "my-project"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/123456789/global/networks/net-a"},
				},
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "123456789"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/my-project/global/networks/net-a"},
				},
			},
			PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
				{
					NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{
						External: "https://www.googleapis.com/compute/v1/projects/my-project/regions/us-central1/networkAttachments/my-na",
					},
				},
			},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_UNSPECIFIED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("PSC_AUTO_CONNECTION_POLICY_STATE_UNSPECIFIED"),
		},
	}

	paths, err := compareInstance(ctx, actual, desired, mapper)
	if err != nil {
		t.Fatalf("compareInstance returned unexpected error: %v", err)
	}
	if len(paths) != 0 {
		t.Errorf("expected zero updatePaths after normalization, got %v", paths)
	}

	// Also verify *_UNSPECIFIED when actual enum pointers are nil.
	actualNilEnums := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{{External: "123456789"}},
		},
	}
	desiredNilEnums := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs:   []refsv1beta1.ProjectRef{{External: "my-project"}},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_UNSPECIFIED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("PSC_AUTO_CONNECTION_POLICY_STATE_UNSPECIFIED"),
		},
	}
	pathsNilEnums, err := compareInstance(ctx, actualNilEnums, desiredNilEnums, mapper)
	if err != nil {
		t.Fatalf("compareInstance returned unexpected error for nil actual enums: %v", err)
	}
	if len(pathsNilEnums) != 0 {
		t.Errorf("expected zero updatePaths for UNSPECIFIED with nil actual enums, got %v", pathsNilEnums)
	}

	// Verify genuine diffs on pscAutoConnections and pscInterfaceConfigs are detected.
	desiredChanged := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "my-project"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/my-project/global/networks/different-net"},
				},
			},
			PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
				{
					NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{
						External: "projects/my-project/regions/us-central1/networkAttachments/different-na",
					},
				},
			},
		},
	}
	changedPaths, err := compareInstance(ctx, actual, desiredChanged, mapper)
	if err != nil {
		t.Fatalf("compareInstance returned unexpected error for genuine changes: %v", err)
	}
	wantChangedPaths := []string{
		"psc_instance_config.psc_interface_configs",
		"psc_instance_config.psc_auto_connections",
	}
	if !reflect.DeepEqual(changedPaths, wantChangedPaths) {
		t.Errorf("compareInstance paths = %v, want %v", changedPaths, wantChangedPaths)
	}

	// Verify unresolvable external project IDs fall back to raw comparison without failing reconciliation.
	actualExternal := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{{External: "555666777"}},
			PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "unreadable-external-project"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/unreadable-external-project/global/networks/ext-net"},
				},
			},
			PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
				{
					NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{
						External: "projects/555666777/regions/us-central1/networkAttachments/ext-na",
					},
				},
			},
		},
	}
	desiredExternal := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjectRefs: []refsv1beta1.ProjectRef{{External: "unreadable-external-project"}},
			PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
				{
					ConsumerProjectRef: &refsv1beta1.ProjectRef{External: "555666777"},
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{External: "projects/555666777/global/networks/ext-net"},
				},
			},
			PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
				{
					NetworkAttachmentResourceRef: &computev1alpha1.ComputeNetworkAttachmentRef{
						External: "projects/unreadable-external-project/regions/us-central1/networkAttachments/ext-na",
					},
				},
			},
		},
	}
	fallbackPaths, err := compareInstance(ctx, actualExternal, desiredExternal, mapper)
	if err != nil {
		t.Fatalf("compareInstance returned unexpected error on unresolvable project ID: %v", err)
	}
	wantFallbackPaths := []string{
		"psc_instance_config.allowed_consumer_projects",
		"psc_instance_config.psc_interface_configs",
		"psc_instance_config.psc_auto_connections",
	}
	if !reflect.DeepEqual(fallbackPaths, wantFallbackPaths) {
		t.Errorf("compareInstance fallback paths = %v, want %v", fallbackPaths, wantFallbackPaths)
	}
}
