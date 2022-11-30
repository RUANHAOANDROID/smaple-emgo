package main

import (
	"emcs-relay-go/logger"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

func LoadStatic() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", http.FileServer(statikFS))
	logger.Log.Info(statikFS)
}
