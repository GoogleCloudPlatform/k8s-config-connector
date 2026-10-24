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
// krm.group: translate.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: TranslateGlossary

package translate

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/translate/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/translate/apiv3"
	translatepb "cloud.google.com/go/translate/apiv3/translatepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.TranslateGlossaryGVK, NewGlossaryModel)
}

func NewGlossaryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGlossary{config: *config}, nil
}

var _ directbase.Model = &modelGlossary{}

type modelGlossary struct {
	config config.ControllerConfig
}

func (m *modelGlossary) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.TranslateGlossary{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.TranslateGlossaryIdentity)

	// Get translate GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &GlossaryAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGlossary) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//translate.googleapis.com/") {
		return nil, nil
	}
	id := &krm.TranslateGlossaryIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &GlossaryAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelGlossary) client(ctx context.Context) (*gcp.TranslationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTranslationRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Glossary client: %w", err)
	}
	return gcpClient, err
}

type GlossaryAdapter struct {
	id        *krm.TranslateGlossaryIdentity
	gcpClient *gcp.TranslationClient
	desired   *krm.TranslateGlossary
	actual    *translatepb.Glossary
}

var _ directbase.Adapter = &GlossaryAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GlossaryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Glossary", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &translatepb.GetGlossaryRequest{
		Name: a.id.String(),
	}
	glossarypb, err := a.gcpClient.GetGlossary(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Glossary %q: %w", a.id, err)
	}

	a.actual = glossarypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GlossaryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Glossary", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := TranslateGlossarySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &translatepb.CreateGlossaryRequest{
		Parent:   a.id.ParentString(),
		Glossary: resource,
	}
	op, err := a.gcpClient.CreateGlossary(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Glossary %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Glossary %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Glossary", "name", a.id)

	status := &krm.TranslateGlossaryStatus{}
	status.ObservedState = TranslateGlossaryObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.TranslateGlossaryIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.Glossary = parsedID.Glossary
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GlossaryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Glossary", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		TranslateGlossarySpec_FromProto,
		TranslateGlossarySpec_ToProto,
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
		req := &translatepb.UpdateGlossaryRequest{
			UpdateMask: updateMask,
			Glossary:   resource,
		}
		op, err := a.gcpClient.UpdateGlossary(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Glossary %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Glossary %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Glossary", "name", a.id)
	}

	status := &krm.TranslateGlossaryStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = TranslateGlossaryObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GlossaryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TranslateGlossary{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TranslateGlossarySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.TranslateGlossaryGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GlossaryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Glossary", "name", a.id)

	req := &translatepb.DeleteGlossaryRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteGlossary(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Glossary, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Glossary %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Glossary", "name", a.id)
	_, err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Glossary %s: %w", a.id, err)
	}
	return true, nil
}
