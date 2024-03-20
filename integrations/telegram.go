package integrations

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/MauricioUhlig/go-ifes-notification-telegram/util"
)

type telegram struct {
	Chat    int    `json:"chat_id"`
	Message string `json:"text"`
}

func SendNotification(message string) {
	var msg telegram = telegram{util.GetChatId(), message}
	send(msg)
}
func send(msg telegram) bool {
	var posturl = "https://api.telegram.org/bot" + util.GetTelegramKey() + "/sendMessage"
	body, _ := json.Marshal(msg)

	r, err := http.Post(posturl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	if r.StatusCode == 200 {
		return true
	}
	return false
}
