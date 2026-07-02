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
	return &modelCloudControl{config: *config}, nil
}

var _ directbase.Model = &modelCloudControl{}

type modelCloudControl struct {
	config config.ControllerConfig
}

func (m *modelCloudControl) client(ctx context.Context) (*gcp.ConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Config client for CloudControl: %w", err)
	}
	return gcpClient, nil
}

func (m *modelCloudControl) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	obj := &krm.CloudSecurityComplianceCloudControl{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.CloudSecurityComplianceCloudControlIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := cloudsecuritycompliance.CloudSecurityComplianceCloudControlSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Name = id.String()

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelCloudControl) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.CloudSecurityComplianceCloudControlIdentity
	gcpClient *gcp.ConfigClient
	desired   *pb.CloudControl
	actual    *pb.CloudControl
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CloudSecurityComplianceCloudControl", "name", a.id.String())

	req := &pb.GetCloudControlRequest{Name: a.id.String()}
	cloudcontrolpb, err := a.gcpClient.GetCloudControl(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("error getting CloudSecurityComplianceCloudControl %q: %w", a.id.String(), err)
	}

	a.actual = cloudcontrolpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CloudSecurityComplianceCloudControl", "name", a.id.String())

	req := &pb.CreateCloudControlRequest{
		Parent:         a.id.ParentString(),
		CloudControlId: a.id.CloudControl,
		CloudControl:   a.desired,
	}
	created, err := a.gcpClient.CreateCloudControl(ctx, req)
	if err != nil {
		return fmt.Errorf("error creating CloudSecurityComplianceCloudControl %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, op, created)
}

func (a *Adapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CloudSecurityComplianceCloudControl", "name", a.id.String())

	paths, err := common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing spec: %w", err)
	}

	if len(paths) == 0 {
		log.V(2).Info("no diff found for CloudSecurityComplianceCloudControl", "name", a.id.String())
		return nil
	}

	report := &structuredreporting.Diff{Object: op.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateCloudControlRequest{
		CloudControl: a.desired,
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}
	updated, err := a.gcpClient.UpdateCloudControl(ctx, req)
	if err != nil {
		return fmt.Errorf("error updating CloudSecurityComplianceCloudControl %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, op, updated)
}

func (a *Adapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CloudSecurityComplianceCloudControl", "name", a.id.String())

	req := &pb.DeleteCloudControlRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteCloudControl(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("error deleting CloudSecurityComplianceCloudControl %q: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CloudControl) error {
	mapCtx := &direct.MapContext{}
	observedState := cloudsecuritycompliance.CloudSecurityComplianceCloudControlObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.CloudSecurityComplianceCloudControlStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = observedState

	return op.UpdateStatus(ctx, status, nil)
}
