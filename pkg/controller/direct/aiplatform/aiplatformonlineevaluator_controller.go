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

// +generated:controller
// krm.group: aiplatform.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AIPlatformOnlineEvaluator

package aiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	"google.golang.org/protobuf/proto"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/pb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformOnlineEvaluatorGVK, NewOnlineEvaluatorModel)
}

func NewOnlineEvaluatorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelOnlineEvaluator{config: *config}, nil
}

var _ directbase.Model = &modelOnlineEvaluator{}

type modelOnlineEvaluator struct {
	config config.ControllerConfig
}

func (m *modelOnlineEvaluator) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformOnlineEvaluator{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AIPlatformOnlineEvaluatorIdentity)

	gcpClient, operationsClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &OnlineEvaluatorAdapter{
		id:               id,
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		desired:          obj,
	}, nil
}

func (m *modelOnlineEvaluator) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.AIPlatformOnlineEvaluatorIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, operationsClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &OnlineEvaluatorAdapter{
		id:               id,
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
	}, nil
}

func (m *modelOnlineEvaluator) client(ctx context.Context, location string) (pb.OnlineEvaluatorServiceClient, longrunningpb.OperationsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}
	endpoint := "aiplatform.googleapis.com:443"
	if location != "" && location != "global" {
		endpoint = fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	}
	opts = append(opts, option.WithEndpoint(endpoint))
	opts = append(opts, option.WithScopes("https://www.googleapis.com/auth/cloud-platform"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing OnlineEvaluator service: %w", err)
	}
	return pb.NewOnlineEvaluatorServiceClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

type OnlineEvaluatorAdapter struct {
	id               *krm.AIPlatformOnlineEvaluatorIdentity
	gcpClient        pb.OnlineEvaluatorServiceClient
	operationsClient longrunningpb.OperationsClient
	desired          *krm.AIPlatformOnlineEvaluator
	actual           *pb.OnlineEvaluator
}

var _ directbase.Adapter = &OnlineEvaluatorAdapter{}

// Find retrieves the GCP resource.
func (a *OnlineEvaluatorAdapter) Find(ctx context.Context) (bool, error) {
	// Rule 2: Find() Pre-Check Guard for service-generated ID
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("getting OnlineEvaluator", "name", a.id.String())

	req := &pb.GetOnlineEvaluatorRequest{
		Name: a.id.String(),
	}
	onlineevaluatorpb, err := a.gcpClient.GetOnlineEvaluator(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting OnlineEvaluator %q: %w", a.id.String(), err)
	}

	a.actual = onlineevaluatorpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *OnlineEvaluatorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating OnlineEvaluator", "parent", a.id.ParentString())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformOnlineEvaluatorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Rule 3: Do not set resource.Name, let GCP generate the identifier.
	resource.Name = ""

	req := &pb.CreateOnlineEvaluatorRequest{
		Parent:          a.id.ParentString(),
		OnlineEvaluator: resource,
	}
	op, err := a.gcpClient.CreateOnlineEvaluator(ctx, req)
	if err != nil {
		return fmt.Errorf("creating OnlineEvaluator in %s: %w", a.id.ParentString(), err)
	}
	completedOp, err := a.waitForLRO(ctx, op)
	if err != nil {
		return fmt.Errorf("OnlineEvaluator in %s waiting creation: %w", a.id.ParentString(), err)
	}

	if completedOp.GetResponse() != nil {
		var created pb.OnlineEvaluator
		// Note: The protobuf package for OnlineEvaluator is compiled under
		// google.cloud.aiplatform.v1beta1.onlineevaluator to avoid GlobalTypes/GlobalFiles
		// collisions with GAPIC apiv1beta1/aiplatformpb. This package rename changes
		// the type URL of the google.protobuf.Any message in the LRO response from
		// type.googleapis.com/google.cloud.aiplatform.v1beta1.OnlineEvaluator to
		// type.googleapis.com/google.cloud.aiplatform.v1beta1.onlineevaluator.OnlineEvaluator.
		// As a result, anypb.UnmarshalTo fails on type URL mismatch; we decode the raw
		// wire bytes using proto.Unmarshal instead.
		if err := proto.Unmarshal(completedOp.GetResponse().GetValue(), &created); err != nil {
			return fmt.Errorf("unmarshaling OnlineEvaluator from LRO response: %w", err)
		}
		if created.GetName() != "" {
			tempID := &krm.AIPlatformOnlineEvaluatorIdentity{}
			if err := tempID.FromExternal(created.GetName()); err != nil {
				return fmt.Errorf("parsing created OnlineEvaluator name %q: %w", created.GetName(), err)
			}
			a.id.OnlineEvaluator = tempID.OnlineEvaluator
		}
	}
	if !a.id.HasIdentitySpecified() {
		return fmt.Errorf("created OnlineEvaluator ID is empty")
	}
	log.V(2).Info("successfully created OnlineEvaluator", "name", a.id.String())

	getReq := &pb.GetOnlineEvaluatorRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetOnlineEvaluator(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching newly created OnlineEvaluator %q: %w", a.id.String(), err)
	}

	return a.setStatus(ctx, createOp, latest)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *OnlineEvaluatorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating OnlineEvaluator", "name", a.id.String())

	// Real GCP omits semconvVersion from GET responses (returns "openTelemetry": {}).
	// We normalize semconvVersion to empty string during comparison so that
	// server-omitted default fields do not trigger spurious diffs on re-reconciliation.
	normalize := func(ctx context.Context, p *pb.OnlineEvaluator) error {
		if co := p.GetCloudObservability(); co != nil && co.GetOpenTelemetry() != nil {
			co.GetOpenTelemetry().SemconvVersion = ""
		}
		return nil
	}

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformOnlineEvaluatorSpec_FromProto,
		AIPlatformOnlineEvaluatorSpec_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.setStatus(ctx, updateOp, a.actual)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)
	resource.Name = a.id.String()
	req := &pb.UpdateOnlineEvaluatorRequest{
		UpdateMask:      updateMask,
		OnlineEvaluator: resource,
	}
	op, err := a.gcpClient.UpdateOnlineEvaluator(ctx, req)
	if err != nil {
		return fmt.Errorf("updating OnlineEvaluator %s: %w", a.id.String(), err)
	}
	_, err = a.waitForLRO(ctx, op)
	if err != nil {
		return fmt.Errorf("OnlineEvaluator %s waiting update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated OnlineEvaluator", "name", a.id.String())

	getReq := &pb.GetOnlineEvaluatorRequest{
		Name: a.id.String(),
	}
	latest, err := a.gcpClient.GetOnlineEvaluator(ctx, getReq)
	if err != nil {
		return fmt.Errorf("fetching updated OnlineEvaluator %q: %w", a.id.String(), err)
	}

	return a.setStatus(ctx, updateOp, latest)
}

func (a *OnlineEvaluatorAdapter) setStatus(ctx context.Context, op directbase.Operation, latest *pb.OnlineEvaluator) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AIPlatformOnlineEvaluatorStatus{}
	if latest != nil {
		status.ObservedState = AIPlatformOnlineEvaluatorObservedState_FromProto(mapCtx, latest)
		if mapCtx.Err() != nil {
			return fmt.Errorf("mapping status: %w", mapCtx.Err())
		}
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *OnlineEvaluatorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformOnlineEvaluator{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformOnlineEvaluatorSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AIPlatformOnlineEvaluatorGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *OnlineEvaluatorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting OnlineEvaluator", "name", a.id.String())

	req := &pb.DeleteOnlineEvaluatorRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteOnlineEvaluator(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent OnlineEvaluator, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting OnlineEvaluator %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted OnlineEvaluator", "name", a.id.String())
	_, err = a.waitForLRO(ctx, op)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete OnlineEvaluator %s: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *OnlineEvaluatorAdapter) waitForLRO(ctx context.Context, op *longrunningpb.Operation) (*longrunningpb.Operation, error) {
	pollInterval := 1 * time.Second
	return common.WaitForOperation(ctx, pollInterval, func(op *longrunningpb.Operation) (bool, error) {
		if !op.Done {
			return false, nil
		}
		if errProto := op.GetError(); errProto != nil {
			return true, fmt.Errorf("operation %s failed: %s (code %d)", op.Name, errProto.GetMessage(), errProto.GetCode())
		}
		return true, nil
	}, func() (*longrunningpb.Operation, error) {
		return a.operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
	})
}
