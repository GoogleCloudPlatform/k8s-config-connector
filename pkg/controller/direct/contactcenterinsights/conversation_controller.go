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

package contactcenterinsights

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/contactcenterinsights/apiv1"
	contactcenterinsightspb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CCInsightsConversationGVK, NewConversationModel)
}

func NewConversationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConversation{config: *config}, nil
}

type modelConversation struct {
	config config.ControllerConfig
}

func (m *modelConversation) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	var err error
	opts, err = m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building contactcenterinsights client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelConversation) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CCInsightsConversation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewCCInsightsConversationIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get contactcenterinsights GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConversationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelConversation) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ConversationAdapter struct {
	id        *krm.CCInsightsConversationIdentity
	gcpClient *gcp.Client
	desired   *krm.CCInsightsConversation
	actual    *contactcenterinsightspb.Conversation
}

var _ directbase.Adapter = &ConversationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
func (a *ConversationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding Conversation", "name", a.id)

	req := &contactcenterinsightspb.GetConversationRequest{
		Name: a.id.String(),
	}
	resp, err := a.gcpClient.GetConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Conversation %q: %w", a.id, err)
	}

	a.actual = resp
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConversationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Conversation", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := CCInsightsConversationSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &contactcenterinsightspb.CreateConversationRequest{
		Parent:       fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		Conversation: resource,
	}
	created, err := a.gcpClient.CreateConversation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Conversation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Conversation", "name", a.id)

	status := &krm.CCInsightsConversationStatus{}
	status.ObservedState = CCInsightsConversationObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ConversationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Conversation", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := CCInsightsConversationSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	{
		var err error
		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		desiredPb.Name = a.id.String()
		req := &contactcenterinsightspb.UpdateConversationRequest{
			UpdateMask:   updateMask,
			Conversation: desiredPb,
		}
		var err error
		updated, err = a.gcpClient.UpdateConversation(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Conversation %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Conversation", "name", a.id)
	}

	status := &krm.CCInsightsConversationStatus{}
	status.ObservedState = CCInsightsConversationObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ConversationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CCInsightsConversation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CCInsightsConversationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Conversation)
	u.SetGroupVersionKind(krm.CCInsightsConversationGVK)

	u.Object = uObj
	return u, nil
}

// Delete deletes the resource in GCP.
// Return true means the resource is successfully deleted.
// Return false means the resource is not deleted yet.
func (a *ConversationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Conversation", "name", a.id)

	req := &contactcenterinsightspb.DeleteConversationRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Conversation, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Conversation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Conversation", "name", a.id)
	return true, nil
}
