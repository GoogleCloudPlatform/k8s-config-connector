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

// +tool:controller
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.Enrollment
// crd.type: EventarcEnrollment
// crd.version: v1alpha1

package eventarc

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/eventarc/apiv1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.EventarcEnrollmentGVK, NewEnrollmentModel)
}

func NewEnrollmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &enrollmentModel{config: *config}, nil
}

var _ directbase.Model = &enrollmentModel{}

type enrollmentModel struct {
	config config.ControllerConfig
}

func (m *enrollmentModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.EventarcEnrollment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEventarcEnrollmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get eventarc GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}
	return &enrollmentAdapter{
		gcpClient: eventarcClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *enrollmentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type enrollmentAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EventarcEnrollmentIdentity
	desired   *krm.EventarcEnrollment
	actual    *pb.Enrollment
	reader    client.Reader
}

var _ directbase.Adapter = &enrollmentAdapter{}

func (a *enrollmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc enrollment", "name", a.id)

	req := &pb.GetEnrollmentRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetEnrollment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc enrollment %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *enrollmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc enrollment", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	desired := a.desired.DeepCopy()
	resource := EventarcEnrollmentSpec_ToProto(mapCtx, &desired.Spec)
	resource.Name = a.id.String()
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateEnrollmentRequest{
		Parent:       fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		Enrollment:   resource,
		EnrollmentId: a.id.Enrollment,
	}

	op, err := a.gcpClient.CreateEnrollment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating eventarc enrollment %q: %w", a.id.String(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc enrollment creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created eventarc enrollment", "name", a.id)

	status := &krm.EventarcEnrollmentStatus{}
	status.ObservedState = EventarcEnrollmentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *enrollmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc enrollment", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	desired := a.desired.DeepCopy()
	resource := EventarcEnrollmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	allowedPaths := make(sets.Set[string])
	allowedPaths.Insert("labels")
	allowedPaths.Insert("annotations")
	allowedPaths.Insert("display_name")
	allowedPaths.Insert("cel_match")
	allowedPaths.Insert("message_bus")
	allowedPaths.Insert("destination")
	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	// Retain updateable fields only
	paths = paths.Intersection(allowedPaths)
	var updated *pb.Enrollment
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		return nil
	} else {
		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		for path := range paths {
			report.AddField(path, nil, nil)
		}
		structuredreporting.ReportDiff(ctx, report)

		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateEnrollmentRequest{
			Enrollment: resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		}

		op, err := a.gcpClient.UpdateEnrollment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating eventarc enrollment %s: %w", a.id.String(), err)
		}

		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for eventarc enrollment update %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated eventarc enrollment", "name", a.id)
	}

	status := &krm.EventarcEnrollmentStatus{}
	status.ObservedState = EventarcEnrollmentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *enrollmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcEnrollment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *EventarcEnrollmentSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Enrollment)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Enrollment)
	u.SetNamespace(obj.Namespace) // This is required KCC controller code convention
	u.SetGroupVersionKind(krm.EventarcEnrollmentGVK)
	u.Object = uObj

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *enrollmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc enrollment", "name", a.id)

	req := &pb.DeleteEnrollmentRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEnrollment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("eventarc enrollment not found", "name", a.id)
			return false, nil // Resource is gone, consider the delete successful.
		}
		return false, fmt.Errorf("deleting eventarc enrollment %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for eventarc enrollment deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted eventarc enrollment", "name", a.id)
	return true, nil
}

func (a *enrollmentAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.MessageBusRef != nil {
		messageBusRef, err := obj.Spec.MessageBusRef.NormalizedExternal(ctx, a.reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.MessageBusRef.External = messageBusRef
	}
	if obj.Spec.DestinationRef != nil {
		destinationRef, err := obj.Spec.DestinationRef.NormalizedExternal(ctx, a.reader, obj.Namespace)
		if err != nil {
			return err
		}
		obj.Spec.DestinationRef.External = destinationRef
	}
	return nil
}
