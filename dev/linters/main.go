package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/linters/loglint"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/linters/jsonunmarshalreuse"
)

func main() {
	multichecker.Main(
		loglint.Analyzer,
		jsonunmarshalreuse.Analyzer,
	)
}
