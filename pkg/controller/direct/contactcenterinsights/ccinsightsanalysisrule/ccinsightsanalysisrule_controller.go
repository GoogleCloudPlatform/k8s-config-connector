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
// proto.message: google.cloud.contactcenterinsights.v1.AnalysisRule
// crd.type: CCInsightsAnalysisRule
// crd.version: v1alpha1

package ccinsightsanalysisrule

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/contactcenterinsights/apiv1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
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
	registry.RegisterModel(krm.CCInsightsAnalysisRuleGVK, NewModel)
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
		return nil, fmt.Errorf("building contactcenterinsights client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CCInsightsAnalysisRule{}
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
	id := idAny.(*krm.CCInsightsAnalysisRuleIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := CCInsightsAnalysisRuleSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if desired.DisplayName == nil || *desired.DisplayName == "" {
		desired.DisplayName = direct.LazyPtr(u.GetName())
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CCInsightsAnalysisRuleIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
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
	gcpClient *gcp.Client
	id        *krm.CCInsightsAnalysisRuleIdentity
	desired   *pb.AnalysisRule
	actual    *pb.AnalysisRule
	reader    client.Reader
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CCInsightsAnalysisRule", "name", a.id)

	req := &pb.GetAnalysisRuleRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetAnalysisRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) || direct.IsBadRequest(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CCInsightsAnalysisRule %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CCInsightsAnalysisRule", "name", a.id)

	a.desired.Name = ""
	req := &pb.CreateAnalysisRuleRequest{
		Parent:       a.id.ParentString(),
		AnalysisRule: a.desired,
	}
	created, err := a.gcpClient.CreateAnalysisRule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CCInsightsAnalysisRule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CCInsightsAnalysisRule", "name", created.Name)

	if err := a.id.FromExternal(created.Name); err != nil {
		return fmt.Errorf("parsing server-assigned analysis rule name %q: %w", created.Name, err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CCInsightsAnalysisRule", "name", a.id)

	diffs, updateMask, err := a.compare(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	a.desired.Name = a.id.String()
	req := &pb.UpdateAnalysisRuleRequest{
		AnalysisRule: a.desired,
		UpdateMask:   updateMask,
	}
	_, err = a.gcpClient.UpdateAnalysisRule(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CCInsightsAnalysisRule %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated CCInsightsAnalysisRule", "name", a.id)

	latest, err := a.gcpClient.GetAnalysisRule(ctx, &pb.GetAnalysisRuleRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting CCInsightsAnalysisRule %s after update: %w", a.id, err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *adapter) compare(ctx context.Context, actual, desired *pb.AnalysisRule) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CCInsightsAnalysisRuleSpec_FromProto, CCInsightsAnalysisRuleSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}

	clonedDesired := proto.Clone(desired).(*pb.AnalysisRule)
	clonedDesired.Name = actual.Name
	maskedActual.Name = actual.Name

	populateDefaults := func(obj *pb.AnalysisRule) {
		if (obj.DisplayName == nil || *obj.DisplayName == "") && actual.DisplayName != nil {
			obj.DisplayName = actual.DisplayName
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.AnalysisRule) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CCInsightsAnalysisRuleStatus{}
	status.ObservedState = CCInsightsAnalysisRuleObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CCInsightsAnalysisRule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CCInsightsAnalysisRuleSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.AnalysisRule)
	u.SetGroupVersionKind(krm.CCInsightsAnalysisRuleGVK)
	u.Object = uObj
	return u, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CCInsightsAnalysisRule", "name", a.id)

	req := &pb.DeleteAnalysisRuleRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAnalysisRule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) || direct.IsBadRequest(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting CCInsightsAnalysisRule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CCInsightsAnalysisRule", "name", a.id)
	return true, nil
}
