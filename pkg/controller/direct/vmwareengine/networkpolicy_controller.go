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
// proto.message: google.cloud.vmwareengine.v1.NetworkPolicy
// crd.type: VMwareEngineNetworkPolicy
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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.VMwareEngineNetworkPolicyGVK, NewNetworkPolicyModel)
}

func NewNetworkPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkPolicyModel{config: *config}, nil
}

var _ directbase.Model = &networkPolicyModel{}

type networkPolicyModel struct {
	config config.ControllerConfig
}

func (m *networkPolicyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.VMwareEngineNetworkPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNetworkPolicyIdentity(ctx, reader, obj)
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
	return &networkPolicyAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *networkPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type networkPolicyAdapter struct {
	gcpClient *gcp.Client
	id        *krm.NetworkPolicyIdentity
	desired   *krm.VMwareEngineNetworkPolicy
	actual    *pb.NetworkPolicy
	reader    client.Reader
}

var _ directbase.Adapter = &networkPolicyAdapter{}

func (a *networkPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting vmwareengine network policy", "name", a.id)

	req := &pb.GetNetworkPolicyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetNetworkPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vmwareengine network policy %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *networkPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating vmwareengine network policy", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineNetworkPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateNetworkPolicyRequest{
		Parent:          a.id.Parent().String(),
		NetworkPolicyId: a.id.ID(),
		NetworkPolicy:   resource,
	}
	op, err := a.gcpClient.CreateNetworkPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vmwareengine network policy %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine network policy %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created vmwareengine network policy in gcp", "name", a.id)

	status := &krm.VMwareEngineNetworkPolicyStatus{}
	status.ObservedState = VMwareEngineNetworkPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *networkPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating vmwareengine network policy", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineNetworkPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.EdgeServicesCIDR != nil && !reflect.DeepEqual(resource.EdgeServicesCidr, a.actual.EdgeServicesCidr) {
		paths = append(paths, "edge_services_cidr")
	}
	if desired.Spec.InternetAccess != nil && resource.GetInternetAccess().GetEnabled() != a.actual.GetInternetAccess().GetEnabled() {
		paths = append(paths, "internet_access.enabled")
	}
	if desired.Spec.ExternalIP != nil && resource.GetExternalIp().GetEnabled() != a.actual.GetExternalIp().GetEnabled() {
		paths = append(paths, "external_ip.enabled")
	}
	if desired.Spec.VMwareEngineNetworkRef != nil && !reflect.DeepEqual(resource.VmwareEngineNetwork, a.actual.VmwareEngineNetwork) {
		paths = append(paths, "vmware_engine_network")
	}

	var updated *pb.NetworkPolicy
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateNetworkPolicyRequest{
			NetworkPolicy: resource,
			UpdateMask:    &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateNetworkPolicy(ctx, req)
		if err != nil {
			return fmt.Errorf("updating vmwareengine network policy %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("vmwareengine network policy %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated vmwareengine network policy", "name", a.id)
	}

	status := &krm.VMwareEngineNetworkPolicyStatus{}
	status.ObservedState = VMwareEngineNetworkPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *networkPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VMwareEngineNetworkPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VMwareEngineNetworkPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VMwareEngineNetworkPolicyGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *networkPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting vmwareengine network policy", "name", a.id)

	req := &pb.DeleteNetworkPolicyRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteNetworkPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BackupVault, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting vmwareengine network policy %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted vmwareengine network policy", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupVault %s: %w", a.id, err)
	}
	return true, nil
}

func (a *networkPolicyAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.VMwareEngineNetworkRef != nil {
		if _, err := obj.Spec.VMwareEngineNetworkRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	return nil
}
