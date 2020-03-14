package main

import (
	"Edwardz43/tgbot/config"
	"Edwardz43/tgbot/message/to"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Command struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

var token string

func init() {
	token = config.GetToken()
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

func send(data interface{}) error {

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(data)

	u := fmt.Sprintf("https://api.telegram.org/%s/sendMessage", token)

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(reqBodyBytes.Bytes()))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("send message err : %v\n", err)
	}

	if resp == nil {
		log.Println("send message no response")
		return fmt.Errorf("send message no response")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var s to.Send

	json.Unmarshal(body, &s)

	log.Println("response from tgbot : ok =", s.OK)

	return nil
}
