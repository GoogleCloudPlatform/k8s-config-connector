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
// proto.message: google.cloud.speech.v2.Recognizer
// crd.type: SpeechRecognizer
// crd.version: v1alpha1

package speech

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/speech/apiv2"
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.SpeechRecognizerGVK, NewRecognizerModel)
}

func NewRecognizerModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &recognizerModel{config: *config}, nil
}

var _ directbase.Model = &recognizerModel{}

type recognizerModel struct {
	config config.ControllerConfig
}

func (m *recognizerModel) client(ctx context.Context, projectID string) (*gcp.Client, error) {
	var opts []option.ClientOption

	config := m.config

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building speech recognizer client: %w", err)
	}

	return gcpClient, err
}

func (m *recognizerModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SpeechRecognizer{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewRecognizerIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &recognizerAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *recognizerModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type recognizerAdapter struct {
	gcpClient *gcp.Client
	id        *krm.RecognizerIdentity
	desired   *krm.SpeechRecognizer
	actual    *pb.Recognizer
	reader    client.Reader
}

var _ directbase.Adapter = &recognizerAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *recognizerAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting speech recognizer", "name", a.id)

	req := &pb.GetRecognizerRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetRecognizer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting speech recognizer %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// normalizeReferences resolves any resource references in the KRM resource.
func (a *recognizerAdapter) normalizeReferences(ctx context.Context) error {
	obj := a.desired
	// resolve resource references
	if obj.Spec.DefaultRecognitionConfig != nil {
		if err := normalizeRecognizerConfigRefs(ctx, a.reader, obj, obj.Spec.DefaultRecognitionConfig); err != nil {
			return err
		}
	}
	return nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *recognizerAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating speech recognizer", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := SpeechRecognizerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateRecognizerRequest{
		Parent:       a.id.Parent().String(),
		RecognizerId: a.id.ID(),
		Recognizer:   resource,
	}
	op, err := a.gcpClient.CreateRecognizer(ctx, req)
	if err != nil {
		return fmt.Errorf("creating speech recognizer %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("speech recognizer %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created speech recognizer in gcp", "name", a.id)

	status := &krm.SpeechRecognizerStatus{}
	status.ObservedState = SpeechRecognizerObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *recognizerAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating speech recognizer", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := SpeechRecognizerSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := []string{}
	if desired.Spec.DisplayName != nil && !reflect.DeepEqual(resource.DisplayName, a.actual.DisplayName) {
		paths = append(paths, "display_name")
	}
	if desired.Spec.Annotations != nil && !reflect.DeepEqual(resource.Annotations, a.actual.Annotations) {
		paths = append(paths, "annotations")
	}
	if desired.Spec.DefaultRecognitionConfig != nil && !recognizerConfigsEqual(resource.DefaultRecognitionConfig, a.actual.DefaultRecognitionConfig) {
		paths = append(paths, "default_recognition_config")
	}

	var updated *pb.Recognizer
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateRecognizerRequest{
			Recognizer: resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateRecognizer(ctx, req)
		if err != nil {
			return fmt.Errorf("updating speech recognizer %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("speech recognizer %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated speech recognizer", "name", a.id)
	}

	status := &krm.SpeechRecognizerStatus{}
	status.ObservedState = SpeechRecognizerObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *recognizerAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SpeechRecognizer{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SpeechRecognizerSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.SpeechRecognizerGVK)
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *recognizerAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting speech recognizer", "name", a.id)

	req := &pb.DeleteRecognizerRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteRecognizer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent speech recognizer, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting speech recognizer %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated deletion for speech recognizer", "name", a.id)

	_, err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of speech recognizer %s: %w", a.id, err)
	}
	return true, nil
}

// normalizeRecognizerConfigRefs resolves references in the RecognitionConfig message.
func normalizeRecognizerConfigRefs(ctx context.Context, reader client.Reader, obj client.Object, config *krm.RecognitionConfig) error {
	if config == nil {
		return nil
	}
	if config.Adaptation != nil {
		for i := range config.Adaptation.PhraseSets {
			if config.Adaptation.PhraseSets[i].PhraseSetRef != nil {
				if _, err := config.Adaptation.PhraseSets[i].PhraseSetRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
					return fmt.Errorf("normalizing PhraseSetRef: %w", err)
				}
			}
		}
	}
	return nil
}

// recognizerConfigsEqual compares two RecognitionConfig protos for equality, handling oneof fields.
func recognizerConfigsEqual(a, b *pb.RecognitionConfig) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Compare AutoDecodingConfig (oneof)
	aAuto, bAuto := a.GetDecodingConfig(), b.GetDecodingConfig()
	if (aAuto == nil) != (bAuto == nil) {
		return false
	}
	if aAuto != nil {
		_, aIsAuto := aAuto.(*pb.RecognitionConfig_AutoDecodingConfig)
		_, bIsAuto := bAuto.(*pb.RecognitionConfig_AutoDecodingConfig)
		if aIsAuto != bIsAuto {
			return false // Different types in oneof
		}
		// If both are AutoDecodingConfig, they are considered equal for this field
	}

	// Compare simple fields
	if a.Model != b.Model {
		return false
	}
	if !reflect.DeepEqual(a.LanguageCodes, b.LanguageCodes) {
		return false
	}

	// Compare nested messages (excluding Adaptation for now)
	if !reflect.DeepEqual(a.GetFeatures(), b.GetFeatures()) {
		return false
	}

	// Compare Adaptation (needs custom comparison)
	if !adaptationConfigsEqual(a.GetAdaptation(), b.GetAdaptation()) {
		return false
	}

	// Compare TranscriptNormalization (nested message)
	if !reflect.DeepEqual(a.GetTranscriptNormalization(), b.GetTranscriptNormalization()) {
		return false
	}

	// Compare TranslationConfig (nested message)
	if !reflect.DeepEqual(a.GetTranslationConfig(), b.GetTranslationConfig()) {
		return false
	}

	return true
}

// adaptationConfigsEqual compares two Adaptation protos for equality.
func adaptationConfigsEqual(a, b *pb.SpeechAdaptation) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Compare PhraseSets
	if !reflect.DeepEqual(a.GetPhraseSets(), b.GetPhraseSets()) {
		return false
	}

	// Compare CustomClasses
	if !reflect.DeepEqual(a.GetCustomClasses(), b.GetCustomClasses()) {
		return false
	}

	return true
}
