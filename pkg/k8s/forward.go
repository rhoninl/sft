package k8s

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rhoninl/sft/pkg/utils/logger"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func PortForwardPod(ctx context.Context, namespace, podName, localPort, remotePort string, readyChan chan struct{}) error {
	clientset, config, err := NewClientSet()
	if err != nil {
		return fmt.Errorf("failed to create clientset: %w", err)
	}

	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(namespace).
		Name(podName).
		SubResource("portforward")

	transport, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return fmt.Errorf("failed to create spdy roundtripper: %w", err)
	}

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())

	ports := []string{fmt.Sprintf("%s:%s", localPort, remotePort)}

	forwarder, err := portforward.New(dialer, ports, ctx.Done(), readyChan, os.Stdout, os.Stderr)
	if err != nil {
		return fmt.Errorf("failed to create port forwarder: %w", err)
	}

	go func() {
		<-readyChan
		fmt.Printf("Port-forward to pod %s on localhost:%s -> remote port %s is ready.\n", podName, localPort, remotePort)
	}()

	go func() {
		if err := forwarder.ForwardPorts(); err != nil {
			logger.Printf("Error: Failed to forward ports: %v", err)
		}
	}()

	return nil
}
