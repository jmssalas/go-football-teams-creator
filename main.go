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

	r.StaticFS("/css", http.Dir("./internal/web/css"))
	r.StaticFS("/js", http.Dir("./internal/web/js"))
	r.StaticFS("/assets", http.Dir("./internal/web/assets"))
	r.StaticFile("/", "./internal/web/index.html")
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			c.File("./internal/web/index.html")
			return
		}

		c.Status(http.StatusNotFound)
	})

	http.ListenAndServe(":8080", r)

	dbAdapter.CloseDb()
}
