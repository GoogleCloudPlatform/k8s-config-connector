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

package typeupdater

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog"
)

type SyncProtoPackageOptions struct {
	ServiceName     string
	APIVersion      string
	ProtoSourcePath string
	APIDirectory    string
	GoPackagePath   string
	LegacyMode      bool
}

type ProtoPackageSyncer struct {
	opts *SyncProtoPackageOptions

	// holds info about Go structs in existing types files, including both generated and manually edited structs.
	// key is the go struct name
	existingGoMessages map[string]messageInfo
	api                *protoapi.Proto // Store the loaded proto API
}

func NewProtoPackageSyncer(opts *SyncProtoPackageOptions) *ProtoPackageSyncer {
	return &ProtoPackageSyncer{
		opts:               opts,
		existingGoMessages: make(map[string]messageInfo),
	}
}

func (s *ProtoPackageSyncer) Run() error {
	// 1. parse the existing go types
	if err := s.parseExistingTypes(); err != nil {
		return err
	}

	// 2. load the proto package
	if err := s.loadProtoPackage(); err != nil {
		return err
	}

	// 3. create the update plans
	plans, err := s.createFieldUpdatePlans()
	if err != nil {
		return fmt.Errorf("creating update plans: %w", err)
	}

	// printUpdatePlans(plans)

	// 4. apply the update plans to update the existing types
	for _, plan := range plans {
		if err := s.applyFieldUpdatePlan(plan); err != nil {
			return fmt.Errorf("applying update plan for field %s in struct %s: %w",
				plan.fieldName, plan.structName, err)
		}
	}

	return nil
}

func (s *ProtoPackageSyncer) parseExistingTypes() error {
	dir, err := typeFilePath(s.opts.APIDirectory, s.opts.APIVersion)
	if err != nil {
		return fmt.Errorf("getting API directory for %q: %w", s.opts.APIVersion, err)
	}

	klog.Infof("Parsing existing types from %q", dir)
	messages, err := extractMessageInfoFromGoFiles(dir)
	if err != nil {
		return err
	}

	s.existingGoMessages = messages
	return nil
}

// typeFilePath returns the path to the types.go file for the given API version
func typeFilePath(apiBaseDir, gv string) (string, error) {
	groupVersion, err := schema.ParseGroupVersion(gv)
	if err != nil {
		return "", fmt.Errorf("parsing APIVersion %q: %w", gv, err)
	}

	goPackagePath := strings.TrimSuffix(groupVersion.Group, ".cnrm.cloud.google.com") + "/" + groupVersion.Version
	packageTokens := strings.Split(goPackagePath, ".")
	return filepath.Join(append([]string{apiBaseDir}, packageTokens...)...), nil
}

func (s *ProtoPackageSyncer) createFieldUpdatePlans() ([]FieldUpdatePlan, error) {
	var plans []FieldUpdatePlan

	// for each existing Go message that has a corresponding proto message
	for goTypeName, msgInfo := range s.existingGoMessages {
		if msgInfo.IsVirtual {
			klog.Infof("Skipping virtual type %s", goTypeName)
			continue
		}

		// find corresponding proto message
		desc, err := s.api.Files().FindDescriptorByName(protoreflect.FullName(msgInfo.ProtoName))
		if err != nil && err != protoregistry.NotFound {
			return nil, fmt.Errorf("finding proto message %s: %w", msgInfo.ProtoName, err)
		}
		if desc == nil {
			klog.Warningf("No proto message found for %s", msgInfo.ProtoName)
			continue
		}
		msgDesc, ok := desc.(protoreflect.MessageDescriptor)
		if !ok {
			return nil, fmt.Errorf("unexpected descriptor type for %s: %T", msgInfo.ProtoName, desc)
		}

		// for each field in the message, create update plan based on exsiting go types and the matching proto field
		for fieldName, fieldInfo := range msgInfo.Fields {
			plan, err := s.createFieldUpdatePlan(msgInfo, fieldInfo, msgDesc)
			if err != nil {
				return nil, fmt.Errorf("creating plan for field %s: %w", fieldName, err)
			}
			if plan != nil {
				plans = append(plans, *plan)
			}
		}
	}

	return plans, nil
}

func (s *ProtoPackageSyncer) loadProtoPackage() error {
	api, err := protoapi.LoadProto(s.opts.ProtoSourcePath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}
	s.api = api
	return nil
}
