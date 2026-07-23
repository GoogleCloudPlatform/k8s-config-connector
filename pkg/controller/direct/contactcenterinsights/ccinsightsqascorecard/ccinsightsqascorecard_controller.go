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

package ccinsightsqascorecard

import (
	"context"
	"fmt"

	contactcenterinsights "cloud.google.com/go/contactcenterinsights/apiv1"
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
	ccimappers "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/contactcenterinsights"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CCInsightsQAScorecardGVK, newModel)
}

func newModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type adapter struct {
	id      *krm.CCInsightsQAScorecardIdentity
	desired *pb.QaScorecard
	actual  *pb.QaScorecard
	gcp     *contactcenterinsights.Client
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*contactcenterinsights.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := contactcenterinsights.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building contactcenterinsights client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.CCInsightsQAScorecard{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ccimappers.CCInsightsQAScorecardSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &adapter{
		id:      id.(*krm.CCInsightsQAScorecardIdentity),
		desired: desired,
		gcp:     gcp,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	req := &pb.GetQaScorecardRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcp.GetQaScorecard(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting qascorecard: %w", err)
	}
	a.actual = actual
	return true, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CCInsightsQAScorecard", "id", a.id)

	scorecard := proto.Clone(a.desired).(*pb.QaScorecard)
	scorecard.Name = a.id.String()

	req := &pb.CreateQaScorecardRequest{
		Parent:        a.id.ParentString(),
		QaScorecard:   scorecard,
		QaScorecardId: a.id.QaScorecard,
	}

	created, err := a.gcp.CreateQaScorecard(ctx, req)
	if err != nil {
		return fmt.Errorf("creating qascorecard %s: %w", a.id.String(), err)
	}

	log.V(2).Info("created CCInsightsQAScorecard", "id", a.id)
	return a.updateStatus(ctx, createOp, created)
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CCInsightsQAScorecard", "id", a.id)

	mapCtx := &direct.MapContext{}
	actualSpec := ccimappers.CCInsightsQAScorecardSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := ccimappers.CCInsightsQAScorecardSpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.QaScorecard)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "id", a.id)
		return nil
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	scorecard := proto.Clone(clonedDesired).(*pb.QaScorecard)
	scorecard.Name = a.id.String()

	req := &pb.UpdateQaScorecardRequest{
		QaScorecard: scorecard,
		UpdateMask:  updateMask,
	}

	updated, err := a.gcp.UpdateQaScorecard(ctx, req)
	if err != nil {
		return fmt.Errorf("updating qascorecard %s: %w", a.id.String(), err)
	}

	log.V(2).Info("updated CCInsightsQAScorecard", "id", a.id)
	return a.updateStatus(ctx, updateOp, updated)
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CCInsightsQAScorecard", "id", a.id)

	req := &pb.DeleteQaScorecardRequest{
		Name: a.id.String(),
	}

	err := a.gcp.DeleteQaScorecard(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting qascorecard %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted CCInsightsQAScorecard", "id", a.id)
	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	mapCtx := &direct.MapContext{}
	spec := ccimappers.CCInsightsQAScorecardSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.CCInsightsQAScorecard{}
	obj.Spec = *spec

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: specObj}
	u.SetGroupVersionKind(krm.CCInsightsQAScorecardGVK)
	return u, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.QaScorecard) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CCInsightsQAScorecardStatus{}
	status.ObservedState = ccimappers.CCInsightsQAScorecardObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
