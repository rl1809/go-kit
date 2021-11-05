package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	mysql "go-kit/database/mysql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// New configures application resources and routes.
func New() (*chi.Mux, error) {

	ctx := context.Background()

	db, err := mysql.DBConnnect()

	queries := mysql.New(db)
	authors, err := queries.ListAuthors(ctx)

	if err!=nil{
		log.Fatal(err)
	}



	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		
		jsonResponse, err := json.Marshal(authors)
		if err!=nil{
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		
		w.Write(jsonResponse)
	})

	return r, nil
}