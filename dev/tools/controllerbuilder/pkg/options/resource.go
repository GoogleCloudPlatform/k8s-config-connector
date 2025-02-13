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

package options

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type Resource struct {
	Kind              string
	ProtoName         string
	SkipScaffoldFiles bool
}

var _ pflag.Value = &Resource{}

func (r *Resource) Type() string { return "resource" }

func (r *Resource) String() string {
	return fmt.Sprintf("%s:%s", r.Kind, r.ProtoName)
}

func (r *Resource) Set(s string) error {
	tokens := strings.Split(s, ":")
	if len(tokens) != 2 || tokens[0] == "" || tokens[1] == "" {
		return fmt.Errorf("expected [KRMKind]:[ProtoResourceName], got %q", s)
	}
	r.Kind = tokens[0]
	r.ProtoName = tokens[1]
	return nil
}

type ResourceList []Resource

var _ pflag.Value = &ResourceList{}

func (r *ResourceList) Type() string { return "resources" }

func (r *ResourceList) String() string {
	var sb strings.Builder
	for _, res := range *r {
		fmt.Fprintf(&sb, "%s:%s", res.Kind, res.ProtoName)
	}
	return sb.String()
}

func (r *ResourceList) Set(s string) error {
	tokens := strings.Split(s, ":")
	if len(tokens) != 2 || tokens[0] == "" || tokens[1] == "" {
		return fmt.Errorf("expected [KRMKind]:[ProtoResourceName], got %q", s)
	}
	*r = append(*r, Resource{
		Kind:      tokens[0],
		ProtoName: tokens[1],
	})
	return nil
}
