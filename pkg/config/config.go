package config

import (
	"log"

	"github.com/weinmann-emt/prometheus-reconfigure/pkg/store"
	"gopkg.in/yaml.v3"
)

type Config struct {
	content map[string]interface{}
}

func (conf *Config) Create(store store.ConfigStore) error {
	yamlString, err := store.Content()
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(yamlString), &conf.content)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (conf *Config) Update(key string, value string) {

}
