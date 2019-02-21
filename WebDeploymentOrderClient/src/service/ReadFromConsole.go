package service

import (
	"bufio"
	"fmt"
	"os"
	"practice-repo/WebDeploymentOrderClient/src/web/RestControllers"
	"strings"
)

func readFromConsole() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Тhe program will stop to read after you are writing an END!")
	var data strings.Builder
	for ; ; {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		if strings.EqualFold(strings.ToLower(input), "end\n") {
			data := []byte(data.String())
			RestControllers.DataHttpController(data)
			break
		}
		data.WriteString(input)
	}
}
