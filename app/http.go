package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *App) ServeHTTP() error {
	if os.Getenv("DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	app.http = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	app.http.Use(cors.New(corsConfig), auth.Middleware())

	app.http.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Nimbus - extensible storage system")
	})

	// GraphQL
	app.http.POST("/graphql", handlers.GQLServerHandler(app.store))
	app.http.GET("/playground", handlers.GQLPlaygroundHandler())

	// Files
	app.http.GET("/files/:id", handlers.FileHandler())
	app.http.GET("/files/download/:id/:name", handlers.FileDownloadHandler())

	log.Println("Nimbus server listening on " + os.Getenv("HOST"))

	srv := &http.Server{
		Addr:           os.Getenv("HOST"),
		Handler:        app.http,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv.ListenAndServe()
}
