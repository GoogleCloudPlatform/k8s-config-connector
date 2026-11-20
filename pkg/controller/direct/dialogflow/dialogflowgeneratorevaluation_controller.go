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
// krm.group: dialogflow.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DialogflowGeneratorEvaluation

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
	registry.RegisterModel(krm.DialogflowGeneratorEvaluationGVK, NewGeneratorEvaluationModel)
}

func NewGeneratorEvaluationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGeneratorEvaluation{config: *config}, nil
}

var _ directbase.Model = &modelGeneratorEvaluation{}

type modelGeneratorEvaluation struct {
	config config.ControllerConfig
}

func (m *modelGeneratorEvaluation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowGeneratorEvaluation{}
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
	id := i.(*krm.DialogflowGeneratorEvaluationIdentity)

	// Get dialogflow GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &GeneratorEvaluationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGeneratorEvaluation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowGeneratorEvaluationIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &GeneratorEvaluationAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelGeneratorEvaluation) client(ctx context.Context, location string) (*gcp.GeneratorEvaluationsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-dialogflow.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewGeneratorEvaluationsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building GeneratorEvaluation client: %w", err)
	}
	return gcpClient, err
}

type GeneratorEvaluationAdapter struct {
	id        *krm.DialogflowGeneratorEvaluationIdentity
	gcpClient *gcp.GeneratorEvaluationsClient
	desired   *krm.DialogflowGeneratorEvaluation
	actual    *dialogflowpb.GeneratorEvaluation
}

var _ directbase.Adapter = &GeneratorEvaluationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GeneratorEvaluationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting GeneratorEvaluation", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &dialogflowpb.GetGeneratorEvaluationRequest{
		Name: a.id.String(),
	}
	generatorevaluationpb, err := a.gcpClient.GetGeneratorEvaluation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GeneratorEvaluation %q: %w", a.id, err)
	}

	a.actual = generatorevaluationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GeneratorEvaluationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GeneratorEvaluation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DialogflowGeneratorEvaluationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dialogflowpb.CreateGeneratorEvaluationRequest{
		Parent:              a.id.ParentString(),
		GeneratorEvaluation: resource,
	}
	op, err := a.gcpClient.CreateGeneratorEvaluation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating GeneratorEvaluation %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("GeneratorEvaluation %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created GeneratorEvaluation", "name", a.id)

	// Fetch full GeneratorEvaluation after LRO completion to populate observedState.
	getReq := &dialogflowpb.GetGeneratorEvaluationRequest{
		Name: created.Name,
	}
	actual, err := a.gcpClient.GetGeneratorEvaluation(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting GeneratorEvaluation %s after creation: %w", a.id, err)
	}
	created = actual

	status := &krm.DialogflowGeneratorEvaluationStatus{}
	status.ObservedState = DialogflowGeneratorEvaluationObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.DialogflowGeneratorEvaluationIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.Evaluation = parsedID.Evaluation
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GeneratorEvaluationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GeneratorEvaluation", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DialogflowGeneratorEvaluationSpec_FromProto,
		DialogflowGeneratorEvaluationSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating DialogflowGeneratorEvaluation is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.DialogflowGeneratorEvaluationStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DialogflowGeneratorEvaluationObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GeneratorEvaluationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowGeneratorEvaluation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowGeneratorEvaluationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DialogflowGeneratorEvaluationGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GeneratorEvaluationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GeneratorEvaluation", "name", a.id)

	req := &dialogflowpb.DeleteGeneratorEvaluationRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteGeneratorEvaluation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent GeneratorEvaluation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting GeneratorEvaluation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted GeneratorEvaluation", "name", a.id)
	return true, nil
}
