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
// krm.group: translate.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: TranslateGlossaryEntry

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
	registry.RegisterModel(krm.TranslateGlossaryEntryGVK, NewGlossaryEntryModel)
}

func NewGlossaryEntryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGlossaryEntry{config: *config}, nil
}

var _ directbase.Model = &modelGlossaryEntry{}

type modelGlossaryEntry struct {
	config config.ControllerConfig
}

func (m *modelGlossaryEntry) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.TranslateGlossaryEntry{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.TranslateGlossaryEntryIdentity)

	// Get translate GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &GlossaryEntryAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelGlossaryEntry) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//translate.googleapis.com/") {
		return nil, nil
	}
	id := &krm.TranslateGlossaryEntryIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &GlossaryEntryAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelGlossaryEntry) client(ctx context.Context) (*gcp.TranslationClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTranslationRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building GlossaryEntry client: %w", err)
	}
	return gcpClient, err
}

type GlossaryEntryAdapter struct {
	id        *krm.TranslateGlossaryEntryIdentity
	gcpClient *gcp.TranslationClient
	desired   *krm.TranslateGlossaryEntry
	actual    *translatepb.GlossaryEntry
}

var _ directbase.Adapter = &GlossaryEntryAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *GlossaryEntryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting GlossaryEntry", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &translatepb.GetGlossaryEntryRequest{
		Name: a.id.String(),
	}
	glossaryentrypb, err := a.gcpClient.GetGlossaryEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GlossaryEntry %q: %w", a.id, err)
	}

	a.actual = glossaryentrypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GlossaryEntryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GlossaryEntry", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := TranslateGlossaryEntrySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &translatepb.CreateGlossaryEntryRequest{
		Parent:        a.id.ParentString(),
		GlossaryEntry: resource,
	}
	created, err := a.gcpClient.CreateGlossaryEntry(ctx, req)
	if err != nil {
		return fmt.Errorf("creating GlossaryEntry %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created GlossaryEntry", "name", a.id)

	status := &krm.TranslateGlossaryEntryStatus{}
	parsedID := &krm.TranslateGlossaryEntryIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.GlossaryEntry = parsedID.GlossaryEntry
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *GlossaryEntryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GlossaryEntry", "name", a.id)

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		TranslateGlossaryEntrySpec_FromProto,
		TranslateGlossaryEntrySpec_ToProto,
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
		req := &translatepb.UpdateGlossaryEntryRequest{
			GlossaryEntry: resource,
		}
		_, err = a.gcpClient.UpdateGlossaryEntry(ctx, req)
		if err != nil {
			return fmt.Errorf("updating GlossaryEntry %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated GlossaryEntry", "name", a.id)
	}

	status := &krm.TranslateGlossaryEntryStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *GlossaryEntryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TranslateGlossaryEntry{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TranslateGlossaryEntrySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.TranslateGlossaryEntryGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *GlossaryEntryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GlossaryEntry", "name", a.id)

	req := &translatepb.DeleteGlossaryEntryRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteGlossaryEntry(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent GlossaryEntry, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting GlossaryEntry %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted GlossaryEntry", "name", a.id)
	return true, nil
}
