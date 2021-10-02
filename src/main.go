package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/johncalvinroberts/furizu/src/files"
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
	router := gin.Default()
	// static server
	router.Use(static.Serve("/", EmbedFolder(embeddedFiles, "client/build")))
	// api
	api := router.Group("/api")
	filesApi := api.Group("/fs")
	filesApi.GET("/", files.Cmd)
	router.Run("localhost:3000")
}
