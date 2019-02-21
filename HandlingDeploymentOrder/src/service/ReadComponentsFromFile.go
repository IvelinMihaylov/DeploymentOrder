package service

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"practice-repo/HandlingDeploymentOrder/src/models"
)

func ReadComponentsFromFile(path string) (*models.Components, error) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var components *models.Components
	err = yaml.Unmarshal(data, &components)

	if components == nil {
		return components, fmt.Errorf("Not found components!")
	}
	return components, err
}
