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

package iamclient

import (
	"context"
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dcliam "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	ProjectKind            = "Project"
	ResourceManagerGroup   = "resourcemanager.cnrm.cloud.google.com"
	ResourceManagerVersion = "v1beta1"

	SQLInstanceKind = "SQLInstance"
	SQLGroup        = "sql.cnrm.cloud.google.com"
	SQLVersion      = "v1beta1"

	LoggingLogSinkKind = "LoggingLogSink"
	LoggingGroup       = "logging.cnrm.cloud.google.com"
	LoggingVersion     = "v1beta1"

	IAMServiceAccountKind = "IAMServiceAccount"
	IAMGroup              = "iam.cnrm.cloud.google.com"
	IAMVersion            = "v1beta1"

	SerivceIdentityKind = "ServiceIdentity"
	ServiceUsageGroup   = "serviceusage.cnrm.cloud.google.com"
	ServiceUsageVersion = "v1beta1"

	BigQueryConnectionConnectionKind = "BigQueryConnectionConnection"
	BigQueryConnectionGroup          = "bigqueryconnection.cnrm.cloud.google.com"
	BigQueryConnectionVersion        = "v1beta1"
)

var (
	ErrNotFound = k8s.ErrIAMNotFound
	logger      = log.Log.WithName("iamclient")

	ProjectGVK = schema.GroupVersionKind{
		Group:   ResourceManagerGroup,
		Version: ResourceManagerVersion,
		Kind:    ProjectKind,
	}
	SQLInstanceGVK = schema.GroupVersionKind{
		Group:   SQLGroup,
		Version: SQLVersion,
		Kind:    SQLInstanceKind,
	}
	LoggingLogSinkGVK = schema.GroupVersionKind{
		Group:   LoggingGroup,
		Version: LoggingVersion,
		Kind:    LoggingLogSinkKind,
	}
	IAMServiceAccountGVK = schema.GroupVersionKind{
		Group:   IAMGroup,
		Version: IAMVersion,
		Kind:    IAMServiceAccountKind,
	}
	ServiceIdentityGVK = schema.GroupVersionKind{
		Group:   ServiceUsageGroup,
		Version: ServiceUsageVersion,
		Kind:    SerivceIdentityKind,
	}
	BigQueryConnectionConnectionGVK = schema.GroupVersionKind{
		Group:   BigQueryConnectionGroup,
		Version: BigQueryConnectionVersion,
		Kind:    BigQueryConnectionConnectionKind,
	}
)

// idTemplateVarsRegex is a regex used to match named tokens in an id template
// (e.g. "{{project}}" and "{{name}}" in "projects/{{project}}/global/networks/{{name}}"
var idTemplateVarsRegex = regexp.MustCompile(`{{[a-z]([a-zA-Z0-9\-_.]*[a-zA-Z0-9])?}}`)

type IAMClient struct {
	TFIAMClient  *TFIAMClient
	DCLIAMClient *DCLIAMClient
	kubeClient   client.Client
}

func New(tfProvider *tfschema.Provider,
	smLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client,
	converter *conversion.Converter,
	dclConfig *mmdcl.Config) *IAMClient {
	tfIAMClient := TFIAMClient{
		kubeClient: kubeClient,
		provider:   tfProvider,
		smLoader:   smLoader,
	}
	dclIAMClient := DCLIAMClient{
		dclClient: &dcliam.Client{
			Config: dclConfig,
		},
		converter:  converter,
		smLoader:   smLoader,
		kubeClient: kubeClient,
	}
	iamClient := IAMClient{
		TFIAMClient:  &tfIAMClient,
		DCLIAMClient: &dclIAMClient,
		kubeClient:   kubeClient,
	}
	return &iamClient
}

func (c *IAMClient) SetPolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	if registry.IsIAMDirect(policyMember.Spec.ResourceReference.GroupVersionKind().GroupKind()) {
		id, err := ResolveMemberIdentity(ctx, policyMember.Spec.Member, policyMember.Spec.MemberFrom, policyMember.GetNamespace(), c.TFIAMClient)
		if err != nil {
			return nil, err
		}

		return direct.SetIAMPolicyMember(ctx, c.kubeClient, policyMember, v1beta1.Member(id))
	}

	if c.isDCLBasedIAMResource(policyMember) {
		return c.DCLIAMClient.SetPolicyMember(ctx, c.TFIAMClient, policyMember)
	}
	return c.TFIAMClient.SetPolicyMember(ctx, policyMember)
}

func (c *IAMClient) GetPolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	if registry.IsIAMDirect(policyMember.Spec.ResourceReference.GroupVersionKind().GroupKind()) {
		id, err := ResolveMemberIdentity(ctx, policyMember.Spec.Member, policyMember.Spec.MemberFrom, policyMember.GetNamespace(), c.TFIAMClient)
		if err != nil {
			return nil, err
		}

		return direct.GetIAMPolicyMember(ctx, c.kubeClient, policyMember, v1beta1.Member(id))
	}
	if c.isDCLBasedIAMResource(policyMember) {
		return c.DCLIAMClient.GetPolicyMember(ctx, c.TFIAMClient, policyMember)
	}
	return c.TFIAMClient.GetPolicyMember(ctx, policyMember)
}

func (c *IAMClient) DeletePolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember) error {
	if registry.IsIAMDirect(policyMember.Spec.ResourceReference.GroupVersionKind().GroupKind()) {
		id, err := ResolveMemberIdentity(ctx, policyMember.Spec.Member, policyMember.Spec.MemberFrom, policyMember.GetNamespace(), c.TFIAMClient)
		if err != nil {
			return err
		}

		return direct.DeleteIAMPolicyMember(ctx, c.kubeClient, policyMember, v1beta1.Member(id))
	}

	if c.isDCLBasedIAMResource(policyMember) {
		return c.DCLIAMClient.DeletePolicyMember(ctx, c.TFIAMClient, policyMember)

	}
	return c.TFIAMClient.DeletePolicyMember(ctx, policyMember)
}

func (c *IAMClient) SetPolicy(ctx context.Context, policy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	if c.isDCLBasedIAMResource(policy) {
		return c.DCLIAMClient.SetPolicy(ctx, policy)
	}
	return c.TFIAMClient.SetPolicy(ctx, policy)
}

func (c *IAMClient) GetPolicy(ctx context.Context, policy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	if c.isDCLBasedIAMResource(policy) {
		return c.DCLIAMClient.GetPolicy(ctx, policy)
	}
	return c.TFIAMClient.GetPolicy(ctx, policy)
}

func (c *IAMClient) DeletePolicy(ctx context.Context, policy *v1beta1.IAMPolicy) error {
	if c.isDCLBasedIAMResource(policy) {
		return c.DCLIAMClient.DeletePolicy(ctx, policy)
	}
	return c.TFIAMClient.DeletePolicy(ctx, policy)
}

func (c *IAMClient) SetAuditConfig(ctx context.Context, auditConfig *v1beta1.IAMAuditConfig) (*v1beta1.IAMAuditConfig, error) {
	if c.isDCLBasedIAMResource(auditConfig) {
		return nil, fmt.Errorf("resource with gvk %v does not have AuditConfig support right now", auditConfig.Spec.ResourceReference.GroupVersionKind())
	}
	return c.TFIAMClient.SetAuditConfig(ctx, auditConfig)
}

func (c *IAMClient) GetAuditConfig(ctx context.Context, auditConfig *v1beta1.IAMAuditConfig) (*v1beta1.IAMAuditConfig, error) {
	if c.isDCLBasedIAMResource(auditConfig) {
		return nil, fmt.Errorf("resource with gvk %v does not have AuditConfig support right now", auditConfig.Spec.ResourceReference.GroupVersionKind())
	}
	return c.TFIAMClient.GetAuditConfig(ctx, auditConfig)
}

func (c *IAMClient) DeleteAuditConfig(ctx context.Context, auditConfig *v1beta1.IAMAuditConfig) error {
	if c.isDCLBasedIAMResource(auditConfig) {
		return fmt.Errorf("resource with gvk %v does not have AuditConfig support right now", auditConfig.Spec.ResourceReference.GroupVersionKind())
	}
	return c.TFIAMClient.DeleteAuditConfig(ctx, auditConfig)
}
