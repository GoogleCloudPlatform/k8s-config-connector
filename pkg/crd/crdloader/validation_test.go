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

package crdloader_test

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestCRDObjectTypes(t *testing.T) {
	// knownInvalidCRDs is a list of CRDs that currently fail the validation.
	// We want to eventually fix these, but for now we allowlist them so the test passes.
	// This allows us to detect new regressions.
	knownInvalidCRDs := map[string]bool{
		"accesscontextmanageraccesslevels.accesscontextmanager.cnrm.cloud.google.com":        true,
		"aiplatformmodels.aiplatform.cnrm.cloud.google.com":                                 true,
		"apigeeenvironments.apigee.cnrm.cloud.google.com":                                   true,
		"apigeeorganizations.apigee.cnrm.cloud.google.com":                                   true,
		"bigqueryconnectionconnections.bigqueryconnection.cnrm.cloud.google.com":            true,
		"bigquerydatapolicies.bigquerydatapolicy.cnrm.cloud.google.com":                     true,
		"bigquerydatatransferconfigs.bigquerydatatransfer.cnrm.cloud.google.com":            true,
		"bigquerytables.bigquery.cnrm.cloud.google.com":                                     true,
		"bigtableauthorizedviews.bigtable.cnrm.cloud.google.com":                            true,
		"bigtablelogicalviews.bigtable.cnrm.cloud.google.com":                               true,
		"bigtablematerializedviews.bigtable.cnrm.cloud.google.com":                          true,
		"cloudbuildtriggers.cloudbuild.cnrm.cloud.google.com":                               true,
		"clouddmsmigrationjobs.clouddms.cnrm.cloud.google.com":                              true,
		"datacatalogentries.datacatalog.cnrm.cloud.google.com":                              true,
		"datacatalogpolicytags.datacatalog.cnrm.cloud.google.com":                           true,
		"dataformrepositories.dataform.cnrm.cloud.google.com":                               true,
		"dataprocjobs.dataproc.cnrm.cloud.google.com":                                       true,
		"dataprocnodegroups.dataproc.cnrm.cloud.google.com":                                 true,
		"datastreamconnectionprofiles.datastream.cnrm.cloud.google.com":                     true,
		"discoveryengineengines.discoveryengine.cnrm.cloud.google.com":                      true,
		"firestorebackupschedules.firestore.cnrm.cloud.google.com":                          true,
		"firestorefields.firestore.cnrm.cloud.google.com":                                   true,
		"iamdenypolicies.iam.cnrm.cloud.google.com":                                         true,
		"memorystoreinstances.memorystore.cnrm.cloud.google.com":                            true,
		"monitoringdashboards.monitoring.cnrm.cloud.google.com":                             true,
		"recaptchaenterprisefirewallpolicies.recaptchaenterprise.cnrm.cloud.google.com":     true,
		"servicenetworkingpeereddnsdomains.servicenetworking.cnrm.cloud.google.com":         true,
		"spannerbackupschedules.spanner.cnrm.cloud.google.com":                              true,
		"vertexaiindexes.vertexai.cnrm.cloud.google.com":                                    true,
	}

	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	for _, crd := range crds {
		t.Run(crd.Name, func(t *testing.T) {
			isKnownInvalid := knownInvalidCRDs[crd.Name]
			invalidVersions := 0
			for _, version := range crd.Spec.Versions {
				if version.Schema == nil || version.Schema.OpenAPIV3Schema == nil {
					continue
				}
				schema := version.Schema.OpenAPIV3Schema
				for name, subProps := range schema.Properties {
					if name == "metadata" {
						continue
					}
					if err := validateProps(&subProps, fmt.Sprintf("%s.%s", version.Name, name)); err != nil {
						if isKnownInvalid {
							t.Logf("KNOWN INVALID: version %s is invalid: %v", version.Name, err)
							invalidVersions++
						} else {
							t.Errorf("version %s is invalid: %v", version.Name, err)
						}
					}
				}
			}
			if isKnownInvalid && invalidVersions == 0 {
				t.Errorf("CRD %s is in knownInvalidCRDs but passed validation; please remove it from the list", crd.Name)
			}
		})
	}
}

func validateProps(props *apiextensions.JSONSchemaProps, path string) error {
	if props.Type == "object" {
		if len(props.Properties) == 0 && props.AdditionalProperties == nil && (props.XPreserveUnknownFields == nil || !*props.XPreserveUnknownFields) {
			return fmt.Errorf("object at %s is missing properties, additionalProperties, or x-kubernetes-preserve-unknown-fields", path)
		}
	}
	for name, subProps := range props.Properties {
		if err := validateProps(&subProps, path+"."+name); err != nil {
			return err
		}
	}
	if props.Items != nil {
		if props.Items.Schema != nil {
			if err := validateProps(props.Items.Schema, path+"[]"); err != nil {
				return err
			}
		}
		for i := range props.Items.JSONSchemas {
			if err := validateProps(&props.Items.JSONSchemas[i], fmt.Sprintf("%s[%d]", path, i)); err != nil {
				return err
			}
		}
	}
	if props.AdditionalProperties != nil && props.AdditionalProperties.Schema != nil {
		if err := validateProps(props.AdditionalProperties.Schema, path+"[*]"); err != nil {
			return err
		}
	}
	for i := range props.AllOf {
		if err := validateProps(&props.AllOf[i], fmt.Sprintf("%s.allOf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i := range props.AnyOf {
		if err := validateProps(&props.AnyOf[i], fmt.Sprintf("%s.anyOf[%d]", path, i)); err != nil {
			return err
		}
	}
	for i := range props.OneOf {
		if err := validateProps(&props.OneOf[i], fmt.Sprintf("%s.oneOf[%d]", path, i)); err != nil {
			return err
		}
	}
	if props.Not != nil {
		if err := validateProps(props.Not, path+".not"); err != nil {
			return err
		}
	}
	return nil
}
