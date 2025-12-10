package main

import (
	"context"
	"log"
	"net/http"

	"github.com/NikDevRych/auth-go/internal/infrastructure/db"
	"github.com/NikDevRych/auth-go/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	urlExample := "postgres://postgres:postgres@localhost:5432/auth_go_db"
	dbpool, err := pgxpool.New(ctx, urlExample)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	if err = dbpool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	repo := db.NewUserRepository(dbpool)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	mux.HandleFunc("POST /signup", handler.SignUp)
	mux.HandleFunc("POST /signin", handler.SignIn)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
