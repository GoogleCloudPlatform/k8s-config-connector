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

package documentai

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/documentai/apiv1"
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DocumentAIProcessorGVK, NewProcessorModel)
}

func NewProcessorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelProcessor{config: *config}, nil
}

var _ directbase.Model = &modelProcessor{}

type modelProcessor struct {
	config config.ControllerConfig
}

func (m *modelProcessor) client(ctx context.Context) (*gcp.DocumentProcessorClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDocumentProcessorRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Processor client: %w", err)
	}
	return gcpClient, err
}

func (m *modelProcessor) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.DocumentAIProcessor{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewProcessorIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get documentai GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ProcessorAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelProcessor) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ProcessorAdapter struct {
	id        *krm.ProcessorIdentity
	gcpClient *gcp.DocumentProcessorClient
	desired   *krm.DocumentAIProcessor
	actual    *pb.Processor
}

var _ directbase.Adapter = &ProcessorAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ProcessorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Processor", "name", a.id)

	req := &pb.GetProcessorRequest{Name: a.id.String()}
	processorpb, err := a.gcpClient.GetProcessor(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Processor %q: %w", a.id, err)
	}

	a.actual = processorpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ProcessorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Processor", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DocumentAIProcessorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateProcessorRequest{
		Parent:    a.id.Parent().String(),
		Processor: resource,
	}
	created, err := a.gcpClient.CreateProcessor(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Processor %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Processor", "name", a.id)

	status := &krm.DocumentAIProcessorStatus{}
	status.ObservedState = DocumentAIProcessorObservedState_FromProto(mapCtx, created)
	// TODO: is the default version also an output field?
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ProcessorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Processor", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := DocumentAIProcessorSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// NOTE: Processor is immutable, so we cannot update it.
	// However it is possible to update DefaultProcessorVersion via calling setdefaultprocessorversion API.
	if desiredPb.DisplayName != a.actual.DisplayName {
		return fmt.Errorf("displayName cannot be updated")
	}
	if desiredPb.Type != a.actual.Type {
		return fmt.Errorf("type cannot be updated")
	}
	if desiredPb.KmsKeyName != a.actual.KmsKeyName {
		return fmt.Errorf("KmsKeyName cannot be updated")
	}

	// ONLY allow updating DefaultProcessorVersion
	if desiredPb.DefaultProcessorVersion == a.actual.DefaultProcessorVersion {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.DocumentAIProcessorStatus{}
		status.ObservedState = DocumentAIProcessorObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.SetDefaultProcessorVersionRequest{
		Processor:               a.id.String(),
		DefaultProcessorVersion: desiredPb.DefaultProcessorVersion,
	}
	op, err := a.gcpClient.SetDefaultProcessorVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Processor %s: %w", a.id.String(), err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting update for Processor %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Processor", "name", a.id.String())

	// the SetDefaultProcessorVersion API does not return the updated Processor, so we need to read it again.
	updated, err := a.gcpClient.GetProcessor(ctx, &pb.GetProcessorRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting updated Processor %s: %w", a.id.String(), err)
	}

	status := &krm.DocumentAIProcessorStatus{}
	status.ObservedState = DocumentAIProcessorObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ProcessorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DocumentAIProcessor{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DocumentAIProcessorSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.DocumentAIProcessorGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ProcessorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Processor", "name", a.id)

	req := &pb.DeleteProcessorRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteProcessor(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Processor %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Processor %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Processor", "name", a.id)
	return true, nil
}
