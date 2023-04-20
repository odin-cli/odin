package config

type Config struct {
	Server Server `json:"server" yaml:"server"`
	Log    Log    `json:"log" yaml:"log"`
}

type Server struct {
	Port int64 `json:"server" yaml:"port"`
}

type Log struct {
	Mode string `json:"mode" yaml:"mode"`
}
