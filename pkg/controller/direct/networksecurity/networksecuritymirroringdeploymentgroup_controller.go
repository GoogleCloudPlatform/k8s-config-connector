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

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	networksecurity "cloud.google.com/go/networksecurity/apiv1"
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

func (m *mirroringDeploymentGroupModel) client(ctx context.Context) (*networksecurity.MirroringClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := networksecurity.NewMirroringRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networksecurity Mirroring REST client: %w", err)
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
	id := &krm.NetworkSecurityMirroringDeploymentGroupIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &mirroringDeploymentGroupAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type mirroringDeploymentGroupAdapter struct {
	gcpClient *networksecurity.MirroringClient
	id        *krm.NetworkSecurityMirroringDeploymentGroupIdentity
	desired   *pb.MirroringDeploymentGroup
	actual    *pb.MirroringDeploymentGroup
}

var _ directbase.Adapter = &mirroringDeploymentGroupAdapter{}

func (a *mirroringDeploymentGroupAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding NetworkSecurityMirroringDeploymentGroup", "name", a.id)

	req := &pb.GetMirroringDeploymentGroupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetMirroringDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkSecurityMirroringDeploymentGroup %s: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *mirroringDeploymentGroupAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkSecurityMirroringDeploymentGroup", "name", a.id)

	req := &pb.CreateMirroringDeploymentGroupRequest{
		Parent:                     a.id.ParentString(),
		MirroringDeploymentGroupId: a.id.MirroringDeploymentGroup,
		MirroringDeploymentGroup:   a.desired,
	}

	op, err := a.gcpClient.CreateMirroringDeploymentGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkSecurityMirroringDeploymentGroup %s: %w", a.id, err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NetworkSecurityMirroringDeploymentGroup %s waiting for creation: %w", a.id, err)
	}

	actual, err := a.gcpClient.GetMirroringDeploymentGroup(ctx, &pb.GetMirroringDeploymentGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting NetworkSecurityMirroringDeploymentGroup %s after creation: %w", a.id, err)
	}

	log.V(2).Info("successfully created NetworkSecurityMirroringDeploymentGroup", "name", a.id)

	return a.updateStatus(ctx, createOp, actual)
}

func (a *mirroringDeploymentGroupAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkSecurityMirroringDeploymentGroup", "name", a.id)

	diffs, updateMask, err := compareMirroringDeploymentGroup(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.MirroringDeploymentGroup)
		desired.Name = a.id.String()

		req := &pb.UpdateMirroringDeploymentGroupRequest{
			MirroringDeploymentGroup: desired,
			UpdateMask:               updateMask,
		}

		op, err := a.gcpClient.UpdateMirroringDeploymentGroup(ctx, req)
		if err != nil {
			return fmt.Errorf("updating NetworkSecurityMirroringDeploymentGroup %s: %w", a.id, err)
		}

		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("NetworkSecurityMirroringDeploymentGroup %s waiting for update: %w", a.id, err)
		}
	}

	latest, err := a.gcpClient.GetMirroringDeploymentGroup(ctx, &pb.GetMirroringDeploymentGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting NetworkSecurityMirroringDeploymentGroup %s after update: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *mirroringDeploymentGroupAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MirroringDeploymentGroup) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityMirroringDeploymentGroupStatus{}
	status.ObservedState = NetworkSecurityMirroringDeploymentGroupObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *mirroringDeploymentGroupAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityMirroringDeploymentGroup{}
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityMirroringDeploymentGroupSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if spec != nil {
		obj.Spec = *spec
	}

	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.MirroringDeploymentGroup

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.MirroringDeploymentGroup)
	u.SetGroupVersionKind(krm.NetworkSecurityMirroringDeploymentGroupGVK)
	return u, nil
}

func (a *mirroringDeploymentGroupAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkSecurityMirroringDeploymentGroup", "name", a.id)

	req := &pb.DeleteMirroringDeploymentGroupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMirroringDeploymentGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkSecurityMirroringDeploymentGroup %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("NetworkSecurityMirroringDeploymentGroup %s waiting for deletion: %w", a.id, err)
	}

	return true, nil
}

// compareMirroringDeploymentGroup compares the actual and desired MirroringDeploymentGroup state, returning the diff and update field mask.
func compareMirroringDeploymentGroup(ctx context.Context, actual, desired *pb.MirroringDeploymentGroup) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkSecurityMirroringDeploymentGroupSpec_v1alpha1_FromProto, NetworkSecurityMirroringDeploymentGroupSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
