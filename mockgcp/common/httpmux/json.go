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

package httpmux

import (
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/klog/v2"
)

func MarshalAsJSON(obj proto.Message) ([]byte, error) {
	return protojson.MarshalOptions{Resolver: &Resolver{}}.Marshal(obj)
}

type Resolver struct {
}

var _ protoregistry.ExtensionTypeResolver = &Resolver{}

func (r *Resolver) FindExtensionByName(message protoreflect.FullName) (protoreflect.ExtensionType, error) {
	return protoregistry.GlobalTypes.FindExtensionByName(r.remapName(message))
}

func (r *Resolver) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
	return protoregistry.GlobalTypes.FindExtensionByNumber(r.remapName(message), field)
}

var _ protoregistry.MessageTypeResolver = &Resolver{}

func (r *Resolver) FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error) {
	return protoregistry.GlobalTypes.FindMessageByName(r.remapName(message))
}

func (r *Resolver) FindMessageByURL(url string) (protoreflect.MessageType, error) {
	// Default to trying to find the message as-is.
	mt, err := protoregistry.GlobalTypes.FindMessageByURL(url)
	if err == nil {
		return mt, nil
	}
	if suffix, ok := strings.CutPrefix(url, "type.googleapis.com/google."); ok {
		s := "type.googleapis.com/mockgcp." + suffix
		mt, err := protoregistry.GlobalTypes.FindMessageByURL(s)
		if err != nil {
			klog.Warningf("FindMessageByURL(%q) failed: %v", s, err)
		} else {
			return mt, nil
		}
	}

	return nil, err
}

func (r *Resolver) remapName(name protoreflect.FullName) protoreflect.FullName {
	// Remap names with a prefix of "google."" to be "mockgcp.", so we can find them.

	s := string(name)
	if strings.HasPrefix(s, "google.") {
		s = "mockgcp." + strings.TrimPrefix(s, "google.")
		return protoreflect.FullName(s)
	}
	return name
}
