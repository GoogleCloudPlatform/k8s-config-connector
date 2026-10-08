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

// Package refs decides whether a CRD spec field ought to be a KCC reference,
// from the two things a CRD keeps about it: the field's path and its
// description.
//
// TestMissingRefs in tests/apichecks applies these rules to every CRD. It
// records each IsReference in missingrefs.txt, a ratchet, and each
// NotRepresentable in refs_not_representable.txt. The rules live in the
// generator module, which tests/apichecks already imports, so generate-types
// can apply them as well.
package refs

import (
	"regexp"
	"strings"
)

// Verdict is what Classify concluded about a field.
type Verdict int

const (
	// NotAReference means nothing suggests the field names another resource.
	NotAReference Verdict = iota
	// IsReference means the field names another resource and a KCC reference
	// can express it.
	IsReference
	// NotRepresentable means the field names another resource, but KCC cannot
	// model it as a reference today.
	NotRepresentable
)

// IsReferenceFieldPath reports whether fieldPath is already a KCC reference, or
// the external or name field inside one.
func IsReferenceFieldPath(fieldPath string) bool {
	for _, suffix := range []string{"Ref", "Refs[]", "Refs", "Ref.external", "Refs[].external", "Ref.name"} {
		if strings.HasSuffix(fieldPath, suffix) {
			return true
		}
	}
	return false
}

// Classify returns the verdict for a spec field from its path and description,
// and for NotRepresentable the reason. Callers skip status fields and paths
// IsReferenceFieldPath accepts before calling it.
func Classify(fieldPath, desc string) (Verdict, string) {
	// A pattern or matcher field describes a set of resources, never one, so it
	// is not a reference of either kind. dlp's cloudStorageRegex.bucketNameRegex
	// is an example: "Regex to test the bucket name against ...
	// (marketing)\d{4}".
	if isPatternField(fieldPath, desc) {
		return NotAReference, ""
	}

	if reason := notRepresentableReason(fieldPath, desc); reason != "" {
		return NotRepresentable, reason
	}

	isRef := false

	// Proto-level resource references (google.api.resource_reference) are resolved
	// in the proto generator where the proto field descriptors are available.
	// In Classify, which operates on KRM field paths and doc strings without direct
	// proto context, heuristics are used instead.

	// A resource-name template in the description marks a reference, as in
	// "projects/{projectID}/locations/{location}/bars/{name}".
	if hasResourceNameTemplate(desc) {
		isRef = true
	}

	// The suffix matches serviceAccount and ServiceAccount alike.
	if strings.HasSuffix(fieldPath, "erviceAccount") {
		isRef = true
	}
	// TODO: how to detect KMS Key

	// A Cloud Storage bucket reference is expressible today:
	// StorageBucketIdentity.FromExternal accepts the bare "gs://<bucket>"
	// form (apis/storage/v1beta1/storagebucket_identity.go).
	// notRepresentableReason has already handled object paths.
	if hasToken(fieldPath, "bucket") && mentionsCloudStorage(desc) {
		isRef = true
	}

	if !isRef {
		return NotAReference, ""
	}

	// No reference is required for a zone, region, location, machine type or
	// accelerator type. These are taxonomy values rather than managed objects.
	// The plural forms are excluded too: a field naming a set of zones is no
	// more a reference than one naming a single zone.
	switch {
	case strings.HasSuffix(fieldPath, ".zone"), strings.HasSuffix(fieldPath, ".zones"),
		strings.HasSuffix(fieldPath, ".region"), strings.HasSuffix(fieldPath, ".regions"),
		strings.HasSuffix(fieldPath, ".location"), strings.HasSuffix(fieldPath, ".locations"),
		strings.HasSuffix(fieldPath, ".machineType"),
		strings.HasSuffix(fieldPath, ".acceleratorType"):
		return NotAReference, ""
	}
	return IsReference, ""
}

// resourceNamePrefixes are the collection segments that start a GCP resource
// name. A description containing one of these followed by a placeholder is
// describing a resource name, i.e. a reference.
var resourceNamePrefixes = []string{"projects/", "locations/", "zones/", "regions/", "organizations/", "folders/"}

// hasResourceNameTemplate reports whether desc contains a resource-name path
// template.
//
// Placeholder syntax varies across APIs, and so does the character before
// the path. notification_channels, for example, reads "Must be of the format
// `projects/<project_id_or_number>/notificationChannels/<channel_id>`", with
// a backtick rather than a space and <> rather than {}. Both forms count.
func hasResourceNameTemplate(desc string) bool {
	// A placeholder immediately after a collection segment is unambiguous: the
	// description is spelling out a resource name. Both brace and angle-bracket
	// syntaxes are used upstream.
	for _, prefix := range resourceNamePrefixes {
		if strings.Contains(desc, prefix+"{") || strings.Contains(desc, prefix+"<") {
			return true
		}
	}

	// Without a placeholder, only a space-delimited "projects/" counts. Widening
	// this to other delimiters or to "locations/" matches ordinary prose, and it
	// produced findings on container.username and allowedLocations, which are
	// not references.
	return strings.Contains(desc, " projects/")
}

// hasToken reports whether the last segment of fieldPath contains tok as a
// whole word, after splitting it at camelCase boundaries and underscores. A
// substring match is wrong here: "security" contains "uri".
func hasToken(fieldPath string, tok string) bool {
	// Take the leaf segment, then split camelCase into lowercase words.
	leaf := fieldPath
	if i := strings.LastIndex(leaf, "."); i >= 0 {
		leaf = leaf[i+1:]
	}
	leaf = strings.TrimSuffix(leaf, "[]")
	leaf = strings.ReplaceAll(leaf, "_", " ")

	for _, w := range splitCamelWords(leaf) {
		if w == tok {
			return true
		}
	}
	return false
}

// splitCamelWords lowercases and splits identifiers on camelCase boundaries,
// keeping acronym runs intact: "outputURIPrefix" -> [output uri prefix].
// A naive split would yield [output u r i prefix] and never match "uri".
func splitCamelWords(s string) []string {
	isUpper := func(r rune) bool { return r >= 'A' && r <= 'Z' }
	isLower := func(r rune) bool { return r >= 'a' && r <= 'z' }

	runes := []rune(s)
	var words []string
	start := 0
	flush := func(end int) {
		w := strings.ToLower(strings.TrimSpace(string(runes[start:end])))
		if w != "" {
			words = append(words, w)
		}
		start = end
	}
	for i := 1; i < len(runes); i++ {
		prev, cur := runes[i-1], runes[i]
		switch {
		case isLower(prev) && isUpper(cur):
			// output|URI
			flush(i)
		case isUpper(prev) && isUpper(cur) && i+1 < len(runes) && isLower(runes[i+1]):
			// URI|Prefix
			flush(i)
		case cur == ' ' || prev == ' ':
			flush(i)
		}
	}
	flush(len(runes))
	return words
}

// mentionsCloudStorage reports whether desc names Cloud Storage or a gs://
// URI.
func mentionsCloudStorage(desc string) bool {
	return strings.Contains(desc, "Cloud Storage") || strings.Contains(desc, "gs://")
}

// isPatternField reports whether a field holds a matcher (regex/glob/filter)
// rather than the identity of one resource. Classify reports such a field as
// neither a reference nor a non-representable one.
func isPatternField(fieldPath, desc string) bool {
	for _, tok := range []string{"regex", "pattern", "filter", "matcher", "matchers"} {
		if hasToken(fieldPath, tok) {
			return true
		}
	}
	trimmed := strings.TrimSpace(desc)
	if strings.HasPrefix(trimmed, "Regex") ||
		strings.HasPrefix(trimmed, "Optional. Regex") ||
		strings.Contains(desc, "Regex to test") {
		return true
	}

	// Wildcard patterns match sets of resources rather than identifying a single
	// target resource, so they are not treated as direct resource references.
	return mentionsWildcard(desc)
}

// mentionsWildcard reports whether a description says its value may be a glob.
//
// It requires the description to say so rather than looking for a literal
// "*", because "*" appears in ordinary prose, as markdown emphasis or a
// footnote marker, far more often than it denotes a glob.
func mentionsWildcard(desc string) bool {
	lower := strings.ToLower(desc)
	for _, phrase := range []string{"wildcard", "may contain wildcards", "glob pattern"} {
		if strings.Contains(lower, phrase) {
			return true
		}
	}
	return false
}

// mentionsContainerImage reports whether desc names a container image
// registry.
//
// notRepresentableReason checks this before Cloud Storage, because image
// descriptions often mention Cloud Storage in passing. executorImageURI reads
// "an image in Artifact Registry ... Google Cloud Storage paths ... will be
// mapped to local paths", and would otherwise be filed as an object path.
func mentionsContainerImage(desc string) bool {
	return strings.Contains(desc, "Artifact Registry") ||
		strings.Contains(desc, "Container Registry") ||
		strings.Contains(desc, "gcr.io") ||
		strings.Contains(desc, "-docker.pkg.dev") ||
		strings.Contains(desc, "container image")
}

// notRepresentableReason returns why a field that looks like a reference
// cannot be modeled as one today, or "" when nothing stops it. TestMissingRefs
// records these in refs_not_representable.txt rather than missingrefs.txt, so
// the ratchet lists only fields someone can fix.
func notRepresentableReason(fieldPath, desc string) string {
	isURIField := hasToken(fieldPath, "uri") || hasToken(fieldPath, "uris") ||
		hasToken(fieldPath, "url") || hasToken(fieldPath, "urls")

	// The registry and bq:// checks come before the Cloud Storage one; see
	// mentionsContainerImage.
	if isURIField && mentionsContainerImage(desc) {
		return "container-image-uri-not-a-storage-object"
	}

	// bq:// is not a GCP resource name. Its arity is ambiguous within a single
	// field (bq://p, bq://p.d and bq://p.d.t address three different kinds), and
	// services disagree on the separator: aiplatform uses dots, and datalabeling
	// documents slashes. KCC has no bq:// parsing.
	if strings.Contains(desc, "bq://") {
		return "bq-scheme-not-a-gcp-resource-name"
	}
	if isURIField && strings.Contains(desc, "BigQuery") {
		return "bigquery-uri-not-a-gcp-resource-name"
	}

	// Cloud Storage object paths and prefixes are not addressable KCC resources:
	// there is no StorageObject CRD, and StorageBucketIdentity rejects anything
	// with a "/" after the bucket. Classify treats a bucket-only field as a
	// reference.
	if isURIField && mentionsCloudStorage(desc) && !hasToken(fieldPath, "bucket") {
		if hasToken(fieldPath, "prefix") || hasToken(fieldPath, "directory") ||
			strings.Contains(desc, "output directory") || strings.Contains(desc, "directory path") {
			return "gcs-prefix-needs-bucket-ref-plus-path"
		}
		return "gcs-object-path-string-for-now-decomposable-as-bucketref-plus-path"
	}

	return ""
}

// NameRule matches a field by its leaf name, and names the reference type the
// field probably points at.
type NameRule struct {
	// Target is the KCC ref type the field points at, for the hint text.
	Target string
	// Match takes the field's leaf name, already stripped of any list suffix.
	Match func(leaf string) bool
}

// eq returns a matcher for a leaf equal to one of names, ignoring case.
func eq(names ...string) func(string) bool {
	return func(leaf string) bool {
		for _, n := range names {
			if strings.EqualFold(leaf, n) {
				return true
			}
		}
		return false
	}
}

// hasSuffix returns a matcher for a leaf longer than suffix that ends in it,
// ignoring case.
func hasSuffix(suffix string) func(string) bool {
	return func(leaf string) bool {
		return len(leaf) > len(suffix) && strings.EqualFold(leaf[len(leaf)-len(suffix):], suffix)
	}
}

// NameRules maps field naming conventions to specific target reference types.
// Rules are evaluated in order; the first matching rule wins.
var NameRules = []NameRule{
	// Every confirmed instance ends in SecretVersion.
	{Target: "SecretManagerSecretVersionRef", Match: hasSuffix("SecretVersion")},
	// Matches common network identifier names as exact words.
	{Target: "ComputeNetworkRef", Match: eq("network", "vpc", "vpcName")},
	{Target: "KMSCryptoKeyRef", Match: eq("kmsKey", "cmekKeyName", "encryptionKey", "kmsKeyName")},
	// Matches fields containing "secret" as a discrete word, excluding "SecretVersion".
	{Target: "SecretManagerSecretVersionRef", Match: wholeWordNotSuffix("secret", "SecretVersion")},
	// Matches exact "project" leaf fields (excluding projectNumber or projectID).
	{Target: "ProjectRef", Match: eq("project")},
}

// wholeWordNotSuffix returns a matcher for a leaf that contains word as a
// camelCase word and does not end in exclude.
func wholeWordNotSuffix(word, exclude string) func(string) bool {
	return func(leaf string) bool {
		if strings.EqualFold(leaf[max(0, len(leaf)-len(exclude)):], exclude) {
			return false
		}
		for _, w := range splitCamelWords(leaf) {
			if w == word {
				return true
			}
		}
		return false
	}
}

// Patterns that identify resource names or URI formats in field descriptions.
var (
	looseResourceName = regexp.MustCompile(`(?i)resource name(\s*\(URI\))?\s+(of|for)\b`)
	looseResourceURL  = regexp.MustCompile(`(?i)resource URL`)
	looseBracketPath  = regexp.MustCompile(`(projects|locations|organizations|folders)/\[`)
	// Taxonomy values (zones, machine types, images) that represent catalogue entries rather than managed resources.
	looseCatalogueEntry = regexp.MustCompile(`(?i)(zone|location|machineType|acceleratorType|image|nodeType|diskType)(URI|Url)?$`)
)

// MatchDescriptionLoose reports whether desc describes a resource name or URI in
// prose or bracketed template format. This is used by generator hints to identify
// potential references for review.
func MatchDescriptionLoose(fieldPath, desc string) bool {
	leaf := fieldPath
	if i := strings.LastIndex(leaf, "."); i >= 0 {
		leaf = leaf[i+1:]
	}
	leaf = strings.TrimSuffix(leaf, "[]")
	if looseCatalogueEntry.MatchString(leaf) {
		return false
	}
	if looseResourceURL.MatchString(desc) || looseBracketPath.MatchString(desc) {
		return true
	}
	// A sub-message's own "name" is excluded (e.g. conversionWorkspace.name).
	return !strings.EqualFold(leaf, "name") && looseResourceName.MatchString(desc)
}

// MatchName returns the reference type indicated by a field's name, if any.
// This is used by the generator when proposing reference scaffolding.
func MatchName(fieldPath string) (string, bool) {
	leaf := fieldPath
	if i := strings.LastIndex(leaf, "."); i >= 0 {
		leaf = leaf[i+1:]
	}
	leaf = strings.TrimSuffix(leaf, "[]")
	for _, r := range NameRules {
		if r.Match(leaf) {
			return r.Target, true
		}
	}
	return "", false
}
