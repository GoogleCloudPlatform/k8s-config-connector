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

package monitoring

import (
	"context"
	"fmt"
	"strings"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.MonitoringAlertPolicyGVK, newAlertPolicyModel)
}

func newAlertPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &alertPolicyModel{config: config}, nil
}

type alertPolicyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &alertPolicyModel{}

type alertPolicyAdapter struct {
	id *krm.MonitoringAlertPolicyIdentity

	desired *pb.AlertPolicy
	actual  *pb.AlertPolicy

	alertPolicyClient *monitoring.AlertPolicyClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &alertPolicyAdapter{}

// AdapterForObject implements the Model interface.
func (m *alertPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	alertPolicyClient, err := gcpClient.newAlertPolicyClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.MonitoringAlertPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, err
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := MonitoringAlertPolicySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.UserLabels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &alertPolicyAdapter{
		id:                id.(*krm.MonitoringAlertPolicyIdentity),
		desired:           desiredProto,
		alertPolicyClient: alertPolicyClient,
	}, nil
}

func (m *alertPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //monitoring.googleapis.com/projects/PROJECT_NUMBER/alertPolicies/ALERT_POLICY_ID
	if !strings.HasPrefix(url, "//monitoring.googleapis.com/") {
		return nil, nil
	}

	id := &krm.MonitoringAlertPolicyIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	if !id.HasIdentitySpecified() {
		return nil, nil
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	alertPolicyClient, err := gcpClient.newAlertPolicyClient(ctx)
	if err != nil {
		return nil, err
	}

	return &alertPolicyAdapter{
		id:                id,
		alertPolicyClient: alertPolicyClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *alertPolicyAdapter) Find(ctx context.Context) (bool, error) {
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &pb.GetAlertPolicyRequest{
		Name: a.fullyQualifiedName(),
	}
	alertPolicy, err := a.alertPolicyClient.GetAlertPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = alertPolicy

	return true, nil
}

// Delete implements the Adapter interface.
func (a *alertPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &pb.DeleteAlertPolicyRequest{
		Name: a.fullyQualifiedName(),
	}

	if err := a.alertPolicyClient.DeleteAlertPolicy(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting alert policy %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *alertPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	if a.id.HasIdentitySpecified() {
		return fmt.Errorf("cannot create alert policy %q: server-generated identity is already specified", a.id.String())
	}

	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	parent := a.id.ParentString()

	req := &pb.CreateAlertPolicyRequest{
		Name:        parent,
		AlertPolicy: a.desired,
	}

	log.V(2).Info("creating alert policy", "req", req)
	created, err := a.alertPolicyClient.CreateAlertPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating alert policy: %w", err)
	}
	log.V(2).Info("created alert policy", "alertPolicy", created)

	return a.updateStatus(ctx, createOp, created)
}

// Update implements the Adapter interface.
func (a *alertPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating object", "u", u)

	diffs, updateMask, err := compareAlertPolicy(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		log.V(2).Info("alert policy has diffs", "diffs", diffs)
		req := &pb.UpdateAlertPolicyRequest{
			AlertPolicy: a.desired,
			UpdateMask:  updateMask,
		}
		req.AlertPolicy.Name = a.fullyQualifiedName()

		log.V(2).Info("updating alert policy", "request", req)
		updated, err := a.alertPolicyClient.UpdateAlertPolicy(ctx, req)
		if err != nil {
			return err
		}
		log.V(2).Info("updated alert policy", "alertPolicy", updated)
		a.actual = updated
	} else {
		log.V(2).Info("alert policy has no diffs")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *alertPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("alertPolicy %q not found", a.fullyQualifiedName())
	}

	u := &unstructured.Unstructured{}

	obj := &krm.MonitoringAlertPolicy{}
	mc := &direct.MapContext{}
	obj.Spec = *MonitoringAlertPolicySpec_FromProto(mc, a.actual)
	if err := mc.Err(); err != nil {
		return nil, fmt.Errorf("error converting alertPolicy from API %w", err)
	}

	obj.Spec.ResourceID = direct.LazyPtr(a.id.AlertPolicy)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.AlertPolicy)
	u.SetGroupVersionKind(krm.MonitoringAlertPolicyGVK)

	export.SetProjectID(u, a.id.Project)
	export.SetLabels(u, a.actual.UserLabels)

	return u, nil
}

func (a *alertPolicyAdapter) fullyQualifiedName() string {
	return a.id.String()
}

func (a *alertPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AlertPolicy) error {
	mapCtx := &direct.MapContext{}
	status := MonitoringAlertPolicyStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareAlertPolicy(ctx context.Context, actual, desired *pb.AlertPolicy) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, MonitoringAlertPolicySpec_FromProto, MonitoringAlertPolicySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = actual.Name
	maskedActual.UserLabels = actual.UserLabels

	populateDefaults := func(obj *pb.AlertPolicy) {
		if obj.Enabled == nil {
			obj.Enabled = &wrapperspb.BoolValue{Value: true}
		}
	}

	clonedDesired := proto.CloneOf(desired)
	clonedDesired.Name = actual.Name

	populateDefaults(clonedDesired)
	populateDefaults(maskedActual)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
