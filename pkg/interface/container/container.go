package container

import (
	"github.com/go-playground/validator"
	"goarch/pkg/infrastructure/goarch"
	"goarch/pkg/infrastructure/gorm"
	"goarch/pkg/shared/config"
	Database "goarch/pkg/shared/database"
	authSvc "goarch/pkg/usecase/auth"
	userSvc "goarch/pkg/usecase/user"
)

type Container struct {
	Config  *config.Config
	UserSvc userSvc.Service
	AuthSvc authSvc.Service

	Validate *validator.Validate
}

func Setup() *Container {
	// ====== Construct Config
	cfg := config.NewConfig("./resources/config.json")

	// ====== Construct Database
	db := Database.New(cfg.Database)

	userRepo := gorm.UserSetup(db)

	crudUserWrapper := goarch.SetupCrudUserWrapper(&cfg.GoarchGrpc)

	userSvc := userSvc.NewService(userRepo, crudUserWrapper)
	authSvc := authSvc.NewService(userRepo)

	return &Container{
		Config:   cfg,
		UserSvc:  userSvc,
		AuthSvc:  authSvc,
		Validate: validator.New(),
	}
}
