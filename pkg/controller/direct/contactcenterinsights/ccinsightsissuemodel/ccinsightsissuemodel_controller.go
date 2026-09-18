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
// proto.service: google.cloud.contactcenterinsights.v1.ContactCenterInsights
// proto.message: google.cloud.contactcenterinsights.v1.IssueModel
// crd.type: CCInsightsIssueModel
// crd.version: v1alpha1

package ccinsightsissuemodel

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/contactcenterinsights/apiv1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CCInsightsIssueModelGVK, NewModel)
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
		return nil, fmt.Errorf("building contactcenterinsights REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CCInsightsIssueModel{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idAny, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idAny.(*krm.CCInsightsIssueModelIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := CCInsightsIssueModelSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type adapter struct {
	gcpClient *gcp.Client
	id        *krm.CCInsightsIssueModelIdentity
	desired   *pb.IssueModel
	actual    *pb.IssueModel
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding CCInsightsIssueModel", "name", a.id)

	req := &pb.GetIssueModelRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetIssueModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CCInsightsIssueModel %s: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CCInsightsIssueModel", "name", a.id)

	a.desired.Name = a.id.String()
	req := &pb.CreateIssueModelRequest{
		Parent:     a.id.ParentString(),
		IssueModel: a.desired,
	}

	op, err := a.gcpClient.CreateIssueModel(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CCInsightsIssueModel %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for CCInsightsIssueModel creation %s: %w", a.id.String(), err)
	}

	// Fetch fully-populated resource after creation LRO completes
	latest, err := a.gcpClient.GetIssueModel(ctx, &pb.GetIssueModelRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting CCInsightsIssueModel after creation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("created CCInsightsIssueModel", "name", a.id)
	return a.updateStatus(ctx, createOp, latest)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CCInsightsIssueModel", "name", a.id)

	mapCtx := &direct.MapContext{}
	actualSpec := CCInsightsIssueModelSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := CCInsightsIssueModelSpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.IssueModel)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	issueModel := proto.Clone(clonedDesired).(*pb.IssueModel)
	issueModel.Name = a.id.String()

	req := &pb.UpdateIssueModelRequest{
		IssueModel: issueModel,
		UpdateMask: updateMask,
	}

	updated, err := a.gcpClient.UpdateIssueModel(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CCInsightsIssueModel %s: %w", a.id.String(), err)
	}

	log.V(2).Info("updated CCInsightsIssueModel", "name", a.id)
	return a.updateStatus(ctx, updateOp, updated)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CCInsightsIssueModel", "name", a.id)

	req := &pb.DeleteIssueModelRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteIssueModel(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting CCInsightsIssueModel %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for CCInsightsIssueModel deletion %s: %w", a.id, err)
	}

	log.V(2).Info("deleted CCInsightsIssueModel", "name", a.id)
	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	mapCtx := &direct.MapContext{}
	spec := CCInsightsIssueModelSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.CCInsightsIssueModel{}
	obj.Spec = *spec

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: specObj}
	u.SetGroupVersionKind(krm.CCInsightsIssueModelGVK)
	return u, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.IssueModel) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CCInsightsIssueModelStatus{}
	status.ObservedState = CCInsightsIssueModelObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}
