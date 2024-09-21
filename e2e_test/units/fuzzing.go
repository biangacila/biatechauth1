package units

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ValuesRequest struct {
	Values []int `json:"values"`
}

func CalculateHighestHandler(w http.ResponseWriter, r *http.Request) {
	var valuesRequest ValuesRequest
	err := json.NewDecoder(r.Body).Decode(&valuesRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var highest int
	for _, value := range valuesRequest.Values {
		if value > highest {
			highest = value
		}
	}

	if highest == 50 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "%d", highest)
	if err != nil {
		return
	}
}
