package handlers_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhuandantas/verifymy-test/internal/models"
	"github.com/rhuandantas/verifymy-test/internal/server/handlers"
	mock_auth "github.com/rhuandantas/verifymy-test/test/mock/auth"
	mock_log "github.com/rhuandantas/verifymy-test/test/mock/log"
	mock_repo "github.com/rhuandantas/verifymy-test/test/mock/repo"
	mock_util "github.com/rhuandantas/verifymy-test/test/mock/util"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

var _ = Describe("Test all handlers methods", func() {
	var (
		mockCtrl    *gomock.Controller
		e           *echo.Echo
		validator   *mock_util.MockValidator
		userRepo    *mock_repo.MockUserRepo
		tokenJwt    *mock_auth.MockToken
		logger      *mock_log.MockSimpleLogger
		userHandler *handlers.UserHandler
		mockUser    models.User
	)

	BeforeEach(func() {
		e = echo.New()
		mockCtrl = gomock.NewController(GinkgoT())
		validator = mock_util.NewMockValidator(mockCtrl)
		userRepo = mock_repo.NewMockUserRepo(mockCtrl)
		tokenJwt = mock_auth.NewMockToken(mockCtrl)
		logger = mock_log.NewMockSimpleLogger(mockCtrl)
		userHandler = handlers.NewUserHandler(validator, userRepo, tokenJwt, logger)
		mockUser = models.User{
			UserId:   1,
			Name:     "Jon Snow",
			Age:      30,
			Email:    "jon@email.com",
			Password: "123456",
			Address:  "rua network",
		}
	})

	AfterEach(func() {
		e.Close()
	})

	Context("Call user create handler", func() {
		It("successfully", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"12345"}`
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
			userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Create(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(200))
		})

		It("json body invalid", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":12345}`
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Create(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("email field required", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","password":"12345"}`
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(errors.New("mock error"))
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Create(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("create user repo fails", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"12345"}`
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
			userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error"))
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Create(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(500))
		})
	})

	Context("Call user update handler", func() {
		It("successfully", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"12345","address":"teste"}`
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
			userRepo.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.Update(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(200))
		})

		It("json body invalid", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":12345,"address":"teste"}`
			req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.Update(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("path param id is missing", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"12345","address":"teste"}`
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
			req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Update(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("email field required", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","password":"12345","address":"teste"}`
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(errors.New("mock error"))
			req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.Update(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("update repo fails", func(ctx SpecContext) {
			userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"12345","address":"teste"}`
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
			userRepo.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error"))
			req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.Update(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(500))
		})
	})

	Context("Call user delete handler", func() {
		It("successfully", func(ctx SpecContext) {
			userRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(true, nil)
			req := httptest.NewRequest(http.MethodDelete, "/users", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.Delete(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(200))
		})

		It("path param id is missing", func(ctx SpecContext) {
			req := httptest.NewRequest(http.MethodDelete, "/users", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Delete(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("delete user repo fails", func(ctx SpecContext) {
			userRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(false, errors.New("mock error"))
			req := httptest.NewRequest(http.MethodDelete, "/users", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.Delete(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(500))
		})
	})

	Context("Call get user by id handler", func() {
		It("successfully", func(ctx SpecContext) {
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&models.User{}, nil)
			req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.GetById(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(200))
		})

		It("path param id is missing", func(ctx SpecContext) {
			req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.GetById(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("get by id repo fails", func(ctx SpecContext) {
			userRepo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error"))
			req := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")
			err := userHandler.GetById(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(500))
		})
	})

	Context("Call get all users handler", func() {
		It("successfully", func(ctx SpecContext) {
			userRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.User{}, nil)
			validator.EXPECT().ValidateStruct(gomock.Any()).Return(nil)
			q := make(url.Values)
			q.Set("page", "0")
			q.Set("size", "10")
			req := httptest.NewRequest(http.MethodGet, "/users?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.GetUsers(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(200))
		})
	})

	Context("Call login handler", func() {
		It("successfully", func(ctx SpecContext) {
			userJSON := `{"email":"jon@labstack.com","password":"12345"}`
			tokenJwt.EXPECT().GenerateToken(gomock.Any()).Return("", nil)
			mockUser.Password = "$2a$10$g1BebFwhxXMbsAn4G7rj4..u6pkVogpGlE8clTY3PXartfNsIDNmG"
			userRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Login(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(200))
		})

		It("email empty", func(ctx SpecContext) {
			userJSON := `{"email":"","password":"12345"}`
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Login(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(401))
		})

		It("password empty", func(ctx SpecContext) {
			userJSON := `{"email":"teste","password":""}`
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Login(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(401))
		})

		It("successfully", func(ctx SpecContext) {
			userJSON := `{"email":"jon@labstack.com","password":12345}`
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Login(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(400))
		})

		It("successfully", func(ctx SpecContext) {
			userJSON := `{"email":"jon@labstack.com","password":"12345"}`
			userRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(nil, errors.New("mock error"))
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Login(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(401))
		})

		It("wrong password", func(ctx SpecContext) {
			userJSON := `{"email":"jon@labstack.com","password":"123453"}`
			mockUser.Password = "$2a$10$g1BebFwhxXMbsAn4G7rj4..u6pkVogpGlE8clTY3PXartfNsIDNmG"
			userRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Login(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(401))
		})

		It("generate token fails", func(ctx SpecContext) {
			userJSON := `{"email":"jon@labstack.com","password":"12345"}`
			tokenJwt.EXPECT().GenerateToken(gomock.Any()).Return("", errors.New("mock error"))
			mockUser.Password = "$2a$10$g1BebFwhxXMbsAn4G7rj4..u6pkVogpGlE8clTY3PXartfNsIDNmG"
			userRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := userHandler.Login(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
			Expect(c.Response().Status).To(Equal(500))
		})
	})

	It("call register handlers", func(ctx SpecContext) {
		userHandler.RegisterRoutes(e)
	})
})
