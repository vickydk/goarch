package context

import (
	"context"
	"errors"
	"time"

	Map "github.com/orcaman/concurrent-map"
	"github.com/spf13/cast"

	Logger "goarch/pkg/shared/logger"
	"goarch/pkg/shared/models"
)

type Context struct {
	Map                       Map.ConcurrentMap
	Logger                    Logger.Logger
	RequestTime               time.Time
	UserSession               models.AccountSession
	XRequestID, XAgent        string
	AppName, AppVersion, IP   string
	Port                      int
	SrcIP, URL, Method        string
	Header, Request, Response interface{}
	ErrorMessage              string
	ResponseCode              int
}

func New(logger Logger.Logger) *Context {
	return &Context{
		RequestTime: time.Now(),
		Logger:      logger,
		Map:         Map.New(),
		Header:      map[string]interface{}{},
		Request:     struct{}{},
	}
}

func (s *Context) SetXRequestID(xRequestID string) *Context {
	s.XRequestID = xRequestID
	return s
}

func (s *Context) SetXAgent(xAgent string) *Context {
	s.XAgent = xAgent
	return s
}

func (s *Context) SetMethod(method string) *Context {
	s.Method = method
	return s
}

func (s *Context) SetAppName(appName string) *Context {
	s.AppName = appName
	return s
}

func (s *Context) SetAppVersion(appVersion string) *Context {
	s.AppVersion = appVersion
	return s
}

func (s *Context) SetURL(url string) *Context {
	s.URL = url
	return s
}

func (s *Context) SetIP(ip string) *Context {
	s.IP = ip
	return s
}

func (s *Context) SetPort(port int) *Context {
	s.Port = port
	return s
}

func (s *Context) SetSrcIP(srcIp string) *Context {
	s.SrcIP = srcIp
	return s
}

func (s *Context) SetHeader(header interface{}) *Context {
	s.Header = header
	return s
}

func (s *Context) SetRequest(request interface{}) *Context {
	s.Request = request
	return s
}

func (s *Context) SetRequestTime(request time.Time) *Context {
	s.RequestTime = request
	return s
}

func (s *Context) SetErrorMessage(errorMessage string) *Context {
	s.ErrorMessage = errorMessage
	return s
}

func (s *Context) SetResponseCode(responseCode int) *Context {
	s.ResponseCode = responseCode
	return s
}

func (s *Context) Get(key string) (data interface{}, err error) {
	data, ok := s.Map.Get(key)
	if !ok {
		err = errors.New("not found")
	}
	return
}

func (s *Context) Put(key string, data interface{}) {
	s.Map.Set(key, data)
}

func (s *Context) Lv1(message ...interface{}) {
	s.Logger.Info(s.toContextLogger("Lv1"), "", formatLogs(message...)...)
}

func (s *Context) Lv2(message ...interface{}) time.Time {
	s.Logger.Info(s.toContextLogger("Lv2"), "", formatLogs(message...)...)
	return time.Now()
}

func (s *Context) Lv3(startProcessTime time.Time, message ...interface{}) {
	stop := time.Now()

	msg := formatLogs(message...)
	msg = append(msg, Logger.ToField("_process_time", stop.Sub(startProcessTime).Nanoseconds()/1000000))

	s.Logger.Info(s.toContextLogger("Lv3"), "", msg...)
}

func (s *Context) Lv4(message ...interface{}) {
	stop := time.Now()
	rt := stop.Sub(s.RequestTime).Nanoseconds() / 1000000

	msg := formatLogs(message...)
	msg = append(msg, Logger.ToField("_response_time", rt))
	msg = append(msg, Logger.ToField("_response_code", s.ResponseCode))

	s.Logger.Info(s.toContextLogger("Lv4"), "", msg...)
}

func formatLogs(message ...interface{}) (logRecord []Logger.Field) {
	for index, msg := range message {
		logRecord = append(logRecord, Logger.ToField("_message_"+cast.ToString(index), msg))
	}

	return
}

func (s *Context) toContextLogger(tag string) (ctx context.Context) {
	ctxVal := Logger.Context{
		ServiceName:    s.AppName,
		ServiceVersion: s.AppVersion,
		ServicePort:    s.Port,
		XRequestID:     s.XRequestID,
		XAgent:         s.XAgent,
		Tag:            tag,
		ReqMethod:      s.Method,
		ReqURI:         s.URL,
		AdditionalData: s.Map.Items(),
		Error:          s.ErrorMessage,
	}
	if tag == "Lv4" {
		ctxVal.Request = Logger.ToField("req", s.Request)
		ctxVal.Response = Logger.ToField("resp", s.Response)
	}

	ctx = Logger.InjectCtx(context.Background(), ctxVal)
	return
}
