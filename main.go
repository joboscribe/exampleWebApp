package main

import (
 log "github.com/sirupsen/logrus"
 "github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

)

var db *sqlx.DB
var logger log.Logger

func main() {

	router := mux.NewRouter()
	router.HandleFunc()
	router.Path("/clients/{id}").HandlerFunc(sessionHandler)


}