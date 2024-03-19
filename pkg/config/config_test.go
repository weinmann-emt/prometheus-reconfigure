package config_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weinmann-emt/prometheus-reconfigure/pkg/config"
	"gopkg.in/yaml.v3"
)

func TestCreate(t *testing.T) {
	store := MockedStore{}
	dut := config.Config{}
	err := dut.Create(store)
	if err != nil {
		fmt.Println(err)
		assert.Equal(t, nil, err, "Create")
	}
}

func TestSomething(t *testing.T) {
	yamlString := `
name: Amit Kumar
age: 34
address:
  city: Chennai
  state: TN
`
	var data map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		fmt.Println(err)
		assert.Equal(t, nil, err, "Error during parse")
	}
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("City:", data["address"].(map[string]interface{})["city"])
	assert.Equal(t, true, true, "No equal")
}
