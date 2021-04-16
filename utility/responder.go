package utility

import (
	"encoding/json"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, jsonObj interface{}) {
	bs, err := json.Marshal(jsonObj)
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(bs)
	w.WriteHeader(status)
	return
}
