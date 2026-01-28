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

package run

import (
	"context"
	"fmt"
	"strings"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"google.golang.org/protobuf/proto"

	gcp "cloud.google.com/go/run/apiv2"

	pb "cloud.google.com/go/run/apiv2/runpb"
	runpb "cloud.google.com/go/run/apiv2/runpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.RunJobGVK, NewJobModel)
}

func NewJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelJob{config: *config}, nil
}

var _ directbase.Model = &modelJob{}

type modelJob struct {
	config config.ControllerConfig
}

func (m *modelJob) client(ctx context.Context) (*gcp.JobsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewJobsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Job client: %w", err)
	}
	return gcpClient, err
}

func (m *modelJob) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.RunJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewJobIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	// TODO: this should not block DELETION.
	if err := ResolveRunJobRefs(ctx, reader, obj); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()
	desired := RunJobSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	// Get run GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &JobAdapter{
		id:                 id,
		gcpClient:          gcpClient,
		desired:            desired,
		lastModifiedCookie: obj.Status.LastModifiedCookie,
	}, nil
}

func (m *modelJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//run.googleapis.com/"); ok {
		// Direct controller for RunJob only handles v2.
		s = strings.TrimPrefix(s, "v2/")

		var id krm.JobIdentity
		if err := id.FromExternal(s); err != nil {
			log.V(2).Error(err, "url did not match RunJob format", "url", url)
			return nil, nil
		}

		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &JobAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type JobAdapter struct {
	id                 *krm.JobIdentity
	gcpClient          *gcp.JobsClient
	desired            *runpb.Job
	actual             *runpb.Job
	lastModifiedCookie *string
}

var _ directbase.Adapter = &JobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *JobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Job", "name", a.id)

	req := &runpb.GetJobRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Job %q: %w", a.id, err)
	}

	a.actual = found
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *JobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Job", "name", a.id)
	req := &runpb.CreateJobRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", a.id.ProjectID, a.id.Location),
		Job:    a.desired,
		JobId:  a.id.ID(),
	}
	op, err := a.gcpClient.CreateJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Job %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of job %q: %w", a.id, err)
	}
	log.V(2).Info("successfully created Job", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.RunJobStatus{}
	status.ObservedState = RunJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	newCookie, err := common.NewCookie(a.desired, created)
	if err != nil {
		return fmt.Errorf("composing cookie: %w", err)
	}
	log.V(2).Info("Job cookie added", "name", a.id, "new-cookie", newCookie.String())
	status.LastModifiedCookie = direct.LazyPtr(newCookie.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *JobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Job", "name", a.id)

	currentCookie, err := common.NewCookie(a.desired, a.actual)
	if err != nil {
		return err
	}
	if currentCookie.Equal(a.lastModifiedCookie) {
		log.V(2).Info("resource is up to date", "name", a.id)
		return a.updateStatus(ctx, a.actual, updateOp)
	}

	// We need to set the name for the update request, but we don't want to modify a.desired
	// as that would change the computed hash (specHash).
	updateJob := proto.Clone(a.desired).(*runpb.Job)
	updateJob.Name = a.actual.Name
	req := &runpb.UpdateJobRequest{
		Job: updateJob,
	}
	op, err := a.gcpClient.UpdateJob(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Job %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Job %s waiting update: %w", a.id, err)
	}
	log.Info("successfully updated Job", "name", a.id)
	return a.updateStatus(ctx, updated, updateOp)
}

func (a *JobAdapter) updateStatus(ctx context.Context, updated *pb.Job, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	status := &krm.RunJobStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = RunJobObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	updatedCookie, err := common.NewCookie(a.desired, updated)
	if err != nil {
		return err
	}
	status.LastModifiedCookie = direct.LazyPtr(updatedCookie.String())
	status.ExternalRef = direct.LazyPtr(a.id.String())

	if !updatedCookie.Equal(a.lastModifiedCookie) {
		log.Info("Job cookie updated", "name", a.id, "old-cookie", direct.ValueOf(a.lastModifiedCookie),
			"new-cookie", updatedCookie.String())
	}

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *JobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.RunJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RunJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ProjectID}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.RunJobGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *JobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Job", "name", a.id)

	name := a.id.String()
	req := &runpb.DeleteJobRequest{Name: name}
	op, err := a.gcpClient.DeleteJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Job, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Job %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Job", "name", a.id)

	if _, err = op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting delete Job %s: %w", a.id, err)
	}
	return true, nil
}
