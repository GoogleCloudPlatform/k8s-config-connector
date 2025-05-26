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

// +tool:controller
// proto.service: google.cloud.speech.v2.Speech
// proto.message: google.cloud.speech.v2.PhraseSet
// crd.type: SpeechPhraseSet
// crd.version: v1alpha1

package speech

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/speech/apiv2"
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.SpeechPhraseSetGVK, NewPhraseSetModel)
}

func NewPhraseSetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &phraseSetModel{config: *config}, nil
}

var _ directbase.Model = &phraseSetModel{}

type phraseSetModel struct {
	config config.ControllerConfig
}

func (m *phraseSetModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SpeechPhraseSet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewPhraseSetIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newSpeechClient(ctx)
	if err != nil {
		return nil, err
	}

	return &phraseSetAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *phraseSetModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type phraseSetAdapter struct {
	gcpClient *gcp.Client
	id        *krm.PhraseSetIdentity
	desired   *krm.SpeechPhraseSet
	actual    *pb.PhraseSet
	reader    client.Reader
}

var _ directbase.Adapter = &phraseSetAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *phraseSetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting speech phraseset", "name", a.id)

	req := &pb.GetPhraseSetRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetPhraseSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting speech phraseset %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *phraseSetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating speech phraseset", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := SpeechPhraseSetSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreatePhraseSetRequest{
		Parent:      a.id.Parent().String(),
		PhraseSetId: a.id.ID(),
		PhraseSet:   resource,
	}
	op, err := a.gcpClient.CreatePhraseSet(ctx, req)
	if err != nil {
		return fmt.Errorf("creating speech phraseset %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("speech phraseset %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created speech phraseset", "name", a.id)

	status := &krm.SpeechPhraseSetStatus{}
	status.ObservedState = SpeechPhraseSetObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *phraseSetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating speech phraseset", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := SpeechPhraseSetSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if !reflect.DeepEqual(resource.Phrases, a.actual.Phrases) {
		paths = append(paths, "phrases")
	}
	/* NOTYET
	if !reflect.DeepEqual(resource.Boost, a.actual.Boost) {
		paths = append(paths, "boost")
	}
	*/
	if !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		paths = append(paths, "display_name")
	}
	if !reflect.DeepEqual(resource.Annotations, a.actual.Annotations) {
		paths = append(paths, "annotations")
	}

	var updated *pb.PhraseSet
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdatePhraseSetRequest{
			PhraseSet:  resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdatePhraseSet(ctx, req)
		if err != nil {
			return fmt.Errorf("updating speech phraseset %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("speech phraseset %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated speech phraseset", "name", a.id)
	}

	status := &krm.SpeechPhraseSetStatus{}
	status.ObservedState = SpeechPhraseSetObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *phraseSetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SpeechPhraseSet{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SpeechPhraseSetSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.SpeechPhraseSetGVK)
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *phraseSetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting speech phraseset", "name", a.id)

	req := &pb.DeletePhraseSetRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePhraseSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent speech phraseset, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting speech phraseset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully initiated deletion speech phraseset", "name", a.id)

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete speech phraseset %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted speech phraseset", "name", a.id)
	return true, nil
}
