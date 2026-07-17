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

	// Since GCP might have added default values or read-only/immutable fields in a.actual,
	// we copy/preserve them in our desired 'cluster' before comparison so that they don't produce false diffs.
	if cluster.Config == nil {
		cluster.Config = &pb.ClusterConfig{}
	}
	if a.actual.Config == nil {
		a.actual.Config = &pb.ClusterConfig{}
	}

	// Copy immutable/server-assigned top-level config fields from a.actual to cluster to avoid false diffs
	cluster.Config.ConfigBucket = a.actual.Config.ConfigBucket
	cluster.Config.TempBucket = a.actual.Config.TempBucket
	cluster.Config.EncryptionConfig = a.actual.Config.EncryptionConfig
	cluster.Config.MetastoreConfig = a.actual.Config.MetastoreConfig
	cluster.Config.DataprocMetricConfig = a.actual.Config.DataprocMetricConfig
	cluster.Config.LifecycleConfig = a.actual.Config.LifecycleConfig
	cluster.Config.AutoscalingConfig = a.actual.Config.AutoscalingConfig

	// Only copy EndpointConfig if it is specified in KRM spec. Otherwise, leave it nil.
	if a.desired.Spec.Config != nil && a.desired.Spec.Config.EndpointConfig != nil {
		cluster.Config.EndpointConfig = a.actual.Config.EndpointConfig
	} else {
		cluster.Config.EndpointConfig = nil
	}

	cluster.Config.GceClusterConfig = a.actual.Config.GceClusterConfig
	// If serviceAccountScopes was not specified in the KRM spec, set it to nil to omit from payload.
	if cluster.Config.GceClusterConfig != nil {
		if a.desired.Spec.Config == nil || a.desired.Spec.Config.GceClusterConfig == nil || len(a.desired.Spec.Config.GceClusterConfig.ServiceAccountScopes) == 0 {
			cluster.Config.GceClusterConfig.ServiceAccountScopes = nil
		}
	}

	cluster.Config.InitializationActions = a.actual.Config.InitializationActions
	if cluster.Config.InitializationActions == nil {
		cluster.Config.InitializationActions = []*pb.NodeInitializationAction{}
	}

	cluster.Config.SoftwareConfig = a.actual.Config.SoftwareConfig
	if cluster.Config.SoftwareConfig != nil {
		if cluster.Config.SoftwareConfig.OptionalComponents == nil {
			cluster.Config.SoftwareConfig.OptionalComponents = []pb.Component{}
		}
	}

	// For MasterConfig, it is also immutable:
	cluster.Config.MasterConfig = a.actual.Config.MasterConfig
	if cluster.Config.MasterConfig != nil {
		if cluster.Config.MasterConfig.Accelerators == nil {
			cluster.Config.MasterConfig.Accelerators = []*pb.AcceleratorConfig{}
		}
	}

	// For WorkerConfig: only NumInstances is mutable, so copy other fields (like MachineTypeUri, DiskConfig, etc.)
	if cluster.Config.WorkerConfig == nil {
		cluster.Config.WorkerConfig = a.actual.Config.WorkerConfig
	} else if a.actual.Config.WorkerConfig != nil {
		cluster.Config.WorkerConfig.ImageUri = a.actual.Config.WorkerConfig.ImageUri
		cluster.Config.WorkerConfig.MachineTypeUri = a.actual.Config.WorkerConfig.MachineTypeUri
		cluster.Config.WorkerConfig.DiskConfig = a.actual.Config.WorkerConfig.DiskConfig
		cluster.Config.WorkerConfig.Preemptibility = a.actual.Config.WorkerConfig.Preemptibility
		cluster.Config.WorkerConfig.MinCpuPlatform = a.actual.Config.WorkerConfig.MinCpuPlatform
		cluster.Config.WorkerConfig.Accelerators = a.actual.Config.WorkerConfig.Accelerators
		cluster.Config.WorkerConfig.InstanceNames = a.actual.Config.WorkerConfig.InstanceNames
	}
	if cluster.Config.WorkerConfig != nil {
		if cluster.Config.WorkerConfig.Accelerators == nil {
			cluster.Config.WorkerConfig.Accelerators = []*pb.AcceleratorConfig{}
		}
	}

	// For SecondaryWorkerConfig: only NumInstances is mutable
	if cluster.Config.SecondaryWorkerConfig == nil {
		cluster.Config.SecondaryWorkerConfig = a.actual.Config.SecondaryWorkerConfig
	} else if a.actual.Config.SecondaryWorkerConfig != nil {
		cluster.Config.SecondaryWorkerConfig.ImageUri = a.actual.Config.SecondaryWorkerConfig.ImageUri
		cluster.Config.SecondaryWorkerConfig.MachineTypeUri = a.actual.Config.SecondaryWorkerConfig.MachineTypeUri
		cluster.Config.SecondaryWorkerConfig.DiskConfig = a.actual.Config.SecondaryWorkerConfig.DiskConfig
		cluster.Config.SecondaryWorkerConfig.Preemptibility = a.actual.Config.SecondaryWorkerConfig.Preemptibility
		cluster.Config.SecondaryWorkerConfig.MinCpuPlatform = a.actual.Config.SecondaryWorkerConfig.MinCpuPlatform
		cluster.Config.SecondaryWorkerConfig.Accelerators = a.actual.Config.SecondaryWorkerConfig.Accelerators
		cluster.Config.SecondaryWorkerConfig.InstanceNames = a.actual.Config.SecondaryWorkerConfig.InstanceNames
	}
	if cluster.Config.SecondaryWorkerConfig != nil {
		if cluster.Config.SecondaryWorkerConfig.Accelerators == nil {
			cluster.Config.SecondaryWorkerConfig.Accelerators = []*pb.AcceleratorConfig{}
		}
	}

	// For SecurityConfig: only UserServiceAccountMapping is mutable
	if cluster.Config.SecurityConfig == nil {
		cluster.Config.SecurityConfig = a.actual.Config.SecurityConfig
	} else if a.actual.Config.SecurityConfig != nil {
		if cluster.Config.SecurityConfig.IdentityConfig == nil {
			cluster.Config.SecurityConfig.IdentityConfig = a.actual.Config.SecurityConfig.IdentityConfig
		}
	}

	// Align output-only fields of the top-level pb.Cluster
	cluster.Status = a.actual.Status
	cluster.StatusHistory = a.actual.StatusHistory
	cluster.ClusterUuid = a.actual.ClusterUuid
	cluster.Metrics = a.actual.Metrics

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
			paths.Delete(path)
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
