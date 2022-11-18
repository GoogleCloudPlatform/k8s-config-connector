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

package servicemappingloader

import (
	"fmt"
	"io"
	"path"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/scripts/resource-autogen/allowlist"
	generatedembed "github.com/GoogleCloudPlatform/k8s-config-connector/scripts/resource-autogen/servicemapping/embed/generated"

	"github.com/ghodss/yaml"
)

func GetServiceMappingMap() (map[string]v1alpha1.ServiceMapping, error) {
	generatedSMMap, err := getGeneratedSMMap()
	if err != nil {
		return nil, fmt.Errorf("error getting all the generated ServiceMapping map: %w", err)
	}
	return getAllowlistedSMMap(generatedSMMap)
}

func getGeneratedSMMap() (map[string]v1alpha1.ServiceMapping, error) {
	baseDirName := "/"
	generatedSMDir, err := generatedembed.Assets.Open(baseDirName)
	if err != nil {
		return nil, fmt.Errorf("error reading generated files in ServiceMapping directory: %v", err)
	}
	defer generatedSMDir.Close()
	generatedFiles, err := generatedSMDir.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("error reading generated files in ServiceMapping directory: %v", err)
	}

	serviceMappings := make(map[string]v1alpha1.ServiceMapping)
	for _, file := range generatedFiles {
		smPath := path.Join(baseDirName, file.Name())
		sm, err := fileToServiceMapping(smPath)
		if err != nil {
			return nil, err
		}
		serviceMappings[sm.Name] = *sm
	}
	return serviceMappings, nil
}

func getAllowlistedSMMap(generatedSMMap map[string]v1alpha1.ServiceMapping) (map[string]v1alpha1.ServiceMapping, error) {
	autoGenAllowlist, err := allowlist.LoadAutoGenAllowList()
	if err != nil {
		return nil, fmt.Errorf("error loading allowlist for autogen resources: %w", err)
	}

	allowlistedSMMap := make(map[string]v1alpha1.ServiceMapping)
	for _, sm := range generatedSMMap {
		if !autoGenAllowlist.HasService(sm.Spec.Name) {
			continue
		}

		allowlistedSM := deepcopy.DeepCopy(sm).(v1alpha1.ServiceMapping)
		rcList := []v1alpha1.ResourceConfig{}
		for _, rc := range sm.Spec.Resources {
			if !autoGenAllowlist.HasTFTypeInService(sm.Spec.Name, rc.Name) {
				continue
			}

			allowlistedRC := deepcopy.DeepCopy(rc).(v1alpha1.ResourceConfig)
			rcList = append(rcList, allowlistedRC)
		}
		allowlistedSM.Spec.Resources = rcList
		allowlistedSMMap[sm.Name] = allowlistedSM
	}
	return allowlistedSMMap, nil
}

func fileToServiceMapping(filePath string) (*v1alpha1.ServiceMapping, error) {
	file, err := generatedembed.Assets.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("error opening file '%v': %v", filePath, err)
	}
	defer file.Close()

	sm, err := readerToServiceMapping(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file '%v' to service mapping: %v", filePath, err)
	}
	return sm, nil
}

func readerToServiceMapping(r io.Reader) (*v1alpha1.ServiceMapping, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	var sm v1alpha1.ServiceMapping
	if err := yaml.Unmarshal(bytes, &sm); err != nil {
		return nil, fmt.Errorf("error unmarshaling byte to service mapping: %v", err)
	}
	return &sm, nil
}
