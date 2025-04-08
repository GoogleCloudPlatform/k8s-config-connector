// Copyright 2024 Google LLC
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
// proto.service: google.cloud.speech.v2.Speech
// proto.message: google.cloud.speech.v2.CustomClass
// crd.type: SpeechCustomClass
// crd.version: v1alpha1

package speech

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/speech/apiv2"
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.SpeechCustomClassGVK, NewCustomClassModel)
}

func NewCustomClassModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &customClassModel{config: *config}, nil
}

var _ directbase.Model = &customClassModel{}

type customClassModel struct {
	config config.ControllerConfig
}

func (m *customClassModel) client(ctx context.Context, projectID string) (*gcp.Client, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building speech customclass client: %w", err)
	}

	return gcpClient, err
}

func (m *customClassModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SpeechCustomClass{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewCustomClassIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &customClassAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *customClassModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type customClassAdapter struct {
	gcpClient *gcp.Client
	id        *krm.CustomClassIdentity
	desired   *krm.SpeechCustomClass
	actual    *pb.CustomClass
	reader    client.Reader
}

var _ directbase.Adapter = &customClassAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *customClassAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting speech customclass", "name", a.id)

	req := &pb.GetCustomClassRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetCustomClass(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting speech customclass %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *customClassAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating speech customclass", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := SpeechCustomClassSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateCustomClassRequest{
		Parent:        a.id.Parent().String(),
		CustomClassId: a.id.ID(),
		CustomClass:   resource,
	}
	op, err := a.gcpClient.CreateCustomClass(ctx, req)
	if err != nil {
		return fmt.Errorf("creating speech customclass %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("speech customclass %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created speech customclass", "name", a.id)

	status := &krm.SpeechCustomClassStatus{}
	status.ObservedState = SpeechCustomClassObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *customClassAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating speech customclass", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := SpeechCustomClassSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if a.desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		paths = append(paths, "display_name")
	}
	if !reflect.DeepEqual(resource.Items, a.actual.Items) {
		paths = append(paths, "items")
	}
	if !reflect.DeepEqual(resource.Annotations, a.actual.Annotations) {
		paths = append(paths, "annotations")
	}

	var updated *pb.CustomClass
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateCustomClassRequest{
			CustomClass: resource,
			UpdateMask:  &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateCustomClass(ctx, req)
		if err != nil {
			return fmt.Errorf("updating speech customclass %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("speech customclass %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated speech customclass", "name", a.id)
	}

	status := &krm.SpeechCustomClassStatus{}
	status.ObservedState = SpeechCustomClassObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *customClassAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SpeechCustomClass{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SpeechCustomClassSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.SpeechCustomClassGVK)
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *customClassAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting speech customclass", "name", a.id)

	req := &pb.DeleteCustomClassRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCustomClass(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent speech customclass, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting speech customclass %s: %w", a.id, err)
	}
	log.V(2).Info("successfully initiated deletion speech customclass", "name", a.id)

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete speech customclass %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted speech customclass", "name", a.id)
	return true, nil
}
