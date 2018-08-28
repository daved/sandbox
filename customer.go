package main

import (
	"context"
	"net/http"

	"github.com/codemodus/chain"
	"github.com/codemodus/mixmux"
)

type customerDataProvider interface {
	SelectCustomer(context.Context, *GetCustomerReq) (*CustomerResp, error)
	InsertCustomer(context.Context, *AddCustomerReq) (*CustomerResp, error)
}

type customerService struct {
	db customerDataProvider
}

func newCustomerService(db customerDataProvider) (*customerService, error) {
	svc := &customerService{
		db: db,
	}

	return svc, nil
}

func (svc *customerService) applyRoutes(m mixmux.Mux) error {
	c := chain.New()

	m.Get("/api/v1/customer", c.EndFn(svc.HandleGetCustomer))
	m.Post("/api/v1/customer", c.EndFn(svc.HandlePostCustomer))

	return nil
}

func (svc *customerService) HandleGetCustomer(w http.ResponseWriter, r *http.Request) {
	// ...
}

func (svc *customerService) HandlePostCustomer(w http.ResponseWriter, r *http.Request) {
	// ...
}
