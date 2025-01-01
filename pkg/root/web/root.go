package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rhoninl/sft/pkg/root/web/apis/shifu"
	"github.com/rhoninl/sft/pkg/root/web/devices"
	"github.com/rhoninl/sft/pkg/root/web/middleware"
	"github.com/rhoninl/sft/pkg/root/web/server"
	"github.com/rhoninl/sft/pkg/utils/logger"
	"github.com/spf13/cobra"
)

var (
	port       string
	serverOnly bool
)

func init() {
	WebCmd.Flags().BoolVarP(&serverOnly, "server-only", "s", false, "Only run the server")
	WebCmd.Flags().StringVarP(&port, "port", "p", "34550", "Port to run the web server on")
}

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web server",
	Long:  `Run a web server that listens on a specified port`,
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		muxNeedInstallShifu := http.NewServeMux()

		muxNeedInstallShifu.HandleFunc("/", HomeHandler)
		muxNeedInstallShifu.HandleFunc("/device", devices.DevicesHandler)
		muxNeedInstallShifu.HandleFunc("/device/{device_name}", devices.DetailHandler)

		mux.Handle("/", middleware.InstallChecker(muxNeedInstallShifu))
		mux.HandleFunc("/api/shifu/checker", shifu.InstallChecker)
		mux.HandleFunc("/api/shifu/install", shifu.InstallShifu)

		logger.Printf("Starting web server on localhost:%s\n", port)

		go func() {
			portInt, err := strconv.Atoi(port)
			if err != nil {
				logger.Printf("Failed to convert port to int: %v\n", err)
			}
			// if err := http.ListenAndServe(":"+port, mux); err != nil {
			// 	logger.Printf("Failed to start web server: localhost:%v\n", err)
			// }
			server.StartGRPCServer(portInt)
		}()

		if !serverOnly {
			// Open the browser to localhost:port
			url := fmt.Sprintf("http://localhost:%s/device", port)
			// if err := browser.Open(url); err != nil {
			// 	logger.Printf("Failed to open browser: %v\n", err)
			// }
			fmt.Println(url)
		}

		select {}
	},
}
