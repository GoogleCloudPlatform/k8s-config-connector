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

// +generated:controller
// krm.group: sql.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: SQLBlueGreenDeployment

package sql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/googleapis/gax-go/v2"

	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/sql/apiv1"
	sqlpb "cloud.google.com/go/sql/apiv1/sqlpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.SQLBlueGreenDeploymentGVK, NewBlueGreenDeploymentModel)
}

func NewBlueGreenDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelBlueGreenDeployment{config: *config}, nil
}

var _ directbase.Model = &modelBlueGreenDeployment{}

type modelBlueGreenDeployment struct {
	config config.ControllerConfig
}

func (m *modelBlueGreenDeployment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SQLBlueGreenDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.SQLBlueGreenDeploymentIdentity)

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if obj.Spec.SourceInstanceRef != nil {
		sqlInstance, err := refsv1beta1.ResolveSQLInstanceRef(ctx, reader, obj, obj.Spec.SourceInstanceRef)
		if err != nil {
			return nil, fmt.Errorf("resolving spec.sourceInstanceRef: %w", err)
		}
		if sqlInstance != nil {
			obj.Spec.SourceInstanceRef.External = fmt.Sprintf("projects/%s/instances/%s", sqlInstance.ProjectID, sqlInstance.SQLInstanceName)
		}
	}

	// Get sql GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	operationsClient, err := m.operationsClient(ctx)
	if err != nil {
		return nil, err
	}
	return &BlueGreenDeploymentAdapter{
		id:               id,
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		desired:          obj,
	}, nil
}

func (m *modelBlueGreenDeployment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//sqladmin.googleapis.com/") {
		return nil, nil
	}
	id := &krm.SQLBlueGreenDeploymentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	operationsClient, err := m.operationsClient(ctx)
	if err != nil {
		return nil, err
	}
	return &BlueGreenDeploymentAdapter{
		id:               id,
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
	}, nil
}
func (m *modelBlueGreenDeployment) client(ctx context.Context) (*gcp.BlueGreenDeploymentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewBlueGreenDeploymentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BlueGreenDeployment client: %w", err)
	}
	return gcpClient, err
}

func (m *modelBlueGreenDeployment) operationsClient(ctx context.Context) (*gcp.SqlOperationsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	operationsClient, err := gcp.NewSqlOperationsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SqlOperations client: %w", err)
	}
	return operationsClient, nil
}

type BlueGreenDeploymentAdapter struct {
	id               *krm.SQLBlueGreenDeploymentIdentity
	gcpClient        *gcp.BlueGreenDeploymentsClient
	operationsClient *gcp.SqlOperationsClient
	desired          *krm.SQLBlueGreenDeployment
	actual           *sqlpb.BlueGreenDeployment
}

var _ directbase.Adapter = &BlueGreenDeploymentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *BlueGreenDeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BlueGreenDeployment", "name", a.id)

	req := &sqlpb.GetBlueGreenDeploymentRequest{
		Name: a.id.String(),
	}
	bluegreendeploymentpb, err := a.gcpClient.GetBlueGreenDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BlueGreenDeployment %q: %w", a.id, err)
	}

	a.actual = bluegreendeploymentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BlueGreenDeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BlueGreenDeployment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SQLBlueGreenDeploymentSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &sqlpb.CreateBlueGreenDeploymentRequest{
		Parent:                a.id.ParentString(),
		BlueGreenDeploymentId: a.id.BlueGreenDeployment,
		BlueGreenDeployment:   resource,
	}
	op, err := a.gcpClient.CreateBlueGreenDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating BlueGreenDeployment %s: %w", a.id, err)
	}
	if _, err := a.pollOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for BlueGreenDeployment %s creation: %w", a.id, err)
	}

	created, err := a.gcpClient.GetBlueGreenDeployment(ctx, &sqlpb.GetBlueGreenDeploymentRequest{
		Name: a.id.String(),
	})
	if err != nil {
		return fmt.Errorf("getting newly created BlueGreenDeployment %s: %w", a.id, err)
	}
	a.actual = created
	log.V(2).Info("successfully created BlueGreenDeployment", "name", a.id)

	status := &krm.SQLBlueGreenDeploymentStatus{}
	status.ObservedState = SQLBlueGreenDeploymentObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *BlueGreenDeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BlueGreenDeployment", "name", a.id)

	normalize := func(ctx context.Context, pb *sqlpb.BlueGreenDeployment) error {
		if pb.GetSourceInstance() != "" {
			tokens := strings.Split(pb.GetSourceInstance(), "/")
			pb.SourceInstance = tokens[len(tokens)-1]
		}
		return nil
	}

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		SQLBlueGreenDeploymentSpec_v1alpha1_FromProto,
		SQLBlueGreenDeploymentSpec_v1alpha1_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating SQLBlueGreenDeployment is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.SQLBlueGreenDeploymentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = SQLBlueGreenDeploymentObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *BlueGreenDeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SQLBlueGreenDeployment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SQLBlueGreenDeploymentSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}
	if obj.Spec.SourceInstanceRef != nil && !strings.Contains(obj.Spec.SourceInstanceRef.External, "/") {
		obj.Spec.SourceInstanceRef.External = fmt.Sprintf("projects/%s/instances/%s", a.id.Project, obj.Spec.SourceInstanceRef.External)
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.SQLBlueGreenDeploymentGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *BlueGreenDeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BlueGreenDeployment", "name", a.id)

	req := &sqlpb.DeleteBlueGreenDeploymentRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteBlueGreenDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent BlueGreenDeployment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BlueGreenDeployment %s: %w", a.id, err)
	}
	if _, err := a.pollOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for BlueGreenDeployment %s deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BlueGreenDeployment", "name", a.id)
	return true, nil
}

func (a *BlueGreenDeploymentAdapter) pollOperation(ctx context.Context, op *sqlpb.Operation) (*sqlpb.Operation, error) {
	if op == nil {
		return nil, nil
	}
	log := klog.FromContext(ctx)
	var err error

	pollingBackoff := gax.Backoff{
		Initial:    2 * time.Second,
		Max:        30 * time.Second,
		Multiplier: 1.5,
	}

	for {
		log.V(2).Info("polling operation", "op", op.GetName(), "status", op.GetStatus())

		if op.GetStatus() == sqlpb.Operation_DONE {
			if opError := op.GetError(); opError != nil && len(opError.GetErrors()) > 0 {
				var errMsgs []string
				for _, e := range opError.GetErrors() {
					errMsgs = append(errMsgs, e.GetMessage())
				}
				return op, fmt.Errorf("operation failed: %s", strings.Join(errMsgs, "; "))
			}
			return op, nil
		}

		if err := gax.Sleep(ctx, pollingBackoff.Pause()); err != nil {
			return op, fmt.Errorf("waiting for operation %s: %w", op.GetName(), err)
		}

		op, err = a.operationsClient.Get(ctx, &sqlpb.SqlOperationsGetRequest{
			Project:   a.id.Project,
			Operation: op.GetName(),
			Location:  a.id.Location,
		})
		if err != nil {
			return op, fmt.Errorf("getting operation %s: %w", op.GetName(), err)
		}
	}
}
