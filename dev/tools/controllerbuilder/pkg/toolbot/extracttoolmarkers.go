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

	"k8s.io/klog/v2"
)

// ExtractToolMarkers extracts tool markers from source code.
type ExtractToolMarkers struct {
}

var _ Extractor = &ExtractToolMarkers{}

// Extract extracts tool markers from source code.
func (x *ExtractToolMarkers) Extract(ctx context.Context, description string, src []byte, filters ...Filter) ([]*DataPoint, error) {
	var dataPoints []*DataPoint

	r := bytes.NewReader(src)
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("scanning code: %w", err)
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimPrefix(line, "//")
			comment = strings.TrimSpace(comment)
			if strings.HasPrefix(comment, "+tool:") {
				klog.V(2).Infof("found tool line %q", comment)
				toolName := strings.TrimPrefix(comment, "+tool:")
				dataPoint := &DataPoint{
					Description: description,
					Type:        toolName,
					Output:      string(src),
				}

				for {
					line, err := br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return nil, fmt.Errorf("scanning code: %w", err)
					}
					line = strings.TrimSpace(line)
					if !strings.HasPrefix(line, "//") {
						break
					}
					toolLine := strings.TrimPrefix(line, "//")
					toolLine = strings.TrimPrefix(toolLine, " ")

					tokens := strings.SplitN(toolLine, ":", 2)
					if len(tokens) == 2 {
						dataPoint.SetInput(tokens[0], strings.TrimSpace(tokens[1]))
					} else {
						return nil, fmt.Errorf("cannot parse tool line %q", toolLine)
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

			if strings.HasPrefix(comment, "+kcc:proto=") {
				klog.V(2).Infof("found tool line %q", comment)
				toolName := "kcc-proto"
				dataPoint := &DataPoint{
					Description: description,
					Type:        toolName,
				}

				proto := strings.TrimPrefix(comment, "+kcc:proto=")
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
		}
	}
	return dataPoints, nil
}
