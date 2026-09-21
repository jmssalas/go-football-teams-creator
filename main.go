package main

import (
	"net/http"
	"os"

	"football-teams-creator/api"
	"football-teams-creator/db"

	"github.com/gin-gonic/gin"
)

func main() {
	dbFilepath := os.Getenv("SQLITE_DATABASE_PATH")
	teamsFilepath := os.Getenv("TEAMS_DATA_PATH")
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is not set")
	}
	if dbFilepath == "" {
		panic("SQLITE_DATABASE_PATH is not set")
	}
	if teamsFilepath == "" {
		panic("TEAMS_DATA_PATH is not set")
	}

	var dbAdapter db.DbAdapter
	dbAdapter.InitDb(dbFilepath)

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

	http.ListenAndServe(":"+port, r)

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
