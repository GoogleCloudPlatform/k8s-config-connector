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

package muteconfig

import (
	"context"
	"fmt"
	"strings"
	"time"

	gcp "cloud.google.com/go/securitycenter/apiv1"
	securitycenterpb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SecurityCenterMuteConfigGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SecurityCenter MuteConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecurityCenterMuteConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &securityCenterMuteConfigAdapter{
		id:        identity.(*krm.SecurityCenterMuteConfigIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//securitycenter.googleapis.com/") {
		return nil, nil
	}

	parsed := &krm.SecurityCenterMuteConfigIdentity{}
	if err := parsed.FromExternal(strings.TrimPrefix(url, "//")); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &securityCenterMuteConfigAdapter{
		id:        parsed,
		gcpClient: gcpClient,
	}, nil
}

type securityCenterMuteConfigAdapter struct {
	id        *krm.SecurityCenterMuteConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.SecurityCenterMuteConfig
	actual    *securitycenterpb.MuteConfig
}

var _ directbase.Adapter = &securityCenterMuteConfigAdapter{}

func (a *securityCenterMuteConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecurityCenterMuteConfig", "name", a.id.String())

	req := &securitycenterpb.GetMuteConfigRequest{
		Name: a.id.String(),
	}
	muteConfig, err := a.gcpClient.GetMuteConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("SecurityCenterMuteConfig not found", "name", a.id.String())
			return false, nil
		}
		return false, fmt.Errorf("getting SecurityCenterMuteConfig %q: %w", a.id.String(), err)
	}

	a.actual = muteConfig
	return true, nil
}

func (a *securityCenterMuteConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecurityCenterMuteConfig", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	muteConfig := SecurityCenterMuteConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &securitycenterpb.CreateMuteConfigRequest{
		Parent:       fmt.Sprintf("organizations/%s", a.id.Organization),
		MuteConfigId: a.id.MuteConfig,
		MuteConfig:   muteConfig,
	}

	created, err := a.gcpClient.CreateMuteConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SecurityCenterMuteConfig %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created SecurityCenterMuteConfig", "name", a.id.String())

	status := &krm.SecurityCenterMuteConfigStatus{}
	status.ObservedState = SecurityCenterMuteConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *securityCenterMuteConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecurityCenterMuteConfig", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desiredProto := SecurityCenterMuteConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredProto.Name = a.id.String()

	if a.desired.Spec.Type != nil {
		desiredType := MuteConfigType_ToProto(a.desired.Spec.Type)
		if desiredType != a.actual.Type {
			return fmt.Errorf("type is immutable and cannot be updated from %s to %s", a.actual.Type, desiredType)
		}
	} else {
		desiredProto.Type = a.actual.Type
	}

	paths, report, err := common.CompareProtoMessageStructuredDiff(desiredProto, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.SecurityCenterMuteConfigStatus{}
		status.ObservedState = SecurityCenterMuteConfigObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.PtrTo(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, report)

	req := &securitycenterpb.UpdateMuteConfigRequest{
		MuteConfig: desiredProto,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}

	updated, err := a.gcpClient.UpdateMuteConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating SecurityCenterMuteConfig %q: %w", a.id.String(), err)
	}

	status := &krm.SecurityCenterMuteConfigStatus{}
	status.ObservedState = SecurityCenterMuteConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *securityCenterMuteConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	mapCtx := &direct.MapContext{}
	spec := SecurityCenterMuteConfigSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.SecurityCenterMuteConfig{}
	obj.Spec = *spec
	obj.Spec.OrganizationRef = &refs.OrganizationRef{
		External: a.id.Organization,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(a.id.MuteConfig)
	u.SetGroupVersionKind(krm.SecurityCenterMuteConfigGVK)
	u.Object = uObj
	return u, nil
}

func (a *securityCenterMuteConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SecurityCenterMuteConfig", "name", a.id.String())

	req := &securitycenterpb.DeleteMuteConfigRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteMuteConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting SecurityCenterMuteConfig %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted SecurityCenterMuteConfig", "name", a.id.String())
	return true, nil
}

func MuteConfigType_ToProto(s *string) securitycenterpb.MuteConfig_MuteConfigType {
	if s == nil {
		return securitycenterpb.MuteConfig_MUTE_CONFIG_TYPE_UNSPECIFIED
	}
	switch *s {
	case "STATIC":
		return securitycenterpb.MuteConfig_STATIC
	case "DYNAMIC":
		return securitycenterpb.MuteConfig_DYNAMIC
	default:
		return securitycenterpb.MuteConfig_MUTE_CONFIG_TYPE_UNSPECIFIED
	}
}

func MuteConfigType_FromProto(t securitycenterpb.MuteConfig_MuteConfigType) *string {
	var s string
	switch t {
	case securitycenterpb.MuteConfig_STATIC:
		s = "STATIC"
	case securitycenterpb.MuteConfig_DYNAMIC:
		s = "DYNAMIC"
	default:
		s = "MUTE_CONFIG_TYPE_UNSPECIFIED"
	}
	return &s
}

func SecurityCenterMuteConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCenterMuteConfigSpec) *securitycenterpb.MuteConfig {
	if in == nil {
		return nil
	}
	out := &securitycenterpb.MuteConfig{}
	if in.Description != nil {
		out.Description = *in.Description
	}
	if in.Filter != nil {
		out.Filter = *in.Filter
	}
	out.Type = MuteConfigType_ToProto(in.Type)
	if in.ExpiryTime != nil {
		out.ExpiryTime = StringTimestamp_ToProto(mapCtx, in.ExpiryTime)
	}
	return out
}

func SecurityCenterMuteConfigSpec_FromProto(mapCtx *direct.MapContext, in *securitycenterpb.MuteConfig) *krm.SecurityCenterMuteConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCenterMuteConfigSpec{}
	out.Description = &in.Description
	out.Filter = &in.Filter
	out.Type = MuteConfigType_FromProto(in.Type)
	out.ExpiryTime = StringTimestamp_FromProto(mapCtx, in.ExpiryTime)
	return out
}

func SecurityCenterMuteConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecurityCenterMuteConfigObservedState) *securitycenterpb.MuteConfig {
	if in == nil {
		return nil
	}
	out := &securitycenterpb.MuteConfig{}
	out.CreateTime = StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	if in.MostRecentEditor != nil {
		out.MostRecentEditor = *in.MostRecentEditor
	}
	return out
}

func SecurityCenterMuteConfigObservedState_FromProto(mapCtx *direct.MapContext, in *securitycenterpb.MuteConfig) *krm.SecurityCenterMuteConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecurityCenterMuteConfigObservedState{}
	out.CreateTime = StringTimestamp_FromProto(mapCtx, in.CreateTime)
	out.UpdateTime = StringTimestamp_FromProto(mapCtx, in.UpdateTime)
	if in.MostRecentEditor != "" {
		out.MostRecentEditor = &in.MostRecentEditor
	}
	return out
}

func StringTimestamp_FromProto(mapCtx *direct.MapContext, ts *timestamppb.Timestamp) *string {
	if ts == nil {
		return nil
	}
	formatted := ts.AsTime().Format(time.RFC3339Nano)
	return &formatted
}

func StringTimestamp_ToProto(mapCtx *direct.MapContext, s *string) *timestamppb.Timestamp {
	if s == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *s)
	if err != nil {
		mapCtx.Errorf("invalid timestamp %q", *s)
	}
	ts := timestamppb.New(t)
	return ts
}
