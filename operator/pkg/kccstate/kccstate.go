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

package kccstate

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
)

// FetchLiveKCCState tries to fetch the ConfigConnector (CC) resource and the ConfigConnectorContext (CCC)
// for the resource's namespace if running in Namespaced mode. It ignores not found errors for CC fetching
// but errors out if KCC is running in Namespaced mode and no CCC is found for the namespace of the resource.
func FetchLiveKCCState(ctx context.Context, c client.Client, resourceNN types.NamespacedName) (v1beta1.ConfigConnector, v1beta1.ConfigConnectorContext, error) {
	var cc v1beta1.ConfigConnector
	if err := c.Get(ctx, types.NamespacedName{
		Name: v1beta1.ConfigConnectorAllowedName,
	}, &cc); err != nil {
		if errors.IsNotFound(err) {
			klog.Infof("%v object is not found", v1beta1.ConfigConnectorAllowedName)
			// if no CC exists, then by definition, KCC cannot be running in namespaced mode;
			return v1beta1.ConfigConnector{}, v1beta1.ConfigConnectorContext{}, nil
		}
		return v1beta1.ConfigConnector{}, v1beta1.ConfigConnectorContext{}, err
	}

	// Namespaced mode is the default mode for the ConfigConnector object.
	if cc.Spec.Mode == "" || cc.Spec.Mode == k8s.NamespacedMode {
		var ccc v1beta1.ConfigConnectorContext
		if err := c.Get(ctx, types.NamespacedName{
			Name:      v1beta1.ConfigConnectorContextAllowedName,
			Namespace: resourceNN.Namespace,
		}, &ccc); err != nil {

			// this should not happen but if we attempt to actuate a resource
			// AND we are running in namespaced mode, not finding a CCC in that namespace
			// is an error in the assumptions that KCC has (i.e. that there is a CCC defined
			// that actively manages resources in that namespace).
			return cc, v1beta1.ConfigConnectorContext{}, err
		}
		return cc, ccc, nil
	}

	return cc, v1beta1.ConfigConnectorContext{}, nil
}
