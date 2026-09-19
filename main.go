package main

import (
	"net/http"

	"football-teams-creator/api"
	"football-teams-creator/db"
)

func main() {
	var dbAdapter db.DbAdapter
	dbAdapter.InitDb("./data/data.db")

	router := api.Router(&dbAdapter)

	http.ListenAndServe(":8080", router)

	dbAdapter.CloseDb()
}
