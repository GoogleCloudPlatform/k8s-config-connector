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

package aiplatform

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformSpecialistPoolGVK, NewAIPlatformSpecialistPoolModel)
}

func NewAIPlatformSpecialistPoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &aiplatformSpecialistPoolModel{config: config}, nil
}

var _ directbase.Model = &aiplatformSpecialistPoolModel{}

type aiplatformSpecialistPoolModel struct {
	config *config.ControllerConfig
}

func (m *aiplatformSpecialistPoolModel) client(ctx context.Context, location string) (*gcp.SpecialistPoolClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewSpecialistPoolClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SpecialistPoolClient client: %w", err)
	}
	return gcpClient, nil
}

func (m *aiplatformSpecialistPoolModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.AIPlatformSpecialistPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(reader.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader.Reader)
	if err != nil {
		return nil, err
	}

	// Always call common.NormalizeReferences to resolve any resource references
	if err := common.NormalizeReferences(ctx, reader.Reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	typedID, ok := id.(*krm.AIPlatformSpecialistPoolIdentity)
	if !ok {
		return nil, fmt.Errorf("expected AIPlatformSpecialistPoolIdentity, got %T", id)
	}

	gcpClient, err := m.client(ctx, typedID.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := AIPlatformSpecialistPoolSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	return &AIPlatformSpecialistPoolAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desiredpb: desiredpb,
		desired:   obj,
	}, nil
}

func (m *aiplatformSpecialistPoolModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AIPlatformSpecialistPoolIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &AIPlatformSpecialistPoolAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type AIPlatformSpecialistPoolAdapter struct {
	id        *krm.AIPlatformSpecialistPoolIdentity
	gcpClient *gcp.SpecialistPoolClient
	desiredpb *pb.SpecialistPool
	desired   *krm.AIPlatformSpecialistPool
	actual    *pb.SpecialistPool
}

var _ directbase.Adapter = &AIPlatformSpecialistPoolAdapter{}

func (a *AIPlatformSpecialistPoolAdapter) Find(ctx context.Context) (bool, error) {
	// Rule 2: Find() Pre-Check Guard for service-generated ID
	if a.id.SpecialistPool == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("getting AIPlatformSpecialistPool", "name", a.id.String())

	req := &pb.GetSpecialistPoolRequest{
		Name: a.id.String(),
	}

	specialistPoolpb, err := a.gcpClient.GetSpecialistPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AIPlatformSpecialistPool %q: %w", a.id.String(), err)
	}

	a.actual = specialistPoolpb

	mapCtx := &direct.MapContext{}
	observedState := AIPlatformSpecialistPoolObservedState_FromProto(mapCtx, specialistPoolpb)
	if mapCtx.Err() != nil {
		return false, fmt.Errorf("mapping from proto to observed state: %w", mapCtx.Err())
	}

	if a.desired != nil {
		a.desired.Status.ObservedState = observedState
		a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	}
	return true, nil
}

func (a *AIPlatformSpecialistPoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AIPlatformSpecialistPool", "parent", a.id.ParentString())

	// Rule 3: Do not set desiredpb.Name, let GCP generate the identifier.
	a.desiredpb.Name = ""

	req := &pb.CreateSpecialistPoolRequest{
		Parent:         a.id.ParentString(),
		SpecialistPool: a.desiredpb,
	}

	op, err := a.gcpClient.CreateSpecialistPool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AIPlatformSpecialistPool in %q: %w", a.id.ParentString(), err)
	}

	log.V(2).Info("successfully started creation of AIPlatformSpecialistPool", "parent", a.id.ParentString())

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AIPlatformSpecialistPool creation in %q: %w", a.id.ParentString(), err)
	}

	if created != nil && created.Name != "" {
		tempID := &krm.AIPlatformSpecialistPoolIdentity{}
		if err := tempID.FromExternal(created.Name); err != nil {
			return fmt.Errorf("parsing created AIPlatformSpecialistPool name %q: %w", created.Name, err)
		}
		a.id.SpecialistPool = tempID.SpecialistPool
	}

	log.V(2).Info("successfully completed creation of AIPlatformSpecialistPool", "name", a.id.String())

	// Fetch fully-populated resource immediately after LRO success
	getReq := &pb.GetSpecialistPoolRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetSpecialistPool(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching newly created AIPlatformSpecialistPool %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *AIPlatformSpecialistPoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AIPlatformSpecialistPool", "name", a.id.String())

	a.desiredpb.Name = a.id.String()

	diffs, updateMask, err := compareSpecialistPool(ctx, a.actual, a.desiredpb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSpecialistPoolRequest{
		SpecialistPool: a.desiredpb,
		UpdateMask:     updateMask,
	}

	op, err := a.gcpClient.UpdateSpecialistPool(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AIPlatformSpecialistPool %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully started update of AIPlatformSpecialistPool", "name", a.id.String())

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AIPlatformSpecialistPool %q update: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully completed update of AIPlatformSpecialistPool", "name", a.id.String())

	getReq := &pb.GetSpecialistPoolRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetSpecialistPool(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching updated AIPlatformSpecialistPool %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *AIPlatformSpecialistPoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformSpecialistPool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformSpecialistPoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.SpecialistPool

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.SpecialistPool)
	u.SetGroupVersionKind(krm.AIPlatformSpecialistPoolGVK)

	return u, nil
}

func (a *AIPlatformSpecialistPoolAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SpecialistPool) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AIPlatformSpecialistPoolStatus{}
	status.ObservedState = AIPlatformSpecialistPoolObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *AIPlatformSpecialistPoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AIPlatformSpecialistPool", "name", a.id.String())

	req := &pb.DeleteSpecialistPoolRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteSpecialistPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AIPlatformSpecialistPool %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for deletion of AIPlatformSpecialistPool %q: %w", a.id.String(), err)
	}

	return true, nil
}

func compareSpecialistPool(ctx context.Context, actual, desired *pb.SpecialistPool) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, AIPlatformSpecialistPoolSpec_FromProto, AIPlatformSpecialistPoolSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}

	return diffs, updateMask, nil
}
