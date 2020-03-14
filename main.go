package main

import (
	"Edwardz43/tgbot/crawl/beauty"
	"Edwardz43/tgbot/message/from"
	"Edwardz43/tgbot/worker"
	"Edwardz43/tgbot/worker/rabbitmqworker"
)

var jobWorker worker.Worker

func main() {
	jobWorker = &rabbitmqworker.Worker{}
	go jobWorker.Do(GetPTTBueaty)
	select {}
	//serve()
}

func GetPTTBueaty(arg ...interface{}) error {

	result := arg[0].(*from.Result)

	if result.Message.Text != "!b" {
		return nil
	}

	crawler := &beauty.Crawler{}
	s := crawler.Get()

	c := &Command{
		ChatID:    result.Message.Chat.ID,
		Text:      s,
		ParseMode: "HTML",
	}

	return send(&c)
}

/*
func serve() {
	http.HandleFunc("/", botHandler)
	log.Fatal(http.ListenAndServe(":5008", nil))
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("READ ERROR : %v", err)
	}

	log.Printf("request %v from proxy\n", string(b))

	w.WriteHeader(200)

	result, err := parseMessage(b)

	if err != nil {
		log.Printf("Parse message error : %v", err)
	}

	if result.Message.Text == "!b" {
		jobWorker.Do(GetPTTBueaty)
	}
}
*/

// func parseMessage(msg []byte) (Result, error) {
// 	u := Result{}
// 	err := json.Unmarshal(msg, &u)
// 	if err != nil {
// 		return Result{}, err
// 	}
// 	return u, nil
// }
