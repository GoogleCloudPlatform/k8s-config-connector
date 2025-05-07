// Copyright 2022 Google LLC
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

package nocache

import (
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// newClientOnlyCacheCCAndCCC is a client builder that caches only CC and CCC objects.
func newClientOnlyCacheCCAndCCC(innerNewClient client.NewClientFunc) client.NewClientFunc {
	return func(config *rest.Config, options client.Options) (client.Client, error) {
		kind := func(gvk schema.GroupVersionKind) *unstructured.Unstructured {
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(gvk)
			return u
		}

		// We cache some objects because we read them often:
		//
		// * CC and CCC: read reconciliation mode, default state-into-spec, etc
		// * CRDs: Build our schema (TODO: should we just list and then restart if these change?)
		//
		// Everything else, we don't want to cache.
		// Unstructured objects don't get cached, but we need an exclude list for some objects we read using typed clients.

		options.Cache.DisableFor = []client.Object{
			kind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Namespace"}),
			kind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Secret"}),
			kind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "ConfigMap"}),
			&iamv1beta1.IAMAuditConfig{},
			&iamv1beta1.IAMPartialPolicy{},
			&iamv1beta1.IAMPolicy{},
			&iamv1beta1.IAMPolicyMember{},
		}

		// Don't cache unstructured objects (this is the default anyway)
		options.Cache.Unstructured = false

		return innerNewClient(config, options)
	}
}

// OnlyCacheCCAndCCC turns off caching for most objects, except for CC and CCC objects.
// We do this so that our memory usage should not grow with the size of objects in the cluster,
// only those we are actively reconciling.
func OnlyCacheCCAndCCC(mgr *manager.Options) {
	innerNewClient := mgr.NewClient
	if mgr.NewClient == nil {
		innerNewClient = client.New
	}
	mgr.NewClient = newClientOnlyCacheCCAndCCC(innerNewClient)
}

var newClientCacheNothing = func(config *rest.Config, options client.Options) (client.Client, error) {
	options.Cache = nil
	return client.New(config, options)
}

// TurnOffAllCaching turns off caching for all objects (including CC and CCC objects).
// We do this so that our memory usage should not grow with the size of objects in the cluster,
// only those we are actively reconciling.
func TurnOffAllCaching(mgr *manager.Options) {
	mgr.NewClient = newClientCacheNothing
}
