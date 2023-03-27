package log_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	"github.com/rhuandantas/verifymy-test/internal/log"
	mock_config "github.com/rhuandantas/verifymy-test/test/mock/config"
)

var _ = Describe("Test Logging methods", func() {
	var (
		mockCtrl *gomock.Controller
		config   *mock_config.MockConfigProvider
		logger   log.SimpleLogger
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		config = mock_config.NewMockConfigProvider(mockCtrl)
		config.EXPECT().GetString(gomock.Any()).Return("info").AnyTimes()
		logger = log.NewLogger(config)
	})

	It("call info successfully", func(ctx SpecContext) {
		logger.Info("info")
	})

	It("call infoF successfully", func(ctx SpecContext) {
		logger.Infof("%s", "info")
	})

	It("call error successfully", func(ctx SpecContext) {
		logger.Error("error")
	})

	It("call errorF successfully", func(ctx SpecContext) {
		logger.Errorf("%s", "error")
	})

	It("call infoF successfully", func(ctx SpecContext) {
		logger.Warn("warn")
	})

	It("call infoF successfully", func(ctx SpecContext) {
		logger.Warnf("%s", "warn")
	})

	It("call infoF successfully", func(ctx SpecContext) {
		logger.Debug("%s", "debug")
	})

	It("call infoF successfully", func(ctx SpecContext) {
		logger.Warnf("%s", "warn")
	})
})
