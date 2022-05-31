package main

import (
	"fmt"
	"golang/database"
	"golang/graph"
	"golang/graph/generated"
	"golang/repository"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	config := &database.Config{
		Host:     "db",
		Port:     "5432",
		Password: "postgresql",
		User:     "postgresql",
		SSLMode:  "disable",
		DBName:   "api",
	}

	db, err := database.NewConnection(config)
	fmt.Println("sentinel6")
	if err != nil {
		fmt.Println("sentinel7")

		panic(err)
	}
	fmt.Println("sentinel8")

	database.Migrate(db)
	fmt.Println("sentinel9")
	repo := repository.NewBookService(db)
	fmt.Println("sentinel10")
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		BookRepository: repo,
	}}))

	fmt.Println("sentinel11")
	http.HandleFunc("/hoge", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":8080", nil)
	fmt.Println("sentinel11")
	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", 5432)

	// log.Fatal(http.ListenAndServe(":5432", nil))
}
