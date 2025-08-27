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
	"fmt"
	"strings"
)

type FolderIdentity struct {
	ResourceID string
}

func (i *FolderIdentity) String() string {
	return "folders/" + i.ResourceID
}

func ParseFolderExternal(external string) (*FolderIdentity, error) {
	tokens := strings.Split(external, "/")
	if len(tokens) == 2 && tokens[0] == "folders" {
		return &FolderIdentity{ResourceID: tokens[1]}, nil
	} else if len(tokens) == 1 {
		return &FolderIdentity{ResourceID: tokens[0]}, nil
	}
	return nil, fmt.Errorf("format of 'folderRef.external'=%q was not known (use folders/{{folderID}} or {{folderID}})", external)
}
