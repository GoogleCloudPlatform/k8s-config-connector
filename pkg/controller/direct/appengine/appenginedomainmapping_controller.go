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

// +tool:controller
// proto.service: google.appengine.v1.DomainMappings
// proto.message: google.appengine.v1.DomainMapping
// crd.type: AppEngineDomainMapping
// crd.version: v1alpha1

package appengine

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	api "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AppEngineDomainMappingGVK, NewAppEngineDomainMappingModel)
}

func NewAppEngineDomainMappingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &domainMappingModel{config: *config}, nil
}

type domainMappingModel struct {
	config config.ControllerConfig
}

var _ directbase.Model = &domainMappingModel{}

func (m *domainMappingModel) client(ctx context.Context) (*api.DomainMappingsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewDomainMappingsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building appengine domain mappings client: %w", err)
	}

	return gcpClient, nil
}

func (m *domainMappingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppEngineDomainMapping{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	domainMappingId := id.(*krm.AppEngineDomainMappingIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &domainMappingAdapter{
		gcpClient: gcpClient,
		id:        domainMappingId,
		desired:   obj,
	}, nil
}

func (m *domainMappingModel) AdapterForURL(ctx context.Context, rawURL string) (directbase.Adapter, error) {
	parsedURL, err := url.Parse(rawURL)
	if err == nil && parsedURL.Path != "" {
		rawURL = strings.TrimPrefix(parsedURL.Path, "/v1/")
		rawURL = strings.TrimPrefix(rawURL, "/")
	}
	id := &krm.AppEngineDomainMappingIdentity{}
	if err := id.FromExternal(rawURL); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &domainMappingAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type domainMappingAdapter struct {
	id        *krm.AppEngineDomainMappingIdentity
	gcpClient *api.DomainMappingsClient
	desired   *krm.AppEngineDomainMapping
	actual    *pb.DomainMapping
}

var _ directbase.Adapter = &domainMappingAdapter{}

func (a *domainMappingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AppEngineDomainMapping", "name", a.id.String())

	req := &pb.GetDomainMappingRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetDomainMapping(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AppEngineDomainMapping %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *domainMappingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AppEngineDomainMapping", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desiredPb := AppEngineDomainMappingSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Id = a.id.DomainMapping

	req := &pb.CreateDomainMappingRequest{
		Parent:        a.id.ParentString(),
		DomainMapping: desiredPb,
	}

	if a.desired.Spec.OverrideStrategy != nil {
		switch *a.desired.Spec.OverrideStrategy {
		case "STRICT":
			req.OverrideStrategy = pb.DomainOverrideStrategy_STRICT
		case "OVERRIDE":
			req.OverrideStrategy = pb.DomainOverrideStrategy_OVERRIDE
		}
	}

	op, err := a.gcpClient.CreateDomainMapping(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AppEngineDomainMapping %s: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineDomainMapping creation %s: %w", a.id.String(), err)
	}

	latest := created
	if v, err := a.gcpClient.GetDomainMapping(ctx, &pb.GetDomainMappingRequest{Name: a.id.String()}); err != nil {
		log.V(2).Info("getting AppEngineDomainMapping after creation failed; falling back to op.Wait result", "name", a.id.String(), "err", err)
	} else {
		latest = v
	}

	log.V(2).Info("successfully created AppEngineDomainMapping", "name", a.id.String())
	return a.updateStatus(ctx, createOp, latest)
}

func (a *domainMappingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AppEngineDomainMapping", "name", a.id.String())

	diffs, updateMask, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AppEngineDomainMappingSpec_v1alpha1_FromProto,
		AppEngineDomainMappingSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for AppEngineDomainMapping", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	mapCtx := &direct.MapContext{}
	desiredPb := AppEngineDomainMappingSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	req := &pb.UpdateDomainMappingRequest{
		Name:          a.id.String(),
		DomainMapping: desiredPb,
		UpdateMask:    updateMask,
	}

	op, err := a.gcpClient.UpdateDomainMapping(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AppEngineDomainMapping %s: %w", a.id.String(), err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for AppEngineDomainMapping update %s: %w", a.id.String(), err)
	}

	latest := updated
	if v, err := a.gcpClient.GetDomainMapping(ctx, &pb.GetDomainMappingRequest{Name: a.id.String()}); err == nil {
		latest = v
	}

	log.V(2).Info("successfully updated AppEngineDomainMapping", "name", a.id.String())
	return a.updateStatus(ctx, updateOp, latest)
}

func (a *domainMappingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AppEngineDomainMapping", "name", a.id.String())

	req := &pb.DeleteDomainMappingRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteDomainMapping(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting AppEngineDomainMapping %s: %w", a.id.String(), err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for AppEngineDomainMapping deletion %s: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *domainMappingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, nil
	}

	mapCtx := &direct.MapContext{}
	spec := AppEngineDomainMappingSpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.Project = direct.LazyPtr(a.id.App)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(krm.AppEngineDomainMappingGVK)
	u.SetName(a.id.DomainMapping)

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("converting spec to unstructured: %w", err)
	}

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec on unstructured: %w", err)
	}

	return u, nil
}

func (a *domainMappingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.DomainMapping) error {
	mapCtx := &direct.MapContext{}
	status := AppEngineDomainMappingStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
