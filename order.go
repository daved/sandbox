package main

import (
	"net/http"

	"github.com/codemodus/chain"
	"github.com/codemodus/mixmux"
)

type orderDataProvider interface{}

type orderService struct {
	db orderDataProvider
}

func newOrderService(db orderDataProvider) (*orderService, error) {
	svc := &orderService{
		db: db,
	}

	return svc, nil
}

func (svc *orderService) applyRoutes(m mixmux.Mux) error {
	c := chain.New()

	m.Get("/api/v1/order", c.EndFn(svc.HandleGetOrder))
	m.Post("/api/v1/order", c.EndFn(svc.HandlePostOrder))

	return nil
}

func (svc *orderService) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	// ...
}

func (svc *orderService) HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	// ...
}
