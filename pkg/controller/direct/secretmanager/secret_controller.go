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

	gcp "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName      = "secretmanager-controller"
	serviceDomain = "//secretmanager.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.SecretManagerSecretGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
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

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SecretManagerSecret{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// TODO: add spec.projectID
	projectID, err := refs.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return nil, err
	}

	var id *SecretManagerSecretIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Project != projectID {
			return nil, fmt.Errorf("SecretManagerSecret %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Project, projectID)
		}
		if id.Secret != resourceID {
			return nil, fmt.Errorf("SecretManagerSecret  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Secret, resourceID)
		}
	}

	if err = normalizeExternal(ctx, reader, u, obj); err != nil {
		return nil, err
	}

	// Get secretmanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *SecretManagerSecretIdentity
	gcpClient *gcp.Client
	desired   *krm.SecretManagerSecret
	actual    *secretmanagerpb.Secret
}

var _ directbase.Adapter = &Adapter{}

func normalizeExternal(ctx context.Context, reader client.Reader, src client.Object, secret *krm.SecretManagerSecret) error {
	if secret.Spec.Replication != nil {
		if secret.Spec.Replication.LegacyAutomatic != nil {
			if secret.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption != nil {
				kmsKeyRef := secret.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption.KmsKeyRef

				kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, src, kmsKeyRef)
				if err != nil {
					return err
				}
				secret.Spec.Replication.LegacyAutomatic.CustomerManagedEncryption.KmsKeyRef = kmsKeyRef
			}
		}
	}
	if len(secret.Spec.TopicRefs) != 0 {
		for _, topicRef := range secret.Spec.TopicRefs {
			if topicRef.PubSubTopicRef != nil {
				pubsubRef, err := refs.ResolvePubSubTopic(ctx, reader, src, topicRef.PubSubTopicRef)
				if err != nil {
					return err
				}
				topicRef.PubSubTopicRef.External = pubsubRef.String()
			}
		}
	}
	return nil

}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting SecretManagerSecret", "name", a.id.FullyQualifiedName())

	req := &secretmanagerpb.GetSecretRequest{Name: a.id.FullyQualifiedName()}
	secretpb, err := a.gcpClient.GetSecret(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecretManagerSecret %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = secretpb
	return true, nil
}

func MergeMap(a, b map[string]string) map[string]string {
	copy := make(map[string]string, len(a))
	for k, v := range a {
		copy[k] = v
	}
	for k, v := range b {
		copy[k] = v
	}
	return copy
}

func ComputeAnnotations(secret *krm.SecretManagerSecret) map[string]string {
	annotations := MergeMap(secret.GetAnnotations(), secret.Spec.Annotations)
	common.RemoveByPrefixes(annotations, "cnrm.cloud.google.com", "alpha.cnrm.cloud.google.com")
	return annotations
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Secret", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecretManagerSecretSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Annotations = ComputeAnnotations(desired)
	resource.Labels = common.ComputeGCPLabels(desired.GetLabels())

	// Add metadata
	req := &secretmanagerpb.CreateSecretRequest{
		Parent:   a.id.Parent.String(),
		SecretId: a.id.Secret,
		Secret:   resource,
	}
	created, err := a.gcpClient.CreateSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Secret %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created Secret", "name", a.id.FullyQualifiedName())

	status := &krm.SecretManagerSecretStatus{}
	status.ObservedState = SecretManagerSecretStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	status.Name = created.Name
	return setStatus(u, status)
}

func topicsEqual(desired []*krm.TopicRef, actual []*secretmanagerpb.Topic) bool {

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

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Secret", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecretManagerSecretSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// the GCP service use *name* to identify the resource.
	resource.Name = a.id.FullyQualifiedName()

	resource.Annotations = ComputeAnnotations(desired)
	resource.Labels = common.ComputeGCPLabels(desired.GetLabels())

	updateMask := &fieldmaskpb.FieldMask{}

	if !reflect.DeepEqual(resource.Annotations, a.actual.Annotations) {
		updateMask.Paths = append(updateMask.Paths, "annotations")
	}
	if !topicsEqual(desired.Spec.TopicRefs, a.actual.GetTopics()) {
		updateMask.Paths = append(updateMask.Paths, "topics")
	}
	if !common.DeepEqual_StringAndTimestampPb(*desired.Spec.ExpireTime, a.actual.GetExpireTime()) {
		updateMask.Paths = append(updateMask.Paths, "expire_time")
	}
	if !reflect.DeepEqual(desired.Spec.Rotation.NextRotationTime, a.actual.GetRotation().GetNextRotationTime()) {
		updateMask.Paths = append(updateMask.Paths, "rotation.next_rotation_time")
	}
	if !reflect.DeepEqual(desired.Spec.Rotation.RotationPeriod, a.actual.GetRotation().GetRotationPeriod()) {
		updateMask.Paths = append(updateMask.Paths, "rotation.rotation_period")
	}
	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.FullyQualifiedName())
		return nil
	}

	req := &secretmanagerpb.UpdateSecretRequest{
		UpdateMask: updateMask,
		Secret:     resource,
	}
	updated, err := a.gcpClient.UpdateSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("Secret %s waiting update: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated Secret", "name", a.id.FullyQualifiedName())

	status := &krm.SecretManagerSecretStatus{}
	status.ObservedState = SecretManagerSecretStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.Name = updated.Name
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecretManagerSecret{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecretManagerSecretSpec_FromProto(mapCtx, a.actual))
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

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting Secret", "name", a.id.FullyQualifiedName())

	req := &secretmanagerpb.DeleteSecretRequest{Name: a.id.FullyQualifiedName()}
	err := a.gcpClient.DeleteSecret(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Secret %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted Secret", "name", a.id.FullyQualifiedName())
	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
		status["name"] = old["name"]
	}

	u.Object["status"] = status

	return nil
}
