package common

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"strings"
)

var client = &http.Client{}

func Get(url string, res interface{}) error {
	request := createRequest(url, "GET", "")
	return exec(request, url, res)
}

func Post(url string, body interface{}, res interface{}) error {
	payload, err := json.Marshal(body)
	if err != nil {
		logrus.Errorf("POST %s error %s")
		return err
	}
	request := createRequest(url, "POST", string(payload))
	return exec(request, url, res)
}

func PostStr(url string, body string, res interface{}) error {
	request := createRequest(url, "POST", body)
	return exec(request, url, res)
}

func Put(url string, body interface{}, res interface{}) error {
	payload, err := json.Marshal(body)
	if err != nil {
		logrus.Errorf("PUT %s error %s")
		return err
	}
	request := createRequest(url, "PUT", string(payload))
	return exec(request, url, res)
}

func exec(request *http.Request, url string, res interface{}) error {
	resp, err := client.Do(request)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logrus.Errorf("POST %s error %s", url, err.Error())
		return err
	}
	respStr, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respStr, res)
	if err != nil {
		logrus.Errorf("POST %s error %s", url, err.Error())
	}
	return err
}

func createRequest(url string, method string, body string) *http.Request {
	request, e := http.NewRequest(method, url, strings.NewReader(body))
	if e != nil {
		logrus.Errorf("create request error %s %s body is %s", url, method, body)
		return nil
	}
	request.Header.Set("Content-Type", ContentTypeJson)
	return request
}
