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
	"testing"
)

func TestTagsTagValueIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    string
		wantErr bool
	}{
		{
			name: "valid full URL",
			ref:  "//cloudresourcemanager.googleapis.com/tagValues/123",
			want: "123",
		},
		{
			name: "valid short URL",
			ref:  "tagValues/123",
			want: "123",
		},
		{
			name:    "invalid format",
			ref:     "tagValues/123/more",
			wantErr: true,
		},
		{
			name:    "invalid prefix",
			ref:     "otherValues/123",
			wantErr: true,
		},
		{
			name:    "empty value",
			ref:     "tagValues/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &TagsTagValueIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i.TagValue != tt.want {
				t.Errorf("FromExternal() got = %v, want %v", i.TagValue, tt.want)
			}
		})
	}
}
