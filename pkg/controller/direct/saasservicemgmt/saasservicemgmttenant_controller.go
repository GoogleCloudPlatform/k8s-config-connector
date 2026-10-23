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

package saasservicemgmt

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1"
	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/saasservicemgmt/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SaaSServiceMgmtTenantGVK, NewTenantModel)
}

func NewTenantModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &tenantModel{config: config}, nil
}

var _ directbase.Model = &tenantModel{}

type tenantModel struct {
	config *config.ControllerConfig
}

func (m *tenantModel) client(ctx context.Context) (*gcp.SaasDeploymentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSaasDeploymentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SaasServiceMgmt deployments client: %w", err)
	}
	return gcpClient, nil
}

func (m *tenantModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SaaSServiceMgmtTenant{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Resolve resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.SaaSServiceMgmtTenantIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert KRM spec to API format
	mapCtx := &direct.MapContext{}
	desired := SaaSServiceMgmtTenantSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())
	desired.Annotations = filterSystemAnnotations(obj.GetAnnotations())

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &SaaSServiceMgmtTenantAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *tenantModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id, match, err := krm.SaaSServiceMgmtTenantIdentityFormat.Parse(url)
	if err != nil || !match {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SaaSServiceMgmtTenantAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type SaaSServiceMgmtTenantAdapter struct {
	id        *krm.SaaSServiceMgmtTenantIdentity
	gcpClient *gcp.SaasDeploymentsClient
	desired   *pb.Tenant
	actual    *pb.Tenant
}

var _ directbase.Adapter = &SaaSServiceMgmtTenantAdapter{}

func (a *SaaSServiceMgmtTenantAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting SaaSServiceMgmtTenant", "name", fqn)

	req := &pb.GetTenantRequest{
		Name: fqn,
	}
	resource, err := a.gcpClient.GetTenant(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SaaSServiceMgmtTenant %q: %w", fqn, err)
	}

	a.actual = resource
	return true, nil
}

func (a *SaaSServiceMgmtTenantAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	parent := a.id.ParentString()
	fqn := a.id.String()
	log.V(2).Info("creating SaaSServiceMgmtTenant", "name", fqn)

	req := &pb.CreateTenantRequest{
		Parent:   parent,
		TenantId: a.id.Tenant,
		Tenant:   a.desired,
	}
	created, err := a.gcpClient.CreateTenant(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SaaSServiceMgmtTenant %s: %w", a.id.Tenant, err)
	}
	log.V(2).Info("successfully created SaaSServiceMgmtTenant", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *SaaSServiceMgmtTenantAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("updating SaaSServiceMgmtTenant", "name", fqn)

	diffs, updateMask, err := compareTenant(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = fqn

		req := &pb.UpdateTenantRequest{
			Tenant:     desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateTenant(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SaaSServiceMgmtTenant %s: %w", fqn, err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *SaaSServiceMgmtTenantAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Tenant) error {
	mapCtx := &direct.MapContext{}
	status := SaaSServiceMgmtTenantStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *SaaSServiceMgmtTenantAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SaaSServiceMgmtTenant{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SaaSServiceMgmtTenantSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Tenant)
	u.SetGroupVersionKind(krm.SaaSServiceMgmtTenantGVK)

	export.SetLabels(u, a.actual.Labels)
	return u, nil
}

func (a *SaaSServiceMgmtTenantAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting SaaSServiceMgmtTenant", "name", fqn)

	req := &pb.DeleteTenantRequest{Name: fqn}
	err := a.gcpClient.DeleteTenant(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent SaaSServiceMgmtTenant, assuming it was already deleted", "name", fqn)
			return true, nil
		}
		return false, fmt.Errorf("deleting SaaSServiceMgmtTenant %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted SaaSServiceMgmtTenant", "name", fqn)
	return true, nil
}

func compareTenant(ctx context.Context, actual, desired *pb.Tenant) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, SaaSServiceMgmtTenantSpec_FromProto, SaaSServiceMgmtTenantSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Labels = actual.Labels
	maskedActual.Annotations = filterSystemAnnotations(actual.Annotations)
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func SaaSServiceMgmtTenantStatus_FromProto(mapCtx *direct.MapContext, in *pb.Tenant) *krm.SaaSServiceMgmtTenantStatus {
	if in == nil {
		return nil
	}
	out := &krm.SaaSServiceMgmtTenantStatus{}
	if in.Name != "" {
		out.ExternalRef = &in.Name
	}
	out.ObservedState = SaaSServiceMgmtTenantObservedState_FromProto(mapCtx, in)
	return out
}
