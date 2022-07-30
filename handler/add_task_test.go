package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/entity"
	"todo/store"
	"todo/testutil"

	"github.com/go-playground/validator"
)

func TestAddTask(t *testing.T) {
	t.Parallel()

	type want struct {
		status  int
		rspFile string
	}

	tests := map[string]struct {
		reqFile string
		want    want
	}{
		"ok": {
			reqFile: "testdata/add_task/ok_req.json.golden",
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/add_task/ok_rsp.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_req.json.golden",
			want: want{
				status:  http.StatusInternalServerError,
				rspFile: "testdata/add_task/bad_rsp.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()

			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader(testutil.LoadFile(t, tt.reqFile)),
			)

			moq := &AddTaskServiceMock{}
			moq.AddTaskFunc = func(
				ctx context.Context,
				title string,
			) (*entity.Task, error) {
				if tt.want.status == http.StatusOK {
					return &entity.Task{
						ID: 1,
					}, nil
				}
				return nil, errors.New("error from mock")
			}

			sut := AddTask{
				Store: &store.TaskStore{
					Tasks: map[entity.TaskID]*entity.Task{},
				},
				Validator: validator.New(),
			}

			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(
				t,
				resp,
				tt.want.status,
				testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
