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

package kccmanager

import (
	"context"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/leaderelection"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type mockManager struct {
	manager.Manager
	started bool
	err     error
}

func (m *mockManager) Start(ctx context.Context) error {
	m.started = true
	return m.err
}

func TestLeaderElectionManager_Start(t *testing.T) {
	// Save original runOrDie and restore after test
	originalRunOrDie := runOrDie
	defer func() { runOrDie = originalRunOrDie }()

	mockMgr := &mockManager{}
	leConfig := &leaderelection.LeaderElectionConfig{
		Callbacks: leaderelection.LeaderCallbacks{},
	}

	leMgr := &leaderElectionManager{
		Manager:  mockMgr,
		leConfig: leConfig,
	}

	// Mock runOrDie to immediately simulate acquiring leadership
	runOrDie = func(ctx context.Context, lec leaderelection.LeaderElectionConfig) {
		// Simulate gaining leadership
		if lec.Callbacks.OnStartedLeading != nil {
			lec.Callbacks.OnStartedLeading(ctx)
		}
	}

	ctx := context.TODO()
	if err := leMgr.Start(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !mockMgr.started {
		t.Error("expected inner manager to be started, but it was not")
	}
}

func TestSetUpMultiClusterLease(t *testing.T) {
	scheme := runtime.NewScheme()
	require.NoError(t, operatorv1beta1.AddToScheme(scheme))

	tests := []struct {
		name          string
		cc            *operatorv1beta1.ConfigConnector
		expectConfig  bool
		expectedLease string
	}{
		{
			name:         "ConfigConnector not found",
			cc:           nil,
			expectConfig: false,
		},
		{
			name: "MultiClusterLease disabled",
			cc: &operatorv1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "configconnector.core.cnrm.cloud.google.com",
				},
				Spec: operatorv1beta1.ConfigConnectorSpec{},
			},
			expectConfig: false,
		},
		{
			name: "MultiClusterLease enabled",
			cc: &operatorv1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "configconnector.core.cnrm.cloud.google.com",
				},
				Spec: operatorv1beta1.ConfigConnectorSpec{
					Experiments: &operatorv1beta1.CCExperiments{
						MultiClusterLease: &operatorv1beta1.MultiClusterLeaseSpec{
							LeaseName:                "test-lease",
							Namespace:                "test-ns",
							ClusterCandidateIdentity: "test-lock",
						},
					},
				},
			},
			expectConfig:  true,
			expectedLease: "test-lease",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			initObjs := []client.Object{}
			if tc.cc != nil {
				initObjs = append(initObjs, tc.cc)
			}

			c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(initObjs...).Build()

			var explicitConfig *operatorv1beta1.MultiClusterLeaseSpec
			if tc.cc != nil && tc.cc.Spec.Experiments != nil {
				explicitConfig = tc.cc.Spec.Experiments.MultiClusterLease
			}

			config, _, err := setUpMultiClusterLease(context.TODO(), nil, scheme, explicitConfig, false, c)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tc.expectConfig {
				if config == nil {
					t.Fatal("expected leader election config, got nil")
				}
				if config.Name != tc.expectedLease {
					t.Errorf("expected lease name %q, got %q", tc.expectedLease, config.Name)
				}
			} else {
				if config != nil {
					t.Errorf("expected no leader election config, got %v", config)
				}
			}
		})
	}
}
