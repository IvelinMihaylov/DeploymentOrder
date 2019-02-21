package service

import (
	"bufio"
	"fmt"
	"os"
)

func PickUpGetDataMode() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose type of read data. NOTE: Pick up with number.")
	inputDataIsReady := true
	for ; inputDataIsReady; {
		fmt.Println("1. Read from file")
		fmt.Println("2. Read from console")
		fmt.Println("3. Auto generated YAML files with Components")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}

		switch input {

		case "1\n":
			readPathOfFile()
			inputDataIsReady = false

		case "2\n":
			readFromConsole()
			inputDataIsReady = false

		case "3\n":
			GeneratedYamlWithComponents()
			inputDataIsReady = false
		}

	}
}
