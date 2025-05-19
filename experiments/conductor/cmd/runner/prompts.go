package runner

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

//go:embed *.txt
var fs embed.FS

func getTemplate(key string) (*template.Template, error) {
	b, err := fs.ReadFile(key)
	if err != nil {
		return nil, fmt.Errorf("reading embedded file %q: %w", key, err)
	}

	t, err := template.New(key).Parse(string(b))
	if err != nil {
		return nil, fmt.Errorf("parsing template %q: %w", key, err)
	}
	return t, nil
}

type MockgcpGenerateGcloudTestPrompt struct {
	GcloudCommand string
	Group         string
	Resource      string
}

func (o *MockgcpGenerateGcloudTestPrompt) Generate() ([]byte, error) {
	key := "mockgcp-generate-gcloud-test.prompt.txt"
	t, err := getTemplate(key)
	if err != nil {
		return nil, err
	}

	var bb bytes.Buffer
	if err := t.Execute(&bb, o); err != nil {
		return nil, fmt.Errorf("running %q template: %w", key, err)
	}

	return bb.Bytes(), nil
}
