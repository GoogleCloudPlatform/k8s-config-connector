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
// proto.service: google.cloud.networkmanagement.v1.ReachabilityService
// proto.message: google.cloud.networkmanagement.v1.ConnectivityTest
// crd.type: NetworkManagementConnectivityTest
// crd.version: v1alpha1

package networkmanagement

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/networkmanagement/apiv1"
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.NetworkManagementConnectivityTestGVK, NewConnectivityTestModel)
}

func NewConnectivityTestModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &connectivityTestModel{config: config}, nil
}

var _ directbase.Model = &connectivityTestModel{}

type connectivityTestModel struct {
	config *config.ControllerConfig
}

func (m *connectivityTestModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkManagementConnectivityTest{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewConnectivityTestIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newReachabilityClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	return &connectivityTestAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *connectivityTestModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type connectivityTestAdapter struct {
	gcpClient *gcp.ReachabilityClient
	id        *krm.ConnectivityTestIdentity
	desired   *krm.NetworkManagementConnectivityTest
	actual    *pb.ConnectivityTest
	reader    client.Reader
}

var _ directbase.Adapter = &connectivityTestAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *connectivityTestAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networkmanagement connectivitytest", "name", a.id)

	req := &pb.GetConnectivityTestRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetConnectivityTest(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networkmanagement connectivitytest %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *connectivityTestAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating networkmanagement connectivitytest", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := NetworkManagementConnectivityTestSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateConnectivityTestRequest{
		Parent:   a.id.Parent().String(),
		TestId:   a.id.ID(),
		Resource: resource,
	}
	op, err := a.gcpClient.CreateConnectivityTest(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networkmanagement connectivitytest %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("networkmanagement connectivitytest %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created networkmanagement connectivitytest in gcp", "name", a.id)

	status := &krm.NetworkManagementConnectivityTestStatus{}
	status.ObservedState = NetworkManagementConnectivityTestObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *connectivityTestAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networkmanagement connectivitytest", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := NetworkManagementConnectivityTestSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String() // Set the name for the update request

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("calculating diff for connectivity test %s: %w", a.id, err)
	}

	topLevelFieldPaths := sets.New[string]()
	for path, _ := range paths {
		tokens := strings.Split(path, ".")
		topLevelFieldPaths.Insert(tokens[0])
	}
	// Remove output-only fields.
	topLevelFieldPaths.Delete("reachability_details")

	var updated *pb.ConnectivityTest
	if len(topLevelFieldPaths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Still update the status to cover the use case of acquisition.
		updated = a.actual
	} else {
		updateMask := &fieldmaskpb.FieldMask{Paths: sets.List(topLevelFieldPaths)}
		log.V(2).Info("updating fields", "name", a.id, "paths", updateMask.Paths)

		req := &pb.UpdateConnectivityTestRequest{
			UpdateMask: updateMask,
			Resource:   resource,
		}
		op, err := a.gcpClient.UpdateConnectivityTest(ctx, req)
		if err != nil {
			return fmt.Errorf("updating networkmanagement connectivitytest %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("networkmanagement connectivitytest %s waiting for update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated networkmanagement connectivitytest", "name", a.id)

		updated = a.actual // Use the fetched resource for status update
	}

	status := &krm.NetworkManagementConnectivityTestStatus{}
	status.ObservedState = NetworkManagementConnectivityTestObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// Set externalRef here to cover the use case of acquisition.
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *connectivityTestAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkManagementConnectivityTest{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkManagementConnectivityTestSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID()) // Use connectivity test id from identity
	u.SetGroupVersionKind(krm.NetworkManagementConnectivityTestGVK)
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *connectivityTestAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networkmanagement connectivitytest", "name", a.id)

	req := &pb.DeleteConnectivityTestRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteConnectivityTest(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent networkmanagement connectivitytest, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting networkmanagement connectivitytest %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully initiated deletion of networkmanagement connectivitytest", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for deletion of networkmanagement connectivitytest %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted networkmanagement connectivitytest", "name", a.id)
	return true, nil
}

func (a *connectivityTestAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired

	if obj.Spec.Source != nil {
		if obj.Spec.Source.ComputeInstanceRef != nil {
			if _, err := obj.Spec.Source.ComputeInstanceRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		if obj.Spec.Source.ComputeNetworkRef != nil {
			external, err := obj.Spec.Source.ComputeNetworkRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
			if err != nil {
				return err
			}
			obj.Spec.Source.ComputeNetworkRef.External = external
		}
		if obj.Spec.Source.ContainerClusterRef != nil {
			if _, err := obj.Spec.Source.ContainerClusterRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		if obj.Spec.Source.SQLInstanceRef != nil {
			instance, err := refs.ResolveSQLInstanceRef(ctx, a.reader, obj, obj.Spec.Source.SQLInstanceRef)
			if err != nil {
				return err
			}
			obj.Spec.Source.SQLInstanceRef.External = instance.String()
		}
		if obj.Spec.Source.ProjectRef != nil {
			projectRef, err := refs.ResolveProject(ctx, a.reader, obj.GetNamespace(), obj.Spec.Source.ProjectRef)
			if err != nil {
				return err
			}
			obj.Spec.Source.ProjectRef.External = projectRef.ProjectID
		}
		if obj.Spec.Source.CloudRunRevision != nil && obj.Spec.Source.CloudRunRevision.RunRevisionRef != nil {
			if _, err := obj.Spec.Source.CloudRunRevision.RunRevisionRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
	}

	if obj.Spec.Destination != nil {
		if obj.Spec.Destination.ComputeInstanceRef != nil {
			if _, err := obj.Spec.Destination.ComputeInstanceRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		if obj.Spec.Destination.ComputeNetworkRef != nil {
			external, err := obj.Spec.Destination.ComputeNetworkRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
			if err != nil {
				return err
			}
			obj.Spec.Destination.ComputeNetworkRef.External = external
		}
		if obj.Spec.Destination.ContainerClusterRef != nil {
			if _, err := obj.Spec.Destination.ContainerClusterRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
		if obj.Spec.Destination.SQLInstanceRef != nil {
			instance, err := refs.ResolveSQLInstanceRef(ctx, a.reader, obj, obj.Spec.Destination.SQLInstanceRef)
			if err != nil {
				return err
			}
			obj.Spec.Destination.SQLInstanceRef.External = instance.String()
		}
		if obj.Spec.Destination.ProjectRef != nil {
			projectRef, err := refs.ResolveProject(ctx, a.reader, obj.GetNamespace(), obj.Spec.Destination.ProjectRef)
			if err != nil {
				return err
			}
			obj.Spec.Destination.ProjectRef.External = projectRef.ProjectID
		}
		if obj.Spec.Destination.CloudRunRevision != nil && obj.Spec.Destination.CloudRunRevision.RunRevisionRef != nil {
			if _, err := obj.Spec.Destination.CloudRunRevision.RunRevisionRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
				return err
			}
		}
	}
	return nil
}
