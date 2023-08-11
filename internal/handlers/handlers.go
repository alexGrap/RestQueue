package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	models "inter/internal"
	"inter/internal/usecase"
	"log"
)

type Server struct {
	Server  *fiber.App
	UseCase models.UseCase
}

func Fabric(countRoutine int, ctx context.Context, closeChan chan bool) models.Rest {
	server := Server{}
	server.Server = fiber.New()
	server.Server.Post("create", server.Post)
	server.Server.Get("get", server.Get)
	server.UseCase = usecase.InitUseCase(countRoutine, ctx, closeChan)
	return &server
}

func (server *Server) Hearing() error {
	err := server.Server.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (server *Server) Post(ctx *fiber.Ctx) error {
	var body models.Task
	if err := ctx.BodyParser(&body); err != nil {
		ctx.Status(400)
		return ctx.JSON(err.Error())
	}
	server.UseCase.Create(body)
	return ctx.SendStatus(200)
}

func (server *Server) Get(ctx *fiber.Ctx) error {
	result := server.UseCase.Get()
	return ctx.JSON(result)
}
