// Copyright 2026 Google LLC
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

package computerefs

import (
	"context"
	"errors"
	"fmt"
	"sync"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ComputeAddressResolver struct {
	config *config.ControllerConfig

	mu                    sync.Mutex
	globalAddressesClient *compute.GlobalAddressesClient
	addressesClient       *compute.AddressesClient
}

func NewComputeAddressResolver(config *config.ControllerConfig) *ComputeAddressResolver {
	return &ComputeAddressResolver{
		config: config,
	}
}

func (r *ComputeAddressResolver) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var errs []error
	if r.globalAddressesClient != nil {
		if err := r.globalAddressesClient.Close(); err != nil {
			errs = append(errs, err)
		} else {
			r.globalAddressesClient = nil
		}
	}
	if r.addressesClient != nil {
		if err := r.addressesClient.Close(); err != nil {
			errs = append(errs, err)
		} else {
			r.addressesClient = nil
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("closing resolver clients: %w", errors.Join(errs...))
	}
	return nil
}

func (r *ComputeAddressResolver) getGlobalAddressesClient(ctx context.Context) (*compute.GlobalAddressesClient, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.globalAddressesClient != nil {
		return r.globalAddressesClient, nil
	}

	opts, err := r.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	client, err := compute.NewGlobalAddressesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building GlobalAddresses client: %w", err)
	}
	r.globalAddressesClient = client
	return client, nil
}

func (r *ComputeAddressResolver) getAddressesClient(ctx context.Context) (*compute.AddressesClient, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.addressesClient != nil {
		return r.addressesClient, nil
	}

	opts, err := r.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	client, err := compute.NewAddressesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Addresses client: %w", err)
	}
	r.addressesClient = client
	return client, nil
}

func (r *ComputeAddressResolver) ResolveComputeAddressIP(ctx context.Context, reader client.Reader, src client.Object, ref *computev1beta1.ComputeAddressRef) (string, error) {
	if ref == nil {
		return "", nil
	}

	// Case 1: Name reference (internal resource managed by KCC)
	if ref.Name != "" {
		if ref.External != "" {
			return "", fmt.Errorf("cannot specify both name and external on reference")
		}

		key := types.NamespacedName{
			Namespace: ref.Namespace,
			Name:      ref.Name,
		}
		if key.Namespace == "" {
			key.Namespace = src.GetNamespace()
		}

		gvk := schema.GroupVersionKind{
			Group:   "compute.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "ComputeAddress",
		}

		computeAddress := &unstructured.Unstructured{}
		computeAddress.SetGroupVersionKind(gvk)
		if err := reader.Get(ctx, key, computeAddress); err != nil {
			if apierrors.IsNotFound(err) {
				return "", k8s.NewReferenceNotFoundError(gvk, key)
			}
			return "", fmt.Errorf("error reading referenced %v %v: %w", gvk.Kind, key, err)
		}

		// Because `spec.address` field is optional, we can't guarantee it always
		// exists in a successfully reconciled ComputeAddress CR, so we should use
		// the `status.address` or `status.observedState.address` instead.
		// We also want to wait for the referenced ComputeAddress to be ready - that's the main reason to get it from status.address or status.observedState.address.
		address, _, err := unstructured.NestedString(computeAddress.Object, "status", "address")
		if err != nil || address == "" {
			address, _, err = unstructured.NestedString(computeAddress.Object, "status", "observedState", "address")
			if err != nil || address == "" {
				return "", k8s.NewReferenceNotReadyError(gvk, key)
			}
		}
		return address, nil
	}

	// Case 2: External reference (external GCP resource or raw URI)
	if ref.External != "" {
		id, err := computev1beta1.ParseComputeAddressExternal(ref.External)
		if err != nil {
			return "", err
		}

		if id.IsGlobal() {
			client, err := r.getGlobalAddressesClient(ctx)
			if err != nil {
				return "", err
			}

			req := &computepb.GetGlobalAddressRequest{
				Project: id.Project,
				Address: id.Address,
			}
			addr, err := client.Get(ctx, req)
			if err != nil {
				return "", fmt.Errorf("getting global address %s/%s from GCP: %w", id.Project, id.Address, err)
			}
			ip := addr.GetAddress()
			if ip == "" {
				return "", fmt.Errorf("global address %s/%s address field is empty in response", id.Project, id.Address)
			}
			return ip, nil
		} else {
			client, err := r.getAddressesClient(ctx)
			if err != nil {
				return "", err
			}

			req := &computepb.GetAddressRequest{
				Project: id.Project,
				Region:  id.Region,
				Address: id.Address,
			}
			addr, err := client.Get(ctx, req)
			if err != nil {
				return "", fmt.Errorf("getting regional address %s/%s/%s from GCP: %w", id.Project, id.Region, id.Address, err)
			}
			ip := addr.GetAddress()
			if ip == "" {
				return "", fmt.Errorf("regional address %s/%s/%s address field is empty in response", id.Project, id.Region, id.Address)
			}
			return ip, nil
		}
	}

	return "", fmt.Errorf("must specify either name or external on reference")
}
