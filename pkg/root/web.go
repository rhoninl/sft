package root

import (
	"fmt"
	"net/http"

	"github.com/rhoninl/sft/pkg/k8s"
	"github.com/rhoninl/sft/pkg/utils/browser"
	"github.com/rhoninl/sft/template"
	"github.com/spf13/cobra"
)

var (
	port string
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web server",
	Long:  `Run a web server that listens on a specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/device", func(w http.ResponseWriter, r *http.Request) {
			devices, err := k8s.GetEdgedevices()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			template.RenderDevicesTemplate(w, devices)
		})

		http.HandleFunc("/device/{device_name}", func(w http.ResponseWriter, r *http.Request) {
			device, err := k8s.GetAllByDeviceName(r.PathValue("device_name"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			template.RenderDetailTemplate(w, device)
		})

		fmt.Printf("Starting web server on localhost:%s\n", port)

		go func() {
			if err := http.ListenAndServe(":"+port, nil); err != nil {
				fmt.Printf("Failed to start web server: localhost:%v\n", err)
			}
		}()

		// Open the browser to localhost:port
		url := fmt.Sprintf("http://localhost:%s/device", port)
		if err := browser.Open(url); err != nil {
			fmt.Printf("Failed to open browser: %v\n", err)
		}

		select {}
	},
}

func init() {
	webCmd.Flags().StringVarP(&port, "port", "p", "34550", "Port to run the web server on")
}
