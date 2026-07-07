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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contactcenterinsights/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/contactcenterinsights/apiv1"
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CCInsightsConversationGVK, NewCCInsightsConversationModel)
}

func NewCCInsightsConversationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &ccInsightsConversationModel{config: *config}, nil
}

var _ directbase.Model = &ccInsightsConversationModel{}

type ccInsightsConversationModel struct {
	config config.ControllerConfig
}

func (m *ccInsightsConversationModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building contactcenterinsights REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *ccInsightsConversationModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CCInsightsConversation{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.CCInsightsConversationIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// Normalize references in the spec
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Map to Proto once for the desired state
	mapCtx := &direct.MapContext{}
	desired := CCInsightsConversationSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ccInsightsConversationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		rawObj:    obj,
		reader:    reader,
	}, nil
}

func (m *ccInsightsConversationModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ccInsightsConversationAdapter struct {
	id        *krm.CCInsightsConversationIdentity
	gcpClient *gcp.Client
	desired   *pb.Conversation
	actual    *pb.Conversation
	rawObj    *krm.CCInsightsConversation
	reader    client.Reader
}

var _ directbase.Adapter = &ccInsightsConversationAdapter{}

func (a *ccInsightsConversationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.CCInsightsConversationGVK)
	u.Object = uObj
	return u, nil
}

func (a *ccInsightsConversationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CCInsightsConversation", "name", a.id)

	req := &pb.GetConversationRequest{Name: a.id.String()}
	conversation, err := a.gcpClient.GetConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CCInsightsConversation %q: %w", a.id, err)
	}

	a.actual = conversation
	return true, nil
}

func (a *ccInsightsConversationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CCInsightsConversation", "name", a.id)

	req := &pb.CreateConversationRequest{
		Parent:         a.id.ParentString(),
		ConversationId: a.id.ID(),
		Conversation:   a.desired,
	}
	req.Conversation.Name = a.id.String()

	created, err := a.gcpClient.CreateConversation(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CCInsightsConversation %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CCInsightsConversation", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ccInsightsConversationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CCInsightsConversation", "name", a.id)

	diffs, updateMask, err := compareCCInsightsConversation(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateConversationRequest{
		Conversation: a.desired,
		UpdateMask:   updateMask,
	}
	req.Conversation.Name = a.id.String()

	updated, err := a.gcpClient.UpdateConversation(ctx, req)
	if err != nil {
		return fmt.Errorf("updating CCInsightsConversation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated CCInsightsConversation", "name", a.id)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *ccInsightsConversationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CCInsightsConversation", "name", a.id)

	req := &pb.DeleteConversationRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteConversation(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting CCInsightsConversation %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted CCInsightsConversation", "name", a.id)
	return true, nil
}

func (a *ccInsightsConversationAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Conversation) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CCInsightsConversationStatus{}
	status.ObservedState = CCInsightsConversationObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}

func compareCCInsightsConversation(ctx context.Context, actual, desired *pb.Conversation) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CCInsightsConversationSpec_FromProto, CCInsightsConversationSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Conversation) {
		// Define and populate GCP/server defaults here if any exist
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
