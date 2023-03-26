package handlers

import (
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	errx "github.com/rhuandantas/verifymy-test/internal/errors"
	"github.com/rhuandantas/verifymy-test/internal/log"
	"github.com/rhuandantas/verifymy-test/internal/models"
	"github.com/rhuandantas/verifymy-test/internal/repo"
	serverErr "github.com/rhuandantas/verifymy-test/internal/server/error"
	"github.com/rhuandantas/verifymy-test/internal/server/middlewares/auth"
	"github.com/rhuandantas/verifymy-test/internal/util"
	"strconv"
)

type UserHandler struct {
	validator *util.CustomValidator
	userRepo  repo.UserRepo
	token     *auth.JwtToken
	logger    log.SimpleLogger
}

func NewUserHandler(validator *util.CustomValidator, userRepo repo.UserRepo, jwt *auth.JwtToken, logger log.SimpleLogger) *UserHandler {
	return &UserHandler{
		validator: validator,
		userRepo:  userRepo,
		token:     jwt,
		logger:    logger,
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
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	if err = uh.validator.ValidateStruct(user); err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	res, err := uh.userRepo.Create(ctx.Request().Context(), user)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.InternalError.New(err.Error()))
	}

	if res != nil {
		res.Password = ""
	}

	return serverErr.ResponseJson(ctx, res)
}

func (uh *UserHandler) update(ctx echo.Context) error {
	var (
		user models.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	if err = uh.validator.ValidateStruct(user); err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	res, err := uh.userRepo.Update(ctx.Request().Context(), id, user)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.InternalError.New(err.Error()))
	}

	res.Password = ""
	return serverErr.ResponseJson(ctx, res)
}

func (uh *UserHandler) delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	res, err := uh.userRepo.Delete(ctx.Request().Context(), id)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.InternalError.New(err.Error()))
	}

	return serverErr.ResponseJson(ctx, res)
}

func (uh *UserHandler) getById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	res, err := uh.userRepo.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.InternalError.New(err.Error()))
	}
	res.Password = ""

	return serverErr.ResponseJson(ctx, res)
}

func (uh *UserHandler) getUsers(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	offset, err := strconv.Atoi(ctx.QueryParam("size"))
	if err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	res, err := uh.userRepo.GetUsers(ctx.Request().Context(), offset, page)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.InternalError.New(err.Error()))
	}

	return serverErr.ResponseJson(ctx, res)
}

func (uh *UserHandler) Login(ctx echo.Context) error {
	var (
		user models.User
		err  error
	)

	if err = ctx.Bind(&user); err != nil {
		return serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	if user.Email == "" || user.Password == "" {
		return serverErr.HandleError(ctx, errx.Unauthorized.New("email or password invalid"))
	}

	res, err := uh.userRepo.GetByEmail(ctx.Request().Context(), user.Email)
	if err != nil {
		return serverErr.HandleError(ctx, errx.Unauthorized.New("email or password invalid"))
	}

	if err = user.VerifyPassword(res.Password); err != nil {
		return serverErr.HandleError(ctx, errx.Unauthorized.New(err.Error()))
	}

	token, err := uh.token.GenerateToken(user.Email)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.InternalError.New(err.Error()))
	}

	return serverErr.ResponseJson(ctx, echo.Map{"token": token})
}
