package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Chat struct {
	Id    int
	Title string
	Tyoe  string
}
type Result struct {
	Message_id  int
	Sender_chat Chat
	Chat        Chat
	Date        int
	Text        string
}
type Message struct {
	Ok     bool
	Result Result
}

func getTelegramInfo(tmurl string) string {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get(tmurl)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	return string(robots)
}

func SendTextMessage(text string) {
	//getTelegramInfo(".../getMe")
	jsonStream := getTelegramInfo("https://api.telegram.org/bot" + conf.TM.BotAPIKey + "/sendMessage?chat_id=" + conf.TM.ChannelID + "&text=" + text)

	// const jsonStream = `
	// 	{"ok":true,
	// 	"result":{
	// 		"message_id":6,
	// 		"sender_chat":{
	// 			"id":-100XXXXXXXXXX,
	// 			"title":"jogtrot",
	// 			"type":"channel"},
	// 		"chat":{
	// 			"id":-100XXXXXXXXXX,
	// 			"title":"jogtrot",
	// 			"type":"channel"},
	// 		"date":1607127591,
	// 		"text":"Golang Test Message"}
	// 	}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	var m Message
	err := dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("text: ", m.Result.Text)
	fmt.Println("ok: ", m.Ok)
}

func SendImageFromTelegram(file_id string) {
	// file, err := os.Open(file_id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// getTelegramInfo("..." + file)
}
