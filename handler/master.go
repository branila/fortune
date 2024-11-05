package handler

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func Master(w http.ResponseWriter, r *http.Request) {
	update, err := parseRequest(r)
	if err != nil {
		log.Printf("Error parsing request: %e", err)
	}

	response, err := sendFortune(update)
	if err != nil {
		log.Printf("Got error %e, response is %s", err, response)
		return
	}

	fmt.Println("Sent fortune to %d", update.Message.Chat.Id)
}

func getFortune() (string, error) {
	fortune, err := exec.Command("fortune").Output()
	if err != nil {
		return "", err
	}

	return string(fortune), nil
}
