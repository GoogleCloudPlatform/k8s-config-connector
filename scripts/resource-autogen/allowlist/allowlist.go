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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
)

var (
	AlphaVersion   = "v1alpha1"
	betaVersion    = "v1beta1"
	alphaAllowlist = []string{}
	betaAllowlist  = []string{
		"BigQuery/Routine",
		"DataCatalog/PolicyTag",
		"DataCatalog/Taxonomy",
		"Tags/TagBinding",
		"Tags/TagKey",
		"Tags/TagValue",
	}
	tfLegacyServiceNames = map[string]string{
		"BigQuery": "bigquery",
	}
)

type AutoGenType struct {
	ServiceName  string
	ResourceName string
	Version      string
}

func (t *AutoGenType) toTFType() string {
	tfType := "google_"
	legacyName, ok := tfLegacyServiceNames[t.ServiceName]
	if ok {
		tfType += legacyName
	} else {
		tfType += text.AsSnakeCase(t.ServiceName)
	}
	tfType += "_" + text.AsSnakeCase(t.ResourceName)
	return tfType
}

func (t *AutoGenType) ToKRMKind() string {
	return t.ServiceName + t.ResourceName
}

func NewAutoGenType(autoGenTypeInString string, version string) (*AutoGenType, error) {
	parts := strings.Split(autoGenTypeInString, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("type for resource auto-generation should be"+
			" in the format '[ServiceName]/[ResourceName], split by one '/',"+
			" but the provided type is %q", autoGenTypeInString)
	}

	if !text.IsPascalCase(parts[0]) || !text.IsPascalCase(parts[1]) {
		return nil, fmt.Errorf("type for resource auto-generation should be"+
			" in the format '[ServiceName]/[ResourceName], both ServiceName"+
			" and ResourceName should be in pascal case, but the provided"+
			" type is %q", autoGenTypeInString)
	}

	return &AutoGenType{ServiceName: parts[0], ResourceName: parts[1], Version: version}, nil
}

type AutoGenAllowlist struct {
	ServiceAndTFTypes  map[string]map[string]*AutoGenType
	ServiceAndKRMKinds map[string]map[string]*AutoGenType
	KRMKinds           map[string]*AutoGenType
}

func (a *AutoGenAllowlist) addAutoGenType(autoGenType *AutoGenType) {
	_, ok := a.ServiceAndTFTypes[autoGenType.ServiceName]
	if !ok {
		a.ServiceAndTFTypes[autoGenType.ServiceName] = make(map[string]*AutoGenType)
		a.ServiceAndKRMKinds[autoGenType.ServiceName] = make(map[string]*AutoGenType)
	}
	TFTypeMap, _ := a.ServiceAndTFTypes[autoGenType.ServiceName]
	KRMKindMap, _ := a.ServiceAndKRMKinds[autoGenType.ServiceName]
	_, ok = TFTypeMap[autoGenType.toTFType()]
	if !ok {
		TFTypeMap[autoGenType.toTFType()] = autoGenType
		KRMKindMap[autoGenType.ToKRMKind()] = autoGenType
		a.KRMKinds[autoGenType.ToKRMKind()] = autoGenType
	}
	return
}

func (a *AutoGenAllowlist) HasService(service string) bool {
	_, ok := a.ServiceAndTFTypes[service]
	return ok
}

func (a *AutoGenAllowlist) GetTFTypeInService(service, tfType string) (*AutoGenType, bool) {
	resourceMap, ok := a.ServiceAndTFTypes[service]
	if !ok {
		return nil, false
	}
	autoGenType, ok := resourceMap[tfType]
	return autoGenType, ok
}

func (a *AutoGenAllowlist) GetKRMKind(krmKind string) (*AutoGenType, bool) {
	autoGenType, ok := a.KRMKinds[krmKind]
	return autoGenType, ok
}

func (a *AutoGenAllowlist) HasKRMKindInService(service, krmKind string) bool {
	resourceMap, ok := a.ServiceAndKRMKinds[service]
	if !ok {
		return false
	}
	_, ok = resourceMap[krmKind]
	return ok
}

func NewAutoGenAllowlist() *AutoGenAllowlist {
	return &AutoGenAllowlist{
		ServiceAndTFTypes:  make(map[string]map[string]*AutoGenType),
		ServiceAndKRMKinds: make(map[string]map[string]*AutoGenType),
		KRMKinds:           make(map[string]*AutoGenType),
	}
}

func LoadAutoGenAllowList() (*AutoGenAllowlist, error) {
	autoGenAllowlist := NewAutoGenAllowlist()
	for _, typeInString := range alphaAllowlist {
		autoGenType, err := NewAutoGenType(typeInString, AlphaVersion)
		if err != nil {
			return nil, fmt.Errorf("error converting the types in the "+
				"alphaAllowlist from string to AutoGenType: %w", err)
		}
		autoGenAllowlist.addAutoGenType(autoGenType)
	}
	for _, typeInString := range betaAllowlist {
		autoGenType, err := NewAutoGenType(typeInString, betaVersion)
		if err != nil {
			return nil, fmt.Errorf("error converting the types in the "+
				"betaAllowlist from string to AutoGenType: %w", err)
		}
		autoGenAllowlist.addAutoGenType(autoGenType)
	}
	return autoGenAllowlist, nil
}
