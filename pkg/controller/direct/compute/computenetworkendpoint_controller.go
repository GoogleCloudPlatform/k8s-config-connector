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
// proto.service: google.cloud.compute.v1.NetworkEndpointGroups
// proto.message: google.cloud.compute.v1.NetworkEndpoint
// crd.type: ComputeNetworkEndpoint
// crd.version: v1alpha1

package compute

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeNetworkEndpointGVK, NewComputeNetworkEndpointModel)
}

func NewComputeNetworkEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeNetworkEndpointModel{config: config}, nil
}

var _ directbase.Model = &computeNetworkEndpointModel{}

type computeNetworkEndpointModel struct {
	config *config.ControllerConfig
}

func (m *computeNetworkEndpointModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeNetworkEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	endpointId := id.(*krm.ComputeNetworkEndpointIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	client, err := gcpClient.newNetworkEndpointGroupsClient(ctx)
	if err != nil {
		return nil, err
	}

	desired, err := buildDesiredEndpoint(obj, endpointId)
	if err != nil {
		return nil, err
	}

	return &ComputeNetworkEndpointAdapter{
		gcpClient: client,
		id:        endpointId,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *computeNetworkEndpointModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeNetworkEndpointAdapter struct {
	gcpClient *compute.NetworkEndpointGroupsClient
	id        *krm.ComputeNetworkEndpointIdentity
	desired   *pb.NetworkEndpoint
	actual    *pb.NetworkEndpoint
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeNetworkEndpointAdapter{}

func (a *ComputeNetworkEndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeNetworkEndpoint", "id", a.id)

	req := &pb.ListNetworkEndpointsNetworkEndpointGroupsRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		NetworkEndpointGroup: a.id.ComputeNetworkEndpointGroup,
	}
	it := a.gcpClient.ListNetworkEndpoints(ctx, req)
	for {
		item, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("listing network endpoints in %s: %w", a.id.ParentString(), err)
		}
		ep := item.GetNetworkEndpoint()
		if ep == nil {
			continue
		}
		if a.matches(ep) {
			a.actual = ep
			return true, nil
		}
	}

	return false, nil
}

func (a *ComputeNetworkEndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeNetworkEndpoint", "id", a.id)

	req := &pb.AttachNetworkEndpointsNetworkEndpointGroupRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		NetworkEndpointGroup: a.id.ComputeNetworkEndpointGroup,
		NetworkEndpointGroupsAttachEndpointsRequestResource: &pb.NetworkEndpointGroupsAttachEndpointsRequest{
			NetworkEndpoints: []*pb.NetworkEndpoint{a.desired},
		},
	}
	err := wait.PollUntilContextTimeout(ctx, 2*time.Second, 2*time.Minute, true, func(ctx context.Context) (bool, error) {
		op, err := a.gcpClient.AttachNetworkEndpoints(ctx, req)
		if err != nil {
			if isConflict(err) {
				log.V(2).Info("conflict attaching network endpoint, retrying", "error", err)
				return false, nil
			}
			return false, err
		}
		if !op.Done() {
			if err := op.Wait(ctx); err != nil {
				if isConflict(err) {
					log.V(2).Info("conflict waiting for network endpoint attach, retrying", "error", err)
					return false, nil
				}
				return false, err
			}
		}
		return true, nil
	})
	if err != nil {
		return fmt.Errorf("attaching ComputeNetworkEndpoint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully attached ComputeNetworkEndpoint", "id", a.id)

	a.actual = a.desired
	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *ComputeNetworkEndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeNetworkEndpoint", "id", a.id)

	diffs, _, err := compareComputeNetworkEndpoint(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("ComputeNetworkEndpoint is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *ComputeNetworkEndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeNetworkEndpoint", "id", a.id)

	target := a.actual
	if target == nil {
		target = a.desired
	}
	req := &pb.DetachNetworkEndpointsNetworkEndpointGroupRequest{
		Project:              a.id.Project,
		Zone:                 a.id.Zone,
		NetworkEndpointGroup: a.id.ComputeNetworkEndpointGroup,
		NetworkEndpointGroupsDetachEndpointsRequestResource: &pb.NetworkEndpointGroupsDetachEndpointsRequest{
			NetworkEndpoints: []*pb.NetworkEndpoint{target},
		},
	}
	err := wait.PollUntilContextTimeout(ctx, 2*time.Second, 2*time.Minute, true, func(ctx context.Context) (bool, error) {
		op, err := a.gcpClient.DetachNetworkEndpoints(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return true, nil
			}
			if isConflict(err) {
				log.V(2).Info("conflict detaching network endpoint, retrying", "error", err)
				return false, nil
			}
			return false, err
		}
		if !op.Done() {
			if err := op.Wait(ctx); err != nil {
				if direct.IsNotFound(err) {
					return true, nil
				}
				if isConflict(err) {
					log.V(2).Info("conflict waiting for network endpoint detach, retrying", "error", err)
					return false, nil
				}
				return false, err
			}
		}
		return true, nil
	})
	if err != nil {
		return false, fmt.Errorf("detaching ComputeNetworkEndpoint %s: %w", a.id, err)
	}
	log.V(2).Info("successfully detached ComputeNetworkEndpoint", "id", a.id)
	return true, nil
}

func (a *ComputeNetworkEndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNetworkEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeNetworkEndpointSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Zone = a.id.Zone
	obj.Spec.NetworkEndpointGroupRef = &computev1beta1.ComputeNetworkEndpointGroupRef{
		Name: a.id.ComputeNetworkEndpointGroup,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj

	var nameParts []string
	if a.actual.GetInstance() != "" {
		nameParts = append(nameParts, lastComponent(a.actual.GetInstance()))
	}
	if a.actual.GetIpAddress() != "" {
		nameParts = append(nameParts, strings.ReplaceAll(a.actual.GetIpAddress(), ".", "-"))
	}
	port := a.id.Port
	if port == "" && a.actual.GetPort() != 0 {
		port = strconv.Itoa(int(a.actual.GetPort()))
	}
	if port != "" {
		nameParts = append(nameParts, port)
	}
	if len(nameParts) > 0 {
		u.SetName(strings.Join(nameParts, "-"))
	} else {
		u.SetName(a.id.ComputeNetworkEndpointGroup)
	}

	u.SetGroupVersionKind(krm.ComputeNetworkEndpointGVK)
	return u, nil
}

func (a *ComputeNetworkEndpointAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.NetworkEndpoint) error {
	status := &krm.ComputeNetworkEndpointStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ComputeNetworkEndpointAdapter) matches(ep *pb.NetworkEndpoint) bool {
	if ep.GetIpAddress() != a.desired.GetIpAddress() {
		return false
	}
	if ep.GetPort() != a.desired.GetPort() {
		return false
	}
	epInstance := lastComponent(ep.GetInstance())
	desiredInstance := lastComponent(a.desired.GetInstance())
	if epInstance != desiredInstance {
		return false
	}
	return true
}

func buildDesiredEndpoint(obj *krm.ComputeNetworkEndpoint, endpointId *krm.ComputeNetworkEndpointIdentity) (*pb.NetworkEndpoint, error) {
	mapCtx := &direct.MapContext{}
	desired := ComputeNetworkEndpointSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if endpointId.Instance != "" {
		desired.Instance = direct.PtrTo(endpointId.Instance)
	}
	if desired.Port == nil && endpointId.Port != "" {
		portVal, err := strconv.Atoi(endpointId.Port)
		if err != nil {
			return nil, fmt.Errorf("invalid port %q: %w", endpointId.Port, err)
		}
		desired.Port = direct.PtrTo(int32(portVal))
	}
	return desired, nil
}

func isConflict(err error) bool {
	if err == nil {
		return false
	}
	if direct.HasHTTPCode(err, 409) {
		return true
	}
	if strings.Contains(err.Error(), "RESOURCE_IN_USE_BY_ANOTHER_OPERATION") ||
		strings.Contains(err.Error(), "already being used by another operation") {
		return true
	}
	return false
}

func compareComputeNetworkEndpoint(ctx context.Context, actual, desired *pb.NetworkEndpoint) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	krmActual := ComputeNetworkEndpointSpec_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := ComputeNetworkEndpointSpec_v1alpha1_ToProto(mapCtx, krmActual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	clonedDesired := proto.Clone(desired).(*pb.NetworkEndpoint)

	if maskedActual.Instance != nil {
		maskedActual.Instance = direct.PtrTo(lastComponent(*maskedActual.Instance))
	}
	if clonedDesired.Instance != nil {
		clonedDesired.Instance = direct.PtrTo(lastComponent(*clonedDesired.Instance))
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
