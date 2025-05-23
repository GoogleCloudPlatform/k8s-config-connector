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

package bigqueryconnection

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/bigquery/connection/apiv1"
	pb "cloud.google.com/go/bigquery/connection/apiv1/connectionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
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
	ctrlName      = "bigqueryconnection-controller"
	serviceDomain = "//bigqueryconnection.googleapis.com"
)

func init() {
	registry.RegisterModel(krm.BigQueryConnectionConnectionGVK, NewModel)
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
		return nil, fmt.Errorf("building bigqueryconnection client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryConnectionConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	connectionRef, err := krm.NewBigQueryConnectionConnectionRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	// Get bigqueryconnection GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        connectionRef,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
		namespace: obj.Namespace,
	}, nil
}

func (a *Adapter) normalizeReference(ctx context.Context) error {
	obj := a.desired
	// Resolve SQLInstanceRef and SQLDatabaseRef
	if obj.Spec.CloudSQLSpec != nil {
		sql := obj.Spec.CloudSQLSpec
		if sql.InstanceRef != nil {
			_, err := sql.InstanceRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
			if err != nil {
				return err
			}
		}
		if sql.DatabaseRef != nil {
			_, err := sql.DatabaseRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
			if err != nil {
				return err
			}
		}
		if sql.Credential != nil {
			if err := refsv1beta1secret.NormalizedSecret(ctx, sql.Credential.SecretRef, a.reader, a.namespace); err != nil {
				return err
			}
		}
	}

	// Resolve SpannerDatabaseRef
	if obj.Spec.CloudSpannerSpec != nil {
		if obj.Spec.CloudSpannerSpec.DatabaseRef != nil {
			database, err := refs.ResolveSpannerDatabaseRef(ctx, a.reader, obj, obj.Spec.CloudSpannerSpec.DatabaseRef)
			if err != nil {
				return err
			}
			obj.Spec.CloudSpannerSpec.DatabaseRef.External = database.String()
		}
	}

	// Resolve Spark.DataprocClusterRef and Spark.MetastoreServiceRef
	if obj.Spec.SparkSpec != nil {
		if obj.Spec.SparkSpec.SparkHistoryServer != nil {
			if obj.Spec.SparkSpec.SparkHistoryServer.DataprocClusterRef != nil {
				cluster, err := refs.ResolveDataprocClusterRef(ctx, a.reader, obj, obj.Spec.SparkSpec.SparkHistoryServer.DataprocClusterRef)
				if err != nil {
					return err
				}
				obj.Spec.SparkSpec.SparkHistoryServer.DataprocClusterRef.External = cluster.String()
			}
		}

		if obj.Spec.SparkSpec.MetastoreService != nil {
			if obj.Spec.SparkSpec.MetastoreService.MetastoreServiceRef != nil {
				service, err := refs.ResolveMetastoreServiceRef(ctx, a.reader, obj, obj.Spec.SparkSpec.MetastoreService.MetastoreServiceRef)
				if err != nil {
					return err
				}
				obj.Spec.SparkSpec.MetastoreService.MetastoreServiceRef.External = service.String()
			}
		}
	}
	return nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.BigQueryConnectionConnectionRef
	gcpClient *gcp.Client
	desired   *krm.BigQueryConnectionConnection
	actual    *pb.Connection
	reader    client.Reader
	namespace string
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)

	log.V(2).Info("getting BigQueryConnectionConnection", "name", a.id.External)

	_, idIsSet, err := a.id.ConnectionID()
	if err != nil {
		return false, err
	}
	if !idIsSet { // resource is not yet created
		return false, nil
	}
	req := &pb.GetConnectionRequest{Name: a.id.External}
	connectionpb, err := a.gcpClient.GetConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryConnectionConnection %q: %w", a.id.External, err)
	}

	a.actual = connectionpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Connection", "name", a.id.External)

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}
	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := BigQueryConnectionConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent()
	if err != nil {
		return err
	}
	req := &pb.CreateConnectionRequest{
		Parent:     parent,
		Connection: resource,
	}
	id, isIsSet, err := a.id.ConnectionID()
	if err != nil {
		return err
	}
	if isIsSet { // during "Create", this means user has specified connection ID in `spec.ResourceID` field.
		req = &pb.CreateConnectionRequest{
			Parent:       parent,
			ConnectionId: id,
			Connection:   resource,
		}
	}
	created, err := a.gcpClient.CreateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Connection %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created Connection", "name", created.Name)

	status := &krm.BigQueryConnectionConnectionStatus{}
	status.ObservedState = BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	tokens := strings.Split(created.Name, "/")
	parent, err = a.id.Parent()
	if err != nil {
		return err
	}
	externalRef := parent + "/connections/" + tokens[5]
	status.ExternalRef = &externalRef
	return setStatus(u, status)
}

// aws.accessRole.Identity is a output-only field in CREATE, it is required in the UPDATE.
func processAwsIdentityMerge(desired *pb.Connection, previouslyApplied *krm.BigQueryConnectionConnectionObservedState) *pb.Connection {
	if previouslyApplied != nil && previouslyApplied.Aws != nil {
		desired.GetAws().GetAccessRole().Identity = direct.ValueOf(previouslyApplied.Aws.AccessRole.Identity)
	}
	return desired
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Connection", "name", a.id.External)

	if err := a.normalizeReference(ctx); err != nil {
		return err
	}
	mapCtx := &direct.MapContext{}
	connection := BigQueryConnectionConnectionSpec_ToProto(mapCtx, &a.desired.Spec)
	connection = processAwsIdentityMerge(connection, a.desired.Status.ObservedState)

	connection.Name = a.actual.Name
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	paths, err := common.CompareProtoMessage(connection, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		return nil
	}
	fqn := a.id.External
	req := &pb.UpdateConnectionRequest{
		Name:       fqn,
		Connection: connection,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}
	updated, err := a.gcpClient.UpdateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Connection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully updated Connection", "name", fqn)

	status := &krm.BigQueryConnectionConnectionStatus{}
	status.ObservedState = BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryConnectionConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryConnectionConnectionSpec_FromProto(mapCtx, a.actual))
	tokens := strings.Split(a.id.External, "connections/")
	obj.Spec.ResourceID = direct.LazyPtr(tokens[len(tokens)-1])
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent, err := a.id.Parent()
	if err != nil {
		return nil, fmt.Errorf("BigQueryConnectionConnection %s parent unset: %w", a.id.External, err)
	}
	if parent != "" {
		tokens := strings.Split(parent, "/")
		if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
			obj.Spec.ProjectRef = &refs.ProjectRef{Name: tokens[1]}
			obj.Spec.Location = tokens[3]
		}
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
	log.V(2).Info("deleting Connection", "name", a.id.External)

	fqn := a.id.External
	req := &pb.DeleteConnectionRequest{Name: fqn}
	if err := a.gcpClient.DeleteConnection(ctx, req); err != nil {
		return false, fmt.Errorf("deleting Connection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted Connection", "name", fqn)
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
		status["serviceGeneratedID"] = old["serviceGeneratedID"]
	}

	u.Object["status"] = status

	return nil
}
