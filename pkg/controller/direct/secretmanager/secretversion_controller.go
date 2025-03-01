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

	"encoding/base64"

	gcp "cloud.google.com/go/secretmanager/apiv1"
	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.SecretManagerSecretVersionGVK, NewSecretVersionModel)
}

func NewSecretVersionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSecretVersion{config: *config}, nil
}

var _ directbase.Model = &modelSecretVersion{}

type modelSecretVersion struct {
	config config.ControllerConfig
}

func (m *modelSecretVersion) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SecretVersion client: %w", err)
	}
	return gcpClient, err
}

func (m *modelSecretVersion) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SecretManagerSecretVersion{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewSecretVersionIdentity(ctx, reader, obj, u)
	if err != nil {
		return nil, err
	}

	// Get secretmanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SecretVersionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelSecretVersion) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type SecretVersionAdapter struct {
	id        *krm.SecretVersionIdentity
	gcpClient *gcp.Client
	desired   *krm.SecretManagerSecretVersion
	actual    *pb.SecretVersion
	reader    client.Reader
}

var _ directbase.Adapter = &SecretVersionAdapter{}

func (a *SecretVersionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SecretManagerSecretVersion", "name", a.id)

	if !a.id.HasKnownID() {
		return false, nil
	}

	req := &pb.GetSecretVersionRequest{Name: a.id.String()}
	secretVersion, err := a.gcpClient.GetSecretVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SecretManagerSecretVersion %q: %w", a.id, err)
	}
	a.actual = secretVersion
	return true, nil
}

func (a *SecretVersionAdapter) normalizeSecretData(ctx context.Context) ([]byte, error) {
	if a.desired.Spec.SecretData == nil {
		return nil, fmt.Errorf("SecretManagerSecretVersion is service generated object."+
			" Creating a new SecretVersion requires `spec.secretData` "+
			"Acquiring an existing SecretVersion requires `spec.resourceID`: %s", a.desired.GetName())
	}
	plain := a.desired.Spec.SecretData.Value
	secretRef := a.desired.Spec.SecretData.ValueFrom
	if plain != nil && secretRef != nil {
		return nil, fmt.Errorf("either spec.secretData.Value or spec.secretData.ValueFrom is required")
	}

	if plain != nil {
		data := base64.StdEncoding.EncodeToString([]byte(*plain))
		return []byte(data), nil
	}
	return refsv1beta1secret.NormalizedLegacySecret(ctx, secretRef.SecretKeyRef, a.reader, a.desired.Namespace)
}

func (a *SecretVersionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SecretVersion", "name", a.id)
	mapCtx := &direct.MapContext{}

	data, err := a.normalizeSecretData(ctx)
	if err != nil {
		return err
	}
	secretPayload := &pb.SecretPayload{
		Data: data,
		// TODO: safer option to pass in the checksum.
		// DataCrc32C:  *int64,
	}
	req := &pb.AddSecretVersionRequest{
		Parent:  a.id.Parent().String(),
		Payload: secretPayload,
	}
	created, err := a.gcpClient.AddSecretVersion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SecretVersion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created SecretVersion", "name", a.id)

	// The default status for newly created resource is "enabled".
	// This saves the waiting time for the next reconciliation.
	if !*a.desired.Spec.Enabled {
		log.V(2).Info("disabling the secret version", "name", a.id)
		req := &pb.DisableSecretVersionRequest{
			Name: created.Name,
			Etag: created.Etag,
		}
		created, err = a.gcpClient.DisableSecretVersion(ctx, req)
		if err != nil {
			return fmt.Errorf("disable SecretVersion %s: %w", a.id, err)
		}
	}
	status := &krm.SecretManagerSecretVersionStatus{}
	status.ObservedState = SecretManagerSecretVersionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if err := updateLegacyFields(status); err != nil {
		return nil
	}
	status.ExternalRef = direct.LazyPtr(a.id.Parent().String() + "/versions/" + *status.Version)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *SecretVersionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SecretVersion", "name", a.id)

	var updated *pb.SecretVersion
	var err error

	switch a.actual.State {
	case pb.SecretVersion_ENABLED:
		if !*a.desired.Spec.Enabled {
			log.V(2).Info("disabling the secret version", "name", a.id)
			req := &pb.DisableSecretVersionRequest{
				Name: a.id.String(),
				Etag: a.actual.Etag,
			}
			updated, err = a.gcpClient.DisableSecretVersion(ctx, req)
			if err != nil {
				return fmt.Errorf("disable SecretVersion %s: %w", a.id, err)
			}
		} else {
			log.V(2).Info("already disabled", "name", a.id)
		}
	case pb.SecretVersion_DISABLED:
		if *a.desired.Spec.Enabled {
			log.V(2).Info("enabling the secret version", "name", a.id)
			req := &pb.EnableSecretVersionRequest{
				Name: a.id.String(),
				Etag: a.actual.Etag,
			}
			updated, err = a.gcpClient.EnableSecretVersion(ctx, req)
			if err != nil {
				return fmt.Errorf("enable SecretVersion %s: %w", a.id, err)
			}
		} else {
			log.V(2).Info("already enabled", "name", a.id)
		}

	case pb.SecretVersion_DESTROYED:
		// If a version is destroyed, it cannot be re-enabled or re-disabled.
		return fmt.Errorf("already destroyed %s", a.id)
	case pb.SecretVersion_STATE_UNSPECIFIED:
		return fmt.Errorf("version unspecified %s", a.id)
	}

	mapCtx := &direct.MapContext{}
	status := &krm.SecretManagerSecretVersionStatus{}
	status.ObservedState = SecretManagerSecretVersionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if err := updateLegacyFields(status); err != nil {
		return nil
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *SecretVersionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.SecretManagerSecretVersion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(SecretManagerSecretVersionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	id, err := krm.ParseSecretVersionExternal(a.actual.Name)
	if err != nil {
		return nil, err
	}
	u.SetName(id.ID())
	u.SetGroupVersionKind(krm.SecretManagerSecretVersionGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *SecretVersionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	// If the secret version is already in DESTROYED state, no need to call DestroySecretVersion.
	if a.actual != nil && a.actual.State == pb.SecretVersion_DESTROYED {
		log.Info("SecretVersion already in DESTROYED state", "name", a.id)
		return true, nil
	}

	log.Info("destroying SecretVersion", "name", a.id)
	req := &pb.DestroySecretVersionRequest{
		Name: a.id.String(), Etag: a.actual.Etag}
	_, err := a.gcpClient.DestroySecretVersion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting SecretVersion %s: %w", a.id, err)
	}
	log.Info("destroyed SecretVersion", "name", a.id)
	return true, nil
}

func updateLegacyFields(status *krm.SecretManagerSecretVersionStatus) error {
	if status.ObservedState == nil {
		return nil
	}
	status.CreateTime = status.ObservedState.CreateTime
	status.DestroyTime = status.ObservedState.DestroyTime
	status.Name = status.ObservedState.Name

	id, err := krm.ParseSecretVersionExternal(*status.ObservedState.Name)
	if err != nil {
		return err
	}
	version := id.ID()
	status.Version = &version
	return nil
}
