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
// proto.service: google.cloud.discoveryengine.v1beta.SampleQuerySetService
// proto.message: google.cloud.discoveryengine.v1beta.SampleQuerySet
// crd.type: DiscoveryEngineSampleQuerySet
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/discoveryengine/apiv1beta"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineSampleQuerySetGVK, NewSampleQuerySetModel)
}

func NewSampleQuerySetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sampleQuerySetModel{config: *config}, nil
}

var _ directbase.Model = &sampleQuerySetModel{}

type sampleQuerySetModel struct {
	config config.ControllerConfig
}

func (m *sampleQuerySetModel) client(ctx context.Context, projectID string) (*gcp.SampleQuerySetClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewSampleQuerySetRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine samplequeryset client: %w", err)
	}

	return gcpClient, err
}

func (m *sampleQuerySetModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineSampleQuerySet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewSampleQuerySetIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineSampleQuerySetSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &sampleQuerySetAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *sampleQuerySetModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type sampleQuerySetAdapter struct {
	gcpClient *gcp.SampleQuerySetClient
	id        *krm.SampleQuerySetIdentity
	desired   *pb.SampleQuerySet
	actual    *pb.SampleQuerySet
}

var _ directbase.Adapter = &sampleQuerySetAdapter{}

func (a *sampleQuerySetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine sample query set", "name", a.id)

	req := &pb.GetSampleQuerySetRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSampleQuerySet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine sample query set %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *sampleQuerySetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine sample query set", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	req := &pb.CreateSampleQuerySetRequest{
		Parent:           a.id.Parent().String(),
		SampleQuerySet:   desired,
		SampleQuerySetId: a.id.ID(),
	}
	created, err := a.gcpClient.CreateSampleQuerySet(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine sample query set %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created discoveryengine sample query set in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *sampleQuerySetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine sample query set", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	diffs, updateMask, err := compareSampleQuerySet(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSampleQuerySetRequest{
		SampleQuerySet: desired,
		UpdateMask:     updateMask,
	}
	updated, err := a.gcpClient.UpdateSampleQuerySet(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine sample query set %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated discoveryengine sample query set", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *sampleQuerySetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineSampleQuerySet{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineSampleQuerySetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.DiscoveryEngineSampleQuerySetGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *sampleQuerySetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine sample query set", "name", a.id)

	req := &pb.DeleteSampleQuerySetRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSampleQuerySet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting discoveryengine sample query set %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine sample query set", "name", a.id)
	return true, nil
}

func (a *sampleQuerySetAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SampleQuerySet) error {
	status := &krm.DiscoveryEngineSampleQuerySetStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineSampleQuerySetObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func compareSampleQuerySet(ctx context.Context, actual, desired *pb.SampleQuerySet) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineSampleQuerySetSpec_FromProto, DiscoveryEngineSampleQuerySetSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.SampleQuerySet) {
		// Populate GCP/server defaults here if any
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
