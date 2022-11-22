package rest

import (
	"context"
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

type RestClient interface {
	Post(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PostWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PostFormData(path string, headers http.Header, payload map[string]string) (body []byte, statusCode int, err error)
	Put(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PutWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	Get(path string, headers http.Header) (body []byte, statusCode int, err error)
	GetWithContext(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error)
	GetWithQueryParam(path string, headers http.Header, queryParam map[string]string) (body []byte, statusCode int, err error)
	GetWithQueryParamAndContext(ctx context.Context, path string, headers http.Header, queryParam map[string]string) (body []byte, statusCode int, err error)
	Delete(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	DeleteWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	Patch(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
	PatchWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error)
}

type client struct {
	options    Options
	httpClient *resty.Client
}

func New(options Options) RestClient {
	httpClient := resty.New()

	//if options.SkipTLS {
	//	httpClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	//}

	if options.SkipCheckRedirect {
		httpClient.SetRedirectPolicy(resty.RedirectPolicyFunc(func(request *http.Request, requests []*http.Request) error {
			return http.ErrUseLastResponse
		}))
	}

	if options.WithProxy {
		httpClient.SetProxy(options.ProxyAddress)
	} else {
		httpClient.RemoveProxy()
	}

	httpClient.SetTimeout(options.Timeout * time.Second)
	httpClient.SetDebug(options.DebugMode)

	return &client{
		options:    options,
		httpClient: httpClient,
	}
}

func (c *client) Post(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetBody(payload)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "Astro-Admin")

	httpResp, httpErr := request.Post(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) PostWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)
	request.SetBody(payload)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "Astro-Admin")

	httpResp, httpErr := request.Post(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) PostFormData(path string, headers http.Header, payload map[string]string) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetFormData(payload)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	httpResp, httpErr := request.Post(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) Put(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	request.SetBody(payload)

	httpResp, httpErr := request.Put(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) PutWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	request.SetBody(payload)

	httpResp, httpErr := request.Put(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) Get(path string, headers http.Header) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()

	for h, val := range headers {
		request.Header[h] = val
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	httpResp, httpErr := request.Get(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) GetWithContext(ctx context.Context, path string, headers http.Header) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	httpResp, httpErr := request.Get(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) GetWithQueryParam(path string, headers http.Header, queryParam map[string]string) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()

	for h, val := range headers {
		request.Header[h] = val
	}
	request.Header.Set("User-Agent", "https://linkaja.id")
	request.SetQueryParams(queryParam)

	httpResp, httpErr := request.Get(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) GetWithQueryParamAndContext(ctx context.Context, path string, headers http.Header, queryParam map[string]string) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}
	request.Header.Set("User-Agent", "https://linkaja.id")
	request.SetQueryParams(queryParam)

	httpResp, httpErr := request.Get(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) Delete(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()

	for h, val := range headers {
		request.Header[h] = val
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	request.SetBody(payload)

	httpResp, httpErr := request.Delete(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) DeleteWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)

	for h, val := range headers {
		request.Header[h] = val
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	request.SetBody(payload)

	httpResp, httpErr := request.Delete(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) Patch(path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetBody(payload)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	httpResp, httpErr := request.Patch(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}

func (c *client) PatchWithContext(ctx context.Context, path string, headers http.Header, payload interface{}) (body []byte, statusCode int, err error) {
	url := c.options.Address + path

	request := c.httpClient.R()
	request.SetContext(ctx)
	request.SetBody(payload)

	for h, val := range headers {
		request.Header[h] = val
	}
	if headers["Content-Type"] == nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "https://linkaja.id")

	httpResp, httpErr := request.Patch(url)

	if httpResp != nil {
		body = httpResp.Body()
	}

	if httpResp != nil && httpResp.StatusCode() != 0 {
		statusCode = httpResp.StatusCode()
	}

	if statusCode == http.StatusOK {
		return body, statusCode, nil
	}

	return body, statusCode, httpErr
}
