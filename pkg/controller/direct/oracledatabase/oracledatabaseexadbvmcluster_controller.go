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

package oracledatabase

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/oracledatabase/apiv1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.OracleDatabaseExadbVMClusterGVK, NewExadbVMClusterModel)
}

func NewExadbVMClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelExadbVMCluster{config: *config}, nil
}

var _ directbase.Model = &modelExadbVMCluster{}

type modelExadbVMCluster struct {
	config config.ControllerConfig
}

func (m *modelExadbVMCluster) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ExadbVMCluster REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelExadbVMCluster) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.OracleDatabaseExadbVMCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := OracleDatabaseExadbVMClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ExadbVMClusterAdapter{
		id:        id.(*krm.OracleDatabaseExadbVMClusterIdentity),
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelExadbVMCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.OracleDatabaseExadbVMClusterIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ExadbVMClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ExadbVMClusterAdapter struct {
	id        *krm.OracleDatabaseExadbVMClusterIdentity
	gcpClient *gcp.Client
	desired   *pb.ExadbVmCluster
	actual    *pb.ExadbVmCluster
}

var _ directbase.Adapter = &ExadbVMClusterAdapter{}

func (a *ExadbVMClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ExadbVMCluster", "name", a.id)

	req := &pb.GetExadbVmClusterRequest{Name: a.id.String()}
	exadbVMClusterpb, err := a.gcpClient.GetExadbVmCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ExadbVMCluster %q: %w", a.id, err)
	}

	a.actual = exadbVMClusterpb
	return true, nil
}

func (a *ExadbVMClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ExadbVMCluster", "name", a.id)

	clonedDesired := proto.Clone(a.desired).(*pb.ExadbVmCluster)

	req := &pb.CreateExadbVmClusterRequest{
		Parent:           a.id.ParentString(),
		ExadbVmClusterId: a.id.ExadbVmCluster,
		ExadbVmCluster:   clonedDesired,
	}
	op, err := a.gcpClient.CreateExadbVmCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ExadbVMCluster %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("ExadbVMCluster %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created ExadbVMCluster", "name", a.id)

	// Fetch full resource state from GCP immediately after LRO success to avoid empty fields in status
	latest, err := a.gcpClient.GetExadbVmCluster(ctx, &pb.GetExadbVmClusterRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting ExadbVMCluster %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *ExadbVMClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ExadbVMCluster", "name", a.id)

	clonedDesired := proto.Clone(a.desired).(*pb.ExadbVmCluster)
	clonedDesired.Name = a.id.String()

	diffs, updateMask, err := compareResource(ctx, a.actual, clonedDesired)
	if err != nil {
		return err
	}

	latest := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", updateMask.Paths)
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdateExadbVmClusterRequest{
			UpdateMask:     updateMask,
			ExadbVmCluster: clonedDesired,
		}
		op, err := a.gcpClient.UpdateExadbVmCluster(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ExadbVMCluster %s: %w", a.id, err)
		}
		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("ExadbVMCluster %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ExadbVMCluster", "name", a.id)

		// Fetch full resource state from GCP immediately after LRO success to avoid empty fields in status
		latest, err = a.gcpClient.GetExadbVmCluster(ctx, &pb.GetExadbVmClusterRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting ExadbVMCluster %s after update: %w", a.id, err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareResource(ctx context.Context, actual, desired *pb.ExadbVmCluster) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := OracleDatabaseExadbVMClusterSpec_FromProto(mapCtx, actual)
	maskedActual := OracleDatabaseExadbVMClusterSpec_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.ExadbVmCluster)

	populateDefaults := func(obj *pb.ExadbVmCluster) {
		// Populate GCP/server defaults here if needed
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *ExadbVMClusterAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ExadbVmCluster) error {
	mapCtx := &direct.MapContext{}
	status := &krm.OracleDatabaseExadbVMClusterStatus{}
	status.ObservedState = OracleDatabaseExadbVMClusterObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ExadbVMClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.OracleDatabaseExadbVMCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(OracleDatabaseExadbVMClusterSpec_FromProto(mapCtx, a.actual))
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
	u.SetName(a.id.ExadbVmCluster)
	u.SetGroupVersionKind(krm.OracleDatabaseExadbVMClusterGVK)

	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

func (a *ExadbVMClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ExadbVMCluster", "name", a.id)

	req := &pb.DeleteExadbVmClusterRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteExadbVmCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ExadbVMCluster, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ExadbVMCluster %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ExadbVMCluster %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted ExadbVMCluster", "name", a.id)
	return true, nil
}
