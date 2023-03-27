package repo_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhuandantas/verifymy-test/internal/models"
	"github.com/rhuandantas/verifymy-test/internal/repo"
	mock_log "github.com/rhuandantas/verifymy-test/test/mock/log"
	mock_repo "github.com/rhuandantas/verifymy-test/test/mock/repo"
	"gorm.io/gorm"
)

var _ = Describe("Test all user repo methods", func() {
	var (
		mockCtrl *gomock.Controller
		log      *mock_log.MockSimpleLogger
		db       *mock_repo.MockDBConnection
		userRepo repo.UserRepo
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		log = mock_log.NewMockSimpleLogger(mockCtrl)
		db = mock_repo.NewMockDBConnection(mockCtrl)
		userRepo = repo.NewUserRepo(db, log)
	})

	Context("Create a user", func() {
		It("successfully", func(ctx SpecContext) {
			db.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			user, err := userRepo.Create(ctx, models.User{})
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
		})
		It("with fail", func(ctx SpecContext) {
			db.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("mock error")})
			_, err := userRepo.Create(ctx, models.User{})
			Expect(err).ToNot(BeNil())
		})
	})

	Context("Update a user", func() {
		It("successfully", func(ctx SpecContext) {
			db.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			user, err := userRepo.Update(ctx, 1, models.User{})
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
		})
		It("with get by id fail", func(ctx SpecContext) {
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("mock error")})
			_, err := userRepo.Update(ctx, 1, models.User{})
			Expect(err).ToNot(BeNil())
		})
		It("with fail", func(ctx SpecContext) {
			db.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("mock error")})
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			_, err := userRepo.Update(ctx, 1, models.User{})
			Expect(err).ToNot(BeNil())
		})
	})

	Context("Find a user", func() {
		It("successfully", func(ctx SpecContext) {
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			user, err := userRepo.GetByID(ctx, 1)
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
		})
		It("with fail", func(ctx SpecContext) {
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("mock error")})
			_, err := userRepo.GetByID(ctx, 1)
			Expect(err).ToNot(BeNil())
		})
		It("with record not found", func(ctx SpecContext) {
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("record not found")})
			_, err := userRepo.GetByID(ctx, 1)
			Expect(err.Error()).To(Equal("User not found with id 1"))
		})
	})

	Context("Delete a user", func() {
		It("successfully", func(ctx SpecContext) {
			db.EXPECT().Delete(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			user, err := userRepo.Delete(ctx, 1)
			Expect(err).To(BeNil())
			Expect(user).ToNot(BeNil())
		})
		It("with get by id fail", func(ctx SpecContext) {
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("mock error")})
			_, err := userRepo.Delete(ctx, 1)
			Expect(err).ToNot(BeNil())
		})
		It("with fail", func(ctx SpecContext) {
			db.EXPECT().Delete(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("mock error")})
			db.EXPECT().First(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			_, err := userRepo.Delete(ctx, 1)
			Expect(err).ToNot(BeNil())
		})
	})

	Context("Get all users", func() {
		It("successfully", func(ctx SpecContext) {
			db.EXPECT().FindAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: nil})
			user, err := userRepo.GetUsers(ctx, 1, 1)
			Expect(err).To(BeNil())
			Expect(user).To(BeEmpty())
		})
		It("with fail", func(ctx SpecContext) {
			db.EXPECT().FindAll(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&gorm.DB{Error: errors.New("mock error")})
			_, err := userRepo.GetUsers(ctx, 1, 1)
			Expect(err).ToNot(BeNil())
		})
	})
})
