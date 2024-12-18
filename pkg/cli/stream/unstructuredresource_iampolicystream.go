// Copyright 2022 Google LLC
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

package stream

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/iamresource"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type IAMFormat string

const (
	IAMFormatPartialPolicy = "partialpolicy"
	IAMFormatPolicy        = "policy"
	IAMFormatPolicyMember  = "policymember"
)

type IAMClient interface {
	SupportsIAM(unstructured *unstructured.Unstructured) (bool, error)
	GetPolicy(ctx context.Context, unstructured *unstructured.Unstructured) (*v1beta1.IAMPolicy, error)
}

type UnstructuredResourceAndIAMPolicyStream struct {
	unstructStream          UnstructuredStream
	iamClient               IAMClient
	iamFormat               IAMFormat
	filterDeletedIAMMembers bool
	nextIAMResources        []*unstructured.Unstructured
}

func NewUnstructuredResourceAndIAMPolicyStream(unstructuredStream UnstructuredStream, iamClient IAMClient, iamFormat IAMFormat, filterDeletedIAMMembers bool) *UnstructuredResourceAndIAMPolicyStream {
	iamPolicyStream := UnstructuredResourceAndIAMPolicyStream{
		unstructStream:          unstructuredStream,
		iamClient:               iamClient,
		iamFormat:               iamFormat,
		filterDeletedIAMMembers: filterDeletedIAMMembers,
	}
	return &iamPolicyStream
}

// This function returns the next unstructured as follows:
// 1. If there at least one value in the iam resources slice, reduce the size of the slice by one and return the value at the head of the slice
// 2. Else, get the next resource, if it supports iam policy fetch the iam policy, if the policy is non-empty save it into the iam resources slice, then return the resource
func (s *UnstructuredResourceAndIAMPolicyStream) Next(ctx context.Context) (*unstructured.Unstructured, error) {
	if len(s.nextIAMResources) > 0 {
		result := s.nextIAMResources[0]
		s.nextIAMResources = s.nextIAMResources[1:]
		return result, nil
	}
	resourceUnstruct, err := s.unstructStream.Next(ctx)
	if err != nil {
		if !errors.Is(err, io.EOF) {
			err = fmt.Errorf("error getting next unstruct: %w", err)
		}
		return nil, err
	}
	// if any error occurs in this function then return error and 'drop' the resourceUnstruct: resources are incomplete
	// without their associated IAMPolicy.
	if err := s.fillNextIAMPolicyIfSupportedAndIfNonEmpty(resourceUnstruct); err != nil {
		return nil, err
	}
	return resourceUnstruct, nil
}

func (s *UnstructuredResourceAndIAMPolicyStream) fillNextIAMPolicyIfSupportedAndIfNonEmpty(u *unstructured.Unstructured) error {
	hasIAMSupport, err := s.iamClient.SupportsIAM(u)
	if err != nil {
		return fmt.Errorf("error determining if resource %v supports iam: %w", u.GroupVersionKind(), err)
	}
	if !hasIAMSupport {
		return nil
	}

	return s.fillNextIAMPolicyIfNonEmpty(u)
}

func (s *UnstructuredResourceAndIAMPolicyStream) fillNextIAMPolicyIfNonEmpty(u *unstructured.Unstructured) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	iamPolicy, err := s.iamClient.GetPolicy(ctx, u)
	if err != nil {
		return fmt.Errorf("error getting iam policy for '%v' with name '%v': %w",
			u.GetKind(), u.GetName(), err)
	}
	if s.filterDeletedIAMMembers {
		filterOutDeletedMembers(iamPolicy)
	}
	// ignore empty policies
	if len(iamPolicy.Spec.Bindings) == 0 {
		return nil
	}
	iamPolicy.SetName(fmt.Sprintf("%v-%v", iamPolicy.GetName(), "iampolicy"))
	res, err := s.policyToUnstructureds(iamPolicy)
	if err != nil {
		return err
	}
	s.nextIAMResources = res
	return nil
}

func filterOutDeletedMembers(iamPolicy *v1beta1.IAMPolicy) {
	newBindings := make([]v1beta1.IAMPolicyBinding, 0)
	for _, binding := range iamPolicy.Spec.Bindings {
		members := make([]v1beta1.Member, 0)
		for _, m := range binding.Members {
			if !isDeletedMembers(m) {
				members = append(members, m)
			}
		}
		if len(members) != 0 {
			binding.Members = members
			newBindings = append(newBindings, binding)
		}
	}
	iamPolicy.Spec.Bindings = newBindings
}

func isDeletedMembers(member v1beta1.Member) bool {
	return strings.HasPrefix(string(member), "deleted:")
}

func (s *UnstructuredResourceAndIAMPolicyStream) policyToUnstructureds(iamPolicy *v1beta1.IAMPolicy) ([]*unstructured.Unstructured, error) {
	switch s.iamFormat {
	case IAMFormatPartialPolicy:
		return policyToPartialPolicyUnstructureds(iamPolicy)
	case IAMFormatPolicy:
		return policyToPolicyUnstructureds(iamPolicy)
	case IAMFormatPolicyMember:
		return policyToPolicyMemberUnstructureds(iamPolicy)
	default:
		panic(fmt.Sprintf("unimplemented iam format: %v", s.iamFormat))
	}
}

func policyToPartialPolicyUnstructureds(iamPolicy *v1beta1.IAMPolicy) ([]*unstructured.Unstructured, error) {
	partialPolicy := iamresource.ConvertIAMPolicyToIAMPartialPolicy(iamPolicy)
	u, err := metaObjectToUnstructured(partialPolicy)
	if err != nil {
		return nil, err
	}
	return []*unstructured.Unstructured{u}, nil
}

func policyToPolicyUnstructureds(iamPolicy *v1beta1.IAMPolicy) ([]*unstructured.Unstructured, error) {
	u, err := metaObjectToUnstructured(iamPolicy)
	if err != nil {
		return nil, err
	}
	return []*unstructured.Unstructured{u}, nil
}

func policyToPolicyMemberUnstructureds(iamPolicy *v1beta1.IAMPolicy) ([]*unstructured.Unstructured, error) {
	iamPolicy.SetName(fmt.Sprintf("%vmember", iamPolicy.GetName()))
	policyMembers := iamresource.SplitPolicy(iamPolicy)
	results := make([]*unstructured.Unstructured, 0, len(policyMembers))
	for _, pm := range policyMembers {
		u, err := metaObjectToUnstructured(&pm)
		if err != nil {
			return nil, err
		}
		results = append(results, u)
	}
	return results, nil
}

func metaObjectToUnstructured(o metav1.Object) (*unstructured.Unstructured, error) {
	return k8s.MarshalObjectAsUnstructured(o)
}
