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

package billing

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/billing/apiv1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BillingAccountGVK, NewBillingAccountModel, registry.CannotBeDeleted())
}

func NewBillingAccountModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBillingAccount{config: *config}, nil
}

var _ directbase.Model = &modelBillingAccount{}

type modelBillingAccount struct {
	config config.ControllerConfig
}

func (m *modelBillingAccount) client(ctx context.Context) (*gcp.CloudBillingClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudBillingRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CloudBilling REST client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBillingAccount) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BillingAccount{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.BillingAccountIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := BillingAccountSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &BillingAccountAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelBillingAccount) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type BillingAccountAdapter struct {
	id        *krm.BillingAccountIdentity
	gcpClient *gcp.CloudBillingClient
	desired   *pb.BillingAccount
	actual    *pb.BillingAccount
}

var _ directbase.Adapter = &BillingAccountAdapter{}

func (a *BillingAccountAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BillingAccount", "name", a.id.String())

	req := &pb.GetBillingAccountRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetBillingAccount(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BillingAccount %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *BillingAccountAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating BillingAccount", "id", fqn)

	desired := proto.Clone(a.desired).(*pb.BillingAccount)
	desired.Name = fqn

	req := &pb.CreateBillingAccountRequest{
		BillingAccount: desired,
	}
	created, err := a.gcpClient.CreateBillingAccount(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BillingAccount %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BillingAccount", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *BillingAccountAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BillingAccount", "name", a.id.String())

	diffs, updateMask, err := compareBillingAccount(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.BillingAccount)
		desired.Name = a.id.String()

		req := &pb.UpdateBillingAccountRequest{
			Name:       a.id.String(),
			Account:    desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateBillingAccount(ctx, req)
		if err != nil {
			return fmt.Errorf("updating BillingAccount %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *BillingAccountAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.BillingAccount) error {
	mapCtx := &direct.MapContext{}
	status := &krm.BillingAccountStatus{}
	status.ObservedState = BillingAccountObservedState_FromProto(mapCtx, latest)
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *BillingAccountAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BillingAccount{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BillingAccountSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.BillingAccountGVK)
	return u, nil
}

func (a *BillingAccountAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BillingAccount is a no-op because billing accounts cannot be deleted in GCP", "name", a.id)
	return true, nil
}

func compareBillingAccount(ctx context.Context, actual, desired *pb.BillingAccount) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BillingAccountSpec_FromProto, BillingAccountSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
