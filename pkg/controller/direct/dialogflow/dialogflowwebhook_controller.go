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

package dialogflow

import (
	"context"
	"fmt"

	cx "cloud.google.com/go/dialogflow/cx/apiv3"
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DialogflowWebhookGVK, NewDialogflowWebhookModel)
}

func NewDialogflowWebhookModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &webhookModel{config: *config}, nil
}

type webhookModel struct {
	config config.ControllerConfig
}

func (m *webhookModel) client(ctx context.Context, location string) (*cx.WebhooksClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-dialogflow.googleapis.com:443", location)))
	}
	gcpClient, err := cx.NewWebhooksRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Webhooks REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *webhookModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DialogflowWebhook{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	if obj.Spec.GenericWebService != nil && obj.Spec.GenericWebService.Password != nil {
		if err := obj.Spec.GenericWebService.Password.NormalizeSecret(ctx, "spec.genericWebService.password", obj.GetNamespace(), reader); err != nil {
			return nil, err
		}
	}
	if obj.Spec.ServiceDirectory != nil && obj.Spec.ServiceDirectory.GenericWebService != nil && obj.Spec.ServiceDirectory.GenericWebService.Password != nil {
		if err := obj.Spec.ServiceDirectory.GenericWebService.Password.NormalizeSecret(ctx, "spec.serviceDirectory.genericWebService.password", obj.GetNamespace(), reader); err != nil {
			return nil, err
		}
	}
	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DialogflowWebhookIdentity)

	mapCtx := &direct.MapContext{}
	desired := DialogflowWebhookSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &webhookAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *webhookModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.DialogflowWebhookIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &webhookAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type webhookAdapter struct {
	id        *krm.DialogflowWebhookIdentity
	gcpClient *cx.WebhooksClient
	desired   *cxpb.Webhook
	actual    *cxpb.Webhook
}

var _ directbase.Adapter = &webhookAdapter{}

func (a *webhookAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id.Webhook == "" {
		return false, nil
	}
	fqn := a.id.String()
	log.V(2).Info("getting DialogflowWebhook", "name", fqn)

	req := &cxpb.GetWebhookRequest{Name: fqn}
	actual, err := a.gcpClient.GetWebhook(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DialogflowWebhook %q: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *webhookAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DialogflowWebhook", "name", fqn)

	parent := a.id.ParentString()
	clonedDesired := proto.Clone(a.desired).(*cxpb.Webhook)
	clonedDesired.Name = ""
	req := &cxpb.CreateWebhookRequest{
		Parent:  parent,
		Webhook: clonedDesired,
	}
	created, err := a.gcpClient.CreateWebhook(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DialogflowWebhook %s: %w", fqn, err)
	}
	log.V(2).Info("successfully created DialogflowWebhook", "name", created.Name)

	if err := a.id.FromExternal(created.Name); err != nil {
		return err
	}
	a.actual = created

	return a.updateStatus(ctx, createOp, created)
}

func (a *webhookAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DialogflowWebhook", "name", fqn)

	maskedActual, err := mappers.OnlySpecFields(a.actual, DialogflowWebhookSpec_FromProto, DialogflowWebhookSpec_ToProto)
	if err != nil {
		return err
	}

	clonedDesired := proto.Clone(a.desired).(*cxpb.Webhook)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", fqn)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = u
	structuredreporting.ReportDiff(ctx, diffs)

	clonedDesired.Name = a.actual.Name
	req := &cxpb.UpdateWebhookRequest{
		Webhook:    clonedDesired,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateWebhook(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DialogflowWebhook %s: %w", fqn, err)
	}
	log.V(2).Info("successfully updated DialogflowWebhook", "name", fqn)

	a.actual = updated
	return a.updateStatus(ctx, updateOp, updated)
}

func (a *webhookAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *cxpb.Webhook) error {
	status := &krm.DialogflowWebhookStatus{}
	status.ExternalRef = &latest.Name
	return op.UpdateStatus(ctx, status, nil)
}

func (a *webhookAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DialogflowWebhook{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DialogflowWebhookSpec_FromProto(mapCtx, a.actual))
	obj.Spec.AgentRef = &krm.DialogflowCXAgentRef{
		External: a.id.ParentString(),
	}
	if a.id.Webhook != "" {
		obj.Spec.ResourceID = &a.id.Webhook
	}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.Webhook)
	u.SetGroupVersionKind(krm.DialogflowWebhookGVK)
	return u, nil
}

func (a *webhookAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	if a.id.Webhook == "" {
		return false, nil
	}
	fqn := a.id.String()
	log.V(2).Info("deleting DialogflowWebhook", "name", fqn)

	req := &cxpb.DeleteWebhookRequest{
		Name: fqn,
	}
	err := a.gcpClient.DeleteWebhook(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting DialogflowWebhook %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted DialogflowWebhook", "name", fqn)
	return true, nil
}
