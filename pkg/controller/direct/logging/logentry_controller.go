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
	registry.RegisterModel(krm.LoggingLogEntryGVK, NewLogEntryModel)
}

func NewLogEntryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLogEntry{config: *config}, nil
}

var _ directbase.Model = &modelLogEntry{}

type modelLogEntry struct {
	config config.ControllerConfig
}

func (m *modelLogEntry) client(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (m *modelLogEntry) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.LoggingLogEntry{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewLogEntryIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get logging GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &LogEntryAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelLogEntry) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LogEntryAdapter struct {
	id        *krm.LogEntryIdentity
	gcpClient interface{}
	desired   *krm.LoggingLogEntry
	actual    *unstructured.Unstructured
}

var _ directbase.Adapter = &LogEntryAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *LogEntryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LogEntry", "name", a.id)

	//req := &loggingpb.GetLogEntryRequest{Name: a.id.String()}
	_, err := fmt.Println("Not implemented")
	if err != nil {
		return false, fmt.Errorf("Not implemented: %w", err)
	}
	a.actual = &unstructured.Unstructured{}
	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LogEntryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating LogEntry", "name", a.id)
	mapCtx := &direct.MapContext{}

	//desired := a.desired.DeepCopy()
	//resource := LoggingLogEntrySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	//req := &loggingpb.CreateLogEntryRequest{
		//Parent:   a.id.Parent().String(),
		//LogEntry: resource,
	//}
	_, err := fmt.Println("Not implemented")
	if err != nil {
		return fmt.Errorf("creating LogEntry %s: %w", a.id, err)
	}
	return createOp.UpdateStatus(ctx, nil, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *LogEntryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LogEntry", "name", a.id)
	mapCtx := &direct.MapContext{}

	//desiredPb := LoggingLogEntrySpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	var err error
	//paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	//if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		//status := &krm.LoggingLogEntryStatus{}
		//status.ObservedState = LoggingLogEntryObservedState_FromProto(mapCtx, a.actual)
		//if mapCtx.Err() != nil {
		//	return mapCtx.Err()
		//}
		return updateOp.UpdateStatus(ctx, nil, nil)
	//}
	//updateMask := &fieldmaskpb.FieldMask{
	//	Paths: sets.List(paths)}

	// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
	req := &loggingpb.UpdateLogEntryRequest{
		//Name:       "Not implemented",
		//UpdateMask: updateMask,
		//LogEntry:   desiredPb,
	}
	_, err := fmt.Println("Not implemented")
	if err != nil {
		return fmt.Errorf("updating LogEntry %s: %w", a.id, err)
	}
	return updateOp.UpdateStatus(ctx, nil, nil)

}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *LogEntryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *LogEntryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LogEntry", "name", a.id)

	//req := &loggingpb.DeleteLogEntryRequest{Name: a.id.String()}
	_, err := fmt.Println("Not implemented")
	if err != nil {
			log.V(2).Info("skipping delete for non-existent LogEntry, assuming it was already deleted", "name", a.id.String())
			return true, nil
	}
	return true, nil
}
