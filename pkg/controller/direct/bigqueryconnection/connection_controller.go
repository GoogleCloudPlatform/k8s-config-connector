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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
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

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryConnectionConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve any resource references:
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	connectionIdentity := identity.(*krm.BigQueryConnectionConnectionIdentity)
	// Get bigqueryconnection GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	adapter := &Adapter{
		id:        connectionIdentity,
		gcpClient: gcpClient,
		reader:    reader,
		namespace: obj.Namespace,
	}

	if err := adapter.normalizeReference(ctx, obj); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := BigQueryConnectionConnectionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// aws.accessRole.Identity is a output-only field in CREATE, it is required in the UPDATE.
	desiredProto = processAwsIdentityMerge(desiredProto, obj.Status.ObservedState)

	adapter.desired = desiredProto
	return adapter, nil
}

func (a *Adapter) normalizeReference(ctx context.Context, obj *krm.BigQueryConnectionConnection) error {
	// Resolve SQLInstanceRef and SQLDatabaseRef
	if obj.Spec.CloudSQLSpec != nil {
		sql := obj.Spec.CloudSQLSpec
		if sql.InstanceRef != nil {
			instance, err := refs.ResolveSQLInstanceRef(ctx, a.reader, obj, sql.InstanceRef)
			if err != nil {
				return err
			}
			sql.InstanceRef.External = instance.ConnectionName()
		}
		if sql.DatabaseRef != nil {
			database, err := refs.ResolveSQLDatabaseRef(ctx, a.reader, obj, sql.DatabaseRef)
			if err != nil {
				return err
			}
			sql.DatabaseRef.External = database.Name()
		}
		if sql.Credential != nil {
			if err := refsv1beta1secret.NormalizedSecret(ctx, sql.Credential.SecretRef, a.reader, a.namespace); err != nil {
				return err
			}
		}
	}

	// Resolve SpannerDatabaseRef
	if obj.Spec.CloudSpannerSpec != nil {
		if ref := obj.Spec.CloudSpannerSpec.DatabaseRef; ref != nil {
			_, err := ref.NormalizedExternal(ctx, a.reader, obj.Namespace)
			if err != nil {
				return err
			}
		}
	}

	// Resolve Spark.DataprocClusterRef and Spark.MetastoreServiceRef
	if obj.Spec.SparkSpec != nil {
		if obj.Spec.SparkSpec.SparkHistoryServer != nil {
			if ref := obj.Spec.SparkSpec.SparkHistoryServer.DataprocClusterRef; ref != nil {
				err := ref.Normalize(ctx, a.reader, obj.GetNamespace())
				if err != nil {
					return err
				}
			}
		}

		if obj.Spec.SparkSpec.MetastoreService != nil {
			if ref := obj.Spec.SparkSpec.MetastoreService.MetastoreServiceRef; ref != nil {
				_, err := ref.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.BigQueryConnectionConnectionIdentity
	gcpClient *gcp.Client
	desired   *pb.Connection
	actual    *pb.Connection
	reader    client.Reader
	namespace string
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	if !a.id.HasIdentitySpecified() { // resource is not yet created
		return false, nil
	}
	fqn := a.id.String()
	log.V(2).Info("getting BigQueryConnectionConnection", "name", fqn)

	req := &pb.GetConnectionRequest{Name: fqn}
	connectionpb, err := a.gcpClient.GetConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryConnectionConnection %q: %w", fqn, err)
	}

	a.actual = connectionpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	fqn := a.id.String()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating Connection", "name", fqn)

	parent := a.id.ParentString()
	req := &pb.CreateConnectionRequest{
		Parent:     parent,
		Connection: a.desired,
	}
	if a.id.HasIdentitySpecified() {
		req.ConnectionId = a.id.Connection
	}
	created, err := a.gcpClient.CreateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Connection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully created Connection", "name", created.Name)

	return a.updateStatus(ctx, createOp, created)
}

// aws.accessRole.Identity is a output-only field in CREATE, it is required in the UPDATE.
func processAwsIdentityMerge(desired *pb.Connection, previouslyApplied *krm.BigQueryConnectionConnectionObservedState) *pb.Connection {
	if previouslyApplied != nil && previouslyApplied.Aws != nil {
		desired.GetAws().GetAccessRole().Identity = direct.ValueOf(previouslyApplied.Aws.AccessRole.Identity)
	}
	return desired
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	fqn := a.id.String()
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Connection", "name", fqn)

	a.desired.Name = a.actual.Name

	diffs, updateMask, err := compareConnection(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}
	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", fqn)
		return nil
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateConnectionRequest{
		Name:       fqn,
		Connection: a.desired,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Connection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully updated Connection", "name", fqn)

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Connection) error {
	mapCtx := &direct.MapContext{}
	observedState := BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status := &krm.BigQueryConnectionConnectionStatus{}
	status.ObservedState = observedState

	parent := a.id.ParentString()
	tokens := strings.Split(latest.Name, "/")
	externalRef := parent + "/connections/" + tokens[len(tokens)-1]
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func compareConnection(ctx context.Context, actual, desired *pb.Connection) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, BigQueryConnectionConnectionSpec_FromProto, BigQueryConnectionConnectionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.Connection)

	populateDefaults := func(obj *pb.Connection) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryConnectionConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryConnectionConnectionSpec_FromProto(mapCtx, a.actual))
	tokens := strings.Split(a.id.String(), "connections/")
	obj.Spec.ResourceID = direct.LazyPtr(tokens[len(tokens)-1])
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent := a.id.ParentString()
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
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("deleting Connection", "name", fqn)

	req := &pb.DeleteConnectionRequest{Name: fqn}
	if err := a.gcpClient.DeleteConnection(ctx, req); err != nil {
		return false, fmt.Errorf("deleting Connection %s: %w", fqn, err)
	}
	log.V(2).Info("successfully deleted Connection", "name", fqn)
	return true, nil
}
