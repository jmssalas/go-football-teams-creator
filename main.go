package main

import (
	"net/http"

	"football-teams-creator/api"
	"football-teams-creator/db"

	"github.com/gin-gonic/gin"
)

func main() {
	var dbAdapter db.DbAdapter
	dbAdapter.InitDb("./data/data.db")

	r := api.Router(&dbAdapter)

	r.StaticFS("/css", http.Dir("./web/css"))
	r.StaticFS("/js", http.Dir("./web/js"))
	r.StaticFS("/assets", http.Dir("./web/assets"))
	r.StaticFile("/", "./web/index.html")
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			c.File("./web/index.html")
			return
		}

		c.Status(http.StatusNotFound)
	})

	http.ListenAndServe(":8080", r)

	dbAdapter.CloseDb()
}
