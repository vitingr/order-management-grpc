package main

import "context"

type store struct {
	// add here our mongodb instance
}

func NewStore() *store {
	return &store{}
}

func (store *store) Create(context.Context) error {
	return nil
}