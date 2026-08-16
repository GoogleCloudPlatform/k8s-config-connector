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

package transcoder

import (
	"context"
	"errors"
	"fmt"

	gcp "cloud.google.com/go/video/transcoder/apiv1"
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/transcoder/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.TranscoderJobGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Transcoder client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.TranscoderJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.TranscoderJobIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := TranscoderJobSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if desired.Labels == nil {
		desired.Labels = make(map[string]string)
	}
	for k, v := range label.NewGCPLabelsFromK8sLabels(u.GetLabels()) {
		desired.Labels[k] = v
	}
	desired.Labels["cnrm-resource-id"] = id.Job

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		k8sName:   id.Job,
		model:     m,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.TranscoderJobIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		k8sName:   id.Job,
		model:     m,
	}, nil
}

type Adapter struct {
	id        *krm.TranscoderJobIdentity
	gcpClient *gcp.Client
	desired   *pb.Job
	actual    *pb.Job
	k8sName   string
	model     *model
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding TranscoderJob", "id", a.id)

	if a.id.Job == a.k8sName {
		// Status.ExternalRef was empty, list jobs with label filter.
		it := a.gcpClient.ListJobs(ctx, &pb.ListJobsRequest{
			Parent: a.id.ParentString(),
			Filter: fmt.Sprintf("labels.cnrm-resource-id:%s", a.k8sName),
		})
		for {
			resp, err := it.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				return false, fmt.Errorf("listing TranscoderJobs: %w", err)
			}
			parsed, match, err := krm.TranscoderJobIdentityFormat.Parse(resp.GetName())
			if err == nil && match {
				a.actual = resp
				a.id.Job = parsed.Job // Use the UUID
				return true, nil
			}
		}
		return false, nil
	}

	req := &pb.GetJobRequest{
		Name: a.id.String(),
	}
	job, err := a.gcpClient.GetJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TranscoderJob %s: %w", a.id.String(), err)
	}

	a.actual = job
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TranscoderJob", "id", a.id)

	req := &pb.CreateJobRequest{
		Parent: a.id.ParentString(),
		Job:    a.desired,
	}
	created, err := a.gcpClient.CreateJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TranscoderJob under %s: %w", a.id.ParentString(), err)
	}

	parsed, match, err := krm.TranscoderJobIdentityFormat.Parse(created.GetName())
	if err == nil && match {
		a.id.Job = parsed.Job
	}
	a.actual = created

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TranscoderJob", "id", a.id)

	diffs, err := a.compareJob(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("TranscoderJob is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Job) error {
	mapCtx := &direct.MapContext{}
	status := krm.TranscoderJobStatus{}
	status.ObservedState = TranscoderJobObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, &status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TranscoderJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TranscoderJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.ResourceID = &a.k8sName

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.k8sName)
	u.SetGroupVersionKind(krm.TranscoderJobGVK)

	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TranscoderJob", "id", a.id)

	req := &pb.DeleteJobRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent TranscoderJob, assuming it was already deleted", "id", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting TranscoderJob %s: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *Adapter) compareJob(ctx context.Context, actual, desired *pb.Job) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, TranscoderJobSpec_FromProto, TranscoderJobSpec_ToProto)
	if err != nil {
		return nil, err
	}
	maskedActual.Name = desired.Name
	maskedActual.Labels = desired.Labels

	diffs, _, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}
