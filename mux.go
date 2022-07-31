package main

import (
	"context"
	"net/http"
	"todo/clock"
	"todo/config"
	"todo/handler"
	"todo/service"
	"todo/store"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset-utf-8")

		_, _ = w.Write([]byte(`{"status": "OK"}`))
	})

	v := validator.New()

	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	r := store.Repository{
		Clocker: clock.RealClocker{},
	}

	at := &handler.AddTask{
		Service: &service.AddTask{
			DB:   db,
			Repo: &r,
		},
		DB:        db,
		Repo:      &r,
		Validator: v,
	}

	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{
		Service: &service.ListTasks{
			DB:   db,
			Repo: &r,
		},
		DB:   db,
		Repo: r,
	}

	mux.Get("/tasks", lt.ServeHTTP)

	ru := &handler.RegisterUser{
		Service: &service.RegisterUser{
			DB:   db,
			Repo: &r,
		},
		Validator: v,
	}

	mux.Post("/register", ru.ServeHTTP)

	return mux, cleanup, nil
}
