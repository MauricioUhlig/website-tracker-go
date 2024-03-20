package main

import (
	"fmt"
	"os"
	"time"

	"github.com/MauricioUhlig/go-ifes-notification-telegram/integrations"
	"github.com/MauricioUhlig/go-ifes-notification-telegram/util"
)

const DATETIME_FORMAT = "2006-01-02T15:04:05"

func main() {
	util.FetchYAML()
	loop()
}

func loop() {
	for {
		changed, err := integrations.Changed(util.GetExpectedString())
		if changed {
			integrations.SendNotification("Atualizou o resultado")
			integrations.SendNotification("Atualizou o resultado")
			integrations.SendNotification("Atualizou o resultado")
			os.Exit(0)
			break
		}
		fmt.Println(time.Now().Format(DATETIME_FORMAT), "- Ainda não", "| Tem Erro:", err != nil)
		time.Sleep(1 * time.Minute)
	}
}
