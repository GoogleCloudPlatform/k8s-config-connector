// Copyright 2024 Google LLC. All Rights Reserved.
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
package beta

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type DeliveryPipeline struct {
	Name           *string                         `json:"name"`
	Uid            *string                         `json:"uid"`
	Description    *string                         `json:"description"`
	Annotations    map[string]string               `json:"annotations"`
	Labels         map[string]string               `json:"labels"`
	CreateTime     *string                         `json:"createTime"`
	UpdateTime     *string                         `json:"updateTime"`
	SerialPipeline *DeliveryPipelineSerialPipeline `json:"serialPipeline"`
	Condition      *DeliveryPipelineCondition      `json:"condition"`
	Etag           *string                         `json:"etag"`
	Project        *string                         `json:"project"`
	Location       *string                         `json:"location"`
	Suspended      *bool                           `json:"suspended"`
}

func (r *DeliveryPipeline) String() string {
	return dcl.SprintResource(r)
}

type DeliveryPipelineSerialPipeline struct {
	empty  bool                                   `json:"-"`
	Stages []DeliveryPipelineSerialPipelineStages `json:"stages"`
}

type jsonDeliveryPipelineSerialPipeline DeliveryPipelineSerialPipeline

func (r *DeliveryPipelineSerialPipeline) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipeline
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipeline
	} else {

		r.Stages = res.Stages

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipeline is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipeline *DeliveryPipelineSerialPipeline = &DeliveryPipelineSerialPipeline{empty: true}

func (r *DeliveryPipelineSerialPipeline) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipeline) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipeline) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStages struct {
	empty            bool                                                   `json:"-"`
	TargetId         *string                                                `json:"targetId"`
	Profiles         []string                                               `json:"profiles"`
	Strategy         *DeliveryPipelineSerialPipelineStagesStrategy          `json:"strategy"`
	DeployParameters []DeliveryPipelineSerialPipelineStagesDeployParameters `json:"deployParameters"`
}

type jsonDeliveryPipelineSerialPipelineStages DeliveryPipelineSerialPipelineStages

func (r *DeliveryPipelineSerialPipelineStages) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStages
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStages
	} else {

		r.TargetId = res.TargetId

		r.Profiles = res.Profiles

		r.Strategy = res.Strategy

		r.DeployParameters = res.DeployParameters

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStages is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStages *DeliveryPipelineSerialPipelineStages = &DeliveryPipelineSerialPipelineStages{empty: true}

func (r *DeliveryPipelineSerialPipelineStages) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStages) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStages) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategy struct {
	empty    bool                                                  `json:"-"`
	Standard *DeliveryPipelineSerialPipelineStagesStrategyStandard `json:"standard"`
	Canary   *DeliveryPipelineSerialPipelineStagesStrategyCanary   `json:"canary"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategy DeliveryPipelineSerialPipelineStagesStrategy

func (r *DeliveryPipelineSerialPipelineStagesStrategy) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategy
	} else {

		r.Standard = res.Standard

		r.Canary = res.Canary

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategy *DeliveryPipelineSerialPipelineStagesStrategy = &DeliveryPipelineSerialPipelineStagesStrategy{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategy) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategy) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyStandard struct {
	empty      bool                                                            `json:"-"`
	Verify     *bool                                                           `json:"verify"`
	Predeploy  *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy  `json:"predeploy"`
	Postdeploy *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy `json:"postdeploy"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyStandard DeliveryPipelineSerialPipelineStagesStrategyStandard

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandard) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyStandard
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyStandard
	} else {

		r.Verify = res.Verify

		r.Predeploy = res.Predeploy

		r.Postdeploy = res.Postdeploy

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyStandard is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyStandard *DeliveryPipelineSerialPipelineStagesStrategyStandard = &DeliveryPipelineSerialPipelineStagesStrategyStandard{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandard) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandard) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandard) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy struct {
	empty   bool     `json:"-"`
	Actions []string `json:"actions"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy
	} else {

		r.Actions = res.Actions

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy struct {
	empty   bool     `json:"-"`
	Actions []string `json:"actions"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy
	} else {

		r.Actions = res.Actions

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanary struct {
	empty                  bool                                                                      `json:"-"`
	RuntimeConfig          *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig          `json:"runtimeConfig"`
	CanaryDeployment       *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment       `json:"canaryDeployment"`
	CustomCanaryDeployment *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment `json:"customCanaryDeployment"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanary DeliveryPipelineSerialPipelineStagesStrategyCanary

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanary) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanary
	} else {

		r.RuntimeConfig = res.RuntimeConfig

		r.CanaryDeployment = res.CanaryDeployment

		r.CustomCanaryDeployment = res.CustomCanaryDeployment

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanary *DeliveryPipelineSerialPipelineStagesStrategyCanary = &DeliveryPipelineSerialPipelineStagesStrategyCanary{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanary) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanary) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig struct {
	empty      bool                                                                       `json:"-"`
	Kubernetes *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes `json:"kubernetes"`
	CloudRun   *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun   `json:"cloudRun"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig
	} else {

		r.Kubernetes = res.Kubernetes

		r.CloudRun = res.CloudRun

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes struct {
	empty              bool                                                                                         `json:"-"`
	GatewayServiceMesh *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh `json:"gatewayServiceMesh"`
	ServiceNetworking  *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking  `json:"serviceNetworking"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes
	} else {

		r.GatewayServiceMesh = res.GatewayServiceMesh

		r.ServiceNetworking = res.ServiceNetworking

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh struct {
	empty                 bool    `json:"-"`
	HttpRoute             *string `json:"httpRoute"`
	Service               *string `json:"service"`
	Deployment            *string `json:"deployment"`
	RouteUpdateWaitTime   *string `json:"routeUpdateWaitTime"`
	StableCutbackDuration *string `json:"stableCutbackDuration"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh
	} else {

		r.HttpRoute = res.HttpRoute

		r.Service = res.Service

		r.Deployment = res.Deployment

		r.RouteUpdateWaitTime = res.RouteUpdateWaitTime

		r.StableCutbackDuration = res.StableCutbackDuration

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking struct {
	empty                      bool    `json:"-"`
	Service                    *string `json:"service"`
	Deployment                 *string `json:"deployment"`
	DisablePodOverprovisioning *bool   `json:"disablePodOverprovisioning"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking
	} else {

		r.Service = res.Service

		r.Deployment = res.Deployment

		r.DisablePodOverprovisioning = res.DisablePodOverprovisioning

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun struct {
	empty                   bool     `json:"-"`
	AutomaticTrafficControl *bool    `json:"automaticTrafficControl"`
	CanaryRevisionTags      []string `json:"canaryRevisionTags"`
	PriorRevisionTags       []string `json:"priorRevisionTags"`
	StableRevisionTags      []string `json:"stableRevisionTags"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun
	} else {

		r.AutomaticTrafficControl = res.AutomaticTrafficControl

		r.CanaryRevisionTags = res.CanaryRevisionTags

		r.PriorRevisionTags = res.PriorRevisionTags

		r.StableRevisionTags = res.StableRevisionTags

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun = &DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment struct {
	empty       bool                                                                          `json:"-"`
	Percentages []int64                                                                       `json:"percentages"`
	Verify      *bool                                                                         `json:"verify"`
	Predeploy   *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy  `json:"predeploy"`
	Postdeploy  *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy `json:"postdeploy"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	} else {

		r.Percentages = res.Percentages

		r.Verify = res.Verify

		r.Predeploy = res.Predeploy

		r.Postdeploy = res.Postdeploy

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy struct {
	empty   bool     `json:"-"`
	Actions []string `json:"actions"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy
	} else {

		r.Actions = res.Actions

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy struct {
	empty   bool     `json:"-"`
	Actions []string `json:"actions"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy
	} else {

		r.Actions = res.Actions

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment struct {
	empty        bool                                                                                   `json:"-"`
	PhaseConfigs []DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs `json:"phaseConfigs"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment
	} else {

		r.PhaseConfigs = res.PhaseConfigs

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs struct {
	empty      bool                                                                                            `json:"-"`
	PhaseId    *string                                                                                         `json:"phaseId"`
	Percentage *int64                                                                                          `json:"percentage"`
	Profiles   []string                                                                                        `json:"profiles"`
	Verify     *bool                                                                                           `json:"verify"`
	Predeploy  *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy  `json:"predeploy"`
	Postdeploy *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy `json:"postdeploy"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs
	} else {

		r.PhaseId = res.PhaseId

		r.Percentage = res.Percentage

		r.Profiles = res.Profiles

		r.Verify = res.Verify

		r.Predeploy = res.Predeploy

		r.Postdeploy = res.Postdeploy

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigs) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy struct {
	empty   bool     `json:"-"`
	Actions []string `json:"actions"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy
	} else {

		r.Actions = res.Actions

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy struct {
	empty   bool     `json:"-"`
	Actions []string `json:"actions"`
}

type jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy
	} else {

		r.Actions = res.Actions

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy = &DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPostdeploy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineSerialPipelineStagesDeployParameters struct {
	empty             bool              `json:"-"`
	Values            map[string]string `json:"values"`
	MatchTargetLabels map[string]string `json:"matchTargetLabels"`
}

type jsonDeliveryPipelineSerialPipelineStagesDeployParameters DeliveryPipelineSerialPipelineStagesDeployParameters

func (r *DeliveryPipelineSerialPipelineStagesDeployParameters) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineSerialPipelineStagesDeployParameters
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineSerialPipelineStagesDeployParameters
	} else {

		r.Values = res.Values

		r.MatchTargetLabels = res.MatchTargetLabels

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineSerialPipelineStagesDeployParameters is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineSerialPipelineStagesDeployParameters *DeliveryPipelineSerialPipelineStagesDeployParameters = &DeliveryPipelineSerialPipelineStagesDeployParameters{empty: true}

func (r *DeliveryPipelineSerialPipelineStagesDeployParameters) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineSerialPipelineStagesDeployParameters) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineSerialPipelineStagesDeployParameters) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineCondition struct {
	empty                   bool                                              `json:"-"`
	PipelineReadyCondition  *DeliveryPipelineConditionPipelineReadyCondition  `json:"pipelineReadyCondition"`
	TargetsPresentCondition *DeliveryPipelineConditionTargetsPresentCondition `json:"targetsPresentCondition"`
	TargetsTypeCondition    *DeliveryPipelineConditionTargetsTypeCondition    `json:"targetsTypeCondition"`
}

type jsonDeliveryPipelineCondition DeliveryPipelineCondition

func (r *DeliveryPipelineCondition) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineCondition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineCondition
	} else {

		r.PipelineReadyCondition = res.PipelineReadyCondition

		r.TargetsPresentCondition = res.TargetsPresentCondition

		r.TargetsTypeCondition = res.TargetsTypeCondition

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineCondition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineCondition *DeliveryPipelineCondition = &DeliveryPipelineCondition{empty: true}

func (r *DeliveryPipelineCondition) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineConditionPipelineReadyCondition struct {
	empty      bool    `json:"-"`
	Status     *bool   `json:"status"`
	UpdateTime *string `json:"updateTime"`
}

type jsonDeliveryPipelineConditionPipelineReadyCondition DeliveryPipelineConditionPipelineReadyCondition

func (r *DeliveryPipelineConditionPipelineReadyCondition) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineConditionPipelineReadyCondition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineConditionPipelineReadyCondition
	} else {

		r.Status = res.Status

		r.UpdateTime = res.UpdateTime

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineConditionPipelineReadyCondition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineConditionPipelineReadyCondition *DeliveryPipelineConditionPipelineReadyCondition = &DeliveryPipelineConditionPipelineReadyCondition{empty: true}

func (r *DeliveryPipelineConditionPipelineReadyCondition) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineConditionPipelineReadyCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineConditionPipelineReadyCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineConditionTargetsPresentCondition struct {
	empty          bool     `json:"-"`
	Status         *bool    `json:"status"`
	MissingTargets []string `json:"missingTargets"`
	UpdateTime     *string  `json:"updateTime"`
}

type jsonDeliveryPipelineConditionTargetsPresentCondition DeliveryPipelineConditionTargetsPresentCondition

func (r *DeliveryPipelineConditionTargetsPresentCondition) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineConditionTargetsPresentCondition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineConditionTargetsPresentCondition
	} else {

		r.Status = res.Status

		r.MissingTargets = res.MissingTargets

		r.UpdateTime = res.UpdateTime

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineConditionTargetsPresentCondition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineConditionTargetsPresentCondition *DeliveryPipelineConditionTargetsPresentCondition = &DeliveryPipelineConditionTargetsPresentCondition{empty: true}

func (r *DeliveryPipelineConditionTargetsPresentCondition) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineConditionTargetsPresentCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineConditionTargetsPresentCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DeliveryPipelineConditionTargetsTypeCondition struct {
	empty        bool    `json:"-"`
	Status       *bool   `json:"status"`
	ErrorDetails *string `json:"errorDetails"`
}

type jsonDeliveryPipelineConditionTargetsTypeCondition DeliveryPipelineConditionTargetsTypeCondition

func (r *DeliveryPipelineConditionTargetsTypeCondition) UnmarshalJSON(data []byte) error {
	var res jsonDeliveryPipelineConditionTargetsTypeCondition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDeliveryPipelineConditionTargetsTypeCondition
	} else {

		r.Status = res.Status

		r.ErrorDetails = res.ErrorDetails

	}
	return nil
}

// This object is used to assert a desired state where this DeliveryPipelineConditionTargetsTypeCondition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDeliveryPipelineConditionTargetsTypeCondition *DeliveryPipelineConditionTargetsTypeCondition = &DeliveryPipelineConditionTargetsTypeCondition{empty: true}

func (r *DeliveryPipelineConditionTargetsTypeCondition) Empty() bool {
	return r.empty
}

func (r *DeliveryPipelineConditionTargetsTypeCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *DeliveryPipelineConditionTargetsTypeCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *DeliveryPipeline) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "clouddeploy",
		Type:    "DeliveryPipeline",
		Version: "beta",
	}
}

func (r *DeliveryPipeline) ID() (string, error) {
	if err := extractDeliveryPipelineFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":            dcl.ValueOrEmptyString(nr.Name),
		"uid":             dcl.ValueOrEmptyString(nr.Uid),
		"description":     dcl.ValueOrEmptyString(nr.Description),
		"annotations":     dcl.ValueOrEmptyString(nr.Annotations),
		"labels":          dcl.ValueOrEmptyString(nr.Labels),
		"create_time":     dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":     dcl.ValueOrEmptyString(nr.UpdateTime),
		"serial_pipeline": dcl.ValueOrEmptyString(nr.SerialPipeline),
		"condition":       dcl.ValueOrEmptyString(nr.Condition),
		"etag":            dcl.ValueOrEmptyString(nr.Etag),
		"project":         dcl.ValueOrEmptyString(nr.Project),
		"location":        dcl.ValueOrEmptyString(nr.Location),
		"suspended":       dcl.ValueOrEmptyString(nr.Suspended),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}", params), nil
}

const DeliveryPipelineMaxPage = -1

type DeliveryPipelineList struct {
	Items []*DeliveryPipeline

	nextToken string

	pageSize int32

	resource *DeliveryPipeline
}

func (l *DeliveryPipelineList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DeliveryPipelineList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDeliveryPipeline(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDeliveryPipeline(ctx context.Context, project, location string) (*DeliveryPipelineList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDeliveryPipelineWithMaxResults(ctx, project, location, DeliveryPipelineMaxPage)

}

func (c *Client) ListDeliveryPipelineWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*DeliveryPipelineList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &DeliveryPipeline{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listDeliveryPipeline(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DeliveryPipelineList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetDeliveryPipeline(ctx context.Context, r *DeliveryPipeline) (*DeliveryPipeline, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractDeliveryPipelineFields(r)

	b, err := c.getDeliveryPipelineRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDeliveryPipeline(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDeliveryPipelineNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractDeliveryPipelineFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDeliveryPipeline(ctx context.Context, r *DeliveryPipeline) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("DeliveryPipeline resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting DeliveryPipeline...")
	deleteOp := deleteDeliveryPipelineOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDeliveryPipeline deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDeliveryPipeline(ctx context.Context, project, location string, filter func(*DeliveryPipeline) bool) error {
	listObj, err := c.ListDeliveryPipeline(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllDeliveryPipeline(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDeliveryPipeline(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDeliveryPipeline(ctx context.Context, rawDesired *DeliveryPipeline, opts ...dcl.ApplyOption) (*DeliveryPipeline, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *DeliveryPipeline
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDeliveryPipelineHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyDeliveryPipelineHelper(c *Client, ctx context.Context, rawDesired *DeliveryPipeline, opts ...dcl.ApplyOption) (*DeliveryPipeline, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyDeliveryPipeline...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractDeliveryPipelineFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.deliveryPipelineDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToDeliveryPipelineDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []deliveryPipelineApiOperation
	if create {
		ops = append(ops, &createDeliveryPipelineOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyDeliveryPipelineDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyDeliveryPipelineDiff(c *Client, ctx context.Context, desired *DeliveryPipeline, rawDesired *DeliveryPipeline, ops []deliveryPipelineApiOperation, opts ...dcl.ApplyOption) (*DeliveryPipeline, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetDeliveryPipeline(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDeliveryPipelineOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDeliveryPipeline(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDeliveryPipelineNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDeliveryPipelineNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDeliveryPipelineDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractDeliveryPipelineFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractDeliveryPipelineFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDeliveryPipeline(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}

func (r *DeliveryPipeline) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "", body, nil
}
