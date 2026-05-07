package gkehub

import (
	"context"
	"fmt"
	"reflect"
	"time"

	gkehubv1 "google.golang.org/api/gkehub/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.GKEHubScopeRBACRoleBindingGVK, NewGKEHubScopeRBACRoleBindingModel)
}

func NewGKEHubScopeRBACRoleBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubScopeRBACRoleBindingModel{config: config}, nil
}

var _ directbase.Model = &gkeHubScopeRBACRoleBindingModel{}

type gkeHubScopeRBACRoleBindingModel struct {
	config *config.ControllerConfig
}

func (m *gkeHubScopeRBACRoleBindingModel) client(ctx context.Context) (*gkeHubClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}
	serviceV1, err := gkehubv1.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building v1 service for gkehub: %w", err)
	}
	gcpClient.config = *m.config
	return &gkeHubClient{
		namespaceClientV1:            gkehubv1.NewProjectsLocationsScopesNamespacesService(serviceV1),
		scopeRBACRoleBindingClientV1: gkehubv1.NewProjectsLocationsScopesRbacrolebindingsService(serviceV1),
		operationClientV1:            gkehubv1.NewProjectsLocationsOperationsService(serviceV1),
	}, nil
}

func (m *gkeHubScopeRBACRoleBindingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	obj := &krm.GKEHubScopeRBACRoleBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.Object.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}

	client, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &gkeHubScopeRBACRoleBindingAdapter{
		id:        id.(*krm.GKEHubScopeRBACRoleBindingIdentity),
		hubClient: client,
		desired:   obj,
		reader:    op.Reader,
	}, nil
}

func (m *gkeHubScopeRBACRoleBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type gkeHubScopeRBACRoleBindingAdapter struct {
	id        *krm.GKEHubScopeRBACRoleBindingIdentity
	hubClient *gkeHubClient
	desired   *krm.GKEHubScopeRBACRoleBinding
	actual    *gkehubv1.RBACRoleBinding
	reader    client.Reader
}

var _ directbase.Adapter = &gkeHubScopeRBACRoleBindingAdapter{}

func (a *gkeHubScopeRBACRoleBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting GKEHubScopeRBACRoleBinding", "name", a.id.String())

	req := a.hubClient.scopeRBACRoleBindingClientV1.Get(a.id.String())
	rbacRoleBinding, err := req.Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
	}

	a.actual = rbacRoleBinding

	mapCtx := &direct.MapContext{}
	status := GKEHubScopeRBACRoleBindingStatus_FromAPI(mapCtx, rbacRoleBinding)
	if mapCtx.Err() != nil {
		return true, mapCtx.Err()
	}

	a.desired.Status.ObservedState = status.ObservedState
	a.desired.Status.ExternalRef = direct.LazyPtr(a.id.String())
	return true, nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GKEHubScopeRBACRoleBinding", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Normalize references before mapping
	if a.desired.Spec.ScopeRef != nil {
		if err := a.desired.Spec.ScopeRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return err
		}
	}

	desiredAPI := GKEHubScopeRBACRoleBindingSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent().String()
	req := a.hubClient.scopeRBACRoleBindingClientV1.Create(parent, desiredAPI)
	req.RbacrolebindingId(a.id.RBACRoleBindingID)
	op, err := req.Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for GKEHubScopeRBACRoleBinding creation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created GKEHubScopeRBACRoleBinding", "name", a.id.String())

	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("getting GKEHubScopeRBACRoleBinding after creation %q: %w", a.id.String(), err)
	}

	return nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GKEHubScopeRBACRoleBinding", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	// Normalize references before mapping
	if a.desired.Spec.ScopeRef != nil {
		if err := a.desired.Spec.ScopeRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return err
		}
	}

	desiredAPI := GKEHubScopeRBACRoleBindingSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// We use a.actual which was fetched in Find()
	updateMask := []string{}
	actualCopy := *a.actual

	if !reflect.DeepEqual(a.actual.Role, desiredAPI.Role) {
		updateMask = append(updateMask, "role")
		actualCopy.Role = desiredAPI.Role
	}
	if !reflect.DeepEqual(a.actual.Labels, desiredAPI.Labels) {
		updateMask = append(updateMask, "labels")
		actualCopy.Labels = desiredAPI.Labels
	}
	
	if len(updateMask) == 0 {
		log.V(2).Info("no updates required for GKEHubScopeRBACRoleBinding", "name", a.id.String())
		return nil
	}

	patchReq := a.hubClient.scopeRBACRoleBindingClientV1.Patch(a.id.String(), &actualCopy)
	updateMaskStr := ""
	for i, m := range updateMask {
		if i > 0 {
			updateMaskStr += ","
		}
		updateMaskStr += m
	}
	patchReq.UpdateMask(updateMaskStr)

	op, err := patchReq.Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for GKEHubScopeRBACRoleBinding update %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated GKEHubScopeRBACRoleBinding", "name", a.id.String())

	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("getting GKEHubScopeRBACRoleBinding after update %q: %w", a.id.String(), err)
	}

	return nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.desired == nil {
		return nil, fmt.Errorf("GKEHubScopeRBACRoleBindingAdapter.Export: desired object is nil")
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(a.desired)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting GKEHubScopeRBACRoleBinding", "name", a.id.String())

	req := a.hubClient.scopeRBACRoleBindingClientV1.Delete(a.id.String())
	op, err := req.Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for GKEHubScopeRBACRoleBinding deletion %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully deleted GKEHubScopeRBACRoleBinding", "name", a.id.String())
	return true, nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) waitForOp(ctx context.Context, op *gkehubv1.Operation) error {
	retryPeriod := 5 * time.Second
	timeoutDuration := 20 * time.Minute
	timeoutAt := time.Now().Add(timeoutDuration)
	for {
		current, err := a.hubClient.operationClientV1.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q failed: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error.Message)
			}
			return nil
		}
		if time.Now().After(timeoutAt) {
			return fmt.Errorf("timeout waiting for operation %q", op.Name)
		}
		time.Sleep(retryPeriod)
	}
}
