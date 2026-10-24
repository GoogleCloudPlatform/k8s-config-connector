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

// +generated:controller
// krm.group: dialogflow.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DialogflowConversationModel

package dialogflow

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowConversationModelGVK, NewConversationModelModel)
}

func NewConversationModelModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConversationModel{config: *config}, nil
}

var _ directbase.Model = &modelConversationModel{}

type modelConversationModel struct {
	config config.ControllerConfig
}

func (m *modelConversationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowConversationModel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DialogflowConversationModelIdentity)

	// Get dialogflow GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConversationModelAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelConversationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowConversationModelIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConversationModelAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelConversationModel) client(ctx context.Context) (*gcp.ConversationModelsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewConversationModelsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ConversationModel client: %w", err)
	}
	return gcpClient, err
}

type ConversationModelAdapter struct {
	id        *krm.DialogflowConversationModelIdentity
	gcpClient *gcp.ConversationModelsClient
	desired   *krm.DialogflowConversationModel
	actual    *dialogflowpb.ConversationModel
}

var _ directbase.Adapter = &ConversationModelAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ConversationModelAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ConversationModel", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &dialogflowpb.GetConversationModelRequest{
		Name: a.id.String(),
	}
	conversationmodelpb, err := a.gcpClient.GetConversationModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConversationModel %q: %w", a.id, err)
	}

	a.actual = conversationmodelpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConversationModelAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ConversationModel", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DialogflowConversationModelSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dialogflowpb.CreateConversationModelRequest{
		Parent:            a.id.ParentString(),
		ConversationModel: resource,
	}
	op, err := a.gcpClient.CreateConversationModel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ConversationModel %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ConversationModel %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ConversationModel", "name", a.id)

	status := &krm.DialogflowConversationModelStatus{}
	status.ObservedState = DialogflowConversationModelObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.DialogflowConversationModelIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.ConversationModel = parsedID.ConversationModel
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConversationModelAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ConversationModel", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DialogflowConversationModelSpec_FromProto,
		DialogflowConversationModelSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating DialogflowConversationModel is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.DialogflowConversationModelStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DialogflowConversationModelObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConversationModelAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowConversationModel{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowConversationModelSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DialogflowConversationModelGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ConversationModelAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ConversationModel", "name", a.id)

	req := &dialogflowpb.DeleteConversationModelRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteConversationModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ConversationModel, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ConversationModel %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ConversationModel", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete ConversationModel %s: %w", a.id, err)
	}
	return true, nil
}
