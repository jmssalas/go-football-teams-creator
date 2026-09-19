package main

import (
	"net/http"

	"football-teams-creator/internals/api"
)

func main() {
	router := api.Router()

	http.ListenAndServe(":8080", router)
}
