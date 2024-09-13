package bigqueryconnection

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigquery/connection/apiv1"

	bigqueryconnectionpb "cloud.google.com/go/bigquery/connection/apiv1/connectionpb"
	"google.golang.org/api/option"
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
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
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

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := refs.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get location
	location := obj.Spec.Location

	var id *BigQueryConnectionConnectionIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Project != projectID {
			return nil, fmt.Errorf("BigQueryConnectionConnection %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Project, projectID)
		}
		if id.Parent.Location != location {
			return nil, fmt.Errorf("BigQueryConnectionConnection %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Location, location)
		}
		if id.Connection != resourceID {
			return nil, fmt.Errorf("BigQueryConnectionConnection  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Connection, resourceID)
		}
	}

	// Get bigqueryconnection GCP client
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
	return nil, nil
}

type Adapter struct {
	id        *BigQueryConnectionConnectionIdentity
	gcpClient *gcp.Client
	desired   *krm.BigQueryConnectionConnection
	actual    *bigqueryconnectionpb.Connection
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting BigQueryConnectionConnection", "name", a.id.FullyQualifiedName())

	req := &bigqueryconnectionpb.GetConnectionRequest{Name: a.id.FullyQualifiedName()}
	Connectionpb, err := a.gcpClient.GetConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigQueryConnectionConnection %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = Connectionpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating Connection", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigQueryConnectionConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &bigqueryconnectionpb.CreateConnectionRequest{
		Parent:     a.id.Parent.String(),
		Connection: resource,
	}
	created, err := a.gcpClient.CreateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Connection %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created Connection", "name", a.id.FullyQualifiedName())

	status := &krm.BigQueryConnectionConnectionStatus{}
	status.ObservedState = BigQueryConnectionConnectionStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating Connection", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.FriendlyName, a.actual.FriendlyName) {
		updateMask.Paths = append(updateMask.Paths, "friendly_name")
	}
	if !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	desired := a.desired.DeepCopy()
	resource := BigQueryConnectionConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &bigqueryconnectionpb.UpdateConnectionRequest{
		Name:       a.id.FullyQualifiedName(),
		Connection: resource,
		UpdateMask: updateMask,
	}
	updated, err := a.gcpClient.UpdateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Connection %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated Connection", "name", a.id.FullyQualifiedName())

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
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Parent.Project}
	obj.Spec.Location = a.id.Parent.Location
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
	log.V(2).Info("deleting Connection", "name", a.id.FullyQualifiedName())

	req := &bigqueryconnectionpb.DeleteConnectionRequest{Name: a.id.FullyQualifiedName()}
	err := a.gcpClient.DeleteConnection(ctx, req)
	if err != nil {
		//	if !strings.Contains(err.Error(), "Error 404: Not found") {
		return false, fmt.Errorf("deleting Connection %s: %w", a.id.FullyQualifiedName(), err)
		//	}
	}
	log.V(2).Info("successfully deleted Connection", "name", a.id.FullyQualifiedName())
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
	}

	u.Object["status"] = status

	return nil
}
