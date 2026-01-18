package jsonunmarshalreuse_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/linters/jsonunmarshalreuse"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, jsonunmarshalreuse.Analyzer, "a")
}
