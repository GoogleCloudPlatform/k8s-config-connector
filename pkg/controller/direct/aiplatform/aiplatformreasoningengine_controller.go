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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformReasoningEngineGVK, NewAIPlatformReasoningEngineModel)
}

func NewAIPlatformReasoningEngineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &aiplatformReasoningEngineModel{config: config}, nil
}

var _ directbase.Model = &aiplatformReasoningEngineModel{}

type aiplatformReasoningEngineModel struct {
	config *config.ControllerConfig
}

func (m *aiplatformReasoningEngineModel) client(ctx context.Context, location string) (*gcp.ReasoningEngineClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewReasoningEngineClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReasoningEngineClient client: %w", err)
	}
	return gcpClient, nil
}

func (m *aiplatformReasoningEngineModel) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.AIPlatformReasoningEngine{}
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

	typedID, ok := id.(*krm.AIPlatformReasoningEngineIdentity)
	if !ok {
		return nil, fmt.Errorf("expected AIPlatformReasoningEngineIdentity, got %T", id)
	}

	gcpClient, err := m.client(ctx, typedID.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredpb := AIPlatformReasoningEngineSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, fmt.Errorf("mapping spec to proto: %w", mapCtx.Err())
	}

	return &AIPlatformReasoningEngineAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desiredpb: desiredpb,
		desired:   obj,
	}, nil
}

func (m *aiplatformReasoningEngineModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AIPlatformReasoningEngineIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &AIPlatformReasoningEngineAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type AIPlatformReasoningEngineAdapter struct {
	id        *krm.AIPlatformReasoningEngineIdentity
	gcpClient *gcp.ReasoningEngineClient
	desiredpb *pb.ReasoningEngine
	desired   *krm.AIPlatformReasoningEngine
	actual    *pb.ReasoningEngine
}

var _ directbase.Adapter = &AIPlatformReasoningEngineAdapter{}

func (a *AIPlatformReasoningEngineAdapter) Find(ctx context.Context) (bool, error) {
	// Rule 2: Find() Pre-Check Guard for service-generated ID
	if a.id.ReasoningEngine == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("getting AIPlatformReasoningEngine", "name", a.id.String())

	req := &pb.GetReasoningEngineRequest{
		Name: a.id.String(),
	}

	reasoningEnginepb, err := a.gcpClient.GetReasoningEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AIPlatformReasoningEngine %q: %w", a.id.String(), err)
	}

	a.actual = reasoningEnginepb

	return true, nil
}

func (a *AIPlatformReasoningEngineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AIPlatformReasoningEngine", "parent", a.id.ParentString())

	// Rule 3: Do not set desiredpb.Name, let GCP generate the identifier.
	a.desiredpb.Name = ""

	req := &pb.CreateReasoningEngineRequest{
		Parent:          a.id.ParentString(),
		ReasoningEngine: a.desiredpb,
	}

	op, err := a.gcpClient.CreateReasoningEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AIPlatformReasoningEngine in %q: %w", a.id.ParentString(), err)
	}

	log.V(2).Info("successfully started creation of AIPlatformReasoningEngine", "parent", a.id.ParentString())

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AIPlatformReasoningEngine creation in %q: %w", a.id.ParentString(), err)
	}

	if created != nil && created.Name != "" {
		tempID := &krm.AIPlatformReasoningEngineIdentity{}
		if err := tempID.FromExternal(created.Name); err != nil {
			return fmt.Errorf("parsing created AIPlatformReasoningEngine name %q: %w", created.Name, err)
		}
		a.id.ReasoningEngine = tempID.ReasoningEngine
	}

	log.V(2).Info("successfully completed creation of AIPlatformReasoningEngine", "name", a.id.String())

	// Fetch fully-populated resource immediately after LRO success
	getReq := &pb.GetReasoningEngineRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetReasoningEngine(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching newly created AIPlatformReasoningEngine %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *AIPlatformReasoningEngineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AIPlatformReasoningEngine", "name", a.id.String())

	diffs, updateMask, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformReasoningEngineSpec_FromProto,
		AIPlatformReasoningEngineSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	a.desiredpb.Name = a.id.String()
	req := &pb.UpdateReasoningEngineRequest{
		ReasoningEngine: a.desiredpb,
		UpdateMask:      updateMask,
	}

	op, err := a.gcpClient.UpdateReasoningEngine(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AIPlatformReasoningEngine %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully started update of AIPlatformReasoningEngine", "name", a.id.String())

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AIPlatformReasoningEngine %q update: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully completed update of AIPlatformReasoningEngine", "name", a.id.String())

	getReq := &pb.GetReasoningEngineRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetReasoningEngine(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching updated AIPlatformReasoningEngine %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *AIPlatformReasoningEngineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformReasoningEngine{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformReasoningEngineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.ReasoningEngine

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.ReasoningEngine)
	u.SetGroupVersionKind(krm.AIPlatformReasoningEngineGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *AIPlatformReasoningEngineAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ReasoningEngine) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AIPlatformReasoningEngineStatus{}
	status.ObservedState = AIPlatformReasoningEngineObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return fmt.Errorf("mapping status: %w", mapCtx.Err())
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *AIPlatformReasoningEngineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AIPlatformReasoningEngine", "name", a.id.String())

	req := &pb.DeleteReasoningEngineRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteReasoningEngine(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AIPlatformReasoningEngine %q: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of AIPlatformReasoningEngine %q: %w", a.id.String(), err)
	}

	return true, nil
}
