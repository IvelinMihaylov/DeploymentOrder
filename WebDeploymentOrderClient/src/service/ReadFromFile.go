package service

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"practice-repo/WebDeploymentOrderClient/src/web/RestControllers"
	"strings"
)

func readPathOfFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give absolute path of file:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	input = strings.TrimRight(input, "\n")
	input = strings.Replace(input, "\\", "\\"+"\\", 0)
	data := readFromFile(input)
	RestControllers.DataHttpController(data)
}

func readFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}
