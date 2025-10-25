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

package protoapi

import (
	"fmt"
	"os"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Proto struct {
	files *protoregistry.Files
}

func LoadProto(p string) (*Proto, error) {
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading %q: %w", p, err)
	}

	fds := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(b, fds); err != nil {
		return nil, fmt.Errorf("unmarshalling %q: %w", p, err)
	}

	files, err := protodesc.NewFiles(fds)
	if err != nil {
		return nil, fmt.Errorf("building file description: %w", err)
	}

	return &Proto{
		files: files,
	}, nil
}

func (p *Proto) SortedFiles() []protoreflect.FileDescriptor {
	var sortedFiles []protoreflect.FileDescriptor
	p.files.RangeFiles(func(f protoreflect.FileDescriptor) bool {
		sortedFiles = append(sortedFiles, f)
		return true
	})
	sort.Slice(sortedFiles, func(i, j int) bool {
		return sortedFiles[i].FullName() < sortedFiles[j].FullName()
	})
	return sortedFiles
}

func (p *Proto) Files() *protoregistry.Files {
	return p.files
}

func (p *Proto) GetFileDescriptorByPackage(protoPackage string) ([]protoreflect.FileDescriptor, error) {
	var files []protoreflect.FileDescriptor
	for _, f := range p.SortedFiles() {
		if string(f.Package()) == protoPackage {
			files = append(files, f)
		}
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("could not find FileDescriptor for package %q", protoPackage)
	}
	return files, nil
}
