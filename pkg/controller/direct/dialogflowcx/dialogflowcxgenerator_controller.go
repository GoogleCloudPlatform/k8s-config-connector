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
// krm.group: dialogflowcx.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: DialogflowCXGenerator

package dialogflowcx

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dialogflow/cx/apiv3"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowCXGeneratorGVK, NewGeneratorModel)
}

func NewGeneratorModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGenerator{config: *config}, nil
}

var _ directbase.Model = &modelGenerator{}

type modelGenerator struct {
	config config.ControllerConfig
}

func (m *modelGenerator) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowCXGenerator{}
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
	id := i.(*krm.DialogflowCXGeneratorIdentity)

	// Get dialogflowcx GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &GeneratorAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGenerator) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowCXGeneratorIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &GeneratorAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelGenerator) client(ctx context.Context, location string) (*gcp.GeneratorsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-dialogflow.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewGeneratorsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Generator client: %w", err)
	}
	return gcpClient, err
}

type GeneratorAdapter struct {
	id        *krm.DialogflowCXGeneratorIdentity
	gcpClient *gcp.GeneratorsClient
	desired   *krm.DialogflowCXGenerator
	actual    *cxpb.Generator
}

var _ directbase.Adapter = &GeneratorAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GeneratorAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Generator", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &cxpb.GetGeneratorRequest{
		Name: a.id.String(),
	}
	generatorpb, err := a.gcpClient.GetGenerator(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Generator %q: %w", a.id, err)
	}

	a.actual = generatorpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GeneratorAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Generator", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DialogflowCXGeneratorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cxpb.CreateGeneratorRequest{
		Parent:    a.id.ParentString(),
		Generator: resource,
	}
	created, err := a.gcpClient.CreateGenerator(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Generator %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Generator", "name", created.GetName())

	status := &krm.DialogflowCXGeneratorStatus{}
	parsedID := &krm.DialogflowCXGeneratorIdentity{}
	if err := parsedID.FromExternal(created.GetName()); err != nil {
		return err
	}
	a.id.Generator = parsedID.Generator
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GeneratorAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Generator", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DialogflowCXGeneratorSpec_FromProto,
		DialogflowCXGeneratorSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &cxpb.UpdateGeneratorRequest{
			UpdateMask: updateMask,
			Generator:  resource,
		}
		_, err = a.gcpClient.UpdateGenerator(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Generator %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Generator", "name", a.id)
	}

	status := &krm.DialogflowCXGeneratorStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GeneratorAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowCXGenerator{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowCXGeneratorSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.DialogflowCXGeneratorGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GeneratorAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Generator", "name", a.id)

	req := &cxpb.DeleteGeneratorRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteGenerator(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Generator, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Generator %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Generator", "name", a.id)
	return true, nil
}
