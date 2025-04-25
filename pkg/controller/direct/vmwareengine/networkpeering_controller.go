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

// +tool:controller
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.NetworkPeering
// crd.type: VMwareEngineNetworkPeering
// crd.version: v1alpha1

package vmwareengine

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/vmwareengine/apiv1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.VMwareEngineNetworkPeeringGVK, NewNetworkPeeringModel)
}

func NewNetworkPeeringModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkPeeringModel{config: *config}, nil
}

var _ directbase.Model = &networkPeeringModel{}

type networkPeeringModel struct {
	config config.ControllerConfig
}

func (m *networkPeeringModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.VMwareEngineNetworkPeering{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNetworkPeeringIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get VMwareEngine GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}
	return &networkPeeringAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *networkPeeringModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type networkPeeringAdapter struct {
	gcpClient *gcp.Client
	id        *krm.NetworkPeeringIdentity
	desired   *krm.VMwareEngineNetworkPeering
	actual    *pb.NetworkPeering
	reader    client.Reader
}

var _ directbase.Adapter = &networkPeeringAdapter{}

func (a *networkPeeringAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting vmwareengine networkpeering", "name", a.id)

	req := &pb.GetNetworkPeeringRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetNetworkPeering(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vmwareengine networkpeering %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *networkPeeringAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating vmwareengine networkpeering", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineNetworkPeeringSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateNetworkPeeringRequest{
		Parent:           a.id.Parent().String(),
		NetworkPeeringId: a.id.ID(),
		NetworkPeering:   resource,
	}
	op, err := a.gcpClient.CreateNetworkPeering(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vmwareengine networkpeering %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine networkpeering %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created vmwareengine networkpeering in gcp", "name", a.id)

	status := &krm.VMwareEngineNetworkPeeringStatus{}
	status.ObservedState = VMwareEngineNetworkPeeringObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *networkPeeringAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating vmwareengine networkpeering", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineNetworkPeeringSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var paths []string
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.PeerNetwork != nil && !reflect.DeepEqual(resource.PeerNetwork, a.actual.PeerNetwork) {
		paths = append(paths, "peer_network")
	}
	if desired.Spec.PeerNetworkType != nil && !reflect.DeepEqual(resource.PeerNetworkType, a.actual.PeerNetworkType) {
		paths = append(paths, "peer_network_type")
	}
	if desired.Spec.ExportCustomRoutes != nil && !reflect.DeepEqual(resource.ExportCustomRoutes, a.actual.ExportCustomRoutes) {
		paths = append(paths, "export_custom_routes")
	}
	if desired.Spec.ImportCustomRoutes != nil && !reflect.DeepEqual(resource.ImportCustomRoutes, a.actual.ImportCustomRoutes) {
		paths = append(paths, "import_custom_routes")
	}
	if desired.Spec.ExchangeSubnetRoutes != nil && !reflect.DeepEqual(resource.ExchangeSubnetRoutes, a.actual.ExchangeSubnetRoutes) {
		paths = append(paths, "exchange_subnet_routes")
	}
	if desired.Spec.ExportCustomRoutesWithPublicIP != nil && !reflect.DeepEqual(resource.ExportCustomRoutesWithPublicIp, a.actual.ExportCustomRoutesWithPublicIp) {
		paths = append(paths, "export_custom_routes_with_public_ip")
	}
	if desired.Spec.ImportCustomRoutesWithPublicIP != nil && !reflect.DeepEqual(resource.ImportCustomRoutesWithPublicIp, a.actual.ImportCustomRoutesWithPublicIp) {
		paths = append(paths, "import_custom_routes_with_public_ip")
	}
	if desired.Spec.PeerMTU != nil && !reflect.DeepEqual(resource.PeerMtu, a.actual.PeerMtu) {
		paths = append(paths, "peer_mtu")
	}
	if desired.Spec.VMwareEngineNetworkRef != nil && !reflect.DeepEqual(resource.VmwareEngineNetwork, a.actual.VmwareEngineNetwork) {
		paths = append(paths, "vmware_engine_network")
	}

	var updated *pb.NetworkPeering
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateNetworkPeeringRequest{
			NetworkPeering: resource,
			UpdateMask:     &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateNetworkPeering(ctx, req)
		if err != nil {
			return fmt.Errorf("updating vmwareengine networkpeering %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("vmwareengine networkpeering %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated vmwareengine networkpeering", "name", a.id)
	}

	status := &krm.VMwareEngineNetworkPeeringStatus{}
	status.ObservedState = VMwareEngineNetworkPeeringObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *networkPeeringAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VMwareEngineNetworkPeering{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VMwareEngineNetworkPeeringSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VMwareEngineNetworkPeeringGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *networkPeeringAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting vmwareengine networkpeering", "name", a.id)

	req := &pb.DeleteNetworkPeeringRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteNetworkPeering(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BackupVault, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting vmwareengine networkpeering %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted vmwareengine networkpeering", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupVault %s: %w", a.id, err)
	}
	return true, nil
}

func (a *networkPeeringAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired

	// normalize reference fields
	if obj.Spec.PeerNetwork != nil {
		if obj.Spec.PeerNetwork.ComputeNetworkRef != nil {
			if err := obj.Spec.PeerNetwork.ComputeNetworkRef.Normalize(ctx, a.reader, obj); err != nil {
				return err
			}
		}
		if obj.Spec.PeerNetwork.VMwareEngineNetworkRef != nil {
			if _, err := obj.Spec.PeerNetwork.VMwareEngineNetworkRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
	}
	if obj.Spec.VMwareEngineNetworkRef != nil {
		if _, err := obj.Spec.VMwareEngineNetworkRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}

	return nil
}
