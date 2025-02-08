// Copyright 2025 Google LLC
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

package paymentgateway

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/paymentgateway/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/paymentgateway/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	paymentgatewaypb "cloud.google.com/go/paymentgateway/v1/paymentgatewaypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.PaymentgatewayMetadataTransactionGVK, NewMetadataTransactionModel)
}

func NewMetadataTransactionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelMetadataTransaction{config: *config}, nil
}

var _ directbase.Model = &modelMetadataTransaction{}

type modelMetadataTransaction struct {
	config config.ControllerConfig
}

func (m *modelMetadataTransaction) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building MetadataTransaction client: %w", err)
	}
	return gcpClient, err
}

func (m *modelMetadataTransaction) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.PaymentgatewayMetadataTransaction{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewMetadataTransactionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get paymentgateway GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &MetadataTransactionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelMetadataTransaction) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type MetadataTransactionAdapter struct {
	id        *krm.MetadataTransactionIdentity
	gcpClient *gcp.Client
	desired   *krm.PaymentgatewayMetadataTransaction
	actual    *paymentgatewaypb.MetadataTransaction
}

var _ directbase.Adapter = &MetadataTransactionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *MetadataTransactionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting MetadataTransaction", "name", a.id)

	req := &paymentgatewaypb.GetMetadataTransactionRequest{Name: a.id}
	metadatatransactionpb, err := a.gcpClient.GetMetadataTransaction(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting MetadataTransaction %q: %w", a.id, err)
	}

	a.actual = metadatatransactionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetadataTransactionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating MetadataTransaction", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := PaymentgatewayMetadataTransactionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &paymentgatewaypb.CreateMetadataTransactionRequest{
		Parent:              a.id.Parent().String(),
		MetadataTransaction: resource,
	}
	op, err := a.gcpClient.CreateMetadataTransaction(ctx, req)
	if err != nil {
		return fmt.Errorf("creating MetadataTransaction %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MetadataTransaction %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created MetadataTransaction", "name", a.id)

	status := &krm.PaymentgatewayMetadataTransactionStatus{}
	status.ObservedState = PaymentgatewayMetadataTransactionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *MetadataTransactionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating MetadataTransaction", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := PaymentgatewayMetadataTransactionSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	// Option 1: This option is good for proto that has `field_mask` for output-only, immutable, required/optional.
	// TODO(contributor): If choosing this option, remove the "Option 2" code.
	{
		var err error
		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
	}

	// Option 2: manually add all mutable fields.
	// TODO(contributor): If choosing this option, remove the "Option 1" code.
	{
		if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
			paths = append(paths, "display_name")
		}
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		status := &krm.PaymentgatewayMetadataTransactionStatus{}
		status.ObservedState = PaymentgatewayMetadataTransactionObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &paymentgatewaypb.UpdateMetadataTransactionRequest{
		Name:                a.id.External,
		UpdateMask:          updateMask,
		MetadataTransaction: desiredPb,
	}
	op, err := a.gcpClient.UpdateMetadataTransaction(ctx, req)
	if err != nil {
		return fmt.Errorf("updating MetadataTransaction %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("MetadataTransaction %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated MetadataTransaction", "name", a.id.External)

	status := &krm.PaymentgatewayMetadataTransactionStatus{}
	status.ObservedState = PaymentgatewayMetadataTransactionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *MetadataTransactionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PaymentgatewayMetadataTransaction{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(PaymentgatewayMetadataTransactionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.PaymentgatewayMetadataTransactionGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *MetadataTransactionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting MetadataTransaction", "name", a.id)

	req := &paymentgatewaypb.DeleteMetadataTransactionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteMetadataTransaction(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting MetadataTransaction %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted MetadataTransaction", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete MetadataTransaction %s: %w", a.id, err)
	}
	return true, nil
}
