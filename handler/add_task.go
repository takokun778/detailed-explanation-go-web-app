package handler

import (
	"encoding/json"
	"net/http"
	"todo/store"

	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
)

type AddTask struct {
	Store     *store.TaskStore
	DB        *sqlx.DB
	Repo      *store.Repository
	Service   AddTaskService
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	err := at.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	t, err := at.Service.AddTask(ctx, b.Title)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	rsp := struct {
		ID int `json:"id"`
	}{
		ID: int(t.ID),
	}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
