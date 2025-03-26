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

// This acronym list contains both acronym (including initialism) and abbreviation.
// - acronyms use all-cap case as Kubernetes API convention suggested. https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#naming-conventions
// - abbreviations use its most known form with upper letters reflecting how it is pronounced.
// - Plural form of acronym should avoid using upper "S" unless it is a const.
var Acronyms = []string{

	"API",
	"BGP", "BYOID",
	"CA", "CDN", "CIDR", "CPU",
	"DNS",
	"EUC",
	"FS",
	"GCE", "GB", "GCS", "GKE",
	"HTML", "HTTP", "HTTPS",
	"IAM", "IAP", "ID", "IP", "IPV4", "IPV6",
	"KMS",
	"MiB",
	"NAT",
	"OAuth2", "OIDC", "OS",
	"PD", "PSC",
	"SQL", "SSH", "SSL", "SSO",
	"TCP", "TLS", "TTL",
	"UDP", "URI", "URL",
	"VTPM", "VM", "VPC", "VIP", "VPN",
	"X509",
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
