package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rhuandantas/verifymy-test/internal/auth"
	"github.com/rhuandantas/verifymy-test/internal/models"
	"github.com/rhuandantas/verifymy-test/internal/repo"
	"github.com/rhuandantas/verifymy-test/internal/util"
	"net/http"
	"strconv"
)

type UserHandler struct {
	validator *util.CustomValidator
	userRepo  repo.UserRepo
	token     *auth.JwtToken
}

func NewUserHandler(validator *util.CustomValidator, userRepo repo.UserRepo, jwt *auth.JwtToken) *UserHandler {
	return &UserHandler{
		validator: validator,
		userRepo:  userRepo,
		token:     jwt,
	}
}

func (uh *UserHandler) RegisterRoutes(server *echo.Echo) {
	g := server.Group("/users", uh.token.VerifyToken)
	g.POST("", uh.create)
	g.PUT("/:id", uh.update)
	g.DELETE("/:id", uh.delete)
	g.GET("/:id", uh.getById)
	g.GET("", uh.getUsers)
	server.POST("/login", uh.Login)
}

func (uh *UserHandler) create(ctx echo.Context) error {
	var (
		user models.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err = uh.validator.ValidateStruct(user); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	userDTO := models.User{
		Name:     user.Name,
		Age:      user.Age,
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Password,
	}

	res, err := uh.userRepo.Create(ctx.Request().Context(), userDTO)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	if res != nil {
		res.Password = ""
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (uh *UserHandler) update(ctx echo.Context) error {
	var (
		user models.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err = uh.validator.ValidateStruct(user); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	res, err := uh.userRepo.Update(ctx.Request().Context(), id, user)
	if err != nil {
		return ctx.String(500, err.Error())
	}

	res.Password = ""
	return ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	res, err := uh.userRepo.Delete(ctx.Request().Context(), id)
	if err != nil {
		return ctx.String(500, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) getById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	res, err := uh.userRepo.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.String(500, err.Error())
	}
	res.Password = ""

	return ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) getUsers(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	offset, err := strconv.Atoi(ctx.QueryParam("size"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	res, err := uh.userRepo.GetUsers(ctx.Request().Context(), offset, page)
	if err != nil {
		return ctx.String(500, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (uh *UserHandler) Login(ctx echo.Context) error {
	var (
		user models.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if user.Email == "" || user.Password == "" {
		return ctx.JSON(http.StatusUnauthorized, errors.New("email or password invalid"))
	}

	res, err := uh.userRepo.GetByEmail(ctx.Request().Context(), user.Email)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, errors.New("email or password invalid"))
	}

	if err = user.VerifyPassword(res.Password); err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	token, err := uh.token.GenerateToken(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.Response().Header().Set("token", token)

	return ctx.JSON(http.StatusOK, "authorized")
}
