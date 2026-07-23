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
// proto.service: google.cloud.networkmanagement.v1.VpcFlowLogsService
// proto.message: google.cloud.networkmanagement.v1.VpcFlowLogsConfig
// crd.type: NetworkManagementVPCFlowLogsConfig
// crd.version: v1alpha1

package networkmanagement

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/networkmanagement/apiv1"
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.NetworkManagementVPCFlowLogsConfigGVK, NewVpcFlowLogsConfigModel)
}

func NewVpcFlowLogsConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &vpcFlowLogsConfigModel{config: config}, nil
}

var _ directbase.Model = &vpcFlowLogsConfigModel{}

type vpcFlowLogsConfigModel struct {
	config *config.ControllerConfig
}

func (m *vpcFlowLogsConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkManagementVPCFlowLogsConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewVpcFlowLogsConfigIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newVpcFlowLogsClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	adapter := &vpcFlowLogsConfigAdapter{
		gcpClient:  gcpClient,
		id:         id,
		desiredKRM: obj,
		reader:     reader,
	}

	if err := adapter.normalizeReferenceFields(ctx); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	adapter.desired = NetworkManagementVPCFlowLogsConfigSpec_ToProto(mapCtx, &adapter.desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return adapter, nil
}

func (m *vpcFlowLogsConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type vpcFlowLogsConfigAdapter struct {
	gcpClient  *gcp.VpcFlowLogsClient
	id         *krm.VpcFlowLogsConfigIdentity
	desiredKRM *krm.NetworkManagementVPCFlowLogsConfig
	desired    *pb.VpcFlowLogsConfig
	actual     *pb.VpcFlowLogsConfig
	reader     client.Reader
}

var _ directbase.Adapter = &vpcFlowLogsConfigAdapter{}

func (a *vpcFlowLogsConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networkmanagement vpcflowlogsconfig", "name", a.id)

	req := &pb.GetVpcFlowLogsConfigRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetVpcFlowLogsConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networkmanagement vpcflowlogsconfig %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *vpcFlowLogsConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating networkmanagement vpcflowlogsconfig", "name", a.id)

	req := &pb.CreateVpcFlowLogsConfigRequest{
		Parent:              a.id.Parent().String(),
		VpcFlowLogsConfigId: a.id.ID(),
		VpcFlowLogsConfig:   a.desired,
	}
	op, err := a.gcpClient.CreateVpcFlowLogsConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networkmanagement vpcflowlogsconfig %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("networkmanagement vpcflowlogsconfig %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created networkmanagement vpcflowlogsconfig in gcp", "name", a.id)

	// Perform GET immediately after LRO to get fully populated status fields
	latest, err := a.gcpClient.GetVpcFlowLogsConfig(ctx, &pb.GetVpcFlowLogsConfigRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created networkmanagement vpcflowlogsconfig %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *vpcFlowLogsConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networkmanagement vpcflowlogsconfig", "name", a.id)

	a.desired.Name = a.id.String()

	diffs, updateMask, err := compareVpcFlowLogsConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var latest *pb.VpcFlowLogsConfig
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		latest = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, diffs)
		log.V(2).Info("updating fields", "name", a.id, "paths", updateMask.Paths)

		req := &pb.UpdateVpcFlowLogsConfigRequest{
			UpdateMask:        updateMask,
			VpcFlowLogsConfig: a.desired,
		}
		op, err := a.gcpClient.UpdateVpcFlowLogsConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating networkmanagement vpcflowlogsconfig %s: %w", a.id.String(), err)
		}
		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("networkmanagement vpcflowlogsconfig %s waiting for update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated networkmanagement vpcflowlogsconfig", "name", a.id)

		// Perform GET immediately after LRO to get fully populated status fields
		latest, err = a.gcpClient.GetVpcFlowLogsConfig(ctx, &pb.GetVpcFlowLogsConfigRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting updated networkmanagement vpcflowlogsconfig %s: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *vpcFlowLogsConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkManagementVPCFlowLogsConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkManagementVPCFlowLogsConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = direct.PtrTo(a.id.Parent().Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkManagementVPCFlowLogsConfigGVK)
	return u, nil
}

func (a *vpcFlowLogsConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networkmanagement vpcflowlogsconfig", "name", a.id)

	req := &pb.DeleteVpcFlowLogsConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteVpcFlowLogsConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent networkmanagement vpcflowlogsconfig, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting networkmanagement vpcflowlogsconfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated deletion of networkmanagement vpcflowlogsconfig", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of networkmanagement vpcflowlogsconfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted networkmanagement vpcflowlogsconfig", "name", a.id)
	return true, nil
}

func (a *vpcFlowLogsConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.VpcFlowLogsConfig) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkManagementVPCFlowLogsConfigStatus{}
	status.ObservedState = NetworkManagementVPCFlowLogsConfigObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *vpcFlowLogsConfigAdapter) normalizeReferenceFields(ctx context.Context) error {
	if a.desiredKRM.Spec.VPNTunnelRef != nil {
		if err := resolveComputeVPNTunnelRef(ctx, a.reader, a.desiredKRM.GetNamespace(), a.desiredKRM.Spec.VPNTunnelRef); err != nil {
			return err
		}
	}
	if a.desiredKRM.Spec.InterconnectAttachmentRef != nil {
		if err := resolveComputeInterconnectAttachmentRef(ctx, a.reader, a.desiredKRM.GetNamespace(), a.desiredKRM.Spec.InterconnectAttachmentRef); err != nil {
			return err
		}
	}
	return nil
}

func compareVpcFlowLogsConfig(ctx context.Context, actual, desired *pb.VpcFlowLogsConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NetworkManagementVPCFlowLogsConfigSpec_FromProto, NetworkManagementVPCFlowLogsConfigSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func resolveComputeVPNTunnelRef(ctx context.Context, reader client.Reader, defaultNamespace string, ref *computev1beta1.ComputeVPNTunnelRef) error {
	if ref == nil {
		return nil
	}
	if ref.External != "" {
		return nil
	}
	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on reference")
	}
	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeVPNTunnel",
	})
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("error reading referenced ComputeVPNTunnel %v: %w", key, err)
	}

	selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
	if selfLink == "" {
		return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	}

	externalRef := refs.TrimComputeURIPrefix(selfLink)
	ref.External = externalRef
	return nil
}

func resolveComputeInterconnectAttachmentRef(ctx context.Context, reader client.Reader, defaultNamespace string, ref *computev1beta1.ComputeInterconnectAttachmentRef) error {
	if ref == nil {
		return nil
	}
	if ref.External != "" {
		return nil
	}
	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on reference")
	}
	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeInterconnectAttachment",
	})
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("error reading referenced ComputeInterconnectAttachment %v: %w", key, err)
	}

	selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
	if selfLink == "" {
		return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
	}

	externalRef := refs.TrimComputeURIPrefix(selfLink)
	ref.External = externalRef
	return nil
}
