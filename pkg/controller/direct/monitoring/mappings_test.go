package monitoring

import (
	"os"
	"testing"
)

func TestMapping(t *testing.T) {
	os.Setenv("CHECK_COVERAGE", "1")
	errs := dashboardMapping.Validate()
	for _, err := range errs {
		t.Errorf("error mapping: %v", err)
	}
}
