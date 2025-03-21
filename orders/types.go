package main

import "context"

type OrdersService interface {
	createOrder(context.Context) error
}

type OrdersStore interface {
	create(context.Context)
}