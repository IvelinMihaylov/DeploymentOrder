package service

import (
	"practice-repo/HandlingDeploymentOrder/src/models"
	"strings"
	"testing"
)

func TestDeterminingByDependencyOrder(t *testing.T) {
	tests := []struct {
		name           string
		components     models.Components
		expectedResult string
		err            error
	}{
		{name: "Test case 1",
			components: models.Components{[]*models.Component{
				{"I", []string{}},
				{"V", []string{}},
				{"T", []string{}},
				{"G", []string{"I", "V", "T"}},
				{"K", []string{"G"}},
			}},
			expectedResult: "IVTGK",
			err:            nil,
		},
		{name: "Test case 1",
			components: models.Components{[]*models.Component{
				{"R", []string{"T"}},
				{"P", []string{"T"}},
				{"V", []string{"T"}},
				{"T", []string{}},
				{"Q", []string{"T"}},
			}},
			expectedResult: "TRPVQ",
			err:            nil,
		},
	}

	for _, test := range tests {
		orderComponents, err := DeterminingComponentsOrder(&test.components)
		if err != test.err {
			t.Error("Error check is not passed\n" + err.Error())
		}
		var result strings.Builder
		for _, component := range orderComponents {
			result.WriteString(component.Name)
		}
		if strings.Compare(result.String(), test.expectedResult) != 0 {
			t.Error("Excepted "+test.expectedResult+",but received ", result.String())
		}
	}
}
