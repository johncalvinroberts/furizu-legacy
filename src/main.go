package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/johncalvinroberts/furizu/src/archives"
	"github.com/johncalvinroberts/furizu/src/users"
	"github.com/johncalvinroberts/furizu/src/utils"
	"github.com/johncalvinroberts/furizu/src/whoami"
)

//go:embed client/build/*
var embeddedFiles embed.FS

type embedFileSystem struct {
	http.FileSystem
}

// needed to fulfill the interface of gin-contrib/static
func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

// embed folder to FS
func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

// run the server
func main() {
	fmt.Println("Starting Server")
	utils.InitEnv()
	utils.InitJWT()
	utils.InitAWS()
	// get table names
	users.InitRepository(utils.FurizuDB, os.Getenv("USERS_TABLE"))
	whoami.InitRepository(utils.FurizuDB, os.Getenv("WHOAMI_CHALLENGES_TABLE"))

	router := gin.Default()
	// static server
	router.Use(static.Serve("/", EmbedFolder(embeddedFiles, "client/build")))
	// api
	api := router.Group("/api")
	// archives
	archivesApi := api.Group("/archives")
	archivesApi.GET("/", archives.FindMany)
	archivesApi.POST("/", archives.Create)
	archivesApi.GET("/:id", archives.FindOne)
	archivesApi.DELETE("/:id", archives.DestroyOne)
	// whoami
	whoamiApi := api.Group("/whoami")
	whoamiApi.GET("/", whoami.Me)
	whoamiApi.POST("/", whoami.Start)
	whoamiApi.PATCH("/redeem", whoami.Redeem)
	whoamiApi.PATCH("/refresh", whoami.Refresh)
	whoamiApi.DELETE("/revoke", whoami.Revoke)

	router.Run("localhost:4000")
}
