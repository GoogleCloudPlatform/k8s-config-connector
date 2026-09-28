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

package protoapi

// Archetype classifies a resource by the shape of its standard methods, which
// decides which controller template, if any, fits it. The value is a letter
// code; the codes are stable, so reports and judgement entries can quote them.
//
// The methods are those of the ResourceAPI. An AIP method whose request lacks
// the AIP shape, such as a GetCluster that takes project_id and cluster_name
// or an UpdateCluster that takes a ClusterUpdate, is there but flagged
// NonstandardRequest. It makes the resource ArchetypeNonstandard: reading it
// as missing instead would, for example, turn a resource that can be updated
// into ArchetypeNoUpdate.
type Archetype string

const (
	// ArchetypeStandard (A) has Get, Create with a client-assigned ID, Update
	// with an update_mask, and Delete.
	ArchetypeStandard Archetype = "A"
	// ArchetypeNoUpdate (B) has Get, Create and Delete but no Update, whether
	// or not Create takes an ID.
	ArchetypeNoUpdate Archetype = "B"
	// ArchetypeServerGeneratedID (C) has Get, Create with no ID field, Update
	// with an update_mask, and Delete.
	//
	// C also catches creates where the client sets the name in the request
	// body, such as a Composer Environment or a Cloud Tasks Queue. The
	// descriptors cannot tell those apart from server-generated IDs, so the
	// generator has to queue C kinds for judgement.
	ArchetypeServerGeneratedID Archetype = "C"
	// ArchetypeSingleton (D) has Get and Update but no Create.
	ArchetypeSingleton Archetype = "D"
	// ArchetypeNonstandard (E) means the API is nonstandard for the resource:
	// it has neither a Get nor a Create, or one of its Get, Create, Update
	// and Delete is flagged NonstandardRequest. It does not follow the
	// compute convention either.
	ArchetypeNonstandard Archetype = "E"
	// ArchetypeUpdateWithoutMask (F) has Get, Create, Update and Delete, but
	// its Update takes no update_mask.
	ArchetypeUpdateWithoutMask Archetype = "F"
	// ArchetypeOther (G) has AIP-shaped methods in a combination that none
	// of the other archetypes covers, for example Get and Create with no
	// Delete.
	ArchetypeOther Archetype = "G"
	// ArchetypeComputeStyle (H) follows the google.cloud.compute convention
	// (StyleCompute).
	ArchetypeComputeStyle Archetype = "H"
	// ArchetypeUnresolved (I) means the message is not in the descriptor set.
	ArchetypeUnresolved Archetype = "I"
)

// AllArchetypes returns every archetype, in letter order.
func AllArchetypes() []Archetype {
	return []Archetype{
		ArchetypeStandard,
		ArchetypeNoUpdate,
		ArchetypeServerGeneratedID,
		ArchetypeSingleton,
		ArchetypeNonstandard,
		ArchetypeUpdateWithoutMask,
		ArchetypeOther,
		ArchetypeComputeStyle,
		ArchetypeUnresolved,
	}
}

// Name returns a short human-readable name for the archetype, such as
// "standard" for A.
func (a Archetype) Name() string {
	switch a {
	case ArchetypeStandard:
		return "standard"
	case ArchetypeNoUpdate:
		return "no-update"
	case ArchetypeServerGeneratedID:
		return "server-generated-id"
	case ArchetypeSingleton:
		return "singleton"
	case ArchetypeNonstandard:
		return "nonstandard"
	case ArchetypeUpdateWithoutMask:
		return "update-without-mask"
	case ArchetypeOther:
		return "other"
	case ArchetypeComputeStyle:
		return "compute-style"
	case ArchetypeUnresolved:
		return "unresolved"
	}
	return "unknown"
}

// ClassifyArchetype returns the archetype of a resource. A nil api, which is
// what a failed lookup produces, is ArchetypeUnresolved.
//
// The checks run in a fixed order and the first match wins: I, H, E, D, A, C,
// B, F, then G for anything left.
func ClassifyArchetype(api *ResourceAPI) Archetype {
	if api == nil || api.Message == nil {
		return ArchetypeUnresolved
	}
	if api.Style == StyleCompute {
		return ArchetypeComputeStyle
	}

	get, create, update, del := api.Get != nil, api.Create != nil, api.Update != nil, api.Delete != nil
	crud := get && create && update && del
	hasID := create && api.Create.IDField != ""
	hasMask := update && api.Update.UpdateMaskField != ""
	switch {
	case !get && !create, hasNonstandardRequest(api):
		return ArchetypeNonstandard
	case get && !create && update:
		return ArchetypeSingleton
	case crud && hasID && hasMask:
		return ArchetypeStandard
	case crud && !hasID && hasMask:
		return ArchetypeServerGeneratedID
	case get && create && !update && del:
		return ArchetypeNoUpdate
	case crud && !hasMask:
		return ArchetypeUpdateWithoutMask
	default:
		return ArchetypeOther
	}
}

// hasNonstandardRequest reports whether any of api's Get, Create, Update and
// Delete is flagged NonstandardRequest.
func hasNonstandardRequest(api *ResourceAPI) bool {
	for _, m := range []*StandardMethod{api.Get, api.Create, api.Update, api.Delete} {
		if m != nil && m.NonstandardRequest {
			return true
		}
	}
	return false
}
