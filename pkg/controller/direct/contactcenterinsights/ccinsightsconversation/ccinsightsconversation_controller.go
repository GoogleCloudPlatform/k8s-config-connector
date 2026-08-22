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
// proto.service: google.cloud.contactcenterinsights.v1.ContactCenterInsights
// proto.message: google.cloud.contactcenterinsights.v1.Conversation
// crd.type: CCInsightsConversation
// crd.version: v1alpha1

package ccinsightsconversation

import (
	"context"
	"fmt"

	contactcenterinsights "cloud.google.com/go/contactcenterinsights/apiv1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CCInsightsConversationGVK, newModel)
}

func newModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type adapter struct {
	id      *krm.CCInsightsConversationIdentity
	desired *pb.Conversation
	actual  *pb.Conversation
	gcp     *contactcenterinsights.Client
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*contactcenterinsights.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := contactcenterinsights.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building contactcenterinsights client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.CCInsightsConversation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := ConversationSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &adapter{
		id:      id.(*krm.CCInsightsConversationIdentity),
		desired: desired,
		gcp:     gcp,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.CCInsightsConversationIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		id:  id,
		gcp: gcp,
	}, nil
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	req := &pb.GetConversationRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcp.GetConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting conversation: %w", err)
	}
	a.actual = actual
	return true, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CCInsightsConversation", "id", a.id)

	conversation := proto.Clone(a.desired).(*pb.Conversation)

	req := &pb.CreateConversationRequest{
		Parent:         a.id.ParentString(),
		ConversationId: a.id.Conversation,
		Conversation:   conversation,
	}

	created, err := a.gcp.CreateConversation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating conversation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("created CCInsightsConversation", "id", a.id)
	return a.updateStatus(ctx, createOp, created)
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CCInsightsConversation", "id", a.id)

	mapCtx := &direct.MapContext{}
	actualSpec := ConversationSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := ConversationSpec_ToProto(mapCtx, actualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	clonedDesired := proto.Clone(a.desired).(*pb.Conversation)

	// For server-defaulted or automatically populated fields that cannot be cleared,
	// if they are not set in the desired state, copy them from actual to avoid false diffs
	// and invalid update mask entries.
	if clonedDesired.StartTime == nil && maskedActual.StartTime != nil {
		clonedDesired.StartTime = maskedActual.StartTime
	}
	if clonedDesired.LanguageCode == "" && maskedActual.LanguageCode != "" {
		clonedDesired.LanguageCode = maskedActual.LanguageCode
	}
	if clonedDesired.AgentId == "" && maskedActual.AgentId != "" {
		clonedDesired.AgentId = maskedActual.AgentId
	}
	if clonedDesired.Expiration == nil && maskedActual.Expiration != nil {
		clonedDesired.Expiration = maskedActual.Expiration
	}

	// The following fields are immutable on GCP, so we must copy them from actual to desired
	// to ensure they are never included in the update mask.
	clonedDesired.ObfuscatedUserId = maskedActual.ObfuscatedUserId
	clonedDesired.DataSource = maskedActual.DataSource
	clonedDesired.Medium = maskedActual.Medium

	if maskedActual.Metadata != nil {
		switch m := maskedActual.Metadata.(type) {
		case *pb.Conversation_CallMetadata_:
			clonedDesired.Metadata = &pb.Conversation_CallMetadata_{
				CallMetadata: proto.Clone(m.CallMetadata).(*pb.Conversation_CallMetadata),
			}
		default:
			// Fallback/safety in case there are other oneof fields in Metadata in the future.
			clonedDesired.Metadata = maskedActual.Metadata
		}
	}

	if clonedDesired.QualityMetadata == nil && maskedActual.QualityMetadata != nil {
		clonedDesired.QualityMetadata = proto.Clone(maskedActual.QualityMetadata).(*pb.Conversation_QualityMetadata)
	}

	// Since the GCP server returns a 'teams' array instead of the singular 'team' string field,
	// and our compiled Go client library version doesn't yet model 'teams', the server-returned 'team'
	// field is empty. We copy the desired 'team' to maskedActual to avoid false diffs.
	if clonedDesired.QualityMetadata != nil && maskedActual.QualityMetadata != nil {
		for i, desiredAgent := range clonedDesired.QualityMetadata.AgentInfo {
			if i < len(maskedActual.QualityMetadata.AgentInfo) {
				actualAgent := maskedActual.QualityMetadata.AgentInfo[i]
				if desiredAgent.Team != "" && actualAgent.Team == "" {
					actualAgent.Team = desiredAgent.Team
				}
			}
		}
	}

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "id", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	conversation := proto.Clone(clonedDesired).(*pb.Conversation)
	conversation.Name = a.id.String()

	req := &pb.UpdateConversationRequest{
		Conversation: conversation,
		UpdateMask:   updateMask,
	}

	updated, err := a.gcp.UpdateConversation(ctx, req)
	if err != nil {
		return fmt.Errorf("updating conversation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("updated CCInsightsConversation", "id", a.id)
	return a.updateStatus(ctx, updateOp, updated)
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CCInsightsConversation", "id", a.id)

	req := &pb.DeleteConversationRequest{
		Name: a.id.String(),
	}

	err := a.gcp.DeleteConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting conversation %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted CCInsightsConversation", "id", a.id)
	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	mapCtx := &direct.MapContext{}
	spec := ConversationSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.CCInsightsConversation{}
	obj.Spec = *spec
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.Conversation

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u := &unstructured.Unstructured{Object: specObj}
	u.SetGroupVersionKind(krm.CCInsightsConversationGVK)
	return u, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Conversation) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CCInsightsConversationStatus{}
	status.ObservedState = ConversationObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
