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
// krm.group: financialservices.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: FinancialServicesPredictionResult

package financialservices

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/financialservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/financialservices/apiv1"
	financialservicespb "cloud.google.com/go/financialservices/apiv1/financialservicespb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FinancialServicesPredictionResultGVK, NewPredictionResultModel)
}

func NewPredictionResultModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPredictionResult{config: *config}, nil
}

var _ directbase.Model = &modelPredictionResult{}

type modelPredictionResult struct {
	config config.ControllerConfig
}

func (m *modelPredictionResult) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.FinancialServicesPredictionResult{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.FinancialServicesPredictionResultIdentity)

	// Get financialservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PredictionResultAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPredictionResult) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//financialservices.googleapis.com/") {
		return nil, nil
	}
	id := &krm.FinancialServicesPredictionResultIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PredictionResultAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelPredictionResult) client(ctx context.Context) (*gcp.AMLClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAMLRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PredictionResult client: %w", err)
	}
	return gcpClient, err
}

type PredictionResultAdapter struct {
	id        *krm.FinancialServicesPredictionResultIdentity
	gcpClient *gcp.AMLClient
	desired   *krm.FinancialServicesPredictionResult
	actual    *financialservicespb.PredictionResult
}

var _ directbase.Adapter = &PredictionResultAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PredictionResultAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PredictionResult", "name", a.id)

	req := &financialservicespb.GetPredictionResultRequest{
		Name: a.id.String(),
	}
	predictionresultpb, err := a.gcpClient.GetPredictionResult(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PredictionResult %q: %w", a.id, err)
	}

	a.actual = predictionresultpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PredictionResultAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PredictionResult", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FinancialServicesPredictionResultSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &financialservicespb.CreatePredictionResultRequest{
		Parent:             a.id.ParentString(),
		PredictionResultId: a.id.PredictionResult,
		PredictionResult:   resource,
	}
	op, err := a.gcpClient.CreatePredictionResult(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PredictionResult %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("PredictionResult %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PredictionResult", "name", a.id)

	status := &krm.FinancialServicesPredictionResultStatus{}
	status.ObservedState = FinancialServicesPredictionResultObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PredictionResultAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PredictionResult", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		FinancialServicesPredictionResultSpec_FromProto,
		FinancialServicesPredictionResultSpec_ToProto,
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
		req := &financialservicespb.UpdatePredictionResultRequest{
			UpdateMask:       updateMask,
			PredictionResult: resource,
		}
		op, err := a.gcpClient.UpdatePredictionResult(ctx, req)
		if err != nil {
			return fmt.Errorf("updating PredictionResult %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("PredictionResult %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated PredictionResult", "name", a.id)
	}

	status := &krm.FinancialServicesPredictionResultStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = FinancialServicesPredictionResultObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PredictionResultAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FinancialServicesPredictionResult{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FinancialServicesPredictionResultSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.FinancialServicesPredictionResultGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PredictionResultAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PredictionResult", "name", a.id)

	req := &financialservicespb.DeletePredictionResultRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeletePredictionResult(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent PredictionResult, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting PredictionResult %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PredictionResult", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete PredictionResult %s: %w", a.id, err)
	}
	return true, nil
}
