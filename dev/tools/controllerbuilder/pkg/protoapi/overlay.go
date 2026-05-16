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

package protoapi

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	pkgRegex      = regexp.MustCompile(`^\s*package\s+([^;]+);`)
	msgRegex      = regexp.MustCompile(`^\s*message\s+([^{\s]+)\s*\{`)
	endMsgRegex   = regexp.MustCompile(`^\s*\}`)
	fieldRegex    = regexp.MustCompile(`^\s*(?:repeated\s+)?(?:optional\s+)?(?:required\s+)?[^\s]+\s+([^\s=]+)\s*=\s*\d+\s*\[(.*)\];`)
	behaviorRegex = regexp.MustCompile(`\(google\.api\.field_behavior\)\s*=\s*([A-Z_]+)`)
)

func ApplyOverlay(fds *descriptorpb.FileDescriptorSet, overlayPath string) error {
	f, err := os.Open(overlayPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return applyOverlayFromReader(fds, f)
}

func applyOverlayFromReader(fds *descriptorpb.FileDescriptorSet, r io.Reader) error {
	scanner := bufio.NewScanner(r)

	var currentPackage string
	var messageStack []string

	fieldBehaviors := make(map[string][]annotations.FieldBehavior)

	for scanner.Scan() {
		line := scanner.Text()
		
		if pkgMatch := pkgRegex.FindStringSubmatch(line); pkgMatch != nil {
			currentPackage = pkgMatch[1]
			continue
		}

		if msgMatch := msgRegex.FindStringSubmatch(line); msgMatch != nil {
			messageStack = append(messageStack, msgMatch[1])
			continue
		}

		if endMsgRegex.MatchString(line) {
			if len(messageStack) > 0 {
				messageStack = messageStack[:len(messageStack)-1]
			}
			continue
		}

		if fieldMatch := fieldRegex.FindStringSubmatch(line); fieldMatch != nil {
			fieldName := fieldMatch[1]
			opts := fieldMatch[2]

			if behaviorMatch := behaviorRegex.FindStringSubmatch(opts); behaviorMatch != nil {
				behaviorStr := behaviorMatch[1]
				var behavior annotations.FieldBehavior
				switch behaviorStr {
				case "OUTPUT_ONLY":
					behavior = annotations.FieldBehavior_OUTPUT_ONLY
				case "REQUIRED":
					behavior = annotations.FieldBehavior_REQUIRED
				case "OPTIONAL":
					behavior = annotations.FieldBehavior_OPTIONAL
				case "IMMUTABLE":
					behavior = annotations.FieldBehavior_IMMUTABLE
				case "INPUT_ONLY":
					behavior = annotations.FieldBehavior_INPUT_ONLY
				default:
					return fmt.Errorf("unknown field behavior: %q", behaviorStr)
				}
				
				fullName := currentPackage
				for _, msg := range messageStack {
					if fullName != "" {
						fullName += "."
					}
					fullName += msg
				}
				if fullName != "" {
					fullName += "."
				}
				fullName += fieldName

				fieldBehaviors[fullName] = append(fieldBehaviors[fullName], behavior)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if len(fieldBehaviors) == 0 {
		return nil
	}

	// Apply modifications to fds
	for _, file := range fds.File {
		applyToMessages(file.GetPackage(), file.MessageType, fieldBehaviors)
	}

	return nil
}

func applyToMessages(pkg string, msgs []*descriptorpb.DescriptorProto, fieldBehaviors map[string][]annotations.FieldBehavior) {
	for _, msg := range msgs {
		msgName := msg.GetName()
		fullMsgName := msgName
		if pkg != "" {
			fullMsgName = pkg + "." + msgName
		}
		
		for _, field := range msg.Field {
			fullName := fullMsgName + "." + field.GetName()
			if behaviors, ok := fieldBehaviors[fullName]; ok {
				if field.Options == nil {
					field.Options = &descriptorpb.FieldOptions{}
				}
				// Append behaviors using proto extension.
                // We overwrite for now, or append if exists?
                // Usually it's better to append if already exists, or just set it.
                // proto.GetExtension could be checked, but for simplicity we can just set.
                // wait, if we just set it, it overrides any existing behaviors which is fine.
				proto.SetExtension(field.Options, annotations.E_FieldBehavior, behaviors)
			}
		}
		
		if len(msg.NestedType) > 0 {
			applyToMessages(fullMsgName, msg.NestedType, fieldBehaviors)
		}
	}
}
