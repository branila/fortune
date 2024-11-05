package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func Master(w http.ResponseWriter, r *http.Request) {
	update, err := parseRequest(r)
	if err != nil {
		log.Printf("Error parsing request: %e", err)
	}

	if !strings.HasPrefix(update.Message.Text, "/fortune") {
		return
	}

	response, err := sendFortune(update)
	if err != nil {
		log.Printf("Got error %e, response is %s", err, response)
		return
	}

	fmt.Printf("Sent fortune to %d", update.Message.Chat.Id)
}

func getFortune() (string, error) {
	fortune, err := exec.Command("fortune").Output()
	if err != nil {
		return "", err
	}

	fortune = []byte(normalizeText(string(fortune)))

	return string(fortune), nil
}

func normalizeText(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")

	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	text = strings.TrimSpace(text)

	return text
}
