// Copyright 2025 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +k8s:deepcopy-gen=false
type TagBindingParent interface {
	String() string
	Prefix() string

	FromExternal(ref string) error
}

var ParentFullResourceMap = map[FullResourcePrefix]TagBindingParent{
	ProjectPrefix: &TagBindingProject{},
}

// TagBindingParent represents the full resource name prefix of a TagBinding parent.
type FullResourcePrefix string

// Predefined TagBinding parent full resource name prefixes for "Project".
var ProjectPrefix FullResourcePrefix = "//cloudresourcemanager.googleapis.com/projects"

var _ TagBindingParent = &TagBindingProject{}

// TagBindingProject represents a TagBinding parent of the form:
// `//cloudresourcemanager.googleapis.com/projects/{{projectNumber}}`
type TagBindingProject struct {
	ProjectNumber string
}

func (p *TagBindingProject) String() string {
	return fmt.Sprintf("%s/%s", ProjectPrefix, p.ProjectNumber)
}

func (p *TagBindingProject) Prefix() string {
	return fmt.Sprintf("%s", ProjectPrefix)
}

func (p *TagBindingProject) FromExternal(ref string) error {
	var found bool
	p.ProjectNumber, found = strings.CutPrefix(ref, p.Prefix()+"/")
	if !found {
		return fmt.Errorf("format of TagBindingProject missing prefix %q. got=%q", p.Prefix(), ref)
	}
	if p.ProjectNumber == "" {
		return fmt.Errorf("projectID was empty in TagBinding parent external=%q", ref)
	}
	return nil
}

var _ refsv1beta1.ExternalNormalizer = &ParentRef{}

func (p *ParentRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if p.External != "" {
		return p.External, nil
	}

	projectNN := types.NamespacedName{
		Name:      p.Name,
		Namespace: p.Namespace,
	}
	projectNumber, err := refsv1beta1.ResolveProjectNumber(ctx, reader, projectNN)
	if err != nil {
		return "", err
	}
	p.External = fmt.Sprintf("%s/%s", ProjectPrefix, projectNumber)
	return p.External, nil
}
