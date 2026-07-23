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

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityDNSThreatDetectorGVK, NewDNSThreatDetectorModel)
}

func NewDNSThreatDetectorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &dnsThreatDetectorModel{config: *config}, nil
}

var _ directbase.Model = &dnsThreatDetectorModel{}

type dnsThreatDetectorModel struct {
	config config.ControllerConfig
}

func (m *dnsThreatDetectorModel) client(ctx context.Context) (pb.DnsThreatDetectorServiceClient, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}

	opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("dialing networksecurity service: %w", err)
	}

	return pb.NewDnsThreatDetectorServiceClient(conn), nil
}

func (m *dnsThreatDetectorModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityDNSThreatDetector{}
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
	desired := NetworkSecurityDNSThreatDetectorSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &dnsThreatDetectorAdapter{
		gcpClient: gcpClient,
		id:        id.(*krm.NetworkSecurityDNSThreatDetectorIdentity),
		desired:   desired,
	}, nil
}

func (m *dnsThreatDetectorModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type dnsThreatDetectorAdapter struct {
	gcpClient pb.DnsThreatDetectorServiceClient
	id        *krm.NetworkSecurityDNSThreatDetectorIdentity
	desired   *pb.DnsThreatDetector
	actual    *pb.DnsThreatDetector
}

var _ directbase.Adapter = &dnsThreatDetectorAdapter{}

func (a *dnsThreatDetectorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting networksecurity dns threat detector", "name", a.id)

	req := &pb.GetDnsThreatDetectorRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetDnsThreatDetector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity dns threat detector %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *dnsThreatDetectorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating networksecurity dns threat detector", "name", a.id)

	parent := a.id.ParentString()
	req := &pb.CreateDnsThreatDetectorRequest{
		Parent:              parent,
		DnsThreatDetectorId: a.id.DNSThreatDetector,
		DnsThreatDetector:   a.desired,
	}
	actual, err := a.gcpClient.CreateDnsThreatDetector(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networksecurity dns threat detector %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created networksecurity dns threat detector", "name", a.id.String())

	return a.updateStatus(ctx, createOp, actual)
}

func (a *dnsThreatDetectorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating networksecurity dns threat detector", "name", a.id)

	diffs, updateMask, err := compareDNSThreatDetector(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	log.V(2).Info("fields need update", "name", a.id, "updateMask", updateMask)

	req := &pb.UpdateDnsThreatDetectorRequest{
		UpdateMask:        updateMask,
		DnsThreatDetector: a.desired,
	}
	req.DnsThreatDetector.Name = a.id.String()

	actual, err := a.gcpClient.UpdateDnsThreatDetector(ctx, req)
	if err != nil {
		return fmt.Errorf("updating networksecurity dns threat detector %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated networksecurity dns threat detector", "name", a.id)

	return a.updateStatus(ctx, updateOp, actual)
}

func compareDNSThreatDetector(ctx context.Context, actual, desired *pb.DnsThreatDetector) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityDNSThreatDetectorSpec_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := NetworkSecurityDNSThreatDetectorSpec_v1alpha1_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *dnsThreatDetectorAdapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.DnsThreatDetector) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityDNSThreatDetectorStatus{}
	status.ObservedState = NetworkSecurityDNSThreatDetectorObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *dnsThreatDetectorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *dnsThreatDetectorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting networksecurity dns threat detector", "name", a.id)

	req := &pb.DeleteDnsThreatDetectorRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteDnsThreatDetector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting networksecurity dns threat detector %s: %w", a.id.String(), err)
	}

	return true, nil
}
