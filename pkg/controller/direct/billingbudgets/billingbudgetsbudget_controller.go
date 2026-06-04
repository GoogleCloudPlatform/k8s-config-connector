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
	"reflect"

	gcp "cloud.google.com/go/billing/budgets/apiv1"
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billingbudgets/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.BillingBudgetsBudgetGVK, newModel)
}

func newModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type adapter struct {
	id      *krm.BillingBudgetsBudgetIdentity
	desired *krm.BillingBudgetsBudget
	actual  *pb.Budget
	gcp     *gcp.BudgetClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*gcp.BudgetClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewBudgetRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building billing budgets client: %w", err)
	}
	return gcpClient, nil
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.BillingBudgetsBudget{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	return &adapter{
		id:      id.(*krm.BillingBudgetsBudgetIdentity),
		desired: obj,
		gcp:     gcpClient,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	if a.id.Budget == "" {
		return false, nil
	}

	req := &pb.GetBudgetRequest{
		Name: a.fullyQualifiedName(),
	}
	budget, err := a.gcp.GetBudget(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = budget
	return true, nil
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &pb.DeleteBudgetRequest{
		Name: a.fullyQualifiedName(),
	}
	err := a.gcp.DeleteBudget(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting budget: %w", err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating budget", "u", u)

	mapCtx := &direct.MapContext{}
	budgetProto := BillingBudgetsBudgetSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateBudgetRequest{
		Parent: fmt.Sprintf("billingAccounts/%s", a.id.BillingAccount),
		Budget: budgetProto,
	}

	created, err := a.gcp.CreateBudget(ctx, req)
	if err != nil {
		return fmt.Errorf("creating budget: %w", err)
	}

	log.V(2).Info("created budget", "budget", created)

	status := BillingBudgetsBudgetStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	statusObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(status)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	if err := unstructured.SetNestedField(u.Object, statusObj, "status"); err != nil {
		return fmt.Errorf("setting status: %w", err)
	}

	return nil
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	log := klog.FromContext(ctx)

	mapCtx := &direct.MapContext{}
	desiredProto := BillingBudgetsBudgetSpec_v1beta1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := &fieldmaskpb.FieldMask{}

	if !reflect.DeepEqual(desiredProto.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(desiredProto.BudgetFilter, a.actual.BudgetFilter) {
		updateMask.Paths = append(updateMask.Paths, "budget_filter")
	}
	if !reflect.DeepEqual(desiredProto.Amount, a.actual.Amount) {
		updateMask.Paths = append(updateMask.Paths, "amount")
	}
	if !reflect.DeepEqual(desiredProto.ThresholdRules, a.actual.ThresholdRules) {
		updateMask.Paths = append(updateMask.Paths, "threshold_rules")
	}
	if !reflect.DeepEqual(desiredProto.NotificationsRule, a.actual.NotificationsRule) {
		updateMask.Paths = append(updateMask.Paths, "notifications_rule")
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for _, path := range updateMask.Paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	desiredProto.Name = a.fullyQualifiedName()

	req := &pb.UpdateBudgetRequest{
		Budget:     desiredProto,
		UpdateMask: updateMask,
	}

	updated, err := a.gcp.UpdateBudget(ctx, req)
	if err != nil {
		return err
	}

	log.V(2).Info("updated budget", "budget", updated)

	status := BillingBudgetsBudgetStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	statusObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(status)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	if err := unstructured.SetNestedField(u.Object, statusObj, "status"); err != nil {
		return fmt.Errorf("setting status: %w", err)
	}

	return nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("billingbudgets budget %q not found", a.fullyQualifiedName())
	}

	mapCtx := &direct.MapContext{}
	spec := BillingBudgetsBudgetSpec_v1beta1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.BillingAccountRef = a.desired.Spec.BillingAccountRef

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting billingbudgets budget spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.Budget)
	u.SetGroupVersionKind(krm.BillingBudgetsBudgetGVK)
	u.SetLabels(a.desired.Labels)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *adapter) fullyQualifiedName() string {
	return fmt.Sprintf("billingAccounts/%s/budgets/%s", a.id.BillingAccount, a.id.Budget)
}
