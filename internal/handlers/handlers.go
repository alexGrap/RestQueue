package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	models "inter/internal"
	"inter/internal/middleware"
	"inter/internal/usecase"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Server  *mux.Router
	UseCase models.UseCase
}

func Fabric(countRoutine int, ctx context.Context, closeChan chan bool) models.Rest {
	server := Server{}
	server.Server = mux.NewRouter().StrictSlash(true)
	server.Server.HandleFunc("/create", server.Post).Methods("POST")
	server.Server.HandleFunc("/get", server.Get).Methods("GET")
	server.UseCase = usecase.InitUseCase(countRoutine, ctx, closeChan)
	return &server
}

func (server *Server) Hearing() error {
	log.Println("Server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", server.Server))
	return nil
}

func (server *Server) Post(w http.ResponseWriter, r *http.Request) {
	var body models.Task
	requestBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(requestBody, &body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(err)
		return
	}
	if !middleware.Validation(body) {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(err)
		return
	}
	server.UseCase.Create(body)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("Note is was created")
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Get(w http.ResponseWriter, r *http.Request) {
	result := server.UseCase.Get()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(err)
	}
}
