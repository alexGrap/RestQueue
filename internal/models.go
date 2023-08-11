package models

import (
	"net/http"
)

type UseCase interface {
	Create(task Task)
	Get() []Task
}

type Rest interface {
	Hearing() error
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type Task struct {
	Place        int
	Status       string
	ElementCount int     `json:"n"`
	Delta        float64 `json:"d"`
	StartElement float64 `json:"n1"`
	L            float64 `json:"l"`
	TTL          float64 `json:"TTL"`
	Iteration    int
	CreationTime string
	StartTime    string
	EndTime      string
}
