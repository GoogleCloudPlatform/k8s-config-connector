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

package refs

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

var (
	_ identity.IdentityV2 = &FolderIdentity{}
)

var FolderIdentityFormat = gcpurls.Template[FolderIdentity]("cloudresourcemanager.googleapis.com", "folders/{FolderID}")

// +k8s:deepcopy-gen=false
type FolderIdentity struct {
	FolderID string
}

func (i *FolderIdentity) String() string {
	return FolderIdentityFormat.ToString(*i)
}

func (i *FolderIdentity) FromExternal(ref string) error {
	parsed, match, err := FolderIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of Folder external=%q was not known (use %s): %w", ref, FolderIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		// Fallback for simple ID as external ref (e.g. just "123456" instead of "folders/123456")
		tokens := strings.Split(ref, "/")
		if len(tokens) == 1 {
			i.FolderID = tokens[0]
			return nil
		}
		return fmt.Errorf("format of Folder external=%q was not known (use %s)", ref, FolderIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *FolderIdentity) Host() string {
	return FolderIdentityFormat.Host()
}
