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

// +tool:controller
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.Pipeline
// crd.type: EventarcPipeline
// crd.version: v1alpha1

package eventarc

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/eventarc/apiv1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.EventarcPipelineGVK, NewPipelineModel)
}

func NewPipelineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &pipelineModel{config: *config}, nil
}

var _ directbase.Model = &pipelineModel{}

type pipelineModel struct {
	config config.ControllerConfig
}

func (m *pipelineModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.EventarcPipeline{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewEventarcPipelineIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := EventarcPipelineSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredPb.Name = id.String()

	// Get eventarc GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}

	return &pipelineAdapter{
		gcpClient: eventarcClient,
		id:        id,
		desired:   desiredPb,
		reader:    reader,
		obj:       obj,
	}, nil
}

func (m *pipelineModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.EventarcPipelineIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	eventarcClient, err := gcpClient.newEventarcClient(ctx)
	if err != nil {
		return nil, err
	}

	return &pipelineAdapter{
		gcpClient: eventarcClient,
		id:        id,
	}, nil
}

type pipelineAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EventarcPipelineIdentity
	desired   *pb.Pipeline
	actual    *pb.Pipeline
	reader    client.Reader
	obj       *krm.EventarcPipeline
}

var _ directbase.Adapter = &pipelineAdapter{}

func (a *pipelineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting eventarc pipeline", "name", a.id.String())

	req := &pb.GetPipelineRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetPipeline(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting eventarc pipeline %q: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *pipelineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating eventarc pipeline", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.Pipeline)
	desired.Name = a.id.String()

	req := &pb.CreatePipelineRequest{
		Parent:     a.id.ParentString(),
		Pipeline:   desired,
		PipelineId: a.id.Pipeline,
	}

	op, err := a.gcpClient.CreatePipeline(ctx, req)
	if err != nil {
		return fmt.Errorf("creating eventarc pipeline %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for eventarc pipeline creation %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created eventarc pipeline", "name", a.id.String())

	// Fetch fully-populated resource after creation
	latest, err := a.gcpClient.GetPipeline(ctx, &pb.GetPipelineRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting eventarc pipeline after creation %q: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *pipelineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating eventarc pipeline", "name", a.id.String())

	diffs, updateMask, err := compareEventarcPipeline(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desiredCopy := proto.Clone(a.desired).(*pb.Pipeline)
		desiredCopy.Name = a.id.String()

		req := &pb.UpdatePipelineRequest{
			Pipeline:   desiredCopy,
			UpdateMask: updateMask,
		}

		op, err := a.gcpClient.UpdatePipeline(ctx, req)
		if err != nil {
			return fmt.Errorf("updating eventarc pipeline %s: %w", a.id.String(), err)
		}

		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for eventarc pipeline update %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated eventarc pipeline", "name", a.id.String())

		latest, err = a.gcpClient.GetPipeline(ctx, &pb.GetPipelineRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting eventarc pipeline after update %s: %w", a.id.String(), err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *pipelineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EventarcPipeline{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EventarcPipelineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Pipeline)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Pipeline)
	u.SetNamespace(obj.Namespace)
	u.SetGroupVersionKind(krm.EventarcPipelineGVK)
	u.Object = uObj

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *pipelineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting eventarc pipeline", "name", a.id.String())

	req := &pb.DeletePipelineRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePipeline(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("eventarc pipeline not found", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting eventarc pipeline %q: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for eventarc pipeline deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted eventarc pipeline", "name", a.id.String())
	return true, nil
}

func (a *pipelineAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Pipeline) error {
	mapCtx := &direct.MapContext{}
	status := &krm.EventarcPipelineStatus{}
	status.ObservedState = EventarcPipelineObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}

func compareEventarcPipeline(ctx context.Context, actual, desired *pb.Pipeline) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, EventarcPipelineSpec_FromProto, EventarcPipelineSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Pipeline)

	populateDefaults := func(obj *pb.Pipeline) {
		if obj.RetryPolicy == nil {
			obj.RetryPolicy = &pb.Pipeline_RetryPolicy{
				MaxAttempts:   5,
				MinRetryDelay: &durationpb.Duration{Seconds: 1},
				MaxRetryDelay: &durationpb.Duration{Seconds: 60},
			}
		} else {
			if obj.RetryPolicy.MaxAttempts == 0 {
				obj.RetryPolicy.MaxAttempts = 5
			}
			if obj.RetryPolicy.MinRetryDelay == nil {
				obj.RetryPolicy.MinRetryDelay = &durationpb.Duration{Seconds: 1}
			}
			if obj.RetryPolicy.MaxRetryDelay == nil {
				obj.RetryPolicy.MaxRetryDelay = &durationpb.Duration{Seconds: 60}
			}
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
