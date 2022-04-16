package controllers

import (
	"log"
	"net/http"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	userId := request.URL.Query().Get("user_id")

	log.Print("About to process ", userId)
}
