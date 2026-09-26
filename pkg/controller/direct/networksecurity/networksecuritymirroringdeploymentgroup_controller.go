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

package networksecurity

import (
	"context"
	"fmt"
	"reflect"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/networksecurity/apiv1"
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityMirroringDeploymentGroupGVK, NewMirroringDeploymentGroupModel)
}

func NewMirroringDeploymentGroupModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &mirroringDeploymentGroupModel{config: *config}, nil
}

var _ directbase.Model = &mirroringDeploymentGroupModel{}

type mirroringDeploymentGroupModel struct {
	config config.ControllerConfig
}

func (m *mirroringDeploymentGroupModel) client(ctx context.Context) (*gcp.MirroringClient, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewMirroringRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networksecurity mirroring client: %w", err)
	}

	return gcpClient, nil
}

func (m *mirroringDeploymentGroupModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityMirroringDeploymentGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityMirroringDeploymentGroupSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &mirroringDeploymentGroupAdapter{
		gcpClient: gcpClient,
		id:        id.(*krm.NetworkSecurityMirroringDeploymentGroupIdentity),
		desired:   desired,
	}, nil
}

func (m *mirroringDeploymentGroupModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type mirroringDeploymentGroupAdapter struct {
	gcpClient *gcp.MirroringClient
	id        *krm.NetworkSecurityMirroringDeploymentGroupIdentity
	desired   *pb.MirroringDeploymentGroup
	actual    *pb.MirroringDeploymentGroup
}

var _ directbase.Adapter = &mirroringDeploymentGroupAdapter{}

func (a *mirroringDeploymentGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting networksecurity mirroring deployment group", "name", a.id)

	req := &pb.GetMirroringDeploymentGroupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMirroringDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity mirroring deployment group %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *mirroringDeploymentGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating networksecurity mirroring deployment group", "name", a.id)

	parent := a.id.ParentString()
	req := &pb.CreateMirroringDeploymentGroupRequest{
		Parent:                     parent,
		MirroringDeploymentGroupId: a.id.MirroringDeploymentGroup,
		MirroringDeploymentGroup:   a.desired,
	}
	op, err := a.gcpClient.CreateMirroringDeploymentGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networksecurity mirroring deployment group %s: %w", a.id.String(), err)
	}

	actual, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("networksecurity mirroring deployment group %s waiting for creation: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created networksecurity mirroring deployment group", "name", a.id.String())

	return a.updateStatus(ctx, createOp, actual)
}

func (a *mirroringDeploymentGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating networksecurity mirroring deployment group", "name", a.id)

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	var paths []string

	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, a.desired.Labels)
		paths = append(paths, "labels")
	}

	if a.desired.Description != a.actual.Description {
		report.AddField("description", a.actual.Description, a.desired.Description)
		paths = append(paths, "description")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	a.desired.Name = a.id.String()
	updateMask := &fieldmaskpb.FieldMask{Paths: paths}

	req := &pb.UpdateMirroringDeploymentGroupRequest{
		UpdateMask:               updateMask,
		MirroringDeploymentGroup: a.desired,
	}

	op, err := a.gcpClient.UpdateMirroringDeploymentGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating networksecurity mirroring deployment group %s: %w", a.id.String(), err)
	}

	actual, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("networksecurity mirroring deployment group %s waiting for update: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated networksecurity mirroring deployment group", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, actual)
}

func (a *mirroringDeploymentGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *mirroringDeploymentGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting networksecurity mirroring deployment group", "name", a.id)

	req := &pb.DeleteMirroringDeploymentGroupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMirroringDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting networksecurity mirroring deployment group %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("networksecurity mirroring deployment group %s waiting for deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *mirroringDeploymentGroupAdapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.MirroringDeploymentGroup) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityMirroringDeploymentGroupStatus{}
	status.ObservedState = NetworkSecurityMirroringDeploymentGroupObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}
