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

package apigateway

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/apigateway/apiv1"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	apigatewaypb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.APIGatewayAPIGVK, NewApiModel)
}

func NewApiModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelApi{config: *config}, nil
}

var _ directbase.Model = &modelApi{}

type modelApi struct {
	config config.ControllerConfig
}

func (m *modelApi) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Api client: %w", err)
	}
	return gcpClient, err
}

func (m *modelApi) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.APIGatewayAPI{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewApiIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get apigateway GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ApiAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelApi) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ApiAdapter struct {
	id        *krm.ApiIdentity
	gcpClient *gcp.Client
	desired   *krm.APIGatewayAPI
	actual    *apigatewaypb.Api
}

var _ directbase.Adapter = &ApiAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ApiAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Api", "name", a.id)

	req := &apigatewaypb.GetApiRequest{Name: a.id.String()}
	apipb, err := a.gcpClient.GetApi(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Api %q: %w", a.id, err)
	}

	a.actual = apipb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApiAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Api", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := APIGatewayAPISpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &apigatewaypb.CreateApiRequest{
		Parent: a.id.Parent().String(),
		Api:    resource,
	}
	op, err := a.gcpClient.CreateApi(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Api %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Api %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Api", "name", a.id)

	status := &krm.APIGatewayAPIStatus{}
	status.ObservedState = APIGatewayAPIObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ApiAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Api", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := APIGatewayAPISpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	var paths sets.Set[string]
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.APIGatewayAPIStatus{}
		status.ObservedState = APIGatewayAPIObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &apigatewaypb.UpdateApiRequest{
		UpdateMask: updateMask,
		Api:        desiredPb,
	}
	op, err := a.gcpClient.UpdateApi(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Api %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Api %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Api", "name", a.id)

	status := &krm.APIGatewayAPIStatus{}
	status.ObservedState = APIGatewayAPIObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ApiAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIGatewayAPI{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(APIGatewayAPISpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.APIGatewayAPIGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ApiAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Api", "name", a.id)

	req := &apigatewaypb.DeleteApiRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteApi(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Api, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Api %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Api", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Api %s: %w", a.id, err)
	}
	return true, nil
}
