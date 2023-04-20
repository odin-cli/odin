package models

type Config struct {
	Templates map[string]Template `yaml:"templates" json:"templates"`
	Publish   *Publish            `yaml:"publish,omitempty" json:"publish,omitempty"`
}

type Template struct {
	Local  string   `yaml:"local,omitempty" json:"local,omitempty"`
	Repo   string   `yaml:"repo,omitempty" json:"repo,omitempty"`
	Config *Project `yaml:"config" json:"config"`
}

type Publish struct {
	Provider    string `yaml:"provider" json:"provider"`
	GitlabToken string `yaml:"gitlab_token" json:"gitlab_token"`
	GithubToken string `yaml:"github_token" json:"github_token"`
	Repo        string `yaml:"repo" json:"repo"`
}
