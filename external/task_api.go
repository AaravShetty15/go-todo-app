package external

import (
	"encoding/json"
	"net/http"
)

type TaskResponse struct {
	Activity string `json:"activity"`
	Type     string `json:"type"`
}

func GetSuggestedTask() (TaskResponse, error) {

	resp, err := http.Get("https://bored-api.appbrewery.com/random")
	if err != nil {
		return TaskResponse{}, err
	}
	defer resp.Body.Close()

	var task TaskResponse
	err = json.NewDecoder(resp.Body).Decode(&task)

	return task, err
}