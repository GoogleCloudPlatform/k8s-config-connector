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

package toolbot

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/annotations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"k8s.io/klog/v2"
)

// ExtractToolMarkers extracts tool markers from source code.
type ExtractToolMarkers struct {
}

var _ Extractor = &ExtractToolMarkers{}

// Extract extracts tool markers from source code.
func (x *ExtractToolMarkers) Extract(ctx context.Context, description string, src []byte, filters ...Filter) ([]*DataPoint, error) {
	var dataPoints []*DataPoint

	// Find file-scoped DataPoints
	{
		markers := []string{"+tool:"}
		annotations, err := annotations.FindFileAnnotations(src, markers)
		if err != nil {
			return nil, err
		}
		for _, annotation := range annotations {
			toolName := strings.TrimPrefix(annotation.Key, "+tool:")
			dataPoint := &DataPoint{
				Description: description,
				Type:        toolName,
				Output:      string(src),
			}

			for k, values := range annotation.Attributes {
				for _, v := range values {
					dataPoint.SetInput(k, v)
				}
			}
			shouldAdd := true
			for _, filter := range filters {
				if !filter(dataPoint) {
					shouldAdd = false
				}
			}
			if shouldAdd {
				dataPoints = append(dataPoints, dataPoint)
			}
		}
	}

	r := bytes.NewReader(src)
	br := bufio.NewReader(r)

	for {
		rawLine, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("scanning code: %w", err)
		}
		line := strings.TrimSpace(rawLine)
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimPrefix(line, "//")
			comment = strings.TrimSpace(comment)

			if proto, ok := codegen.GetProtoMessageFromAnnotation(comment); ok {
				klog.V(2).Infof("found tool line %q", comment)
				toolName := "kcc-proto"
				if strings.Contains(comment, codegen.KCCProtoMessageAnnotationSpec) {
					toolName = "kcc-proto-spec"
				}
				if strings.Contains(comment, codegen.KCCProtoMessageAnnotationObservedState) {
					toolName = "kcc-proto-observedstate"
				}
				dataPoint := &DataPoint{
					Description: description,
					Type:        toolName,
				}

				dataPoint.SetInput("proto.message", proto)

				var bb bytes.Buffer
				for {
					line, err := br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return nil, fmt.Errorf("scanning code: %w", err)
					}

					bb.WriteString(line)

					s := strings.TrimSpace(line)
					if strings.HasPrefix(s, "}") {
						break
					}
				}
				dataPoint.Output = bb.String()

				shouldAdd := true
				for _, filter := range filters {
					if !filter(dataPoint) {
						shouldAdd = false
					}
				}
				if shouldAdd {
					dataPoints = append(dataPoints, dataPoint)
				}
			}

			if strings.HasPrefix(comment, "+function-gen:special-mapper") {
				klog.V(2).Infof("found tool line %q", comment)
				toolName := "+function-gen:special-mapper"
				dataPoint := &DataPoint{
					Description: description,
					Type:        toolName,
				}

				var bb bytes.Buffer

				// Include the tool directive
				bb.WriteString(rawLine)

				inHeader := true
				openBrackets := 0
				for {
					rawLine, err := br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return nil, fmt.Errorf("scanning code: %w", err)
					}

					if inHeader {
						line := strings.TrimSpace(rawLine)
						if strings.HasPrefix(line, "//") {
							toolLine := strings.TrimPrefix(line, "//")
							toolLine = strings.TrimPrefix(toolLine, " ")
							tokens := strings.SplitN(toolLine, ":", 2)
							if len(tokens) == 2 {
								dataPoint.SetInput(tokens[0], strings.TrimSpace(tokens[1]))
							} else {
								return nil, fmt.Errorf("cannot parse tool line %q", toolLine)
							}
						} else {
							inHeader = false
						}
					}

					bb.WriteString(rawLine)

					if strings.HasSuffix(line, "{") {
						openBrackets++
					}

					if strings.HasPrefix(line, "}") {
						openBrackets--
						if openBrackets == 0 {
							break
						}
					}
				}
				dataPoint.Output = bb.String()

				shouldAdd := true
				for _, filter := range filters {
					if !filter(dataPoint) {
						shouldAdd = false
					}
				}
				if shouldAdd {
					dataPoints = append(dataPoints, dataPoint)
				}
			}
		}
	}
	return dataPoints, nil
}
