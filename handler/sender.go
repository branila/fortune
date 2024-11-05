package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/branila/fortune/types"
)

func sendFortune(update types.Update) (string, error) {
	token := os.Getenv("FORTUNE_TOKEN")
	if token == "" {
		return "", fmt.Errorf("FORTUNE_TOKEN is not set")
	}

	api := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	fortune, err := getFortune()
	if err != nil {
		fmt.Printf("Error getting fortune: %e", err)
		return "", err
	}

	response, err := http.PostForm(
		api,
		url.Values{
			"chat_id": {strconv.Itoa(update.Message.Chat.Id)},
			"text":    {fortune},
		},
	)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	body := string(bodyBytes)

	return body, nil
}
