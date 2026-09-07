// Copyright 2026 Google LLC
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

package configconnector

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func loadComponentManifest(t *testing.T, relPaths ...string) *manifest.Objects {
	t.Helper()
	ctx := context.Background()
	m := testcontroller.ParseObjects(ctx, t, testcontroller.ClusterModeComponents)
	for _, relPath := range relPaths {
		raw, err := os.ReadFile(filepath.Join("..", "..", "..", "..", relPath))
		if err != nil {
			t.Fatalf("failed to read %s: %v", relPath, err)
		}
		extra, err := manifest.ParseObjects(ctx, string(raw))
		if err != nil {
			t.Fatalf("failed to parse %s: %v", relPath, err)
		}
		m.Items = append(m.Items, extra.Items...)
	}
	return m
}

func TestMCLSmoke(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	r := &Reconciler{}

	componentFiles := []string{
		"config/installbundle/components/mcl-controller/crd.yaml",
		"config/installbundle/components/mcl-controller/deployment.yaml",
		"config/installbundle/components/syncer-controller/crd.yaml",
		"config/installbundle/components/syncer-controller/deployment.yaml",
	}

	// 1. When MultiClusterLease is enabled, MCL and Syncer CRDs + Deployments are retained
	// and selectCRDsByVersion preserves both CRDs.
	mEnabled := loadComponentManifest(t, componentFiles...)
	if err := r.selectCRDsByVersion(mEnabled, "v1beta1"); err != nil {
		t.Fatalf("selectCRDsByVersion failed: %v", err)
	}

	ccEnabled := &corev1beta1.ConfigConnector{
		Spec: corev1beta1.ConfigConnectorSpec{
			Mode: "cluster",
			Experiments: &corev1beta1.CCExperiments{
				MultiClusterLease: &corev1beta1.MultiClusterLeaseSpec{
					LeaseName:                "test-lease",
					Namespace:                "cnrm-system",
					ClusterCandidateIdentity: "test-cluster",
				},
			},
		},
	}
	if err := r.applyExperiments(ctx, ccEnabled, mEnabled); err != nil {
		t.Fatalf("applyExperiments failed: %v", err)
	}

	foundMCLCRD := false
	foundSyncerCRD := false
	foundMCLDeployment := false
	foundSyncerDeployment := false
	foundConfiguredManager := false

	for _, item := range mEnabled.Items {
		switch {
		case item.Kind == "CustomResourceDefinition" && item.GetName() == "multiclusterleases.multicluster.core.cnrm.cloud.google.com":
			foundMCLCRD = true
		case item.Kind == "CustomResourceDefinition" && item.GetName() == "krmsyncers.syncer.gkelabs.io":
			foundSyncerCRD = true
		case item.Kind == "Deployment" && item.GetName() == "multiclusterlease-controller-manager":
			foundMCLDeployment = true
		case item.Kind == "Deployment" && item.GetName() == "syncer-controller-manager":
			foundSyncerDeployment = true
		case IsControllerManagerStatefulSet(item):
			foundConfiguredManager = true
			templateAnnotations, _, _ := unstructured.NestedStringMap(item.UnstructuredObject().Object, "spec", "template", "metadata", "annotations")
			if got := templateAnnotations["cnrm.cloud.google.com/lease-name"]; got != "test-lease" {
				t.Errorf("unexpected lease-name annotation: got %q, want %q", got, "test-lease")
			}
		}
	}

	if !foundMCLCRD {
		t.Errorf("expected multiclusterleases CRD to be present when MultiClusterLease is enabled")
	}
	if !foundSyncerCRD {
		t.Errorf("expected krmsyncers CRD to be present when MultiClusterLease is enabled")
	}
	if !foundMCLDeployment {
		t.Errorf("expected multiclusterlease-controller-manager Deployment to be present when MultiClusterLease is enabled")
	}
	if !foundSyncerDeployment {
		t.Errorf("expected syncer-controller-manager Deployment to be present when MultiClusterLease is enabled")
	}
	if !foundConfiguredManager {
		t.Errorf("expected cnrm-controller-manager StatefulSet to be present")
	}

	// 2. When MultiClusterLease is disabled, MCL and Syncer components are filtered out.
	mDisabled := loadComponentManifest(t, componentFiles...)
	ccDisabled := &corev1beta1.ConfigConnector{
		Spec: corev1beta1.ConfigConnectorSpec{
			Mode: "cluster",
		},
	}
	if err := r.applyExperiments(ctx, ccDisabled, mDisabled); err != nil {
		t.Fatalf("applyExperiments (disabled) failed: %v", err)
	}
	for _, item := range mDisabled.Items {
		labels := item.UnstructuredObject().GetLabels()
		if comp := labels[k8s.KCCSystemComponentLabel]; comp == k8s.KCCMCLControllerComponent || comp == k8s.KCCSyncerControllerComponent {
			t.Errorf("expected component %s (%s/%s) to be filtered out when MultiClusterLease is disabled", comp, item.Kind, item.GetName())
		}
	}
}
