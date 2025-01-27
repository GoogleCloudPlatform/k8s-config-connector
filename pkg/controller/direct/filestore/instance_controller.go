// Copyright 2025 Google LLC
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

package filestore

// +tool:controller
// proto.service: google.cloud.filestore.v1.CloudFilestoreManager
// proto.message: google.cloud.filestore.v1.Instance
// crd.type: FilestoreInstance
// crd.version: v1alpha1

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	api "cloud.google.com/go/filestore/apiv1"
	pb "cloud.google.com/go/filestore/apiv1/filestorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.FilestoreInstanceGVK, NewInstanceModel)
}

func NewInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &instanceModel{config: *config}, nil
}

var _ directbase.Model = &instanceModel{}

type instanceModel struct {
	config config.ControllerConfig
}

func (m *instanceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.FilestoreInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.InstanceIdentityForObject(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := FilestoreInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newCloudFilestoreManagerClient(ctx)
	if err != nil {
		return nil, err
	}

	return &instanceAdapter{
		client:  client,
		id:      id,
		desired: desired,
	}, nil
}

func (m *instanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//file.googleapis.com/") {
		id, err := krm.ParseInstanceIdentityExternal(url)
		if err != nil {
			log.V(2).Error(err, "url did not match FilestoreInstance format", "url", url)
		} else {

			gcpClient, err := newGCPClient(ctx, &m.config)
			if err != nil {
				return nil, err
			}
			client, err := gcpClient.newCloudFilestoreManagerClient(ctx)
			if err != nil {
				return nil, err
			}

			return &instanceAdapter{
				client: client,
				id:     id,
			}, nil
		}
	}
	return nil, nil
}

type instanceAdapter struct {
	client  *api.CloudFilestoreManagerClient
	id      *krm.InstanceIdentity
	desired *pb.Instance
	actual  *pb.Instance
}

var _ directbase.Adapter = &instanceAdapter{}

func (a *instanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting filestore instance", "name", a.id)

	req := &pb.GetInstanceRequest{Name: a.id.FullyQualifiedName()}
	actual, err := a.client.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting filestore instance %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *instanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating filestore instance", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	// We don't set name for CreateInstance, it is inferred

	req := &pb.CreateInstanceRequest{
		Parent:     a.id.ProjectAndLocationParent.String(),
		Instance:   desired,
		InstanceId: a.id.Instance,
	}
	op, err := a.client.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating filestore instance %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of filestore instance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created filestore instance in gcp", "name", a.id)

	status := &krm.FilestoreInstanceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = FilestoreInstanceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *instanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating filestore instance", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	// TODO: Which fields are updateable?
	if !reflect.DeepEqual(a.desired.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	req := &pb.UpdateInstanceRequest{
		UpdateMask: updateMask,
		Instance:   desired,
	}
	op, err := a.client.UpdateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("updating filestore instance %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated filestore instance", "name", a.id)

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of filestore instance %s: %w", a.id, err)
	}

	status := &krm.FilestoreInstanceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = FilestoreInstanceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *instanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.FilestoreInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FilestoreInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ProjectID}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Instance)
	u.SetGroupVersionKind(krm.FilestoreInstanceGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *instanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting filestore instance", "name", a.id)

	req := &pb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.client.DeleteInstance(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting filestore instance %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted filestore instance", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of filestore instance %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
