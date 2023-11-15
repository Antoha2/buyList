package transport

import (
	"context"
	"net/http"

	buyService "github.com/antoha2/buylist/service"
)

type webImpl struct {
	Transport
	service buyService.UnitService
	server  *http.Server
}

func NewHTTP(buyService buyService.UnitService) *webImpl {
	return &webImpl{
		service: buyService,
	}
}

type Transport interface {
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
