package k8s

import (
	"fmt"
	"net/http"
	"os"

	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func PortForwardPod(namespace, podName, localPort, remotePort string) error {
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
		return fmt.Errorf("failed to create spdy roundtripper: %v", err)
	}

	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())

	ports := []string{fmt.Sprintf("%s:%s", localPort, remotePort)}
	stopChan := make(chan struct{}, 1)
	readyChan := make(chan struct{})

	forwarder, err := portforward.New(dialer, ports, stopChan, readyChan, os.Stdout, os.Stderr)
	if err != nil {
		return fmt.Errorf("failed to create port forwarder: %v", err)
	}

	go func() {
		<-readyChan
		fmt.Printf("Port-forward to pod %s on localhost:%s -> remote port %s is ready.\n", podName, localPort, remotePort)
	}()

	err = forwarder.ForwardPorts()
	if err != nil {
		return fmt.Errorf("failed to forward ports: %v", err)
	}

	return nil
}
