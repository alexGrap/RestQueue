package models

import (
	"github.com/gofiber/fiber/v2"
)

type UseCase interface {
	Create(task Task)
	Get() []Task
	Handle()
}

type Rest interface {
	Hearing() error
	Post(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
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
