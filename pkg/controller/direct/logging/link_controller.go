// Copyright 2024 Google LLC
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

package logging

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(user): Update the import with the google cloud client
	gcp "cloud.google.com/go/logging/apiv1"

	// TODO(user): Update the import with the google cloud client api protobuf
	loggingpb "cloud.google.com/go/logging/v2/loggingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.LoggingLinkGVK, NewLoggingLinkModel)
}

func NewLoggingLinkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLink{config: *config}, nil
}

var _ directbase.Model = &modelLoggingLink{}

type modelLoggingLink struct {
	config config.ControllerConfig
}

func (m *modelLoggingLink) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Link client: %w", err)
	}
	return gcpClient, err
}

func (m *modelLoggingLink) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.LoggingLink{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewLoggingLinkRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get logging GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &LoggingLinkAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelLoggingLink) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LoggingLinkAdapter struct {
	id        *krm.LoggingLinkRef
	gcpClient *gcp.Client
	desired   *krm.LoggingLink
	actual    *loggingpb.Link
}

var _ directbase.Adapter = &LoggingLinkAdapter{}

func (a *LoggingLinkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLink", "name", a.id.External)

	req := &loggingpb.GetLinkRequest{Name: a.id.External}
	linkpb, err := a.gcpClient.GetLink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLink %q: %w", a.id.External, err)
	}

	a.actual = linkpb
	return true, nil
}

func (a *LoggingLinkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Link", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LoggingLinkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Complete the gcp "CREATE" or "INSERT" request with required fields.
	parent, err := a.id.Parent()
	if err != nil {
		return err
	}
	req := &loggingpb.CreateLinkRequest{
		Parent: parent.String(),
		Link:   resource,
	}
	op, err := a.gcpClient.CreateLink(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Link %s: %w", a.id.External, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Link %s waiting creation: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created Link", "name", a.id.External)

	status := &krm.LoggingLinkStatus{}
	status.ObservedState = LoggingLinkStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLinkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Link", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LoggingLinkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Update the field if applicable.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		return nil
	}
	// TODO(user): Complete the gcp "UPDATE" or "PATCH" request with required fields.
	req := &loggingpb.UpdateLinkRequest{
		Name:       a.id.External,
		UpdateMask: updateMask,
		Link:       resource,
	}
	op, err := a.gcpClient.UpdateLink(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Link %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Link %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated Link", "name", a.id.External)

	status := &krm.LoggingLinkStatus{}
	status.ObservedState = LoggingLinkStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLinkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLink{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLinkSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// TODO(user): Update other resource references
	parent, err := a.id.Parent()
	if err != nil {
		return nil, err
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: parent.String()}
	obj.Spec.Location = parent.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Id)
	u.SetGroupVersionKind(krm.LoggingLinkGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *LoggingLinkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Link", "name", a.id.External)

	req := &loggingpb.DeleteLinkRequest{Name: a.id.External}
	op, err := a.gcpClient.DeleteLink(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Link %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully deleted Link", "name", a.id.External)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Link %s: %w", a.id.External, err)
	}
	return true, nil
}
