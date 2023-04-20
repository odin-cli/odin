package create

import (
	"github.com/odin-cli/odin/pkg/app"
	"github.com/odin-cli/odin/utils"
	"github.com/spf13/cobra"
	"os"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [name]",
		Example: "create auth --template default\ncreate auth --template https://github.com/owner/template",
		Short:   "Create microservice by template",
		Run:     run,
	}

	cmd.Flags().StringP("template", "t", "default", "Create project based of template")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		utils.PrintError("[name] is required")
		os.Exit(1)
	}
	config := utils.ReadConfig()
	name := args[0]
	template, _ := cmd.Flags().GetString("template")

	odinApp := app.New(config)
	odinApp.CreateProject(template, name)
}
