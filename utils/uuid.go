package utils

import (
	"context"

	"github.com/google/uuid"
)

type ContextTag string

const (
	LogIDTag ContextTag = "logid"
)

func NewLogID() string {
	return uuid.New().String()
}

func NewCtxWithLogID() context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, LogIDTag, NewLogID)
}
