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

// +tool:controller
// proto.service: google.cloud.dataproc.v1.ClusterController
// proto.message: google.cloud.dataproc.v1.Cluster
// crd.type: DataprocCluster
// crd.version: v1beta1

package dataproc

import (
	"context"
	"fmt"
	"sort"
	"strings"

	dataproc "cloud.google.com/go/dataproc/v2/apiv1"
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataprocClusterGVK, NewDataprocClusterModel)
}

func NewDataprocClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dataprocClusterModel{config: *config}, nil
}

var _ directbase.Model = &dataprocClusterModel{}

type dataprocClusterModel struct {
	config config.ControllerConfig
}

func (m *dataprocClusterModel) client(ctx context.Context) (*dataproc.ClusterControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	return dataproc.NewClusterControllerRESTClient(ctx, opts...)
}

func (m *dataprocClusterModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataprocCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.GetIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	clusterID, ok := id.(*krm.DataprocClusterIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", id)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &dataprocClusterAdapter{
		gcpClient: gcpClient,
		id:        clusterID,
		desired:   obj,
	}, nil
}

func (m *dataprocClusterModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

var _ directbase.Adapter = &dataprocClusterAdapter{}

type dataprocClusterAdapter struct {
	gcpClient *dataproc.ClusterControllerClient
	id        *krm.DataprocClusterIdentity
	desired   *krm.DataprocCluster
	actual    *pb.Cluster
}

func (a *dataprocClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataproc cluster", "name", a.id)

	req := &pb.GetClusterRequest{
		ProjectId:   a.id.Project,
		Region:      a.id.Region,
		ClusterName: a.id.Cluster,
	}
	actual, err := a.gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataproc cluster %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *dataprocClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataproc cluster", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	cluster := DataprocClusterSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	cluster.ClusterName = a.id.Cluster
	cluster.ProjectId = a.id.Project

	cluster.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		cluster.Labels[k] = v
	}
	cluster.Labels["managed-by-cnrm"] = "true"

	req := &pb.CreateClusterRequest{
		ProjectId: a.id.Project,
		Region:    a.id.Region,
		Cluster:   cluster,
	}
	op, err := a.gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataproc cluster %s: %w", a.id.String(), err)
	}

	log.V(2).Info("waiting for dataproc cluster creation", "name", a.id)
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for dataproc cluster %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created dataproc cluster in gcp", "name", a.id)

	latest, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{
		ProjectId:   a.id.Project,
		Region:      a.id.Region,
		ClusterName: a.id.Cluster,
	})
	if err != nil {
		return fmt.Errorf("getting latest dataproc cluster after creation: %w", err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *dataprocClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataproc cluster", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	cluster := DataprocClusterSpec_v1beta1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	cluster.ClusterName = a.id.Cluster
	cluster.ProjectId = a.id.Project

	cluster.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		cluster.Labels[k] = v
	}
	cluster.Labels["managed-by-cnrm"] = "true"

	// Populate defaults to avoid false drift
	if cluster.Config == nil {
		cluster.Config = a.actual.Config
	} else {
		if cluster.Config.ConfigBucket == "" {
			cluster.Config.ConfigBucket = a.actual.Config.ConfigBucket
		}
		if cluster.Config.TempBucket == "" {
			cluster.Config.TempBucket = a.actual.Config.TempBucket
		}
		if cluster.Config.GceClusterConfig == nil {
			cluster.Config.GceClusterConfig = a.actual.Config.GceClusterConfig
		} else {
			if cluster.Config.GceClusterConfig.ServiceAccountScopes == nil {
				cluster.Config.GceClusterConfig.ServiceAccountScopes = a.actual.Config.GceClusterConfig.ServiceAccountScopes
			}
			if cluster.Config.GceClusterConfig.NetworkUri == "" {
				cluster.Config.GceClusterConfig.NetworkUri = a.actual.Config.GceClusterConfig.NetworkUri
			}
			if cluster.Config.GceClusterConfig.ShieldedInstanceConfig == nil {
				cluster.Config.GceClusterConfig.ShieldedInstanceConfig = a.actual.Config.GceClusterConfig.ShieldedInstanceConfig
			}
			if cluster.Config.GceClusterConfig.InternalIpOnly == nil {
				cluster.Config.GceClusterConfig.InternalIpOnly = a.actual.Config.GceClusterConfig.InternalIpOnly
			}
			if cluster.Config.GceClusterConfig.ZoneUri == "" {
				cluster.Config.GceClusterConfig.ZoneUri = a.actual.Config.GceClusterConfig.ZoneUri
			} else {
				// Convert Zone to ZoneUri
				cluster.Config.GceClusterConfig.ZoneUri = fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/%s", a.id.Project, direct.ValueOf(a.desired.Spec.Config.GceClusterConfig.Zone))
			}
		}
		if cluster.Config.InitializationActions == nil {
			cluster.Config.InitializationActions = a.actual.Config.InitializationActions
		}
		if cluster.Config.SoftwareConfig == nil {
			cluster.Config.SoftwareConfig = a.actual.Config.SoftwareConfig
		} else {
			if cluster.Config.SoftwareConfig.OptionalComponents == nil {
				cluster.Config.SoftwareConfig.OptionalComponents = []pb.Component{}
			}
			if cluster.Config.SoftwareConfig.Properties == nil {
				cluster.Config.SoftwareConfig.Properties = a.actual.Config.SoftwareConfig.Properties
			}
		}
		// For MasterConfig, it is immutable
		cluster.Config.MasterConfig = alignInstanceGroupConfig(cluster.Config.MasterConfig, a.actual.Config.MasterConfig, a.actual.Config.GceClusterConfig.ZoneUri)

		// For WorkerConfig: only NumInstances is mutable, so copy other fields (like MachineTypeUri, DiskConfig, etc.)
		cluster.Config.WorkerConfig = alignInstanceGroupConfig(cluster.Config.WorkerConfig, a.actual.Config.WorkerConfig, a.actual.Config.GceClusterConfig.ZoneUri)

		// For SecondaryWorkerConfig: only NumInstances is mutable
		cluster.Config.SecondaryWorkerConfig = alignInstanceGroupConfig(cluster.Config.SecondaryWorkerConfig, a.actual.Config.SecondaryWorkerConfig, a.actual.Config.GceClusterConfig.ZoneUri)
		if cluster.Config.SecondaryWorkerConfig != nil {
			if cluster.Config.SecondaryWorkerConfig.IsPreemptible == false {
				cluster.Config.SecondaryWorkerConfig.IsPreemptible = true
			}
			if cluster.Config.SecondaryWorkerConfig.ManagedGroupConfig == nil && a.actual.Config.SecondaryWorkerConfig != nil {
				cluster.Config.SecondaryWorkerConfig.ManagedGroupConfig = a.actual.Config.SecondaryWorkerConfig.ManagedGroupConfig
			}
		}
		// For SecurityConfig: only UserServiceAccountMapping is mutable
		if cluster.Config.SecurityConfig == nil {
			cluster.Config.SecurityConfig = a.actual.Config.SecurityConfig
		} else if cluster.Config.SecurityConfig.IdentityConfig == nil {
			cluster.Config.SecurityConfig.IdentityConfig = a.actual.Config.SecurityConfig.IdentityConfig
		}
	}

	// Align output-only fields of the top-level pb.Cluster
	cluster.Status = a.actual.Status
	cluster.StatusHistory = a.actual.StatusHistory
	cluster.ClusterUuid = a.actual.ClusterUuid
	cluster.Metrics = a.actual.Metrics

	// todo: replace CompareProtoMessageStructuredDiff with CompareBrownfieldSpec
	// Compute the diff using common.CompareProtoMessageStructuredDiff
	paths, diffs, err := common.CompareProtoMessageStructuredDiff(cluster, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing cluster spec: %w", err)
	}

	// Filter out any paths that are not mutable on update.
	// Only allow the following paths in the updateMask:
	allowedPaths := map[string]bool{
		"labels":                             true,
		"config.worker_config.num_instances": true,
		"config.secondary_worker_config.num_instances":                        true,
		"config.autoscaling_config.policy_uri":                                true,
		"config.lifecycle_config.auto_delete_ttl":                             true,
		"config.lifecycle_config.auto_delete_time":                            true,
		"config.lifecycle_config.idle_delete_ttl":                             true,
		"config.security_config.identity_config.user_service_account_mapping": true,
	}

	for path := range paths {
		if !allowedPaths[path] {
			return fmt.Errorf("field %s is immutable", path)
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no diff detected, skipping update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	pathsList := paths.UnsortedList()
	sort.Strings(pathsList)
	updateMask := &fieldmaskpb.FieldMask{
		Paths: pathsList,
	}

	req := &pb.UpdateClusterRequest{
		ProjectId:   a.id.Project,
		Region:      a.id.Region,
		ClusterName: a.id.Cluster,
		Cluster:     cluster,
		UpdateMask:  updateMask,
	}
	op, err := a.gcpClient.UpdateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating dataproc cluster %s: %w", a.id.String(), err)
	}

	log.V(2).Info("waiting for dataproc cluster update", "name", a.id)
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for dataproc cluster %s update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated dataproc cluster", "name", a.id)

	latest, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{
		ProjectId:   a.id.Project,
		Region:      a.id.Region,
		ClusterName: a.id.Cluster,
	})
	if err != nil {
		return fmt.Errorf("getting latest dataproc cluster after update: %w", err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *dataprocClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataproc cluster", "name", a.id)

	req := &pb.DeleteClusterRequest{
		ProjectId:   a.id.Project,
		Region:      a.id.Region,
		ClusterName: a.id.Cluster,
	}
	op, err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting dataproc cluster %s: %w", a.id.String(), err)
	}

	log.V(2).Info("waiting for dataproc cluster deletion", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for dataproc cluster %s deletion: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted dataproc cluster", "name", a.id)

	return true, nil
}

// Export implements the Adapter interface.
func (a *dataprocClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DataprocCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataprocClusterSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &parent.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Region
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Cluster)
	u.SetGroupVersionKind(krm.DataprocClusterGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

func (a *dataprocClusterAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Cluster) error {
	mapCtx := &direct.MapContext{}
	status := DataprocClusterStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(krm.DataprocClusterIdentityFormatRelative.ToString(*a.id))
	return op.UpdateStatus(ctx, status, nil)
}

func alignInstanceGroupConfig(desired, actual *pb.InstanceGroupConfig, zoneUri string) *pb.InstanceGroupConfig {
	if desired == nil {
		return actual
	}
	if desired.ImageUri == "" {
		desired.ImageUri = actual.ImageUri
	}
	if desired.MachineTypeUri == "" {
		desired.MachineTypeUri = actual.MachineTypeUri
	} else if !strings.HasPrefix(desired.MachineTypeUri, "https://") {
		desired.MachineTypeUri = fmt.Sprintf("%s/machineTypes/%s", zoneUri, desired.MachineTypeUri)
	}
	if desired.DiskConfig == nil {
		desired.DiskConfig = actual.DiskConfig
	} else {
		if desired.DiskConfig.BootDiskSizeGb == 0 {
			desired.DiskConfig.BootDiskSizeGb = actual.DiskConfig.BootDiskSizeGb
		}
		if desired.DiskConfig.BootDiskType == "" {
			desired.DiskConfig.BootDiskType = actual.DiskConfig.BootDiskType
		}
	}
	if desired.Preemptibility == pb.InstanceGroupConfig_PREEMPTIBILITY_UNSPECIFIED {
		desired.Preemptibility = actual.Preemptibility
	}
	if desired.MinCpuPlatform == "" {
		desired.MinCpuPlatform = actual.MinCpuPlatform
	}
	if desired.Accelerators == nil {
		desired.Accelerators = actual.Accelerators
	}
	if desired.InstanceNames == nil {
		desired.InstanceNames = actual.InstanceNames
	}
	if desired.StartupConfig == nil {
		desired.StartupConfig = actual.StartupConfig
	}
	return desired
}
