package goarch

import (
	"fmt"
	"net/http"
	"os"

	"goarch/pkg/shared/config"
	"goarch/pkg/shared/rest"
	ctxSess "goarch/pkg/shared/utils/context"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/encoding/json"
)

type goarchAPI struct {
	config *config.GoarchAPIConfig
	client rest.RestClient
}

func NewWrapper(config *config.GoarchAPIConfig) Wrapper {
	if config == nil {
		panic("erp config is nil")
	}

	return &goarchAPI{
		config: config,
		client: rest.New(config.RestOptions),
	}
}

func (w *goarchAPI) GetUserDetail(ctx *ctxSess.Context, userId int64) (out GetUserDetailResponse, err error) {

	path := fmt.Sprintf(w.config.Path.GetUserDetail+"/%d", userId)

	headers := http.Header{}
	headers.Set(echo.HeaderXRequestID, ctx.XRequestID)

	body, _, httpErr := w.client.Get(path, headers)

	if os.IsTimeout(httpErr) {
		err = httpErr
		return
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return
	}

	return
}
