package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/johncalvinroberts/furizu/app/users"
	"github.com/johncalvinroberts/furizu/app/utils"
	"github.com/johncalvinroberts/furizu/app/whoami"

	"github.com/johncalvinroberts/furizu/app/graph"
	"github.com/johncalvinroberts/furizu/app/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const defaultPort = "4000"

//go:embed client/build/*
var embeddedFiles embed.FS

type embedFileSystem struct {
	http.FileSystem
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
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := gin.Default()
	// static server
	router.Use(static.Serve("/", embedFolder(embeddedFiles, "client/build")))
	router.Use(utils.CORSMiddleware())
	router.Use(ginContextToContextMiddleware())
	router.GET("/playground", playgroundHandler())
	router.POST("/query", graphQlHandler())

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(router.Run("localhost:" + port))
}

// needed to fulfill the interface of gin-contrib/static
func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

// embed folder to FS
func embedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphQlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
