// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:controller
// proto.service: google.cloud.dataproc.v1.NodeGroupController
// proto.message: google.cloud.dataproc.v1.NodeGroup
// crd.type: DataprocNodeGroup
// crd.version: v1alpha1

package dataproc

import (
	"context"
	"fmt"

	dataproc "cloud.google.com/go/dataproc/v2/apiv1"
	dataprocpb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DataprocNodeGroupGVK, NewNodeGroupModel)
}

func NewNodeGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &nodeGroupModel{config: *config}, nil
}

var _ directbase.Model = &nodeGroupModel{}

type nodeGroupModel struct {
	config config.ControllerConfig
}

func (m *nodeGroupModel) Client(ctx context.Context) (*dataproc.NodeGroupControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := dataproc.NewNodeGroupControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc nodegroup client: %w", err)
	}

	return gcpClient, err
}

func (m *nodeGroupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataprocNodeGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNodeGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.Client(ctx)
	if err != nil {
		return nil, err
	}

	return &nodeGroupAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
	}, nil
}

func (m *nodeGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type nodeGroupAdapter struct {
	gcpClient *dataproc.NodeGroupControllerClient
	id        *krm.NodeGroupIdentity
	desired   *krm.DataprocNodeGroup
	actual    *dataprocpb.NodeGroup
}

var _ directbase.Adapter = &nodeGroupAdapter{}

func (a *nodeGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting dataproc nodegroup", "name", a.id)

	req := &dataprocpb.GetNodeGroupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetNodeGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting dataproc nodegroup %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *nodeGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating dataproc nodegroup", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataprocNodeGroupSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataprocpb.CreateNodeGroupRequest{
		Parent:      fmt.Sprintf("projects/%s/regions/%s/clusters/%s", a.id.Project, a.id.Region, a.id.Cluster),
		NodeGroup:   resource,
		NodeGroupId: a.id.NodeGroup,
	}
	op, err := a.gcpClient.CreateNodeGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating dataproc nodegroup %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("dataproc nodegroup %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created dataproc nodegroup in gcp", "name", a.id)

	status := &krm.DataprocNodeGroupStatus{}
	status.ObservedState = DataprocNodeGroupObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *nodeGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating dataproc nodegroup", "name", a.id)

	desiredSize := int32(0)
	if a.desired.Spec.NodeGroupConfig != nil && a.desired.Spec.NodeGroupConfig.NumInstances != nil {
		desiredSize = *a.desired.Spec.NodeGroupConfig.NumInstances
	}

	actualSize := int32(0)
	if a.actual.NodeGroupConfig != nil {
		actualSize = a.actual.NodeGroupConfig.NumInstances
	}

	if desiredSize != actualSize {
		log.V(2).Info("resizing dataproc nodegroup", "name", a.id, "from", actualSize, "to", desiredSize)
		req := &dataprocpb.ResizeNodeGroupRequest{
			Name: a.id.String(),
			Size: desiredSize,
		}
		op, err := a.gcpClient.ResizeNodeGroup(ctx, req)
		if err != nil {
			return fmt.Errorf("resizing dataproc nodegroup %s: %w", a.id.String(), err)
		}
		resized, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("dataproc nodegroup %s waiting resize: %w", a.id.String(), err)
		}
		a.actual = resized
	}

	desiredpb := DataprocNodeGroupSpec_v1alpha1_ToProto(&direct.MapContext{}, &a.desired.Spec)

	desiredCopy := proto.Clone(desiredpb).(*dataprocpb.NodeGroup)
	actualCopy := proto.Clone(a.actual).(*dataprocpb.NodeGroup)

	if desiredCopy.NodeGroupConfig != nil && actualCopy.NodeGroupConfig != nil {
		desiredCopy.NodeGroupConfig.NumInstances = actualCopy.NodeGroupConfig.NumInstances
	}

	paths, err := common.CompareProtoMessage(desiredCopy, actualCopy, common.BasicDiff)
	if err != nil {
		return err
	}
	delete(paths, ".name")

	if len(paths) != 0 {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)
		log.V(2).Info("This resource only supports size (numInstances) update", "name", a.id.String())
	}

	mapCtx := &direct.MapContext{}
	status := &krm.DataprocNodeGroupStatus{}
	status.ObservedState = DataprocNodeGroupObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *nodeGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting dataproc nodegroup (no-op on GCP as Delete is not supported)", "name", a.id)
	return true, nil
}

func (a *nodeGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	desired := &krm.DataprocNodeGroup{}
	mapCtx := &direct.MapContext{}
	desired.Spec = *DataprocNodeGroupSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, fmt.Errorf("converting DataprocNodeGroup to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetGroupVersionKind(krm.DataprocNodeGroupGVK)
	u.SetName(a.id.NodeGroup)
	return u, nil
}
