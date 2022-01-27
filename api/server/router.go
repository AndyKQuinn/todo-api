package server

import (
	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
)

// Router is our main router
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/graphql", ExecuteQueryGraphql).Methods("POST", "OPTIONS")

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/api/graphql")
	if err != nil {
		panic(err)
	}
	router.Handle("/api/graphiql", graphiqlHandler)
	return router
}
