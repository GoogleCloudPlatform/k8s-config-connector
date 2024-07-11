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
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

type CSVExporter struct {
	protoDirectory string
	toolEntries    []*toolEntry
}

func NewCSVExporter(protoDirectory string) (*CSVExporter, error) {
	x := &CSVExporter{
		protoDirectory: protoDirectory,
	}

	return x, nil
}

type toolEntry struct {
	Name     string
	FilePath string
	Values   map[string]string
}

func (x *CSVExporter) visitGoFile(ctx context.Context, p string) error {

	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading file %q: %w", p, err)
	}
	r := bytes.NewReader(b)
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("scanning file %q: %w", p, err)
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimPrefix(line, "//")
			comment = strings.TrimSpace(comment)
			if strings.HasPrefix(comment, "+tool:") {
				klog.V(2).Infof("found tool line %q", comment)
				toolName := strings.TrimPrefix(comment, "+tool:")
				toolEntry := &toolEntry{
					Name:     toolName,
					FilePath: p,
					Values:   make(map[string]string),
				}
				for {
					line, err := br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return fmt.Errorf("scanning file %q: %w", p, err)
					}
					line = strings.TrimSpace(line)
					if !strings.HasPrefix(line, "//") {
						break
					}
					toolLine := strings.TrimPrefix(line, "//")
					toolLine = strings.TrimPrefix(toolLine, " ")

					tokens := strings.SplitN(toolLine, ":", 2)
					if len(tokens) == 2 {
						toolEntry.Values[tokens[0]] = strings.TrimSpace(tokens[1])
					} else {
						return fmt.Errorf("cannot parse tool line %q", toolLine)
					}
				}
				x.toolEntries = append(x.toolEntries, toolEntry)
			}
		}
	}
	return nil
}

func (x *CSVExporter) VisitCode(ctx context.Context, srcDir string) error {
	if err := filepath.WalkDir(srcDir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch filepath.Ext(p) {
		case ".go":
			// OK
		default:
			return nil
		}
		// klog.Infof("%v", p)
		if err := x.visitGoFile(ctx, p); err != nil {
			return fmt.Errorf("processing file %q: %w", p, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory tree: %w", err)
	}

	return nil
}

func (x *CSVExporter) WriteCSV(ctx context.Context, outputDir string) error {
	log := klog.FromContext(ctx)

	toolsByName := make(map[string][]*toolEntry)
	for _, toolEntry := range x.toolEntries {
		toolsByName[toolEntry.Name] = append(toolsByName[toolEntry.Name], toolEntry)
	}

	for toolName, tools := range toolsByName {
		var bb bytes.Buffer
		csvFile := csv.NewWriter(&bb)

		columnSet := sets.NewString()
		columnSet.Insert("out")

		for _, tool := range tools {
			for k := range tool.Values {
				columnSet.Insert("in." + k)
			}
		}

		if columnSet.Has("in.proto.service") {
			columnSet.Insert("in!proto.service.definition")
		}

		columns := columnSet.List()

		csvFile.Write(columns)
		csvFile.Flush()

		for _, tool := range tools {
			row := make([]string, len(columns))
			for i, column := range columns {
				switch column {
				case "out":
					b, err := os.ReadFile(tool.FilePath)
					if err != nil {
						return fmt.Errorf("reading file %q: %w", tool.FilePath, err)
					}
					row[i] = string(b)

				case "in!proto.service.definition":
					service := tool.Values["proto.service"]
					protoService, err := x.getProtoForService(ctx, service)
					if err != nil {
						return fmt.Errorf("getting proto for service %q: %w", service, err)
					}
					row[i] = strings.Join(protoService.Definition, "\n")
				default:
					if strings.HasPrefix(column, "in.") {
						row[i] = tool.Values[strings.TrimPrefix(column, "in.")]
					} else {
						return fmt.Errorf("unknown column %q", column)
					}
				}
			}
			csvFile.Write(row)
			csvFile.Flush()
		}

		if err := csvFile.Error(); err != nil {
			return fmt.Errorf("writing to csv: %w", err)
		}

		outFilePath := filepath.Join(outputDir, toolName+".csv")
		log.Info("writing CSV", "path", outFilePath)
		if err := os.WriteFile(outFilePath, bb.Bytes(), 0644); err != nil {
			return fmt.Errorf("writing to file %q: %w", outFilePath, err)
		}
	}

	return nil
}

type protoService struct {
	FilePath   string
	Definition []string
}

func (x *CSVExporter) getProtoForService(ctx context.Context, serviceName string) (*protoService, error) {
	var matches []*protoService
	if err := filepath.WalkDir(x.protoDirectory, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch filepath.Ext(p) {
		case ".proto":
			// OK
		default:
			return nil
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
		}
		r := bytes.NewReader(b)
		br := bufio.NewReader(r)

		packageName := ""

		for {
			line, err := br.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("scanning file %q: %w", p, err)
			}
			line = strings.TrimSuffix(line, "\n")

			tokens := strings.Fields(line)

			if len(tokens) >= 2 && tokens[0] == "package" {
				packageName = strings.TrimSuffix(tokens[1], ";")
			}

			if len(tokens) >= 2 && tokens[0] == "service" {
				found := packageName + "." + tokens[1]

				if found != serviceName {
					continue
				}

				match := &protoService{FilePath: p}
				indent := 0
				for {
					match.Definition = append(match.Definition, line)
					for _, r := range line {
						if r == '{' {
							indent++
						}
						if r == '}' {
							indent--
						}
					}
					if indent == 0 {
						break
					}
					line, err = br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return fmt.Errorf("scanning file %q: %w", p, err)
					}
					line = strings.TrimSuffix(line, "\n")
				}
				matches = append(matches, match)
			}
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("walking directory tree: %w", err)
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("service %q not found", serviceName)
	}
	if len(matches) > 1 {
		return nil, fmt.Errorf("found multiple services with name %q", serviceName)
	}
	return matches[0], nil
}
