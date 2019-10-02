package utils

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SendRequest(url string, headers map[string]string, resultType interface{}) error {
	var transCfg = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transCfg}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	jsonString := string(bodyBytes)
	return json.Unmarshal([]byte(jsonString), &resultType)
}
