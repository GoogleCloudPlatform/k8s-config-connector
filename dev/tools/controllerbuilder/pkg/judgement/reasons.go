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
	// ReasonPossibleReference: the proto field has a
	// google.api.resource_reference annotation. The detail names the type.
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

// ReasonRequiredNotEnforced: the proto marks a spec field REQUIRED, but the
// CRD leaves it optional. generate-types --emit-required-from-proto writes one
// per field, for each Kind in the run. Making the field required breaks
// objects that omit it, so a person decides: an alpha Kind may take the break,
// and a beta or GA Kind keeps the field optional.
const ReasonRequiredNotEnforced = "required-not-enforced"
