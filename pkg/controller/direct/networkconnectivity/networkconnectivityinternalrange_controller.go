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
// proto.service: google.cloud.networkconnectivity.v1.InternalRange
// proto.message: google.cloud.networkconnectivity.v1.InternalRange
// crd.type: NetworkConnectivityInternalRange
// crd.version: v1alpha1

package networkconnectivity

import (
	"context"
	"fmt"
	"reflect"

	gcpapi "cloud.google.com/go/networkconnectivity/apiv1"
	pb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/networkconnectivityrefs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityInternalRangeGVK, NewInternalRangeModel)
}

func NewInternalRangeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &internalRangeModel{config: *config}, nil
}

var _ directbase.Model = &internalRangeModel{}

type internalRangeModel struct {
	config config.ControllerConfig
}

func (m *internalRangeModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkConnectivityInternalRange{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idIdentity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idIdentity.(*networkconnectivityrefs.NetworkConnectivityInternalRangeIdentity)

	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve projectID")
	}

	if err := common.VisitFields(obj, &refNormalizer{ctx: ctx, src: obj, project: *projectRef, kube: reader}); err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newInternalRangeClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkConnectivityInternalRangeSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &internalRangeAdapter{
		gcpClient: client,
		id:        id,
		desired:   desiredProto,
	}, nil
}

func (m *internalRangeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type internalRangeAdapter struct {
	gcpClient *gcpapi.InternalRangeClient
	id        *networkconnectivityrefs.NetworkConnectivityInternalRangeIdentity
	desired   *pb.InternalRange
	actual    *pb.InternalRange
}

var _ directbase.Adapter = &internalRangeAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *internalRangeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networkconnectivity internalrange", "name", a.id)
	fqn := a.id.String()
	req := &pb.GetInternalRangeRequest{
		Name: fqn,
	}
	actual, err := a.gcpClient.GetInternalRange(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networkconnectivity internalrange %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *internalRangeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating networkconnectivity internalrange", "name", a.id)

	fqn := a.id.String()
	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)

	req := &pb.CreateInternalRangeRequest{
		Parent:          parent,
		InternalRangeId: a.id.InternalRange,
		InternalRange:   a.desired,
	}

	op, err := a.gcpClient.CreateInternalRange(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networkconnectivity internalrange %s: %w", fqn, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for create of networkconnectivity internalrange %q: %w", fqn, err)
	}

	log.V(2).Info("successfully created networkconnectivity internalrange in gcp", "name", a.id)

	resourceID := lastComponent(created.GetName())
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	mapCtx := &direct.MapContext{}
	status := &krm.NetworkConnectivityInternalRangeStatus{}
	status.ObservedState = NetworkConnectivityInternalRangeObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *internalRangeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networkconnectivity internalrange", "name", a.id)
	fqn := a.id.String()

	paths := []string{}

	if !reflect.DeepEqual(a.desired.GetDescription(), a.actual.GetDescription()) {
		paths = append(paths, "description")
	}
	if !reflect.DeepEqual(a.desired.GetLabels(), a.actual.GetLabels()) {
		paths = append(paths, "labels")
	}
	if !reflect.DeepEqual(a.desired.GetOverlaps(), a.actual.GetOverlaps()) {
		paths = append(paths, "overlaps")
	}
	if !reflect.DeepEqual(a.desired.GetTargetCidrRange(), a.actual.GetTargetCidrRange()) {
		paths = append(paths, "target_cidr_range")
	}
	if a.desired.GetPrefixLength() != a.actual.GetPrefixLength() {
		paths = append(paths, "prefix_length")
	}

	if len(paths) == 0 {
		mapCtx := &direct.MapContext{}
		status := &krm.NetworkConnectivityInternalRangeStatus{}
		status.ObservedState = NetworkConnectivityInternalRangeObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	a.desired.Name = fqn
	req := &pb.UpdateInternalRangeRequest{
		InternalRange: a.desired,
		UpdateMask:    &fieldmaskpb.FieldMask{Paths: paths},
	}

	op, err := a.gcpClient.UpdateInternalRange(ctx, req)
	if err != nil {
		return fmt.Errorf("updating networkconnectivity internalrange %s: %w", fqn, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of networkconnectivity internalrange %q: %w", fqn, err)
	}

	a.actual = updated

	mapCtx := &direct.MapContext{}
	status := &krm.NetworkConnectivityInternalRangeStatus{}
	status.ObservedState = NetworkConnectivityInternalRangeObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *internalRangeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	mapCtx := &direct.MapContext{}
	spec := NetworkConnectivityInternalRangeSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{}
	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.NetworkConnectivityInternalRangeGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *internalRangeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networkconnectivity internalrange", "name", a.id)
	fqn := a.id.String()
	req := &pb.DeleteInternalRangeRequest{
		Name: fqn,
	}
	op, err := a.gcpClient.DeleteInternalRange(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent networkconnectivity internalrange, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting networkconnectivity internalrange %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted networkconnectivity internalrange", "name", a.id)

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting delete networkconnectivity internalrange %s: %w", fqn, err)
	}
	return true, nil
}
