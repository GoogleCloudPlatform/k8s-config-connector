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
// proto.service: google.cloud.compute.v1.ExternalVpnGateways
// proto.message: google.cloud.compute.v1.ExternalVpnGateway
// crd.type: ComputeExternalVPNGateway
// crd.version: v1beta1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeExternalVPNGatewayGVK, NewComputeExternalVPNGatewayModel)
}

func NewComputeExternalVPNGatewayModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &ComputeExternalVPNGatewayModel{config: config}, nil
}

var _ directbase.Model = &ComputeExternalVPNGatewayModel{}

type ComputeExternalVPNGatewayModel struct {
	config *config.ControllerConfig
}

func (m *ComputeExternalVPNGatewayModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeExternalVPNGateway{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	externalVpnGatewaysClient, err := gcpClient.newExternalVpnGatewaysClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ComputeExternalVPNGatewaySpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	return &ComputeExternalVPNGatewayAdapter{
		gcpClient: externalVpnGatewaysClient,
		id:        id.(*krm.ComputeExternalVPNGatewayIdentity),
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *ComputeExternalVPNGatewayModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.ComputeExternalVPNGatewayIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	externalVpnGatewaysClient, err := gcpClient.newExternalVpnGatewaysClient(ctx)
	if err != nil {
		return nil, err
	}

	return &ComputeExternalVPNGatewayAdapter{
		id:        id,
		gcpClient: externalVpnGatewaysClient,
	}, nil
}

type ComputeExternalVPNGatewayAdapter struct {
	gcpClient *compute.ExternalVpnGatewaysClient
	id        *krm.ComputeExternalVPNGatewayIdentity
	desired   *pb.ExternalVpnGateway
	actual    *pb.ExternalVpnGateway
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeExternalVPNGatewayAdapter{}

func (a *ComputeExternalVPNGatewayAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeExternalVPNGateway", "name", a.id)

	req := &pb.GetExternalVpnGatewayRequest{
		Project:            a.id.Project,
		ExternalVpnGateway: a.id.ExternalVPNGateway,
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeExternalVPNGateway %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *ComputeExternalVPNGatewayAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeExternalVPNGateway", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = direct.LazyPtr(a.id.ExternalVPNGateway)

	req := &pb.InsertExternalVpnGatewayRequest{
		Project:                    a.id.Project,
		ExternalVpnGatewayResource: desired,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeExternalVPNGateway %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting ComputeExternalVPNGateway %s creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute ComputeExternalVPNGateway in gcp", "name", a.id)

	getReq := &pb.GetExternalVpnGatewayRequest{
		Project:            a.id.Project,
		ExternalVpnGateway: a.id.ExternalVPNGateway,
	}
	created, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting ComputeExternalVPNGateway %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeExternalVPNGatewayAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeExternalVPNGateway", "name", a.id)

	diffs, _, err := compareExternalVPNGateway(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	fields := diffs.FieldsByID()
	_, labelsChanged := fields["labels"]
	delete(fields, "labels")

	if len(fields) > 0 {
		// Surfacing exact diff back to the user
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		// Since fields in ComputeExternalVPNGateway are immutable, any other change is not supported.
		return fmt.Errorf("ComputeExternalVPNGateway is immutable and cannot be updated")
	}

	if labelsChanged {
		log.V(2).Info("updating ComputeExternalVPNGateway labels", "name", a.id)
		req := &pb.SetLabelsExternalVpnGatewayRequest{
			Project:  a.id.Project,
			Resource: a.id.ExternalVPNGateway,
			GlobalSetLabelsRequestResource: &pb.GlobalSetLabelsRequest{
				Labels:           a.desired.Labels,
				LabelFingerprint: a.actual.LabelFingerprint,
			},
		}
		op, err := a.gcpClient.SetLabels(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ComputeExternalVPNGateway labels %s: %w", a.id, err)
		}
		if err = op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting ComputeExternalVPNGateway %s labels update: %w", a.id, err)
		}

		// Retrieve latest state after update
		getReq := &pb.GetExternalVpnGatewayRequest{
			Project:            a.id.Project,
			ExternalVpnGateway: a.id.ExternalVPNGateway,
		}
		latest, err := a.gcpClient.Get(ctx, getReq)
		if err != nil {
			return fmt.Errorf("getting ComputeExternalVPNGateway %s after labels update: %w", a.id, err)
		}
		return a.updateStatus(ctx, updateOp, latest)
	}

	return nil
}

func (a *ComputeExternalVPNGatewayAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeExternalVPNGateway{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeExternalVPNGatewaySpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ExternalVPNGateway)
	u.SetGroupVersionKind(krm.ComputeExternalVPNGatewayGVK)

	u.Object = uObj
	return u, nil
}

func (a *ComputeExternalVPNGatewayAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeExternalVPNGateway", "name", a.id)

	req := &pb.DeleteExternalVpnGatewayRequest{
		Project:            a.id.Project,
		ExternalVpnGateway: a.id.ExternalVPNGateway,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeExternalVPNGateway %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute ComputeExternalVPNGateway", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of ComputeExternalVPNGateway %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *ComputeExternalVPNGatewayAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ExternalVpnGateway) error {
	mapCtx := &direct.MapContext{}
	status := ComputeExternalVPNGatewayStatus_v1beta1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return op.UpdateStatus(ctx, status, nil)
}

func compareExternalVPNGateway(ctx context.Context, actual, desired *pb.ExternalVpnGateway) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeExternalVPNGatewaySpec_v1beta1_FromProto, ComputeExternalVPNGatewaySpec_v1beta1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
