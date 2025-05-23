package v1alpha1

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// GroupVersion is group version used to register these objects
var GroupVersion = schema.GroupVersion{Group: "accesscontextmanager.cnrm.cloud.google.com", Version: "v1alpha1"}

// SchemeBuilder is used to add go types to the GroupVersionKind scheme
var SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

// AddToScheme adds the types in this group-version to the given scheme.
func AddToScheme(s *scheme.Scheme) error {
	return SchemeBuilder.AddToScheme(s)
}

const AccessPolicyKind = "AccessContextManagerAccessPolicy"

// +k8s:deepcopy-gen=true
type AccessPolicyParent struct {
	OrganizationID string
}

func (p *AccessPolicyParent) String() string {
	if p.OrganizationID == "" {
		return ""
	}
	return fmt.Sprintf("organizations/%s", p.OrganizationID)
}

func (p *AccessPolicyParent) Set(val string) error {
	parsed, err := ParseAccessPolicyExternal(val)
	if err != nil {
		return err
	}
	*p = parsed
	return nil
}

func (p *AccessPolicyParent) Type() string {
	return "AccessPolicyParent"
}

// ParseAccessPolicyExternal parses the external resource name into an AccessPolicyParent.
// Expected format: organizations/{organization_id}
func ParseAccessPolicyExternal(externalName string) (AccessPolicyParent, error) {
	parts := strings.Split(externalName, "/")
	if len(parts) == 2 && parts[0] == "organizations" {
		return AccessPolicyParent{OrganizationID: parts[1]}, nil
	}
	return AccessPolicyParent{}, fmt.Errorf("cannot parse %s: expected format organizations/{organization_id}", externalName)
}

// +k8s:deepcopy-gen=true
type AccessPolicyIdentity struct {
	AccessPolicyParent
	Name string
}

func NewAccessPolicyIdentity(parent AccessPolicyParent, name string) AccessPolicyIdentity {
	return AccessPolicyIdentity{
		AccessPolicyParent: parent,
		Name:               name,
	}
}

func (i *AccessPolicyIdentity) String() string {
	// The format "{parent}/accessPolicies/{name}" is based on the GCP API structure.
	// For AccessPolicy, the parent is an organization, and the resource name is the policy ID.
	// Example from google.identity.accesscontextmanager.v1.GetAccessPolicyRequest:
	// Format `accessPolicies/{policy_id}` which is relative to the parent organization.
	// The full path would be organizations/{org_id}/accessPolicies/{policy_id}
	// However, the GetAccessPolicy RPC uses /v1/{name=accessPolicies/*} where name includes the parent.
	// The 'name' field in AccessPolicy proto is "accessPolicies/{access_policy}" which does not include the "organizations/{id}" part.
	// The 'parent' field in AccessPolicy proto is "organizations/{organization_id}".
	// The ListAccessPolicies RPC takes parent "organizations/{org_id}" and returns policies with name "accessPolicies/{policy_id}".
	// The GetAccessPolicy RPC takes name "accessPolicies/{policy_id}" (which is the resource name relative to service, not its full path).
	// The CRD uses spec.parent.organizationRef for the parent.
	// For KCC identity, we usually construct the full path to the resource if the 'name' field in API is just the final segment.
	// Given the API structure `get: "/v1/{name=accessPolicies/*}"` for GetAccessPolicy, where 'name' means the full name.
	// And `parent` in `AccessPolicy` message: `organizations/{organization_id}`.
	// And `name` in `AccessPolicy` message: `accessPolicies/{access_policy}`.
	// The actual full resource name on GCP is `accessPolicies/{numeric_id}`, where this numeric_id is the policy ID.
	// The `parent` specified in the AccessPolicy is `organizations/{organization_id}`.
	// The URL for Get/Delete is /v1/{name=accessPolicies/*}
	// The URL for Create is /v1/accessPolicies (parent is implicit in the request body)
	// It seems the convention for `String()` is to produce the `name` that would be used in a `Get` request.
	// The proto defines `AccessPolicy` name as `accessPolicies/{access_policy}`. This name does NOT include the `organizations/{org_id}` prefix.
	// The `parent` field of `AccessPolicy` is `organizations/{organization_id}`.
	// The `name` field in KCC spec is `metadata.name`. This usually maps to the `{access_policy}` part.
	// Let's stick to the convention that `AccessPolicyIdentity.String()` returns the value that can be used in a `Get` request for the `name` field.
	// The GetAccessPolicyRequest takes a `name` argument: Format `accessPolicies/{policy_id}`.
	// This is what we should return. The parent context is handled by the client usually.
	// However, KCC often uses the fully qualified name for identity.
	// Let's review an existing similar resource, e.g. BigQueryDataset.
	// DatasetId includes ProjectId. String() returns projects/{project_id}/datasets/{dataset_id}.
	// `name` for Dataset is `projects/{{project}}/datasets/{{datasetId}}`.
	// The `name` field of `AccessPolicy` message is `accessPolicies/{access_policy}`.
	// The `parent` field of `AccessPolicy` message is `organizations/{organization_id}`.
	// The `GetAccessPolicyRequest` message has a `name` field, format `accessPolicies/{policy_id}`.
	// It is not organizations/{org_id}/accessPolicies/{policy_id}.
	// This is confusing. Let's assume the string representation should be the `name` as expected by the `Get` RPC.
	// The `name` in `GetAccessPolicyRequest` is `accessPolicies/{policy_id}`.
	// The controller will resolve the parent for the create call.
	// The terraform provider uses "name" as accessPolicies/{policy_id} for get, update, delete.
	// For create, it uses "parent" (organizations/{org_id}) and a body with title.
	// So the `name` is indeed `accessPolicies/{i.Name}` where i.Name is the numeric ID.
	// The `parent` from `i.AccessPolicyParent.String()` is `organizations/{org_id}`.
	// The URL for `get` is `/v1/{name=accessPolicies/*}`. This means `name` parameter IS `accessPolicies/{id}`.
	// The confusion is whether the parent part should be prepended.
	// Given the AccessPolicy message pattern is "accessPolicies/{access_policy}",
	// this is its canonical name *within the service*.
	// For a Get request, the path is /v1/{name=accessPolicies/*}.
	// It seems the `name` field in such Get requests *is* `accessPolicies/{id}`.
	// Let's verify this against other KCC resources.
	// For `IAMServiceAccount`, identity string is `projects/%s/serviceAccounts/%s`. The Get takes `name=projects/{project}/serviceAccounts/{email}`.
	// This suggests the String() method *should* include the parent.
	// The `name` in `AccessPolicy` proto definition is `accessPolicies/{access_policy}`. This *is* the resource name.
	// Let's re-evaluate `ParseAccessPolicyExternal`. It expects `organizations/{id}`. This is parsing the PARENT.
	// So, the `AccessPolicyIdentity.String()` should return the fully qualified name if that's what is used for Get/Delete.
	// The Get RPC is `get: "/v1/{name=accessPolicies/*}"`.
	// If `name` is `accessPolicies/123`, this means the resource "123" of type "accessPolicies".
	// It does not mean `organizations/foo/accessPolicies/123`.
	// Okay, the `AccessPolicy` resource name *is* "accessPolicies/{id}". It does not embed its parent unlike some other GCP resources.
	// The `parent` field in the `AccessPolicy` resource specifies its location in hierarchy but is not part of its `name`.
	return fmt.Sprintf("accessPolicies/%s", i.Name)
}

func (i *AccessPolicyIdentity) GetName() string {
	return i.Name
}

func (i *AccessPolicyIdentity) GetParent() AccessPolicyParent {
	return i.AccessPolicyParent
}

func (i *AccessPolicyIdentity) GetKind() string {
	return AccessPolicyKind
}

func (i *AccessPolicyIdentity) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    GroupVersion.Group,
		Version:  GroupVersion.Version,
		Resource: "accesscontextmanageraccesspolicies",
	}
}

func (i *AccessPolicyIdentity) Set(val string) error {
	if val == "" {
		return fmt.Errorf("empty value for identity")
	}
	// Expects "accessPolicies/{name}"
	parts := strings.Split(val, "/")
	if len(parts) == 2 && parts[0] == "accessPolicies" {
		i.Name = parts[1]
		// The parent needs to be set separately or through a different mechanism
		// as it's not part of this specific string format for AccessPolicy name.
		// For now, we leave parent as is or expect it to be uninitialized / handled elsewhere.
		// This Set method is for parsing the resource's own name, not its parent's name.
		// However, the interface might expect to parse the full identity including parent.
		// Given the String() method now returns "accessPolicies/{name}", this Set should parse that.
		// If String() were to return "organizations/{org_id}/accessPolicies/{name}",
		// then Set would need to parse that.
		// Let's assume for now that AccessPolicyParent is already populated or will be from spec.
		return nil
	}
	return fmt.Errorf("invalid identity format, expected accessPolicies/{name}: %s", val)
}

func (i *AccessPolicyIdentity) Type() string {
	return "AccessPolicyIdentity"
}
