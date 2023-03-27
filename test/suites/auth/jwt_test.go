package auth_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhuandantas/verifymy-test/internal/server/middlewares/auth"
	mock_config "github.com/rhuandantas/verifymy-test/test/mock/config"
)

var _ = Describe("Test auth jwt methods", func() {
	var (
		mockCtrl *gomock.Controller
		config   *mock_config.MockConfigProvider
		jwtToken auth.Token
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		config = mock_config.NewMockConfigProvider(mockCtrl)
		jwtToken = auth.NewJwtToken(config)
	})

	It("generate token successfully", func(ctx SpecContext) {
		config.EXPECT().GetEnv("AUTH_SECRET").Return("secret")
		token, err := jwtToken.GenerateToken("email")
		Expect(err).To(BeNil())
		Expect(token).ToNot(BeEmpty())
	})

})
