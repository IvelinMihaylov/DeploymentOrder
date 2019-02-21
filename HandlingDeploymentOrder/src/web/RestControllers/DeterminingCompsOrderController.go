package RestControllers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"practice-repo/HandlingDeploymentOrder/src/models"
	"practice-repo/HandlingDeploymentOrder/src/service"
	"strings"
	"time"
)

func OrderedComps(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	var components models.Components
	err = yaml.Unmarshal(body, &components)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	orderComponents, err := service.DeterminingComponentsOrder(&components)

	var nameOfOrderedComponents []string
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	for _, component := range orderComponents {
		nameOfOrderedComponents = append(nameOfOrderedComponents, component.Name)
	}
	response := "{\n\"icon_emoji\": \":ghost:\",\n\"text\": \"Result of operation in " + time.Now().Format("2006-01-02 15:04:05") + " is:\n" + strings.Join(nameOfOrderedComponents, " -> ") + "\n\"\n}"
	SendMessageInSlackChannel("https://hooks.slack.com/services/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", (response))
	fmt.Fprintf(w, strings.Join(nameOfOrderedComponents, " -> "))
}
