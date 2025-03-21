package main

import "context"

type Store struct {
}

func NewStore() *Store {
	return &Store{}
}

func (store *Store) create(context.Context) {}