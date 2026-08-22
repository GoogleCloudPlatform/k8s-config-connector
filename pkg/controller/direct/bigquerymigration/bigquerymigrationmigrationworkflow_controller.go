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

package bigquerymigration

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/bigquery/migration/apiv2alpha"
	pb "cloud.google.com/go/bigquery/migration/apiv2alpha/migrationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerymigration/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName      = "bigquerymigration-controller"
	serviceDomain = "//bigquerymigration.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.BigQueryMigrationMigrationWorkflowGVK, NewModel)
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
		return nil, fmt.Errorf("building bigquerymigration client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryMigrationMigrationWorkflow{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// NormalizeReferences
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	workflowIdentity := identity.(*krm.BigQueryMigrationMigrationWorkflowIdentity)

	// Get bigquerymigration GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := BigQueryMigrationMigrationWorkflowSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:        workflowIdentity,
		gcpClient: gcpClient,
		desired:   desired,
		reader:    reader,
		namespace: obj.Namespace,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.BigQueryMigrationMigrationWorkflowIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type Adapter struct {
	id        *krm.BigQueryMigrationMigrationWorkflowIdentity
	gcpClient *gcp.Client
	desired   *pb.MigrationWorkflow
	actual    *pb.MigrationWorkflow
	reader    client.Reader
	namespace string
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	fqn := a.id.String()
	log.V(2).Info("getting BigQueryMigrationMigrationWorkflow", "name", fqn)

	req := &pb.GetMigrationWorkflowRequest{Name: fqn}
	workflow, err := a.gcpClient.GetMigrationWorkflow(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryMigrationMigrationWorkflow %q: %w", fqn, err)
	}

	a.actual = workflow
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	parent := a.id.ParentString()
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigQueryMigrationMigrationWorkflow", "name", fqn)

	req := &pb.CreateMigrationWorkflowRequest{
		Parent:            parent,
		MigrationWorkflow: a.desired,
	}

	created, err := a.gcpClient.CreateMigrationWorkflow(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BigQueryMigrationMigrationWorkflow %s: %w", fqn, err)
	}
	log.V(2).Info("successfully created BigQueryMigrationMigrationWorkflow", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigQueryMigrationMigrationWorkflow (checking for changes)", "name", fqn)

	diffs, _, err := compareResource(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("BigQueryMigrationMigrationWorkflow is immutable and cannot be updated")
	}

	log.V(2).Info("no field needs update", "name", fqn)
	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting BigQueryMigrationMigrationWorkflow", "name", fqn)

	req := &pb.DeleteMigrationWorkflowRequest{Name: fqn}
	if err := a.gcpClient.DeleteMigrationWorkflow(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting BigQueryMigrationMigrationWorkflow %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted BigQueryMigrationMigrationWorkflow", "name", fqn)
	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryMigrationMigrationWorkflow{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryMigrationMigrationWorkflowSpec_FromProto(mapCtx, a.actual))

	// Populate ResourceID, ProjectRef, Location
	tokens := strings.Split(a.id.String(), "workflows/")
	obj.Spec.ResourceID = direct.LazyPtr(tokens[len(tokens)-1])
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent := a.id.ParentString()
	if parent != "" {
		tokens := strings.Split(parent, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
			obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{Name: tokens[1]}
			obj.Spec.Location = tokens[3]
		}
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MigrationWorkflow) error {
	mapCtx := &direct.MapContext{}
	status := BigQueryMigrationMigrationWorkflowObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	tokens := strings.Split(latest.Name, "/")
	externalRef := a.id.ParentString() + "/workflows/" + tokens[len(tokens)-1]

	krmStatus := &krm.BigQueryMigrationMigrationWorkflowStatus{
		ObservedState: status,
		ExternalRef:   &externalRef,
	}

	return op.UpdateStatus(ctx, krmStatus, nil)
}

func compareResource(ctx context.Context, actual, desired *pb.MigrationWorkflow) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BigQueryMigrationMigrationWorkflowSpec_FromProto, BigQueryMigrationMigrationWorkflowSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
