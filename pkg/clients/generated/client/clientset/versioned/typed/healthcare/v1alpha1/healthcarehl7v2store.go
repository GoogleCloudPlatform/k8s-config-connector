// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/healthcare/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HealthcareHL7V2StoresGetter has a method to return a HealthcareHL7V2StoreInterface.
// A group's client should implement this interface.
type HealthcareHL7V2StoresGetter interface {
	HealthcareHL7V2Stores(namespace string) HealthcareHL7V2StoreInterface
}

// HealthcareHL7V2StoreInterface has methods to work with HealthcareHL7V2Store resources.
type HealthcareHL7V2StoreInterface interface {
	Create(ctx context.Context, healthcareHL7V2Store *v1alpha1.HealthcareHL7V2Store, opts v1.CreateOptions) (*v1alpha1.HealthcareHL7V2Store, error)
	Update(ctx context.Context, healthcareHL7V2Store *v1alpha1.HealthcareHL7V2Store, opts v1.UpdateOptions) (*v1alpha1.HealthcareHL7V2Store, error)
	UpdateStatus(ctx context.Context, healthcareHL7V2Store *v1alpha1.HealthcareHL7V2Store, opts v1.UpdateOptions) (*v1alpha1.HealthcareHL7V2Store, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HealthcareHL7V2Store, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HealthcareHL7V2StoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HealthcareHL7V2Store, err error)
	HealthcareHL7V2StoreExpansion
}

// healthcareHL7V2Stores implements HealthcareHL7V2StoreInterface
type healthcareHL7V2Stores struct {
	client rest.Interface
	ns     string
}

// newHealthcareHL7V2Stores returns a HealthcareHL7V2Stores
func newHealthcareHL7V2Stores(c *HealthcareV1alpha1Client, namespace string) *healthcareHL7V2Stores {
	return &healthcareHL7V2Stores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the healthcareHL7V2Store, and returns the corresponding healthcareHL7V2Store object, and an error if there is any.
func (c *healthcareHL7V2Stores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HealthcareHL7V2Store, err error) {
	result = &v1alpha1.HealthcareHL7V2Store{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HealthcareHL7V2Stores that match those selectors.
func (c *healthcareHL7V2Stores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HealthcareHL7V2StoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HealthcareHL7V2StoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested healthcareHL7V2Stores.
func (c *healthcareHL7V2Stores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a healthcareHL7V2Store and creates it.  Returns the server's representation of the healthcareHL7V2Store, and an error, if there is any.
func (c *healthcareHL7V2Stores) Create(ctx context.Context, healthcareHL7V2Store *v1alpha1.HealthcareHL7V2Store, opts v1.CreateOptions) (result *v1alpha1.HealthcareHL7V2Store, err error) {
	result = &v1alpha1.HealthcareHL7V2Store{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(healthcareHL7V2Store).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a healthcareHL7V2Store and updates it. Returns the server's representation of the healthcareHL7V2Store, and an error, if there is any.
func (c *healthcareHL7V2Stores) Update(ctx context.Context, healthcareHL7V2Store *v1alpha1.HealthcareHL7V2Store, opts v1.UpdateOptions) (result *v1alpha1.HealthcareHL7V2Store, err error) {
	result = &v1alpha1.HealthcareHL7V2Store{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		Name(healthcareHL7V2Store.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(healthcareHL7V2Store).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *healthcareHL7V2Stores) UpdateStatus(ctx context.Context, healthcareHL7V2Store *v1alpha1.HealthcareHL7V2Store, opts v1.UpdateOptions) (result *v1alpha1.HealthcareHL7V2Store, err error) {
	result = &v1alpha1.HealthcareHL7V2Store{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		Name(healthcareHL7V2Store.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(healthcareHL7V2Store).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the healthcareHL7V2Store and deletes it. Returns an error if one occurs.
func (c *healthcareHL7V2Stores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *healthcareHL7V2Stores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched healthcareHL7V2Store.
func (c *healthcareHL7V2Stores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HealthcareHL7V2Store, err error) {
	result = &v1alpha1.HealthcareHL7V2Store{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("healthcarehl7v2stores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
