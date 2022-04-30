package utils

import (
	"log"
)

type MutationReturn struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
