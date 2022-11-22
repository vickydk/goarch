package auth

import (
	"goarch/pkg/shared/constants"
	ctxSess "goarch/pkg/shared/utils/context"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) checkPassword(ctxSess *ctxSess.Context, entityPass, inputPass string) (err error) {
	//check password
	if err = bcrypt.CompareHashAndPassword([]byte(entityPass), []byte(inputPass)); err != nil {
		ctxSess.ErrorMessage = "failed compare password: " + err.Error()
		err = constants.ErrorPasswordNotMatch
		return
	}
	return
}
