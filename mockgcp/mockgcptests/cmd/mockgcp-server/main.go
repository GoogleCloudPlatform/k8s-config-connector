package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockgcptests"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	if err := run(context.Background()); err != nil {

		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}

func run(ctx context.Context) error {
	listen := "127.0.0.1:0"

	errChan := make(chan error, 8)

	flag.StringVar(&listen, "listen", listen, "endpoint on which to run proxy")

	klog.InitFlags(nil)

	flag.Parse()

	var kubeClient client.Client // TODO: We should replace this, it didn't work
	mockGCP, err := mockgcp.NewMockRoundTripper(ctx, kubeClient, storage.NewInMemoryStorage())
	if err != nil {
		return fmt.Errorf("building mockgcp: %v", err)
	}

	go func() {
		if err := mockGCP.Run(ctx); err != nil {
			errChan <- fmt.Errorf("error from mockgcp: %w", err)
		}
	}()

	roundTripper := http.RoundTripper(mockGCP)

	httpClient := &http.Client{Transport: roundTripper}

	proxy := mockgcptests.NewProxy(httpClient)

	proxyListener, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("net.Listen failed: %w", err)
	}

	httpServer := &http.Server{}
	httpServer.Handler = proxy

	go func() {
		if err := httpServer.Serve(proxyListener); err != nil {
			if err != http.ErrServerClosed {
				errChan <- fmt.Errorf("error from proxy server: %w", err)
			}
		}
	}()

	proxyEndpoint := proxyListener.Addr().(*net.TCPAddr)

	gcloudConfig := proxy.BuildGcloudConfig(proxyEndpoint, mockGCP)
	for _, envvar := range gcloudConfig.EnvVars {
		fmt.Fprintf(os.Stdout, "export %v\n", envvar)
	}

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			errChan <- err
		}
	}()

	ret := <-errChan

	httpServer.Close()

	return ret
}
