package main

import (
	gcontext "github.com/reinhartf/AR-Backend/context"
	h "github.com/reinhartf/AR-Backend/handler"
	"github.com/reinhartf/AR-Backend/resolver"
	"github.com/reinhartf/AR-Backend/schema"
	"github.com/reinhartf/AR-Backend/service"
	"log"
	"net/http"

	"github.com/reinhartf/AR-Backend/loader"
	graphql "github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
)

func main() {
	config := gcontext.LoadConfig(".")

	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	ctx := context.Background()
	log := service.NewLogger(config)
	roleService := service.NewRoleService(db, log)
	userService := service.NewUserService(db, roleService, log)
	authService := service.NewAuthService(config, log)

	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "roleService", roleService)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "authService", authService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	http.Handle("/login", h.AddContext(ctx, h.Login()))

	loggerHandler := &h.LoggerHandler{config.DebugMode}
	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(h.Authenticate(&h.GraphQL{Schema: graphqlSchema, Loaders: loader.NewLoaderCollection()}))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
