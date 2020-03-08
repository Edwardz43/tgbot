package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var token string

func init() {
	token = getToken()
}

func getUpdates() ([]byte, error) {

	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/%s/getUpdates", token))

	if err != nil {
		log.Panic(err)
	}

	form, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return form, nil
}
