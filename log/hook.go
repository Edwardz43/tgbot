package log

import (
	"Edwardz43/tgbot/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	index string
	url   string
)

func init() {
	index = config.GetESIndex()
	url = fmt.Sprintf("%s/%s/_doc", config.GetESURL(), index)
}

// Emit emits a log to elasticsearch host
func Emit(c *Content) error {

	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(c)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes.Bytes()))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("send ES log err : %v\n", err)
	}

	if resp == nil {
		return fmt.Errorf("send message no response")
	}

	b := resp.Body

	msg, err := ioutil.ReadAll(b)
	if err != nil {
		return err
	}

	log.Println(string(msg))

	return nil
}

// Content is a log content data model
type Content struct {
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
	Line     int       `json:"line"`
	FileName string    `json:"file_name"`
	Function string    `json:"function"`
}
