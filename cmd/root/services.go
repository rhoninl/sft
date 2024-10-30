package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
}

var serviceCmd = &cobra.Command{
	Use:     "services",
	Aliases: []string{"svc", "service"},
	Short:   "deviceShifu service addresses",
	Long:    `show deviceShifu service addresses`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WIP - service addresses")
	},
}
