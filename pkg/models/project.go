package models

type Project struct {
	ApiVersion string           `yaml:"apiVersion"`
	Params     []ProjectParam   `yaml:"params"`
	Actions    []ProjectActions `yaml:"actions"`
}

type ProjectParam struct {
	Name string `yaml:"name"`
}

type ProjectActions struct {
	Action  string `yaml:"action"`
	Find    string `yaml:"find,omitempty"`
	Replace string `yaml:"replace,omitempty"`
}
