package main

import (
	"context"
	"log"
	"net/http"
	"database/sql"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"

	dbfront "github.com/anz-bank/test-rig-example/gen/dbfront"
	dbfront_impl "github.com/anz-bank/test-rig-example/pkg/dbfront_impl"
)

func initDb() (*sql.DB, error) {
	psqlInfo := "host=dbfront_db port=5432 user=someuser password=somepassword dbname=somedb sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

func loadServices(ctx context.Context) error {
	router := chi.NewRouter()

	db, err := initDb()
	if err != nil {
		return err
	}
	defer db.Close()

	serviceInterface := dbfront_impl.GetServiceInterfaceImplementation()

	genCallbacks := dbfront_impl.GetCallback(db)

	serviceHandler := dbfront.NewServiceHandler(genCallbacks, &serviceInterface)

	serviceRouter := dbfront.NewServiceRouter(genCallbacks, serviceHandler)
	serviceRouter.WireRoutes(ctx, router)

	log.Println("starting dbfront on :8080")
	return http.ListenAndServe(":8080", router)
}

func main() {
	log.Fatal(loadServices(context.Background()))
}
