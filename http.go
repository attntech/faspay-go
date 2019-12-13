package faspay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SendPost(headers map[string][]string, param interface{}, url string) (result map[string]interface{}, err error) {

	jsonPayload, err := json.Marshal(param)
	if err != nil {
		return
	}

	fmt.Println(strings.NewReader(string(jsonPayload)), " INI PAYLOAD")
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jsonPayload)))

	if len(headers) > 0 {
		req.Header = headers
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return
	}
	return
}
