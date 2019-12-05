package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Env struct
type Env struct {
	Mail struct {
		Hostname string `yaml:"hostname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		From     string `yaml:"from"`
	}
}

// Vars configuration
func (env Env) Vars() (Env, error) {
	f, err := os.Open("env.yml")
	if err != nil {
		return env, err
	}

	var e Env
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}
