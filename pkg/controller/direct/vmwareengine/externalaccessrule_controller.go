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
// proto.message: google.cloud.vmwareengine.v1.ExternalAccessRule
// crd.type: VMwareEngineExternalAccessRule
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.VMwareEngineExternalAccessRuleGVK, NewExternalAccessRuleModel)
}

func NewExternalAccessRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &externalAccessRuleModel{config: *config}, nil
}

var _ directbase.Model = &externalAccessRuleModel{}

type externalAccessRuleModel struct {
	config config.ControllerConfig
}

func (m *externalAccessRuleModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.VMwareEngineExternalAccessRule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewExternalAccessRuleIdentity(ctx, reader, obj)
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
	return &externalAccessRuleAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *externalAccessRuleModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type externalAccessRuleAdapter struct {
	gcpClient *gcp.Client
	id        *krm.ExternalAccessRuleIdentity
	desired   *krm.VMwareEngineExternalAccessRule
	actual    *pb.ExternalAccessRule
	reader    client.Reader
}

var _ directbase.Adapter = &externalAccessRuleAdapter{}

func (a *externalAccessRuleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting vmwareengine external access rule", "name", a.id)

	req := &pb.GetExternalAccessRuleRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetExternalAccessRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vmwareengine external access rule %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *externalAccessRuleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating vmwareengine external access rule", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineExternalAccessRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateExternalAccessRuleRequest{
		Parent:               a.id.Parent().String(),
		ExternalAccessRuleId: a.id.ID(),
		ExternalAccessRule:   resource,
	}
	op, err := a.gcpClient.CreateExternalAccessRule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vmwareengine external access rule %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine external access rule %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created vmwareengine external access rule in gcp", "name", a.id)

	status := &krm.VMwareEngineExternalAccessRuleStatus{}
	status.ObservedState = VMwareEngineExternalAccessRuleObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *externalAccessRuleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating vmwareengine external access rule", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEngineExternalAccessRuleSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.Priority != nil && !reflect.DeepEqual(resource.Priority, a.actual.Priority) {
		paths = append(paths, "priority")
	}
	if desired.Spec.Action != nil && !reflect.DeepEqual(resource.Action, a.actual.Action) {
		paths = append(paths, "action")
	}
	if desired.Spec.IPProtocol != nil && !reflect.DeepEqual(resource.IpProtocol, a.actual.IpProtocol) {
		paths = append(paths, "ip_protocol")
	}
	if len(desired.Spec.SourcePorts) != 0 && !reflect.DeepEqual(resource.SourcePorts, a.actual.SourcePorts) {
		paths = append(paths, "source_ports")
	}
	if len(desired.Spec.DestinationPorts) != 0 && !reflect.DeepEqual(resource.DestinationPorts, a.actual.DestinationPorts) {
		paths = append(paths, "destination_ports")
	}
	// TODO: handle source_ip_ranges and destination_ip_ranges which contain oneof fields which requires special handling
	// e.g. https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/7e47fab0c19e7908b88ad1d5e1e6063d27e35fc4/pkg/controller/direct/gkebackup/backupplan_controller.go#L192

	var updated *pb.ExternalAccessRule
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateExternalAccessRuleRequest{
			ExternalAccessRule: resource,
			UpdateMask:         &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateExternalAccessRule(ctx, req)
		if err != nil {
			return fmt.Errorf("updating vmwareengine external access rule %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("vmwareengine external access rule %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated vmwareengine external access rule", "name", a.id)
	}

	status := &krm.VMwareEngineExternalAccessRuleStatus{}
	status.ObservedState = VMwareEngineExternalAccessRuleObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *externalAccessRuleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VMwareEngineExternalAccessRule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VMwareEngineExternalAccessRuleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.NetworkPolicyRef = &krm.NetworkPolicyRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VMwareEngineExternalAccessRuleGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *externalAccessRuleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting vmwareengine external access rule", "name", a.id)

	req := &pb.DeleteExternalAccessRuleRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteExternalAccessRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BackupVault, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting vmwareengine external access rule %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted vmwareengine external access rule", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete BackupVault %s: %w", a.id, err)
	}
	return true, nil
}

func (a *externalAccessRuleAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired

	for i := range obj.Spec.SourceIPRanges {
		if obj.Spec.SourceIPRanges[i].ExternalAddressRef != nil {
			_, err := obj.Spec.SourceIPRanges[i].ExternalAddressRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
			if err != nil {
				return err
			}
		}
	}

	for i := range obj.Spec.DestinationIPRanges {
		if obj.Spec.DestinationIPRanges[i].ExternalAddressRef != nil {
			_, err := obj.Spec.DestinationIPRanges[i].ExternalAddressRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
