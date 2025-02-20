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

package logging

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client
	gcp "cloud.google.com/go/logging/apiv2"

	// TODO(contributor): Update the import with the google cloud client api protobuf
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.LoggingLogExclusionGVK, NewLogExclusionModel)
}

func NewLogExclusionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLogExclusion{config: *config}, nil
}

var _ directbase.Model = &modelLogExclusion{}

type modelLogExclusion struct {
	config config.ControllerConfig
}

func (m *modelLogExclusion) client(ctx context.Context) (*loggingpb.ConfigServiceV2Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := loggingpb.NewConfigServiceV2Client(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("building LogExclusion client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelLogExclusion) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.LoggingLogExclusion{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewLogExclusionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get logging GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &LogExclusionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelLogExclusion) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LogExclusionAdapter struct {
	id        *krm.LogExclusionIdentity
	gcpClient *loggingpb.ConfigServiceV2Client
	desired   *krm.LoggingLogExclusion
	actual    *loggingpb.LogExclusion
}

var _ directbase.Adapter = &LogExclusionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *LogExclusionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LogExclusion", "name", a.id)

	req := &loggingpb.GetExclusionRequest{Name: a.id.String()}
	logexclusionpb, err := a.gcpClient.GetExclusion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LogExclusion %q: %w", a.id, err)
	}

	a.actual = logexclusionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LogExclusionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LogExclusion", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := LoggingLogExclusionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &loggingpb.CreateExclusionRequest{
		Parent:       a.id.Parent().String(),
		Exclusion: resource,
	}
if _, err := a.gcpClient.CreateExclusion(ctx, req); err != nil {
		return fmt.Errorf("creating LogExclusion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created LogExclusion", "name", a.id)
        status := &krm.LoggingLogExclusionStatus{}
        // TODO: Fix observed state
	//status.ObservedState = LoggingLogExclusionObservedState_FromProto(mapCtx, &loggingpb.LogExclusion{})
	//if mapCtx.Err() != nil {
	//	return mapCtx.Err()
	//}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LogExclusionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LogExclusion", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := LoggingLogExclusionSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.LoggingLogExclusionStatus{}
		status.ObservedState = LoggingLogExclusionObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &loggingpb.UpdateLogExclusionRequest{
		Name:         a.id.String(),
		UpdateMask:   updateMask,
		Exclusion: desiredPb,
	}
if _, err = a.gcpClient.UpdateExclusion(ctx, req); err != nil {
		return fmt.Errorf("updating LogExclusion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated LogExclusion", "name", a.id)

	status := &krm.LoggingLogExclusionStatus{}
        // TODO: Fix observed state
	//status.ObservedState = LoggingLogExclusionObservedState_FromProto(mapCtx, &loggingpb.LogExclusion{})
	//if mapCtx.Err() != nil {
	//	return mapCtx.Err()
	//}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *LogExclusionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogExclusion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogExclusionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLogExclusionGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *LogExclusionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LogExclusion", "name", a.id)

	req := &loggingpb.DeleteExclusionRequest{Name: a.id.String()}
if _, err := a.gcpClient.DeleteExclusion(ctx, req); err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent LogExclusion, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return fmt.Errorf("deleting LogExclusion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LogExclusion", "name", a.id)
        return true, nil
}
