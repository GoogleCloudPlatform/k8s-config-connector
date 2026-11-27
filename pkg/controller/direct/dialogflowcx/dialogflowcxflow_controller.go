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
// krm.kind: DialogflowCXFlow

package dialogflowcx

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/dialogflow/cx/apiv3"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowCXFlowGVK, NewFlowModel)
}

func NewFlowModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelFlow{config: *config}, nil
}

var _ directbase.Model = &modelFlow{}

type modelFlow struct {
	config config.ControllerConfig
}

func (m *modelFlow) client(ctx context.Context, location string) (*gcp.FlowsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-dialogflow.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewFlowsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Flows REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelFlow) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowCXFlow{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DialogflowCXFlowIdentity)

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &FlowAdapter{
		id:         id,
		gcpClient:  gcpClient,
		desiredKRM: obj,
	}, nil
}

func (m *modelFlow) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowCXFlowIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &FlowAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type FlowAdapter struct {
	id         *krm.DialogflowCXFlowIdentity
	gcpClient  *gcp.FlowsClient
	desiredKRM *krm.DialogflowCXFlow
	actual     *cxpb.Flow
}

var _ directbase.Adapter = &FlowAdapter{}

func (a *FlowAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DialogflowCXFlow", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &cxpb.GetFlowRequest{
		Name: a.id.String(),
	}
	flowpb, err := a.gcpClient.GetFlow(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DialogflowCXFlow %q: %w", a.id, err)
	}

	a.actual = flowpb
	return true, nil
}

func (a *FlowAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DialogflowCXFlow", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desiredKRM.DeepCopy()
	resource := DialogflowCXFlowSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &cxpb.CreateFlowRequest{
		Parent: a.id.ParentString(),
		Flow:   resource,
	}
	if desired.Spec.LanguageCode != nil {
		req.LanguageCode = *desired.Spec.LanguageCode
	}
	created, err := a.gcpClient.CreateFlow(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DialogflowCXFlow %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DialogflowCXFlow", "name", created.GetName())

	parsedID := &krm.DialogflowCXFlowIdentity{}
	if err := parsedID.FromExternal(created.GetName()); err != nil {
		return err
	}
	a.id.Flow = parsedID.Flow
	a.actual = created

	status := &krm.DialogflowCXFlowStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.Flow)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *FlowAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DialogflowCXFlow", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desiredKRM.Spec,
		a.actual,
		DialogflowCXFlowSpec_FromProto,
		DialogflowCXFlowSpec_ToProto,
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
		req := &cxpb.UpdateFlowRequest{
			UpdateMask: updateMask,
			Flow:       resource,
		}
		if a.desiredKRM.Spec.LanguageCode != nil {
			req.LanguageCode = *a.desiredKRM.Spec.LanguageCode
		}
		updated, err := a.gcpClient.UpdateFlow(ctx, req)
		if err != nil {
			return fmt.Errorf("updating DialogflowCXFlow %s: %w", a.id, err)
		}
		a.actual = updated
		log.V(2).Info("successfully updated DialogflowCXFlow", "name", a.id)
	}

	status := &krm.DialogflowCXFlowStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.Flow)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *FlowAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowCXFlow{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowCXFlowSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Parent = direct.LazyPtr(a.id.ParentString())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Flow)
	u.SetGroupVersionKind(krm.DialogflowCXFlowGVK)

	return u, nil
}

func (a *FlowAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DialogflowCXFlow", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	req := &cxpb.DeleteFlowRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteFlow(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DialogflowCXFlow, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DialogflowCXFlow %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DialogflowCXFlow", "name", a.id)
	return true, nil
}
