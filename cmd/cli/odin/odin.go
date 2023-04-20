package main

import (
	"github.com/odin-cli/odin/cmd/commands/config"
	"github.com/odin-cli/odin/cmd/commands/create"
	"github.com/spf13/cobra"
)

var _commands = []*cobra.Command{
	config.New(),
	create.New(),
}

func main() {
	var odinCmd = &cobra.Command{
		Use:   "odin",
		Short: "",
		Long:  "",
	}
	odinCmd.AddCommand(_commands...)
	if err := odinCmd.Execute(); err != nil {

	}
}
