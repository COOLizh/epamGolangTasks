package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ans := Answer{
		r.Host,
		r.UserAgent(),
		r.RequestURI,
		r.Header,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(ans)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}
