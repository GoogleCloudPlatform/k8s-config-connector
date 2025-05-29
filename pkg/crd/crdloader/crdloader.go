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

package crdloader

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/util/paths"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/ghodss/yaml" //nolint:depguard
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CrdLoader struct {
	kubeClient client.Client
}

func New(kubeClient client.Client) *CrdLoader {
	return &CrdLoader{
		kubeClient: kubeClient,
	}
}

// Find a matching CRD in the API server
func (l *CrdLoader) GetCRDForKind(kind string) (*apiextensions.CustomResourceDefinition, error) {
	return l.GetCRD("", "", kind)
}

// Find a matching CRD in the API server
func (l *CrdLoader) GetCRDForGVK(gvk schema.GroupVersionKind) (*apiextensions.CustomResourceDefinition, error) {
	return l.GetCRD(gvk.Group, gvk.Version, gvk.Kind)
}

// Find a matching CRD in the API server, the group and version parameters are optional
func (l *CrdLoader) GetCRD(group, version, kind string) (*apiextensions.CustomResourceDefinition, error) {
	if kind == "" {
		return nil, fmt.Errorf("invalid argument: 'kind' must contain a value")
	}
	if group == "" || version == "" {
		return l.getCRDViaList(group, version, kind)
	}
	return l.getCRDViaGet(group, version, kind)
}

func (l *CrdLoader) getCRDViaList(group, version, kind string) (*apiextensions.CustomResourceDefinition, error) {
	listOptions := client.ListOptions{
		Raw: &metav1.ListOptions{},
	}
	crds := make([]apiextensions.CustomResourceDefinition, 0)
	for ok := true; ok; ok = listOptions.Raw.Continue != "" {
		var list apiextensions.CustomResourceDefinitionList
		if err := l.kubeClient.List(context.TODO(), &list, &listOptions); err != nil {
			return nil, fmt.Errorf("error listing CRDs for GVK %v: %w", formatGVK(group, version, kind), err)
		}
		crds = append(crds, list.Items...)
		listOptions.Raw.Continue = list.Continue
	}
	return getMatchingCRD(group, version, kind, crds)
}

func (l *CrdLoader) getCRDViaGet(group, version, kind string) (*apiextensions.CustomResourceDefinition, error) {
	lowercasePluralKind := strings.ToLower(text.Pluralize(kind))
	var crd apiextensions.CustomResourceDefinition
	nn := types.NamespacedName{Name: fmt.Sprintf("%v.%v", lowercasePluralKind, group)}
	if err := l.kubeClient.Get(context.TODO(), nn, &crd); err != nil {
		return nil, fmt.Errorf("error getting CRD for GVK %v: %w", formatGVK(group, version, kind), err)
	}
	return &crd, nil
}

// Find a matching CRD from disk
func GetCRDForKind(kind string) (*apiextensions.CustomResourceDefinition, error) {
	return GetCRD("", "", kind)
}

// Find a matching CRD from disk
func GetCRDForGVK(gvk schema.GroupVersionKind) (*apiextensions.CustomResourceDefinition, error) {
	return GetCRD(gvk.Group, gvk.Version, gvk.Kind)
}

// Find a matching CRD from disk, the group and version parameters are optional
func GetCRD(group, version, kind string) (*apiextensions.CustomResourceDefinition, error) {
	crds, err := LoadCRDs()
	if err != nil {
		return nil, fmt.Errorf("error loading CRDs: %w", err)
	}
	return getMatchingCRD(group, version, kind, crds)
}

func getMatchingCRD(group, version, kind string, crds []apiextensions.CustomResourceDefinition) (*apiextensions.CustomResourceDefinition, error) {
	var match *apiextensions.CustomResourceDefinition
	for _, crd := range crds {
		if isMatch(group, version, kind, crd) {
			if match == nil {
				crd := crd
				match = &crd
			} else {
				return nil, fmt.Errorf("ambiguous result: multiple CRDs match GVK parameter of %v", formatGVK(group, version, kind))
			}
		}
	}
	if match == nil {
		return nil, fmt.Errorf("no CRD matches GVK parameter of %v", formatGVK(group, version, kind))
	}
	return match, nil
}

func isMatch(group, version, kind string, crd apiextensions.CustomResourceDefinition) bool {
	if crd.Spec.Names.Kind != kind {
		return false
	}
	if group != "" {
		if crd.Spec.Group != group {
			return false
		}
	}

	if version != "" {
		foundVersion := false
		for i := range crd.Spec.Versions {
			if crd.Spec.Versions[i].Name == version {
				foundVersion = true
			}
		}
		if !foundVersion {
			return false
		}
	}
	return true
}

func formatGVK(group, version, kind string) string {
	if group == "" {
		group = "nil"
	}
	if version == "" {
		version = "nil"
	}
	return fmt.Sprintf("{%v, %v, %v}", group, version, kind)
}

// LoadCRDs returns the list of CRDs Config Connector manages that map to GCP
// resources.
func LoadCRDs() ([]apiextensions.CustomResourceDefinition, error) {
	return loadCRDs(repo.GetCRDsPath())
}

// LoadAllCRDs returns the list of all the CRDs Config Connector manages,
// including the ones map to GCP resources, and the ones used to manage the
// Config Connector operator.
func LoadAllCRDs() ([]apiextensions.CustomResourceDefinition, error) {
	results := make([]apiextensions.CustomResourceDefinition, 0)
	loadFuncs := []func() ([]apiextensions.CustomResourceDefinition, error){LoadCRDs, loadOperatorCRDs}
	for _, loadFunc := range loadFuncs {
		crds, err := loadFunc()
		if err != nil {
			return nil, err
		}
		results = append(results, crds...)
	}
	return results, nil
}

func loadOperatorCRDs() ([]apiextensions.CustomResourceDefinition, error) {
	allCRDs := make([]apiextensions.CustomResourceDefinition, 0)
	for _, p := range paths.GetOperatorCRDsPaths() {
		crds, err := loadCRDs(p)
		if err != nil {
			return nil, err
		}
		allCRDs = append(allCRDs, crds...)
	}
	return allCRDs, nil
}

func loadCRDs(crdsRoot string) ([]apiextensions.CustomResourceDefinition, error) {
	files, err := ioutil.ReadDir(crdsRoot)
	if err != nil {
		return nil, fmt.Errorf("error listing directory '%v': %w", crdsRoot, err)
	}
	results := make([]apiextensions.CustomResourceDefinition, 0)
	for _, crdFile := range files {
		crd, err := FileToCRD(path.Join(crdsRoot, crdFile.Name()))
		if err != nil {
			return nil, err
		}
		results = append(results, *crd)
	}
	return results, nil
}

func FileToCRD(fileName string) (*apiextensions.CustomResourceDefinition, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file '%v': %w", fileName, err)
	}
	var crd apiextensions.CustomResourceDefinition
	err = yaml.Unmarshal(bytes, &crd)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling '%v' to CRD: %w", fileName, err)
	}
	return &crd, nil
}
