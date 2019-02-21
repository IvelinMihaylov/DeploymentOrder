package service

import (
	"fmt"
	"strings"
	"testing"
)

func TestYamlFileReader(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedResult string
		err            error
	}{
		{
			name:           "Read yaml file.",
			path:           "../../../../resource/YamlFileReader_Test.yaml",
			expectedResult: "ADECBDBEFB",
			err:            nil,
		},
		{
			name:           "Read empty yaml file.",
			path:           "../../../../resource/YamlFileReader_Test_EmptyFile.yaml",
			expectedResult: "nil",
			err:            fmt.Errorf("Not found components!")},
	}
	for _, test := range tests {
		components, err := ReadComponentsFromFile(test.path)
		if err != nil {
			if strings.Compare(test.err.Error(), err.Error()) != 0 {
				t.Errorf(test.err.Error())
			}
		}
		var result strings.Builder
		if err == nil {
			for _, component := range components.Components {
				result.WriteString(component.Name)
				for _, require := range component.Requirements {
					result.WriteString(require)
				}
			}
			resultOfComponents := result.String()
			if strings.Compare(resultOfComponents, test.expectedResult) != 0 {
				t.Errorf("Excepcted "+test.expectedResult+",but received ", resultOfComponents)
			}
		}
	}

}
