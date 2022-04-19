package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type URLAddress string

func (u URLAddress) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", "3000", u)
	return []byte(url), nil
}

type URL struct {
	URL         URLAddress `json:"url"`
	Method      string     `json:"method"`
	Description string     `json:"description"`
}

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	data := []URL{
		{
			URL:         URLAddress("/"),
			Method:      "GET",
			Description: "See USERS",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func Start() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":3000", mux)
}
