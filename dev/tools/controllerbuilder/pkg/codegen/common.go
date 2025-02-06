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

package codegen

import "strings"

const (
	// KCCProtoMessageAnnotation is used for go structs that map to proto messages
	KCCProtoMessageAnnotation = "+kcc:proto"

	// KCCProtoFieldAnnotation is used for go struct fields that map to proto fields
	KCCProtoFieldAnnotation = "+kcc:proto:field"

	// KCCProtoIgnoreAnnotation is used for go struct fields that are ignored
	KCCProtoIgnoreAnnotation = "+kcc:proto:ignore"
)

// special-case proto messages that are currently not mapped to KRM Go structs
var protoMessagesNotMappedToGoStruct = map[string]string{
	"google.protobuf.Timestamp":   "string",
	"google.protobuf.Duration":    "string",
	"google.protobuf.Int64Value":  "int64",
	"google.protobuf.StringValue": "string",
	"google.protobuf.BoolValue":   "bool",
	"google.protobuf.Struct":      "map[string]string",
}

var Acronyms = []string{
	"ID", "HTML", "URL", "URI", "HTTP", "HTTPS", "SSH",
	"IP", "GB", "FS", "PD", "KMS", "GCE", "VTPM", "SSL",
	"CA", "CPU", "SQL", "PSC", "DNS",
}

// IsAcronym returns true if the given string is an acronym
func IsAcronym(s string) bool {
	for _, acronym := range Acronyms {
		if strings.EqualFold(s, acronym) {
			return true
		}
	}
	return false
}
