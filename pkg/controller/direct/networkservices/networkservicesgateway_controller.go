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

package networkservices

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/networkservices/apiv1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1beta1"
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
	registry.RegisterModel(krm.NetworkServicesGatewayGVK, NewNetworkServicesGatewayModel)
}

func NewNetworkServicesGatewayModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNetworkServicesGateway{config: *config}, nil
}

var _ directbase.Model = &modelNetworkServicesGateway{}

type modelNetworkServicesGateway struct {
	config config.ControllerConfig
}

func (m *modelNetworkServicesGateway) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NetworkServicesGateway client: %w", err)
	}
	return gcpClient, err
}

func (m *modelNetworkServicesGateway) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkServicesGateway{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewNetworkServicesGatewayIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesGatewaySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Name = id.String()
	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &NetworkServicesGatewayAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelNetworkServicesGateway) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkServicesGatewayAdapter struct {
	id           *krm.NetworkServicesGatewayIdentity
	gcpClient    *gcp.Client
	desiredProto *pb.Gateway
	actual       *pb.Gateway
}

var _ directbase.Adapter = &NetworkServicesGatewayAdapter{}

func (a *NetworkServicesGatewayAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NetworkServicesGateway", "name", a.id)

	req := &pb.GetGatewayRequest{Name: a.id.String()}
	gatewaypb, err := a.gcpClient.GetGateway(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkServicesGateway %q: %w", a.id, err)
	}

	a.actual = gatewaypb
	return true, nil
}

func (a *NetworkServicesGatewayAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkServicesGateway", "name", a.id)

	req := &pb.CreateGatewayRequest{
		Parent:    a.id.ParentString(),
		GatewayId: a.id.Gateway,
		Gateway:   a.desiredProto,
	}
	op, err := a.gcpClient.CreateGateway(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkServicesGateway %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NetworkServicesGateway %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created NetworkServicesGateway", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *NetworkServicesGatewayAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkServicesGateway", "name", a.id)

	diffs, updateMask, err := compareGateway(ctx, a.actual, a.desiredProto)
	if err != nil {
		return fmt.Errorf("comparing NetworkServicesGateway %s: %w", a.id, err)
	}

	latest := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no changes detected for NetworkServicesGateway", "name", a.id)
	} else {
		// Report exact diffs
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateGatewayRequest{
			UpdateMask: updateMask,
			Gateway:    a.desiredProto,
		}
		op, err := a.gcpClient.UpdateGateway(ctx, req)
		if err != nil {
			return fmt.Errorf("updating NetworkServicesGateway %s: %w", a.id, err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("NetworkServicesGateway %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated NetworkServicesGateway", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *NetworkServicesGatewayAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesGateway{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesGatewaySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Gateway)
	u.SetGroupVersionKind(krm.NetworkServicesGatewayGVK)

	u.Object = uObj
	return u, nil
}

func (a *NetworkServicesGatewayAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkServicesGateway", "name", a.id)

	req := &pb.DeleteGatewayRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteGateway(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent NetworkServicesGateway, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkServicesGateway %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted NetworkServicesGateway", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete NetworkServicesGateway %s: %w", a.id, err)
	}
	return true, nil
}

func compareGateway(ctx context.Context, actual, desired *pb.Gateway) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkServicesGatewaySpec_FromProto, NetworkServicesGatewaySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = actual.Labels

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Gateway) {
		// Add any server-side or GCP defaults if known, or leave empty
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *NetworkServicesGatewayAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Gateway) error {
	mapCtx := &direct.MapContext{}
	status := NetworkServicesGatewayStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
