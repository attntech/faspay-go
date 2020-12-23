package faspay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SendPost(headers map[string][]string, param interface{}, urlPost string) (result map[string]interface{}, err error) {

	jsonPayload, err := json.Marshal(param)
	if err != nil {
		return
	}

	fmt.Println(strings.NewReader(string(jsonPayload)), " INI PAYLOAD")
	req, _ := http.NewRequest("POST", urlPost, strings.NewReader(string(jsonPayload)))

	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
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

func SendPostOvo(param *OVODirectRequest, urlPost string) (result map[string]interface{}, err error) {

	form := url.Values{
		"trx_id":     {param.TrxID},
		"ovo_number": {param.OVONumber},
		"signature":  {param.Signature},
	}
	req, _ := http.NewRequest("POST", urlPost, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
