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

package networksecurity

import (
	"context"
	"fmt"
	"reflect"
	"time"

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecuritySACRealmGVK, NewSACRealmModel)
}

func NewSACRealmModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sacRealmModel{config: *config}, nil
}

var _ directbase.Model = &sacRealmModel{}

type sacRealmModel struct {
	config config.ControllerConfig
}

func (m *sacRealmModel) client(ctx context.Context) (pb.SSERealmServiceClient, longrunningpb.OperationsClient, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}

	opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing networksecurity service: %w", err)
	}

	return pb.NewSSERealmServiceClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

func (m *sacRealmModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecuritySACRealm{}
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
	mapCtx := &direct.MapContext{}
	desired := NetworkSecuritySACRealmSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &sacRealmAdapter{
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		id:               id.(*krm.NetworkSecuritySACRealmIdentity),
		desired:          desired,
	}, nil
}

func (m *sacRealmModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type sacRealmAdapter struct {
	gcpClient        pb.SSERealmServiceClient
	operationsClient longrunningpb.OperationsClient
	id               *krm.NetworkSecuritySACRealmIdentity
	desired          *pb.SACRealm
	actual           *pb.SACRealm
}

var _ directbase.Adapter = &sacRealmAdapter{}

func (a *sacRealmAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting networksecurity sac realm", "name", a.id)

	req := &pb.GetSACRealmRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSACRealm(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity sac realm %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *sacRealmAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating networksecurity sac realm", "name", a.id)

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateSACRealmRequest{
		Parent:     parent,
		SacRealmId: a.id.Sac_realm,
		SacRealm:   a.desired,
	}
	op, err := a.gcpClient.CreateSACRealm(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networksecurity sac realm %s: %w", a.id.String(), err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("networksecurity sac realm %s waiting for creation: %w", a.id.String(), err)
	}

	actual, err := a.gcpClient.GetSACRealm(ctx, &pb.GetSACRealmRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting networksecurity sac realm after creation: %w", err)
	}

	log.V(2).Info("successfully created networksecurity sac realm", "name", a.id.String())

	return a.updateStatus(ctx, createOp, actual)
}

func (a *sacRealmAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating networksecurity sac realm", "name", a.id)

	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) ||
		a.desired.SecurityService != a.actual.SecurityService {
		return fmt.Errorf("NetworkSecuritySACRealm is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *sacRealmAdapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.SACRealm) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecuritySACRealmStatus{}
	status.ObservedState = NetworkSecuritySACRealmObservedState_v1alpha1_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *sacRealmAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *sacRealmAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting networksecurity sac realm", "name", a.id)

	req := &pb.DeleteSACRealmRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteSACRealm(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting networksecurity sac realm %s: %w", a.id.String(), err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("networksecurity sac realm %s waiting for deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *sacRealmAdapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	if op.Done {
		if op.GetError() != nil {
			return fmt.Errorf("operation failed: %s", op.GetError().GetMessage())
		}
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
			current, err := a.operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
			if err != nil {
				return fmt.Errorf("getting operation %q: %w", op.Name, err)
			}
			if current.Done {
				if current.GetError() != nil {
					return fmt.Errorf("operation failed: %s", current.GetError().GetMessage())
				}
				return nil
			}
		}
	}
}
