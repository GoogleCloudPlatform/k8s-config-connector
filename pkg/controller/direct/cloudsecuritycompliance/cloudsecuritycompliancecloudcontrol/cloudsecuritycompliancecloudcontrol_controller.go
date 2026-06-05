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

package cloudsecuritycompliancecloudcontrol

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsecuritycompliance/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudsecuritycompliance"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/cloudsecuritycompliance/apiv1"
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudSecurityComplianceCloudControlGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.ConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CloudControl client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudSecurityComplianceCloudControl{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.CloudSecurityComplianceCloudControlIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.CloudSecurityComplianceCloudControlIdentity
	gcpClient *gcp.ConfigClient
	desired   *krm.CloudSecurityComplianceCloudControl
	actual    *pb.CloudControl
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CloudControl", "name", a.id)

	req := &pb.GetCloudControlRequest{Name: a.id.String()}
	cloudcontrolpb, err := a.gcpClient.GetCloudControl(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CloudControl %q: %w", a.id, err)
	}

	a.actual = cloudcontrolpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CloudControl", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := cloudsecuritycompliance.CloudSecurityComplianceCloudControlSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("organizations/%s/locations/%s", a.id.Organization, a.id.Location)
	req := &pb.CreateCloudControlRequest{
		Parent:         parent,
		CloudControlId: a.id.CloudControl,
		CloudControl:   resource,
	}
	created, err := a.gcpClient.CreateCloudControl(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CloudControl %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CloudControl", "name", a.id)

	status := &krm.CloudSecurityComplianceCloudControlStatus{}
	status.ObservedState = cloudsecuritycompliance.CloudSecurityComplianceCloudControlObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CloudControl", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := cloudsecuritycompliance.CloudSecurityComplianceCloudControlSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.CloudSecurityComplianceCloudControlStatus{}
		status.ObservedState = cloudsecuritycompliance.CloudSecurityComplianceCloudControlObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateCloudControlRequest{
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		CloudControl: desiredPb,
	}
	updated, err := a.gcpClient.UpdateCloudControl(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CloudControl %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated CloudControl", "name", a.id.String())

	status := &krm.CloudSecurityComplianceCloudControlStatus{}
	status.ObservedState = cloudsecuritycompliance.CloudSecurityComplianceCloudControlObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudSecurityComplianceCloudControl{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(cloudsecuritycompliance.CloudSecurityComplianceCloudControlSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.OrganizationRef = &refs.OrganizationRef{External: a.id.Organization}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.CloudSecurityComplianceCloudControlGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CloudControl", "name", a.id)

	req := &pb.DeleteCloudControlRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteCloudControl(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CloudControl, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CloudControl %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CloudControl", "name", a.id)
	return true, nil
}
