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
// krm.kind: DialogflowCXWebhook

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
	registry.RegisterModel(krm.DialogflowCXWebhookGVK, NewWebhookModel)
}

func NewWebhookModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWebhook{config: *config}, nil
}

var _ directbase.Model = &modelWebhook{}

type modelWebhook struct {
	config config.ControllerConfig
}

func (m *modelWebhook) client(ctx context.Context, location string) (*gcp.WebhooksClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-dialogflow.googleapis.com", location)))
	}
	gcpClient, err := gcp.NewWebhooksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Webhooks REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelWebhook) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowCXWebhook{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.DialogflowCXWebhookIdentity)

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &WebhookAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWebhook) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//dialogflow.googleapis.com/") {
		return nil, nil
	}
	id := &krm.DialogflowCXWebhookIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &WebhookAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type WebhookAdapter struct {
	id        *krm.DialogflowCXWebhookIdentity
	gcpClient *gcp.WebhooksClient
	desired   *krm.DialogflowCXWebhook
	actual    *cxpb.Webhook
}

var _ directbase.Adapter = &WebhookAdapter{}

func (a *WebhookAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	log.V(2).Info("getting DialogflowCXWebhook", "name", a.id)

	req := &cxpb.GetWebhookRequest{
		Name: a.id.String(),
	}
	webhookpb, err := a.gcpClient.GetWebhook(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DialogflowCXWebhook %q: %w", a.id, err)
	}

	a.actual = webhookpb
	return true, nil
}

func (a *WebhookAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DialogflowCXWebhook", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := DialogflowCXWebhookSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = ""

	req := &cxpb.CreateWebhookRequest{
		Parent:  a.id.ParentString(),
		Webhook: resource,
	}
	created, err := a.gcpClient.CreateWebhook(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DialogflowCXWebhook %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created DialogflowCXWebhook", "name", created.GetName())

	parsedID := &krm.DialogflowCXWebhookIdentity{}
	if err := parsedID.FromExternal(created.GetName()); err != nil {
		return err
	}
	resourceID := parsedID.Webhook
	if a.desired.Spec.ResourceID == nil || *a.desired.Spec.ResourceID == "" {
		if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
			return err
		}
	}
	a.id.Webhook = resourceID
	a.actual = created

	return a.updateStatus(ctx, createOp, created)
}

func (a *WebhookAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DialogflowCXWebhook", "name", a.id)

	diffs, updateMask, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		DialogflowCXWebhookSpec_FromProto,
		DialogflowCXWebhookSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for DialogflowCXWebhook", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	mapCtx := &direct.MapContext{}
	desiredpb := DialogflowCXWebhookSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredpb.Name = a.id.String()

	req := &cxpb.UpdateWebhookRequest{
		Webhook:    desiredpb,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateWebhook(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DialogflowCXWebhook %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated DialogflowCXWebhook", "name", a.id)

	a.actual = updated
	return a.updateStatus(ctx, updateOp, updated)
}

func (a *WebhookAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *cxpb.Webhook) error {
	status := &krm.DialogflowCXWebhookStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.Name = direct.LazyPtr(a.id.Webhook)
	return op.UpdateStatus(ctx, status, nil)
}

func (a *WebhookAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowCXWebhook{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowCXWebhookSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent := a.id.ParentString()
	obj.Spec.Parent = &parent
	if a.id.Webhook != "" {
		obj.Spec.ResourceID = &a.id.Webhook
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Webhook)
	u.SetGroupVersionKind(krm.DialogflowCXWebhookGVK)

	return u, nil
}

func (a *WebhookAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	fqn := a.id.String()
	log.V(2).Info("deleting DialogflowCXWebhook", "name", fqn)

	req := &cxpb.DeleteWebhookRequest{
		Name: fqn,
	}
	err := a.gcpClient.DeleteWebhook(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DialogflowCXWebhook, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DialogflowCXWebhook %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DialogflowCXWebhook", "name", a.id)
	return true, nil
}
