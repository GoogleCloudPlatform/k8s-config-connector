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

	//"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	//"net/http"

	gcp "cloud.google.com/go/logging/apiv2"
	//"google.golang.org/api/option"
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"

	//"google.golang.org/api/option"
	//"google.golang.org/protobuf/types/known/fieldmaskpb"

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

func (m *modelLoggingLink) client(ctx context.Context) (*gcp.ConfigClient, error) {

	gcpClient, err := gcp.NewConfigRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building Logging Config client: %w", err)
	}
	return gcpClient, err
}

func (m *modelLoggingLink) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.LoggingLink{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	linkIdentity, err := krm.NewLinkIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &LoggingLinkAdapter{
		id:        linkIdentity,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

/*
func () ResolveExternalRef(externalRef string) {

}
*/

func (m *modelLoggingLink) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LoggingLinkAdapter struct {
	id        *krm.LinkIdentity
	gcpClient *gcp.ConfigClient
	desired   *krm.LoggingLink
	actual    *loggingpb.Link
}

var _ directbase.Adapter = &LoggingLinkAdapter{}

func (a *LoggingLinkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLink", "name", a.id)

	req := &loggingpb.GetLinkRequest{Name: a.id.String()}
	linkpb, err := a.gcpClient.GetLink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLink %q: %w", a.id, err)
	}

	a.actual = linkpb
	return true, nil
}

func (a *LoggingLinkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating LoggingLink", "u", u)

	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LoggingLinkSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent()

	resourceID := direct.ValueOf(desired.Spec.ResourceID)
	if resourceID == "" {
		log.V(2).Info("ResourceID is not set, will use metadata.name")
		resourceID = desired.Name
	} else {
		log.V(2).Info("ResourceID is set, will use")
		resourceID = *desired.Spec.ResourceID
	}

	req := &loggingpb.CreateLinkRequest{
		Parent: parent.String(),
		Link:   resource,
		LinkId: resourceID,
	}
	op, err := a.gcpClient.CreateLink(ctx, req)

	if err != nil {
		return fmt.Errorf("creating Link %s: %w\n", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		// there is an implicit dependency on the bucket being active, this wait captures that
		return fmt.Errorf("Link %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Link", "name", a.id)

	status := &krm.LoggingLinkStatus{}
	status.ObservedState = LoggingLinkObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLinkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {

	log := klog.FromContext(ctx)
	log.V(2).Info("updating Logging Link", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := LoggingLinkSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.LoggingLinkStatus{}
		status.ObservedState = LoggingLinkObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	} else {
		return fmt.Errorf("update operation not supported for resource %v %v",
			a.desired.GroupVersionKind(), k8s.GetNamespacedName(a.desired))
	}
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
	/* TODO Bucket Ref?
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
	*/

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLinkGVK)

	// TODO uObj is set in the commented out code
	//u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *LoggingLinkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Link", "name", a.id)

	req := &loggingpb.DeleteLinkRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteLink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent LoggingLink, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting Link %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Link", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Link %s: %w", a.id, err)
	}
	return true, nil
}
