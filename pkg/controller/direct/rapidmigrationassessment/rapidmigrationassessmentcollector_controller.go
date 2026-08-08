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

package rapidmigrationassessment

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/rapidmigrationassessment/apiv1"
	pb "cloud.google.com/go/rapidmigrationassessment/apiv1/rapidmigrationassessmentpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/rapidmigrationassessment/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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
	registry.RegisterModel(krm.RapidMigrationAssessmentCollectorGVK, NewModel)
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
		return nil, fmt.Errorf("building rapidmigrationassessment client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.RapidMigrationAssessmentCollector{}
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
	collectorIdentity := identity.(*krm.RapidMigrationAssessmentCollectorIdentity)
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        collectorIdentity,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.RapidMigrationAssessmentCollectorIdentity
	gcpClient *gcp.Client
	desired   *krm.RapidMigrationAssessmentCollector
	actual    *pb.Collector
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if a.id.Collector == "" { // resource is not yet created
		return false, nil
	}
	fqn := a.id.String()
	log.V(2).Info("getting RapidMigrationAssessmentCollector", "name", fqn)

	req := &pb.GetCollectorRequest{Name: fqn}
	collectorpb, err := a.gcpClient.GetCollector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting RapidMigrationAssessmentCollector %q: %w", fqn, err)
	}

	a.actual = collectorpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating RapidMigrationAssessmentCollector", "name", fqn)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := RapidMigrationAssessmentCollectorSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	populateProtoDefaults(resource)

	parent := a.id.ParentString()
	req := &pb.CreateCollectorRequest{
		Parent:      parent,
		CollectorId: a.id.Collector,
		Collector:   resource,
	}
	op, err := a.gcpClient.CreateCollector(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Collector %s: %w", fqn, err)
	}
	_, err = op.Wait(ctx)
	if err != nil && !isProtoResolutionError(err) {
		return fmt.Errorf("waiting Collector %s creation: %w", fqn, err)
	}

	// Always perform GET operation after Create to fetch the fully-populated resource
	latest, err := a.gcpClient.GetCollector(ctx, &pb.GetCollectorRequest{Name: fqn})
	if err != nil {
		return fmt.Errorf("getting Collector %s after creation: %w", fqn, err)
	}

	log.V(2).Info("successfully created Collector", "name", latest.Name)

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating RapidMigrationAssessmentCollector", "name", fqn)

	mapCtx := &direct.MapContext{}
	desiredPb := RapidMigrationAssessmentCollectorSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	populateProtoDefaults(desiredPb)

	diffs, updateMask, err := compareCollector(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", fqn)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = u
	structuredreporting.ReportDiff(ctx, diffs)

	desiredPb.Name = a.actual.Name
	req := &pb.UpdateCollectorRequest{
		Collector:  desiredPb,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateCollector(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Collector %s: %w", fqn, err)
	}
	_, err = op.Wait(ctx)
	if err != nil && !isProtoResolutionError(err) {
		return fmt.Errorf("waiting Collector %s update: %w", fqn, err)
	}

	// Always perform GET operation after Update to fetch the fully-populated resource
	latest, err := a.gcpClient.GetCollector(ctx, &pb.GetCollectorRequest{Name: fqn})
	if err != nil {
		return fmt.Errorf("getting Collector %s after update: %w", fqn, err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting RapidMigrationAssessmentCollector", "name", fqn)

	req := &pb.DeleteCollectorRequest{Name: fqn}
	op, err := a.gcpClient.DeleteCollector(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting RapidMigrationAssessmentCollector %q: %w", fqn, err)
	}

	_, err = op.Wait(ctx)
	if err != nil && !isProtoResolutionError(err) {
		return false, fmt.Errorf("waiting RapidMigrationAssessmentCollector %q deletion: %w", fqn, err)
	}

	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("collector %q not found", a.id.String())
	}

	mapCtx := &direct.MapContext{}
	spec := RapidMigrationAssessmentCollectorSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// ProjectRef and Location are not read-back from GCP but we can populate them from our identity
	spec.ProjectRef = &refsv1beta1.ProjectRef{Name: a.id.Project}
	spec.Location = &a.id.Location

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting collector spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.Collector)
	u.SetGroupVersionKind(krm.RapidMigrationAssessmentCollectorGVK)
	u.SetLabels(a.actual.Labels)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Collector) error {
	mapCtx := &direct.MapContext{}
	status := &krm.RapidMigrationAssessmentCollectorStatus{}
	status.ObservedState = RapidMigrationAssessmentCollectorObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}

func compareCollector(ctx context.Context, actual, desired *pb.Collector) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, RapidMigrationAssessmentCollectorSpec_FromProto, RapidMigrationAssessmentCollectorSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore identifier field

	clonedDesired := proto.Clone(desired).(*pb.Collector)

	populateProtoDefaults(maskedActual)
	populateProtoDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func populateProtoDefaults(obj *pb.Collector) {
	if obj.ExpectedAssetCount == 0 {
		obj.ExpectedAssetCount = 100
	}
	if obj.CollectionDays == 0 {
		obj.CollectionDays = 30
	}
}

func isProtoResolutionError(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "unable to resolve") || strings.Contains(s, "v1main")
}
