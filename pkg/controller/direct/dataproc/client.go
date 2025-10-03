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

// +tool:controller-client
// proto.service: google.cloud.dataproc.v1.AutoscalingPolicyService
// proto.service: google.cloud.dataproc.v1.BatchController
// proto.service: google.cloud.dataproc.v1.ClusterController
// proto.service: google.cloud.dataproc.v1.JobController
// proto.service: google.cloud.dataproc.v1.NodeGroupController
// proto.service: google.cloud.dataproc.v1.SessionController
// proto.service: google.cloud.dataproc.v1.SessionTemplateController
// proto.service: google.cloud.dataproc.v1.WorkflowTemplateService

package dataproc

import (
	"context"
	"fmt"

	api "cloud.google.com/go/dataproc/v2/apiv1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
)

type gcpClient struct {
	config *config.ControllerConfig
}

func newGCPClient(ctx context.Context, config *config.ControllerConfig) (*gcpClient, error) {
	gcpClient := &gcpClient{
		config: config,
	}
	return gcpClient, nil
}

func (m *gcpClient) newAutoscalingPolicyClient(ctx context.Context) (*api.AutoscalingPolicyClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewAutoscalingPolicyRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc autoscalingpolicy client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newBatchControllerClient(ctx context.Context) (*api.BatchControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewBatchControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc batchcontroller client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newClusterControllerClient(ctx context.Context) (*api.ClusterControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewClusterControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc clustercontroller client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newJobControllerClient(ctx context.Context) (*api.JobControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewJobControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc jobcontroller client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newNodeGroupControllerClient(ctx context.Context) (*api.NodeGroupControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewNodeGroupControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc nodegroupcontroller client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newSessionControllerClient(ctx context.Context) (*api.SessionControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewSessionControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc sessioncontroller client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newSessionTemplateControllerClient(ctx context.Context) (*api.SessionTemplateControllerClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewSessionTemplateControllerRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc sessiontemplatecontroller client: %w", err)
	}
	return client, err
}

func (m *gcpClient) newWorkflowTemplateClient(ctx context.Context) (*api.WorkflowTemplateClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	client, err := api.NewWorkflowTemplateRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building dataproc workflowtemplate client: %w", err)
	}
	return client, err
}
