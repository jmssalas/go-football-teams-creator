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
	teamsFilepath := "./data/teams.json"

	r := api.Router(&dbAdapter, teamsFilepath)

	r.StaticFS("/css", http.Dir("./web/css")).Use(noCacheMiddleware())
	r.StaticFS("/js", http.Dir("./web/js")).Use(noCacheMiddleware())
	r.StaticFS("/assets", http.Dir("./web/assets")).Use(noCacheMiddleware())
	r.StaticFile("/", "./web/index.html").Use(noCacheMiddleware())
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

func noCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Next()
	}
}
