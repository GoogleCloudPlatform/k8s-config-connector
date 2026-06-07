/*
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudbuild

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/cloudbuild/apiv2"
	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv2/cloudbuildpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.CloudBuildConnectionGVK, NewConnectionModel)
}

func NewConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelConnection{config: *config}, nil
}

var _ directbase.Model = &modelConnection{}

type modelConnection struct {
	config config.ControllerConfig
}

func (m *modelConnection) client(ctx context.Context) (*gcp.RepositoryManagerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRepositoryManagerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudbuild connection client: %w", err)
	}
	return gcpClient, err
}

func (m *modelConnection) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudBuildConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	typedID, ok := id.(*krm.CloudBuildConnectionIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *krm.CloudBuildConnectionIdentity, got %T", id)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ConnectionAdapter{
		id:        typedID,
		gcpClient: gcpClient,
		reader:    reader,
		desired:   obj,
	}, nil
}

func (m *modelConnection) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //cloudbuild.googleapis.com/projects/<project>/locations/<location>/connections/<id>
	if !strings.HasPrefix(url, "//cloudbuild.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//cloudbuild.googleapis.com/"), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connections" {
		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}

		return &ConnectionAdapter{
			id: &krm.CloudBuildConnectionIdentity{
				Project:    tokens[1],
				Location:   tokens[3],
				Connection: tokens[5],
			},
			gcpClient: gcpClient,
		}, nil
	}

	return nil, nil
}

type ConnectionAdapter struct {
	id        *krm.CloudBuildConnectionIdentity
	gcpClient *gcp.RepositoryManagerClient
	reader    client.Reader
	desired   *krm.CloudBuildConnection
	actual    *cloudbuildpb.Connection
}

var _ directbase.Adapter = &ConnectionAdapter{}

func (a *ConnectionAdapter) Find(ctx context.Context) (bool, error) {
	req := &cloudbuildpb.GetConnectionRequest{Name: a.id.String()}
	connectionpb, err := a.gcpClient.GetConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting cloudbuildconnection %q: %w", a.id.String(), err)
	}

	a.actual = connectionpb
	return true, nil
}

func (a *ConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	err := a.resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating connection", "u", u)

	desired := a.desired.DeepCopy()

	mapCtx := &direct.MapContext{}
	conn := CloudBuildConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	conn.Name = a.id.String()
	req := &cloudbuildpb.CreateConnectionRequest{
		Parent:       fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		ConnectionId: a.id.Connection,
		Connection:   conn,
	}
	op, err := a.gcpClient.CreateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildconnection %s creating failed: %w", conn.Name, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("cloudbuildconnection %s waiting creation failed: %w", conn.Name, err)
	}

	status := &krm.CloudBuildConnectionStatus{}
	status.ObservedState = CloudBuildConnectionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return setStatus(u, status)
}

func (a *ConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	err := a.resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)

	desired := a.desired.DeepCopy()
	mapCtx := &direct.MapContext{}
	conn := CloudBuildConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	conn.Name = a.id.String()
	conn.Etag = a.actual.Etag

	paths, err := common.CompareProtoMessage(conn, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &cloudbuildpb.UpdateConnectionRequest{
		Connection: conn,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}
	op, err := a.gcpClient.UpdateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildconnection %s updating failed: %w", conn.Name, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("cloudbuildconnection %s waiting update failed: %w", conn.Name, err)
	}
	status := &krm.CloudBuildConnectionStatus{}
	status.ObservedState = CloudBuildConnectionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return fmt.Errorf("update connection status %w", mapCtx.Err())
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return setStatus(u, status)
}

func (a *ConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudBuildConnection{}
	obj.SetGroupVersionKind(krm.CloudBuildConnectionGVK)
	obj.SetName(a.actual.Name)

	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudBuildConnectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	return u, nil
}

func (a *ConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &cloudbuildpb.DeleteConnectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConnection(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting cloudbuildconnection %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete cloudbuildconnection %s: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *ConnectionAdapter) resolveDependencies(ctx context.Context, reader client.Reader, obj *krm.CloudBuildConnection) error {
	spec := &obj.Spec
	if spec.GithubConfig != nil && spec.GithubConfig.AuthorizerCredential != nil {
		ref := spec.GithubConfig.AuthorizerCredential.OauthTokenSecretVersionRef
		resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, ref)
		if err != nil {
			return err
		}
		spec.GithubConfig.AuthorizerCredential.OauthTokenSecretVersionRef = resolved
	}
	if spec.GithubEnterpriseConfig != nil {
		if spec.GithubEnterpriseConfig.PrivateKeySecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.GithubEnterpriseConfig.PrivateKeySecretVersionRef)
			if err != nil {
				return err
			}
			spec.GithubEnterpriseConfig.PrivateKeySecretVersionRef = resolved
		}
		if spec.GithubEnterpriseConfig.WebhookSecretSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.GithubEnterpriseConfig.WebhookSecretSecretVersionRef)
			if err != nil {
				return err
			}
			spec.GithubEnterpriseConfig.WebhookSecretSecretVersionRef = resolved
		}
	}
	if spec.GitlabConfig != nil {
		if spec.GitlabConfig.WebhookSecretSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.GitlabConfig.WebhookSecretSecretVersionRef)
			if err != nil {
				return err
			}
			spec.GitlabConfig.WebhookSecretSecretVersionRef = resolved
		}
		if spec.GitlabConfig.ReadAuthorizerCredential != nil && spec.GitlabConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.GitlabConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef)
			if err != nil {
				return err
			}
			spec.GitlabConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef = resolved
		}
		if spec.GitlabConfig.AuthorizerCredential != nil && spec.GitlabConfig.AuthorizerCredential.UserTokenSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.GitlabConfig.AuthorizerCredential.UserTokenSecretVersionRef)
			if err != nil {
				return err
			}
			spec.GitlabConfig.AuthorizerCredential.UserTokenSecretVersionRef = resolved
		}
	}
	if spec.BitbucketDataCenterConfig != nil {
		if spec.BitbucketDataCenterConfig.WebhookSecretSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.BitbucketDataCenterConfig.WebhookSecretSecretVersionRef)
			if err != nil {
				return err
			}
			spec.BitbucketDataCenterConfig.WebhookSecretSecretVersionRef = resolved
		}
		if spec.BitbucketDataCenterConfig.ReadAuthorizerCredential != nil && spec.BitbucketDataCenterConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.BitbucketDataCenterConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef)
			if err != nil {
				return err
			}
			spec.BitbucketDataCenterConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef = resolved
		}
		if spec.BitbucketDataCenterConfig.AuthorizerCredential != nil && spec.BitbucketDataCenterConfig.AuthorizerCredential.UserTokenSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.BitbucketDataCenterConfig.AuthorizerCredential.UserTokenSecretVersionRef)
			if err != nil {
				return err
			}
			spec.BitbucketDataCenterConfig.AuthorizerCredential.UserTokenSecretVersionRef = resolved
		}
	}
	if spec.BitbucketCloudConfig != nil {
		if spec.BitbucketCloudConfig.WebhookSecretSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.BitbucketCloudConfig.WebhookSecretSecretVersionRef)
			if err != nil {
				return err
			}
			spec.BitbucketCloudConfig.WebhookSecretSecretVersionRef = resolved
		}
		if spec.BitbucketCloudConfig.ReadAuthorizerCredential != nil && spec.BitbucketCloudConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.BitbucketCloudConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef)
			if err != nil {
				return err
			}
			spec.BitbucketCloudConfig.ReadAuthorizerCredential.UserTokenSecretVersionRef = resolved
		}
		if spec.BitbucketCloudConfig.AuthorizerCredential != nil && spec.BitbucketCloudConfig.AuthorizerCredential.UserTokenSecretVersionRef != nil {
			resolved, err := refs.ResolveSecretManagerSecretVersionRef(ctx, reader, obj, spec.BitbucketCloudConfig.AuthorizerCredential.UserTokenSecretVersionRef)
			if err != nil {
				return err
			}
			spec.BitbucketCloudConfig.AuthorizerCredential.UserTokenSecretVersionRef = resolved
		}
	}
	return nil
}
