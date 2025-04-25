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

package kccconfig

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
)

type Configuration struct {
	cc  *v1beta1.ConfigConnector
	ccc *v1beta1.ConfigConnectorContext
}

// FetchLiveKCCState tries to fetch the ConfigConnector (CC) resource and the ConfigConnectorContext (CCC)
// for the resource's namespace if running in Namespaced mode. It ignores not found errors for CC fetching
// but errors out if KCC is running in Namespaced mode and no CCC is found for the namespace of the resource.
func FetchLiveKCCState(ctx context.Context, c client.Client, resourceNN types.NamespacedName) (*Configuration, error) {
	config := &Configuration{}

	cc := &v1beta1.ConfigConnector{}
	if err := c.Get(ctx, types.NamespacedName{
		Name: operatork8s.ConfigConnectorAllowedName,
	}, cc); err != nil {
		if errors.IsNotFound(err) {
			klog.Infof("%v object is not found", operatork8s.ConfigConnectorAllowedName)
			// if no CC exists, then by definition, KCC cannot be running in namespaced mode;
			return config, nil
		}
		return nil, err
	} else {
		config.cc = cc
	}

	// Namespaced mode is the default mode for the ConfigConnector object.
	if cc.Spec.Mode == "" || cc.Spec.Mode == v1beta1.NamespacedMode {
		ccc := &v1beta1.ConfigConnectorContext{}
		if err := c.Get(ctx, types.NamespacedName{
			Name:      operatork8s.ConfigConnectorContextAllowedName,
			Namespace: resourceNN.Namespace,
		}, ccc); err != nil {

			// this should not happen but if we attempt to actuate a resource
			// AND we are running in namespaced mode, not finding a CCC in that namespace
			// is an error in the assumptions that KCC has (i.e. that there is a CCC defined
			// that actively manages resources in that namespace).
			return config, err
		} else {
			config.ccc = ccc
		}
		return config, nil
	}

	return config, nil
}

func (c *Configuration) DefaultStateIntoSpecAnnotation() (string, error) {
	annotationValue := apis.StateIntoSpecDefaultValueV1Beta1

	if c.cc != nil && c.cc.Spec.StateIntoSpec != nil {
		switch *c.cc.Spec.StateIntoSpec {
		case v1beta1.StateIntoSpecMerge:
			annotationValue = apis.StateMergeIntoSpec
		case v1beta1.StateIntoSpecAbsent:
			annotationValue = apis.StateAbsentInSpec

		default:
			return "", fmt.Errorf("invalid value %q for spec.stateIntoSpec in ConfigConnector, should be Absent or Merge (Absent recommended)", *c.cc.Spec.StateIntoSpec)
		}
	}

	if c.ccc != nil && c.ccc.Spec.StateIntoSpec != nil {
		switch *c.ccc.Spec.StateIntoSpec {
		case v1beta1.StateIntoSpecMerge:
			annotationValue = apis.StateMergeIntoSpec
		case v1beta1.StateIntoSpecAbsent:
			annotationValue = apis.StateAbsentInSpec

		default:
			return "", fmt.Errorf("invalid value %q for spec.stateIntoSpec in ConfigConnectorContext, should be Absent or Merge (Absent recommended)", *c.ccc.Spec.StateIntoSpec)
		}
	}
	return annotationValue, nil
}

// ActuationMode maps to the API actuation mode, but avoids a dependency on the API itself
type ActuationMode string

const (
	ActuationModeReconciling ActuationMode = ActuationMode(v1beta1.Reconciling)
	ActuationModePaused      ActuationMode = ActuationMode(v1beta1.Paused)
)

// ActuationMode looks at CC and CCC to see if they specify an actuationMode.
// - If both CC & CCC specify a actuationMode in Namespaced mode, we defer to the CCC's value.
//   - If only CC specifies a actuationMode in Namespaced mode, we defer to the CC's value.
//
// - If both CC & CCC specify an actuationMode in cluster mode, the CCC specification is irrelevant.
// - If neither CC nor CCC specify a actuationMode, we defer to the default value defined in apis.
func (c *Configuration) ActuationMode() ActuationMode {
	if c.cc != nil && c.cc.Spec.Mode == v1beta1.NamespacedMode {
		if c.ccc != nil && c.ccc.Spec.Actuation != "" {
			return ActuationMode(c.ccc.Spec.Actuation)
		}
	}

	// if no CCC exists or doesn't define a value, defer to the CC's value.
	if c.cc != nil && c.cc.Spec.Actuation != "" {
		return ActuationMode(c.cc.Spec.Actuation)
	}

	return ActuationMode(v1beta1.DefaultActuationMode())
}
