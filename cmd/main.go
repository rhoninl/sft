package main

import (
	"github.com/rhoninl/shifucli/cmd/root"
)

func main() {
	root.RootCmd.Execute()
	// cobra.CheckErr(root.RootCmd.Execute())
}
