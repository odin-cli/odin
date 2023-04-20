package app

import (
	"fmt"
	"github.com/odin-cli/odin/pkg/models"
	"github.com/odin-cli/odin/utils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type OdinApp struct {
	config *models.Config
	logger *zap.Logger
}

func New(config *models.Config) *OdinApp {
	return &OdinApp{
		config: config,
		logger: utils.NewLogger(),
	}
}

func (o *OdinApp) CreateProject(templateName string, name string) {
	template, ok := o.config.Templates[templateName]
	if !ok {
		o.logger.Warn(templateName + " not found in $HOME/.odin.yaml")
		template = models.Template{
			Local: templateName,
		}
		if strings.Contains(templateName, "https://") {
			template.Local = templateName
		}
	}

	if template.Local != "" {
		o.logger.Info("Local template")
		o.createProjectLocal(template.Local, name)
	} else {
		o.logger.Info("Remote template")
		o.createProjectRemote(template.Repo, name)
	}

	project := o.readProject(name, template)
	params := map[string]string{
		"name": name,
	}
	if project != nil {
		for _, param := range project.Params {
			fmt.Printf("Set value to '%v': ", param.Name)
			var value string
			fmt.Scanln(&value)
			params[param.Name] = value
		}

		replaces := map[string]string{}

		for _, action := range project.Actions {
			switch action.Action {
			case "replace":
				replaces[action.Find] = utils.GetTemplateValues(action.Replace, params)
			}
		}
		if len(replaces) != 0 {
			o.replace(name, replaces)
		}
	}
	o.logger.Info("DONE!")
}

func (o *OdinApp) replace(name string, replaces map[string]string) {
	o.logger.Info("Replacing...")
	files := utils.GetFilesInDir(name)
	for _, file := range files {
		content, _ := os.ReadFile(file)
		contentStr := string(content)
		for find, replace := range replaces {
			contentStr = strings.ReplaceAll(contentStr, find, replace)
		}
		os.WriteFile(file, []byte(contentStr), 0755)
	}
}

func (o *OdinApp) createProjectRemote(template string, name string) {
	o.logger.Info("Clonning project")
	utils.Command("git", "clone", template, name)
	utils.Command("rm", "-rf", fmt.Sprintf("%v/.git", name))
}

func (o *OdinApp) createProjectLocal(template string, name string) {
	o.logger.Info("Copy template")
	utils.Command("cp", "-r", template, name)
}

func (o *OdinApp) readProject(name string, template models.Template) *models.Project {
	var project *models.Project
	if template.Config != nil {
		// Configuration odin project in odin.config.yaml
		return template.Config
	}
	content, err := os.ReadFile(fmt.Sprintf("%v/.odin.yaml", name))
	if err != nil {
		o.logger.Warn(".odin.yaml not found")
		return nil
	}
	if err := yaml.Unmarshal(content, &project); err != nil {
		o.logger.Error("error to read .odin.yaml", zap.Error(err))
		return nil
	}
	return project
}
