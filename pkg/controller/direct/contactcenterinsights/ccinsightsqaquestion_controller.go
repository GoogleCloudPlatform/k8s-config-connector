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
// krm.group: contactcenterinsights.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CCInsightsQaQuestion

package contactcenterinsights

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/contactcenterinsights/apiv1"
	contactcenterinsightspb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CCInsightsQaQuestionGVK, NewQaQuestionModel)
}

func NewQaQuestionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelQaQuestion{config: *config}, nil
}

var _ directbase.Model = &modelQaQuestion{}

type modelQaQuestion struct {
	config config.ControllerConfig
}

func (m *modelQaQuestion) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CCInsightsQaQuestion{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.CCInsightsQaQuestionIdentity)

	// Get contactcenterinsights GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &QaQuestionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelQaQuestion) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//contactcenterinsights.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CCInsightsQaQuestionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &QaQuestionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelQaQuestion) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building QaQuestion client: %w", err)
	}
	return gcpClient, err
}

type QaQuestionAdapter struct {
	id        *krm.CCInsightsQaQuestionIdentity
	gcpClient *gcp.Client
	desired   *krm.CCInsightsQaQuestion
	actual    *contactcenterinsightspb.QaQuestion
}

var _ directbase.Adapter = &QaQuestionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *QaQuestionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting QaQuestion", "name", a.id)

	req := &contactcenterinsightspb.GetQaQuestionRequest{
		Name: a.id.String(),
	}
	qaquestionpb, err := a.gcpClient.GetQaQuestion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting QaQuestion %q: %w", a.id, err)
	}

	a.actual = qaquestionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *QaQuestionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating QaQuestion", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CCInsightsQaQuestionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &contactcenterinsightspb.CreateQaQuestionRequest{
		Parent:       a.id.ParentString(),
		QaQuestionId: a.id.QaQuestion,
		QaQuestion:   resource,
	}
	created, err := a.gcpClient.CreateQaQuestion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating QaQuestion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created QaQuestion", "name", a.id)

	status := &krm.CCInsightsQaQuestionStatus{}
	status.ObservedState = CCInsightsQaQuestionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *QaQuestionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating QaQuestion", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CCInsightsQaQuestionSpec_FromProto,
		CCInsightsQaQuestionSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &contactcenterinsightspb.UpdateQaQuestionRequest{
			UpdateMask: updateMask,
			QaQuestion: resource,
		}
		updated, err = a.gcpClient.UpdateQaQuestion(ctx, req)
		if err != nil {
			return fmt.Errorf("updating QaQuestion %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated QaQuestion", "name", a.id)
	}

	status := &krm.CCInsightsQaQuestionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CCInsightsQaQuestionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *QaQuestionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CCInsightsQaQuestion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CCInsightsQaQuestionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.CCInsightsQaQuestionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *QaQuestionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting QaQuestion", "name", a.id)

	req := &contactcenterinsightspb.DeleteQaQuestionRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteQaQuestion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent QaQuestion, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting QaQuestion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted QaQuestion", "name", a.id)
	return true, nil
}
