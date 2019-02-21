package service

import (
	"fmt"
	"practice-repo/HandlingDeploymentOrder/src/models"
	"strings"
)

func DeterminingComponentsOrder(components *models.Components) ([]*models.Component, error) {

	var orderedComponents []*models.Component
	var err error

	for _, currentComponent := range components.Components {

		if !contains(orderedComponents, currentComponent) {

			findRequireComponentByName, err := foundComponentByName(currentComponent.Name, components)

			if err != nil {
				return orderedComponents, err
			}

			orderedComponents, err = determiningOrderedByRequiresOfComponents(findRequireComponentByName, components, orderedComponents, []*models.Component{findRequireComponentByName})

			if err != nil {
				return orderedComponents, err
			}
		}
	}

	return orderedComponents, err
}

func determiningOrderedByRequiresOfComponents(component *models.Component, components *models.Components, orderedComponents []*models.Component, circleDefender []*models.Component) ([]*models.Component, error) {

	var err error

	if cap(component.Requirements) != 0 {

		for _, requireComponentName := range component.Requirements {

			findRequireComponentByName, err := foundComponentByName(requireComponentName, components)

			if err != nil {
				return orderedComponents, err
			}

			if contains(circleDefender, findRequireComponentByName) {
				for index, value := range circleDefender{
					if strings.EqualFold(value.Name,findRequireComponentByName.Name) {
						circleDefender = append(circleDefender, findRequireComponentByName)
						return circleDefender[index:], fmt.Errorf("Recursive circular dependencies!\n")
					}
				}
			}

			if !contains(orderedComponents, findRequireComponentByName) {

				circleDefender = append(circleDefender, findRequireComponentByName)
				orderedComponents, err = determiningOrderedByRequiresOfComponents(findRequireComponentByName, components, orderedComponents, circleDefender)

				if err != nil {
					return orderedComponents, err
				}
			}
		}
	}
	if !contains(orderedComponents, component) {
		orderedComponents = append(orderedComponents, component)
	}
	return orderedComponents, err
}

func foundComponentByName(name string, components *models.Components) (*models.Component, error) {

	for _, component := range components.Components {

		if strings.EqualFold(component.Name, name) {
			return component, nil
		}
	}
	return nil, fmt.Errorf("Not found required component from list of components.")
}

func contains(a []*models.Component, x *models.Component) bool {

	for _, n := range a {

		if strings.EqualFold(n.Name, x.Name) {
			return true
		}
	}
	return false
}
