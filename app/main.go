package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/johncalvinroberts/furizu/app/users"
	"github.com/johncalvinroberts/furizu/app/utils"
	"github.com/johncalvinroberts/furizu/app/whoami"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/johncalvinroberts/furizu/app/graph"
	"github.com/johncalvinroberts/furizu/app/graph/generated"
)

const defaultPort = "4000"

//go:embed client/build/*
var embeddedFiles embed.FS

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
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", http.FileServer(getFileSystem()))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func getFileSystem() http.FileSystem {

	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "client/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
