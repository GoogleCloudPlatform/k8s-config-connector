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
// proto.service: google.cloud.discoveryengine.v1beta.LicenseConfigService
// proto.message: google.cloud.discoveryengine.v1beta.LicenseConfig
// crd.type: DiscoveryEngineLicenseConfig
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1beta"
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineLicenseConfigGVK, NewLicenseConfigModel, registry.CannotBeDeleted())
}

func NewLicenseConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &licenseConfigModel{config: *config}, nil
}

var _ directbase.Model = &licenseConfigModel{}

type licenseConfigModel struct {
	config config.ControllerConfig
}

func (m *licenseConfigModel) client(ctx context.Context, projectID string) (*gcp.LicenseConfigClient, error) {
	var opts []option.ClientOption

	config := m.config

	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewLicenseConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine license config client: %w", err)
	}

	return gcpClient, err
}

func (m *licenseConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineLicenseConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DiscoveryEngineLicenseConfigIdentity)

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineLicenseConfigSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &licenseConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *licenseConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		url = strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
	}
	if strings.HasPrefix(url, "discoveryengine.googleapis.com/") {
		url = strings.TrimPrefix(url, "discoveryengine.googleapis.com/")
	}

	id := &krm.DiscoveryEngineLicenseConfigIdentity{}
	if err := id.FromExternal(url); err != nil {
		log.V(2).Error(err, "url did not match DiscoveryEngineLicenseConfig format", "url", url)
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &licenseConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
	}, nil
}

type licenseConfigAdapter struct {
	gcpClient *gcp.LicenseConfigClient
	id        *krm.DiscoveryEngineLicenseConfigIdentity
	desired   *discoveryenginepb.LicenseConfig
	actual    *discoveryenginepb.LicenseConfig
}

var _ directbase.Adapter = &licenseConfigAdapter{}

func (a *licenseConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DiscoveryEngineLicenseConfig", "name", a.id.String())

	req := &discoveryenginepb.GetLicenseConfigRequest{Name: a.id.String()}
	licenseConfig, err := a.gcpClient.GetLicenseConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DiscoveryEngineLicenseConfig %q: %w", a.id.String(), err)
	}

	a.actual = licenseConfig
	return true, nil
}

func (a *licenseConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DiscoveryEngineLicenseConfig", "name", a.id.String())

	desired := proto.Clone(a.desired).(*discoveryenginepb.LicenseConfig)
	desired.Name = a.id.String()

	req := &discoveryenginepb.CreateLicenseConfigRequest{
		Parent:          a.id.ParentString(),
		LicenseConfig:   desired,
		LicenseConfigId: a.id.LicenseConfig,
	}
	created, err := a.gcpClient.CreateLicenseConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating DiscoveryEngineLicenseConfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created DiscoveryEngineLicenseConfig", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *licenseConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DiscoveryEngineLicenseConfig", "name", a.id.String())

	diffs, updateMask, err := compareLicenseConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for DiscoveryEngineLicenseConfig", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	desired := proto.Clone(a.desired).(*discoveryenginepb.LicenseConfig)
	desired.Name = a.id.String()

	req := &discoveryenginepb.UpdateLicenseConfigRequest{
		LicenseConfig: desired,
		UpdateMask:    updateMask,
	}
	updated, err := a.gcpClient.UpdateLicenseConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating DiscoveryEngineLicenseConfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated DiscoveryEngineLicenseConfig", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *licenseConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DiscoveryEngineLicenseConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineLicenseConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.LicenseConfig)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.LicenseConfig)
	u.SetGroupVersionKind(krm.DiscoveryEngineLicenseConfigGVK)

	return u, nil
}

func (a *licenseConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// DiscoveryEngineLicenseConfig does not have a delete method, but we register the model with CannotBeDeleted.
	// Since CannotBeDeleted is used, Delete should theoretically be unreachable or we can just return true, nil.
	return true, nil
}

func compareLicenseConfig(ctx context.Context, actual, desired *discoveryenginepb.LicenseConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineLicenseConfigSpec_v1alpha1_FromProto, DiscoveryEngineLicenseConfigSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.Clone(desired).(*discoveryenginepb.LicenseConfig)

	populateDefaults := func(obj *discoveryenginepb.LicenseConfig) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *licenseConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *discoveryenginepb.LicenseConfig) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DiscoveryEngineLicenseConfigStatus{}
	status.ObservedState = DiscoveryEngineLicenseConfigObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
