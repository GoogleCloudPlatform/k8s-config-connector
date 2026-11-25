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
// krm.group: geminidataanalytics.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: GeminiDataAnalyticsDataAgent

package geminidataanalytics

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/geminidataanalytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/geminidataanalytics/apiv1"
	geminidataanalyticspb "cloud.google.com/go/geminidataanalytics/apiv1/geminidataanalyticspb"
	"cloud.google.com/go/iam/apiv1/iampb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.GeminiDataAnalyticsDataAgentGVK, NewDataAgentModel,
		registry.WithIAM(
			krm.GeminiDataAnalyticsDataAgentIdentityFormat.CanonicalForm(),
		),
	)
}

func NewDataAgentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDataAgent{config: *config}, nil
}

var _ directbase.Model = &modelDataAgent{}

type modelDataAgent struct {
	config config.ControllerConfig
}

func (m *modelDataAgent) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.GeminiDataAnalyticsDataAgent{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.GeminiDataAnalyticsDataAgentIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get geminidataanalytics GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataAgentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDataAgent) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//geminidataanalytics.googleapis.com/") {
		return nil, nil
	}
	id := &krm.GeminiDataAnalyticsDataAgentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DataAgentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDataAgent) client(ctx context.Context) (*gcp.DataAgentClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataAgentRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DataAgent client: %w", err)
	}
	return gcpClient, err
}

type DataAgentAdapter struct {
	id        *krm.GeminiDataAnalyticsDataAgentIdentity
	gcpClient *gcp.DataAgentClient
	desired   *krm.GeminiDataAnalyticsDataAgent
	actual    *geminidataanalyticspb.DataAgent
}

var _ directbase.Adapter = &DataAgentAdapter{}
var _ direct.IAMAdapter = &DataAgentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DataAgentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DataAgent", "name", a.id)

	req := &geminidataanalyticspb.GetDataAgentRequest{
		Name: a.id.String(),
	}
	dataagentpb, err := a.gcpClient.GetDataAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataAgent %q: %w", a.id, err)
	}

	a.actual = dataagentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAgentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DataAgent", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GeminiDataAnalyticsDataAgentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &geminidataanalyticspb.CreateDataAgentRequest{
		Parent:      a.id.ParentString(),
		DataAgentId: a.id.DataAgent,
		DataAgent:   resource,
	}
	op, err := a.gcpClient.CreateDataAgent(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DataAgent %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("DataAgent %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created DataAgent", "name", a.id)

	status := &krm.GeminiDataAnalyticsDataAgentStatus{}
	status.ObservedState = GeminiDataAnalyticsDataAgentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DataAgentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DataAgent", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		GeminiDataAnalyticsDataAgentSpec_FromProto,
		GeminiDataAnalyticsDataAgentSpec_ToProto,
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
		req := &geminidataanalyticspb.UpdateDataAgentRequest{
			UpdateMask: updateMask,
			DataAgent:  resource,
		}
		op, err := a.gcpClient.UpdateDataAgent(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DataAgent %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("DataAgent %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DataAgent", "name", a.id)
	}

	status := &krm.GeminiDataAnalyticsDataAgentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = GeminiDataAnalyticsDataAgentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DataAgentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GeminiDataAnalyticsDataAgent{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GeminiDataAnalyticsDataAgentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.GeminiDataAnalyticsDataAgentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DataAgentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DataAgent", "name", a.id)

	req := &geminidataanalyticspb.DeleteDataAgentRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteDataAgent(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent DataAgent, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DataAgent %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DataAgent", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete DataAgent %s: %w", a.id, err)
	}
	return true, nil
}

func (a *DataAgentAdapter) GetIAMPolicy(ctx context.Context) (*iampb.Policy, error) {
	if a.id == nil {
		return nil, fmt.Errorf("cannot get iam policy for missing resource")
	}

	req := &iampb.GetIamPolicyRequest{
		Resource: a.id.String(),
	}
	policy, err := a.gcpClient.GetIamPolicy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting iam policy for %q: %w", a.id.String(), err)
	}

	return policy, nil
}

func (a *DataAgentAdapter) SetIAMPolicy(ctx context.Context, policy *iampb.Policy) (*iampb.Policy, error) {
	if a.id == nil {
		return nil, fmt.Errorf("cannot set iam policy for missing resource")
	}

	req := &iampb.SetIamPolicyRequest{
		Resource: a.id.String(),
		Policy:   policy,
	}
	newPolicy, err := a.gcpClient.SetIamPolicy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("setting iam policy for %q: %w", a.id.String(), err)
	}

	return newPolicy, nil
}
