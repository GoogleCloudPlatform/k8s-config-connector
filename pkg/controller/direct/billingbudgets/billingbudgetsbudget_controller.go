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

package billingbudgets

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/billing/budgets/apiv1"
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billingbudgets/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BillingBudgetsBudgetGVK, NewBillingBudgetsBudgetModel)
}

func NewBillingBudgetsBudgetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBillingBudgetsBudget{config: *config}, nil
}

var _ directbase.Model = &modelBillingBudgetsBudget{}

type modelBillingBudgetsBudget struct {
	config config.ControllerConfig
}

func (m *modelBillingBudgetsBudget) client(ctx context.Context) (*gcp.BudgetClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewBudgetRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BillingBudgets budget client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBillingBudgetsBudget) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BillingBudgetsBudget{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := BillingBudgetsBudgetSpec_v1beta1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredPb.Name = id.String()

	return &BillingBudgetsBudgetAdapter{
		id:        id.(*krm.BillingBudgetsBudgetIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelBillingBudgetsBudget) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type BillingBudgetsBudgetAdapter struct {
	id        *krm.BillingBudgetsBudgetIdentity
	gcpClient *gcp.BudgetClient
	desired   *pb.Budget
	actual    *pb.Budget
}

var _ directbase.Adapter = &BillingBudgetsBudgetAdapter{}

func (a *BillingBudgetsBudgetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BillingBudgetsBudget", "name", a.id)

	req := &pb.GetBudgetRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetBudget(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BillingBudgetsBudget %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *BillingBudgetsBudgetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating BillingBudgetsBudget", "id", fqn)

	parent := "billingAccounts/" + a.id.BillingAccount

	req := &pb.CreateBudgetRequest{
		Parent: parent,
		Budget: a.desired,
	}
	created, err := a.gcpClient.CreateBudget(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BillingBudgetsBudget %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BillingBudgetsBudget", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *BillingBudgetsBudgetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BillingBudgetsBudget", "name", a.id.String())

	diffs, updateMask, err := compareBillingBudgetsBudget(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateBudgetRequest{
			Budget:     a.desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateBudget(ctx, req)
		if err != nil {
			return fmt.Errorf("updating BillingBudgetsBudget %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *BillingBudgetsBudgetAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Budget) error {
	mapCtx := &direct.MapContext{}
	status := BillingBudgetsBudgetStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *BillingBudgetsBudgetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BillingBudgetsBudget{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BillingBudgetsBudgetSpec_v1beta1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.BillingBudgetsBudgetGVK)
	return u, nil
}

func (a *BillingBudgetsBudgetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BillingBudgetsBudget", "name", a.id)

	req := &pb.DeleteBudgetRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteBudget(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent BillingBudgetsBudget, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting BillingBudgetsBudget %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BillingBudgetsBudget", "name", a.id)
	return true, nil
}

func compareBillingBudgetsBudget(ctx context.Context, actual, desired *pb.Budget) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BillingBudgetsBudgetSpec_v1beta1_FromProto, BillingBudgetsBudgetSpec_v1beta1_ToProto)
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
