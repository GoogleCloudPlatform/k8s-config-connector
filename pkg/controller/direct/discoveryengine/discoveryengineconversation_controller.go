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
// proto.service: google.cloud.discoveryengine.v1.ConversationalSearchService
// proto.message: google.cloud.discoveryengine.v1.Conversation
// crd.type: DiscoveryEngineConversation
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
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
	registry.RegisterModel(krm.DiscoveryEngineConversationGVK, NewConversationModel)
}

func NewConversationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &conversationModel{config: *config}, nil
}

var _ directbase.Model = &conversationModel{}

type conversationModel struct {
	config config.ControllerConfig
}

func (m *conversationModel) client(ctx context.Context, projectID string) (*gcp.ConversationalSearchClient, error) {
	var opts []option.ClientOption

	config := m.config

	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewConversationalSearchRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine conversational search client: %w", err)
	}

	return gcpClient, err
}

func (m *conversationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineConversation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DiscoveryEngineConversationIdentity)

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineConversationSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &conversationAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *conversationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
		id := &krm.DiscoveryEngineConversationIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineConversation format", "url", url)
			return nil, nil
		}
		gcpClient, err := m.client(ctx, id.Project)
		if err != nil {
			return nil, err
		}
		return &conversationAdapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type conversationAdapter struct {
	gcpClient *gcp.ConversationalSearchClient
	id        *krm.DiscoveryEngineConversationIdentity
	desired   *pb.Conversation
	actual    *pb.Conversation
}

var _ directbase.Adapter = &conversationAdapter{}

func (a *conversationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting discoveryengine conversation", "name", fqn)

	req := &pb.GetConversationRequest{Name: fqn}
	actual, err := a.gcpClient.GetConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine conversation %q from gcp: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *conversationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine conversation", "name", a.id)

	desired := proto.Clone(a.desired).(*pb.Conversation)
	desired.Name = a.id.String()

	req := &pb.CreateConversationRequest{
		Parent:       a.id.ParentString(),
		Conversation: desired,
	}
	created, err := a.gcpClient.CreateConversation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine conversation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created discoveryengine conversation in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *conversationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine conversation", "name", a.id)

	diffs, updateMask, err := compareConversation(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.Conversation)
		desired.Name = a.id.String()

		req := &pb.UpdateConversationRequest{
			Conversation: desired,
			UpdateMask:   updateMask,
		}

		updated, err := a.gcpClient.UpdateConversation(ctx, req)
		if err != nil {
			return fmt.Errorf("updating discoveryengine conversation %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareConversation(ctx context.Context, actual, desired *pb.Conversation) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineConversationSpec_v1alpha1_FromProto, DiscoveryEngineConversationSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore name if needed

	clonedDesired := proto.Clone(desired).(*pb.Conversation)

	populateDefaults := func(obj *pb.Conversation) {
		if obj.State == pb.Conversation_STATE_UNSPECIFIED {
			obj.State = pb.Conversation_IN_PROGRESS
		}
		if obj.UserPseudoId == "" && actual != nil && actual.UserPseudoId != "" {
			// If not specified in desired, and GCP assigned one, accept it
			obj.UserPseudoId = actual.UserPseudoId
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

func (a *conversationAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Conversation) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DiscoveryEngineConversationStatus{}
	status.ObservedState = DiscoveryEngineConversationObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(latest.GetName())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *conversationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineConversation{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineConversationSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{
		External: fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", a.id.Project, a.id.Location, a.id.Collection, a.id.DataStore),
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Conversation)
	u.SetGroupVersionKind(krm.DiscoveryEngineConversationGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *conversationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting discoveryengine conversation", "name", fqn)

	req := &pb.DeleteConversationRequest{Name: fqn}
	err := a.gcpClient.DeleteConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting discoveryengine conversation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine conversation", "name", a.id)

	return true, nil
}
