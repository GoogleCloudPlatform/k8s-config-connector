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

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/tools/gapic/pkg/openapi"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/tools/gapic/pkg/protogen"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	klog.InitFlags(nil)

	protoVersion := 3
	flag.IntVar(&protoVersion, "proto-version", protoVersion, "use proto version (2 or 3)")
	protoPackage := ""
	flag.StringVar(&protoPackage, "proto-package", protoPackage, "protobuf package to generate")
	flag.Parse()

	if protoPackage == "" {
		return fmt.Errorf("must specify --proto-package")
	}

	p := flag.Args()[0]
	b, err := os.ReadFile(p)
	if err != nil {
		return fmt.Errorf("reading %q: %w", p, err)
	}
	doc := &openapi.Document{}
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(doc); err != nil {
		return fmt.Errorf("parsing json %q (with DisallowUnknownFields): %w", p, err)
	}

	c := protogen.NewOpenAPIConverter(protoPackage, doc)
	fileDescriptor, err := c.Convert(ctx)
	if err != nil {
		return fmt.Errorf("convert failed: %w", err)
	}

	filesDescriptorSet := &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{fileDescriptor},
	}

	// for k := range c.imports {
	// 	file, err := resolveImport(k)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	files, err := protodesc.FileOptions{AllowUnresolvable: true}.NewFiles(filesDescriptorSet)
	if err != nil {
		return fmt.Errorf("protodesc.NewFiles failed: %w", err)
	}

	files.RangeFiles(func(file protoreflect.FileDescriptor) bool {
		pw := protogen.NewProtoWriter(os.Stdout)
		pw.SetComments(&c.Comments)
		pw.SetProtoVersion(protoVersion)
		pw.WriteFile(file)
		if err := pw.Error(); err != nil {
			klog.Fatalf("error rendering proto: %v", err)
		}

		return true
	})

	return nil
}

func PtrTo[T any](val T) *T {
	return &val
}
