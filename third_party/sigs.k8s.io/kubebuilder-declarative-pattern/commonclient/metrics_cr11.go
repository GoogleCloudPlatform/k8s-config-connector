//go:build controllerruntime_11 || controllerruntime_12 || controllerruntime_13 || controllerruntime_14 || controllerruntime_15

package commonclient

import (
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
)

// SetMetricsBindAddress sets the metrics address on options independent of
// manager options version
func SetMetricsBindAddress(options *ctrl.Options, bindAddress string) error {
	if options == nil {
		return fmt.Errorf("unable to set metrics bind address on non-existent manager options")
	}
	options.MetricsBindAddress = bindAddress
	return nil
}
