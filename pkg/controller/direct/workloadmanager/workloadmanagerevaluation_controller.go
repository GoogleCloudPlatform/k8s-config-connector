// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses///LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workloadmanager

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workloadmanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/workloadmanager/apiv1"
	workloadmanagerpb "cloud.google.com/go/workloadmanager/apiv1/workloadmanagerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.WorkloadManagerEvaluationGVK, NewEvaluationModel)
}

func NewEvaluationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEvaluation{config: *config}, nil
}

var _ directbase.Model = &modelEvaluation{}

type modelEvaluation struct {
	config config.ControllerConfig
}

func (m *modelEvaluation) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Evaluation client: %w", err)
	}
	return gcpClient, err
}

func (m *modelEvaluation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.WorkloadManagerEvaluation{}
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

	// Get workloadmanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EvaluationAdapter{
		id:        id.(*krm.WorkloadManagerEvaluationIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEvaluation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EvaluationAdapter struct {
	id        *krm.WorkloadManagerEvaluationIdentity
	gcpClient *gcp.Client
	desired   *krm.WorkloadManagerEvaluation
	actual    *workloadmanagerpb.Evaluation
}

var _ directbase.Adapter = &EvaluationAdapter{}

// Find retrieves the GCP resource.
func (a *EvaluationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Evaluation", "name", a.id)

	req := &workloadmanagerpb.GetEvaluationRequest{Name: a.id.String()}
	evaluationpb, err := a.gcpClient.GetEvaluation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Evaluation %q: %w", a.id, err)
	}

	a.actual = evaluationpb
	return true, nil
}

// Create creates the resource in GCP.
func (a *EvaluationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Evaluation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkloadManagerEvaluationSpec_ToProto(mapCtx, &desired.Spec)
	if desired.Spec.KmsKeyRef != nil {
		resource.KmsKey = desired.Spec.KmsKeyRef.External
	}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &workloadmanagerpb.CreateEvaluationRequest{
		Parent:       a.id.ParentString(),
		EvaluationId: a.id.Evaluation,
		Evaluation:   resource,
	}
	op, err := a.gcpClient.CreateEvaluation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Evaluation %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Evaluation %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Evaluation", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

// Update updates the resource in GCP.
func (a *EvaluationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Evaluation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := WorkloadManagerEvaluationSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if a.desired.Spec.KmsKeyRef != nil {
		desiredPb.KmsKey = a.desired.Spec.KmsKeyRef.External
	}
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	diffs, updateMask, err := compareResource(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	updated := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", updateMask.Paths)
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &workloadmanagerpb.UpdateEvaluationRequest{
			UpdateMask: updateMask,
			Evaluation: desiredPb,
		}
		op, err := a.gcpClient.UpdateEvaluation(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Evaluation %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Evaluation %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Evaluation", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func compareResource(ctx context.Context, actual, desired *workloadmanagerpb.Evaluation) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := WorkloadManagerEvaluationSpec_FromProto(mapCtx, actual)
	if actual.GetKmsKey() != "" {
		spec.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: actual.GetKmsKey()}
	}
	maskedActual := WorkloadManagerEvaluationSpec_ToProto(mapCtx, spec)
	if spec.KmsKeyRef != nil {
		maskedActual.KmsKey = spec.KmsKeyRef.External
	}
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}

	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*workloadmanagerpb.Evaluation)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *EvaluationAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *workloadmanagerpb.Evaluation) error {
	mapCtx := &direct.MapContext{}
	status := &krm.WorkloadManagerEvaluationStatus{}
	status.ObservedState = WorkloadManagerEvaluationObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EvaluationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.WorkloadManagerEvaluation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(WorkloadManagerEvaluationSpec_FromProto(mapCtx, a.actual))
	if a.actual.GetKmsKey() != "" {
		obj.Spec.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: a.actual.GetKmsKey()}
	}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Evaluation)
	u.SetGroupVersionKind(krm.WorkloadManagerEvaluationGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service.
func (a *EvaluationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Evaluation", "name", a.id)

	req := &workloadmanagerpb.DeleteEvaluationRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEvaluation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Evaluation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Evaluation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Evaluation", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Evaluation %s: %w", a.id, err)
	}
	return true, nil
}
