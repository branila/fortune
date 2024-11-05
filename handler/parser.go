package handler

import (
	"encoding/json"
	"net/http"

	"github.com/branila/fortune/types"
)

func parseRequest(r *http.Request) (types.Update, error) {
	var update types.Update

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return types.Update{}, err
	}

	return update, nil
}
