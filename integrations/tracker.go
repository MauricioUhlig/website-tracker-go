package integrations

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MauricioUhlig/go-ifes-notification-telegram/util"
)

func Changed(last_update string) (bool, error) {
	body, err := request()
	if body == "" {
		return false, nil
	}
	return !strings.Contains(body, last_update), err
}

func request() (string, error) {
	request, err := http.Get(util.GetUrl())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err.Error())
	}
	return string(body), nil
}
