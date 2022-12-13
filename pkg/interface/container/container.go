package container

import (
	"github.com/go-playground/validator"
	"goarch/pkg/infrastructure/goarch"
	"goarch/pkg/infrastructure/gorm"
	goarchApi "goarch/pkg/infrastructure/http/goarch"
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
	dbMaster, err := Database.New(cfg.Database.Master)
	if err != nil {
		panic(err)
	}
	dbSlave, _ := Database.New(cfg.Database.Slave)
	if dbSlave == nil {
		dbSlave = dbMaster
	}

	userRepo := gorm.UserSetup(dbMaster, dbSlave)

	crudUserWrapper := goarch.SetupCrudUserWrapper(&cfg.GoarchGrpc)
	goArchAPIWrapper := goarchApi.NewWrapper(&cfg.GoarchAPIConfig)

	userSvc := userSvc.NewService(userRepo, crudUserWrapper, goArchAPIWrapper)
	authSvc := authSvc.NewService(userRepo)

	return &Container{
		Config:   cfg,
		UserSvc:  userSvc,
		AuthSvc:  authSvc,
		Validate: validator.New(),
	}
}
