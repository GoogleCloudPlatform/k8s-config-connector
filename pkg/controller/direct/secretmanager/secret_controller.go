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
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"

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
	ctrlName = "secretmanager-controller"
)

func init() {
	rg := &SecretReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.SecretManagerSecretGVK, NewModel, rg)
}

type SecretReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &SecretReconcileGate{}

func (r *SecretReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	if r.optIn.ShouldReconcile(o) {
		return true
	}
	obj := &krm.SecretManagerSecret{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &obj); err != nil {
		return false
	}
	return obj.Spec.Labels != nil
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &secretModel{config: *config}, nil
}

var _ directbase.Model = &secretModel{}

type secretModel struct {
	config config.ControllerConfig
}

func (m *secretModel) client(ctx context.Context) (*gcp.Client, error) {
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

func (m *secretModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SecretManagerSecret{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSecretIdentity(ctx, reader, obj, u)
	if err != nil {
		return nil, err
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

func (m *secretModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.SecretIdentity
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
		if secret.Spec.Replication.UserManaged != nil {
			for _, r := range secret.Spec.Replication.UserManaged.Replicas {
				if r.CustomerManagedEncryption != nil {
					kmsKeyRef := r.CustomerManagedEncryption.KmsKeyRef
					kmsKeyRef, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, src, kmsKeyRef)
					if err != nil {
						return err
					}
					r.CustomerManagedEncryption.KmsKeyRef = kmsKeyRef
				}
			}
		}
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

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting SecretManagerSecret", "name", a.id)

	req := &secretmanagerpb.GetSecretRequest{Name: a.id.String()}
	secretpb, err := a.gcpClient.GetSecret(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecretManagerSecret %q: %w", a.id, err)
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

func (a *Adapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Secret", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecretManagerSecretSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Annotations = ComputeAnnotations(desired)
	// GCP service does not allow setting version aliases during Secret creation.
	resource.VersionAliases = nil
	req := &secretmanagerpb.CreateSecretRequest{
		Parent:   a.id.Parent().String(),
		SecretId: a.id.ID(),
		Secret:   resource,
	}
	created, err := a.gcpClient.CreateSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Secret %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Secret", "name", a.id)

	status := &krm.SecretManagerSecretStatus{}
	status.ObservedState = SecretManagerSecretStatusObservedState_FromProto(mapCtx, created)
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

func (a *Adapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Secret", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := SecretManagerSecretSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Annotations = ComputeAnnotations(desired)
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

	req := &secretmanagerpb.UpdateSecretRequest{
		UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		Secret:     resource,
	}
	updated, err := a.gcpClient.UpdateSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("Secret %s waiting update: %w", a.id, err)
	}
	log.V(2).Info("successfully updated Secret", "name", a.id)

	status := &krm.SecretManagerSecretStatus{}
	status.ObservedState = SecretManagerSecretStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Set externalRef. This field is empty after migration from the TF-based controller to the direct.
	external := a.id.String()
	status.ExternalRef = &external

	status.Name = updated.Name
	return op.UpdateStatus(ctx, status, nil)
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
	log.V(2).Info("deleting Secret", "name", a.id)

	req := &secretmanagerpb.DeleteSecretRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSecret(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Secret %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Secret", "name", a.id)
	return true, nil
}
