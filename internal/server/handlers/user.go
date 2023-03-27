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
	validator util.Validator
	userRepo  repo.UserRepo
	token     auth.Token
	logger    log.SimpleLogger
}

func NewUserHandler(validator util.Validator, userRepo repo.UserRepo, jwt auth.Token, logger log.SimpleLogger) *UserHandler {
	return &UserHandler{
		validator: validator,
		userRepo:  userRepo,
		token:     jwt,
		logger:    logger,
	}
}

func (uh *UserHandler) RegisterRoutes(server *echo.Echo) {
	g := server.Group("/users", uh.token.VerifyToken)
	g.POST("", uh.Create)
	g.PUT("/:id", uh.Update)
	g.DELETE("/:id", uh.Delete)
	g.GET("/:id", uh.GetById)
	g.GET("", uh.GetUsers)
	server.POST("/login", uh.Login)
}

// Create godoc
// @Summary Create a new user.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "user struct"
// @Security JWT
// @Success 	 200  {object} models.User
// @Failure      400,401,404,500  {object}  error.ErrorResponse
// @Router /users [post]
func (uh *UserHandler) Create(ctx echo.Context) error {
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

// Update godoc
// @Summary Update a user.
// @Tags Users
// @Accept json
// @Produce json
// @Param        id   path      int  true  "user id"
// @Param user body models.User true "user struct"
// @Security JWT
// @Success 	 200  {object} models.User
// @Failure      400,401,404,500  {object}  error.ErrorResponse
// @Router /users/{id} [put]
func (uh *UserHandler) Update(ctx echo.Context) error {
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

// Delete godoc
// @Summary      Delete a user by id
// @Description  Delete user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "user id"
// @Security JWT
// @Success      200  {string}  "user deleted"
// @Failure      400,401,404,500  {object}  error.ErrorResponse
// @Router       /users/{id} [Delete]
func (uh *UserHandler) Delete(ctx echo.Context) error {
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

// GetById godoc
// @Summary      Retrieve a user by id
// @Description  get user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "user id"
// @Security JWT
// @Success      200  {object}  models.User
// @Failure      400,401,404,500  {object}  error.ErrorResponse
// @Router       /users/{id} [get]
func (uh *UserHandler) GetById(ctx echo.Context) error {
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

// GetUsers godoc
// @Summary      Retrieve all users
// @Description  get all users
// @Tags         Users
// @Produce      json
// @Param        page   query      int  true  "page number"
// @Param        size   query      int  true  "size number"
// @Security JWT
// @Success      200  {array}  models.User
// @Failure      400,401,404,500  {object}  error.ErrorResponse
// @Router       /users [get]
func (uh *UserHandler) GetUsers(ctx echo.Context) error {
	pagination, err := uh.GetPagination(ctx)
	if err != nil {
		return err
	}

	res, err := uh.userRepo.GetUsers(ctx.Request().Context(), pagination.Size, pagination.Page)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.InternalError.New(err.Error()))
	}

	return serverErr.ResponseJson(ctx, res)
}

// Login godoc
// @Summary      Login
// @Tags         Auth
// @Produce      json
// @Param user body models.User true "user struct"
// @Success      200  {string} token ""
// @Failure      400,401,404,500  {object}  error.ErrorResponse
// @Router       /login [post]
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

func (uh *UserHandler) GetPagination(ctx echo.Context) (*models.Pagination, error) {
	var (
		pagination models.Pagination
		err        error
	)

	if err = ctx.Bind(&pagination); err != nil {
		return nil, serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	if err = uh.validator.ValidateStruct(pagination); err != nil {
		return nil, serverErr.HandleError(ctx, errx.BadRequest.New(err.Error()))
	}

	return &pagination, nil
}
