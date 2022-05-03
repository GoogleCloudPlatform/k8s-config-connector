//go:build integration
// +build integration

package create

import (
	"regexp"
	"strings"
	"testing"

	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var matchEverythingRegex = regexp.MustCompile(".*")

func TestNames(t *testing.T) {
	samples := loadSamplesOntoUnstructs(t, matchEverythingRegex)
	for _, s := range samples {
		for _, r := range s.Resources {
			validateResourceName(t, s.Name, r)
		}
	}
}

func TestLicenses(t *testing.T) {
	beginsWithCopyrightRegex := regexp.MustCompile("^# Copyright 20[0-9]{2} Google LLC.*")
	samples := mapSampleNamesToFilePaths(t, matchEverythingRegex)
	for sampleName, files := range samples {
		for _, f := range files {
			b := testcontroller.ReadFileToBytes(t, f)
			if !beginsWithCopyrightRegex.Match(b) {
				t.Errorf("file '%v' in sample '%v' does not contain a license header", f, sampleName)
			}
		}
	}
}

func validateResourceName(t *testing.T, sampleName string, u *unstructured.Unstructured) {
	if !resourceNameMustAdhereToSampleGuidelines(u) {
		return
	}
	allowedNameFragments := []string{"sample", "dep"}
	for _, nf := range allowedNameFragments {
		if strings.Contains(u.GetName(), nf) {
			return
		}
	}
	// in addition to this naming scheme following the sample guidelines, the create sample test is looking for either
	// "sample" or "dep" to "uniqify" the name of a sample
	t.Errorf("invalid resource name '%v' in sample '%v': resource name must contain one of {%v} to be valid",
		u.GetName(), sampleName, strings.Join(allowedNameFragments, ","))
}

func resourceNameMustAdhereToSampleGuidelines(u *unstructured.Unstructured) bool {
	// resources that do not have to adhere to the sample guidelines, each resource should have a comment as to why it
	// is OK if it doesn't adhere
	disabledResources := map[string]bool{
		// Service names MUST be the actual URI of their service, i.e. pubsub.googleapis.com
		"Service": true,
	}
	_, ok := disabledResources[u.GroupVersionKind().Kind]
	return !ok
}
