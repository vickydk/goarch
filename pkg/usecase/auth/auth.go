package auth

import (
	domainUser "goarch/pkg/domain/user"
	"goarch/pkg/shared/session"
	ctxSess "goarch/pkg/shared/utils/context"
)

type service struct {
	userRepo domainUser.Repository
}

func NewService(userRepo domainUser.Repository) Service {
	s := &service{
		userRepo: userRepo,
	}
	if s.userRepo == nil {
		panic("please provide user repo")
	}
	return s
}

func (s *service) Login(ctxSess *ctxSess.Context, req *LoginRequest) (res *LoginResponse, err error) {
	entity, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return
	}

	if err = s.checkPassword(ctxSess, entity.Password, req.Password); err != nil {
		ctxSess.ErrorMessage = err.Error()
		return nil, err
	}

	tokenString, _ := session.NewBearerToken(&entity)
	refreshToken, _ := session.RefreshToken(&entity)

	res = &LoginResponse{
		Id:    entity.ID,
		Name:  entity.Name,
		Email: entity.Email,
		Token: Token{
			AccessToken:  tokenString,
			RefreshToken: refreshToken,
		},
	}

	return
}
