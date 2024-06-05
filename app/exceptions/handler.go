package exceptions

import (
	"github.com/goal-web/auth"
	"github.com/goal-web/contracts"
	"github.com/goal-web/http"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"github.com/goal-web/validation"
	"reflect"
	"runtime/debug"
	"strings"
)

type ExceptionHandler struct {
	dontReportExceptions []reflect.Type
}

func NewHandler() contracts.ExceptionHandler {
	return &ExceptionHandler{utils.ToTypes([]contracts.Exception{})}
}

func (handler *ExceptionHandler) Handle(exception contracts.Exception) any {
	logs.WithException(exception).Warn("报错了")
	switch e := exception.(type) {
	case http.Exception: // http 支持在异常处理器返回响应
		return handler.handleHttpException(e)
	case auth.Exception: // http 支持在异常处理器返回响应
		return http.NewJsonResponse(contracts.Fields{
			"msg":  "未登录",
			"code": 401,
		}, 401)
	case *validation.Exception:
		return handler.renderValidationException(e)
	default:
		debug.PrintStack()
	}

	logs.WithException(exception).
		WithField("exception", reflect.TypeOf(exception).String()).
		Error("ExceptionHandler")

	if httpException, isHttpException := exception.(http.Exception); isHttpException {
		logs.WithException(httpException).WithFields(contracts.Fields{}).Debug("http请求报错")
	}

	if handler.ShouldReport(exception) {
		handler.Report(exception)
	}

	return contracts.Fields{
		"msg": exception.Error(),
	}
}

func (handler *ExceptionHandler) handleHttpException(exception http.Exception) any {
	switch e := exception.Exception.(type) {
	case *validation.Exception:
		return handler.renderValidationException(e)
	default:
		if !strings.Contains(exception.Error(), "404") {
			debug.PrintStack()
		}
		return contracts.Fields{
			"path": exception.Request.Path(),
			"msg":  e.Error(),
		}
	}
}

func (handler *ExceptionHandler) renderValidationException(exception *validation.Exception) any {
	return contracts.Fields{
		"msg":    exception.Error(),
		"fields": exception.Param,
		"errors": exception.Errors,
	}
}

func (handler *ExceptionHandler) Report(exception contracts.Exception) {
}
func (handler *ExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	return !utils.IsInstanceIn(exception, handler.dontReportExceptions...)
}
