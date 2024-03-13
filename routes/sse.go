package routes

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/http/sse"
	sse2 "github.com/qbhy/goal-piplin/app/http/sse"
)

func Sse(router contracts.HttpRouter) {
	// 自定义 sse 控制器
	router.Get(sse.New("/api/notify", sse2.Notify{}))

	router.Get("/send-sse", func(sse contracts.Sse, request contracts.HttpRequest) error {
		return sse.Send(uint64(request.GetInt64("fd")), request.GetString("msg"))
	})

	router.Get("/close-sse", func(sse contracts.Sse, request contracts.HttpRequest) error {
		return sse.Close(uint64(request.GetInt64("fd")))
	})
}
