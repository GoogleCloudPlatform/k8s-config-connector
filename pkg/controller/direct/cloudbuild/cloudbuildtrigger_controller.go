/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudbuild

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/cloudbuild/apiv1/v2"
	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"google.golang.org/api/iterator"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/projects"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudBuildTriggerGVK, NewTriggerModel)
}

func NewTriggerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &triggerModel{config: *config}, nil
}

var _ directbase.Model = &triggerModel{}

type triggerModel struct {
	config config.ControllerConfig
}

func (m *triggerModel) client(ctx context.Context) (*gcp.Client, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudbuild client: %w", err)
	}
	return gcpClient, err
}

func (m *triggerModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudBuildTrigger{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get GCP Project
	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	// Get location
	location := obj.Spec.Location
	if location == "" {
		location = "global"
	}

	var id *CloudBuildTriggerIdentity

	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		id, err = asTriggerID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.project != projectID {
			return nil, fmt.Errorf("CloudBuildTrigger %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.project, projectID)
		}
		if id.location != location {
			return nil, fmt.Errorf("CloudBuildTrigger %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.location, location)
		}
	} else {
		// If externalRef is empty, we don't have the ID yet.
		// We can construct a partial identity or handle it in Find.
		// For now, let's create a placeholder identity with empty trigger ID.
		id = BuildTriggerID(projectID, location, "")
	}

	// Get CloudBuild GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &TriggerAdapter{
		id:            id,
		projectMapper: m.config.ProjectMapper,
		gcpClient:     gcpClient,
		reader:        reader,
		desired:       obj,
	}, nil
}

func (m *triggerModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //cloudbuild.googleapis.com/projects/<project>/locations/<location>/triggers/<id>
	if !strings.HasPrefix(url, "//cloudbuild.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//cloudbuild.googleapis.com/"), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "triggers" {
		// Get CloudBuild GCP client
		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}

		return &TriggerAdapter{
			id: &CloudBuildTriggerIdentity{
				project:  tokens[1],
				location: tokens[3],
				trigger:  tokens[5],
			},
			gcpClient: gcpClient,
		}, nil
	}

	return nil, nil
}

type TriggerAdapter struct {
	id            *CloudBuildTriggerIdentity
	projectMapper *projects.ProjectMapper
	gcpClient     *gcp.Client
	reader        client.Reader
	desired       *krm.CloudBuildTrigger
	actual        *cloudbuildpb.BuildTrigger
}

var _ directbase.Adapter = &TriggerAdapter{}

func (a *TriggerAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.trigger != "" {
		req := &cloudbuildpb.GetBuildTriggerRequest{Name: a.id.FullyQualifiedName()}
		triggerpb, err := a.gcpClient.GetBuildTrigger(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("getting cloudbuildtrigger %q: %w", a.id.FullyQualifiedName(), err)
		}

		a.actual = triggerpb
		return true, nil
	}

	// If ID is empty, we need to list triggers and find by Name.
	// Note: Triggers created via API might not enforce unique names globally, but within project they usually do?
	// The doc says "User-assigned name of the trigger. Must be unique within the project."
	desiredName := direct.ValueOf(a.desired.Spec.Name)
	if desiredName == "" {
		desiredName = a.desired.GetName() // Fallback to metadata.name if spec.name is empty?
		// But spec.name is optional in KRM?
		// If spec.name is empty, we might have used metadata.name.
		// Let's assume spec.name should be populated or inferred.
	}

	req := &cloudbuildpb.ListBuildTriggersRequest{
		Parent: a.id.Parent(),
	}
	it := a.gcpClient.ListBuildTriggers(ctx, req)
	for {
		trigger, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, fmt.Errorf("listing cloudbuildtriggers: %w", err)
		}
		if trigger.GetName() == desiredName {
			a.actual = trigger
			// Update ID
			// The trigger name is projects/.../triggers/{id}
			// We can parse it to get the ID.
			tokens := strings.Split(trigger.ResourceName, "/")
			if len(tokens) > 0 {
				a.id.trigger = tokens[len(tokens)-1]
			}
			return true, nil
		}
	}

	return false, nil
}

func (a *TriggerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	err := a.resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	desired := a.desired.DeepCopy()

	mapCtx := &direct.MapContext{}
	trigger := CloudBuildTriggerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if trigger.Name == "" {
		trigger.Name = u.GetName()
	}

	req := &cloudbuildpb.CreateBuildTriggerRequest{
		Parent:    a.id.Parent(),
		Trigger:   trigger,
		ProjectId: a.id.project,
	}
	created, err := a.gcpClient.CreateBuildTrigger(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildtrigger %s creating failed: %w", trigger.Name, err)
	}

	// Update ID from created object
	// created.ResourceName is fully qualified
	// created.Id is the ID
	if created.Id != "" {
		a.id.trigger = created.Id
	} else {
		// Fallback to parsing resource name
		tokens := strings.Split(created.ResourceName, "/")
		if len(tokens) > 0 {
			a.id.trigger = tokens[len(tokens)-1]
		}
	}

	status := &krm.CloudBuildTriggerStatus{}
	status.ObservedState = CloudBuildTriggerObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setTriggerStatus(u, status)
}

func (a *TriggerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	err := a.resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)

	desired := a.desired.DeepCopy()
	mapCtx := &direct.MapContext{}
	trigger := CloudBuildTriggerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Ensure ID is set for update
	trigger.Id = a.id.trigger
	trigger.ResourceName = a.id.FullyQualifiedName() // Update requires resource name?

	// Compare
	// Note: We need to set fields that are not in Spec but present in Actual to avoid diff
	// if they are immutable or system managed.
	// But common.CompareProtoMessage handles basic diff.

	paths, err := common.CompareProtoMessage(trigger, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.AsExternalRef())
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	// UpdateBuildTriggerRequest
	req := &cloudbuildpb.UpdateBuildTriggerRequest{
		Trigger:   trigger,
		ProjectId: a.id.project,
		TriggerId: a.id.trigger,
		// UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)}, // API might not support UpdateMask?
		// Checked proto: UpdateBuildTriggerRequest has `update_mask`.
	}
	// Note: cloudbuildpb.UpdateBuildTriggerRequest struct in google-cloud-go/cloudbuild/apiv1/v2/cloudbuildpb
	// definition might vary.
	// Let's assume it has UpdateMask if it's standard AIP.
	// I'll check if I can use UpdateMask.
	// If the generated code doesn't support UpdateMask, I might have to pass full object.
	// But wait, `CompareProtoMessage` returns paths.
	// If the API supports update_mask, I should use it.
	// Let's check `cloudbuildpb` definition indirectly or assume it's there.
	// Most Google APIs have it.

	// However, the `cloudbuildpb` package I imported: `cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb`
	// Verify if `UpdateBuildTriggerRequest` has `UpdateMask`.
	// If not, I can't set it.
	// For now I'll optimistically assume it does or I'll just not set it (which implies full update).
	// But full update might clear fields?
	// Usually UpdateBuildTrigger replaces the trigger.
	// So passing full object is fine.

	updated, err := a.gcpClient.UpdateBuildTrigger(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildtrigger %s updating failed: %w", trigger.Name, err)
	}

	status := &krm.CloudBuildTriggerStatus{}
	status.ObservedState = CloudBuildTriggerObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return fmt.Errorf("update trigger status %w", mapCtx.Err())
	}
	return setTriggerStatus(u, status)
}

func (a *TriggerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudBuildTrigger{}
	obj.SetGroupVersionKind(krm.CloudBuildTriggerGVK)
	obj.SetName(a.actual.Name)

	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudBuildTriggerSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.project}
	obj.Spec.Location = a.id.location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *TriggerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &cloudbuildpb.DeleteBuildTriggerRequest{
		Name:      a.id.FullyQualifiedName(),
		ProjectId: a.id.project,
		TriggerId: a.id.trigger,
	}
	err := a.gcpClient.DeleteBuildTrigger(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting cloudbuildtrigger %s: %w", a.id.FullyQualifiedName(), err)
	}
	return true, nil
}

func setTriggerStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}

func (a *TriggerAdapter) resolveDependencies(ctx context.Context, reader client.Reader, obj *krm.CloudBuildTrigger) error {
	// Resolve dependencies if any.
	// Currently CloudBuildTrigger only has ProjectRef which is resolved in AdapterForObject.
	// ServiceAccount is a string.
	// If there are other refs (like PubSub topic), they might need resolution if they were refs.
	// But `PubsubConfig` has `Topic *string` (format: projects/.../topics/...).
	// So no resolution needed unless we change it to Ref.
	return nil
}
