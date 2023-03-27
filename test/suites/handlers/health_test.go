package handlers_test

import (
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhuandantas/verifymy-test/internal/server/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Test all handlers methods", func() {
	var (
		e           *echo.Echo
		healthCheck *handlers.HealthCheck
	)

	BeforeEach(func() {
		e = echo.New()
		healthCheck = handlers.NewHealthCheck()
	})

	AfterEach(func() {
		e.Close()
	})

	Context("Call health check", func() {
		It("readiness", func(ctx SpecContext) {
			req := httptest.NewRequest(http.MethodGet, "/readiness", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := healthCheck.Readiness(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
		})
		It("liveness", func(ctx SpecContext) {
			req := httptest.NewRequest(http.MethodGet, "/liveness", strings.NewReader(""))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := healthCheck.Liveness(c)
			Expect(err).To(BeNil())
			Expect(c.Response()).ToNot(BeNil())
		})
		It("liveness", func(ctx SpecContext) {
			healthCheck.RegisterHealth(e)
		})
	})
})
