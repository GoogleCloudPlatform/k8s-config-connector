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

package v1beta1

// MemorystoreInstanceRef defines the resource reference to MemorystoreInstance, which "External" field
// holds the GCP identifier for the KRM object.
type MemorystoreInstanceRef struct {
	// A reference to an externally managed MemorystoreInstance resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/instances/{{instanceID}}".
	External string `json:"external,omitempty"`

	// The name of a MemorystoreInstance resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MemorystoreInstance resource.
	Namespace string `json:"namespace,omitempty"`
}

// MemorystoreInstanceServiceAttachmentRef defines the resource reference to MemorystoreInstance managed ServiceAttachment
// which "ServiceAttachmentExternal" field holds the GCP identifier for the KRM object.
type MemorystoreInstanceServiceAttachmentRef struct {
	// The MemorystoreInstance managed ServiceAttachment of the form "projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}"
	ServiceAttachmentExternal string `json:"serviceAttachmentExternal,omitempty"`

	// The MemorystoreInstance resource that the ServiceAttachment is attached to.
	MemorystoreInstanceRef *MemorystoreInstanceRef `json:"memorystoreInstanceRef,omitempty"`

	// The index of the PSC attachment details in the MemorystoreInstance resource.
	PscAttachmentDetailsIndex *int `json:"pscAttachmentDetailsIndex,omitempty"`
}

// DeepCopyInto is copying the receiver, writing into out. in must be non-nil.
func (in *MemorystoreInstanceServiceAttachmentRef) DeepCopyInto(out *MemorystoreInstanceServiceAttachmentRef) {
	*out = *in
	if in.MemorystoreInstanceRef != nil {
		in, out := &in.MemorystoreInstanceRef, &out.MemorystoreInstanceRef
		*out = new(MemorystoreInstanceRef)
		**out = **in
	}
	if in.PscAttachmentDetailsIndex != nil {
		in, out := &in.PscAttachmentDetailsIndex, &out.PscAttachmentDetailsIndex
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is copying the receiver, creating a new MemorystoreInstanceServiceAttachmentRef.
func (in *MemorystoreInstanceServiceAttachmentRef) DeepCopy() *MemorystoreInstanceServiceAttachmentRef {
	if in == nil {
		return nil
	}
	out := new(MemorystoreInstanceServiceAttachmentRef)
	in.DeepCopyInto(out)
	return out
}
