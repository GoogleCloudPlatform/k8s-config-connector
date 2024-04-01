package livestate

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
)

// FetchLiveKCCState tries to fetch the ConfigConnector (CC) resource and the ConfigConnectorContext (CCC)
// for the resource's namespace if running in Namespaced mode. It ignores not found errors for CC fetching
// but errors out if KCC is running in Namespaced mode and no CCC is found for the namespace of the resource.
func FetchLiveKCCState(ctx context.Context, c client.Client, resourceNN types.NamespacedName) (v1beta1.ConfigConnector, v1beta1.ConfigConnectorContext, error) {
	var cc v1beta1.ConfigConnector
	if err := c.Get(ctx, types.NamespacedName{
		Name: k8s.ConfigConnectorAllowedName,
	}, &cc); err != nil {
		if errors.IsNotFound(err) {
			// if no CC exists, then by definition, KCC cannot be running in namespaced mode;
			return v1beta1.ConfigConnector{}, v1beta1.ConfigConnectorContext{}, nil
		}
		return v1beta1.ConfigConnector{}, v1beta1.ConfigConnectorContext{}, err
	}

	if cc.Spec.Mode == k8s.NamespacedMode {
		var ccc v1beta1.ConfigConnectorContext
		if err := c.Get(ctx, types.NamespacedName{
			Name:      k8s.ConfigConnectorContextAllowedName,
			Namespace: resourceNN.Namespace,
		}, &ccc); err != nil {

			// this should not happen but if we attempt to actuate a resource
			// AND we are running in namespaced mode, not finding a CCC in that namespace
			// is an error in the assumptions that KCC has (i.e. that there is a CCC defined
			// that actively manages resources in that namespace).
			return cc, v1beta1.ConfigConnectorContext{}, err
		}
		return cc, ccc, nil
	}

	return cc, v1beta1.ConfigConnectorContext{}, nil
}
