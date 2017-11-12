package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SpencerSharkey/gomc/query"
	"github.com/julienschmidt/httprouter"
)

// SimpleQuery - GET /query/{addr}
func SimpleQuery(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	req := query.NewRequest()
	req.Connect(params[0].Value)
	res, _ := req.Simple()

	json.NewEncoder(w).Encode(res)
}

// FullQuery - GET /query/{addr}/full
func FullQuery(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	req := query.NewRequest()
	req.Connect(params[0].Value)
	res, _ := req.Full()

	json.NewEncoder(w).Encode(res)
}

func main() {
	router := httprouter.New()
	router.GET("/query/:addr", SimpleQuery)
	router.GET("/query/:addr/full", FullQuery)
	log.Fatal(http.ListenAndServe(":8080", router))
}
