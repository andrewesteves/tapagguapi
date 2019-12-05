package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// EnvConfig struct
type EnvConfig struct {
	Mail struct {
		Hostname string `yaml:"hostname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		From     string `yaml:"from"`
	}
	Idiom struct {
		Lang string `yaml:"lang"`
	}
}

// Vars configuration
func (env EnvConfig) Vars() (EnvConfig, error) {
	f, err := os.Open("env.yml")
	if err != nil {
		return env, err
	}

	var e EnvConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}
