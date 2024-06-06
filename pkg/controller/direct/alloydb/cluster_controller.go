// Copyright 2024 Google LLC
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
	"fmt"
	"reflect"
	"strings"
	"time"

	api "google.golang.org/api/alloydb/v1beta"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

// AddClusterController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddClusterController(mgr manager.Manager, config *controller.Config, opts directbase.Deps) error {
	gvk := krm.AlloyDBClusterGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &clusterModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m, opts)
}

type clusterModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &clusterModel{}

type clusterAdapter struct {
	projectID  string
	location   string
	resourceID string

	desired *api.Cluster
	actual  *api.Cluster

	client *api.Service
}

var _ directbase.Adapter = &clusterAdapter{}

// AdapterForObject implements the Model interface.
func (m *clusterModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	klog.FromContext(ctx).V(0).Info("creating adapter", "u", u)
	client, err := m.newAlloyDBAdminClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.AlloyDBCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	mapCtx := &MapContext{
		//	kube: kube,
	}
	desired := ClusterSpecToApi(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &clusterAdapter{
		resourceID: resourceID,
		projectID:  projectID,
		location:   location,
		desired:    desired,
		client:     client,
	}, nil
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &clusterAdapter{}

// Find implements the Adapter interface.
func (a *clusterAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}
	cluster, err := a.client.Projects.Locations.Clusters.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = cluster

	return true, nil
}

// Delete implements the Adapter interface.
func (a *clusterAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	op, err := a.client.Projects.Locations.Clusters.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting cluster %s: %w", a.fullyQualifiedName(), err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("cluster deletion failed: %w", err)
	}
	return true, nil
}

func (a *clusterAdapter) waitForOp(ctx context.Context, op *api.Operation) error {
	for {
		current, err := a.client.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			} else {
				return nil
			}
		}
		time.Sleep(2 * time.Second)
	}
}

// Create implements the Adapter interface.
func (a *clusterAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	parent := "projects/" + a.projectID + "/locations/" + a.location
	cluster := a.desired

	log.V(0).Info("creating cluster", "cluster", cluster)
	var op *api.Operation
	var err error
	// Default ClusterType to be PRIMARY
	if cluster.ClusterType == "SECONDARY" {
		op, err = a.client.Projects.Locations.Clusters.Createsecondary(parent, cluster).ClusterId(a.resourceID).Context(ctx).Do()
	} else {
		op, err = a.client.Projects.Locations.Clusters.Create(parent, cluster).ClusterId(a.resourceID).Context(ctx).Do()
	}
	if err != nil {
		return fmt.Errorf("creating cluster: %w", err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for cluster create %s: %w", a.fullyQualifiedName(), err)
	}

	created, err := a.client.Projects.Locations.Clusters.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created cluster: %w", err)
	}

	log.V(0).Info("created cluster", "cluster", created)

	resourceID := created.Name
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	mapCtx := &MapContext{
		// kube: kube,
	}
	observedState := ClusterStatusFromApi(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, observedState)
}

// Update implements the Adapter interface.
func (a *clusterAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)
	var latest *api.Cluster

	updateMask := &fieldmaskpb.FieldMask{}
	if a.desired.DisplayName != a.actual.DisplayName {
		log.V(0).Info("change detected: displayName")
		updateMask.Paths = append(updateMask.Paths, "displayName")
	}
	if a.desired.ClusterType != a.actual.ClusterType {
		log.V(0).Info("change detected: clusterType")
		updateMask.Paths = append(updateMask.Paths, "clusterType")
	}
	if !reflect.DeepEqual(a.desired.EncryptionConfig, a.actual.EncryptionConfig) {
		log.V(0).Info("change detected: encryption_config")
		updateMask.Paths = append(updateMask.Paths, "encryption_config")
	}
	if a.desired.Network != a.actual.Network {
		log.V(0).Info("change detected: network")
		updateMask.Paths = append(updateMask.Paths, "network")
	}
	if !reflect.DeepEqual(a.desired.NetworkConfig, a.actual.NetworkConfig) {
		log.V(0).Info("change detected: networkConfig")
		updateMask.Paths = append(updateMask.Paths, "networkConfig")
	}
	if !reflect.DeepEqual(a.desired.InitialUser, a.actual.InitialUser) {
		log.V(0).Info("change detected: initialUser")
		updateMask.Paths = append(updateMask.Paths, "initialUser")
	}
	if !reflect.DeepEqual(a.desired.ContinuousBackupConfig, a.actual.ContinuousBackupConfig) {
		log.V(0).Info("change detected: continuousBackupConfig")
		updateMask.Paths = append(updateMask.Paths, "continuousBackupConfig")
	}
	if !reflect.DeepEqual(a.desired.AutomatedBackupPolicy, a.actual.AutomatedBackupPolicy) {
		log.V(0).Info("change detected: automatedBackupPolicy")
		updateMask.Paths = append(updateMask.Paths, "automatedBackupPolicy")
	}
	if !reflect.DeepEqual(a.desired.SecondaryConfig, a.actual.SecondaryConfig) {
		log.V(0).Info("change detected: secondaryConfig")
		updateMask.Paths = append(updateMask.Paths, "secondaryConfig")
	}
	if len(updateMask.Paths) != 0 {
		cluster := a.desired
		var clusterName string
		op, err := a.client.Projects.Locations.Clusters.Patch(clusterName, cluster).UpdateMask(strings.Join(updateMask.Paths, ",")).Context(ctx).Do()
		if err != nil {
			return err
		}

		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("waiting for cluster update %s: %w", a.fullyQualifiedName(), err)
		}

		updated, err := a.client.Projects.Locations.Clusters.Get(a.fullyQualifiedName()).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated cluster: %w", err)
		}
		log.V(0).Info("updated cluster", "cluster", updated)
		latest = updated
	} else {
		latest = a.actual
	}

	mapCtx := &MapContext{
		// kube: kube,
	}
	observedState := ClusterStatusFromApi(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, observedState)
}

func (a *clusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (a *clusterAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/clusters/%s", a.projectID, a.location, a.resourceID)
}
