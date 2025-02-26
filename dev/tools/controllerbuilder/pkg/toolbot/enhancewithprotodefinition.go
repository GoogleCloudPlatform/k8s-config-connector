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
	"os"
	"path/filepath"
	"strings"

	"k8s.io/klog/v2"
)

type protoService struct {
	FilePath   string
	Definition []string
}

type protoMessage struct {
	FilePath   string
	Definition []string
}

// EnhanceWithProtoDefinition is an enhancer that adds the definition of a proto message or service to the data point.
type EnhanceWithProtoDefinition struct {
	protoDirectory string
	messages       map[string]*protoMessage
	services       map[string]*protoService
}

// NewEnhanceWithProtoDefinition creates a new EnhanceWithProtoDefinition.
func NewEnhanceWithProtoDefinition(protoDirectory string) (*EnhanceWithProtoDefinition, error) {
	x := &EnhanceWithProtoDefinition{
		protoDirectory: protoDirectory,
		messages:       make(map[string]*protoMessage),
		services:       make(map[string]*protoService),
	}
	if err := x.findProtoMessages(); err != nil {
		return nil, err
	}
	if err := x.findProtoServices(); err != nil {
		return nil, err
	}
	return x, nil
}

var _ Enhancer = &EnhanceWithProtoDefinition{}

// EnhanceDataPoint enhances the data point by adding the definition of the proto message or service.
func (x *EnhanceWithProtoDefinition) EnhanceDataPoint(ctx context.Context, p *DataPoint) error {
	service := p.Input["proto.service"]
	if service != "" {
		protoService := x.services[service]
		if protoService != nil {
			p.SetInput("proto.service.definition", strings.Join(protoService.Definition, "\n"))
		} else {
			klog.Infof("unable to find proto service %q", service)
			return nil
			//return fmt.Errorf("unable to find proto service %q", service)
		}
	}

	message := p.Input["proto.message"]
	if message != "" {
		protoMessage := x.messages[message]
		if protoMessage != nil {
			p.SetInput("proto.message.definition", strings.Join(protoMessage.Definition, "\n"))
		} else {
			klog.Infof("unable to find proto message %q", message)
		}
	}

	return nil
}

// findProtoMessages finds all the proto services in the proto directory.
func (x *EnhanceWithProtoDefinition) findProtoServices() error {
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
				serviceName := packageName + "." + tokens[1]

				service := &protoService{FilePath: p}
				indent := 0
				for {
					service.Definition = append(service.Definition, line)
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
				if _, found := x.services[serviceName]; found {
					return fmt.Errorf("found duplication definition for service %q", serviceName)
				}
				x.services[serviceName] = service
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory tree: %w", err)
	}

	return nil
}

// findProtoMessages finds all the proto messages in the proto directory.
func (x *EnhanceWithProtoDefinition) findProtoMessages() error {
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

			if len(tokens) >= 2 && tokens[0] == "message" {
				messageName := packageName + "." + tokens[1]

				message := &protoMessage{FilePath: p}
				indent := 0
				for {
					message.Definition = append(message.Definition, line)
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
				x.messages[messageName] = message
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("walking directory tree: %w", err)
	}

	return nil
}
