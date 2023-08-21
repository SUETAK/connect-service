package main

import (
	"connectrpc.com/connect"
	"database/sql"
	"example/interfaces/proto/eliza/v1/elizav1connect"

	"example/interfaces/interceptor"
	"example/usecase"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	address := "0.0.0.0:8000"
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	openaiKey := os.Getenv("OPENAI_KEY")
	if openaiKey == "" {
		log.Fatal("OPENAI_KEY environment variable must be set")
	}

	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	var v string
	if err := db.NewSelect().ColumnExpr("version()").Scan(context.Background(), &v); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	var issuer, keyPath string
	issuer = "eliza"
	keyPath = "./private.pem"
	authInterceptor := connect.WithInterceptors(interceptor.NewAuthInterceptor(issuer, keyPath))

	eliza := usecase.NewElizaServer(openaiKey)

	mux := http.NewServeMux()
	path, handler := elizav1connect.NewElizaServiceHandler(eliza, authInterceptor)
	mux.Handle(path, handler)

	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
	err = http.ListenAndServe(
		address,
		corsHandler,
	)
	if err != nil {
		panic(err)
	}
}
