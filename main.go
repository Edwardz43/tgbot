package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var form []byte
	var err error
	var update Update

	if form, err = getUpdates(); err != nil {
		log.Printf("Get update error : %v", err)
	}

	if update, err = parseMessage(form); err != nil {
		log.Printf("Parse message error : %v", err)
	}
	if !update.OK {
		log.Println("Get update failed")
	}

	c := &Conn{}
	connStr := getMongoConnStr()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	c.setup(client, "Test", "tgbot")

	for _, v := range update.ResultList {

		filter := bson.M{
			"updateid": v.UpdateID,
		}
		_, err := c.upsert(filter, v)
		if err != nil {
			log.Printf("Insert error :%v", err)
			continue
		}
		log.Println("Update success")
	}
}

func serve() {
	http.HandleFunc("/bot", botHandler)
	http.ListenAndServe(":8080", nil)
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("READ ERROR : %v", err)
	}

	log.Printf("request %v", string(b))

	w.WriteHeader(200)
	// w.Write([]byte("hello"))
}

func parseMessage(msg []byte) (Update, error) {
	u := Update{}
	err := json.Unmarshal(msg, &u)
	if err != nil {
		return Update{}, err
	}
	return u, nil
}
