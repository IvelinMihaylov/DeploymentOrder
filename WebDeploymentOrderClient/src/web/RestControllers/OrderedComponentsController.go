package RestControllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DataHttpController(data []byte) {
	request, err := http.NewRequest("POST", "http://localhost:8081/", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err.Error())
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		response.Body.Close()
		fmt.Println(string(data))
	}
}
