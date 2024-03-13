package sse

import (
	"github.com/goal-web/contracts"
)

type Notify struct {
}

func (c Notify) OnConnect(request contracts.HttpRequest, fd uint64) error {

	return nil
}

func (c Notify) OnClose(fd uint64) {
	// todo: 实现解绑
}
