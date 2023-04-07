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

package allowlist

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
)

var (
	AlphaVersion = "v1alpha1"
	betaVersion  = "v1beta1"
	// alphaAllowlist holds the list of the resources to be allowlisted as
	// v1alpha1 CRDs. The format is '[terraform_product_name]/[terraform_type_name]'.
	// 'google_[terraform_product_name]' should be the prefix of
	// '[terraform_type_name]'.
	alphaAllowlist = []string{
		"access_context_manager/google_access_context_manager_service_perimeter_resource",
	}
	// betaAllowlist holds the list of the resources to be allowlisted as
	// v1beta1 CRDs. The format is '[terraform_product_name]/[terraform_type_name]'.
	// 'google_[terraform_product_name]' should be the prefix of
	// '[terraform_type_name]'.
	betaAllowlist = []string{
		"bigquery/google_bigquery_routine",
		"data_catalog/google_data_catalog_policy_tag",
		"data_catalog/google_data_catalog_taxonomy",
		"tags/google_tags_tag_binding",
		"tags/google_tags_tag_key",
		"tags/google_tags_tag_value",
	}
)

type AutoGenType struct {
	ServiceNameInLC string
	KRMKindName     string
	TFTypeName      string
	Version         string
}

func (a *AutoGenType) loadKRMKindFromSM(smAndRCMap map[string]map[string]string) error {
	service := a.ServiceNameInLC
	tfType := a.TFTypeName
	rcMap, ok := smAndRCMap[a.ServiceNameInLC]
	if !ok {
		return fmt.Errorf("can't find allowlisted service %v "+
			"in generated service mappings", service)
	}
	krmKind, ok := rcMap[tfType]
	if !ok {
		return fmt.Errorf("can't find allowlisted type %v "+
			"under service %v in auto-generated service mappings",
			tfType, service)
	}
	a.KRMKindName = krmKind
	return nil
}

func newAutoGenType(autoGenTypeInString string, version string) (*AutoGenType, error) {
	parts := strings.Split(autoGenTypeInString, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("type for resource auto-generation should be"+
			" in the format '[terraform_product_name]/[terraform_type_name]', split by one '/',"+
			" but the provided type is %q", autoGenTypeInString)
	}

	if !text.IsSnakeCase(parts[0]) && !text.IsSnakeCase(parts[1]) {
		return nil, fmt.Errorf("type for resource auto-generation should be"+
			" in the format '[terraform_product_name]/[terraform_type_name]', both terraform_product_name"+
			" and terraform_type_name should be in snake case, but the provided"+
			" type is %q", autoGenTypeInString)
	}

	return &AutoGenType{
		ServiceNameInLC: strings.Replace(parts[0], "_", "", -1),
		TFTypeName:      parts[1],
		Version:         version,
	}, nil
}

type AutoGenAllowlist struct {
	ServiceAndTFTypes  map[string]map[string]*AutoGenType
	ServiceAndKRMKinds map[string]map[string]*AutoGenType
	KRMKinds           map[string]*AutoGenType
}

func (l *AutoGenAllowlist) HasService(serviceNameInLC string) bool {
	_, ok := l.ServiceAndTFTypes[serviceNameInLC]
	return ok
}

func (l *AutoGenAllowlist) GetTFTypeInService(serviceNameInLC, tfType string) (*AutoGenType, bool) {
	resourceMap, ok := l.ServiceAndTFTypes[serviceNameInLC]
	if !ok {
		return nil, false
	}
	autoGenType, ok := resourceMap[tfType]
	return autoGenType, ok
}

func (l *AutoGenAllowlist) GetKRMKind(krmKind string) (*AutoGenType, bool) {
	autoGenType, ok := l.KRMKinds[krmKind]
	return autoGenType, ok
}

func (l *AutoGenAllowlist) HasKRMKindInService(serviceNameInLC, krmKind string) bool {
	resourceMap, ok := l.ServiceAndKRMKinds[serviceNameInLC]
	if !ok {
		return false
	}
	_, ok = resourceMap[krmKind]
	return ok
}

func (l *AutoGenAllowlist) addAutoGenType(autoGenType *AutoGenType) error {
	_, ok := l.ServiceAndTFTypes[autoGenType.ServiceNameInLC]
	if !ok {
		l.ServiceAndTFTypes[autoGenType.ServiceNameInLC] = make(map[string]*AutoGenType)
		l.ServiceAndKRMKinds[autoGenType.ServiceNameInLC] = make(map[string]*AutoGenType)
	}
	TFTypeMap, _ := l.ServiceAndTFTypes[autoGenType.ServiceNameInLC]
	KRMKindMap, _ := l.ServiceAndKRMKinds[autoGenType.ServiceNameInLC]
	_, ok = TFTypeMap[autoGenType.TFTypeName]
	if ok {
		return fmt.Errorf("TF type %v has already been allowlisted under "+
			"service %v", autoGenType.TFTypeName, autoGenType.ServiceNameInLC)
	}

	TFTypeMap[autoGenType.TFTypeName] = autoGenType
	KRMKindMap[autoGenType.KRMKindName] = autoGenType
	l.KRMKinds[autoGenType.KRMKindName] = autoGenType
	return nil
}

func NewAutoGenAllowlist() *AutoGenAllowlist {
	return &AutoGenAllowlist{
		ServiceAndTFTypes:  make(map[string]map[string]*AutoGenType),
		ServiceAndKRMKinds: make(map[string]map[string]*AutoGenType),
		KRMKinds:           make(map[string]*AutoGenType),
	}
}

func LoadAutoGenAllowList(generatedSMMap map[string]v1alpha1.ServiceMapping) (*AutoGenAllowlist, error) {
	smAndRCMap := getGeneratedSMAndRCMap(generatedSMMap)
	autoGenAllowlist := NewAutoGenAllowlist()
	for _, typeInString := range alphaAllowlist {
		autoGenType, err := newAutoGenType(typeInString, AlphaVersion)
		if err != nil {
			return nil, fmt.Errorf("error converting allowlisted type %v from string to AutoGenType: %w", typeInString, err)
		}
		if err := autoGenType.loadKRMKindFromSM(smAndRCMap); err != nil {
			return nil, fmt.Errorf("error loading KRMKind for allowlisted type %v: %w", typeInString, err)
		}
		if err := autoGenAllowlist.addAutoGenType(autoGenType); err != nil {
			return nil, fmt.Errorf("error adding AutoGenType for %v into the AutoGenAllowlist: %w", typeInString, err)
		}
	}
	for _, typeInString := range betaAllowlist {
		autoGenType, err := newAutoGenType(typeInString, betaVersion)
		if err != nil {
			return nil, fmt.Errorf("error converting allowlisted type %v from string to AutoGenType: %w", typeInString, err)
		}
		if err := autoGenType.loadKRMKindFromSM(smAndRCMap); err != nil {
			return nil, fmt.Errorf("error loading KRMKind for allowlisted type %v: %w", typeInString, err)
		}
		if err := autoGenAllowlist.addAutoGenType(autoGenType); err != nil {
			return nil, fmt.Errorf("error adding AutoGenType for %v into the AutoGenAllowlist: %w", typeInString, err)
		}
	}
	return autoGenAllowlist, nil
}

func getGeneratedSMAndRCMap(generatedSMMap map[string]v1alpha1.ServiceMapping) map[string]map[string]string {
	smAndRCMap := make(map[string]map[string]string)
	for smName, sm := range generatedSMMap {
		service := strings.TrimSuffix(smName, ".cnrm.cloud.google.com")
		generatedRCMap := make(map[string]string)
		for _, rc := range sm.Spec.Resources {
			generatedRCMap[rc.Name] = rc.Kind
		}
		smAndRCMap[service] = generatedRCMap
	}
	return smAndRCMap
}
