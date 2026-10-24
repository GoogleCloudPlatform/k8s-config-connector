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
// proto.service: google.appengine.v1.Firewall
// proto.message: google.appengine.v1.FirewallRule
// crd.type: AppEngineFirewallRule
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
	registry.RegisterModel(krm.AppEngineFirewallRuleGVK, NewAppEngineFirewallRuleModel)
}

func NewAppEngineFirewallRuleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

var _ directbase.Model = &model{}

func (m *model) client(ctx context.Context) (*api.FirewallClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewFirewallRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building appengine firewall client: %w", err)
	}

	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppEngineFirewallRule{}
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
	ruleId := id.(*krm.AppEngineFirewallRuleIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        ruleId,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, rawURL string) (directbase.Adapter, error) {
	parsedURL, err := url.Parse(rawURL)
	if err == nil && parsedURL.Path != "" {
		rawURL = strings.TrimPrefix(parsedURL.Path, "/v1/")
		rawURL = strings.TrimPrefix(rawURL, "/")
	}
	id := &krm.AppEngineFirewallRuleIdentity{}
	if err := id.FromExternal(rawURL); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type adapter struct {
	id        *krm.AppEngineFirewallRuleIdentity
	gcpClient *api.FirewallClient
	desired   *krm.AppEngineFirewallRule
	actual    *pb.FirewallRule
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AppEngineFirewallRule", "name", a.id)

	req := &pb.GetIngressRuleRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetIngressRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AppEngineFirewallRule %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AppEngineFirewallRule", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := AppEngineFirewallRuleSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateIngressRuleRequest{
		Parent: fmt.Sprintf("apps/%s", a.id.App),
		Rule:   desiredPb,
	}

	created, err := a.gcpClient.CreateIngressRule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AppEngineFirewallRule %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created AppEngineFirewallRule", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AppEngineFirewallRule", "name", a.id)

	diffs, updateMask, err := common.CompareBrownfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AppEngineFirewallRuleSpec_v1alpha1_FromProto,
		AppEngineFirewallRuleSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for AppEngineFirewallRule", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	mapCtx := &direct.MapContext{}
	desiredPb := AppEngineFirewallRuleSpec_v1alpha1_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.UpdateIngressRuleRequest{
		Name:       a.id.String(),
		Rule:       desiredPb,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateIngressRule(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AppEngineFirewallRule %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated AppEngineFirewallRule", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AppEngineFirewallRule", "name", a.id)

	req := &pb.DeleteIngressRuleRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteIngressRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting AppEngineFirewallRule %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted AppEngineFirewallRule", "name", a.id.String())

	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AppEngineFirewallRule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AppEngineFirewallRuleSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Project = &a.id.App
	obj.Spec.ResourceID = direct.LazyPtr(a.id.IngressRule)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.IngressRule)
	u.SetGroupVersionKind(krm.AppEngineFirewallRuleGVK)

	return u, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.FirewallRule) error {
	status := &krm.AppEngineFirewallRuleStatus{}
	return op.UpdateStatus(ctx, status, nil)
}
