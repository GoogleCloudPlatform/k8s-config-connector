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
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DataplexGlossaryTerm

package dataplex

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dataplex/apiv1"
	dataplexpb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DataplexGlossaryTermGVK, NewGlossaryTermModel)
}

func NewGlossaryTermModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGlossaryTerm{config: *config}, nil
}

var _ directbase.Model = &modelGlossaryTerm{}

type modelGlossaryTerm struct {
	config config.ControllerConfig
}

func (m *modelGlossaryTerm) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataplexGlossaryTerm{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DataplexGlossaryTermIdentity)

	// Get dataplex GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &GlossaryTermAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGlossaryTerm) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelGlossaryTerm) client(ctx context.Context, location string) (*gcp.BusinessGlossaryClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("https://dataplex.%s.rep.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewBusinessGlossaryRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building GlossaryTerm client: %w", err)
	}
	return gcpClient, err
}

type GlossaryTermAdapter struct {
	id        *krm.DataplexGlossaryTermIdentity
	gcpClient *gcp.BusinessGlossaryClient
	desired   *krm.DataplexGlossaryTerm
	actual    *dataplexpb.GlossaryTerm
}

var _ directbase.Adapter = &GlossaryTermAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GlossaryTermAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting GlossaryTerm", "name", a.id)

	req := &dataplexpb.GetGlossaryTermRequest{
		Name: a.id.String(),
	}
	glossarytermpb, err := a.gcpClient.GetGlossaryTerm(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GlossaryTerm %q: %w", a.id, err)
	}

	a.actual = glossarytermpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GlossaryTermAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GlossaryTerm", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DataplexGlossaryTermSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &dataplexpb.CreateGlossaryTermRequest{
		Parent: a.id.ParentString(),
		TermId: a.id.GlossaryTerm,
		Term:   resource,
	}
	created, err := a.gcpClient.CreateGlossaryTerm(ctx, req)
	if err != nil {
		return fmt.Errorf("creating GlossaryTerm %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created GlossaryTerm", "name", a.id)

	status := &krm.DataplexGlossaryTermStatus{}
	status.ObservedState = DataplexGlossaryTermObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GlossaryTermAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GlossaryTerm", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DataplexGlossaryTermSpec_FromProto,
		DataplexGlossaryTermSpec_ToProto,
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
		req := &dataplexpb.UpdateGlossaryTermRequest{
			UpdateMask: updateMask,
			Term:       resource,
		}
		updated, err = a.gcpClient.UpdateGlossaryTerm(ctx, req)
		if err != nil {
			return fmt.Errorf("updating GlossaryTerm %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated GlossaryTerm", "name", a.id)
	}

	status := &krm.DataplexGlossaryTermStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DataplexGlossaryTermObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GlossaryTermAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataplexGlossaryTerm{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataplexGlossaryTermSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DataplexGlossaryTermGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GlossaryTermAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GlossaryTerm", "name", a.id)

	req := &dataplexpb.DeleteGlossaryTermRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteGlossaryTerm(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent GlossaryTerm, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting GlossaryTerm %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted GlossaryTerm", "name", a.id)
	return true, nil
}
