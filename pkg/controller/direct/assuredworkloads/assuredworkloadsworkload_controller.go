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

package assuredworkloads

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/assuredworkloads/apiv1"
	pb "cloud.google.com/go/assuredworkloads/apiv1/assuredworkloadspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/assuredworkloads/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.AssuredWorkloadsWorkloadGVK, NewAssuredWorkloadsWorkloadModel)
}

func NewAssuredWorkloadsWorkloadModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAssuredWorkloadsWorkload{config: *config}, nil
}

var _ directbase.Model = &modelAssuredWorkloadsWorkload{}

type modelAssuredWorkloadsWorkload struct {
	config config.ControllerConfig
}

func (m *modelAssuredWorkloadsWorkload) client(ctx context.Context, location string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" {
		endpoint := fmt.Sprintf("https://%s-assuredworkloads.googleapis.com", location)
		opts = append(opts, option.WithEndpoint(endpoint))
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AssuredWorkloads client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAssuredWorkloadsWorkload) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AssuredWorkloadsWorkload{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identityVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identityVal.(*krm.AssuredWorkloadsWorkloadIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", identityVal)
	}

	// Always call common.NormalizeReferences to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx, obj.Spec.Location)
	if err != nil {
		return nil, err
	}
	return &AssuredWorkloadsWorkloadAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelAssuredWorkloadsWorkload) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AssuredWorkloadsWorkloadAdapter struct {
	id        *krm.AssuredWorkloadsWorkloadIdentity
	gcpClient *gcp.Client
	desired   *krm.AssuredWorkloadsWorkload
	actual    *pb.Workload
	reader    client.Reader
}

var _ directbase.Adapter = &AssuredWorkloadsWorkloadAdapter{}

func (a *AssuredWorkloadsWorkloadAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AssuredWorkloadsWorkload", "name", a.id)

	req := &pb.GetWorkloadRequest{Name: a.id.String()}
	workload, err := a.gcpClient.GetWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AssuredWorkloadsWorkload %q: %w", a.id, err)
	}

	a.actual = workload
	return true, nil
}

func (a *AssuredWorkloadsWorkloadAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AssuredWorkloadsWorkload", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()

	resource := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateWorkloadRequest{
		Parent:   a.id.ParentString(),
		Workload: resource,
	}
	op, err := a.gcpClient.CreateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AssuredWorkloadsWorkload %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AssuredWorkloadsWorkload %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created AssuredWorkloadsWorkload", "name", a.id)

	// Fetch the fully-populated resource after LRO finishes
	getReq := &pb.GetWorkloadRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetWorkload(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting fully-populated AssuredWorkloadsWorkload after creation: %w", err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *AssuredWorkloadsWorkloadAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AssuredWorkloadsWorkload", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredPb := AssuredWorkloadsWorkloadSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	diffs, updateMask, err := compareResource(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateWorkloadRequest{
		UpdateMask: updateMask,
		Workload:   desiredPb,
	}
	_, err = a.gcpClient.UpdateWorkload(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AssuredWorkloadsWorkload %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated AssuredWorkloadsWorkload", "name", a.id)

	// Fetch the fully-populated resource after update
	getReq := &pb.GetWorkloadRequest{Name: a.id.String()}
	latest, err := a.gcpClient.GetWorkload(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting fully-populated AssuredWorkloadsWorkload after update: %w", err)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *AssuredWorkloadsWorkloadAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AssuredWorkloadsWorkload{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AssuredWorkloadsWorkloadSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Workload)
	u.SetGroupVersionKind(krm.AssuredWorkloadsWorkloadGVK)
	u.SetAnnotations(map[string]string{
		"cnrm.cloud.google.com/organization-id": a.id.Organization,
	})

	u.Object = uObj
	return u, nil
}

func (a *AssuredWorkloadsWorkloadAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AssuredWorkloadsWorkload", "name", a.id)

	req := &pb.DeleteWorkloadRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteWorkload(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent AssuredWorkloadsWorkload, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AssuredWorkloadsWorkload %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AssuredWorkloadsWorkload", "name", a.id)
	return true, nil
}

func (a *AssuredWorkloadsWorkloadAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Workload) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AssuredWorkloadsWorkloadStatus{}
	status.ObservedState = AssuredWorkloadsWorkloadObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func compareResource(ctx context.Context, actual, desired *pb.Workload) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, AssuredWorkloadsWorkloadSpec_FromProto, AssuredWorkloadsWorkloadSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Workload) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
