package main

// GetCustomerReq ...
type GetCustomerReq struct{}

// CustomerResp ...
type CustomerResp struct{}

// AddCustomerReq ...
type AddCustomerReq struct{}

// GetOrderReq ...
type GetOrderReq struct{}

// OrderResp ...
type OrderResp struct{}

// AddOrderReq ...
type AddOrderReq struct{}

// FndOrdersReq ...
type FndOrdersReq struct{}

// OrdersResp ...
type OrdersResp struct {
	Orders []*OrderResp
}
