package RestControllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMessageInSlackChannel(url string, message string) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(message)))
	if err != nil {
		fmt.Println(err.Error())
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
	} else {
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		response.Body.Close()
	}
}
