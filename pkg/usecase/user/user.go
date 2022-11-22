package user

import (
	"strings"
	"time"

	domainUser "goarch/pkg/domain/user"
	"goarch/pkg/infrastructure/goarch"
	goarchApi "goarch/pkg/infrastructure/http/goarch"
	"goarch/pkg/shared/constants"
	"goarch/pkg/shared/utils"
	ctxSess "goarch/pkg/shared/utils/context"
)

type service struct {
	userRepo         domainUser.Repository
	crudUserWrapper  goarch.CrudUserWrapper
	goarchApiWrapper goarchApi.Wrapper
}

func NewService(userRepo domainUser.Repository, crudUserWrapper goarch.CrudUserWrapper, goarchApiWrapper goarchApi.Wrapper) Service {
	s := &service{
		userRepo:         userRepo,
		crudUserWrapper:  crudUserWrapper,
		goarchApiWrapper: goarchApiWrapper,
	}
	if s.userRepo == nil {
		panic("please provide user repo")
	}
	if s.crudUserWrapper == nil {
		panic("please provide user grpc wrapper")
	}
	if s.goarchApiWrapper == nil {
		panic("please provide user API wrapper")
	}
	return s
}

func (s *service) RegisterUser(ctxSess *ctxSess.Context, req *RegisterRequest) (res *User, err error) {
	out, err := s.crudUserWrapper.RegisterUser(ctxSess, req.Email, req.Password, req.Name)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}

	res = &User{
		Email: strings.ToLower(out.Email),
		Name:  out.Name,
	}

	return
}

func (s *service) GetUserDetail(ctxSess *ctxSess.Context) (resp User, err error) {
	out, err := s.goarchApiWrapper.GetUserDetail(ctxSess, ctxSess.UserSession.AccountID)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorDataNotFound
		return
	}

	resp = User{
		Email: out.Email,
		Name:  out.Name,
	}

	return
}

func (s *service) GetUserDetailAPI(ctxSess *ctxSess.Context, userID int64) (resp User, err error) {
	entity, err := s.userRepo.FindById(userID)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorDataNotFound
		return
	}

	resp = User{
		Email: entity.Email,
		Name:  entity.Name,
	}

	return
}

func (s *service) RegisterUserGrpc(ctxSess *ctxSess.Context, req *RegisterRequest) (res *User, err error) {
	password := utils.HashAndSalt([]byte(req.Password))
	err = s.userRepo.Save(&domainUser.Entity{
		Email:    strings.ToLower(req.Email),
		Name:     req.Name,
		Password: password,
	})
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}

	res = &User{
		Email: strings.ToLower(req.Email),
		Name:  req.Name,
	}

	return
}

func (s *service) ResetPassword(ctxSess *ctxSess.Context, req *ResetPasswordReq) (err error) {
	entity, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorGeneral
		return
	}

	entity.Password = utils.HashAndSalt([]byte(req.Password))
	entity.UpdatedAt = time.Now()
	err = s.userRepo.Save(&entity)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorGeneral
		return
	}

	return
}

func (s *service) UpdateName(ctxSess *ctxSess.Context, req *UpdateNameReq) (err error) {
	entity, err := s.userRepo.FindById(ctxSess.UserSession.AccountID)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorGeneral
		return
	}

	entity.Name = req.Name
	entity.UpdatedAt = time.Now()
	err = s.userRepo.Save(&entity)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorGeneral
		return
	}

	return
}

func (s *service) UpdatePassword(ctxSess *ctxSess.Context, req *UpdatePasswordReq) (err error) {
	entity, err := s.userRepo.FindById(ctxSess.UserSession.AccountID)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorGeneral
		return
	}

	if err = s.checkPassword(ctxSess, entity.Password, req.OldPassword); err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorPasswordNotMatch
		return
	}

	entity.Password = utils.HashAndSalt([]byte(req.Password))
	entity.UpdatedAt = time.Now()
	err = s.userRepo.Save(&entity)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		err = constants.ErrorGeneral
		return
	}

	return
}
