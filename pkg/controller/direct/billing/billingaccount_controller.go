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
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/billing/apiv1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BillingAccountGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.CloudBillingClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudBillingRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CloudBilling client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BillingAccount{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize ParentRef if specified
	if obj.Spec.ParentRef != nil {
		if err := obj.Spec.ParentRef.Normalize(ctx, reader, obj.Namespace); err != nil {
			return nil, err
		}
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.BillingAccountIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.BillingAccountIdentity
	gcpClient *gcp.CloudBillingClient
	desired   *krm.BillingAccount
	actual    *pb.BillingAccount
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BillingAccount", "name", a.id.String())

	req := &pb.GetBillingAccountRequest{Name: a.id.String()}
	billingAccount, err := a.gcpClient.GetBillingAccount(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BillingAccount %q: %w", a.id.String(), err)
	}

	a.actual = billingAccount
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating BillingAccount", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BillingAccountSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.String()

	req := &pb.CreateBillingAccountRequest{
		BillingAccount: resource,
	}
	if desired.Spec.ParentRef != nil {
		req.Parent = desired.Spec.ParentRef.External
	}

	created, err := a.gcpClient.CreateBillingAccount(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BillingAccount %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created BillingAccount", "name", a.id.String())

	status := &krm.BillingAccountStatus{}
	status.ObservedState = BillingAccountObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BillingAccount", "name", a.id.String())

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BillingAccountSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		report.AddField("display_name", a.actual.DisplayName, resource.DisplayName)
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Update only if mutable fields (displayName) actually changed.
	var updated *pb.BillingAccount
	var err error
	if len(updateMask.Paths) > 0 {
		req := &pb.UpdateBillingAccountRequest{
			Name:       a.id.String(),
			Account:    resource,
			UpdateMask: updateMask,
		}
		updated, err = a.gcpClient.UpdateBillingAccount(ctx, req)
		if err != nil {
			return fmt.Errorf("updating BillingAccount %s: %w", a.id.String(), err)
		}
	} else {
		updated = a.actual
	}

	// Move billing account if parent changed
	if resource.Parent != "" && resource.Parent != a.actual.Parent {
		log.V(2).Info("moving BillingAccount", "name", a.id.String(), "oldParent", a.actual.Parent, "newParent", resource.Parent)
		moveReq := &pb.MoveBillingAccountRequest{
			Name:              a.id.String(),
			DestinationParent: resource.Parent,
		}
		updated, err = a.gcpClient.MoveBillingAccount(ctx, moveReq)
		if err != nil {
			return fmt.Errorf("moving BillingAccount %s: %w", a.id.String(), err)
		}
	}

	log.V(2).Info("successfully updated BillingAccount", "name", a.id.String())

	status := &krm.BillingAccountStatus{}
	status.ObservedState = BillingAccountObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return setStatus(updateOp.GetUnstructured(), status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("No-op Delete for BillingAccount (not supported by GCP API)", "name", a.id.String())
	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
