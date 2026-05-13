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

package secretmanager

import (
	"context"
	"fmt"
	"reflect"
	"regexp"

	gcp "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/go-logr/logr"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	regionalCtrlName = "secretmanager-controller"
)

func init() {
	registry.RegisterModel(krm.SecretManagerRegionalSecretGVK, NewRegionalSecretModel)
}

func NewRegionalSecretModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &regionalSecretModel{config: *config}, nil
}

var _ directbase.Model = &regionalSecretModel{}

type regionalSecretModel struct {
	config config.ControllerConfig
}

func (m *regionalSecretModel) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Secret client: %w", err)
	}
	return gcpClient, err
}

func (m *regionalSecretModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.SecretManagerRegionalSecret{}

	copied := u.DeepCopy()
	if err := label.ComputeLabels(copied); err != nil {
		return nil, err
	}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(copied.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	typedID, ok := id.(*krm.SecretManagerRegionalSecretIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *krm.SecretManagerRegionalSecretIdentity, got %T", id)
	}

	if err = regionalNormalizeExternal(ctx, reader, copied, obj); err != nil {
		return nil, err
	}

	// Get secretmanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &RegionalSecretAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *regionalSecretModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type RegionalSecretAdapter struct {
	id        *krm.SecretManagerRegionalSecretIdentity
	gcpClient *gcp.Client
	desired   *krm.SecretManagerRegionalSecret
	actual    *secretmanagerpb.Secret
}

var _ directbase.Adapter = &RegionalSecretAdapter{}

func regionalNormalizeExternal(ctx context.Context, reader client.Reader, src client.Object, secret *krm.SecretManagerRegionalSecret) error {
	if secret.Spec.CustomerManagedEncryption != nil {
		kmsKeyRef := secret.Spec.CustomerManagedEncryption.KmsKeyRef
		kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, src, kmsKeyRef)
		if err != nil {
			return err
		}
		secret.Spec.CustomerManagedEncryption.KmsKeyRef = kmsKeyRef
	}
	if len(secret.Spec.TopicRefs) != 0 {
		for _, topicRef := range secret.Spec.TopicRefs {
			if topicRef.PubSubTopicRef != nil {
				_, err := topicRef.PubSubTopicRef.NormalizedExternal(ctx, reader, src.GetNamespace())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (a *RegionalSecretAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecretManagerRegionalSecret", "name", a.id)

	req := &secretmanagerpb.GetSecretRequest{Name: a.id.String()}
	secretpb, err := a.gcpClient.GetSecret(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecretManagerRegionalSecret %q: %w", a.id, err)
	}

	a.actual = secretpb
	return true, nil
}

func regionalMergeMap(a, b map[string]string) map[string]string {
	copy := make(map[string]string, len(a))
	for k, v := range a {
		copy[k] = v
	}
	for k, v := range b {
		copy[k] = v
	}
	return copy
}

// Annotation keys must be between 1 and 63 characters long
// have a UTF-8 encoding of maximum 128 bytes
// begin and end with an alphanumeric character ([a-z0-9A-Z]),
// may have dashes (-), underscores (_), dots (.), and alphanumerics in between these symbols.
func regionalIsValidAnnotation(s string) bool {
	var validPattern = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\.\-_]*[a-zA-Z0-9]$`)

	// A string is valid if it matches the pattern AND its length is between 1 and 63.
	return validPattern.MatchString(s) && len(s) >= 1 && len(s) <= 63
}

func regionalComputeAnnotations(secret *krm.SecretManagerRegionalSecret, log *logr.Logger) map[string]string {
	annotations := regionalMergeMap(secret.GetAnnotations(), secret.Spec.Annotations)
	for key := range annotations {
		if !regionalIsValidAnnotation(key) {
			log.V(2).Info("Remove annotation with invalid key", "key", key)
			delete(annotations, key)
		}
	}
	return annotations
}

func (a *RegionalSecretAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Secret", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecretManagerRegionalSecretSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Annotations = regionalComputeAnnotations(desired, &log)
	// GCP service does not allow setting version aliases during Secret creation.
	resource.VersionAliases = nil
	req := &secretmanagerpb.CreateSecretRequest{
		Parent:   "projects/" + a.id.Project + "/locations/" + a.id.Location,
		SecretId: a.id.Secret,
		Secret:   resource,
	}
	created, err := a.gcpClient.CreateSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Secret %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Secret", "name", a.id)

	status := &krm.SecretManagerRegionalSecretStatus{}
	status.ObservedState = SecretManagerRegionalSecretStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	external := a.id.String()
	status.ExternalRef = &external
	status.Name = created.Name
	if err = op.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}
	// VersionAliases cannot be set in Creation, requeing the result to update the versionAlias without waiting for the reconciliation interval.
	if a.desired.Spec.VersionAliases != nil {
		op.RequeueRequested = true
	}
	return nil
}

func regionalTopicsEqual(desired []*krm.TopicRef, actual []*secretmanagerpb.Topic) bool {

	externalsDesired := sets.Set[string]{}
	externalsActual := sets.Set[string]{}
	for _, topicRef := range desired {
		externalsDesired.Insert(topicRef.PubSubTopicRef.External)
	}
	for _, topic := range actual {
		externalsActual.Insert(topic.GetName())
	}
	return reflect.DeepEqual(sets.List(externalsDesired), sets.List(externalsActual))
}

func (a *RegionalSecretAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Secret", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecretManagerRegionalSecretSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Annotations = regionalComputeAnnotations(desired, &log)
	// the GCP service use *name* to identify the resource.
	resource.Name = a.id.String()
	resource.Etag = a.actual.Etag
	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if paths.Has("ttl") {
		paths = paths.Delete("ttl")
		resource.Expiration = a.actual.Expiration
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	report := &structuredreporting.Diff{Object: op.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &secretmanagerpb.UpdateSecretRequest{
		UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		Secret:     resource,
	}
	updated, err := a.gcpClient.UpdateSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("Secret %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Secret", "name", a.id)

	status := &krm.SecretManagerRegionalSecretStatus{}
	status.ObservedState = SecretManagerRegionalSecretStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Set externalRef. This field is empty after migration from the TF-based controller to the direct.
	external := a.id.String()
	status.ExternalRef = &external

	status.Name = updated.Name
	return op.UpdateStatus(ctx, status, nil)
}

func (a *RegionalSecretAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecretManagerRegionalSecret{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecretManagerRegionalSecretSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the RegionalSecretAdapter interface.
func (a *RegionalSecretAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Secret", "name", a.id)

	req := &secretmanagerpb.DeleteSecretRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSecret(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Secret %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Secret", "name", a.id)
	return true, nil
}
