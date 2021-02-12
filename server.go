package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/gorm"

	"github.com/jrmbrgs/goodkarma-api/library/database"
	"github.com/jrmbrgs/goodkarma-api/library/graph/generated"
	"github.com/jrmbrgs/goodkarma-api/library/graph/resolvers"
	"github.com/jrmbrgs/goodkarma-api/repositories"
	"github.com/jrmbrgs/goodkarma-api/services"
)

const defaultPort = "5000"

var db *gorm.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := initDatabase()
	srv := initServer(db)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDatabase() *gorm.DB {

	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		panic(err)
	}
	database.RunMigrations(db)
	return db
}

func initServer(db *gorm.DB) *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolvers.Resolver{
					HumanService: &services.HumanService{
						Repo: &repositories.HumanRepository{
							Database: db,
						},
					},
				},
			},
		),
	)
}
