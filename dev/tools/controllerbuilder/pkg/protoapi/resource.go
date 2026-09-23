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
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ParentStyle names the shape of a resource's parent, as declared by the
// google.api.resource pattern.
type ParentStyle string

const (
	// ParentProjectLocation is "projects/*/locations/*" - the most common shape,
	// and the one the identity template assumes unconditionally today.
	ParentProjectLocation ParentStyle = "project_location"
	// ParentProject is "projects/*".
	ParentProject ParentStyle = "project"
	// ParentOrganization is "organizations/*".
	ParentOrganization ParentStyle = "organization"
	// ParentFolder is "folders/*".
	ParentFolder ParentStyle = "folder"
	// ParentMulti indicates a resource that supports multiple parent hierarchies (e.g. project, folder, or organization).
	ParentMulti ParentStyle = "multi"
	// ParentOther is anything else: deeper nesting, or a parent we do not model.
	// ParentOther indicates non-standard parent hierarchy patterns requiring custom handling.
	ParentOther ParentStyle = "other"
	// ParentUnknown means the proto declared no pattern for us to read.
	ParentUnknown ParentStyle = "unknown"
)

// ResourceMetadata holds resource naming and hierarchy information extracted
// from google.api.resource proto annotations.
type ResourceMetadata struct {
	// Patterns holds all declared resource name patterns from the descriptor.
	Patterns []string
	// Pattern is the primary declared resource name pattern, e.g.
	// "projects/{project}/locations/{location}/lbTrafficExtensions/{extension}".
	Pattern string
	// Plural is the declared plural, when set. Only about a quarter of annotated
	// messages set it, so Collection is usually the more reliable source.
	Plural string
	// Collection is the segment naming this resource's collection, taken from the
	// pattern, e.g. "lbTrafficExtensions". This is the value that belongs in a
	// resource name, preserving the API's own casing.
	Collection string
	// ParentPath is the parent's literal collection segments, joined, e.g.
	// "projects/locations". ParentStyle collapses the uncommon shapes into
	// "other", so this is what tells a human triaging one what it actually is.
	ParentPath string
	// ParentStyle classifies ParentPath into a shape the templates can render.
	ParentStyle ParentStyle
}

// GetResourceMetadata extracts google.api.resource metadata from a message descriptor.
// It returns nil if the message does not carry a resource descriptor with a pattern.
func GetResourceMetadata(msg protoreflect.MessageDescriptor) *ResourceMetadata {
	if msg == nil {
		return nil
	}
	v := proto.GetExtension(msg.Options(), annotations.E_Resource)
	rd, _ := v.(*annotations.ResourceDescriptor)
	if rd == nil {
		return nil
	}
	patterns := rd.GetPattern()
	if len(patterns) == 0 {
		return nil
	}

	md := &ResourceMetadata{
		Patterns: patterns,
		Pattern:  patterns[0],
		Plural:   rd.GetPlural(),
	}
	md.Collection, md.ParentPath = splitPattern(md.Pattern)
	if md.Collection == "" && md.Plural != "" {
		md.Collection = md.Plural
	}
	md.ParentStyle = classifyParent(md.ParentPath)
	if len(patterns) > 1 {
		for _, p := range patterns[1:] {
			_, pPath := splitPattern(p)
			if classifyParent(pPath) != md.ParentStyle {
				md.ParentStyle = ParentMulti
				break
			}
		}
	}
	return md
}

// splitPattern separates a resource name pattern into the resource's own
// collection segment and its parent's collection segments, joined back into a
// path.
//
// "projects/{project}/locations/{location}/foos/{foo}"
//
//	-> collection "foos", parent "projects/locations"
func splitPattern(pattern string) (collection string, parentPath string) {
	segs := strings.Split(pattern, "/")

	// Literal segments alternate with {placeholders}. Anything that is not a
	// placeholder is a collection name.
	var literals []string
	for _, s := range segs {
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "{") {
			continue
		}
		literals = append(literals, s)
	}
	if len(literals) == 0 {
		return "", ""
	}

	// A pattern may end in a literal with no trailing placeholder (e.g.
	// "projects/{project}/locations"); in that case there is no distinct
	// collection for the resource itself.
	last := segs[len(segs)-1]
	if !strings.HasPrefix(last, "{") {
		return "", strings.Join(literals, "/")
	}
	return literals[len(literals)-1], strings.Join(literals[:len(literals)-1], "/")
}

func classifyParent(parentPath string) ParentStyle {
	switch parentPath {
	case "":
		return ParentUnknown
	case "projects/locations":
		return ParentProjectLocation
	case "projects":
		return ParentProject
	case "organizations":
		return ParentOrganization
	case "folders":
		return ParentFolder
	default:
		return ParentOther
	}
}
