package config

import (
	"fmt"
	"github.com/odin-cli/odin/pkg/models"
	"github.com/odin-cli/odin/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Create config file",
		Run: func(cmd *cobra.Command, args []string) {
			home, err := os.UserHomeDir()
			if err != nil {
				log.Fatal("Error to load UserHomeDir:", err.Error())
				os.Exit(1)
			}

			configFile := fmt.Sprintf("%v/.odin.config.yaml", home)

			// Check if file config exists
			if _, err := os.ReadFile(configFile); err == nil {
				utils.PrintWarn("Config file %v exists! Can you override?", configFile)
			}

			config := models.Config{
				Templates: map[string]models.Template{
					"default": models.Template{
						Repo: "https://github.com/lbernardo/fx-webserver-example",
						Config: &models.Project{
							Params: []models.ProjectParam{
								{
									Name: "package",
								},
							},
							Actions: []models.ProjectActions{
								{
									Action:  "replace",
									Find:    "github.com/lbernardo/fx-webserver",
									Replace: "{{ .package }}",
								},
							},
						},
					},
				},
			}

			content, _ := yaml.Marshal(&config)

			if err := os.WriteFile(configFile, content, 0755); err != nil {
				utils.PrintError("Error to create configFile! Error: %v", err.Error())
				os.Exit(1)
			}

			utils.PrintSuccess("File %v created", configFile)
			fmt.Println(string(content))
		},
	}
}
