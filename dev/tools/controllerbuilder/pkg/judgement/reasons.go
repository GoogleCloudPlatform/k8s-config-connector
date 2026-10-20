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

package judgement

// Reasons for an entry that says a string field may need to be a reference.
// The generator writes these, and TestMissingRefs uses IsReferenceReason to
// decide which [refs] findings an open entry covers. Keeping both in one place
// means the two cannot drift apart.
const (
	// ReasonAuthoritativeProtoReference: the proto field has a
	// google.api.resource_reference annotation. The detail names the type.
	ReasonAuthoritativeProtoReference = "authoritative-proto-reference"
	// ReasonPossibleReference: legacy alias for ReasonAuthoritativeProtoReference
	// retained for backward compatibility with existing judgement_queue.yaml files.
	ReasonPossibleReference = "possible-reference"
	// ReasonPossibleReferenceByDescription: refs.Classify, the rule
	// TestMissingRefs applies, says the field is a reference. The detail says
	// which of its rules matched.
	ReasonPossibleReferenceByDescription = "possible-reference-by-description"
	// ReasonPossibleReferenceByDescriptionLoose: a looser description rule
	// matched. TestMissingRefs would not flag the field.
	ReasonPossibleReferenceByDescriptionLoose = "possible-reference-by-description-loose"
	// ReasonPossibleReferenceByName: the field name matches a known
	// reference type.
	ReasonPossibleReferenceByName = "possible-reference-by-name"
	// ReasonPossibleReferenceBySibling: the field name matches a resource the
	// same service declares.
	ReasonPossibleReferenceBySibling = "possible-reference-by-sibling"
)

var referenceReasons = map[string]bool{
	ReasonAuthoritativeProtoReference:         true,
	ReasonPossibleReference:                   true,
	ReasonPossibleReferenceByDescription:      true,
	ReasonPossibleReferenceByDescriptionLoose: true,
	ReasonPossibleReferenceByName:             true,
	ReasonPossibleReferenceBySibling:          true,
}

// ReasonReferenceNotRepresentable: refs.Classify says the field names another
// resource, but KCC cannot express it as a reference today, so it stays a
// string. This is not a reference reason. TestMissingRefs lists such a field
// in refs_not_representable.txt and never fails on it, so an open entry has
// nothing to hide.
const ReasonReferenceNotRepresentable = "reference-not-representable"

// IsReferenceReason reports whether reason says a field may need to be a
// reference.
func IsReferenceReason(reason string) bool {
	return referenceReasons[reason]
}

// Reasons for a source link in the header of <kind>_types.go that
// generate-types could not verify. The link carries a
// "+kcc:guess=source-link reason=<reason>" line until someone fixes it.
const (
	// ReasonVerifyResourceDocsLink: no candidate for the resource's REST
	// reference page loaded as that page, or the check could not run.
	ReasonVerifyResourceDocsLink = "verify-resource-docs-link"
	// ReasonVerifyServiceDocsLink: the service docs link is missing, did not
	// load, or is a generic landing page.
	ReasonVerifyServiceDocsLink = "verify-service-docs-link"
)

// IsSourceLinkReason reports whether reason is about a source link.
func IsSourceLinkReason(reason string) bool {
	return reason == ReasonVerifyResourceDocsLink || reason == ReasonVerifyServiceDocsLink
}

// ReasonUntriagedBulkGeneration marks a Kind that nobody has reviewed yet.
// PrepopulateSpec queues one for each new Kind. While the entry is open, each
// generate-types run with --prepopulate-spec checks the Kind again.
const ReasonUntriagedBulkGeneration = "untriaged-bulk-generation"

// Reasons for identity and reference generation.
const (
	// ReasonIdentityMultiPattern: google.api.resource declares more than one
	// pattern; the generated identity uses the pattern matching the Spec layout
	// (or the first pattern).
	ReasonIdentityMultiPattern = "identity-multi-pattern"
	// ReasonIdentityRootUnknown: the resource declares no google.api.resource
	// pattern, or the Spec does not model the root/parent segments of the pattern.
	ReasonIdentityRootUnknown = "identity-root-unknown"
	// ReasonIdentityNotInCAI: the template URL is not listed in
	// docs/ai/metadata/cloudassetinventory_names.jsonl.
	ReasonIdentityNotInCAI = "identity-not-in-cai"
	// ReasonIdentityServerGeneratedID: the Create RPC has no *_id parameter (or
	// the Spec comment marks ResourceID server-generated), so the generated
	// identity implements identity.ServerGeneratedIdentity.
	ReasonIdentityServerGeneratedID = "identity-server-generated-id"
)

var identityReasons = map[string]bool{
	ReasonIdentityMultiPattern:      true,
	ReasonIdentityRootUnknown:       true,
	ReasonIdentityNotInCAI:          true,
	ReasonIdentityServerGeneratedID: true,
}

// IsIdentityReason reports whether reason is about an identity or reference file.
func IsIdentityReason(reason string) bool {
	return identityReasons[reason]
}

// ReasonRequiredNotEnforced: the proto marks a spec field REQUIRED, but the
// CRD leaves it optional. generate-types --emit-required-from-proto writes one
// per field, for each Kind in the run. Making the field required breaks
// objects that omit it, so a person decides: an alpha Kind may take the break,
// and a beta or stable Kind keeps the field optional.
const ReasonRequiredNotEnforced = "required-not-enforced"

// ReasonEtagInSpec: a top-level etag field was generated in Spec because the
// proto does not mark it OUTPUT_ONLY.
const ReasonEtagInSpec = "etag-in-spec"

// ReasonPossibleStatusFieldInSpec: an optional top-level field with a
// well-known status, lifecycle, or identifier name (such as state, status,
// startTime, endTime, lastUpdateTime, createTime, updateTime, deleteTime, id,
// uid, selfLink, or reconciling) was generated in Spec because the proto does
// not mark it OUTPUT_ONLY.
const ReasonPossibleStatusFieldInSpec = "possible-status-field-in-spec"

// ReasonLocationOrParentRef: a location, region, or zone field in Spec may be
// replaced by a parent/location reference or kept as a plain string.
const ReasonLocationOrParentRef = "location-or-parent-ref"

// Reasons for deprecated proto resources, fields, and messages.
const (
	// ReasonDeprecatedResource: the proto marks the resource message itself
	// deprecated (option deprecated = true).
	ReasonDeprecatedResource = "deprecated-resource"
	// ReasonDeprecatedField: the proto marks the field deprecated.
	ReasonDeprecatedField = "deprecated-field"
	// ReasonDeprecatedMessage: the proto marks the field's message type
	// deprecated (option deprecated = true).
	ReasonDeprecatedMessage = "deprecated-message"
)

// Reasons for spec and observed-state field review items.
const (
	// ReasonWellKnownProtoType: the field uses a google.protobuf.* message type
	// that was generated as a Go struct rather than mapped to a scalar or
	// special KRM type.
	ReasonWellKnownProtoType = "well-known-proto-type"
	// ReasonSensitiveField: the field name or path indicates sensitive/secret
	// data generated as a plain value.
	ReasonSensitiveField = "sensitive-field"
	// ReasonMessageMapDerived: a map<string, Message> field was generated as
	// map[string]<GoType> under --emit-message-maps.
	ReasonMessageMapDerived = "message-map-derived"
	// ReasonUnsupportedFieldType: the generator could not produce a Go type for
	// the proto field.
	ReasonUnsupportedFieldType = "unsupported-field-type"
	// ReasonEmptyObservedState: no fields were generated into
	// status.observedState for the resource.
	ReasonEmptyObservedState = "empty-observedstate"
	// ReasonAmbiguousProtoName: the --resource proto name matched more than one
	// listed service package.
	ReasonAmbiguousProtoName = "ambiguous-proto-name"
	// ReasonServerSetFieldPlaced: a field on the server-set allowlist was moved
	// to ObservedState because the message carries no field_behavior.
	ReasonServerSetFieldPlaced = "server-set-field-placed"
	// ReasonObservedStateIdentityFieldOmitted: an OUTPUT_ONLY identity field
	// (such as name) was omitted from ObservedState.
	ReasonObservedStateIdentityFieldOmitted = "observedstate-identity-field-omitted"
	// ReasonWellKnownOutputOnlyPattern: a field's proto comment starts with
	// "Output only." or "[Output Only]" without a field_behavior annotation.
	ReasonWellKnownOutputOnlyPattern = "well-known-output-only-pattern-in-comment"
	// ReasonPossibleOutputOnlyPattern: a field's proto comment mentions
	// "output only" without starting with a well-known prefix.
	ReasonPossibleOutputOnlyPattern = "possible-output-only-pattern-in-comment"
)

