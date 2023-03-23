package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rhuandantas/verifymy-test/internal/repo"
)

type UserHandler struct {
	userRepo repo.UserRepo
}

func NewUserHandler(userRepo repo.UserRepo) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (uh *UserHandler) RegisterRoutes(server *echo.Echo) {
	server.POST("/users", uh.create)
	server.PUT("/users/:id", uh.update)
	server.DELETE("/users/:id", uh.delete)
	server.GET("/users", uh.delete)
}

func (uh *UserHandler) create(ctx echo.Context) error {
	return ctx.String(201, "Created successfully")
}

func (uh *UserHandler) update(ctx echo.Context) error {
	return ctx.String(200, "Updated successfully")
}

func (uh *UserHandler) delete(ctx echo.Context) error {
	return nil
}

func (uh *UserHandler) getById(ctx echo.Context) error {
	return nil
}

func (uh *UserHandler) getByEmail(ctx echo.Context) error {
	return nil
}
