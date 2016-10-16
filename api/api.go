package api

import (
	"log"
	"net/http"

	"github.com/maliceio/malice/config"
)

// StartAPI starts the malice API
func StartAPI() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(config.Conf.Web.API, router))
}
