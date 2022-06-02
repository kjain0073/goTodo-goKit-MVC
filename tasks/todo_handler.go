package tasks

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/kjain0073/go-Todo/models"
)

type (
	CreateTodoRequest struct {
		Title string `json:"title"`
	}
	CreateTodoResponse struct {
		Ok string `json:"ok"`
	}

	GetTodosRequest struct {
	}

	GetTodosResponse struct {
		TodoList []models.TodoDto
	}

	DeleteTodoRequest struct {
		Id string `json:"id"`
	}
	DeleteTodoResponse struct {
		Ok string `json:"ok"`
	}

	UpdateTodoRequest struct {
		Id        string `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	UpdateTodoResponse struct {
		Ok string `json:"ok"`
	}
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		endpoints.CreateTodo,
		decodeCreateTodoReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		endpoints.GetTodos,
		decodeGetTodosReq,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteTodo,
		decodeDeleteTodoReq,
		encodeResponse,
	))

	r.Methods("PUT").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateTodo,
		decodeUpdateTodoReq,
		encodeResponse,
	))
	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetTodosReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetTodosRequest
	return req, nil

}

func decodeUpdateTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UpdateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteTodoReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req DeleteTodoRequest
	vars := mux.Vars(r)

	req = DeleteTodoRequest{
		Id: vars["id"],
	}
	return req, nil
}
