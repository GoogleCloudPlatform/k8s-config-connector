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

package webhook

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/certclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/generator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook/cert/writer"

	admissionregistration "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	ValidatingWebhookConfigurationName = "validating-webhook.cnrm.cloud.google.com"
	MutatingWebhookConfigurationName   = "mutating-webhook.cnrm.cloud.google.com"
	CommonWebhookServiceName           = "cnrm-validating-webhook"
)

func RegisterCommonWebhooks(mgr manager.Manager, nocacheClient client.Client) error {
	fmt.Println("starting up webhooks")
	whCfgs, err := GetCommonWebhookConfigs()
	if err != nil {
		return fmt.Errorf("error getting common wehbook configs: %w", err)
	}
	return register(
		ValidatingWebhookConfigurationName,
		MutatingWebhookConfigurationName,
		CommonWebhookServiceName,
		"cnrm-webhook-manager",
		whCfgs,
		mgr,
		nocacheClient,
	)
}

func GetCommonWebhookConfigs() ([]Config, error) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		return nil, fmt.Errorf("error getting new service mapping loader: %w", err)
	}
	dclSchemaLoader, err := dclschemaloader.New()
	if err != nil {
		return nil, fmt.Errorf("error getting new dcl schema loader: %w", err)
	}
	serviceMetadataLoader := metadata.New()
	allGVKs, err := supportedgvks.All(smLoader, serviceMetadataLoader)
	if err != nil {
		return nil, fmt.Errorf("error loading all supported GVKs: %w", err)
	}
	allResourcesRules := getRulesFromResources(allGVKs)
	dynamicResourcesRules := getRulesFromResources(supportedgvks.AllDynamicTypes(smLoader, serviceMetadataLoader))
	handwrittenIamResourcesRules := getRulesFromResources(supportedgvks.BasedOnHandwrittenIAMTypes())
	resourcesWithOverridesRules := getRulesForResourcesWithCustomValidation(allGVKs)
	whCfgs := []Config{
		{
			Name:          "deny-immutable-field-updates.cnrm.cloud.google.com",
			Path:          "/deny-immutable-field-updates",
			Type:          Validating,
			HandlerFunc:   NewRequestLoggingHandler(NewImmutableFieldsValidatorHandler(smLoader, dclSchemaLoader, serviceMetadataLoader), "immutable fields validation"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(
				allResourcesRules,
				admissionregistration.Update,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "deny-unknown-fields.cnrm.cloud.google.com",
			Path:          "/deny-unknown-fields",
			Type:          Validating,
			HandlerFunc:   NewRequestLoggingHandler(NewNoUnknownFieldsValidatorHandler(smLoader), "unknown fields validation"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(
				allResourcesRules,
				admissionregistration.Create,
				admissionregistration.Update,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "iam-validation.cnrm.cloud.google.com",
			Path:          "/iam-validation",
			Type:          Validating,
			HandlerFunc:   NewRequestLoggingHandler(NewIAMValidatorHandler(smLoader, serviceMetadataLoader, dclSchemaLoader), "iam validation"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(handwrittenIamResourcesRules,
				admissionregistration.Create,
				admissionregistration.Update,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "iam-defaulter.cnrm.cloud.google.com",
			Path:          "/iam-defaulter",
			Type:          Mutating,
			HandlerFunc:   NewRequestLoggingHandler(NewIAMDefaulter(smLoader, serviceMetadataLoader), "iam defaulter"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(handwrittenIamResourcesRules,
				admissionregistration.Create,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "container-annotation-handler.cnrm.cloud.google.com",
			Path:          "/container-annotation-handler",
			Type:          Mutating,
			HandlerFunc:   NewRequestLoggingHandler(NewContainerAnnotationHandler(smLoader, dclSchemaLoader, serviceMetadataLoader), "container annotation handler"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(
				dynamicResourcesRules,
				admissionregistration.Create,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "management-conflict-annotation-defaulter.cnrm.cloud.google.com",
			Path:          "/management-conflict-annotation-defaulter",
			Type:          Mutating,
			HandlerFunc:   NewRequestLoggingHandler(NewManagementConflictAnnotationDefaulter(smLoader, dclSchemaLoader, serviceMetadataLoader), "management conflict annotation defaulter"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(
				dynamicResourcesRules,
				admissionregistration.Create,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "generic-defaulter.cnrm.cloud.google.com",
			Path:          "/generic-defaulter",
			Type:          Mutating,
			HandlerFunc:   NewRequestLoggingHandler(NewGenericDefaulter(), "generic defaulter"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(
				dynamicResourcesRules,
				admissionregistration.Create,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "resource-validation.cnrm.cloud.google.com",
			Path:          "/resource-validation",
			Type:          Validating,
			HandlerFunc:   NewRequestLoggingHandler(NewResourceValidatorHandler(), "resource validation"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(resourcesWithOverridesRules,
				admissionregistration.Create,
				admissionregistration.Update,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
		{
			Name:          "state-into-spec-validation.cnrm.cloud.google.com",
			Path:          "/state-into-spec-validation",
			Type:          Validating,
			HandlerFunc:   NewRequestLoggingHandler(NewStateIntoSpecAnnotationValidatorHandler(), "state-into-spec validation"),
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(allResourcesRules,
				admissionregistration.Create,
				admissionregistration.Update,
			),
			SideEffects: admissionregistration.SideEffectClassNone,
		},
	}
	return whCfgs, nil
}

func RegisterAbandonOnUninstallWebhook(mgr manager.Manager, nocacheClient client.Client) error {
	whCfgs := []Config{
		{
			Name:        "abandon-on-uninstall.cnrm.cloud.google.com",
			Path:        "/abandon-on-uninstall",
			Type:        Validating,
			HandlerFunc: NewAbandonOnCRDUninstallWebhook(),
			ObjectSelector: &metav1.LabelSelector{
				// The MatchLabels will not match anything with the value "no-op"
				// specified. We want the webhook to intercept nothing before we
				// work on cleaning up the webhook from existing clusters.
				MatchLabels: map[string]string{
					crdgeneration.ManagedByKCCLabel: "no-op",
				},
			},
			FailurePolicy: admissionregistration.Fail,
			Rules: getRulesForOperationTypes(
				getRulesFromResources([]schema.GroupVersionKind{
					{
						Group:   apiextensions.GroupName,
						Version: apiextensions.SchemeGroupVersion.Version,
						Kind:    "CustomResourceDefinition",
					},
				}),
				admissionregistration.Delete,
			),
			SideEffects: admissionregistration.SideEffectClassNoneOnDryRun,
		},
	}
	return register(
		"abandon-on-uninstall.cnrm.cloud.google.com",
		"",
		"abandon-on-uninstall",
		"cnrm-deletiondefender",
		whCfgs,
		mgr,
		nocacheClient,
	)
}

func register(validatingWebhookConfigurationName, mutatingWebhookConfigurationName, serviceName, componentName string,
	whCfgs []Config, mgr manager.Manager, nocacheClient client.Client) error {
	ctx := context.TODO()

	validatingWebhookCfg, mutatingWebhookCfg := GenerateWebhookManifests(
		validatingWebhookConfigurationName,
		mutatingWebhookConfigurationName,
		serviceName,
		whCfgs,
	)
	manifests := make([]client.Object, 0)
	if validatingWebhookCfg != nil {
		manifests = append(manifests, validatingWebhookCfg)
	}
	if mutatingWebhookCfg != nil {
		manifests = append(manifests, mutatingWebhookCfg)
	}
	svc := generateService(
		serviceName,
		map[string]string{
			k8s.KCCComponentLabel: componentName,
			k8s.KCCSystemLabel:    "true",
		})
	writerOpts := writer.SecretCertWriterOptions{
		Client: nocacheClient,
		Secret: &types.NamespacedName{
			Name:      formatSecretName(serviceName),
			Namespace: k8s.SystemNamespace,
		},
	}
	certWriter, err := writer.NewSecretCertWriter(writerOpts)
	if err != nil {
		return fmt.Errorf("error creating secret cert writer: %w", err)
	}
	certClient, err := certclient.New(certclient.Options{
		WebhookManifests: manifests,
		Service:          svc,
		KubeClient:       nocacheClient,
		CertWriter:       certWriter,
	})
	if err != nil {
		return fmt.Errorf("error creating cert client: %w", err)
	}
	// Do an initial call so we can guarantee the webhook configuration resources are
	// registered in the API server before marking this container as ready.
	if err := certClient.RefreshCertsAndInstall(ctx); err != nil {
		return fmt.Errorf("error refreshing certs and installing manifests: %w", err)
	}
	if err := mgr.Add(certClient); err != nil {
		return fmt.Errorf("error registering cert client with manager: %w", err)
	}
	if err := persistCertificatesToDisk(certWriter, svc); err != nil {
		return err
	}
	// Set up the HTTP server
	s := webhook.NewServer(webhook.Options{
		CertDir:  certDir,
		CertName: writer.ServerCertName,
		KeyName:  writer.ServerKeyName,
		Port:     ServicePort,
	})
	for _, whCfg := range whCfgs {
		handler := whCfg.HandlerFunc(mgr)
		s.Register(whCfg.Path, &admission.Webhook{Handler: handler})
	}
	if err := mgr.Add(s); err != nil {
		return fmt.Errorf("error adding webhook server to manager: %w", err)
	}
	return nil
}

func formatSecretName(serviceName string) string {
	// both webhook and deletiondefender use this package to register their admission handler(s) and the two cannot
	// share the same certificate so append a service-name suffix
	return fmt.Sprintf("%v-%v", certSecretName, serviceName)
}

// the webhook.Server requires the certificates to be present on the local filesystem so fetch them from the API
// server and persist them to disk
func persistCertificatesToDisk(certWriter writer.CertWriter, svc *corev1.Service) error {
	dnsName := getDNSNameForService(svc)
	artifacts, _, err := certWriter.EnsureCert(dnsName)
	if err != nil {
		return fmt.Errorf("error ensuring certificate: %w", err)
	}

	return writeCertificates(artifacts, certDir, writer.ServerCertName, writer.ServerKeyName)
}

func writeCertificates(artifacts *generator.Artifacts, dir, certName, keyName string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("error removing cert dir '%v': %w", dir, err)
	}
	if err := os.MkdirAll(dir, 0777); err != nil {
		return fmt.Errorf("error creating cert dir '%v': %w", dir, err)
	}
	perms := os.FileMode(0644)
	certPath := path.Join(dir, certName)
	if err := os.WriteFile(certPath, artifacts.Cert, perms); err != nil {
		return fmt.Errorf("error writing certificate to '%v': %w", certPath, err)
	}
	keyPath := path.Join(dir, keyName)
	if err := os.WriteFile(keyPath, artifacts.Key, perms); err != nil {
		return fmt.Errorf("error writing key to '%v': %w", keyPath, err)
	}
	return nil
}

func getDNSNameForService(svc *corev1.Service) string {
	// the following line of logic for calculating the dnsName is taken from provisioner.dnsNameFromClientConfig(...)
	// that code is not easily reused
	return generator.ServiceToCommonName(svc.Namespace, svc.Name)
}

func getRulesForResourcesWithCustomValidation(allGVKs []schema.GroupVersionKind) map[string]*admissionregistration.Rule {
	resources := make([]schema.GroupVersionKind, 0)
	for _, gvk := range allGVKs {
		if resourceoverrides.Handler.HasConfigValidate(gvk.Kind) {
			resources = append(resources, gvk)
		}
	}
	return getRulesFromResources(resources)
}

func getRulesFromResources(resources []schema.GroupVersionKind) map[string]*admissionregistration.Rule {
	groupToRule := make(map[string]*admissionregistration.Rule)
	for _, gvk := range resources {
		rule, ok := groupToRule[gvk.Group]
		if !ok {
			rule = &admissionregistration.Rule{
				APIGroups:   []string{gvk.Group},
				APIVersions: []string{},
				Resources:   []string{},
			}
			groupToRule[gvk.Group] = rule
		}
		rule.APIVersions = slice.IncludeString(rule.APIVersions, gvk.Version)
		rule.Resources = slice.IncludeString(rule.Resources, text.Pluralize(strings.ToLower(gvk.Kind)))
	}
	return groupToRule
}

func getRulesForOperationTypes(groupToRule map[string]*admissionregistration.Rule,
	ops ...admissionregistration.OperationType) []admissionregistration.RuleWithOperations {
	rules := make([]admissionregistration.RuleWithOperations, 0)
	for _, rule := range groupToRule {
		rules = append(rules, admissionregistration.RuleWithOperations{
			Operations: ops,
			Rule:       *rule,
		})
	}
	return rules
}
