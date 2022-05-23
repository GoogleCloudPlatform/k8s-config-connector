//go:build integration
// +build integration

package create

import (
	"regexp"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
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
	// Service resources should specify the service to enable (e.g.
	// pubsub.googleapis.com) via spec.resourceID instead of metadata.name.
	// Output a targeted error message for this case since it is an easy
	// mistake to make.
	if u.GetKind() == "Service" {
		if strings.HasSuffix(u.GetName(), ".com") {
			t.Fatalf("invalid metadata.name value '%v' for Service resource in sample '%v': "+
				"use %v instead of metadata.name to specify the service to enable",
				u.GetName(), sampleName, k8s.ResourceIDFieldPath)
		}
	}

	allowedNameFragments := []string{"sample", "dep"}
	for _, nf := range allowedNameFragments {
		if strings.Contains(u.GetName(), nf) {
			return
		}
	}
	// In addition to this naming scheme following the sample guidelines, the
	// create sample test looks for either "sample" or "dep" to "uniqify" the
	// name of a sample
	t.Errorf("invalid metadata.name value '%v' in sample '%v': must contain one of {%v} to be valid",
		u.GetName(), sampleName, strings.Join(allowedNameFragments, ","))
}
