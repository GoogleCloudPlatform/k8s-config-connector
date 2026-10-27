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

package ha

import (
	"context"
	"fmt"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ConfigConnectorResourceName = "configconnector.core.cnrm.cloud.google.com"
)

// ClusterHAState holds the observed HA-related resources on a single cluster.
type ClusterHAState struct {
	ContextName       string
	ConfigConnector   *operatorv1beta1.ConfigConnector
	MultiClusterLease *mclv1alpha1.MultiClusterLease
	KRMSyncers        []syncerv1alpha1.KRMSyncer
}

// FetchAllClusterHAStates fetches the HA state across all configured clusters.
func FetchAllClusterHAStates(ctx context.Context, handles []*ClusterHandle) ([]*ClusterHAState, error) {
	states := make([]*ClusterHAState, 0, len(handles))
	for _, h := range handles {
		state, err := FetchClusterHAState(ctx, h)
		if err != nil {
			return nil, err
		}
		states = append(states, state)
	}
	return states, nil
}

// FetchClusterHAState reads the ConfigConnector, MultiClusterLease, and KRMSyncer objects from a cluster.
func FetchClusterHAState(ctx context.Context, handle *ClusterHandle) (*ClusterHAState, error) {
	if handle == nil || handle.Client == nil {
		return nil, fmt.Errorf("cluster handle or client is nil")
	}

	state := &ClusterHAState{
		ContextName: handle.ContextName,
	}

	cc, err := fetchConfigConnector(ctx, handle)
	if err != nil {
		return nil, err
	}
	state.ConfigConnector = cc

	mclSpec := multiClusterLeaseSpec(cc)
	if mclSpec == nil || mclSpec.LeaseName == "" || mclSpec.Namespace == "" {
		return state, nil
	}

	lease, err := fetchMultiClusterLease(ctx, handle, mclSpec.Namespace, mclSpec.LeaseName)
	if err != nil {
		return nil, err
	}
	state.MultiClusterLease = lease

	syncers, err := listKRMSyncers(ctx, handle, mclSpec.Namespace)
	if err != nil {
		return nil, err
	}
	state.KRMSyncers = syncers

	return state, nil
}

func fetchConfigConnector(ctx context.Context, handle *ClusterHandle) (*operatorv1beta1.ConfigConnector, error) {
	cc := &operatorv1beta1.ConfigConnector{}
	key := types.NamespacedName{Name: ConfigConnectorResourceName}
	err := handle.Client.Get(ctx, key, cc)
	if err == nil {
		return cc, nil
	}
	if isResourceOrKindAbsent(err) {
		return nil, nil
	}
	return nil, fmt.Errorf("fetching ConfigConnector on cluster %q: %w", handle.ContextName, err)
}

func multiClusterLeaseSpec(cc *operatorv1beta1.ConfigConnector) *operatorv1beta1.MultiClusterLeaseSpec {
	if cc == nil || cc.Spec.Experiments == nil {
		return nil
	}
	return cc.Spec.Experiments.MultiClusterLease
}

func fetchMultiClusterLease(ctx context.Context, handle *ClusterHandle, namespace, name string) (*mclv1alpha1.MultiClusterLease, error) {
	lease := &mclv1alpha1.MultiClusterLease{}
	key := types.NamespacedName{Namespace: namespace, Name: name}
	err := handle.Client.Get(ctx, key, lease)
	if err == nil {
		return lease, nil
	}
	if isResourceOrKindAbsent(err) {
		return nil, nil
	}
	return nil, fmt.Errorf("fetching MultiClusterLease %s on cluster %q: %w", key, handle.ContextName, err)
}

func listKRMSyncers(ctx context.Context, handle *ClusterHandle, namespace string) ([]syncerv1alpha1.KRMSyncer, error) {
	syncerList := &syncerv1alpha1.KRMSyncerList{}
	err := handle.Client.List(ctx, syncerList, client.InNamespace(namespace))
	if err == nil {
		return syncerList.Items, nil
	}
	if isResourceOrKindAbsent(err) {
		return nil, nil
	}
	return nil, fmt.Errorf("listing KRMSyncers in namespace %q on cluster %q: %w", namespace, handle.ContextName, err)
}

func isResourceOrKindAbsent(err error) bool {
	return apierrors.IsNotFound(err) || meta.IsNoMatchError(err) || runtime.IsNotRegisteredError(err)
}

// ValidateMultiClusterLeaseConfigs checks that all clusters have consistent MultiClusterLease
// settings (identical leaseName and namespace, and unique clusterCandidateIdentity per cluster).
func ValidateMultiClusterLeaseConfigs(states []*ClusterHAState) error {
	if len(states) < 2 {
		return fmt.Errorf("at least 2 clusters are required for multi-cluster HA validation, got %d", len(states))
	}

	firstSpec, err := extractAndValidateLeaseSpec(states[0])
	if err != nil {
		return err
	}
	firstCluster := states[0].ContextName

	identities := make(map[string]string, len(states))
	for _, s := range states {
		mcl, err := extractAndValidateLeaseSpec(s)
		if err != nil {
			return err
		}
		if err := validateSameLeaseTarget(firstCluster, firstSpec, s.ContextName, mcl); err != nil {
			return err
		}
		if prevCluster, exists := identities[mcl.ClusterCandidateIdentity]; exists {
			return fmt.Errorf("duplicate clusterCandidateIdentity %q shared by clusters %q and %q",
				mcl.ClusterCandidateIdentity, prevCluster, s.ContextName)
		}
		identities[mcl.ClusterCandidateIdentity] = s.ContextName
	}

	return nil
}

func extractAndValidateLeaseSpec(s *ClusterHAState) (*operatorv1beta1.MultiClusterLeaseSpec, error) {
	if s == nil {
		return nil, fmt.Errorf("cluster HA state is nil")
	}
	if s.ConfigConnector == nil {
		return nil, fmt.Errorf("cluster %q is missing the %q ConfigConnector object", s.ContextName, ConfigConnectorResourceName)
	}
	mcl := multiClusterLeaseSpec(s.ConfigConnector)
	if mcl == nil {
		return nil, fmt.Errorf("cluster %q does not have spec.experiments.multiClusterLease configured on ConfigConnector", s.ContextName)
	}
	if mcl.LeaseName == "" {
		return nil, fmt.Errorf("cluster %q has an empty spec.experiments.multiClusterLease.leaseName", s.ContextName)
	}
	if mcl.Namespace == "" {
		return nil, fmt.Errorf("cluster %q has an empty spec.experiments.multiClusterLease.namespace", s.ContextName)
	}
	if mcl.ClusterCandidateIdentity == "" {
		return nil, fmt.Errorf("cluster %q has an empty spec.experiments.multiClusterLease.clusterCandidateIdentity", s.ContextName)
	}
	return mcl, nil
}

func validateSameLeaseTarget(firstCluster string, expected *operatorv1beta1.MultiClusterLeaseSpec, cluster string, actual *operatorv1beta1.MultiClusterLeaseSpec) error {
	if actual.LeaseName != expected.LeaseName {
		return fmt.Errorf("mismatched leaseName across clusters: cluster %q has %q, but cluster %q has %q",
			firstCluster, expected.LeaseName, cluster, actual.LeaseName)
	}
	if actual.Namespace != expected.Namespace {
		return fmt.Errorf("mismatched lease namespace across clusters: cluster %q has %q, but cluster %q has %q",
			firstCluster, expected.Namespace, cluster, actual.Namespace)
	}
	return nil
}
