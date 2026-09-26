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
	// KCCProtoMessageAnnotationMisc is used for go structs that map to proto messages, but are not top-level Spec structs
	KCCProtoMessageAnnotationMisc = "+kcc:proto"
	// KCCProtoMessageAnnotationSpec is used for top-level Spec structs that map to proto messages
	KCCProtoMessageAnnotationSpec = "+kcc:spec:proto"
	// KCCProtoMessageAnnotationObservedState is used for top-level ObservedState structs that map to proto messages
	KCCProtoMessageAnnotationObservedState = "+kcc:observedstate:proto"
	// KCCProtoMessageAnnotationObservedState is used for top-level Status structs that map to proto messages
	// (Used by legacy controllers)
	KCCProtoMessageAnnotationStatus = "+kcc:status:proto"

	// KCCProtoFieldAnnotation is used for go struct fields that map to proto fields
	KCCProtoFieldAnnotation = "+kcc:proto:field"

	// KCCProtoIgnoreAnnotation is used for go struct fields that are ignored
	KCCProtoIgnoreAnnotation = "+kcc:proto:ignore"
)

// GetProtoMessageFromAnnotation will extract a proto message annotation, including the spec and observedstate "subclasses"
func GetProtoMessageFromAnnotation(commentLine string) (string, bool) {
	protoMessage, _, ok := GetProtoMessageAndKindFromAnnotation(commentLine)
	return protoMessage, ok
}

// GetProtoMessageAndKindFromAnnotation extracts a proto message annotation from a comment line,
// returning both the proto message full name and the specific annotation kind that matched
// (e.g. "+kcc:proto" or "+kcc:observedstate:proto").
//
// Identifying the matched annotation enables callers to determine whether the annotated Go type
// represents a spec struct (e.g. "Foo") or an ObservedState struct (e.g. "FooObservedState")
// when both refer to the same underlying protobuf message.
func GetProtoMessageAndKindFromAnnotation(commentLine string) (protoMessage string, annotationKind string, ok bool) {
	trimmed := strings.TrimPrefix(commentLine, "//")
	trimmed = strings.TrimSpace(trimmed)
	for _, annotation := range []string{
		KCCProtoMessageAnnotationMisc,
		KCCProtoMessageAnnotationSpec,
		KCCProtoMessageAnnotationObservedState,
		KCCProtoMessageAnnotationStatus,
	} {
		if strings.HasPrefix(trimmed, annotation+"=") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, annotation+"=")), annotation, true
		}
	}
	return "", "", false
}

// special-case proto messages that are currently not mapped to KRM Go structs
var protoMessagesNotMappedToGoStruct = map[string]string{
	"google.protobuf.Timestamp":         "string",
	"google.protobuf.Duration":          "string",
	"google.protobuf.Int64Value":        "int64",
	"google.protobuf.StringValue":       "string",
	"google.protobuf.BoolValue":         "bool",
	"google.protobuf.Struct":            "apiextensionsv1.JSON",
	"google.rpc.Status":                 "common.Status",
	"google.cloud.connectors.v1.Secret": "secretmanagerv1beta1.SecretRef",
}

// QualifierImports maps the package qualifier of each Go type in
// protoMessagesNotMappedToGoStruct to its import path. If an entry in that map
// uses an external package, declare its import path here.
var QualifierImports = map[string]string{
	"apiextensionsv1":      "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1",
	"common":               "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common",
	"secretmanagerv1beta1": "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1",
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
	"FS", "FQDN",
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

// AcronymCasing returns the correctly-cased form of a name token, and whether
// it was an acronym at all.
//
// With plurals set it also matches a plural acronym, which IsAcronym does not:
// EqualFold("Uris", "URI") is false, so the generator has always written
// RelatedUris and KbArticleIds. TestCRDsAcronyms reads the same Acronyms list
// but strips a trailing "s" before matching, so it asks for RelatedURIs and
// KbArticleIDs and records the difference in acronyms.txt. Singular has always
// worked; the same struct has SupportURL from support_url.
//
// Off by default because switching it on renames 83 fields across 23 packages,
// 34 of them in v1beta1, and renaming a served field is a worse break than
// moving one.
func AcronymCasing(token string, plurals bool) (string, bool) {
	if IsAcronym(token) {
		return strings.ToUpper(token), true
	}
	if !plurals || len(token) < 2 {
		return "", false
	}
	if last := token[len(token)-1]; last != 's' && last != 'S' {
		return "", false
	}
	if stem := token[:len(token)-1]; IsAcronym(stem) {
		return strings.ToUpper(stem) + "s", true
	}
	return "", false
}
