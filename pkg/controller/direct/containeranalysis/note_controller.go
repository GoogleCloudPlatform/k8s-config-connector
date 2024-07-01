// Copyright 2024 Google LLC
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

package containeranalysis

import (
	"context"
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	api "cloud.google.com/go/containeranalysis/apiv1"
	"google.golang.org/api/option"
	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containeranalysis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

func init() {
	registry.RegisterModel(krm.ContainerAnalysisNoteGVK, getNoteModel)
}

func getNoteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &noteModel{config: config}, nil
}

type noteModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &noteModel{}

type noteAdapter struct {
	resourceID string
	projectID  string

	desired *krm.ContainerAnalysisNote
	actual  *grafeaspb.Note

	client *api.Client
}

var _ directbase.Adapter = &noteAdapter{}

func (m *noteModel) client(ctx context.Context) (*api.Client, error) {
	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.GRPCUnaryClientInterceptor != nil {
		opts = append(opts, option.WithGRPCDialOption(grpc.WithUnaryInterceptor(m.config.GRPCUnaryClientInterceptor)))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	gcpClient, err := api.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building grafeas client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *noteModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	klog.FromContext(ctx).V(0).Info("creating adapter", "u", u)
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.ContainerAnalysisNote{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Get GCP Project
	projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project ID")
	}

	return &noteAdapter{
		resourceID: resourceID,
		projectID:  projectID,
		desired:    obj,
		client:     gcpClient,
	}, nil
}

func (m *noteModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &noteAdapter{}

// Find implements the Adapter interface.
func (a *noteAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}
	req := &grafeaspb.GetNoteRequest{Name: a.fullyQualifiedName()}
	note, err := a.client.GetGrafeasClient().GetNote(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	a.actual = note
	return true, nil
}

// Delete implements the Adapter interface.
func (a *noteAdapter) Delete(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &grafeaspb.DeleteNoteRequest{Name: a.fullyQualifiedName()}
	err := a.client.GetGrafeasClient().DeleteNote(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting note %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *noteAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	if a.projectID == "" {
		return fmt.Errorf("project is empty")
	}

	desired := a.desired.DeepCopy()
	note := Note_ToProto(desired)
	note.Name = a.fullyQualifiedName()

	log.V(0).Info("creating note", "note", note.Name)

	req := &grafeaspb.CreateNoteRequest{
		Parent: "projects/" + a.projectID,
		NoteId: a.resourceID,
		Note:   note,
	}
	createdNote, err := a.client.GetGrafeasClient().CreateNote(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to create note %s: %w", note.Name, err)
	}
	log.V(0).Info("created note", "note", createdNote.Name)

	// Set resourceID
	resourceID := a.resourceID
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	status := &krm.ContainerAnalysisNoteStatus{}
	status.CreateTime = ToOpenAPIDateTime(createdNote.CreateTime)
	status.UpdateTime = ToOpenAPIDateTime(createdNote.UpdateTime)

	return setStatus(u, status)
}

// Update implements the Adapter interface.
func (a *noteAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	desired := a.desired.DeepCopy()
	note := Note_ToProto(desired)
	note.Name = a.fullyQualifiedName()

	updateMask := &fieldmaskpb.FieldMask{}
	if note.ShortDescription != a.actual.ShortDescription {
		updateMask.Paths = append(updateMask.Paths, "short_description")
	}
	if note.LongDescription != a.actual.LongDescription {
		updateMask.Paths = append(updateMask.Paths, "long_description")
	}
	if !reflect.DeepEqual(note.RelatedUrl, a.actual.RelatedUrl) {
		updateMask.Paths = append(updateMask.Paths, "related_url")
	}

	if !reflect.DeepEqual(note.ExpirationTime, a.actual.ExpirationTime) {
		updateMask.Paths = append(updateMask.Paths, "expiration_time")
	}
	if !reflect.DeepEqual(note.RelatedNoteNames, a.actual.RelatedNoteNames) {
		updateMask.Paths = append(updateMask.Paths, "related_note_names")
	}
	if !reflect.DeepEqual(note.GetAttestation(), a.actual.GetAttestation()) {
		updateMask.Paths = append(updateMask.Paths, "attestation")
	}
	if len(updateMask.Paths) == 0 {
		klog.Warningf("unexpected empty update mask, desired: %v, actual: %v", a.desired, a.actual)
		return nil
	}

	req := &grafeaspb.UpdateNoteRequest{
		Name:       note.Name,
		Note:       note,
		UpdateMask: updateMask,
	}
	updatedNote, err := a.client.GetGrafeasClient().UpdateNote(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update note %s: %w", note.Name, err)
	}
	log.V(0).Info("updated note", "note", updatedNote.Name)

	status := &krm.ContainerAnalysisNoteStatus{}
	status.CreateTime = ToOpenAPIDateTime(updatedNote.GetCreateTime())
	status.UpdateTime = ToOpenAPIDateTime(updatedNote.GetUpdateTime())
	return setStatus(u, status)

}

func (a *noteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (a *noteAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/notes/%s", a.projectID, a.resourceID)
}
