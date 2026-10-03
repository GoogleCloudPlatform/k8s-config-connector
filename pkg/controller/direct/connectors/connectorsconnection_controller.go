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

package connectors

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	api "google.golang.org/api/connectors/v1"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connectors/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/connectors/pb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ConnectorsConnectionGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Connectors REST client: %w", err)
	}

	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConnectorsConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	if obj.Spec.ServiceAccountRef != nil {
		if err := obj.Spec.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return nil, fmt.Errorf("resolving serviceAccountRef: %w", err)
		}
	}
	mapCtx := &direct.MapContext{}
	desired := ConnectorsConnectionSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        id.(*krm.ConnectorsConnectionIdentity),
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type adapter struct {
	gcpClient *api.Service
	id        *krm.ConnectorsConnectionIdentity
	desired   *pb.Connection
	actual    *pb.Connection

	reader client.Reader
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	if a.id.Connection == "" {
		return false, nil
	}

	restObj, err := a.gcpClient.Projects.Locations.Connections.Get(a.id.String()).View("FULL").Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ConnectorsConnection %q from gcp: %w", a.id.String(), err)
	}

	actual, err := RESTToProto(restObj)
	if err != nil {
		return false, err
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating ConnectorsConnection", "name", a.id)

	a.desired.Name = a.id.String()
	restObj, err := ProtoToREST(a.desired)
	if err != nil {
		return err
	}

	op, err := a.gcpClient.Projects.Locations.Connections.Create(a.id.ParentString(), restObj).ConnectionId(a.id.Connection).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating ConnectorsConnection %s: %w", a.id.String(), err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for ConnectorsConnection %s creation: %w", a.id.String(), err)
	}

	log.Info("successfully created ConnectorsConnection in gcp", "name", a.id)

	// Fetch the fully populated object after LRO completes and wait for it to finish initialization
	latestREST, err := a.waitForConnectionReady(ctx)
	if err != nil {
		return fmt.Errorf("fetching created ConnectorsConnection %s: %w", a.id.String(), err)
	}

	created, err := RESTToProto(latestREST)
	if err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating ConnectorsConnection", "name", a.id)

	desired := proto.Clone(a.desired).(*pb.Connection)
	desired.Name = a.id.String()

	diffs, updateMask, err := a.compare(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.Info("no diff detected, skipping update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	restObj, err := ProtoToREST(desired)
	if err != nil {
		return err
	}

	maskStr := strings.Join(updateMask.GetPaths(), ",")
	op, err := a.gcpClient.Projects.Locations.Connections.Patch(a.id.String(), restObj).UpdateMask(maskStr).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating ConnectorsConnection %s: %w", a.id.String(), err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for ConnectorsConnection %s update: %w", a.id.String(), err)
	}

	log.Info("successfully updated ConnectorsConnection", "name", a.id)

	// Fetch the fully populated object after LRO completes and wait for it to finish updating
	latestREST, err := a.waitForConnectionReady(ctx)
	if err != nil {
		return fmt.Errorf("fetching updated ConnectorsConnection %s: %w", a.id.String(), err)
	}

	updated, err := RESTToProto(latestREST)
	if err != nil {
		return err
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting ConnectorsConnection", "name", a.id)

	op, err := a.gcpClient.Projects.Locations.Connections.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting ConnectorsConnection %s: %w", a.id.String(), err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting for ConnectorsConnection %s deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConnectorsConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConnectorsConnectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Connection)
	u.SetGroupVersionKind(krm.ConnectorsConnectionGVK)

	u.Object = uObj
	return u, nil
}

func (a *adapter) waitForOperation(ctx context.Context, op *api.Operation) error {
	if op.Done {
		if op.Error != nil {
			return fmt.Errorf("operation completed with error: %s", op.Error.Message)
		}
		return nil
	}
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
		latest, err := a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			if direct.IsNotFound(err) {
				return nil
			}
			return fmt.Errorf("getting operation %q: %w", op.Name, err)
		}
		if latest.Done {
			if latest.Error != nil {
				return fmt.Errorf("operation completed with error: %s", latest.Error.Message)
			}
			return nil
		}
	}
}

func (a *adapter) waitForConnectionReady(ctx context.Context) (*api.Connection, error) {
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		conn, err := a.gcpClient.Projects.Locations.Connections.Get(a.id.String()).View("FULL").Context(ctx).Do()
		if err != nil {
			return nil, fmt.Errorf("fetching ConnectorsConnection %s: %w", a.id.String(), err)
		}
		if conn.Status == nil || (conn.Status.State != "CREATING" && conn.Status.State != "UPDATING") {
			return conn, nil
		}
		time.Sleep(2 * time.Second)
	}
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Connection) error {
	status := &krm.ConnectorsConnectionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ConnectorsConnectionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) compare(ctx context.Context, actual, desired *pb.Connection) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	clonedActual := proto.Clone(actual).(*pb.Connection)
	clonedDesired := proto.Clone(desired).(*pb.Connection)

	maskedActual, err := mappers.OnlySpecFields(clonedActual, ConnectorsConnectionSpec_FromProto, ConnectorsConnectionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = clonedDesired.Name

	if clonedDesired.NodeConfig == nil && maskedActual.NodeConfig != nil {
		clonedDesired.NodeConfig = maskedActual.NodeConfig
	}
	if clonedDesired.LockConfig == nil && maskedActual.LockConfig != nil {
		clonedDesired.LockConfig = maskedActual.LockConfig
	}
	if clonedDesired.AuthConfig == nil && maskedActual.AuthConfig != nil {
		clonedDesired.AuthConfig = maskedActual.AuthConfig
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func RESTToProto(in *api.Connection) (*pb.Connection, error) {
	if in == nil {
		return nil, nil
	}
	bytes, err := in.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling REST struct: %w", err)
	}
	out := &pb.Connection{}
	unmarshalOpts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
	if err := unmarshalOpts.Unmarshal(bytes, out); err != nil {
		return nil, fmt.Errorf("error unmarshaling proto struct: %w", err)
	}
	return out, nil
}

func ProtoToREST(in *pb.Connection) (*api.Connection, error) {
	if in == nil {
		return nil, nil
	}
	marshalOpts := protojson.MarshalOptions{
		UseProtoNames:   false,
		EmitUnpopulated: false,
	}
	bytes, err := marshalOpts.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("error marshaling proto struct: %w", err)
	}
	out := &api.Connection{}
	if err := json.Unmarshal(bytes, out); err != nil {
		return nil, fmt.Errorf("error unmarshaling REST struct: %w", err)
	}
	return out, nil
}
